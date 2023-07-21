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

// GCPAzChaos is the Schema for the GCPAzChaos API
type GCPAzChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GCPAzChaosSpec   `json:"spec"`
	Status GCPAzChaosStatus `json:"status,omitempty"`
}

var (
	_ InnerObjectWithSelector = (*GCPAzChaos)(nil)
	_ InnerObject             = (*GCPAzChaos)(nil)
)

// GCPAzChaosSpec is the content of the specification for a GCPAzChaos
type GCPAzChaosSpec struct {
	// ContainerSelector specifies target
	GCPAzSelector `json:",inline"`

	// Duration represents the duration of the chaos action
	// +optional
	Duration *string `json:"duration,omitempty"`

	// RemoteCluster represents the remote cluster where the chaos will be deployed
	// +optional
	RemoteCluster string `json:"remoteCluster,omitempty"`
}

// GCPAzChaosStatus represents the status of a GCPAzChaos
type GCPAzChaosStatus struct {
	ChaosStatus `json:",inline"`
}

type GCPAzSelector struct {
	// Project defines the GCP project.
	Project string `json:"project"`

	// Zone defines the zone of GCP project.
	Zone string `json:"zone"`

	// Filter defines the filter used when fetching GCP instance groups.
	// Filter reference: https://google.aip.dev/160
	// +optional
	Filter string `json:"filter"`
}

// GetSelectorSpecs is a getter for selectors
func (obj *GCPAzChaos) GetSelectorSpecs() map[string]interface{} {
	return map[string]interface{}{
		".": &obj.Spec.GCPAzSelector,
	}
}

func (obj *GCPAzSelector) Id() string {
	json, _ := json.Marshal(obj)
	return string(json)
}
