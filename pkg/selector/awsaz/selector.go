package awsaz

import (
	"context"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
)

type SelectImpl struct{}

func (impl *SelectImpl) Select(ctx context.Context, azureSelector *v1alpha1.AWSAZSelector) ([]*v1alpha1.AWSAZSelector, error) {
	return []*v1alpha1.AWSAZSelector{azureSelector}, nil
}

func New() *SelectImpl {
	return &SelectImpl{}
}
