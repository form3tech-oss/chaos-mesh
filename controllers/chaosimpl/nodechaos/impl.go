// Copyright 2021 Chaos Mesh Authors.
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
//

package nodechaos

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"go.uber.org/fx"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	impltypes "github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/types"
	"github.com/chaos-mesh/chaos-mesh/controllers/utils/chaosdaemon"
	"github.com/chaos-mesh/chaos-mesh/pkg/chaosdaemon/pb"
)

var _ impltypes.ChaosImpl = (*Impl)(nil)

type Impl struct {
	client.Reader
	Log     logr.Logger
	Builder *chaosdaemon.ChaosDaemonClientBuilder
}

func (impl *Impl) Apply(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	record := records[index]

	ciliumContainerID, err := impl.ciliumContainerID(ctx, record.Id)
	if err != nil {
		return v1alpha1.NotInjected, fmt.Errorf("determing cilium-agent container ID: %w", err)
	}

	client, err := impl.Builder.BuildNodeClient(ctx, record.Id)
	if err != nil {
		return v1alpha1.NotInjected, fmt.Errorf("building chaos-daemon client: %w", err)
	}

	_, err = client.ApplyNodeChaos(ctx, &pb.ApplyNodeChaosRequest{ContainerId: ciliumContainerID})
	if err != nil {
		return v1alpha1.NotInjected, fmt.Errorf("applying node chaos: %w", err)
	}

	return v1alpha1.Injected, nil
}

func (impl *Impl) Recover(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	record := records[index]

	ciliumContainerID, err := impl.ciliumContainerID(ctx, record.Id)
	if err != nil {
		return v1alpha1.Injected, fmt.Errorf("determing cilium-agent container ID: %w", err)
	}

	client, err := impl.Builder.BuildNodeClient(ctx, record.Id)
	if err != nil {
		return v1alpha1.Injected, fmt.Errorf("building chaos-daemon client: %w", err)
	}

	_, err = client.RecoverNodeChaos(ctx, &pb.RecoverNodeChaosRequest{ContainerId: ciliumContainerID})
	if err != nil {
		return v1alpha1.Injected, fmt.Errorf("applying node chaos: %w", err)
	}

	return v1alpha1.NotInjected, nil
}

func (impl *Impl) ciliumContainerID(ctx context.Context, nodeName string) (string, error) {
	ciliumDS := &appsv1.DaemonSet{}
	err := impl.Get(ctx, types.NamespacedName{Name: "cilium", Namespace: "kube-system"}, ciliumDS)
	if err != nil {
		return "", fmt.Errorf("getting cilium daemonset: %w", err)
	}

	labelSelector := labels.SelectorFromSet(labels.Set(ciliumDS.Spec.Selector.MatchLabels))

	podList := &corev1.PodList{}
	err = impl.List(ctx, podList, &client.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return "", fmt.Errorf("listing cilium pods: %w", err)
	}

	pods := []corev1.Pod{}
	for _, pod := range podList.Items {
		if pod.Spec.NodeName == nodeName {
			pods = append(pods, pod)
		}
	}

	if len(pods) != 1 {
		return "", fmt.Errorf("received unexpected number of cilium pods: %d", len(podList.Items))
	}

	ciliumPod := pods[0]

	var ciliumContainerID string
	for _, status := range ciliumPod.Status.ContainerStatuses {
		if status.Name == "cilium-agent" {
			ciliumContainerID = status.ContainerID
			break
		}
	}

	if ciliumContainerID == "" {
		return "", fmt.Errorf("retrieving cilium-agent container id")
	}

	return ciliumContainerID, nil
}

type ImplParams struct {
	fx.In

	Reader  client.Reader `name:"no-cache"`
	Builder *chaosdaemon.ChaosDaemonClientBuilder
	Logger  logr.Logger
}

func NewImpl(params ImplParams) *impltypes.ChaosImplPair {
	return &impltypes.ChaosImplPair{
		Name:   "nodechaos",
		Object: &v1alpha1.NodeChaos{},
		Impl: &Impl{
			Reader:  params.Reader,
			Log:     params.Logger.WithName("nodechaos"),
			Builder: params.Builder,
		},
	}
}

var Module = fx.Provide(
	fx.Annotated{
		Group:  "impl",
		Target: NewImpl,
	},
	NewImpl,
)
