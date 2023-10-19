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

package resourcescale

import (
	"context"
	"encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
)

type SelectImpl struct{}

type ResourceSpec struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
	Replicas  int32  `json:"replicas,omitempty"`
}

func (dep *ResourceSpec) Id() string {
	b, _ := json.Marshal(dep)
	return string(b)
}

func (impl *SelectImpl) Select(ctx context.Context, selector *v1alpha1.ResourceScaleSelector) ([]*ResourceSpec, error) {
	client, err := kubernetesClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create kubernetes client: %w", err)
	}

	// TODO: support different scalable resources
	scale, err := client.AppsV1().Deployments(selector.Namespace).GetScale(ctx, selector.Name, v1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get deployment scale for %q: %w", selector.Namespace+"/"+selector.Name, err)
	}

	dep := &ResourceSpec{
		Namespace: selector.Namespace,
		Name:      selector.Name,
		Replicas:  scale.Spec.Replicas,
	}

	return []*ResourceSpec{dep}, nil
}

func New() *SelectImpl {
	return &SelectImpl{}
}

func kubernetesClient() (*kubernetes.Clientset, error) {
	config, err := ctrl.GetConfig()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}
