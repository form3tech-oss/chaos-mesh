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

package nodechaos

import (
	"context"
	"strconv"

	"github.com/go-logr/logr"
	"go.uber.org/fx"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	impltypes "github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/types"
	"github.com/chaos-mesh/chaos-mesh/controllers/utils/chaosdaemon"
	"github.com/chaos-mesh/chaos-mesh/pkg/chaosdaemon/pb"
)

type Impl struct {
	client.Client
	Log logr.Logger

	ChaosDaemonClientBuilder *chaosdaemon.ChaosDaemonClientBuilder
}

func (impl *Impl) Apply(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	impl.Log.Info("apply nodechaos", "node", records[index].Id)

	nodeChaos := obj.(*v1alpha1.NodeChaos)

	pbClient, err := impl.ChaosDaemonClientBuilder.NodeName(records[index].Id).Build(ctx, &types.NamespacedName{
		Namespace: nodeChaos.Namespace,
		Name:      nodeChaos.Name,
	})
	if err != nil {
		return v1alpha1.NotInjected, err
	}
	defer pbClient.Close()

	tcs := []*pb.Tc{}
	lossPct, err := strconv.ParseFloat(nodeChaos.Spec.NetworkLoss.Percent, 32)
	if err != nil {
		return v1alpha1.NotInjected, err
	}
	netem := pb.Netem{
		Loss:     float32(lossPct),
		LossCorr: 100.0,
	}
	tcs = append(tcs, &pb.Tc{Type: pb.Tc_NETEM, Netem: &netem, Device: nodeChaos.Spec.NetworkLoss.Device})
	impl.Log.Info("nodechaos apply", "tcs", tcs)

	_, err = pbClient.SetTcs(ctx, &pb.TcsRequest{
		Tcs:         tcs,
		ContainerId: "host",
		EnterNS:     true,
	})
	if err != nil {
		return v1alpha1.NotInjected, err
	}

	return v1alpha1.Injected, nil
}

func (impl *Impl) Recover(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	impl.Log.Info("recover nodechaos", "node", records[index].Id)

	nodeChaos := obj.(*v1alpha1.NodeChaos)

	pbClient, err := impl.ChaosDaemonClientBuilder.NodeName(records[index].Id).Build(ctx, &types.NamespacedName{
		Namespace: nodeChaos.Namespace,
		Name:      nodeChaos.Name,
	})
	if err != nil {
		return v1alpha1.Injected, err
	}
	defer pbClient.Close()

	_, err = pbClient.SetTcs(ctx, &pb.TcsRequest{
		Tcs:         []*pb.Tc{},
		ContainerId: "host",
		EnterNS:     true,
	})
	if err != nil {
		return v1alpha1.Injected, err
	}

	return v1alpha1.NotInjected, nil
}

func NewImpl(c client.Client, log logr.Logger, clientBuilder *chaosdaemon.ChaosDaemonClientBuilder) *impltypes.ChaosImplPair {
	return &impltypes.ChaosImplPair{
		Name:   "nodechaos",
		Object: &v1alpha1.NodeChaos{},
		Impl: &Impl{
			Client:                   c,
			Log:                      log.WithName("nodechaos"),
			ChaosDaemonClientBuilder: clientBuilder,
		},
	}
}

var Module = fx.Provide(
	fx.Annotated{
		Group:  "impl",
		Target: NewImpl,
	},
)
