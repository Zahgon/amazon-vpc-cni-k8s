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
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
)

type AutoScaling interface {
	DescribeAutoScalingGroup(ctx context.Context, autoScalingGroupName string) ([]types.AutoScalingGroup, error)
	GetASGForInstance(ctx context.Context, instanceID string) (string, error)
	SetDesiredCapacity(ctx context.Context, asgName string, desired int32) error
}

// Directly using the client to interact with the service instead of an interface.
type defaultAutoScaling struct {
	client *autoscaling.Client
}

func NewAutoScaling(cfg aws.Config) AutoScaling {
	_ = "STUB: not implemented"
	return *new(AutoScaling)
}

func (d defaultAutoScaling) DescribeAutoScalingGroup(ctx context.Context, autoScalingGroupName string) ([]types.AutoScalingGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d defaultAutoScaling) GetASGForInstance(ctx context.Context, instanceID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d defaultAutoScaling) SetDesiredCapacity(ctx context.Context, asgName string, desired int32) error {
	_ = "STUB: not implemented"
	return nil
}
