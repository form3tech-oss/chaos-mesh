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
	"errors"
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
	impl.Log.Info("Starting hypervisor recovery")
	cloudstackchaos := obj.(*v1alpha1.CloudStackHostChaos)
	spec := cloudstackchaos.Spec

	client, err := utils.GetCloudStackClient(ctx, impl.Client, cloudstackchaos)
	if err != nil {
		return v1alpha1.Injected, fmt.Errorf("creating cloudstack api client: %w", err)
	}
	impl.Log.Info("Parsing selector")

	var selector v1alpha1.CloudStackHostChaosSelector
	if err := json.Unmarshal([]byte(records[index].Id), &selector); err != nil {
		return v1alpha1.Injected, fmt.Errorf("decoding selector: %w", err)
	}

	impl.Log.Info("Looking for hosts to recover", "selector", records[index].Id)

	params := utils.SelectorToListParams(&selector)
	params.SetOutofbandmanagementenabled(true)
	params.SetOutofbandmanagementpowerstate("Off")

	resp, err := client.Host.ListHosts(params)
	if err != nil {
		return v1alpha1.Injected, fmt.Errorf("listing hosts: %w", err)
	}

	impl.Log.Info("Found hosts to start", "len", len(resp.Hosts))

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
	if err := impl.uncordonK8sNodes(ctx, spec.DryRun); err != nil {
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
func (impl *Impl) getK8sNodesWhenReady(ctx context.Context) ([]v1.Node, error) {
	ticker := time.NewTicker(UpCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			nodes := v1.NodeList{}
			err := impl.List(ctx, &nodes)
			if err != nil {
				impl.Log.Error(err, "failed to list nodes")
				continue
			}
			unreadyNodes := []string{}
			for _, node := range nodes.Items {
				for _, condition := range node.Status.Conditions {
					if condition.Type == v1.NodeReady && condition.Status != v1.ConditionTrue {
						unreadyNodes = append(unreadyNodes, node.Name)
						break
					}
				}
			}
			if len(unreadyNodes) > 0 {
				impl.Log.Info("nodes not ready", "nodes", strings.Join(unreadyNodes, ", "))
			} else {
				return nodes.Items, nil
			}

		case <-time.After(UpCheckTimeout):
			return nil, errors.New("timed out waiting for nodes to be ready")
		}
	}
}

func (impl *Impl) uncordonK8sNodes(ctx context.Context, dryRun bool) error {
	impl.Log.Info("Will uncordon ready nodes")
	nodes, err := impl.getK8sNodesWhenReady(ctx)
	if err != nil {
		impl.Log.Error(err, "nodes not ready")
		return nil
	}
	impl.Log.Info("Found ready nodes", "len", len(nodes))

	for _, node := range nodes {
		newTaints := []v1.Taint{}
		for _, t := range node.Spec.Taints {
			if t.Effect != v1.TaintEffectNoSchedule || t.TimeAdded == nil {
				newTaints = append(newTaints, t)
			}
		}

		if len(newTaints) == len(node.Spec.Taints) {
			continue
		}

		impl.Log.Info("Removing taints", "node", node.Name, "dryRun", dryRun)
		if dryRun {
			continue
		}
		node.Spec.Taints = newTaints
		err := impl.Update(ctx, &node)
		if err != nil {
			impl.Log.Error(err, "failed to uncordon node", "node", node.Name)
		}
		impl.Log.Info("Removed taints", "node", node.Name, "dryRun", dryRun)
	}
	return nil
}

func (impl *Impl) startVMs(client *cloudstack.CloudStackClient, dryRun bool) error {
	impl.Log.Info("Will start stopped VMs")
	params := client.VirtualMachine.NewListVirtualMachinesParams()
	params.SetState(StateStopped)

	resp, err := client.VirtualMachine.ListVirtualMachines(params)
	if err != nil {
		return err
	}
	impl.Log.Info("Found vms to start", "len", len(resp.VirtualMachines))

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
	impl.Log.Info("Will destroy stuck system VMs")
	params := client.SystemVM.NewListSystemVmsParams()
	params.SetState(StateStopped)
	resp, err := client.SystemVM.ListSystemVms(params)
	if err != nil {
		return err
	}
	impl.Log.Info("Found system vms to start", "len", len(resp.SystemVms))

	for _, vm := range resp.SystemVms {
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
	return nil
}
