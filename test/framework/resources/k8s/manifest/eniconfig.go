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

package manifest

import (
	"github.com/aws/amazon-vpc-cni-k8s/pkg/apis/crd/v1alpha1"
)

type ENIConfigBuilder struct {
	name          string
	subnetID      string
	securityGroup []string
}

func NewENIConfigBuilder() *ENIConfigBuilder { _ = "STUB: not implemented"; return nil }

func (e *ENIConfigBuilder) Name(name string) *ENIConfigBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (e *ENIConfigBuilder) SubnetID(subnetID string) *ENIConfigBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (e *ENIConfigBuilder) SecurityGroup(securityGroup []string) *ENIConfigBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (e *ENIConfigBuilder) Build() (*v1alpha1.ENIConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
