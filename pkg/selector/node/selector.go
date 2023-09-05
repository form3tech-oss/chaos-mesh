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

package node

import (
	"context"
	"fmt"

	"go.uber.org/fx"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/go-logr/logr"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/chaos-mesh/chaos-mesh/pkg/selector/generic"
)

// const nodeSelectorName = "node"

type SelectImpl struct {
	r      client.Reader
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
		return nil, nil
	}

	selector := nodeSelector.Selector

	var nodes []*Node
	if len(selector.Names) > 0 {
		for _, name := range selector.Names {
			var node v1.Node
			if err := impl.r.Get(ctx, types.NamespacedName{Name: name}, &node); err != nil {
				return nil, err
			}
			nodes = append(nodes, &Node{node.Name})
		}
	}

	opts := &client.ListOptions{
		LabelSelector: labels.SelectorFromSet(selector.LabelSelectors),
		FieldSelector: fields.SelectorFromSet(selector.FieldSelectors),
	}

	var nodeList v1.NodeList

	if len(selector.LabelSelectors) > 0 || len(selector.FieldSelectors) > 0 {
		if err := impl.r.List(ctx, &nodeList, opts); err != nil {
			return nil, fmt.Errorf("listing nodes: %w", err)
		}

		for _, node := range nodeList.Items {
			nodes = append(nodes, &Node{node.Name})
		}
	}

	nodes, err := filterNodesByMode(nodes, nodeSelector.Mode, nodeSelector.Value)
	if err != nil {
		return nil, fmt.Errorf("filtering nodes by mode %q and value %q: %w", nodeSelector.Mode, nodeSelector.Value, err)
	}

	return nodes, nil
}

func filterNodesByMode(nodes []*Node, mode v1alpha1.SelectorMode, value string) ([]*Node, error) {
	indexes, err := generic.FilterObjectsByMode(mode, value, len(nodes))
	if err != nil {
		return nil, err
	}

	filteredNodes := make([]*Node, 0, len(nodes))

	for _, index := range indexes {
		index := index
		filteredNodes = append(filteredNodes, nodes[index])
	}
	return filteredNodes, nil
}

type Params struct {
	fx.In

	Reader client.Reader `name:"no-cache"`
}

func New(params Params, logger logr.Logger) *SelectImpl {
	return &SelectImpl{
		r:      params.Reader,
		logger: logger.WithName("node-selector"),
	}
}
