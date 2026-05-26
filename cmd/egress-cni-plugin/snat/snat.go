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

package snat

import (
	"net"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/iptableswrapper"
)

func iptRules(target, src net.IP, multicastRange, chain, comment string, useRandomFully, useHashRandom bool) [][]string {
	_ = "STUB: not implemented"
	return nil

	// Accept/ignore multicast (just because we can)
}

// SNAT

// Add NAT entries to iptables for POD egress IPv6/IPv4 traffic
func Add(ipt iptableswrapper.IPTablesIface, nodeIP, src net.IP, multicastRange, chain, comment, rndSNAT string) error {
	_ = "STUB: not implemented"
	//Defaults to `random-fully` unless a different option is explicitly set via
	//`AWS_VPC_K8S_CNI_RANDOMIZESNAT`. If the underlying iptables version doesn't support
	//'random-fully`, we will fall back to `random`.
	return nil
}

// Del removes rules added by snat
func Del(ipt iptableswrapper.IPTablesIface, src net.IP, chain, comment string) (err error) {
	_ = "STUB: not implemented"
	return nil
}
