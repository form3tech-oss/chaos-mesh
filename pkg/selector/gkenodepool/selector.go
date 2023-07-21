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

package gkenodepool

import (
	"context"
	"encoding/json"
	"fmt"

	container "google.golang.org/api/container/v1"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
)

type SelectImpl struct{}

type NodePool struct {
	Name        string                        `json:"name,omitempty"`
	Autoscaling container.NodePoolAutoscaling `json:"autoscaling,omitempty"`
}

func (np *NodePool) Id() string {
	b, _ := json.Marshal(np)
	return string(b)
}

func (impl *SelectImpl) Select(ctx context.Context, selector *v1alpha1.GKENodePoolSelector) ([]*NodePool, error) {
	nodepools, err := fetchGKENodePools(ctx, selector.Project, selector.Cluster, selector.Location)
	if err != nil {
		return nil, err
	}

	return nodepools, nil
}

func fetchGKENodePools(ctx context.Context, project string, cluster string, location string) ([]*NodePool, error) {
	containerSvc, err := container.NewService(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error creating container service: %w", err)
	}

	nodepoolSvc := container.NewProjectsLocationsClustersNodePoolsService(containerSvc)
	resp, err := nodepoolSvc.List(fmt.Sprintf("projects/%s/locations/%s/clusters/%s", project, location, cluster)).Do()
	if err != nil {
		return nil, fmt.Errorf("error listing node pools for cluster %q: %w", cluster, err)
	}

	var nodepools []*NodePool
	for _, nodepool := range resp.NodePools {
		req := nodepoolSvc.Get(fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", project, location, cluster, nodepool.Name))
		req.Fields("autoscaling", "name")

		nodepool, err := req.Do()
		if err != nil {
			return nil, fmt.Errorf("error getting node pool %q for cluster %q: %w", nodepool.Name, cluster, err)
		}

		if nodepool.Autoscaling == nil {
			continue
		}

		nodepools = append(nodepools, &NodePool{
			Name:        nodepool.Name,
			Autoscaling: *nodepool.Autoscaling,
		})
	}

	return nodepools, nil
}

func New() *SelectImpl {
	return &SelectImpl{}
}
