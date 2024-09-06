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

	HostStartingPhase = "HostStarting"
	HostStartedPhase  = "HostStarted"
	VMsStartedPhase   = "VMsStarted"
	NodesReady        = "NodesReady"
)

var retryOpts = []retry.Option{retry.Attempts(12), retry.Delay(5 * time.Second), retry.DelayType(retry.FixedDelay), retry.LastErrorOnly(true)}

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

	cloudstackchaos.Status.Host = h.Id

	vmResp, err := retry.DoWithData(func() (*cloudstack.ListVirtualMachinesResponse, error) {
		params := client.VirtualMachine.NewListVirtualMachinesParams()
		params.SetHostid(h.Id)
		return client.VirtualMachine.ListVirtualMachines(params)
	})
	if err != nil {
		return v1alpha1.NotInjected, errors.Wrapf(err, "list vms on host %s", h.Name)
	}
	vms := []string{}
	for _, vm := range vmResp.VirtualMachines {
		vms = append(vms, vm.Name)
	}
	cloudstackchaos.Status.VMs = vms

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
	cloudstackchaos := obj.(*v1alpha1.CloudStackHostChaos)

	if cloudstackchaos.Status.Host == "" {
		impl.Log.Info("Nothing to recover")
		return v1alpha1.NotInjected, nil
	}

	hostName := cloudstackchaos.Status.Host
	vms := cloudstackchaos.Status.VMs
	spec := cloudstackchaos.Spec

	record := records[index]

	impl.Log.Info("Hypervisor recovery", "host", hostName, "vms", vms, "phase", record.Phase, "dry-run", spec.DryRun)

	if spec.DryRun {
		return v1alpha1.NotInjected, nil
	}

	client, err := utils.GetCloudStackClient(ctx, impl.Client, cloudstackchaos)
	if err != nil {
		return v1alpha1.Injected, errors.Wrap(err, "creating cloudstack api client")
	}

	switch record.Phase {
	case v1alpha1.Injected:
		if err := impl.startHost(client, hostName); err != nil {
			return v1alpha1.Injected, err
		}
		return HostStartingPhase, nil

	case HostStartingPhase:
		if err := impl.ensureStartedHost(client, hostName); err != nil {
			return HostStartingPhase, err
		}

		return HostStartedPhase, nil

	case HostStartedPhase:
		if err := impl.startVMs(client, vms); err != nil {
			return HostStartedPhase, errors.Wrapf(err, "failed to start vms on host %s", hostName)
		}

		return VMsStartedPhase, nil

	case VMsStartedPhase:
		if err := impl.uncordonK8sNodes(ctx, vms); err != nil {
			// jump back to HostStartedPhase to make sure all VMs are started
			return HostStartedPhase, errors.Wrapf(err, "failed to uncordon nodes on host %s", hostName)
		}

		return NodesReady, nil

	case NodesReady:
		if err := impl.destroyStuckSystemVMs(client); err != nil {
			return NodesReady, errors.Wrap(err, "failed to destroy stuck system VMs")
		}
		return v1alpha1.NotInjected, nil

	default:
		panic("unknown recovery phase: " + record.Phase)
	}
}

func NewImpl(c client.Client, log logr.Logger) *Impl {
	return &Impl{
		Client: c,
		Log:    log.WithName("hvstop"),
	}
}

func isK8sNodeReady(node v1.Node) bool {
	for _, condition := range node.Status.Conditions {
		if condition.Type == v1.NodeReady && condition.Status != v1.ConditionTrue {
			return false
		}
	}
	return true
}
func contains(names []string, name string) bool {
	for _, n := range names {
		if n == name {
			return true
		}
	}
	return false
}

func (impl *Impl) getK8sNodes(ctx context.Context, names []string) ([]v1.Node, error) {
	return retry.DoWithData(func() ([]v1.Node, error) {
		nodeList := v1.NodeList{}
		err := impl.List(ctx, &nodeList)
		if err != nil {
			return nil, errors.Wrap(err, "failed to list nodes")
		}
		matchingNodes := []v1.Node{}
		for _, node := range nodeList.Items {
			if !contains(names, node.Name) {
				continue
			}
			matchingNodes = append(matchingNodes, node)
		}
		return matchingNodes, nil
	}, retryOpts...)
}

func (impl *Impl) uncordonK8sNodes(ctx context.Context, names []string) error {
	impl.Log.Info("Will uncordon ready nodes")
	nodes, err := impl.getK8sNodes(ctx, names)
	if err != nil {
		return err
	}
	for _, node := range nodes {
		if !isK8sNodeReady(node) {
			return errors.Errorf("node %s is not ready", node.Name)
		}
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

		impl.Log.Info("Uncordon unschedulable node", "node", node.Name)

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

func (impl *Impl) startHost(client *cloudstack.CloudStackClient, hostName string) error {
	host, err := retry.DoWithData(func() (*cloudstack.Host, error) {
		host, _, err := client.Host.GetHostByName(hostName)
		return host, err
	}, retryOpts...)
	if err != nil {
		return err
	}
	return retry.Do(func() error {
		_, err := client.OutofbandManagement.IssueOutOfBandManagementPowerAction(client.OutofbandManagement.NewIssueOutOfBandManagementPowerActionParams(ActionOn, host.Id))
		return err
	}, retryOpts...)
}

func (impl *Impl) ensureStartedHost(client *cloudstack.CloudStackClient, hostName string) error {
	host, err := retry.DoWithData(func() (*cloudstack.Host, error) {
		host, _, err := client.Host.GetHostByName(hostName)
		return host, err
	}, retryOpts...)
	if err != nil {
		return err
	}

	if host.State != StateUp {
		return errors.Errorf("host %s is not up", hostName)
	}
	return nil
}

func (impl *Impl) startVMs(client *cloudstack.CloudStackClient, names []string) error {
	impl.Log.Info("Will start stopped VMs")
	params := client.VirtualMachine.NewListVirtualMachinesParams()
	params.SetState(StateStopped)

	resp, err := retry.DoWithData(func() (*cloudstack.ListVirtualMachinesResponse, error) {
		return client.VirtualMachine.ListVirtualMachines(params)
	}, retryOpts...)
	if err != nil {
		return errors.Wrap(err, "Failed to list stopped VMs")
	}

	for _, vm := range resp.VirtualMachines {
		if !contains(names, vm.Name) {
			continue
		}

		impl.Log.Info("Starting VM", "id", vm.Id, "name", vm.Name)

		startParams := client.VirtualMachine.NewStartVirtualMachineParams(vm.Id)
		startParams.SetConsiderlasthost(true) // try to schedule to the same host

		if err := retry.Do(func() error {
			_, err := client.VirtualMachine.StartVirtualMachine(startParams)
			return err
		}, retryOpts...); err != nil {
			return errors.Wrapf(err, "failed to start stopped vm %s", vm.Name)
		}
		impl.Log.Info("Started VM", "id", vm.Id, "name", vm.Name)
	}

	return nil
}

func (impl *Impl) destroyStuckSystemVMs(client *cloudstack.CloudStackClient) error {
	impl.Log.Info("Will destroy stuck system VMs")
	params := client.SystemVM.NewListSystemVmsParams()
	params.SetState(StateStopped)

	resp, err := retry.DoWithData(func() (*cloudstack.ListSystemVmsResponse, error) {
		return client.SystemVM.ListSystemVms(params)
	}, retryOpts...)
	if err != nil {
		return errors.Wrap(err, "Failed to list system VMs")
	}

	for _, vm := range resp.SystemVms {
		impl.Log.Info("Destroying system VM", "id", vm.Id, "name", vm.Name)

		params := client.SystemVM.NewDestroySystemVmParams(vm.Id)
		err := retry.Do(func() error {
			_, err := client.SystemVM.DestroySystemVm(params)
			return err
		}, retryOpts...)

		if err != nil {
			return errors.Wrapf(err, "failed to destroy system vm %s", vm.Name)
		}
		impl.Log.Info("Destroyed system VM", "id", vm.Id, "name", vm.Name)
	}

	return nil
}
