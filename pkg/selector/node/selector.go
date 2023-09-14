package node

import (
	"context"
	"errors"

	"github.com/go-logr/logr"
	"go.uber.org/fx"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/chaos-mesh/chaos-mesh/controllers/config"
	"github.com/chaos-mesh/chaos-mesh/pkg/selector/generic"
	genericlabel "github.com/chaos-mesh/chaos-mesh/pkg/selector/generic/label"
	"github.com/chaos-mesh/chaos-mesh/pkg/selector/generic/registry"
)

var ErrNoNodeSelected = errors.New("no node is selected")

type SelectImpl struct {
	generic.Option

	c client.Client
	r client.Reader

	logger logr.Logger
}

type Params struct {
	fx.In

	Client client.Client
	Reader client.Reader `name:"no-cache"`
}

func New(params Params, logger logr.Logger) *SelectImpl {
	return &SelectImpl{
		Option: generic.Option{
			ClusterScoped:         config.ControllerCfg.ClusterScoped,
			TargetNamespace:       config.ControllerCfg.TargetNamespace,
			EnableFilterNamespace: config.ControllerCfg.EnableFilterNamespace,
		},
		c:      params.Client,
		r:      params.Reader,
		logger: logger.WithName("node-selector"),
	}
}

type Node struct {
	v1.Node
}

func (node *Node) Id() string {
	return node.Name
}

func (impl *SelectImpl) Select(ctx context.Context, selector *v1alpha1.NodeSelector) ([]*Node, error) {
	if selector == nil {
		return []*Node{}, nil
	}

	nodes, err := selectAndFilterNodes(ctx, impl.c, impl.r, selector, impl.logger)
	if err != nil {
		return nil, err
	}

	var result []*Node
	for _, node := range nodes {
		result = append(result, &Node{node})
	}

	return result, nil
}

func selectAndFilterNodes(ctx context.Context, c client.Client, r client.Reader, spec *v1alpha1.NodeSelector, logger logr.Logger) ([]v1.Node, error) {
	nodes, err := selectNodes(ctx, c, r, spec.Selector, logger)
	if err != nil {
		return nil, err
	}

	if len(nodes) == 0 {
		return nil, ErrNoNodeSelected
	}

	return nodes, nil
}

func selectNodes(ctx context.Context, c client.Client, r client.Reader, selector v1alpha1.NodeSelectorSpec, logger logr.Logger) ([]v1.Node, error) {
	selectorRegistry := newSelectorRegistry()
	selectorChain, err := registry.Parse(selectorRegistry, selector.GenericSelectorSpec, generic.Option{EnableFilterNamespace: false})
	if err != nil {
		return nil, err
	}

	return listNodes(ctx, c, r, selector, selectorChain, logger)
}

func listNodes(ctx context.Context, c client.Client, r client.Reader, selector v1alpha1.NodeSelectorSpec, selectorChain generic.SelectorChain, logger logr.Logger) ([]v1.Node, error) {
	var nodes []v1.Node

	err := selectorChain.ListObjects(c, r, func(listFunc generic.ListFunc, opts client.ListOptions) error {
		var nodeList v1.NodeList

		if err := listFunc(ctx, &nodeList, &opts); err != nil {
			return err
		}
		nodes = append(nodes, nodeList.Items...)
		return nil
	})
	if err != nil {
		return nil, err
	}

	filterList := make([]v1.Node, 0, len(nodes))
	for _, node := range nodes {
		if selectorChain.Match(&node) {
			filterList = append(filterList, node)
		}
	}

	return filterList, nil
}

func newSelectorRegistry() registry.Registry {
	return map[string]registry.SelectorFactory{
		genericlabel.Name: genericlabel.New,
	}
}
