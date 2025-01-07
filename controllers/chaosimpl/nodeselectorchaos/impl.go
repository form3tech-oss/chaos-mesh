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

package nodeselectorchaos

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	"go.uber.org/fx"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"

	chaosimpltypes "github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/types"
	"github.com/chaos-mesh/chaos-mesh/controllers/utils/controller"
)

type Impl struct {
	client.Client
	Log logr.Logger
}

func (i *Impl) Apply(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	chaos, ok := obj.(*v1alpha1.NodeSelectorChaos)

	if !ok {
		err := errors.New("not NodeSelectorChaos")
		i.Log.Error(err, "casting InnerObject to NodeSelectorChaos")
		return v1alpha1.NotInjected, err
	}

	name, err := controller.ParseNamespacedName(records[index].Id)
	if err != nil {
		i.Log.Error(err, "parsing record name")
		return v1alpha1.NotInjected, err
	}

	var deployment v1.Deployment
	err = i.Client.Get(ctx, name, &deployment)
	if err != nil {
		i.Log.Error(err, "getting deployment")
		return v1alpha1.NotInjected, err
	}

	data := []byte(fmt.Sprintf(`{"spec": {"template": {"spec": {"nodeSelector": {"%s" :"%s"}}}}}`, chaos.Spec.Key, chaos.Spec.Value))
	patch := client.RawPatch(types.MergePatchType, data)
	err = i.Client.Patch(ctx, &deployment, patch)
	if err != nil {
		i.Log.Error(err, "patching deployment")
		return v1alpha1.NotInjected, err
	}

	return v1alpha1.Injected, nil
}

func (i *Impl) Recover(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	chaos, ok := obj.(*v1alpha1.NodeSelectorChaos)

	if !ok {
		err := errors.New("not NodeSelectorChaos")
		i.Log.Error(err, "casting InnerObject to NodeSelectorChaos")
		return v1alpha1.NotInjected, err
	}

	name, err := controller.ParseNamespacedName(records[index].Id)
	if err != nil {
		i.Log.Error(err, "parsing record name")
		return v1alpha1.Injected, err
	}

	var deployment v1.Deployment
	err = i.Client.Get(ctx, name, &deployment)
	if err != nil {
		i.Log.Error(err, "getting deployment")
		return v1alpha1.Injected, err
	}

	escapedKey := strings.ReplaceAll(chaos.Spec.Key, "/", "~1")
	data := []byte(fmt.Sprintf(`[{"op": "remove", "path": "/spec/template/spec/nodeSelector/%s"}]`, escapedKey))
	patch := client.RawPatch(types.JSONPatchType, data)
	err = i.Client.Patch(ctx, &deployment, patch)
	if err != nil {
		i.Log.Error(err, "patching deployment")
		return v1alpha1.Injected, err
	}

	return v1alpha1.NotInjected, nil
}

func NewImpl(c client.Client, log logr.Logger) *chaosimpltypes.ChaosImplPair {
	return &chaosimpltypes.ChaosImplPair{
		Name:       "nodeselectorchaos",
		Object:     &v1alpha1.NodeSelectorChaos{},
		Impl:       &Impl{c, log.WithName("nodeselectorchaos")},
		ObjectList: &v1alpha1.NodeSelectorChaosList{},
	}
}

var Module = fx.Provide(
	fx.Annotated{
		Group:  "impl",
		Target: NewImpl,
	},
)
