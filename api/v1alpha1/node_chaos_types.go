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

type NodeChaosAction string

const (
	NodeNetworkLossAction NodeChaosAction = "network-loss"
)

var _ InnerObject = (*NodeChaos)(nil)
var _ InnerObjectWithSelector = (*NodeChaos)(nil)

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="action",type=string,JSONPath=`.spec.action`
// +kubebuilder:printcolumn:name="duration",type=string,JSONPath=`.spec.duration`
// +chaos-mesh:experiment
type NodeChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec NodeChaosSpec `json:"spec"`

	Status NodeChaosStatus `json:"status,omitempty"`
}

func (nc *NodeChaos) GetSelectorSpecs() map[string]interface{} {
	return map[string]interface{}{
		".": &nc.Spec.NodeSelector,
	}
}

type NodeChaosSpec struct {
	// +kubebuilder:validation:Enum=network-loss
	Action NodeChaosAction `json:"action"`

	// +ui:form:when=action=='network-loss'
	// +optional
	NetworkLoss *NetworkLossSpec `json:"network-loss,omitempty"`

	NodeSelector `json:",inline"`

	// Duration represents the duration of the chaos action
	// +optional
	Duration *string `json:"duration,omitempty"`

	// RemoteCluster represents the remote cluster where the chaos will be deployed
	// +optional
	RemoteCluster string `json:"remoteCluster,omitempty"`
}

type NodeSelector struct {
	Selector NodeSelectorSpec `json:"selector,omitempty"`

	// Mode defines the mode to run chaos action.
	// Supported mode: one / all / fixed / fixed-percent / random-max-percent
	// +kubebuilder:validation:Enum=one;all;fixed;fixed-percent;random-max-percent
	Mode SelectorMode `json:"mode"`

	// Value is required when the mode is set to `FixedMode` / `FixedPercentMode` / `RandomMaxPercentMode`.
	// If `FixedMode`, provide an integer of physical machines to do chaos action.
	// If `FixedPercentMode`, provide a number from 0-100 to specify the percent of physical machines the server can do chaos action.
	// IF `RandomMaxPercentMode`,  provide a number from 0-100 to specify the max percent of pods to do chaos action
	// +optional
	Value string `json:"value,omitempty"`
}

type NodeSelectorSpec struct {
	GenericSelectorSpec `json:",inline"`
}

type NodeChaosStatus struct {
	ChaosStatus `json:",inline"`
}
