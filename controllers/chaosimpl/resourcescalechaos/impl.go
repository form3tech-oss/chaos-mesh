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

package resourcescalechaos

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-logr/logr"
	"go.uber.org/fx"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	impltypes "github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/types"
)

var _ impltypes.ChaosImpl = (*Impl)(nil)

// ScalableResource is an interface for resource structs allowing to get and update scaling of the resource
type ScalableResource interface {
	GetScale(ctx context.Context, deploymentName string, options metav1.GetOptions) (*autoscalingv1.Scale, error)
	UpdateScale(ctx context.Context, deploymentName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (*autoscalingv1.Scale, error)
}

type ResourceSpecs struct {
	Namespace       string                `json:"namespace,omitempty"`
	Name            string                `json:"name,omitempty"`
	ResourceType    v1alpha1.ResourceType `json:"resourceType,omitempty"`
	ApplyReplicas   int                   `json:"applyReplicas,omitempty"`
	RecoverReplicas int                   `json:"recoverReplicas,omitempty"`
}

type Impl struct {
	client.Client
	Log logr.Logger
}

func (impl *Impl) Apply(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	impl.Log.Info("resourcescalechaos Apply", "namespace", obj.GetNamespace(), "name", obj.GetName())

	var spec ResourceSpecs
	if err := json.Unmarshal([]byte(records[index].Id), &spec); err != nil {
		return v1alpha1.NotInjected, err
	}

	client, err := impl.kubernetesClient()
	if err != nil {
		return v1alpha1.NotInjected, fmt.Errorf("failed to create kubernetes client: %w", err)
	}

	switch spec.ResourceType {
	case v1alpha1.ResourceTypeDeployment:
	case v1alpha1.ResourceTypeReplicaSet:
	case v1alpha1.ResourceTypeStatefulSet:
		res, err := impl.getScalingResource(client, spec)
		if err != nil {
			return v1alpha1.NotInjected, fmt.Errorf("failed to get scaling resource %s: %w", spec.ResourceType, err)
		}

		if err = impl.scaleResource(ctx, res, spec.Name, spec.ApplyReplicas); err != nil {
			return v1alpha1.NotInjected, fmt.Errorf("failed to scale resource %s: %w", spec.ResourceType, err)
		}
	case v1alpha1.ResourceTypeDaemonSet:
		return v1alpha1.NotInjected, fmt.Errorf("not yet implemented")
	}

	return v1alpha1.Injected, nil
}

func (impl *Impl) Recover(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	impl.Log.Info("resourcescalechaos Recover", "namespace", obj.GetNamespace(), "name", obj.GetName())

	var spec ResourceSpecs
	if err := json.Unmarshal([]byte(records[index].Id), &spec); err != nil {
		return v1alpha1.Injected, err
	}

	client, err := impl.kubernetesClient()
	if err != nil {
		return v1alpha1.Injected, fmt.Errorf("failed to create kubernetes client: %w", err)
	}

	res, err := impl.getScalingResource(client, spec)
	if err != nil {
		return v1alpha1.Injected, fmt.Errorf("failed to get scaling resource %s: %w", spec.ResourceType, err)
	}

	if err = impl.scaleResource(ctx, res, spec.Name, spec.RecoverReplicas); err != nil {
		return v1alpha1.Injected, fmt.Errorf("failed to scale resource %s: %w", spec.ResourceType, err)
	}

	return v1alpha1.NotInjected, nil
}

func (i *Impl) scaleResource(ctx context.Context, res ScalableResource, resourceName string, desiredScale int) error {
	scale, err := res.GetScale(ctx, resourceName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get scale %s: %w", resourceName, err)
	}

	scale.Spec.Replicas = int32(desiredScale)

	_, err = res.UpdateScale(ctx, resourceName, scale, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to update scale %s: %w", resourceName, err)
	}
}

func (i *Impl) getScalingResource(client *kubernetes.Clientset, spec ResourceSpecs) (ScalableResource, error) {
	switch spec.ResourceType {
	case v1alpha1.ResourceTypeDeployment:
		return client.AppsV1().Deployments(spec.Namespace), nil
	case v1alpha1.ResourceTypeReplicaSet:
		return client.AppsV1().ReplicaSets(spec.Namespace), nil
	case v1alpha1.ResourceTypeStatefulSet:
		return client.AppsV1().StatefulSets(spec.Namespace), nil
	}

	return nil, fmt.Errorf("failed to get scaling resource client for resource: %s", spec.ResourceType)
}

func (impl *Impl) kubernetesClient() (*kubernetes.Clientset, error) {
	config, err := ctrl.GetConfig()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func NewImpl(c client.Client, log logr.Logger) *impltypes.ChaosImplPair {
	return &impltypes.ChaosImplPair{
		Name:   "resourcescalechaos",
		Object: &v1alpha1.DeploymentChaos{},
		Impl: &Impl{
			Client: c,
			Log:    log.WithName("resourcescalechaos"),
		},
		ObjectList: &v1alpha1.DeploymentChaosList{},
	}
}

var Module = fx.Provide(
	fx.Annotated{
		Group:  "impl",
		Target: NewImpl,
	},
)
