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

package statuscheck

import (
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/chaos-mesh/chaos-mesh/controllers/config"
	"github.com/chaos-mesh/chaos-mesh/controllers/utils/builder"
	"github.com/chaos-mesh/chaos-mesh/controllers/utils/catrust"
	"github.com/chaos-mesh/chaos-mesh/controllers/utils/recorder"
)

type Params struct {
	fx.In

	Mgr             ctrl.Manager
	CertLoader      *catrust.CACertLoader
	KubeClient      client.Client
	Logger          logr.Logger
	RecorderBuilder *recorder.RecorderBuilder
}

func Bootstrap(params Params) error {
	if !config.ShouldSpawnController("statuscheck") {
		return nil
	}

	certPool, err := params.CertLoader.Load()
	if err != nil {
		return errors.Wrap(err, "loading CA certs")
	}

	eventRecorder := params.RecorderBuilder.Build("statuscheck")
	manager := NewManager(params.Logger.WithName("statuscheck-manager"), eventRecorder, certPool, newExecutor)

	return builder.Default(params.Mgr).
		For(&v1alpha1.StatusCheck{}).
		Named("statuscheck").
		Complete(NewReconciler(params.Logger.WithName("statuscheck-reconciler"), params.KubeClient, eventRecorder, manager))
}
