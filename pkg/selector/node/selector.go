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

package node

import (
	"context"
	"fmt"

	"go.uber.org/fx"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/go-logr/logr"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
)

// const nodeSelectorName = "node"

type SelectImpl struct {
	c      client.Client
	logger logr.Logger
}

type Node struct {
	Name string
}

func (n *Node) Id() string {
	return n.Name
}

func (impl *SelectImpl) Select(ctx context.Context, nodeSelector *v1alpha1.NodeSelector) ([]*Node, error) {
	if nodeSelector == nil {
		return []*Node{}, nil
	}

	var nodes []*Node
	if len(nodeSelector.Selector.Names) > 0 {
		for _, name := range nodeSelector.Selector.Names {
			var node v1.Node
			if err := impl.c.Get(ctx, types.NamespacedName{Name: name}, &node); err != nil {
				return nil, err
			}
			nodes = append(nodes, &Node{node.Name})
		}
	}

	opts := &client.ListOptions{
		LabelSelector: labels.SelectorFromSet(nodeSelector.Selector.LabelSelectors),
		// FieldSelector: fields.SelectorFromSet(nodeSelector.Selector.FieldSelectors),
	}

	var nodeList v1.NodeList

	if len(nodeSelector.Selector.LabelSelectors) > 0 || len(nodeSelector.Selector.FieldSelectors) > 0 {
		if err := impl.c.List(ctx, &nodeList, opts); err != nil {
			return []*Node{}, fmt.Errorf("listing nodes: %w", err)
		}

		for _, node := range nodeList.Items {
			nodes = append(nodes, &Node{node.Name})
		}
	}

	return nodes, nil
}

// func (s *SelectImpl) Match(obj client.Object) bool {
// 	node := obj.(*v1.Node)
// 	for _, n := range s.nodes {
// 		if n.Name == node.Name {
// 			return true
// 		}
// 	}
// 	return false
// }

// // if both setting Nodes and NodeSelectors, the node list will be combined.
// func newNodeSelector(ctx context.Context, c client.Client, spec v1alpha1.PodSelectorSpec) (generic.Selector, error) {
// 	if len(spec.Nodes) == 0 && len(spec.NodeSelectors) == 0 {
// 		return &SelectImpl{nodes: []v1.Node{}}, nil
// 	}
// 	var nodes []v1.Node
// 	if len(spec.Nodes) > 0 {
// 		for _, name := range spec.Nodes {
// 			var node v1.Node
// 			if err := c.Get(ctx, types.NamespacedName{Name: name}, &node); err != nil {
// 				return nil, err
// 			}
// 			nodes = append(nodes, node)
// 		}
// 	}
// 	if len(spec.NodeSelectors) > 0 {
// 		var nodeList v1.NodeList
// 		if err := c.List(ctx, &nodeList, &client.ListOptions{
// 			LabelSelector: labels.SelectorFromSet(spec.NodeSelectors),
// 		}); err != nil {
// 			return nil, err
// 		}
// 		nodes = append(nodes, nodeList.Items...)
// 	}
// 	return &SelectImpl{nodes: nodes}, nil
// }

type Params struct {
	fx.In

	Client client.Client
}

func New(params Params, logger logr.Logger) *SelectImpl {
	return &SelectImpl{
		c:      params.Client,
		logger: logger.WithName("node-selector"),
	}
}
