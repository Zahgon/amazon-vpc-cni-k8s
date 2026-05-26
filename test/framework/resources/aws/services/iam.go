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
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type PolicyDocument struct {
	Version   string
	Statement []StatementEntry
}

type StatementEntry struct {
	Effect   string
	Action   []string
	Resource string
}

type IAM interface {
	AttachRolePolicy(ctx context.Context, policyArn string, roleName string) error
	DetachRolePolicy(ctx context.Context, policyARN string, roleName string) error
	CreatePolicy(ctx context.Context, policyName string, policyDocument string) (*iam.CreatePolicyOutput, error)
	DeletePolicy(ctx context.Context, policyARN string) error
	GetInstanceProfile(ctx context.Context, instanceProfileName string) (*iam.GetInstanceProfileOutput, error)
	GetRolePolicy(ctx context.Context, policyName string, role string) (*iam.GetRolePolicyOutput, error)
	PutRolePolicy(ctx context.Context, policyDocument string, policyName string, roleName string) error
	ListPolicies(ctx context.Context, scope string) (*iam.ListPoliciesOutput, error)
}

type defaultIAM struct {
	client *iam.Client
}

func (d *defaultIAM) AttachRolePolicy(ctx context.Context, policyARN string, roleName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultIAM) DetachRolePolicy(ctx context.Context, policyARN string, roleName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultIAM) CreatePolicy(ctx context.Context, policyName string, policyDocument string) (*iam.CreatePolicyOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultIAM) DeletePolicy(ctx context.Context, policyARN string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultIAM) GetRolePolicy(ctx context.Context, role string, policyName string) (*iam.GetRolePolicyOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultIAM) PutRolePolicy(ctx context.Context, policyDocument string, policyName string, roleName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultIAM) GetInstanceProfile(ctx context.Context, instanceProfileName string) (*iam.GetInstanceProfileOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultIAM) ListPolicies(ctx context.Context, scope string) (*iam.ListPoliciesOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewIAM(cfg aws.Config) IAM { _ = "STUB: not implemented"; return *new(IAM) }
