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
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

type CloudFormation interface {
	WaitTillStackCreated(ctx context.Context, stackName string, stackParams []types.Parameter, templateBody string) (*cloudformation.DescribeStacksOutput, error)
	WaitTillStackDeleted(ctx context.Context, stackName string) error
}

// Directly using the client instead of the Interface.
type defaultCloudFormation struct {
	client *cloudformation.Client
}

func NewCloudFormation(cfg aws.Config) CloudFormation {
	_ = "STUB: not implemented"
	return *new(CloudFormation)
}

func (d *defaultCloudFormation) WaitTillStackCreated(ctx context.Context, stackName string, stackParams []types.Parameter, templateBody string) (*cloudformation.DescribeStacksOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Using the provided ctx, ctx.Done() allows wait.PollImmediateUtil to cancel

func (d *defaultCloudFormation) WaitTillStackDeleted(ctx context.Context, stackName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Using the provided ctx, ctx.Done() allows wait.PollImmediateUtil to cancel if required.
