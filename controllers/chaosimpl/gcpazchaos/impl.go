package gcpazchaos

import (
	"context"
	"encoding/json"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	compute "google.golang.org/api/compute/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/gcpchaos/utils"
	impltypes "github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/types"
)

var _ impltypes.ChaosImpl = (*Impl)(nil)

type Impl struct {
	client.Client
	Log logr.Logger
}

func (impl *Impl) Apply(ctx context.Context, index int, records []*v1alpha1.Record, chaos v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	gcpazchaos, ok := chaos.(*v1alpha1.GCPAzChaos)
	if !ok {
		err := errors.New("chaos is not gcpchaos")
		impl.Log.Error(err, "chaos is not GCPChaos", "chaos", chaos)
		return v1alpha1.NotInjected, err
	}
	computeService, err := utils.GetComputeService(ctx, impl.Client, nil)
	if err != nil {
		impl.Log.Error(err, "fail to get the compute service")
		return v1alpha1.NotInjected, err
	}

	var selected v1alpha1.GCPAzSelector
	err = json.Unmarshal([]byte(records[index].Id), &selected)
	if err != nil {
		impl.Log.Error(err, "fail to unmarshal the selector")
		return v1alpha1.NotInjected, err
	}

	groupSizes, err := GetInstanceGroupSizes(computeService, impl.Log, selected.Project, selected.Zone)
	if err != nil {
		impl.Log.Error(err, "fail to get current instance group sizes")
		return v1alpha1.NotInjected, err
	}

	// Backup instance group sizes for later recovery.
	gcpazchaos.Status.InstanceGroupSizes = make(map[string]int64)
	for name, size := range groupSizes {
		gcpazchaos.Status.InstanceGroupSizes[name] = size
	}

	// Resize all groups to zero.
	for name := range groupSizes {
		groupSizes[name] = 0
	}

	err = ResizeInstanceGroups(computeService, impl.Log, selected.Project, selected.Zone, groupSizes)
	if err != nil {
		impl.Log.Error(err, "fail to resize instance groups")
		return v1alpha1.NotInjected, err
	}

	return v1alpha1.Injected, nil
}

func (impl *Impl) Recover(ctx context.Context, index int, records []*v1alpha1.Record, chaos v1alpha1.InnerObject) (v1alpha1.Phase, error) {
	gcpazchaos, ok := chaos.(*v1alpha1.GCPAzChaos)
	if !ok {
		err := errors.New("chaos is not gcpchaos")
		impl.Log.Error(err, "chaos is not GCPChaos", "chaos", chaos)
		return v1alpha1.Injected, err
	}
	computeService, err := utils.GetComputeService(ctx, impl.Client, nil)
	if err != nil {
		impl.Log.Error(err, "fail to get the compute service")
		return v1alpha1.Injected, err
	}

	var selected v1alpha1.GCPAzSelector
	err = json.Unmarshal([]byte(records[index].Id), &selected)
	if err != nil {
		impl.Log.Error(err, "fail to unmarshal the selector")
		return v1alpha1.Injected, err
	}

	err = ResizeInstanceGroups(computeService, impl.Log, selected.Project, selected.Zone, gcpazchaos.Status.InstanceGroupSizes)
	if err != nil {
		impl.Log.Error(err, "fail to resize instance groups")
		return v1alpha1.Injected, err
	}

	return v1alpha1.NotInjected, nil
}

func NewImpl(c client.Client, log logr.Logger) *Impl {
	return &Impl{
		Client: c,
		Log:    log.WithName("gcpazchaos"),
	}
}

func GetInstanceGroupSizes(svc *compute.Service, logger logr.Logger, project string, zone string) (map[string]int64, error) {
	instanceGroups, err := svc.InstanceGroups.List(project, zone).Do()
	if err != nil {
		logger.Error(err, "fail to list instance groups", "project", project, "zone", zone)
		return nil, err
	}

	sizes := make(map[string]int64)
	for _, group := range instanceGroups.Items {
		sizes[group.Name] = group.Size
	}
	return sizes, nil
}

func ResizeInstanceGroups(svc *compute.Service, logger logr.Logger, project string, zone string, instanceGroups map[string]int64) error {
	for name, size := range instanceGroups {
		logger.Info("pretending to resize instance group", "project", project, "zone", zone, "name", name, "size", size)
		// _, err := svc.InstanceGroupManagers.Resize(project, zone, name, size).Do()
		// if err != nil {
		// 	logger.Error(err, "fail to resize instance group", "project", project, "zone", zone, "name", name, "size", size)
		// }
	}
	return nil
}

var Module = fx.Provide(
	fx.Annotated{
		Group:  "impl",
		Target: NewImpl,
	},
)
