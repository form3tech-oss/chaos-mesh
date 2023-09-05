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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +chaos-mesh:base
// +chaos-mesh:webhook:enableUpdate

// NodeNetworkChaos is the Schema for the NodeNetworkChaos API
type NodeNetworkChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the behavior of a node network chaos experiment
	Spec NodeNetworkChaosSpec `json:"spec"`

	// +optional
	// Most recently observed status of the chaos experiment about nodes
	Status NodeNetworkChaosStatus `json:"status,omitempty"`
}

// NodeNetworkChaosSpec defines the desired state of NodeNetworkChaos
type NodeNetworkChaosSpec struct {
}

// NodeNetworkChaosStatus defines the observed state of NodeNetworkChaos
type NodeNetworkChaosStatus struct {
	FailedMessage string `json:"failedMessage,omitempty"`

	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

// +kubebuilder:object:root=true

// NodeNetworkChaosList contains a list of NodeNetworkChaos
type NodeNetworkChaosList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeNetworkChaos `json:"items"`
}
