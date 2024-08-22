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
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/apache/cloudstack-go/v2/cloudstack"
	"github.com/avast/retry-go/v4"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
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
)

var retryOpts = []retry.Option{retry.Attempts(10), retry.MaxDelay(5 * time.Second), retry.DelayType(retry.FixedDelay), retry.LastErrorOnly(true)}
var waitRetryOpts = []retry.Option{retry.Attempts(20), retry.MaxDelay(30 * time.Second), retry.DelayType(retry.FixedDelay), retry.LastErrorOnly(true)}

func (impl *Impl) Apply(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	cloudstackchaos := obj.(*v1alpha1.CloudStackHostChaos)
	spec := cloudstackchaos.Spec

	client, err := utils.GetCloudStackClient(ctx, impl.Client, cloudstackchaos)
	if err != nil {
		return v1alpha1.NotInjected, errors.Wrap(err, "creating cloudstack api client")
	}

	var selector v1alpha1.CloudStackHostChaosSelector
	if err := json.Unmarshal([]byte(records[index].Id), &selector); err != nil {
		return v1alpha1.NotInjected, errors.Wrapf(err, "decoding selector: %s", records[index].Id)
	}

	params := utils.SelectorToListParams(&selector)
	params.SetOutofbandmanagementenabled(true)
	params.SetOutofbandmanagementpowerstate("On")

	resp, err := retry.DoWithData(func() (*cloudstack.ListHostsResponse, error) {
		return client.Host.ListHosts(params)
	})
	if err != nil {
		impl.Log.Error(err, "Failed to list matching hosts", "selector", records[index].Id)
		return v1alpha1.NotInjected, errors.Wrap(err, "listing hosts")
	}

	if len(resp.Hosts) == 0 {
		impl.Log.Info("No hosts matching criteria")
		return v1alpha1.Injected, nil
	}

	h := resp.Hosts[rand.Intn(len(resp.Hosts))]

	impl.Log.Info("Stopping host", "id", h.Id, "name", h.Name, "dry-run", spec.DryRun)

	if !spec.DryRun {
		params := client.OutofbandManagement.NewIssueOutOfBandManagementPowerActionParams(ActionOff, h.Id)
		if err := retry.Do(func() error {
			_, err := client.OutofbandManagement.IssueOutOfBandManagementPowerAction(params)
			return err
		}, retryOpts...); err != nil {
			impl.Log.Error(err, "Failed to stop the host", "name", h.Name)
			return v1alpha1.NotInjected, errors.Wrapf(err, "stopping host %s", h.Name)
		}

		impl.Log.Info("Stopped host", "id", h.Id, "name", h.Name)
	}

	return v1alpha1.Injected, nil
}

func (impl *Impl) Recover(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	impl.Log.Info("Starting hypervisor recovery")

	cloudstackchaos := obj.(*v1alpha1.CloudStackHostChaos)
	spec := cloudstackchaos.Spec

	client, err := utils.GetCloudStackClient(ctx, impl.Client, cloudstackchaos)
	if err != nil {
		return v1alpha1.Injected, errors.Wrap(err, "creating cloudstack api client")
	}

	var selector v1alpha1.CloudStackHostChaosSelector
	if err := json.Unmarshal([]byte(records[index].Id), &selector); err != nil {
		return v1alpha1.Injected, errors.Wrapf(err, "decoding selector: %s", records[index].Id)
	}

	params := utils.SelectorToListParams(&selector)
	params.SetOutofbandmanagementenabled(true)
	params.SetOutofbandmanagementpowerstate("Off")

	resp, err := retry.DoWithData(func() (*cloudstack.ListHostsResponse, error) {
		return client.Host.ListHosts(params)
	}, retryOpts...)
	if err != nil {
		impl.Log.Error(err, "Failed to list offline hosts", "selector", records[index].Id)
		return v1alpha1.Injected, errors.Wrap(err, "listing hosts")
	}

	for _, h := range resp.Hosts {
		impl.Log.Info("Starting host", "id", h.Id, "name", h.Name, "dry-run", spec.DryRun)

		if spec.DryRun {
			continue
		}
		if err := retry.Do(func() error {
			_, err := client.OutofbandManagement.IssueOutOfBandManagementPowerAction(client.OutofbandManagement.NewIssueOutOfBandManagementPowerActionParams(ActionOn, h.Id))
			return err
		}, retryOpts...); err != nil {
			impl.Log.Error(err, "Failed to start host", "host", h.Name)
			return v1alpha1.Injected, errors.Wrapf(err, "starting host %s", h.Name)
		}

		if err := waitForHostToBeUp(client, h.Id); err != nil {
			impl.Log.Error(err, "Host failed to start", "host", h.Name)
			return v1alpha1.Injected, err
		}

		impl.Log.Info("Started host", "id", h.Id, "name", h.Name)
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
	return retry.Do(func() error {
		vm, _, err := client.VirtualMachine.GetVirtualMachineByID(vmId)
		if err != nil {
			return errors.Wrapf(err, "failed to query status for vm %s", vmId)
		}
		if vm.State == StateRunning {
			return nil
		}

		return errors.Errorf("VM %s is not running", vmId)
	}, waitRetryOpts...)
}
func waitForHostToBeUp(client *cloudstack.CloudStackClient, hostId string) error {
	return retry.Do(func() error {
		host, _, err := client.Host.GetHostByID(hostId)
		if err != nil {
			return errors.Wrapf(err, "failed to query status for host %s", hostId)
		}
		if host.State == StateUp {
			return nil
		}
		return errors.Errorf("host %s is not up", hostId)
	}, waitRetryOpts...)
}
func (impl *Impl) getK8sNodesWhenReady(ctx context.Context) ([]v1.Node, error) {
	return retry.DoWithData(func() ([]v1.Node, error) {
		nodeList := v1.NodeList{}
		err := impl.List(ctx, &nodeList)
		if err != nil {
			return nil, errors.Wrap(err, "failed to list nodes")
		}
		unreadyNodes := []string{}
		for _, node := range nodeList.Items {
			for _, condition := range node.Status.Conditions {
				if condition.Type == v1.NodeReady && condition.Status != v1.ConditionTrue {
					unreadyNodes = append(unreadyNodes, node.Name)
					break
				}
			}
		}
		if len(unreadyNodes) > 0 {
			return nil, errors.Errorf("nodes %s not ready", strings.Join(unreadyNodes, ", "))
		}
		return nodeList.Items, nil

	}, waitRetryOpts...)
}

func (impl *Impl) uncordonK8sNodes(ctx context.Context, dryRun bool) error {
	impl.Log.Info("Will uncordon ready nodes")
	nodes, err := impl.getK8sNodesWhenReady(ctx)
	if err != nil {
		impl.Log.Error(err, "Nodes not ready")
		return nil
	}

	for _, node := range nodes {
		isApplicable := false
		for _, t := range node.Spec.Taints {
			if t.Effect == v1.TaintEffectNoSchedule && t.TimeAdded != nil {
				isApplicable = true
			}
		}
		if !isApplicable || !node.Spec.Unschedulable {
			continue
		}

		impl.Log.Info("Uncordon unschedulable node", "node", node.Name, "dryRun", dryRun)
		if dryRun {
			continue
		}

		err := retry.Do(func() error {
			nodeItem := v1.Node{}
			if err := impl.Get(ctx, types.NamespacedName{
				Namespace: node.Namespace,
				Name:      node.Name,
			}, &nodeItem); err != nil {
				return retry.Unrecoverable(errors.Wrapf(err, "could not find node %s", node.Name))
			}
			if !nodeItem.Spec.Unschedulable {
				return nil // nothing to do
			}
			nodeItem.Spec.Unschedulable = false
			return impl.Update(ctx, &nodeItem)
		}, retryOpts...)

		if err != nil {
			impl.Log.Error(err, "Failed to uncordon node", "node", node.Name)
		} else {
			impl.Log.Info("Uncordoned node", "node", node.Name)
		}
	}
	return nil
}

func (impl *Impl) startVMs(client *cloudstack.CloudStackClient, dryRun bool) error {
	impl.Log.Info("Will start stopped VMs")
	params := client.VirtualMachine.NewListVirtualMachinesParams()
	params.SetState(StateStopped)

	resp, err := retry.DoWithData(func() (*cloudstack.ListVirtualMachinesResponse, error) {
		return client.VirtualMachine.ListVirtualMachines(params)
	}, retryOpts...)
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

			if retry.Do(func() error {
				_, err := client.VirtualMachine.StartVirtualMachine(startParams)
				return err
			}, retryOpts...); err != nil {
				impl.Log.Error(err, "failed to start stopped vm", "name", vm.Name)
			}

			if err := waitForVmToBeRunning(client, vm.Id); err != nil {
				impl.Log.Error(err, "failed to wait for vm to be running", "name", vm.Name)
			} else {
				impl.Log.Info("Started VM", "id", vm.Id, "name", vm.Name)
			}

		}(vm)
	}
	wg.Wait()

	return nil
}

func (impl *Impl) destroyStuckSystemVMs(client *cloudstack.CloudStackClient, dryRun bool) error {
	impl.Log.Info("Will destroy stuck system VMs")
	params := client.SystemVM.NewListSystemVmsParams()
	params.SetState(StateStopped)

	resp, err := retry.DoWithData(func() (*cloudstack.ListSystemVmsResponse, error) {
		return client.SystemVM.ListSystemVms(params)
	})
	if err != nil {
		return err
	}

	for _, vm := range resp.SystemVms {
		impl.Log.Info("Destroying system VM", "id", vm.Id, "name", vm.Name, "dryRun", dryRun)
		if dryRun {
			continue
		}

		params := client.SystemVM.NewDestroySystemVmParams(vm.Id)
		err := retry.Do(func() error {
			_, err := client.SystemVM.DestroySystemVm(params)
			return err
		}, retryOpts...)

		if err != nil {
			impl.Log.Error(err, "Failed to destroy system vm", "id", vm.Id, "name", vm.Name)
		} else {
			impl.Log.Info("Destroyed system VM", "id", vm.Id, "name", vm.Name)
		}
	}

	return nil
}
