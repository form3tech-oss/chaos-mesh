// Copyright 2023 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package certificatechaos

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/avast/retry-go/v4"
	cmv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	"github.com/go-logr/logr"
	"go.uber.org/fx"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	impltypes "github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/types"
	"github.com/chaos-mesh/chaos-mesh/controllers/utils/controller"
)

var _ impltypes.ChaosImpl = (*Impl)(nil)

const FluxSuspended = "Not Injected/FluxSuspended"
const CertUpdated = "Not Injected/CertUpdated"
const CertReady = "Not Injected/CertReady"
const RevertedCerts = "Injected/CertUpdated"

const restartTimeAnnotation = "chaos-mesh.org/certificateChaosAt"

type Impl struct {
	client.Client
	Log logr.Logger
}

func (impl *Impl) Apply(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	impl.Log.Info("certificate chaos Apply", "namespace", obj.GetNamespace(), "name", obj.GetName())

	chaos, ok := obj.(*v1alpha1.CertificateChaos)
	if !ok {
		err := errors.New("chaos is not CertificateChaos")
		impl.Log.Error(err, "chaos is not CertificateChaos", "chaos", obj)
		return v1alpha1.NotInjected, err
	}

	if chaos.Status.Instances == nil {
		chaos.Status.Instances = make(map[string]v1alpha1.Instance)
	}

	record := records[index]
	namespacedName, err := controller.ParseNamespacedName(record.Id)
	if err != nil {
		return v1alpha1.NotInjected, err
	}

	switch record.Phase {
	case v1alpha1.NotInjected:
		var cert cmv1.Certificate
		err = impl.Get(ctx, namespacedName, &cert)
		if err != nil {
			if apiErrors.IsNotFound(err) {
				return v1alpha1.Injected, nil
			}
			return v1alpha1.NotInjected, err
		}

		// Find and suspend Flux resources
		if entity, ok := getManagedBy(&cert); ok {
			impl.Log.Info("Suspending Flux", "resource", entity.NamespacedName())
			if err = impl.suspend(ctx, entity, true); err != nil {
				innerErr := errors.New("failed to suspend Flux resource")
				impl.Log.Error(err, "failed to suspend Flux resource", "resource", entity)
				return v1alpha1.NotInjected, innerErr
			}

			chaos.Status.Instances[record.Id] = v1alpha1.Instance{FluxResource: entity}
		}
		return FluxSuspended, nil

	case FluxSuspended:
		var cert cmv1.Certificate
		err = impl.Get(ctx, namespacedName, &cert)
		if err != nil {
			if apiErrors.IsNotFound(err) {
				return v1alpha1.Injected, nil
			}
			return v1alpha1.NotInjected, err
		}

		// Update actual certificate
		if err = impl.updateCertificate(ctx, &cert, chaos.Spec.CertificateExpiry, chaos.Spec.RenewBefore); err != nil {
			impl.Log.Error(err, "Updating Certificate", "resource", cert.Name)
			return record.Phase, err
		}
		newInstance := chaos.Status.Instances[record.Id]
		newInstance.OriginalExpiry = cert.Spec.Duration
		newInstance.OriginalRenewBefore = cert.Spec.RenewBefore
		newInstance.SecretName = cert.Spec.SecretName
		chaos.Status.Instances[record.Id] = newInstance
		return CertUpdated, nil

	case CertUpdated:
		impl.Log.Info("Checking if Certificate is ready", "certificate", namespacedName.String())

		instance := chaos.Status.Instances[record.Id]
		var readyAt metav1.Time
		if readyAt, ok = impl.getCertificateReadyAt(ctx, namespacedName, chaos); !ok {
			return CertUpdated, errors.New("certificate not yet ready")
		}

		instance.CertificateReadyAt = &readyAt
		chaos.Status.Instances[record.Id] = instance
		return CertReady, nil

	case CertReady:
		impl.Log.Info("Finding related PODs", "certificate", namespacedName.String())
		listOptions := &client.ListOptions{
			Namespace: namespacedName.Namespace,
		}
		podsList := &v1.PodList{}
		err := impl.Client.List(ctx, podsList, listOptions)
		if err != nil {
			impl.Log.Error(err, "Finding related PODs failed", "certificate", namespacedName.String())
			return CertReady, err
		}
		instance := chaos.Status.Instances[record.Id]

		owners, err := impl.getPodOwnersUsingSecret(ctx, podsList, instance.SecretName)
		if err != nil {
			return record.Phase, err
		}

		for owner := range owners {
			err = impl.Restart(ctx, owner, namespacedName.Namespace, instance.CertificateReadyAt)
			if err != nil {
				return record.Phase, fmt.Errorf("restarting %s/%s: %w", owner.Kind, owner.Name, err)
			}
		}

	default:
		panic("unknown phase: " + record.Phase)
	}

	return v1alpha1.Injected, nil
}

func (impl *Impl) getPodOwnersUsingSecret(ctx context.Context, podsList *v1.PodList, secretName string) (map[Dependent]bool, error) {
	owners := make(map[Dependent]bool)
	for _, pod := range podsList.Items {
		if usesVolume(pod, secretName) {
			for _, ref := range pod.GetOwnerReferences() {
				if ref.APIVersion == "apps/v1" && ref.Kind == "ReplicaSet" {
					// need to find the owner ref of the replicaset
					var rs appsv1.ReplicaSet
					rsNamespacedName := types.NamespacedName{Name: ref.Name, Namespace: pod.Namespace}
					err := impl.Client.Get(ctx, rsNamespacedName, &rs)
					if err != nil {
						return nil, fmt.Errorf("getting replicaset %s/%s: %w", pod.Namespace, ref.Name, err)
					}

					for _, ref := range rs.GetOwnerReferences() {
						owners[Dependent{
							APIVersion: ref.APIVersion,
							Kind:       ref.Kind,
							Name:       ref.Name,
						}] = true
					}
				} else {
					owners[Dependent{
						APIVersion: ref.APIVersion,
						Kind:       ref.Kind,
						Name:       ref.Name,
					}] = true
				}
			}
		}

	}
	return owners, nil
}

func (impl *Impl) getCertificateReadyAt(ctx context.Context, namespacedName types.NamespacedName, chaos *v1alpha1.CertificateChaos) (metav1.Time, bool) {
	var timestamp metav1.Time
	err := retry.Do(
		func() error {
			var cert cmv1.Certificate
			if innerErr := impl.Get(ctx, namespacedName, &cert); innerErr != nil {
				return innerErr
			}
			for _, cond := range cert.Status.Conditions {
				if cond.Reason == "Ready" && cond.Status == "True" {
					// TODO what if certificate was already at the specified duration? Then it won't update
					if cond.LastTransitionTime.After(chaos.CreationTimestamp.Time.Add(-3 * time.Second)) {
						timestamp = metav1.NewTime((*cond.LastTransitionTime).Time)
						return nil
					}
					break
				}
			}

			return errors.New("certificate not yet ready")
		},
		retry.Attempts(3),
		retry.OnRetry(func(n uint, err error) {
			impl.Log.Info("Certificate not yet ready", "certificate", namespacedName.String())
		}),
		retry.Delay(time.Second),
	)
	return timestamp, err == nil
}

func usesVolume(pod v1.Pod, secretName string) bool {
	for _, volume := range pod.Spec.Volumes {
		if volume.Secret != nil && volume.Secret.SecretName == secretName {
			return true
		}
		if volume.Projected != nil {
			for _, source := range volume.Projected.Sources {
				if source.Secret != nil && source.Secret.Name == secretName {
					return true
				}
			}
		}
	}
	return false
}

func (impl *Impl) updateCertificate(ctx context.Context, cert *cmv1.Certificate, expiry, renewBefore *metav1.Duration) error {
	updated := cert.DeepCopy()

	updated.Spec.Duration.Duration = expiry.Duration
	updated.Spec.RenewBefore.Duration = renewBefore.Duration

	impl.Log.Info(
		"Patching certificate",
		"namespace", cert.Namespace,
		"name", cert.Name,
		"certificateExpiry", *expiry,
		"renewBefore", *renewBefore,
	)
	err := impl.Patch(ctx, updated, client.MergeFrom(cert))
	if err != nil {
		return fmt.Errorf("patching certificate %s/%s: %w", cert.Namespace, cert.Name, err)
	}

	return nil
}

func getManagedBy(cert *cmv1.Certificate) (v1alpha1.FluxResource, bool) {
	entity := v1alpha1.FluxResource{}
	for k, v := range cert.GetLabels() {
		// TODO versions should probably not be hardcoded. We can probably do that later on though
		// One idea would be to query the api-resources for it?
		switch k {
		case "helm.toolkit.fluxcd.io/namespace":
			entity.Namespace = v
			entity.Group = "helm.toolkit.fluxcd.io"
			entity.Version = "v2"
			entity.Kind = "helmrelease"
		case "helm.toolkit.fluxcd.io/name":
			entity.Name = v
			entity.Group = "helm.toolkit.fluxcd.io"
			entity.Version = "v2"
			entity.Kind = "helmrelease"
		case "kustomize.toolkit.fluxcd.io/namespace":
			entity.Namespace = v
			entity.Group = "kustomize.toolkit.fluxcd.io"
			entity.Version = "v1"
			entity.Kind = "kustomization"
		case "kustomize.toolkit.fluxcd.io/name":
			entity.Name = v
			entity.Group = "kustomize.toolkit.fluxcd.io"
			entity.Version = "v1"
			entity.Kind = "kustomization"
		}
	}

	if entity.Name != "" && entity.Namespace != "" {
		return entity, true
	}
	return v1alpha1.FluxResource{}, false
}

func (impl *Impl) suspend(ctx context.Context, e v1alpha1.FluxResource, state bool) error {
	var nilFluxResource v1alpha1.FluxResource
	if e == nilFluxResource {
		return nil
	}
	u := &unstructured.Unstructured{}
	u.Object = map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":      e.Name,
			"namespace": e.Namespace,
		},
	}
	u.SetGroupVersionKind(e.GVK())
	patch := []byte(fmt.Sprintf(`{"spec":{"suspend":%v}}`, state))
	return impl.Client.Patch(ctx, u, client.RawPatch(types.MergePatchType, patch))
}

func (impl *Impl) Recover(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	impl.Log.Info("certificate chaos Recover", "namespace", obj.GetNamespace(), "name", obj.GetName())

	chaos, ok := obj.(*v1alpha1.CertificateChaos)
	if !ok {
		err := errors.New("chaos is not CertificateChaos")
		impl.Log.Error(err, "chaos is not CertificateChaos", "chaos", obj)
		return v1alpha1.Injected, err
	}

	if chaos.Status.Instances == nil {
		impl.Log.Info("No Instances to recover")
		return v1alpha1.NotInjected, nil
	}

	record := records[index]
	if instance, ok := chaos.Status.Instances[record.Id]; ok {
		switch record.Phase {
		case v1alpha1.Injected:
			namespacedName, err := controller.ParseNamespacedName(record.Id)
			if err != nil {
				return v1alpha1.Injected, err
			}

			var cert cmv1.Certificate
			err = impl.Get(ctx, namespacedName, &cert)
			if err != nil {
				if apiErrors.IsNotFound(err) {
					return v1alpha1.Injected, nil
				}
				return v1alpha1.Injected, err
			}
			if err = impl.updateCertificate(ctx, &cert, instance.OriginalExpiry, instance.OriginalRenewBefore); err != nil {
				impl.Log.Error(err, "Updating Certificate", "resource", cert.Name)
				return record.Phase, err
			}

			return RevertedCerts, nil

		case RevertedCerts:
			err := impl.suspend(ctx, instance.FluxResource, false)
			if err != nil {
				innerErr := errors.New("failed to unsuspend Flux resource")
				impl.Log.Error(err, "failed to unsuspend Flux resource", "resource", instance.FluxResource)
				return v1alpha1.Injected, innerErr
			}

		default:
			panic("unknown recovery phase: " + record.Phase)
		}

	}

	return v1alpha1.NotInjected, nil
}

type Dependent struct {
	APIVersion string
	Kind       string
	Name       string
}

func (owner Dependent) toUnstructured(namespace string) unstructured.Unstructured {
	u := unstructured.Unstructured{}
	u.SetAPIVersion(owner.APIVersion)
	u.SetKind(owner.Kind)
	u.SetName(owner.Name)
	u.SetNamespace(namespace)
	return u

}

func (impl *Impl) Restart(ctx context.Context, owner Dependent, namespace string, certificateAt *metav1.Time) error {
	resourceType := strings.ToLower(owner.Kind)
	if resourceType != string(v1alpha1.DaemonSetResourceType) &&
		resourceType != string(v1alpha1.DeploymentResourceType) &&
		resourceType != string(v1alpha1.StatefulSetResourceType) {
		impl.Log.Info("Can't restart this resource", "resource", owner.Kind, "name", owner.Name, "namespace", namespace)
		return nil
	}

	u := owner.toUnstructured(namespace)
	_ = impl.Client.Get(context.Background(), client.ObjectKey{
		Namespace: u.GetNamespace(),
		Name:      u.GetName(),
	}, &u)
	if annotation, ok := u.GetAnnotations()[restartTimeAnnotation]; ok {
		if restartedAt, err := time.Parse(time.RFC3339, annotation); err == nil && restartedAt.After(certificateAt.Time) {
			impl.Log.Info("Skipping restart. Already happened", "resource", owner.Kind,
				"name", owner.Name,
				"namespace", namespace,
				"restartedAt", restartedAt.String(),
				"certificateAt", certificateAt.String())
			return nil
		}
	}

	u = owner.toUnstructured(namespace)

	now := time.Now().UTC()
	data := []byte(fmt.Sprintf(
		`{
				"spec": {"template": {"metadata": {"annotations": {"kubectl.kubernetes.io/restartedAt": "%s"}}}},
			 	"metadata": {"annotations": {"%s": "%s"}}
			}`,
		now.Format("20060102150405"),
		restartTimeAnnotation,
		now.Format(time.RFC3339),
	))

	impl.Log.Info("Patching", "resource", owner.Kind, "name", owner.Name, "namespace", namespace)
	return impl.Client.Patch(ctx, &u, client.RawPatch(types.StrategicMergePatchType, data))
}

func NewImpl(c client.Client, log logr.Logger) *impltypes.ChaosImplPair {
	return &impltypes.ChaosImplPair{
		Name:   "certificatechaos",
		Object: &v1alpha1.CertificateChaos{},
		Impl: &Impl{
			Client: c,
			Log:    log.WithName("certificatechaos"),
		},
		ObjectList: &v1alpha1.CertificateChaosList{},
	}
}

var Module = fx.Provide(
	fx.Annotated{
		Group:  "impl",
		Target: NewImpl,
	},
)
