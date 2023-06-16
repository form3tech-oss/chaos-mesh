package aws_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/ptr"
	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/chaos-mesh/chaos-mesh/pkg/selector"
	"github.com/chaos-mesh/chaos-mesh/pkg/selector/aws"
	"github.com/stretchr/testify/require"
)

type StubClient struct {
	Input  *ec2.DescribeInstancesInput
	Output *ec2.DescribeInstancesOutput
}

func (s StubClient) DescribeInstances(ctx context.Context, in *ec2.DescribeInstancesInput, opt ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error) {
	s.Input = in
	return s.Output, nil
}
func TestSelect(t *testing.T) {
	ctx := context.Background()

	sel := &v1alpha1.AWSSelector{
		Filters: []*v1alpha1.AWSFilter{{
			Name:   "tag:Stack",
			Values: []string{"staging"},
		}},
		Mode: v1alpha1.OneMode,
	}

	ec2Client := StubClient{
		Output: &ec2.DescribeInstancesOutput{
			Reservations: []ec2types.Reservation{{
				Instances: []ec2types.Instance{{
					InstanceId: ptr.String("1111"),
				}}}, {
				Instances: []ec2types.Instance{{
					InstanceId: ptr.String("2222"),
				}}}, {
				Instances: []ec2types.Instance{{
					InstanceId: ptr.String("3333"),
				}},
			}},
		},
	}
	s := selector.New(
		selector.SelectorParams{
			AWSSelector: aws.New(ec2Client),
		})

	result, err := s.Select(ctx, sel)

	require.NoError(t, err)
	require.NotNil(t, result)

	require.Len(t, result, 1)
	require.Subset(t,
		[]string{"1111", "2222", "3333"},
		[]string{result[0].(*v1alpha1.AWSSelector).Ec2Instance},
	)
}
