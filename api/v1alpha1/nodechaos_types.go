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

// +kubebuilder:object:root=true
// +chaos-mesh:experiment

// NodeChaos is the control script`s spec.
type NodeChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the behavior of a node chaos experiment
	Spec NodeChaosSpec `json:"spec"`

	// +optional
	// Most recently observed status of the chaos experiment about nodes
	Status NodeChaosStatus `json:"status,omitempty"`
}

var _ InnerObject = (*NodeChaos)(nil)

// NodeChaosSpec defines the attributes that a user creates on a chaos experiment about nodes.
type NodeChaosSpec struct {
	NodeSelector `json:",inline"`

	// Duration represents the duration of the chaos action.
	// It is required when the action is `PodFailureAction`.
	// A duration string is a possibly signed sequence of
	// decimal numbers, each with optional fraction and a unit suffix,
	// such as "300ms", "-1.5h" or "2h45m".
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// +optional
	Duration *string `json:"duration,omitempty" webhook:"Duration"`

	// RemoteCluster represents the remote cluster where the chaos will be deployed
	// +optional
	RemoteCluster string `json:"remoteCluster,omitempty"`
}

// PodChaosStatus represents the current status of the chaos experiment about pods.
type NodeChaosStatus struct {
	ChaosStatus `json:",inline"`
}

func (obj *NodeChaos) GetSelectorSpecs() map[string]interface{} {
	return map[string]interface{}{
		".": &obj.Spec.NodeSelector,
	}
}
