// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package utils

import (
	"github.com/aws/amazon-vpc-cni-k8s/test/framework"
)

const (
	CreateNodeGroupCFNTemplate        = "/testdata/amazon-eks-nodegroup.yaml"
	CreateManagedNodeGroupCFNTemplate = "/testdata/amazon-eks-managed-nodegroup.yaml"
	NodeImageIdSSMParam               = "/aws/service/eks/optimized-ami/%s/amazon-linux-2/recommended/image_id"
	ManagedNodeGroupNameLabelKey      = "eks.amazonaws.com/nodegroup"
)

type NodeGroupProperties struct {
	// Required to verify the node is up and ready
	NgLabelKey string
	NgLabelVal string
	// ASG Size
	AsgSize       int
	NodeGroupName string
	// If custom networking is set then max pod
	// will be set on Kubelet extra arguments
	IsCustomNetworkingEnabled bool
	// Subnet where the node group will be created
	Subnet       []string
	InstanceType string
	KeyPairName  string

	// optional: specify container runtime
	ContainerRuntime string

	NodeImageId string
}

type ClusterVPCConfig struct {
	PublicSubnetList   []string
	AvailZones         []string
	PublicRouteTableID string
	PrivateSubnetList  []string
}

type AWSAuthMapRole struct {
	Groups   []string `yaml:"groups"`
	RoleArn  string   `yaml:"rolearn"`
	UserName string   `yaml:"username"`
}

// Create self managed node group stack
func CreateAndWaitTillSelfManagedNGReady(f *framework.Framework, properties NodeGroupProperties) error {
	_ = "STUB: not implemented"
	return nil
}

// Update the AWS Auth Config with the Node Instance Role

// Wait till the node group have joined the cluster and are ready

func DeleteAndWaitTillSelfManagedNGStackDeleted(f *framework.Framework, properties NodeGroupProperties) error {
	_ = "STUB: not implemented"
	return nil
}

// Create managed node group stack
func CreateAndWaitTillManagedNGReady(f *framework.Framework, properties NodeGroupProperties) error {
	_ = "STUB: not implemented"
	return nil
}

// Wait till the node group have joined the cluster and are ready

func GetClusterVPCConfig(f *framework.Framework) (*ClusterVPCConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// user provided the info so we don't need to look it up

func TerminateInstances(f *framework.Framework) error { _ = "STUB: not implemented"; return nil }

// Find the ASG owning the first instance. Assumes all nodes belong to the same ASG.

// Scale the ASG to 0 so it terminates all instances through its own lifecycle

// Wait until ASG has actually finished terminating its instances before scaling

// Force-delete stale Node objects so the scheduler/aws-node don't wait on wedged kubelets.

// Wait until ASG reports `expected` instances InService.
