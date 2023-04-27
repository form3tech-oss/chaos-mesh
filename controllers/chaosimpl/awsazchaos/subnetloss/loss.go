package subnetloss

import (
	"context"
	"fmt"

	"github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/awsazchaos/awsdrclient"
	"github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/awsazchaos/ctxutil"
	"github.com/go-logr/logr"
)

// AWSSubnetsLoss will simulate loss of certain AWS subnets by setting up
// Network ACLs that completely lock down the subnets. It can either affect
// all subnets in the stack (if az == "") or subnets in a single az (if az != "")
type AWSSubnetsLoss struct {
	stack  string // Which stack to test
	az     string // Which AZ to affect, empty means all AZs
	curb   bool
	client *awsdrclient.StackScopedDRClient
	log    logr.Logger
}

func (a *AWSSubnetsLoss) String() string {
	if a.az == "" {
		return "AWS Region Loss"
	}
	return "AWS AZ Loss"
}

func NewAWSAzLoss(ctx context.Context, stack string, az string, log logr.Logger) (*AWSSubnetsLoss, error) {
	client, err := awsdrclient.New(stack, log, awsdrclient.StackScopedDRClientOptions{DryRun: false})
	if err != nil {
		return nil, err
	}

	curb := ctxutil.GetOptionalBool(ctx, ctxutil.CtxKeyCurbFlag)

	return &AWSSubnetsLoss{
		stack:  stack,
		az:     az,
		client: client,
		curb:   curb,
		log:    log,
	}, nil
}
func (a *AWSSubnetsLoss) GetSubnetToACL(ctx context.Context) (map[string]string, error) {
	return a.client.DescribeNetworkAclsForStackSubnets(ctx, a.az)
}

func (a *AWSSubnetsLoss) Start(ctx context.Context, originalSubnetToACL map[string]string) error {
	vpc, err := a.client.DescribeMainVPC(ctx)
	if err != nil {
		return err
	}

	emptyAclId, err := a.client.CreateOrGetEmptyNetworkAcl(ctx, vpc.VpcId, a.curb)
	if err != nil {
		return err
	}
	a.log.Info(fmt.Sprintf("Created empty NACL with ID: %s", emptyAclId))

	for sID, aclID := range originalSubnetToACL {
		a.log.Info(fmt.Sprintf("replacing Network ACL %s of subnet %s with block-all ACL %s", aclID, sID, emptyAclId))
		// Replace ACL and keep track of old association
		/*if err := a.client.ReplaceNetworkAclForSubnet(ctx, sID, emptyAclId); err != nil {
			// Maybe do not clean up here and depend on the next apply run to take care of it
			a.attemptCleanUp(ctx, originalSubnetToACL, emptyAclId)
			return fmt.Errorf("error replacing Network ACL for subnet (%s): %w", sID, err)
		}*/
	}

	return nil
}

func (a *AWSSubnetsLoss) Stop(ctx context.Context, originalSubnetToACL map[string]string) error {
	vpc, err := a.client.DescribeMainVPC(ctx)
	if err != nil {
		return err
	}
	emptyAclId, err := a.client.GetEmptyNetworkACL(ctx, vpc.VpcId)
	if err != nil {
		return err
	}
	err = a.cleanUp(ctx, originalSubnetToACL, emptyAclId)
	if err != nil {
		return fmt.Errorf("error cleaning up resources while stopping simulation: %w", err)
	}
	return nil
}

func (a *AWSSubnetsLoss) cleanUp(ctx context.Context, associations map[string]string, emptyACLID string) error {
	a.log.Info("cleaning up resources")

	if len(associations) > 0 {
		a.log.Info("restoring Network ACL associations")
		for subnetId, originalACLID := range associations {
			a.log.Info(fmt.Sprintf("restoring subnet (%s) to its original Network ACL (%s)", subnetId, originalACLID))
			/*if err := a.client.ReplaceNetworkAclForSubnet(ctx, subnetId, originalACLID); err != nil {
				return err
			}*/
		}
	}

	if emptyACLID != "" {
		a.log.Info(fmt.Sprintf("deleting empty Network ACL (%s)", emptyACLID))
		err := a.client.DeleteNetworkAcl(ctx, emptyACLID)
		if err != nil {
			return err
		}
	}

	a.log.Info("clean-up completed")
	return nil
}

func (a *AWSSubnetsLoss) attemptCleanUp(ctx context.Context, associations map[string]string, emptyACLID string) {
	err := a.cleanUp(ctx, associations, emptyACLID)
	if err != nil {
		a.log.Error(err, "error during clean-up")
	}
}
