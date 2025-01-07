// Copyright 2021 Chaos Mesh Authors.PodPVCChaos
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

package deployment

import (
	"context"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
)

type Deployment struct {
	v1.Deployment
}

func (d *Deployment) Id() string {
	return (types.NamespacedName{
		Name:      d.Name,
		Namespace: d.Namespace,
	}).String()
}

type SelectImpl struct{}

func (impl *SelectImpl) Select(ctx context.Context, selector *v1alpha1.DeploymentSelector) ([]*Deployment, error) {
	if selector == nil {
		return []*Deployment{}, nil
	}

	client, err := kubernetesClient()
	if err != nil {
		return []*Deployment{}, err
	}

	var deployments []*Deployment
	for namespace, names := range selector.Deployments {
		deploymentList, err := client.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
		if err != nil {
			return []*Deployment{}, err
		}

		matched := match(deploymentList, names)
		deployments = append(deployments, matched...)

	}

	return deployments, nil
}

func match(list *v1.DeploymentList, names []string) []*Deployment {
	var deployments []*Deployment
	for _, selectorName := range names {
		for _, deployment := range list.Items {
			if selectorName == deployment.Name {
				deployments = append(deployments, &Deployment{
					Deployment: deployment,
				})
			}
		}
	}

	return deployments
}

func kubernetesClient() (*kubernetes.Clientset, error) {
	config, err := ctrl.GetConfig()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func New() *SelectImpl {
	return &SelectImpl{}
}
