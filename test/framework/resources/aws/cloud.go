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

package aws

import (
	"github.com/aws/amazon-vpc-cni-k8s/test/framework/resources/aws/services"
)

type CloudConfig struct {
	VpcID       string
	Region      string
	EKSEndpoint string
}

type Cloud interface {
	EKS() services.EKS
	EC2() services.EC2
	IAM() services.IAM
	AutoScaling() services.AutoScaling
	CloudFormation() services.CloudFormation
	CloudWatch() services.CloudWatch
}

type defaultCloud struct {
	cfg            CloudConfig
	ec2            services.EC2
	eks            services.EKS
	iam            services.IAM
	autoScaling    services.AutoScaling
	cloudFormation services.CloudFormation
	cloudWatch     services.CloudWatch
}

func NewCloud(config CloudConfig) (Cloud, error) {
	_ = "STUB: not implemented"
	return *new(Cloud), nil
}

func (c *defaultCloud) EC2() services.EC2 { _ = "STUB: not implemented"; return *new(services.EC2) }

func (c *defaultCloud) AutoScaling() services.AutoScaling {
	_ = "STUB: not implemented"
	return *new(services.AutoScaling)
}

func (c *defaultCloud) CloudFormation() services.CloudFormation {
	_ = "STUB: not implemented"
	return *new(services.CloudFormation)
}

func (c *defaultCloud) EKS() services.EKS { _ = "STUB: not implemented"; return *new(services.EKS) }

func (c *defaultCloud) IAM() services.IAM { _ = "STUB: not implemented"; return *new(services.IAM) }

func (c *defaultCloud) CloudWatch() services.CloudWatch {
	_ = "STUB: not implemented"
	return *new(services.CloudWatch)
}
