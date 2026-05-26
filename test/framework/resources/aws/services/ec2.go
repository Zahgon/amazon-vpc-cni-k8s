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

	"github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type EC2 interface {
	DescribeInstanceType(ctx context.Context, instanceType string) ([]types.InstanceTypeInfo, error)
	DescribeInstance(ctx context.Context, instanceID string) (types.Instance, error)
	DescribeVPC(ctx context.Context, vpcID string) (*ec2.DescribeVpcsOutput, error)
	DescribeNetworkInterface(ctx context.Context, interfaceIDs []string) (*ec2.DescribeNetworkInterfacesOutput, error)
	AuthorizeSecurityGroupIngress(ctx context.Context, groupID string, protocol string, fromPort int, toPort int, cidrIP string, sourceSG bool) error
	RevokeSecurityGroupIngress(ctx context.Context, groupID string, protocol string, fromPort int, toPort int, cidrIP string, sourceSG bool) error
	AuthorizeSecurityGroupEgress(ctx context.Context, groupID string, protocol string, fromPort int, toPort int, cidrIP string) error
	RevokeSecurityGroupEgress(ctx context.Context, groupID string, protocol string, fromPort int, toPort int, cidrIP string) error
	AssociateVPCCIDRBlock(ctx context.Context, vpcId string, cidrBlock string) (*ec2.AssociateVpcCidrBlockOutput, error)
	TerminateInstance(ctx context.Context, instanceIDs []string) error
	DisAssociateVPCCIDRBlock(ctx context.Context, associationID string) error
	DescribeSubnets(ctx context.Context, subnetIDs []string) (*ec2.DescribeSubnetsOutput, error)
	CreateSubnet(ctx context.Context, cidrBlock string, vpcID string, az string) (*ec2.CreateSubnetOutput, error)
	DeleteSubnet(ctx context.Context, subnetID string) error
	DescribeRouteTables(ctx context.Context, subnetID string) (*ec2.DescribeRouteTablesOutput, error)
	DescribeRouteTablesWithVPCID(ctx context.Context, vpcID string) (*ec2.DescribeRouteTablesOutput, error)
	CreateSecurityGroup(ctx context.Context, groupName string, description string, vpcID string) (*ec2.CreateSecurityGroupOutput, error)
	DeleteSecurityGroup(ctx context.Context, groupID string) error
	AssociateRouteTableToSubnet(ctx context.Context, routeTableId string, subnetID string) error
	CreateKey(ctx context.Context, keyName string) (*ec2.CreateKeyPairOutput, error)
	DeleteKey(ctx context.Context, keyName string) error
	DescribeKey(ctx context.Context, keyName string) (*ec2.DescribeKeyPairsOutput, error)
	ModifyNetworkInterfaceSecurityGroups(ctx context.Context, securityGroupIds []string, networkInterfaceId *string) (*ec2.ModifyNetworkInterfaceAttributeOutput, error)
	DescribeAvailabilityZones(ctx context.Context) (*ec2.DescribeAvailabilityZonesOutput, error)
	CreateTags(ctx context.Context, resourceIds []string, tags []types.Tag) (*ec2.CreateTagsOutput, error)
	DeleteTags(ctx context.Context, resourceIds []string, tags []types.Tag) (*ec2.DeleteTagsOutput, error)
}

type defaultEC2 struct {
	client *ec2.Client
}

func (d *defaultEC2) DescribeInstanceType(ctx context.Context, instanceType string) ([]types.InstanceTypeInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEC2) DescribeAvailabilityZones(ctx context.Context) (*ec2.DescribeAvailabilityZonesOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEC2) ModifyNetworkInterfaceSecurityGroups(ctx context.Context, securityGroupIds []string, networkInterfaceId *string) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEC2) DescribeInstance(ctx context.Context, instanceID string) (types.Instance, error) {
	_ = "STUB: not implemented"
	return *new(types.Instance), nil
}

func (d *defaultEC2) AuthorizeSecurityGroupIngress(ctx context.Context, groupID string, protocol string, fromPort int, toPort int, cidrIP string, sourceSG bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultEC2) RevokeSecurityGroupIngress(ctx context.Context, groupID string, protocol string, fromPort int, toPort int, cidrIP string, sourceSG bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultEC2) AuthorizeSecurityGroupEgress(ctx context.Context, groupID string, protocol string, fromPort int, toPort int, cidrIP string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultEC2) RevokeSecurityGroupEgress(ctx context.Context, groupID string, protocol string, fromPort int, toPort int, cidrIP string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultEC2) DescribeNetworkInterface(ctx context.Context, interfaceIDs []string) (*ec2.DescribeNetworkInterfacesOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEC2) AssociateVPCCIDRBlock(ctx context.Context, vpcId string, cidrBlock string) (*ec2.AssociateVpcCidrBlockOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEC2) DisAssociateVPCCIDRBlock(ctx context.Context, associationID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultEC2) CreateSubnet(ctx context.Context, cidrBlock string, vpcID string, az string) (*ec2.CreateSubnetOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEC2) DescribeSubnets(ctx context.Context, subnetIDs []string) (*ec2.DescribeSubnetsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEC2) DescribeRouteTablesWithVPCID(ctx context.Context, vpcID string) (*ec2.DescribeRouteTablesOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEC2) DeleteSubnet(ctx context.Context, subnetID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultEC2) DescribeRouteTables(ctx context.Context, subnetID string) (*ec2.DescribeRouteTablesOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEC2) AssociateRouteTableToSubnet(ctx context.Context, routeTableId string, subnetID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultEC2) DeleteSecurityGroup(ctx context.Context, groupID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultEC2) CreateSecurityGroup(ctx context.Context, groupName string, description string, vpcID string) (*ec2.CreateSecurityGroupOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEC2) CreateKey(ctx context.Context, keyName string) (*ec2.CreateKeyPairOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEC2) DeleteKey(ctx context.Context, keyName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultEC2) DescribeKey(ctx context.Context, keyName string) (*ec2.DescribeKeyPairsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEC2) TerminateInstance(ctx context.Context, instanceIDs []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultEC2) DescribeVPC(ctx context.Context, vpcID string) (*ec2.DescribeVpcsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEC2) CreateTags(ctx context.Context, resourceIds []string, tags []types.Tag) (*ec2.CreateTagsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultEC2) DeleteTags(ctx context.Context, resourceIds []string, tags []types.Tag) (*ec2.DeleteTagsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewEC2(cfg aws.Config) EC2 { _ = "STUB: not implemented"; return *new(EC2) }
