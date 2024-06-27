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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

// +kubebuilder:object:root=true
// +chaos-mesh:experiment

// CertificateChaos is the control script`s spec.
type CertificateChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the behavior of a certificate chaos experiment
	Spec CertificateChaosSpec `json:"spec"`

	// +optional
	// Most recently observed status of the chaos experiment about pods
	Status CertificateChaosStatus `json:"status,omitempty"`
}

var _ InnerObjectWithCustomStatus = (*CertificateChaos)(nil)
var _ InnerObjectWithSelector = (*CertificateChaos)(nil)
var _ InnerObject = (*CertificateChaos)(nil)

// CertificateChaosSpec defines the attributes that a user creates on a chaos experiment about pods.
type CertificateChaosSpec struct {
	CertificateSelector `json:"selector"`

	// Duration represents the duration of the chaos action.
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// +optional
	// +kubebuilder:default="90m"
	Duration *string `json:"duration,omitempty" webhook:"Duration"`

	// CertificateExpiry represents the expiry period for the requested certificate.
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// +optional
	// +kubebuilder:default="1h"
	CertificateExpiry *metav1.Duration `json:"certificateExpiry,omitempty"`

	// RenewBefore represents when the cert-manager should rotate the certificate.
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// +optional
	// +kubebuilder:default="30m"
	RenewBefore *metav1.Duration `json:"renewBefore,omitempty"`

	// RemoteCluster represents the remote cluster where the chaos will be deployed
	// +optional
	RemoteCluster string `json:"remoteCluster,omitempty"`
}

type CertificateSelector struct {
	GenericSelectorSpec `json:",inline"`
}

// CertificateChaosStatus represents the current status of the chaos experiment about pods.
type CertificateChaosStatus struct {
	ChaosStatus `json:",inline"`

	// Instances keeps track of the state for each certificate
	// +optional
	Instances map[string]Instance `json:"affectedFluxResources,omitempty"`
}

type Instance struct {
	FluxResource        FluxResource     `json:"fluxResource"`
	OriginalExpiry      *metav1.Duration `json:"originalExpiry,omitempty"`
	OriginalRenewBefore *metav1.Duration `json:"originalRenewBefore,omitempty"`
	CertificateReadyAt  *metav1.Time     `json:"certificateReadyAt,omitempty"`
	SecretName          string           `json:"secretName,omitempty"`
}

type FluxResource struct {
	Group     string `json:"group"`
	Version   string `json:"version"`
	Kind      string `json:"kind"`
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

func (e *FluxResource) NamespacedName() string {
	return types.NamespacedName{Name: e.Name, Namespace: e.Namespace}.String()
}

func (e *FluxResource) GVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   e.Group,
		Kind:    e.Kind,
		Version: e.Version,
	}
}

func (obj *CertificateChaos) GetSelectorSpecs() map[string]interface{} {
	return map[string]interface{}{
		".": &obj.Spec.CertificateSelector,
	}
}

func (obj *CertificateChaos) GetCustomStatus() interface{} {
	return &obj.Status.Instances
}
