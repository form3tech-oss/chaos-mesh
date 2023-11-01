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

package nodenetworkchaos

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/chaos-mesh/chaos-mesh/controllers/utils/chaosdaemon"
	"github.com/chaos-mesh/chaos-mesh/controllers/utils/recorder"
)

type Reconciler struct {
	client.Client
	Recorder                 recorder.ChaosRecorder
	Log                      logr.Logger
	ChaosDaemonClientBuilder *chaosdaemon.ChaosDaemonClientBuilder
}

func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	obj := &v1alpha1.NodeChaos{}

	if err := r.Client.Get(ctx, req.NamespacedName, obj); err != nil {
		if apierrors.IsNotFound(err) {
			r.Log.Info("chaos not found", "nodenetworkchaos", req.NamespacedName)
		} else {
			r.Log.Error(err, "unable to get chaos", "nodenetworkchaos", req.NamespacedName)
		}

		return ctrl.Result{}, nil
	}

	r.Log.Info("updating nodenetworkchaos", "obj", obj)

	// pbClient, err := r.ChaosDaemonClientBuilder.BuildNodeClient(ctx,

	return ctrl.Result{Requeue: false}, fmt.Errorf("not finished")
}
