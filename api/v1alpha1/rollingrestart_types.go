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
//

package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +chaos-mesh:experiment
// +chaos-mesh:oneshot=

// RollingRestartChaos is the Schema for the rolling restart API
type RollingRestartChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the behavior of a pod chaos experiment
	Spec RollingRestartChaosSpec `json:"spec"`

	// +optional
	// Most recently observed status of the chaos experiment about pods
	Status RollingRestartChaosStatus `json:"status,omitempty"`
}

// RollingRestartChaosSpec defines the desired state of RollingRestartChaos
type RollingRestartChaosSpec struct {
	RollingRestartSelector `json:",inline"`

	// RemoteCluster represents the remote cluster where the chaos will be deployed
	// +optional
	RemoteCluster string `json:"remoteCluster,omitempty"`
}

type RollingRestartSelector struct {
	GenericSelectorSpec `json:",inline"`
}

// RollingRestartChaosStatus defines the observed state of RollingRestartChaos
type RollingRestartChaosStatus struct {
	ChaosStatus `json:",inline"`
}

func (obj *RollingRestartChaos) GetSelectorSpecs() map[string]interface{} {
	return map[string]interface{}{
		".": &obj.Spec.RollingRestartSelector,
	}
}

func (selector *RollingRestartSelector) Id() string {
	json, _ := json.Marshal(selector)
	return string(json)
}
