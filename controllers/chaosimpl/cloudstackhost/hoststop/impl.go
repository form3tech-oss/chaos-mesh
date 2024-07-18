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

package hoststop

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/cloudstackhost/utils"
)

type Impl struct {
	client.Client
	Log logr.Logger
}

const (
	ActionOff = "OFF"
	ActionOn  = "ON"

	StateUp      = "Up"
	StateStopped = "Stopped"

	UpCheckInterval = 30 * time.Second
	UpCheckTimeout  = 10 * time.Minute
)

func (impl *Impl) Apply(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	cloudstackchaos := obj.(*v1alpha1.CloudStackHostChaos)
	spec := cloudstackchaos.Spec

	client, err := utils.GetCloudStackClient(ctx, impl.Client, cloudstackchaos)
	if err != nil {
		return v1alpha1.NotInjected, fmt.Errorf("creating cloudstack api client: %w", err)
	}

	var selector v1alpha1.CloudStackHostChaosSelector
	if err := json.Unmarshal([]byte(records[index].Id), &selector); err != nil {
		return v1alpha1.NotInjected, fmt.Errorf("decoding selector: %w", err)
	}

	params := utils.SelectorToListParams(&selector)
	params.SetOutofbandmanagementenabled(true)
	params.SetOutofbandmanagementpowerstate("On")

	resp, err := client.Host.ListHosts(params)
	if err != nil {
		return v1alpha1.NotInjected, fmt.Errorf("listing hosts: %w", err)
	}

	if len(resp.Hosts) == 0 {
		return v1alpha1.NotInjected, fmt.Errorf("no hosts returned matching criteria")
	}

	h := resp.Hosts[rand.Intn(len(resp.Hosts))]

	impl.Log.Info("Stopping host", "id", h.Id, "name", h.Name, "dry-run", spec.DryRun)

	if !spec.DryRun {
		_, err := client.OutofbandManagement.IssueOutOfBandManagementPowerAction(client.OutofbandManagement.NewIssueOutOfBandManagementPowerActionParams(ActionOff, h.Id))
		if err != nil {
			return v1alpha1.NotInjected, fmt.Errorf("stopping host %s: %w", h.Name, err)
		}
	}

	impl.Log.Info("Stopped host", "id", h.Id, "name", h.Name, "dry-run", spec.DryRun)

	return v1alpha1.Injected, nil
}

func (impl *Impl) Recover(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	cloudstackchaos := obj.(*v1alpha1.CloudStackHostChaos)
	spec := cloudstackchaos.Spec

	client, err := utils.GetCloudStackClient(ctx, impl.Client, cloudstackchaos)
	if err != nil {
		return v1alpha1.Injected, fmt.Errorf("creating cloudstack api client: %w", err)
	}

	var selector v1alpha1.CloudStackHostChaosSelector
	if err := json.Unmarshal([]byte(records[index].Id), &selector); err != nil {
		return v1alpha1.Injected, fmt.Errorf("decoding selector: %w", err)
	}

	params := utils.SelectorToListParams(&selector)
	params.SetOutofbandmanagementenabled(true)
	params.SetOutofbandmanagementpowerstate("Off")

	resp, err := client.Host.ListHosts(params)
	if err != nil {
		return v1alpha1.Injected, fmt.Errorf("listing hosts: %w", err)
	}

	if len(resp.Hosts) == 0 {
		return v1alpha1.Injected, fmt.Errorf("no hosts returned matching criteria")
	}

	for _, h := range resp.Hosts {
		impl.Log.Info("Starting host", "id", h.Id, "name", h.Name, "dry-run", spec.DryRun)

		if !spec.DryRun {
			_, err := client.OutofbandManagement.IssueOutOfBandManagementPowerAction(client.OutofbandManagement.NewIssueOutOfBandManagementPowerActionParams(ActionOn, h.Id))
			if err != nil {
				return v1alpha1.Injected, fmt.Errorf("starting host %s: %w", h.Name, err)
			}

			if err := waitForHostToBeUp(client, h.Id); err != nil {
				return v1alpha1.Injected, err
			}

			impl.Log.Info("Started host", "id", h.Id, "name", h.Name, "dry-run", spec.DryRun)
		}
	}

	if !spec.DryRun {
		if err := impl.startStoppedVMs(client); err != nil {
			return v1alpha1.Injected, err
		}

		if err := impl.destroyStuckSystemVMs(client); err != nil {
			return v1alpha1.Injected, err
		}
	}

	return v1alpha1.NotInjected, nil
}

func NewImpl(c client.Client, log logr.Logger) *Impl {
	return &Impl{
		Client: c,
		Log:    log.WithName("hvstop"),
	}
}

func waitForHostToBeUp(client *cloudstack.CloudStackClient, hostId string) error {
	ticker := time.NewTicker(UpCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			host, _, err := client.Host.GetHostByID(hostId)
			if err != nil {
				return fmt.Errorf("failed to query status for host %s: %w", hostId, err)
			}
			if host.State == StateUp {
				return nil
			}
		case <-time.After(UpCheckTimeout):
			return fmt.Errorf("timed out waiting for host %s to be up", hostId)
		}
	}
}

func (impl *Impl) startStoppedVMs(client *cloudstack.CloudStackClient) error {
	resp, err := client.VirtualMachine.ListVirtualMachines(client.VirtualMachine.NewListVirtualMachinesParams())
	if err != nil {
		return err
	}
	for _, vm := range resp.VirtualMachines {
		if vm.State == StateStopped {
			impl.Log.Info("Starting VM", "id", vm.Id)
			_, err := client.VirtualMachine.StartVirtualMachine(client.VirtualMachine.NewStartVirtualMachineParams(vm.Id))
			if err != nil {
				return fmt.Errorf("failed to start stopped vm %s: %w", vm.Name, err)
			}
		}
	}

	return nil
}

func (impl *Impl) destroyStuckSystemVMs(client *cloudstack.CloudStackClient) error {
	resp, err := client.SystemVM.ListSystemVms(client.SystemVM.NewListSystemVmsParams())
	if err != nil {
		return err
	}
	for _, vm := range resp.SystemVms {
		if vm.State == StateStopped {
			impl.Log.Info("Destroying system VM", "id", vm.Id)
			_, err := client.SystemVM.DestroySystemVm(client.SystemVM.NewDestroySystemVmParams(vm.Id))
			if err != nil {
				return fmt.Errorf("failed to destroy system vm %s: %w", vm.Id, err)
			}
		}
	}
	return nil
}
