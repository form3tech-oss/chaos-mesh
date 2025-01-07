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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	_ InnerObject             = (*NodeSelectorChaos)(nil)
	_ InnerObjectWithSelector = (*NodeSelectorChaos)(nil)
)

// +kubebuilder:object:root=true
// +chaos-mesh:experiment
type NodeSelectorChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeSelectorChaosSpec   `json:"spec"`
	Status NodeSelectorChaosStatus `json:"status"`
}

type NodeSelectorChaosSpec struct {
	DeploymentSelectorSpec `json:"selector"`
	// Key is the name of the key that will be applied to the deployment's nodeSelector field.
	Key string `json:"key"`
	// Value is the value assigned to the provided key.
	Value string `json:"value"`
	// Duration represents the duration of the chaos
	// +optional
	Duration *string `json:"duration,omitempty"`
	// RemoteCluster represents the remote cluster where the chaos will be deployed
	// +optional
	RemoteCluster string `json:"remoteCluster,omitempty"`
}
type NodeSelectorChaosStatus struct {
	ChaosStatus `json:",inline"`
}

func (obj *NodeSelectorChaos) GetSelectorSpecs() map[string]interface{} {
	return map[string]interface{}{
		".": &obj.Spec.DeploymentSelector,
	}
}
