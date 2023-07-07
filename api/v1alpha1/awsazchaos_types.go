package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="duration",type=string,JSONPath=`.spec.duration`
// +chaos-mesh:experiment

// AWSAzChaos is the Schema for the helloworldchaos API
type AWSAzChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AWSAzChaosSpec   `json:"spec"`
	Status AWSAzChaosStatus `json:"status,omitempty"`
}

var (
	_ InnerObjectWithCustomStatus = (*AWSAzChaos)(nil)
	_ InnerObjectWithSelector     = (*AWSAzChaos)(nil)
	_ InnerObject                 = (*AWSAzChaos)(nil)
)

// AWSAzChaosSpec is the content of the specification for a AWSAzChaos
type AWSAzChaosSpec struct {
	// ContainerSelector specifies target
	AWSAZSelector `json:",inline"`

	// Duration represents the duration of the chaos action
	// +optional
	Duration *string `json:"duration,omitempty"`

	// RemoteCluster represents the remote cluster where the chaos will be deployed
	// +optional
	RemoteCluster string `json:"remoteCluster,omitempty"`
}

// AWSAzChaosStatus represents the status of a HelloWorldChaos
type AWSAzChaosStatus struct {
	ChaosStatus `json:",inline"`
	// SubnetToACL represents the connection between a subnet and its Network ACL
	SubnetToACL map[string]string `json:"subnetToACL,omitempty"`
}

type AWSAZSelector struct {
	// TODO: it would be better to split them into multiple different selector and implementation
	// but to keep the minimal modification on current implementation, it hasn't been splited.

	// AWSRegion defines the region of aws.
	Stack string `json:"stack"`

	// AvailabilityZone indicates the Availability zone to be taken down
	AvailabilityZone string `json:"az"`
}

// GetSelectorSpecs is a getter for selectors
func (obj *AWSAzChaos) GetSelectorSpecs() map[string]interface{} {
	return map[string]interface{}{
		".": &obj.Spec.AWSAZSelector,
	}
}

func (obj *AWSAZSelector) Id() string {
	// TODO: handle the error here
	// or ignore it is enough ?
	json, _ := json.Marshal(obj)

	return string(json)
}

func (obj *AWSAzChaos) GetCustomStatus() interface{} {
	return &obj.Status.SubnetToACL
}
