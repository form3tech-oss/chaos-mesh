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

package utils

import (
	"context"
	"fmt"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
)

// GetCloudStackClient is used to get a CloudStack client.
func GetCloudStackClient(ctx context.Context, cli client.Client, cloudstackchaos *v1alpha1.CloudStackHostChaos) (*cloudstack.CloudStackClient, error) {
	apiConfig := cloudstackchaos.Spec.APIConfig

	var secret v1.Secret
	if err := cli.Get(ctx, types.NamespacedName{Namespace: cloudstackchaos.Namespace, Name: apiConfig.SecretName}, &secret); err != nil {
		return nil, fmt.Errorf("retrieving secret for cloudstack api client: %w", err)
	}

	apiKey, ok := secret.Data[apiConfig.APIKeyField]
	if !ok {
		return nil, fmt.Errorf("field %s not found in secret %s", apiConfig.APIKeyField, apiConfig.SecretName)
	}

	apiSecret, ok := secret.Data[apiConfig.APISecretField]
	if !ok {
		return nil, fmt.Errorf("field %s not found in secret %s", apiConfig.APIKeyField, apiConfig.SecretName)
	}

	return cloudstack.NewAsyncClient(
		cloudstackchaos.Spec.APIConfig.Address,
		string(apiKey),
		string(apiSecret),
		apiConfig.VerifySSL,
	), nil
}

func SelectorToListParams(s *v1alpha1.CloudStackHostChaosSelector) *cloudstack.ListHostsParams {
	params := &cloudstack.ListHostsParams{}

	if s.Hypervisor != nil {
		params.SetHypervisor(*s.Hypervisor)
	}

	if s.ID != nil {
		params.SetId(*s.ID)
	}

	if s.Keyword != nil {
		params.SetKeyword(*s.Keyword)
	}

	if s.Name != nil {
		params.SetName(*s.Name)
	}

	if s.ZoneID != nil {
		params.SetZoneid(*s.ZoneID)
	}

	if s.ClusterID != nil {
		params.SetClusterid(*s.ClusterID)
	}

	return params
}
