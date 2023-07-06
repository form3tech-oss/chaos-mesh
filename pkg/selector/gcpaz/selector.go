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

package gcpaz

import (
	"context"
	"encoding/json"

	compute "google.golang.org/api/compute/v1"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
)

type SelectImpl struct{}

type InstanceGroup struct {
	Project string `json:"project,omitempty"`
	Zone    string `json:"zone,omitempty"`
	Name    string `json:"name,omitempty"`
	Size    int64  `json:"size,omitempty"`
}

func (ig *InstanceGroup) Id() string {
	b, _ := json.Marshal(ig)
	return string(b)
}

func (impl *SelectImpl) Select(ctx context.Context, selector *v1alpha1.GCPAzSelector) ([]*InstanceGroup, error) {
	groupSizes, err := getInstanceGroupSizes(ctx, selector.Project, selector.Zone, selector.Filter)
	if err != nil {
		return nil, err
	}

	var igs []*InstanceGroup
	for name, size := range groupSizes {
		igs = append(igs, &InstanceGroup{
			Project: selector.Project,
			Zone:    selector.Zone,
			Name:    name,
			Size:    size,
		})
	}

	return igs, nil
}

func New() *SelectImpl {
	return &SelectImpl{}
}

func getInstanceGroupSizes(ctx context.Context, project string, zone string, filter string) (map[string]int64, error) {
	computeService, err := compute.NewService(ctx)
	if err != nil {
		return nil, err
	}

	req := computeService.InstanceGroupManagers.List(project, zone)
	if filter != "" {
		req.Filter(filter)
	}

	instanceGroups, err := req.Do()
	if err != nil {
		return nil, err
	}

	sizes := make(map[string]int64)
	for _, group := range instanceGroups.Items {
		sizes[group.Name] = group.TargetSize
	}
	return sizes, nil
}
