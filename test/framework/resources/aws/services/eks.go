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

package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

type EKS interface {
	DescribeCluster(ctx context.Context, clusterName string) (*eks.DescribeClusterOutput, error)
	CreateAddon(ctx context.Context, addonInput AddonInput) (*eks.CreateAddonOutput, error)
	DescribeAddonVersions(ctx context.Context, addonInput AddonInput) (*eks.DescribeAddonVersionsOutput, error)
	DescribeAddon(ctx context.Context, addonInput AddonInput) (*eks.DescribeAddonOutput, error)
	DeleteAddon(ctx context.Context, addOnInput AddonInput) (*eks.DeleteAddonOutput, error)
	GetLatestVersion(ctx context.Context, addonInput AddonInput) (string, error)
}

type defaultEKS struct {
	client *eks.Client
}

// Internal Addon Input struct
// subset of eks.AddonInput
// used by ginkgo tests
type AddonInput struct {
	AddonName    string
	ClusterName  string
	AddonVersion string
	K8sVersion   string
}

func NewEKS(cfg aws.Config, endpoint string) (EKS, error) {
	_ = "STUB: not implemented"
	return *new(EKS), nil
}

// EKS Custom endpoint resolver needs PartitionID, SingingRegion and URL for handling STS requests.
// TODO: default to "aws" partition for now as it handled only tests. Provide option to pass partitionID.

// Fallback to default endpoint resolution for non EKS Services.

func (d *defaultEKS) CreateAddon(ctx context.Context, addonInput AddonInput) (*eks.CreateAddonOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEKS) DeleteAddon(ctx context.Context, addonInput AddonInput) (*eks.DeleteAddonOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEKS) DescribeAddonVersions(ctx context.Context, addonInput AddonInput) (*eks.DescribeAddonVersionsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEKS) DescribeAddon(ctx context.Context, addonInput AddonInput) (*eks.DescribeAddonOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEKS) DescribeCluster(ctx context.Context, clusterName string) (*eks.DescribeClusterOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEKS) GetLatestVersion(ctx context.Context, addonInput AddonInput) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
