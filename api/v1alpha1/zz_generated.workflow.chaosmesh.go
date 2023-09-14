// Copyright Chaos Mesh Authors.
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

// Code generated by chaos-builder. DO NOT EDIT.

package v1alpha1


import (
	"github.com/pkg/errors"
)


const (
	TypeAWSAzChaos TemplateType = "AWSAzChaos"
	TypeAWSChaos TemplateType = "AWSChaos"
	TypeAzureChaos TemplateType = "AzureChaos"
	TypeBlockChaos TemplateType = "BlockChaos"
	TypeDeploymentChaos TemplateType = "DeploymentChaos"
	TypeDNSChaos TemplateType = "DNSChaos"
	TypeGCPAzChaos TemplateType = "GCPAzChaos"
	TypeGCPChaos TemplateType = "GCPChaos"
	TypeGKENodePoolChaos TemplateType = "GKENodePoolChaos"
	TypeHTTPChaos TemplateType = "HTTPChaos"
	TypeIOChaos TemplateType = "IOChaos"
	TypeJVMChaos TemplateType = "JVMChaos"
	TypeK8SChaos TemplateType = "K8SChaos"
	TypeKernelChaos TemplateType = "KernelChaos"
	TypeNetworkChaos TemplateType = "NetworkChaos"
	TypePhysicalMachineChaos TemplateType = "PhysicalMachineChaos"
	TypePodChaos TemplateType = "PodChaos"
	TypeStressChaos TemplateType = "StressChaos"
	TypeTimeChaos TemplateType = "TimeChaos"

)

var allChaosTemplateType = []TemplateType{
	TypeSchedule,
	TypeAWSAzChaos,
	TypeAWSChaos,
	TypeAzureChaos,
	TypeBlockChaos,
	TypeDeploymentChaos,
	TypeDNSChaos,
	TypeGCPAzChaos,
	TypeGCPChaos,
	TypeGKENodePoolChaos,
	TypeHTTPChaos,
	TypeIOChaos,
	TypeJVMChaos,
	TypeK8SChaos,
	TypeKernelChaos,
	TypeNetworkChaos,
	TypePhysicalMachineChaos,
	TypePodChaos,
	TypeStressChaos,
	TypeTimeChaos,

}

type EmbedChaos struct {
	// +optional
	AWSAzChaos *AWSAzChaosSpec `json:"awsazChaos,omitempty"`
	// +optional
	AWSChaos *AWSChaosSpec `json:"awsChaos,omitempty"`
	// +optional
	AzureChaos *AzureChaosSpec `json:"azureChaos,omitempty"`
	// +optional
	BlockChaos *BlockChaosSpec `json:"blockChaos,omitempty"`
	// +optional
	DeploymentChaos *DeploymentChaosSpec `json:"deploymentChaos,omitempty"`
	// +optional
	DNSChaos *DNSChaosSpec `json:"dnsChaos,omitempty"`
	// +optional
	GCPAzChaos *GCPAzChaosSpec `json:"gcpazChaos,omitempty"`
	// +optional
	GCPChaos *GCPChaosSpec `json:"gcpChaos,omitempty"`
	// +optional
	GKENodePoolChaos *GKENodePoolChaosSpec `json:"gkenodepoolChaos,omitempty"`
	// +optional
	HTTPChaos *HTTPChaosSpec `json:"httpChaos,omitempty"`
	// +optional
	IOChaos *IOChaosSpec `json:"ioChaos,omitempty"`
	// +optional
	JVMChaos *JVMChaosSpec `json:"jvmChaos,omitempty"`
	// +optional
	K8SChaos *K8SChaosSpec `json:"k8sChaos,omitempty"`
	// +optional
	KernelChaos *KernelChaosSpec `json:"kernelChaos,omitempty"`
	// +optional
	NetworkChaos *NetworkChaosSpec `json:"networkChaos,omitempty"`
	// +optional
	PhysicalMachineChaos *PhysicalMachineChaosSpec `json:"physicalmachineChaos,omitempty"`
	// +optional
	PodChaos *PodChaosSpec `json:"podChaos,omitempty"`
	// +optional
	StressChaos *StressChaosSpec `json:"stressChaos,omitempty"`
	// +optional
	TimeChaos *TimeChaosSpec `json:"timeChaos,omitempty"`

}

func (it *EmbedChaos) SpawnNewObject(templateType TemplateType) (GenericChaos, error) {
	switch templateType {
	case TypeAWSAzChaos:
		result := AWSAzChaos{}
		result.Spec = *it.AWSAzChaos
		return &result, nil
	case TypeAWSChaos:
		result := AWSChaos{}
		result.Spec = *it.AWSChaos
		return &result, nil
	case TypeAzureChaos:
		result := AzureChaos{}
		result.Spec = *it.AzureChaos
		return &result, nil
	case TypeBlockChaos:
		result := BlockChaos{}
		result.Spec = *it.BlockChaos
		return &result, nil
	case TypeDeploymentChaos:
		result := DeploymentChaos{}
		result.Spec = *it.DeploymentChaos
		return &result, nil
	case TypeDNSChaos:
		result := DNSChaos{}
		result.Spec = *it.DNSChaos
		return &result, nil
	case TypeGCPAzChaos:
		result := GCPAzChaos{}
		result.Spec = *it.GCPAzChaos
		return &result, nil
	case TypeGCPChaos:
		result := GCPChaos{}
		result.Spec = *it.GCPChaos
		return &result, nil
	case TypeGKENodePoolChaos:
		result := GKENodePoolChaos{}
		result.Spec = *it.GKENodePoolChaos
		return &result, nil
	case TypeHTTPChaos:
		result := HTTPChaos{}
		result.Spec = *it.HTTPChaos
		return &result, nil
	case TypeIOChaos:
		result := IOChaos{}
		result.Spec = *it.IOChaos
		return &result, nil
	case TypeJVMChaos:
		result := JVMChaos{}
		result.Spec = *it.JVMChaos
		return &result, nil
	case TypeK8SChaos:
		result := K8SChaos{}
		result.Spec = *it.K8SChaos
		return &result, nil
	case TypeKernelChaos:
		result := KernelChaos{}
		result.Spec = *it.KernelChaos
		return &result, nil
	case TypeNetworkChaos:
		result := NetworkChaos{}
		result.Spec = *it.NetworkChaos
		return &result, nil
	case TypePhysicalMachineChaos:
		result := PhysicalMachineChaos{}
		result.Spec = *it.PhysicalMachineChaos
		return &result, nil
	case TypePodChaos:
		result := PodChaos{}
		result.Spec = *it.PodChaos
		return &result, nil
	case TypeStressChaos:
		result := StressChaos{}
		result.Spec = *it.StressChaos
		return &result, nil
	case TypeTimeChaos:
		result := TimeChaos{}
		result.Spec = *it.TimeChaos
		return &result, nil

	default:
		return nil, errors.Wrapf(errInvalidValue, "unknown template type %s", templateType)
	}
}

func (it *EmbedChaos) RestoreChaosSpec(root interface{}) error {
	switch chaos := root.(type) {
	case *AWSAzChaos:
		*it.AWSAzChaos = chaos.Spec
		return nil
	case *AWSChaos:
		*it.AWSChaos = chaos.Spec
		return nil
	case *AzureChaos:
		*it.AzureChaos = chaos.Spec
		return nil
	case *BlockChaos:
		*it.BlockChaos = chaos.Spec
		return nil
	case *DeploymentChaos:
		*it.DeploymentChaos = chaos.Spec
		return nil
	case *DNSChaos:
		*it.DNSChaos = chaos.Spec
		return nil
	case *GCPAzChaos:
		*it.GCPAzChaos = chaos.Spec
		return nil
	case *GCPChaos:
		*it.GCPChaos = chaos.Spec
		return nil
	case *GKENodePoolChaos:
		*it.GKENodePoolChaos = chaos.Spec
		return nil
	case *HTTPChaos:
		*it.HTTPChaos = chaos.Spec
		return nil
	case *IOChaos:
		*it.IOChaos = chaos.Spec
		return nil
	case *JVMChaos:
		*it.JVMChaos = chaos.Spec
		return nil
	case *K8SChaos:
		*it.K8SChaos = chaos.Spec
		return nil
	case *KernelChaos:
		*it.KernelChaos = chaos.Spec
		return nil
	case *NetworkChaos:
		*it.NetworkChaos = chaos.Spec
		return nil
	case *PhysicalMachineChaos:
		*it.PhysicalMachineChaos = chaos.Spec
		return nil
	case *PodChaos:
		*it.PodChaos = chaos.Spec
		return nil
	case *StressChaos:
		*it.StressChaos = chaos.Spec
		return nil
	case *TimeChaos:
		*it.TimeChaos = chaos.Spec
		return nil

	default:
		return errors.Wrapf(errInvalidValue, "unknown chaos %#v", root)
	}
}

func (it *EmbedChaos) SpawnNewList(templateType TemplateType) (GenericChaosList, error) {
	switch templateType {
	case TypeAWSAzChaos:
		result := AWSAzChaosList{}
		return &result, nil
	case TypeAWSChaos:
		result := AWSChaosList{}
		return &result, nil
	case TypeAzureChaos:
		result := AzureChaosList{}
		return &result, nil
	case TypeBlockChaos:
		result := BlockChaosList{}
		return &result, nil
	case TypeDeploymentChaos:
		result := DeploymentChaosList{}
		return &result, nil
	case TypeDNSChaos:
		result := DNSChaosList{}
		return &result, nil
	case TypeGCPAzChaos:
		result := GCPAzChaosList{}
		return &result, nil
	case TypeGCPChaos:
		result := GCPChaosList{}
		return &result, nil
	case TypeGKENodePoolChaos:
		result := GKENodePoolChaosList{}
		return &result, nil
	case TypeHTTPChaos:
		result := HTTPChaosList{}
		return &result, nil
	case TypeIOChaos:
		result := IOChaosList{}
		return &result, nil
	case TypeJVMChaos:
		result := JVMChaosList{}
		return &result, nil
	case TypeK8SChaos:
		result := K8SChaosList{}
		return &result, nil
	case TypeKernelChaos:
		result := KernelChaosList{}
		return &result, nil
	case TypeNetworkChaos:
		result := NetworkChaosList{}
		return &result, nil
	case TypePhysicalMachineChaos:
		result := PhysicalMachineChaosList{}
		return &result, nil
	case TypePodChaos:
		result := PodChaosList{}
		return &result, nil
	case TypeStressChaos:
		result := StressChaosList{}
		return &result, nil
	case TypeTimeChaos:
		result := TimeChaosList{}
		return &result, nil

	default:
		return nil, errors.Wrapf(errInvalidValue, "unknown template type %s", templateType)
	}
}

func (in *AWSAzChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}
func (in *AWSChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}
func (in *AzureChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}
func (in *BlockChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}
func (in *DeploymentChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}
func (in *DNSChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}
func (in *GCPAzChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}
func (in *GCPChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}
func (in *GKENodePoolChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}
func (in *HTTPChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}
func (in *IOChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}
func (in *JVMChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}
func (in *K8SChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}
func (in *KernelChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}
func (in *NetworkChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}
func (in *PhysicalMachineChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}
func (in *PodChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}
func (in *StressChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}
func (in *TimeChaosList) GetItems() []GenericChaos {
	var result []GenericChaos
	for _, item := range in.Items {
		item := item
		result = append(result, &item)
	}
	return result
}

