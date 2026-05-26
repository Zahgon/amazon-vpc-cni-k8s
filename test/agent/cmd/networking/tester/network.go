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

package tester

import (
	"net"

	"github.com/aws/amazon-vpc-cni-k8s/test/agent/pkg/input"
	"github.com/vishvananda/netlink"
)

// TestNetworkingSetupForRegularPod tests networking set by the CNI Plugin for a list of Pod is as
// expected
func TestNetworkingSetupForRegularPod(podNetworkingValidationInput input.PodNetworkingValidationInput) []error {
	_ = "STUB: not implemented"
	return nil
}

// Get the list of IP rules

// Do validation for each Pod and if validation fails instead of failing
// entire test add errors to a list for all the failing Pods

// For each Pod validate the Pod networking

// For each pod categorize the rules into rules for the main route table
// and for non main route table

// Get the veth pair for pod in host network namespace

// Validate MTU value if it is set to true

// Verify IP Link for the Pod is UP

// Get the IP Rules to/from the Pod IP and categorize them into main and non-main table rules

// Both Pod with IP from Primary and Secondary ENI will have 1 rule for main route table

// Verify main table route for pod IP go through the veth pair when destination is Pod IP

// Verify that the link index for the route is the same as the veth pair index

// Pod with IP from Secondary ENI will have additional rule for destination to each
// VPC Cidr block

// Finally validate that the route table for secondary ENI has the right routes

// Route 1 should route all traffic through Gateway via Secondary ENI

// Route 2 should route all traffic intended for Gateway IP through Secondary ENI

// TODO: validate iptables rules get setup correctly

// TestNetworkingSetupForPods using security groups
func TestNetworkingSetupForPodsUsingSecurityGroup(podNetworkingValidationInput input.PodNetworkingValidationInput) []error {
	_ = "STUB: not implemented"
	return nil
}

// Get the list of IP rules

// Get the veth pair for pod in host network namespace

// Check if branch ENI exists for given pod

// Check if branchENI is in UP state

// Validate MTU value if it is set to true

// Verify IP Link for the Pod is UP

// TestNetworkTearedDownForRegularPods test pod networking is correctly teared down by the CNI Plugin
// The test assumes that the IP assigned to the older Pod is not assigned to a new Pod while this test
// is being executed
func TestNetworkTearedDownForRegularPods(podNetworkingValidationInput input.PodNetworkingValidationInput) []error {
	_ = "STUB: not implemented"
	return nil
}

// Get the list of IP rules

// Make sure the veth pair doesn't exist anymore

// Make sure there's no more rules either to or from the Pod's IPv4 Address

// Test the next pod if even a single leaked rule if found

// Make sure there's no route to Pod IP Address

// TestNetworkingForPods using security groups is teared down correctly
func TestNetworkTearedDownForPodsUsingSecurityGroup(podNetworkingValidationInput input.PodNetworkingValidationInput) []error {
	_ = "STUB: not implemented"
	return nil
}

// Get the list of IP rules

// Check if branchENI's are cleanup

// Make sure the veth pair doesn't exist anymore

// Found leaked veth pair

// check if vlanTable rule exists for leaked veth pair

func isRuleToOrFromIP(rule netlink.Rule, ip net.IP) bool { _ = "STUB: not implemented"; return false }

func getHostVethPairName(input input.Pod, vethPrefix string) string {
	_ = "STUB: not implemented"
	return ""
}
