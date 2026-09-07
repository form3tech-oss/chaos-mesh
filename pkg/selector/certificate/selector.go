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

package certificate

import (
	"context"

	cmv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/chaos-mesh/chaos-mesh/controllers/config"
	"github.com/chaos-mesh/chaos-mesh/pkg/log"
	"github.com/chaos-mesh/chaos-mesh/pkg/selector/generic"
	genericannotation "github.com/chaos-mesh/chaos-mesh/pkg/selector/generic/annotation"
	genericfield "github.com/chaos-mesh/chaos-mesh/pkg/selector/generic/field"
	genericlabel "github.com/chaos-mesh/chaos-mesh/pkg/selector/generic/label"
	genericnamespace "github.com/chaos-mesh/chaos-mesh/pkg/selector/generic/namespace"
	"github.com/chaos-mesh/chaos-mesh/pkg/selector/generic/registry"
)

type SelectImpl struct {
	c client.Client
	r client.Reader

	generic.Option
}

type Certificate struct {
	cmv1.Certificate
}

func (cert *Certificate) Id() string {
	return (types.NamespacedName{
		Name:      cert.Name,
		Namespace: cert.Namespace,
	}).String()
}

func (impl *SelectImpl) Select(ctx context.Context, selector *v1alpha1.CertificateSelector) ([]*Certificate, error) {
	if selector == nil {
		return []*Certificate{}, nil
	}

	selectorRegistry := newSelectorRegistry()
	selectorChain, err := registry.Parse(selectorRegistry, selector.GenericSelectorSpec, generic.Option{
		ClusterScoped:         impl.ClusterScoped,
		TargetNamespace:       impl.TargetNamespace,
		EnableFilterNamespace: impl.EnableFilterNamespace,
	})
	if err != nil {
		return nil, err
	}

	certs, err := listCertificates(ctx, impl.c, impl.r, selector.GenericSelectorSpec, selectorChain, impl.EnableFilterNamespace)

	var result []*Certificate
	for _, cert := range certs {
		result = append(result, &Certificate{Certificate: cert})
	}

	return result, nil
}

func newSelectorRegistry() registry.Registry {
	return map[string]registry.SelectorFactory{
		genericlabel.Name:      genericlabel.New,
		genericnamespace.Name:  genericnamespace.New,
		genericfield.Name:      genericfield.New,
		genericannotation.Name: genericannotation.New,
	}
}

func listCertificates(ctx context.Context, c client.Client, r client.Reader, spec v1alpha1.GenericSelectorSpec,
	selectorChain generic.SelectorChain, enableFilterNamespace bool) ([]cmv1.Certificate, error) {
	var certs []cmv1.Certificate
	namespaceCheck := make(map[string]bool)
	logger, err := log.NewDefaultZapLogger()
	if err != nil {
		return certs, errors.Wrap(err, "failed to create logger")
	}
	if err := selectorChain.ListObjects(c, r,
		func(listFunc generic.ListFunc, opts client.ListOptions) error {
			var certList cmv1.CertificateList
			if len(spec.Namespaces) > 0 {
				for _, namespace := range spec.Namespaces {
					if enableFilterNamespace {
						allow, ok := namespaceCheck[namespace]
						if !ok {
							allow = genericnamespace.CheckNamespace(ctx, c, namespace, logger)
							namespaceCheck[namespace] = allow
						}
						if !allow {
							continue
						}
					}

					opts.Namespace = namespace
					if err := listFunc(ctx, &certList, &opts); err != nil {
						logger.Error(err, "list func errored", "namespace", namespace)
						return err
					}
					certs = append(certs, certList.Items...)
				}
			} else {
				// in fact, this will never happen
				if err := listFunc(ctx, &certList, &opts); err != nil {
					logger.Error(err, "list func errored")
					return err
				}
				certs = append(certs, certList.Items...)
			}
			return nil
		}); err != nil {
		return nil, err
	}

	filterCerts := make([]cmv1.Certificate, 0, len(certs))
	for _, cert := range certs {
		cert := cert
		if selectorChain.Match(&cert) {
			filterCerts = append(filterCerts, cert)
		}
	}
	return filterCerts, nil
}

type Params struct {
	fx.In

	Client client.Client
	Reader client.Reader `name:"no-cache"`
}

func New(params Params) *SelectImpl {
	return &SelectImpl{
		params.Client,
		params.Reader,
		generic.Option{
			ClusterScoped:         config.ControllerCfg.ClusterScoped,
			TargetNamespace:       config.ControllerCfg.TargetNamespace,
			EnableFilterNamespace: config.ControllerCfg.EnableFilterNamespace,
		},
	}
}
