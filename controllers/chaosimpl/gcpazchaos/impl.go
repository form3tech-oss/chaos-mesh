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

package gcpazchaos

import (
	"context"
	"encoding/json"

	"github.com/go-logr/logr"
	"go.uber.org/fx"
	compute "google.golang.org/api/compute/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	impltypes "github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/types"
)

var _ impltypes.ChaosImpl = (*Impl)(nil)

type InstanceGroup struct {
	Project string `json:"project,omitempty"`
	Zone    string `json:"zone,omitempty"`
	Name    string `json:"name,omitempty"`
	Size    int64  `json:"size,omitempty"`
}

type Impl struct {
	client.Client
	Log logr.Logger
}

func (impl *Impl) Apply(ctx context.Context, index int, records []*v1alpha1.Record, chaos v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	var ig InstanceGroup
	err := json.Unmarshal([]byte(records[index].Id), &ig)
	if err != nil {
		impl.Log.Error(err, "fail to decode instance group json")
		return v1alpha1.NotInjected, err
	}

	computeService, err := compute.NewService(ctx)
	if err != nil {
		impl.Log.Error(err, "fail to get the compute service")
		return v1alpha1.NotInjected, err
	}

	err = resizeInstanceGroup(computeService, impl.Log, ig.Project, ig.Zone, ig.Name, 0)
	if err != nil {
		return v1alpha1.NotInjected, err
	}

	return v1alpha1.Injected, nil
}

func (impl *Impl) Recover(ctx context.Context, index int, records []*v1alpha1.Record, chaos v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	var ig InstanceGroup
	err := json.Unmarshal([]byte(records[index].Id), &ig)
	if err != nil {
		impl.Log.Error(err, "fail to decode instance group json")
		return v1alpha1.Injected, err
	}

	computeService, err := compute.NewService(ctx)
	if err != nil {
		impl.Log.Error(err, "fail to get the compute service")
		return v1alpha1.Injected, err
	}

	err = resizeInstanceGroup(computeService, impl.Log, ig.Project, ig.Zone, ig.Name, ig.Size)
	if err != nil {
		return v1alpha1.Injected, err
	}

	return v1alpha1.NotInjected, nil
}

func resizeInstanceGroup(svc *compute.Service, logger logr.Logger, project string, zone string, name string, size int64) error {
	logger.Info("resizing instance group", "project", project, "zone", zone, "name", name, "size", size)
	if _, err := svc.InstanceGroupManagers.Resize(project, zone, name, size).Do(); err != nil {
		logger.Error(err, "fail to resize instance groups", "project", project, "zone", zone, "name", name, "size", size)
		return err
	}
	return nil
}

func NewImpl(c client.Client, log logr.Logger) *impltypes.ChaosImplPair {
	return &impltypes.ChaosImplPair{
		Name:   "gcpazchaos",
		Object: &v1alpha1.GCPAzChaos{},
		Impl: &Impl{
			Client: c,
			Log:    log.WithName("gcpazchaos"),
		},
		ObjectList: &v1alpha1.GCPAzChaosList{},
	}
}

var Module = fx.Provide(
	fx.Annotated{
		Group:  "impl",
		Target: NewImpl,
	},
)
