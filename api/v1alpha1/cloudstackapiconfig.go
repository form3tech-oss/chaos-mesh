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

type CloudStackAPIConfig struct {
	// Address defines the address of the CloudStack instsance.
	Address string `json:"address"`

	// VerifySSL defines whether certificates should be verified when connecting to the API.
	// +optional
	VerifySSL bool `json:"verifySSL,omitempty"`

	// SecretName defines the name of the secret where the API credentials are stored.
	SecretName string `json:"secretName"`

	// APIKeyField defines the key under which the value for API key is stored inside the secret.
	// +optional
	APIKeyField string `json:"apiKeyField,omitempty" default:"api-key"`

	// APISecretField defines the key under which the value for API secret is stored inside the secret.
	// +optional
	APISecretField string `json:"apiSecretField,omitempty" default:"api-secret"`
}
