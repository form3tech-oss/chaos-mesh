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
package chaosdaemon

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/chaos-mesh/chaos-mesh/pkg/bpm"
	"github.com/chaos-mesh/chaos-mesh/pkg/chaosdaemon/pb"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *DaemonServer) ApplyNodeChaos(ctx context.Context, req *pb.ApplyNodeChaosRequest) (*pb.ApplyNodeChaosResponse, error) {
	log := s.getLoggerFromContext(ctx)
	log.Info("ApplyNodeChaos", "request", req)

	pid, err := s.crClient.GetPidFromContainerID(ctx, req.ContainerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get pid from containerID error: %v", err)
	}

	log.Info("ApplyNodeChaos", "pid", pid)

	client := newCiliumClient(log, pid)
	err = client.applyPolicy(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "apply cilium policy: %v", err)
	}

	return &pb.ApplyNodeChaosResponse{InjectionId: 0}, nil
}

func (s *DaemonServer) RecoverNodeChaos(ctx context.Context, req *pb.RecoverNodeChaosRequest) (*emptypb.Empty, error) {
	log := s.getLoggerFromContext(ctx)
	log.Info("RecoverNodeChaos", "request", req)

	pid, err := s.crClient.GetPidFromContainerID(ctx, req.ContainerId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get pid from containerID error: %v", err)
	}

	log.Info("RecoverNodeChaos", "pid", pid)

	client := newCiliumClient(log, pid)
	err = client.recoverPolicy(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "delete cilium policy: %v", err)
	}

	return &emptypb.Empty{}, nil
}

type ciliumClient struct {
	log logr.Logger
	pid uint32
}

func newCiliumClient(log logr.Logger, pid uint32) ciliumClient {
	return ciliumClient{
		log,
		pid,
	}
}

func (c *ciliumClient) applyPolicy(ctx context.Context) error {
	policyPath, err := c.writePolicyFile(ctx)
	if err != nil {
		return errors.Wrap(err, "write policy file")
	}

	processBuilder := bpm.DefaultProcessBuilder("cilium", "policy", "import", policyPath).SetContext(ctx).SetNS(c.pid, bpm.MountNS).SetNS(c.pid, bpm.PidNS)
	cmd := processBuilder.Build(ctx)
	output, err := cmd.CombinedOutput()
	if err != nil {
		c.log.Error(err, "applyPolicy", "output", string(output))
		return errors.Wrap(err, "executing command: cilium policy import")
	}

	err = c.cleanupPolicyFile(ctx, policyPath)
	if err != nil {
		return errors.Wrap(err, "clean up policy file")
	}

	return nil
}

func (c *ciliumClient) recoverPolicy(ctx context.Context) error {
	processBuilder := bpm.DefaultProcessBuilder("cilium", "policy", "delete", "tech.form3.chaos.injection-id=chaos-test").SetContext(ctx).SetNS(c.pid, bpm.PidNS).SetNS(c.pid, bpm.MountNS)
	cmd := processBuilder.Build(ctx)
	output, err := cmd.CombinedOutput()
	if err != nil {
		c.log.Error(err, "recoverPolicy", "output", string(output))
		return errors.Wrap(err, "executing command: cilium policy delete")
	}
	return nil
}

func (c *ciliumClient) writePolicyFile(ctx context.Context) (string, error) {
	policy := `[{
    "endpointSelector": {
      "matchExpressions": [
        {
          "key": "k8s:app.kubernetes.io/name",
          "operator": "NotIn",
          "values": ["chaos-mesh"]
        }
      ]
    },
    "ingressDeny": [{"fromEntities": ["all"]}],
    "egressDeny": [{"toEntities": ["all"]}],
    "labels": [
      {
        "key": "tech.form3.chaos.injection-id",
        "value": "chaos-test",
        "source": "chaos-mesh"
      }
    ]
  }
]`

	policy64 := base64.StdEncoding.EncodeToString([]byte(policy))

	processBuilder := bpm.DefaultProcessBuilder("sh", "-c", fmt.Sprintf("echo '%s' | base64 -d > /tmp/chaos.json", policy64)).SetContext(ctx).SetNS(c.pid, bpm.MountNS)
	cmd := processBuilder.Build(ctx)
	output, err := cmd.CombinedOutput()
	if err != nil {
		c.log.Error(err, "writePolicyFile", "output", string(output))
		return "", errors.Wrap(err, "executing command")
	}

	return "/tmp/chaos.json", nil
}

func (c *ciliumClient) cleanupPolicyFile(ctx context.Context, policyPath string) error {
	processBuilder := bpm.DefaultProcessBuilder("/bin/rm", "-f", policyPath).SetContext(ctx).SetNS(c.pid, bpm.MountNS)
	cmd := processBuilder.Build(ctx)
	output, err := cmd.CombinedOutput()
	if err != nil {
		c.log.Error(err, "cleanupPolicyFile", "output", string(output))
		return errors.Wrap(err, "executing command")
	}

	return nil
}
