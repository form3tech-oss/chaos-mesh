package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="duration",type=string,JSONPath=`.spec.duration`
// +chaos-mesh:experiment

// GCPAzChaos is the Schema for the helloworldchaos API
type GCPAzChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GCPAzChaosSpec   `json:"spec"`
	Status GCPAzChaosStatus `json:"status,omitempty"`
}

var (
	_ InnerObjectWithCustomStatus = (*GCPAzChaos)(nil)
	_ InnerObjectWithSelector     = (*GCPAzChaos)(nil)
	_ InnerObject                 = (*GCPAzChaos)(nil)
)

// GCPAzChaosSpec is the content of the specification for a GCPAzChaos
type GCPAzChaosSpec struct {
	// ContainerSelector specifies target
	GCPAzSelector `json:",inline"`

	// Duration represents the duration of the chaos action
	// +optional
	Duration *string `json:"duration,omitempty"`

	RemoteCluster string `json:"remoteCluster,omitempty"`
}

// GCPAzChaosStatus represents the status of a HelloWorldChaos
type GCPAzChaosStatus struct {
	ChaosStatus `json:",inline"`

	// InstanceGroupSizes represents a map of instance group names to their sizes
	InstanceGroupSizes map[string]int64 `json:"instanceGroupSizes,omitempty"`
}

type GCPAzSelector struct {
	// Project defines the GCP project.
	Project string `json:"stack"`

	// AZ indicates the Availability zone to be taken down
	Zone string `json:"zone"`
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

func (obj *GCPAzChaos) GetCustomStatus() interface{} {
	return &obj.Status.InstanceGroupSizes
}
