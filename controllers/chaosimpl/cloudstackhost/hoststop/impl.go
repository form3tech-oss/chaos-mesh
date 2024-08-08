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
	"strings"
	"sync"
	"time"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
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

	ColorBlue  = "blue"
	ColorGreen = "green"

	StateUp      = "Up"
	StateRunning = "Running"
	StateStopped = "Stopped"

	UpCheckInterval = 30 * time.Second
	UpCheckTimeout  = 10 * time.Minute

	DownCheckInterval = 30 * time.Second
	DownCheckTimeout  = 2 * time.Minute
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

	if err := impl.startVMs(client, spec.DryRun); err != nil {
		return v1alpha1.Injected, err
	}
	if err := impl.destroyStuckSystemVMs(client, spec.DryRun); err != nil {
		return v1alpha1.Injected, err
	}
	if err := impl.uncordonK8sNodes(spec.DryRun); err != nil {
		return v1alpha1.Injected, err
	}

	return v1alpha1.NotInjected, nil
}

func NewImpl(c client.Client, log logr.Logger) *Impl {
	return &Impl{
		Client: c,
		Log:    log.WithName("hvstop"),
	}
}

func waitForVmToBeRunning(client *cloudstack.CloudStackClient, vmId string) error {
	ticker := time.NewTicker(UpCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			vm, _, err := client.VirtualMachine.GetVirtualMachineByID(vmId)
			if err != nil {
				return fmt.Errorf("failed to query status for vm %s: %w", vmId, err)
			}
			if vm.State == StateRunning {
				return nil
			}
		case <-time.After(UpCheckTimeout):
			return fmt.Errorf("timed out waiting for vm %s to be up", vmId)
		}
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

func (impl *Impl) uncordonK8sNodes(dryRun bool) error {
	nodes := v1.NodeList{}
	err := impl.List(context.TODO(), &nodes)
	if err != nil {
		impl.Log.Error(err, "failed to list nodes")
		return nil
	}
	activeCluster := ""
	for _, node := range nodes.Items {
		color := ""
		if strings.Contains(node.Name, ColorGreen) {
			color = ColorGreen
		} else if strings.Contains(node.Name, ColorBlue) {
			color = ColorBlue
		} else {
			continue
		}
		if len(node.Spec.Taints) == 0 {
			activeCluster = color
			break
		}
	}
	if activeCluster == "" {
		impl.Log.Info("failed to determine active cluster, not uncordoning nodes")
		return nil
	}

	impl.Log.Info("Uncordoning nodes", "cluster", activeCluster)

	for _, node := range nodes.Items {
		if !strings.Contains(node.Name, activeCluster) {
			continue
		}
		if len(node.Spec.Taints) == 0 {
			continue
		}

		impl.Log.Info("Removing taints", "node", node.Name, "dryRun", dryRun)
		if dryRun {
			continue
		}
		node.Spec.Taints = nil
		err := impl.Update(context.TODO(), &node)
		if err != nil {
			impl.Log.Error(err, "failed to uncordon node", "node", node.Name)
		}
		impl.Log.Info("Removed taints", "node", node.Name, "dryRun", dryRun)
	}
	return nil
}

func (impl *Impl) startVMs(client *cloudstack.CloudStackClient, dryRun bool) error {
	params := client.VirtualMachine.NewListVirtualMachinesParams()
	params.SetState(StateStopped)

	resp, err := client.VirtualMachine.ListVirtualMachines(params)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	for _, vm := range resp.VirtualMachines {
		impl.Log.Info("Starting VM", "id", vm.Id, "name", vm.Name, "dryRun", dryRun)

		if dryRun {
			continue
		}
		wg.Add(1)

		go func(vm *cloudstack.VirtualMachine) {
			defer wg.Done()
			startParams := client.VirtualMachine.NewStartVirtualMachineParams(vm.Id)
			startParams.SetConsiderlasthost(true) // try to schedule to the same host

			_, err = client.VirtualMachine.StartVirtualMachine(startParams)
			if err != nil {
				impl.Log.Error(err, "failed to start stopped vm", "name", vm.Name)
			}
			err := waitForVmToBeRunning(client, vm.Id)
			if err != nil {
				impl.Log.Error(err, "failed to wait for vm to be running", "name", vm.Name)
			}
			impl.Log.Info("Started VM", "id", vm.Id, "name", vm.Name, "dryRun", dryRun)

		}(vm)
	}
	wg.Wait()

	return nil
}

func (impl *Impl) destroyStuckSystemVMs(client *cloudstack.CloudStackClient, dryRun bool) error {
	resp, err := client.SystemVM.ListSystemVms(client.SystemVM.NewListSystemVmsParams())
	if err != nil {
		return err
	}
	for _, vm := range resp.SystemVms {
		if vm.State == StateStopped {
			impl.Log.Info("Destroying system VM", "id", vm.Id, "name", vm.Name, "dryRun", dryRun)
			if dryRun {
				continue
			}
			_, err := client.SystemVM.DestroySystemVm(client.SystemVM.NewDestroySystemVmParams(vm.Id))
			if err != nil {
				return fmt.Errorf("failed to destroy system vm %s: %w", vm.Id, err)
			}
			impl.Log.Info("Destroyed system VM", "id", vm.Id, "name", vm.Name, "dryRun", dryRun)
		}
	}
	return nil
}
