// Copyright 2022 Chaos Mesh Authors.
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

package gkenodepoolchaos

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-logr/logr"
	"go.uber.org/fx"
	container "google.golang.org/api/container/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	impltypes "github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/types"
	"github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/utils"
)

var _ impltypes.ChaosImpl = (*Impl)(nil)

type NodePool struct {
	Name        string                        `json:"name,omitempty"`
	Autoscaling container.NodePoolAutoscaling `json:"autoscaling,omitempty"`
}

type Impl struct {
	client.Client
	Log logr.Logger
}

func (impl *Impl) Apply(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	gkenodepoolchaos := obj.(*v1alpha1.GKENodePoolChaos)
	project := gkenodepoolchaos.Spec.Project
	cluster := gkenodepoolchaos.Spec.Cluster
	location := gkenodepoolchaos.Spec.Location

	var nodepool NodePool
	if err := json.Unmarshal([]byte(records[index].Id), &nodepool); err != nil {
		return v1alpha1.NotInjected, err
	}

	impl.Log.Info(fmt.Sprintf("disabling autoscaling for nodepool %q", nodepool.Name))

	nodepool.Autoscaling.Enabled = false
	nodepool.Autoscaling.MinNodeCount = 0
	nodepool.Autoscaling.MaxNodeCount = 0

	if err := updateNodePoolAutoscaling(ctx, impl.Log, project, cluster, location, nodepool); err != nil {
		return v1alpha1.NotInjected, err
	}

	return v1alpha1.Injected, nil
}

func (impl *Impl) Recover(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	gkenodepoolchaos := obj.(*v1alpha1.GKENodePoolChaos)
	project := gkenodepoolchaos.Spec.Project
	cluster := gkenodepoolchaos.Spec.Cluster
	location := gkenodepoolchaos.Spec.Location

	var nodepool NodePool
	if err := json.Unmarshal([]byte(records[index].Id), &nodepool); err != nil {
		return v1alpha1.NotInjected, err
	}

	impl.Log.Info(fmt.Sprintf("enabling autoscaling for nodepool %q", nodepool.Name))

	if err := updateNodePoolAutoscaling(ctx, impl.Log, project, cluster, location, nodepool); err != nil {
		return v1alpha1.NotInjected, err
	}

	return v1alpha1.NotInjected, nil
}

func updateNodePoolAutoscaling(ctx context.Context, log logr.Logger, project string, cluster string, location string, nodepool NodePool) error {
	log.Info("updating nodepool autoscaler",
		"project", project,
		"cluster", cluster,
		"location", location,
		"nodepool", nodepool.Name,
		"autoscaling", nodepool.Autoscaling)

	containerSvc, err := container.NewService(ctx)
	if err != nil {
		return fmt.Errorf("error creating container service: %w", err)
	}

	nodepoolSvc := container.NewProjectsLocationsClustersNodePoolsService(containerSvc)

	name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", project, location, cluster, nodepool.Name)

	req := &container.SetNodePoolAutoscalingRequest{
		Name:        name,
		Autoscaling: &nodepool.Autoscaling,
	}
	if _, err := nodepoolSvc.SetAutoscaling(req.Name, req).Do(); err != nil {
		return fmt.Errorf("error updating node pool autoscaler: %w", err)
	}

	return nil
}

// NewImpl returns a new GKENodePoolChaos implementation instance.
func NewImpl(c client.Client, log logr.Logger, decoder *utils.ContainerRecordDecoder) *impltypes.ChaosImplPair {
	return &impltypes.ChaosImplPair{
		Name:   "gkenodepoolchaos",
		Object: &v1alpha1.GKENodePoolChaos{},
		Impl: &Impl{
			Client: c,
			Log:    log.WithName("gkenodepoolchaos"),
		},
		ObjectList: &v1alpha1.GKENodePoolChaosList{},
	}
}

var Module = fx.Provide(
	fx.Annotated{
		Group:  "impl",
		Target: NewImpl,
	},
)
