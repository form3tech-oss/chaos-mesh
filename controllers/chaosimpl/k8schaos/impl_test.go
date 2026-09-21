// Copyright 2026 Chaos Mesh Authors.
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

package k8schaos

import (
	"strings"
	"testing"
)

const apiObjects = `---
apiVersion: cilium.io/v2
kind: CiliumClusterwideNetworkPolicy
metadata:
  name: first-object
spec:
  endpointSelector: {}
  egressDeny:
    - toCIDR:
      - 192.168.13.0/24
---
apiVersion: cilium.io/v2
kind: CiliumClusterwideNetworkPolicy
metadata:
  name: second-object
spec:
  endpointSelector: {}
`

func TestResourceForIndex(t *testing.T) {
	impl := &Impl{}

	t.Run("nested mappings marshal back to JSON", func(t *testing.T) {
		resource, err := impl.resourceForIndex(apiObjects, 0)
		if err != nil {
			t.Fatalf("resourceForIndex(apiObjects, 0) unexpected error: %v", err)
		}

		// The dynamic client marshals the object on Create, which fails when the
		// YAML decoder produces map[interface{}]interface{} for nested mappings.
		if _, err := resource.MarshalJSON(); err != nil {
			t.Errorf("MarshalJSON() unexpected error: %v", err)
		}

		if got, want := resource.GetName(), "first-object"; got != want {
			t.Errorf("resourceForIndex(apiObjects, 0) name = %q, want %q", got, want)
		}
	})

	t.Run("index selects the matching document", func(t *testing.T) {
		resource, err := impl.resourceForIndex(apiObjects, 1)
		if err != nil {
			t.Fatalf("resourceForIndex(apiObjects, 1) unexpected error: %v", err)
		}

		if got, want := resource.GetName(), "second-object"; got != want {
			t.Errorf("resourceForIndex(apiObjects, 1) name = %q, want %q", got, want)
		}
	})

	t.Run("index out of range", func(t *testing.T) {
		_, err := impl.resourceForIndex(apiObjects, 2)
		if err == nil {
			t.Fatal("resourceForIndex(apiObjects, 2) expected error, got nil")
		}

		if !strings.Contains(err.Error(), "no resource for index") {
			t.Errorf("resourceForIndex(apiObjects, 2) error = %v, want it to mention no resource for index", err)
		}
	})
}
