package awsdrclient

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	autoscalingTypes "github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/awsazchaos/ctxutil"
	"github.com/go-logr/logr"
)

func NewPtr[T any](val T) *T {
	return &val
}

type StackScopedDRClient struct {
	stack string

	ec2StackFilters         []ec2types.Filter
	ec2Client               *ec2.Client
	autoscalingStackFilters []autoscalingTypes.Filter
	autoscalingClient       *autoscaling.Client
	dryRun                  bool
	log                     logr.Logger
}

type StackScopedDRClientOptions struct {
	DryRun bool
}

func New(stack string, log logr.Logger, options ...StackScopedDRClientOptions) (*StackScopedDRClient, error) {
	if len(options) > 1 {
		return nil, fmt.Errorf("merging of StackScopedDRClientOptions is not supported, specify at most one options struct")
	}

	sess, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	ec2Client := ec2.NewFromConfig(sess)
	autoscalingClient := autoscaling.NewFromConfig(sess)

	dryRun := len(options) == 1 && options[0].DryRun

	return &StackScopedDRClient{
		stack: stack,
		ec2StackFilters: []ec2types.Filter{{
			Name:   NewPtr("tag:Stack"),
			Values: []string{stack},
		}},
		autoscalingStackFilters: []autoscalingTypes.Filter{{
			Name:   NewPtr("tag:Stack"),
			Values: []string{stack},
		}},
		ec2Client:         ec2Client,
		autoscalingClient: autoscalingClient,
		dryRun:            dryRun,
		log:               log,
	}, nil
}

func (a *StackScopedDRClient) DescribeMainVPC(ctx context.Context) (ec2types.Vpc, error) {
	vpcs, err := a.ec2Client.DescribeVpcs(ctx, &ec2.DescribeVpcsInput{
		Filters: append([]ec2types.Filter{
			{
				//main vpc has Name=${stack}
				Name:   NewPtr("tag:Name"),
				Values: []string{a.stack},
			},
		}, a.ec2StackFilters...),
		MaxResults: int32(1000),
	})

	if err != nil {
		return ec2types.Vpc{}, err
	}

	if len(vpcs.Vpcs) != 1 {
		return ec2types.Vpc{}, fmt.Errorf("got %d VPCs for stack %s, expected 1", len(vpcs.Vpcs), a.stack)
	}

	return vpcs.Vpcs[0], nil
}

func (a *StackScopedDRClient) DescribeSubnets(ctx context.Context, vpcId string) ([]ec2types.Subnet, error) {
	subnets, err := a.ec2Client.DescribeSubnets(ctx, &ec2.DescribeSubnetsInput{
		Filters: append([]ec2types.Filter{
			{
				Name:   NewPtr("vpc-id"),
				Values: []string{vpcId},
			},
		}, a.ec2StackFilters...),
		MaxResults: int32(1000),
	})

	if err != nil {
		return nil, err
	}
	return subnets.Subnets, nil
}

func (a *StackScopedDRClient) tagsForDRResources(simulationId string, resourceType ec2types.ResourceType) []ec2types.TagSpecification {
	return []ec2types.TagSpecification{
		{
			ResourceType: resourceType,
			Tags: []ec2types.Tag{
				{
					Key:   NewPtr("Stack"),
					Value: &a.stack,
				},
				{
					Key:   NewPtr("DisasterRecoveryResource"),
					Value: NewPtr("true"),
				},
				{
					Key:   NewPtr("DisasterRecoverySimulationId"),
					Value: NewPtr(simulationId),
				},
			},
		},
	}
}

func isSubnetIdInSubnetsAssociations(subnetId string, subnetAssociations []ec2types.NetworkAclAssociation) bool {
	for _, assoc := range subnetAssociations {
		if *assoc.SubnetId == subnetId {
			return true
		}
	}
	return false
}

func (a *StackScopedDRClient) DescribeNetworkAclsForStackSubnets(ctx context.Context) (map[string]string, error) {
	vpc, err := a.DescribeMainVPC(ctx)
	if err != nil {
		return nil, err
	}

	subnets, err := a.DescribeSubnets(ctx, *vpc.VpcId)
	if err != nil {
		return nil, err
	}

	networkAcls, err := a.ec2Client.DescribeNetworkAcls(ctx, &ec2.DescribeNetworkAclsInput{
		DryRun: false,
		Filters: []ec2types.Filter{
			{
				Name:   NewPtr("vpc-id"),
				Values: []string{*vpc.VpcId},
			},
		},
		MaxResults: int32(1000),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to describe network ACLs for ")
	}

	subnetToACL := make(map[string]string)

	// check if all subnets have a corresponding network Acl
	for _, subnet := range subnets {
		aclID := a.getACLAssociatedToSubnet(subnet, networkAcls.NetworkAcls)
		if aclID == "" {
			return nil, fmt.Errorf("subnet %s does not have a corresponding ACL", *subnet.SubnetId)
		}
		subnetToACL[*subnet.SubnetId] = aclID
	}

	return subnetToACL, nil
}

func (a *StackScopedDRClient) getACLAssociatedToSubnet(subnet ec2types.Subnet, acls []ec2types.NetworkAcl) string {
	for _, networkAcl := range acls {
		if isSubnetIdInSubnetsAssociations(*subnet.SubnetId, networkAcl.Associations) {
			a.log.Info(fmt.Sprintf("subnet %s corresponds to network acl %s", *subnet.SubnetId, *networkAcl.NetworkAclId))
			return *networkAcl.NetworkAclId
		}
	}
	return ""
}

func (a *StackScopedDRClient) deleteAssociatedNetworkAclEntries(ctx context.Context, entry AssociatedNetworkAclEntry) error {
	_, err := a.ec2Client.DeleteNetworkAclEntry(ctx, &ec2.DeleteNetworkAclEntryInput{
		Egress:       entry.Egress,
		NetworkAclId: &entry.NetworkAclId,
		RuleNumber:   entry.RuleNumber,
		DryRun:       a.dryRun,
	})
	if err != nil {

		return fmt.Errorf("error deleting DR network entry %v, cause: %w", entry, err)
	}
	a.log.Info(fmt.Sprintf("deleted network acl entry %+v", entry))
	return nil
}

type AssociatedNetworkAclEntry struct {
	Egress       bool
	NetworkAclId string
	RuleNumber   int32
}

func (a *StackScopedDRClient) DeleteNetworkAclEntries(ctx context.Context, entries []AssociatedNetworkAclEntry) error {
	var entriesThatWereNotDeleted []AssociatedNetworkAclEntry
	var errors []error

	for _, entry := range entries {
		if err := a.deleteAssociatedNetworkAclEntries(ctx, entry); err != nil {
			entriesThatWereNotDeleted = append(entriesThatWereNotDeleted, entry)
			errors = append(errors, err)
			a.log.Error(err, fmt.Sprintf("deleting network acl %v failed", entry))
		}
	}

	if len(errors) != 0 {
		return fmt.Errorf("error clearing DR network ACL entries. "+
			"The following entries have not been deleted, please do it manually: %v. Errors: %v", entriesThatWereNotDeleted, errors)
	}

	return nil
}

// CreateOrGetEmptyNetworkAcl creates a new Network ACL without only the default "deny-all"
// rules in the given VPC. It returns the ID of the Network ACL created.
func (a *StackScopedDRClient) CreateOrGetEmptyNetworkAcl(ctx context.Context, vpcId *string, curb bool) (networkAclId string, err error) {
	exisitngEmptyNetworkACL, err := a.GetEmptyNetworkACL(ctx, vpcId)
	if err != nil {
		return "", err
	}

	if exisitngEmptyNetworkACL != "" {
		a.log.Info(fmt.Sprintf("using the existing block-all network ACL %s", exisitngEmptyNetworkACL))
		return exisitngEmptyNetworkACL, nil
	}

	acl, err := a.ec2Client.CreateNetworkAcl(ctx, &ec2.CreateNetworkAclInput{
		VpcId:             vpcId,
		DryRun:            a.dryRun,
		TagSpecifications: a.tagsForDRResources(ctx.Value(ctxutil.CtxKeySimulationId).(string), ec2types.ResourceTypeNetworkAcl),
	})
	if err != nil {
		return "", fmt.Errorf("error creating Network Acl: %w", err)
	}

	a.log.Info(fmt.Sprintf("created empty Network ACL: %s", *acl.NetworkAcl.NetworkAclId))
	if curb {
		a.log.Info(fmt.Sprintf("scenario curbed; adding allow-all entries to ACL"))
		for _, isEgress := range []bool{true, false} {
			egress := isEgress
			_, err := a.ec2Client.CreateNetworkAclEntry(ctx, &ec2.CreateNetworkAclEntryInput{
				Egress:       egress,
				NetworkAclId: acl.NetworkAcl.NetworkAclId,
				Protocol:     NewPtr("-1"),
				RuleAction:   ec2types.RuleActionAllow,
				RuleNumber:   int32(100),
				CidrBlock:    NewPtr("0.0.0.0/0"),
			})
			if err != nil {
				return "", fmt.Errorf("error creating entries in ACL: %w", err)
			}
		}
	}

	return *acl.NetworkAcl.NetworkAclId, nil
}

func (a *StackScopedDRClient) GetEmptyNetworkACL(ctx context.Context, vpcId *string) (string, error) {
	existingEmptyNetworkAcl, err := a.ec2Client.DescribeNetworkAcls(ctx, &ec2.DescribeNetworkAclsInput{
		DryRun: false,
		Filters: []ec2types.Filter{
			{
				Name:   NewPtr("vpc-id"),
				Values: []string{*vpcId},
			},
			{
				Name:   NewPtr("tag:DisasterRecoveryResource"),
				Values: []string{"true"},
			},
		},
		MaxResults: int32(1),
	})
	if err != nil {
		return "", err
	}
	if len(existingEmptyNetworkAcl.NetworkAcls) == 0 {
		return "", nil
	}
	return *existingEmptyNetworkAcl.NetworkAcls[0].NetworkAclId, nil
}

// DeleteNetworkAcl deletes a Network ACL.
func (a *StackScopedDRClient) DeleteNetworkAcl(ctx context.Context, networkAclId string) error {
	_, err := a.ec2Client.DeleteNetworkAcl(ctx, &ec2.DeleteNetworkAclInput{
		NetworkAclId: NewPtr(networkAclId),
		DryRun:       a.dryRun,
	})
	if err != nil {
		return fmt.Errorf("error deleting Network Acl: %w", err)
	}

	a.log.Info(fmt.Sprintf("deleted Network ACL: %s", networkAclId))
	return nil
}

type NetworkAclAssociation struct {
	AclAssociationId string
	AclId            string
}

func (a *StackScopedDRClient) ReplaceNetworkAclForSubnet(ctx context.Context, subnetId string, newNetworkAclId string) (err error) {

	a.log.Info(fmt.Sprintf("replacing Network ACL for subnet (%s) with new acl ID: %s", subnetId, newNetworkAclId))

	acls, err := a.ec2Client.DescribeNetworkAcls(ctx, &ec2.DescribeNetworkAclsInput{
		DryRun: false,
		Filters: []ec2types.Filter{
			{
				Name: NewPtr("association.subnet-id"),
				Values: []string{
					subnetId,
				},
			},
		},
		MaxResults: int32(1000),
	})
	if err != nil {
		return err
	}
	if len(acls.NetworkAcls) != 1 {
		return fmt.Errorf("expected a single Network ACL association for subnet (%s), got %d", subnetId, len(acls.NetworkAcls))
	}

	aclIdBeforeReplace := *acls.NetworkAcls[0].NetworkAclId // Undo needs to restore old ACL
	if aclIdBeforeReplace == newNetworkAclId {
		a.log.Info(fmt.Sprintf("not replacing Network ACL for subnet (%s) which is already attached to the desired aclID (%s)",
			subnetId, newNetworkAclId))
		return nil
	}

	var associationIdToReplace *string
	for _, a := range acls.NetworkAcls[0].Associations {
		if *a.SubnetId == subnetId {
			associationIdToReplace = a.NetworkAclAssociationId
			break
		}
	}
	if associationIdToReplace == nil {
		return fmt.Errorf("couldn't find associationId for subnet (%s) and ACL (%s)", subnetId, *acls.NetworkAcls[0].NetworkAclId)
	}

	newNetworkAssociationId, err := a.ReplaceNetworkAcl(ctx, *associationIdToReplace, newNetworkAclId)
	if err != nil {
		return fmt.Errorf("unable to replace network ACL for subnet (%s): %w", subnetId, err)
	}
	a.log.Info(fmt.Sprintf("replaced Network ACL for subnet (%s) having AclId (%s) and AclAssociationId (%s) with new NetworkAclId (%s) and new AclAssociationId(%s)",
		subnetId, aclIdBeforeReplace, *associationIdToReplace, newNetworkAclId, newNetworkAssociationId))
	return nil
}

func (a *StackScopedDRClient) ReplaceNetworkAcl(ctx context.Context, networkAssociationId, networkAclId string) (newNetworkAssociationId string, err error) {
	a.log.Info(fmt.Sprintf("replacing Network ACL for association (%s) with: %s", networkAssociationId, networkAclId))
	newAssociation, err := a.ec2Client.ReplaceNetworkAclAssociation(ctx, &ec2.ReplaceNetworkAclAssociationInput{
		AssociationId: NewPtr(networkAssociationId),
		NetworkAclId:  NewPtr(networkAclId),
		DryRun:        a.dryRun,
	})
	if err != nil {
		return "", err
	}

	return *newAssociation.NewAssociationId, nil
}

/*func (a *StackScopedDRClient) DescribeAutoscalingGroups(ctx context.Context) ([]*AutoScalingGroupState, error) {
	asgs, err := a.autoscalingClient.DescribeAutoScalingGroups(ctx, &autoscaling.DescribeAutoScalingGroupsInput{
		Filters: append([]autoscalingTypes.Filter{
			{
				Name:   NewPtr("tag:Cluster"),
				Values: []string{fmt.Sprintf("%s-eks-general-blue-01", a.stack)},
			},
		}, a.autoscalingStackFilters...),
		MaxRecords: NewPtr(int32(100)),
	})
	if err != nil {
		return nil, err
	}

	results := make([]*AutoScalingGroupState, len(asgs.AutoScalingGroups))
	for i, asg := range asgs.AutoScalingGroups {
		results[i] = &AutoScalingGroupState{
			AutoScalingGroupName: *asg.AutoScalingGroupName,
			AvailabilityZones:    asg.AvailabilityZones,
			DesiredCapacity:      *asg.DesiredCapacity,
			MaxSize:              *asg.MaxSize,
			MinSize:              *asg.MinSize,
		}
	}

	return results, nil
}*/

/*func (a *StackScopedDRClient) ScaleAutoscalingGroups(ctx context.Context, desiredStates []*AutoScalingGroupState) error {
	_, errors := parallel.ExecuteInParallel(desiredStates, func(group *AutoScalingGroupState) (interface{}, error) {
		log.Infof("scaling Auto Scaling Group %s to min %d, desired %d, max %d)",
			group.AutoScalingGroupName, group.MinSize, group.DesiredCapacity, group.MaxSize)
		_, err := a.autoscalingClient.UpdateAutoScalingGroup(ctx, &autoscaling.UpdateAutoScalingGroupInput{
			AutoScalingGroupName: &group.AutoScalingGroupName,
			DesiredCapacity:      &group.DesiredCapacity,
			MaxSize:              &group.MaxSize,
			MinSize:              &group.MinSize,
		})
		if err != nil {
			return nil, fmt.Errorf("error scaling Auto Scaling Group: %w", err)
		}
		return nil, nil
	})

	if len(errors) != 0 {
		return fmt.Errorf("error(s) during scaling Auto Scaling Groups: %v", errors)
	}

	return nil
}*/
