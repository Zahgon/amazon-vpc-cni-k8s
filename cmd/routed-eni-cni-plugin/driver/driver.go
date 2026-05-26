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

// Package driver is the CNI network driver setting up iptables, routes and rules
package driver

import (
	"net"
	"time"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/sgpp"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/ipwrapper"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/netlinkwrapper"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/nswrapper"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/procsyswrapper"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/utils/logger"
)

const (
	WAIT_INTERVAL = 50 * time.Millisecond

	//Time duration CNI waits for an IPv6 address assigned to an interface
	//to move to stable state before error'ing out.
	v6DADTimeout                = 10 * time.Second
	MAX_MAC_GENERATION_ATTEMPTS = 10
)

type VirtualInterfaceMetadata struct {
	IPAddress         *net.IPNet
	DeviceNumber      int
	RouteTable        int
	HostVethName      string
	ContainerVethName string
}

// NetworkAPIs defines network API calls
type NetworkAPIs interface {
	// SetupPodNetwork sets up pod network for normal ENI based pods
	SetupPodNetwork(vethMetadata []VirtualInterfaceMetadata, netnsPath string, mtu int, log logger.Logger) error
	// TeardownPodNetwork clean up pod network for normal ENI based pods
	TeardownPodNetwork(vethMetadata []VirtualInterfaceMetadata, log logger.Logger) error
	// SetupBranchENIPodNetwork sets up pod network for branch ENI based pods
	SetupBranchENIPodNetwork(vethMetadata VirtualInterfaceMetadata, netnsPath string, vlanID int, eniMAC string,
		subnetGW string, parentIfIndex int, mtu int, podSGEnforcingMode sgpp.EnforcingMode, log logger.Logger) error
	// TeardownBranchENIPodNetwork cleans up pod network for branch ENI based pods
	TeardownBranchENIPodNetwork(vethMetadata VirtualInterfaceMetadata, vlanID int, podSGEnforcingMode sgpp.EnforcingMode, log logger.Logger) error
}

type linuxNetwork struct {
	netLink netlinkwrapper.NetLink
	ns      nswrapper.NS
	procSys procsyswrapper.ProcSys
}

// New creates linuxNetwork object
func New() NetworkAPIs { _ = "STUB: not implemented"; return *new(NetworkAPIs) }

// createVethPairContext wraps the parameters and the method to create the
// veth pair to attach the container namespace
type createVethPairContext struct {
	contVethName string
	hostVethName string
	ipAddr       *net.IPNet
	netLink      netlinkwrapper.NetLink
	ip           ipwrapper.IP
	mtu          int
	procSys      procsyswrapper.ProcSys
	index        int
	log          logger.Logger
	hostMACAddr  net.HardwareAddr
}

func newCreateVethPairContext(contVethName string, hostVethName string, ipAddr *net.IPNet, mtu int, index int, hostMACAddr net.HardwareAddr, log logger.Logger) *createVethPairContext {
	_ = "STUB: not implemented"
	return nil
}

// run defines the closure to execute within the container's namespace to create the veth pair
func (createVethContext *createVethPairContext) run(hostNS ns.NetNS) error {
	_ = "STUB: not implemented"
	return nil
}

// Explicitly set the veth to UP state, because netlink doesn't always do that on all the platforms with net.FlagUp.
// veth won't get a link local address unless it's set to UP state.

// Explicitly set the veth to UP state, because netlink doesn't always do that on all the platforms with net.FlagUp.
// veth won't get a link local address unless it's set to UP state.

// this means it's a V6 IP address

// Enable v6 support on Container's veth interface.

// Enable v6 support on Container's lo interface inside the Pod networking namespace.

// Add a connected route to a dummy next hop (169.254.1.1 or fe80::1)
// # ip route show
// default via 169.254.1.1 dev eth0
// 169.254.1.1 dev eth0

// If Index  > 0 that means it has multiple IPs. Add IP rule + add default route to

// Add a from interface rule

// Add a default route via dummy next hop(169.254.1.1 or fe80::1). Then all outgoing traffic will be routed by this
// default route via dummy next hop (169.254.1.1 or fe80::1)

// add static ARP entry for default gateway
// we are using routed mode on the host and container need this static ARP entry to resolve its default gateway.
// IP address family is derived from the IP address passed to the function (v4 or v6)

// if IP is not IPv4 or a v4 in v6 address, it return nil

// Now that the everything has been successfully set up in the container, move the "host" end of the
// veth into the host namespace.

// SetupPodNetwork wires up linux networking for a pod's network
// we expect v4Addr and v6Addr to have correct IPAddress Family.
func (n *linuxNetwork) SetupPodNetwork(vethMetadata []VirtualInterfaceMetadata, netnsPath string, mtu int, log logger.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// TeardownPodNetwork cleanup ip rules
func (n *linuxNetwork) TeardownPodNetwork(vethMetadata []VirtualInterfaceMetadata, log logger.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// Route table ID for primary ENI was previously calculated as (Network 0, Device 0) => (0* MaxENI + 0 + 1)
// which is why we only take action if the route table is not 1

// SetupBranchENIPodNetwork sets up the network ns for pods requesting its own security group
// we expect v4Addr and v6Addr to have correct IPAddress Family.
func (n *linuxNetwork) SetupBranchENIPodNetwork(vethMetadata VirtualInterfaceMetadata, netnsPath string,
	vlanID int, eniMAC string, subnetGW string, parentIfIndex int, mtu int, podSGEnforcingMode sgpp.EnforcingMode, log logger.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// clean up any previous hostVeth ip rule recursively. (when pod with same name are recreated multiple times).
//
// per our understanding, previous we obtain vlanID from pod spec, it could be possible the vlanID is already updated when deleting old pod, thus the hostVeth been cleaned up during oldPod deletion is incorrect.
// now since we obtain vlanID from prevResult during pod deletion, we should be able to correctly purge hostVeth during pod deletion and thus don't need this logic.
// this logic is kept here for safety purpose.

// If IPv4 it returns the IP address back

// TeardownBranchENIPodNetwork tears down the vlan and corresponding ip rules.
func (n *linuxNetwork) TeardownBranchENIPodNetwork(vethMetadata VirtualInterfaceMetadata, vlanID int, _ sgpp.EnforcingMode, log logger.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// to handle the migration between different enforcingMode, we try to clean up rules under both mode since the pod might be setup with a different mode.

// setupVeth sets up veth for the pod.
func (n *linuxNetwork) setupVeth(hostVethName string, contVethName string, netnsPath string, ipAddr *net.IPNet, mtu int, log logger.Logger, index int) (netlink.Link, error) {
	_ = "STUB: not implemented"
	// Clean up if hostVeth exists.
	return *new(netlink.Link), nil
}

// For IPv6, host veth sysctls must be set to:
// 1. accept_ra=0
// 2. accept_redirects=1
// 3. forwarding=0

// Explicitly set the veth to UP state, because netlink doesn't always do that on all the platforms with net.FlagUp.
// veth won't get a link local address unless it's set to UP state.

// setupVlan sets up the vlan interface for branchENI, and configures default routes in specified route table
func (n *linuxNetwork) setupVlan(vlanID int, eniMAC string, subnetGW string, parentIfIndex int, rtTable int, log logger.Logger) (netlink.Link, error) {
	_ = "STUB: not implemented"
	return *new(netlink.Link), nil
}

// 1. clean up if vlan already exists (necessary when trunk ENI changes).

// 2. add new vlan link

// 3. Set IPv6 sysctls
//    accept_ra=0
//    accept_redirects=1
//    forwarding=0

// 4. bring up the vlan

// 5. create default routes for vlan

func (n *linuxNetwork) teardownVlan(vlanID int, log logger.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// setupIPBasedContainerRouteRules setups the routes and route rules for containers based on IP.
// traffic to container(to containerAddr) will be routed via the `main` route table.
// traffic from container(from containerAddr) will be routed via the specified rtTable.
func (n *linuxNetwork) setupIPBasedContainerRouteRules(hostVeth netlink.Link, containerAddr *net.IPNet, rtTable int, log logger.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *linuxNetwork) teardownIPBasedContainerRouteRules(containerAddr *net.IPNet, rtTable int, log logger.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// note: older version CNI sets up multiple CIDR based from container rule, so we recursively delete them to be backwards-compatible.

// routes will be automatically deleted by kernel when the hostVeth is deleted.
// we try to delete route and only log a warning even deletion failed.

// setupIIFBasedContainerRouteRules setups the routes and route rules for containers based on input network interface.
// traffic to container(iif hostVlan) will be routed via the specified rtTable.
// traffic from container(iif hostVeth) will be routed via the specified rtTable.
func (n *linuxNetwork) setupIIFBasedContainerRouteRules(hostVeth netlink.Link, containerAddr *net.IPNet, hostVlan netlink.Link, rtTable int, log logger.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *linuxNetwork) teardownIIFBasedContainerRouteRules(rtTable int, family int, log logger.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// buildRoutesForVlan builds routes required for the vlan link.
func buildRoutesForVlan(vlanTableID int, vlanIndex int, gw net.IP) []netlink.Route {
	_ = "STUB: not implemented"
	return nil
}

// Add a direct link route for the pod vlan link only.

// buildVlanLinkName builds the name for vlan link.
func buildVlanLinkName(vlanID int) string { _ = "STUB: not implemented"; return "" }

// buildVlanLink builds vlan link for the pod.
func buildVlanLink(vlanName string, vlanID int, parentIfIndex int, eniMAC string) *netlink.Vlan {
	_ = "STUB: not implemented"
	return nil
}

type MACGenerator struct {
	netlink   netlinkwrapper.NetLink
	randMACfn func() string
}

func NewMACGenerator() MACGenerator { _ = "STUB: not implemented"; return *new(MACGenerator) }

func generateRandomMAC() string { _ = "STUB: not implemented"; return "" }

// Set the local bit and unset the multicast bit

// generateUniqueRandomMAC will compare randomly generated Mac to mac addresses of veth already present in host.
func (m MACGenerator) generateUniqueRandomMAC() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
