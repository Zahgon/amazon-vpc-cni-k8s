// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
package vpc

import (
	"errors"

	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/utils/logger"
)

type NetworkCard struct {
	// max number of interfaces supported per card
	MaximumNetworkInterfaces int64
	// the index of current card
	NetworkCardIndex   int64
	NetworkPerformance string
}

// InstanceTypeLimits keeps track of limits for an instance type
type InstanceTypeLimits struct {
	ENILimit                int
	IPv4Limit               int
	DefaultNetworkCardIndex int
	NetworkCards            []NetworkCard
	HypervisorType          string
	IsBareMetal             bool
}

var ErrInstanceTypeNotExist = errors.New("instance type does not exist")
var ErrNoInfo = errors.New("no info on instance type due to not being publicly available")

var log = logger.Get()

func New(eniLimit int, ipv4Limit int, defaultNetworkCardIndex int, networkCards []NetworkCard,
	hypervisorType string, isBareMetalInstance bool) InstanceTypeLimits {
	_ = "STUB: not implemented"
	return *new(InstanceTypeLimits)
}

func GetENILimit(instanceType string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func GetIPv4Limit(instanceType string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func GetDefaultNetworkCardIndex(instanceType string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GetHypervisorType(instanceType string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetIsBareMetal(instanceType string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func GetNetworkCards(instanceType string) ([]NetworkCard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetInstance(instanceType string) (InstanceTypeLimits, bool) {
	_ = "STUB: not implemented"
	return *new(InstanceTypeLimits), false
}

func SetInstance(instanceType ec2types.InstanceType, eniLimit int, ipv4Limit int, defaultNetworkCardIndex int, networkCards []NetworkCard, hypervisorType ec2types.InstanceTypeHypervisor, isBareMetalInstance bool) {
	_ = "STUB: not implemented"
	return
}
