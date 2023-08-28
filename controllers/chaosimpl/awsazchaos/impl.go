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

package awsazchaos

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-logr/logr"
	"go.uber.org/fx"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/awsazchaos/ctxutil"
	"github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/awsazchaos/subnetloss"
	impltypes "github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/types"
)

type Impl struct {
	client.Client
	Log logr.Logger
}

const (
	waitForApplySync v1alpha1.Phase = "Not Injected/Wait"
)

// Apply applies KernelChaos
func (impl *Impl) Apply(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	impl.Log.Info("Apply awsazchaos chaos")

	awsAZChaos := obj.(*v1alpha1.AWSAzChaos)
	ctx = context.WithValue(ctx, ctxutil.CtxKeySimulationId, awsAZChaos.Name)

	var selected v1alpha1.AWSAZSelector
	record := records[index]
	err := json.Unmarshal([]byte(record.Id), &selected)
	if err != nil {
		impl.Log.Error(err, "fail to unmarshal the selector")
		return v1alpha1.NotInjected, err
	}

	azLoss, err := subnetloss.NewAWSAzLoss(ctx, selected.Stack, selected.AvailabilityZone, impl.Log)
	if err != nil {
		impl.Log.Error(err, "fail to create NewAWSAzLoss")
		return v1alpha1.NotInjected, err
	}

	phase := record.Phase
	if phase == waitForApplySync {
		impl.Log.Info(fmt.Sprintf("Applying awsazchaos chaos for stack (%s) and AZ (%s)", selected.Stack, selected.AvailabilityZone))
		err := azLoss.Start(ctx, awsAZChaos.Status.SubnetToACL)
		if err != nil {
			impl.Log.Error(err, "fail to start NewAWSAzLoss")
			return waitForApplySync, err
		}
		return v1alpha1.Injected, nil
	}

	subnetToACL, err := azLoss.GetSubnetToACL(ctx)
	if err != nil {
		impl.Log.Error(err, "fail to get initial state")
		return v1alpha1.NotInjected, err
	}
	awsAZChaos.Status.SubnetToACL = subnetToACL
	return waitForApplySync, nil
}

// Recover means the reconciler recovers the chaos action
func (impl *Impl) Recover(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	impl.Log.Info("Recover awsazchaos chaos")

	awsAZChaos := obj.(*v1alpha1.AWSAzChaos)
	ctx = context.WithValue(ctx, ctxutil.CtxKeySimulationId, awsAZChaos.Name)

	var selected v1alpha1.AWSAZSelector
	err := json.Unmarshal([]byte(records[index].Id), &selected)
	if err != nil {
		impl.Log.Error(err, "fail to unmarshal the selector")
		return v1alpha1.Injected, err
	}

	azLoss, err := subnetloss.NewAWSAzLoss(ctx, selected.Stack, selected.AvailabilityZone, impl.Log)
	if err != nil {
		impl.Log.Error(err, "fail to create NewAWSAzLoss")
		return v1alpha1.Injected, err
	}
	impl.Log.Info(fmt.Sprintf("Recovering awsazchaos chaos for stack (%s) and AZ (%s)", selected.Stack, selected.AvailabilityZone))
	err = azLoss.Stop(ctx, awsAZChaos.Status.SubnetToACL)
	if err != nil {
		impl.Log.Error(err, fmt.Sprintf("failed to recover awsazchaos chaos for stack (%s) and AZ (%s)", selected.Stack, selected.AvailabilityZone))
		return v1alpha1.Injected, err
	}
	return v1alpha1.NotInjected, nil
}

func NewImpl(c client.Client, log logr.Logger) *impltypes.ChaosImplPair {
	return &impltypes.ChaosImplPair{
		Name:   "awsazchaos",
		Object: &v1alpha1.AWSAzChaos{},
		Impl: &Impl{
			Client: c,
			Log:    log.WithName("awsazchaos"),
		},
		ObjectList: &v1alpha1.AWSAzChaosList{},
	}
}

var Module = fx.Provide(
	fx.Annotated{
		Group:  "impl",
		Target: NewImpl,
	},
)
