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
// +kubebuilder:resource:shortName=csh
// +kubebuilder:printcolumn:name="action",type=string,JSONPath=`.spec.action`
// +kubebuilder:printcolumn:name="duration",type=string,JSONPath=`.spec.duration`
// +chaos-mesh:experiment
// +chaos-mesh:oneshot=in.Spec.Action==HostStop

// CloudStackHostChaos is the Schema for the cloudstackchaos API.
type CloudStackHostChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudStackHostChaosSpec   `json:"spec"`
	Status CloudStackHostChaosStatus `json:"status,omitempty"`
}

var (
	_ InnerObjectWithSelector = (*CloudStackHostChaos)(nil)
	_ InnerObject             = (*CloudStackHostChaos)(nil)
)

// CloudStackHostChaosAction represents the chaos action about cloudstack.
type CloudStackHostChaosAction string

const (
	// HostStop represents the chaos action of stopping the host.
	HostStop CloudStackHostChaosAction = "host-stop"
)

// CloudStackHostChaosSpec is the content of the specification for a CloudStackChaos.
type CloudStackHostChaosSpec struct {
	// APIConfig defines the configuration ncessary to connect to the CloudStack API.
	APIConfig CloudStackAPIConfig `json:"apiConfig"`

	// Selector defines the parameters that can be used to select target VMs.
	Selector CloudStackHostChaosSelector `json:"selector"`

	// DryRun defines whether the chaos should run a dry-run mode.
	// +optional
	DryRun bool `json:"dryRun,omitempty"`

	// Action defines the specific cloudstack chaos action.
	// Supported action: host-stop
	// Default action: host-stop
	// +kubebuilder:validation:Enum=host-stop
	Action CloudStackHostChaosAction `json:"action"`

	// Duration represents the duration of the chaos action.
	// +optional
	Duration *string `json:"duration,omitempty" webhook:"Duration"`

	// RemoteCluster represents the remote cluster where the chaos will be deployed
	// +optional
	RemoteCluster string `json:"remoteCluster,omitempty"`
}

// CloudStackHostChaosStatus represents the status of a CloudStackChaos.
type CloudStackHostChaosStatus struct {
	ChaosStatus `json:",inline"`

	// Instances keeps track of the affected hosts and vms
	// +optional
	Instances map[string]CloudStackHostAffected `json:"affectedHosts,omitempty"`
}

type CloudStackHostAffected struct {
	Name string   `json:"name,omitempty"`
	VMs  []string `json:"vms,omitempty"`
}

type CloudStackHostChaosSelector struct {
	// Hypervisor defines the target hypervisor.
	// +optional
	Hypervisor *string `json:"hypervisor,omitempty"`

	// ID defines the ID of the host.
	// +optional
	ID *string `json:"id,omitempty"`

	// Keyword defines the keyword to list the VMs by.
	// +optional
	Keyword *string `json:"keyword,omitempty"`

	// Name defines the name of the host
	// +optiional
	Name *string `json:"name,omitempty"`

	// ZoneID defines the availability zone the host belongs to.
	// +optional
	ZoneID *string `json:"zoneId,omitempty"`

	// ClusterID defines the cluster the host belongs to.
	// +optional
	ClusterID *string `json:"clusterId,omitempty"`
}

func (selector *CloudStackHostChaosSelector) Id() string {
	v, _ := json.Marshal(selector)
	return string(v)
}

func (obj *CloudStackHostChaos) GetSelectorSpecs() map[string]interface{} {
	return map[string]interface{}{".": &obj.Spec.Selector}
}
