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

package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="duration",type=string,JSONPath=`.spec.duration`
// +chaos-mesh:experiment

// GKENodePoolChaos is the Schema for the gkenodepoolchaos API
type GKENodePoolChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GKENodePoolChaosSpec   `json:"spec"`
	Status GKENodePoolChaosStatus `json:"status,omitempty"`
}

// GKENodePoolChaosSpec defines the desired state of GKENodePoolChaos
type GKENodePoolChaosSpec struct {
	// Duration represents the duration of the chaos
	// +optional
	Duration *string `json:"duration,omitempty"`

	GKENodePoolSelector `json:",inline"`

	// RemoteCluster represents the remote cluster where the chaos will be deployed
	// +optional
	RemoteCluster string `json:"remoteCluster,omitempty"`
}

type GKENodePoolSelector struct {
	// Project defines the ID of GCP project.
	Project string `json:"project"`

	// Cluster defines the Kubernetes cluster to target.
	Cluster string `json:"cluster"`

	// Location defines the location/region of the Kubernetes cluster.
	Location string `json:"location"`
}

// GKENodePoolChaosStatus defines the observed state of HelloWorldChaos
type GKENodePoolChaosStatus struct {
	ChaosStatus `json:",inline"`
}

// GetSelectorSpecs is a getter for selectors
func (obj *GKENodePoolChaos) GetSelectorSpecs() map[string]interface{} {
	return map[string]interface{}{
		".": &obj.Spec.GKENodePoolSelector,
	}
}

func (selector *GKENodePoolSelector) Id() string {
	json, _ := json.Marshal(selector)
	return string(json)
}
