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

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/chaos-mesh/chaos-mesh/pkg/selector/generic"
)

// EC2Client defines the minimum client interface required for this package
type EC2Client interface {
	DescribeInstances(context.Context, *ec2.DescribeInstancesInput, ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
}

type SelectImpl struct {
	EC2Client EC2Client
}

func (impl *SelectImpl) Select(ctx context.Context, awsSelector *v1alpha1.AWSSelector) ([]*v1alpha1.AWSSelector, error) {
	if len(awsSelector.Filters) == 0 {
		return []*v1alpha1.AWSSelector{awsSelector}, nil
	}

	instances := []*v1alpha1.AWSSelector{}

	// we have filters, so we should lookup the cloud resources

	// TODO: for now, lazy load the client if not set - I'm unsure how to pass it in the main application
	if impl.EC2Client == nil {
		ec2client, err := newEc2Client(ctx, awsSelector)
		if err != nil {
			return nil, fmt.Errorf("failed to create client: %w", err)
		}
		impl.EC2Client = ec2client
	}

	result, err := impl.EC2Client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
		Filters: buildEc2Filters(awsSelector.Filters),
	})
	if err != nil {
		return instances, err
	}
	for _, r := range result.Reservations {
		// Set the Ec2Instance, and copy over the other attributes, except the filter
		instances = append(instances, &v1alpha1.AWSSelector{
			Ec2Instance: *r.Instances[0].InstanceId,
			Endpoint:    awsSelector.Endpoint,
			AWSRegion:   awsSelector.AWSRegion,
			EbsVolume:   awsSelector.EbsVolume,
			DeviceName:  awsSelector.DeviceName,
		})
	}
	mode := awsSelector.Mode
	value := awsSelector.Value

	filteredInstances, err := filterInstancesByMode(instances, mode, value)
	if err != nil {
		return nil, err
	}

	return filteredInstances, nil
}

func New() *SelectImpl {
	return &SelectImpl{}
}

func buildEc2Filters(filters []*v1alpha1.AWSFilter) []ec2types.Filter {

	ec2Filters := []ec2types.Filter{}
	for _, filter := range filters {
		ec2Filters = append(ec2Filters, ec2types.Filter{
			Name:   aws.String(filter.Name),
			Values: filter.Values,
		})
	}
	return ec2Filters
}

func newEc2Client(ctx context.Context, awsSelector *v1alpha1.AWSSelector) (*ec2.Client, error) {

	opts := []func(*awscfg.LoadOptions) error{
		awscfg.WithRegion(awsSelector.AWSRegion),
	}

	if awsSelector.Endpoint != nil {
		opts = append(opts, awscfg.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{URL: *awsSelector.Endpoint, SigningRegion: region}, nil
		})))
	}

	// TODO: no access to secret here, need to solve this
	// if awschaos.Spec.SecretName != nil {
	// 	secret := &v1.Secret{}
	// 	err := impl.Client.Get(ctx, types.NamespacedName{
	// 		Name:      *awschaos.Spec.SecretName,
	// 		Namespace: awschaos.Namespace,
	// 	}, secret)
	// 	if err != nil {
	// 		impl.Log.Error(err, "fail to get cloud secret")
	// 		return v1alpha1.NotInjected, err
	// 	}
	// 	opts = append(opts, awscfg.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
	// 		string(secret.Data["aws_access_key_id"]),
	// 		string(secret.Data["aws_secret_access_key"]),
	// 		"",
	// 	)))
	// }

	cfg, err := awscfg.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return ec2.NewFromConfig(cfg), nil
}

// filterInstancesByMode filters instances by mode from a list
func filterInstancesByMode(instances []*v1alpha1.AWSSelector, mode v1alpha1.SelectorMode, value string) ([]*v1alpha1.AWSSelector, error) {
	indexes, err := generic.FilterObjectsByMode(mode, value, len(instances))
	if err != nil {
		return nil, err
	}

	var filtered []*v1alpha1.AWSSelector

	for _, index := range indexes {
		index := index
		filtered = append(filtered, instances[index])
	}
	return filtered, nil
}
