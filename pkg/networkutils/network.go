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

// Package networkutils is a collection of iptables and netlink functions
package networkutils

import (
	"net"
	"time"

	"github.com/coreos/go-iptables/iptables"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/sgpp"

	"golang.org/x/sys/unix"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/utils/logger"

	"github.com/vishvananda/netlink"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/iptableswrapper"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/netlinkwrapper"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/nswrapper"
)

const (
	// Vlan rule priority
	VlanRulePriority = 10

	// Local rule, needs to come after the pod ENI rules
	localRulePriority = 20

	// Rule priority for traffic destined to pod IP
	ToContainerRulePriority = 512

	// From Interface priority for multi-homed pods
	FromInterfaceRulePriority = 1

	// 513 - 1023, can be used for priority lower than fromPodRule but higher than default nonVPC CIDR rule

	// 1024 is reserved for (ip rule not to <VPC's subnet> table main)
	hostRulePriority = 1024

	// 1025 - 1534 can be used as priority lower than externalServiceIpRulePriority but higher than default nonVPC CIDR rule

	// Rule priority for traffic destined to explicit IP CIDR
	externalServiceIpRulePriority = 1535

	// Rule priority for traffic from pod
	FromPodRulePriority = 1536

	// Rule priority for traffic from primary IP on secondary ENI
	FromPrimaryIPofENIRulePriority = 32765

	// Main route table
	mainRoutingTable = unix.RT_TABLE_MAIN

	// Local route table
	localRouteTable = unix.RT_TABLE_LOCAL

	// This environment is used to specify whether an external NAT gateway will be used to provide SNAT of
	// secondary ENI IP addresses. If set to "true", the SNAT iptables rule and off-VPC ip rule will not
	// be installed and will be removed if they are already installed. Defaults to false.
	envExternalSNAT = "AWS_VPC_K8S_CNI_EXTERNALSNAT"

	// This environment is used to specify a comma-separated list of IPv4 CIDRs to exclude from SNAT. An additional rule
	// will be written to the iptables for each item. If an item is not an ipv4 range it will be skipped.
	// Defaults to empty.
	envExcludeSNATCIDRs = "AWS_VPC_K8S_CNI_EXCLUDE_SNAT_CIDRS"

	// This environment is used to specify a comma-separated list of IPv4 CIDRs that require routing lookup in
	// main routing table. An IP rule is created for each CIDR.
	envExternalServiceCIDRs = "AWS_EXTERNAL_SERVICE_CIDRS"

	// This environment is used to specify weather the SNAT rule added to iptables should randomize port allocation for
	// outgoing connections. If set to "hashrandom" the SNAT iptables rule will have the "--random" flag added to it.
	// Use "prng" if you want to use pseudo random numbers, i.e. "--random-fully".
	// Default is "prng".
	envRandomizeSNAT = "AWS_VPC_K8S_CNI_RANDOMIZESNAT"

	// envNodePortSupport is the name of environment variable that configures whether we implement support for
	// NodePorts on the primary ENI. This requires that we add additional iptables rules and loosen the kernel's
	// RPF check as described below. Defaults to true.
	envNodePortSupport = "AWS_VPC_CNI_NODE_PORT_SUPPORT"

	// envConnmark is the name of the environment variable that overrides the default connection mark, used to
	// mark traffic coming from the primary ENI so that return traffic can be forced out of the same interface.
	// Without using a mark, NodePort DNAT and our source-based routing do not work together if the target pod
	// behind the node port is not on the main ENI. In that case, the un-DNAT is done after the source-based
	// routing, resulting in the packet being sent out of the pod's ENI, when the NodePort traffic should be
	// sent over the main ENI.
	envConnmark = "AWS_VPC_K8S_CNI_CONNMARK"

	// defaultConnmark is the default value for the connmark described above. Note: the mark space is a little crowded,
	// - kube-proxy uses 0x0000c000
	// - Calico uses 0xffff0000.
	defaultConnmark = 0x80

	// envMTU gives a way to configure the MTU size for new ENIs attached. Range is from 576 to 9001.
	envMTU     = "AWS_VPC_ENI_MTU"
	defaultMTU = 9001
	minMTUv4   = 576

	// envVethPrefix is the environment variable to configure the prefix of the host side veth device names
	envVethPrefix = "AWS_VPC_K8S_CNI_VETHPREFIX"

	// envVethPrefixDefault is the default value for the veth prefix
	envVethPrefixDefault = "eni"

	// envEnIpv6Egress is the environment variable to enable IPv6 egress support on EKS v4 cluster
	envEnIpv6Egress = "ENABLE_V6_EGRESS"

	// number of retries to add a route
	maxRetryRouteAdd = 5

	retryRouteAddInterval = 5 * time.Second

	// number of attempts to find an ENI by MAC address after it is attached
	maxAttemptsLinkByMac = 5

	retryLinkByMacInterval = 3 * time.Second
)

var log = logger.Get()

// NetworkAPIs defines the host level and the ENI level network related operations
type NetworkAPIs interface {
	// SetupNodeNetwork performs node level network configuration
	SetupHostNetwork(vpcCIDRs []string, primaryMAC string, primaryAddr *net.IP, enablePodENI bool,
		v6Enabled bool) error
	// SetupENINetwork performs ENI level network configuration. Not needed on the primary ENI
	SetupENINetwork(eniIP string, eniMAC string, networkCard int, eniSubnetCIDR string, maxENIPerNIC int, isTrunkENI bool, routeTableID int, isRuleConfigured bool) error
	// UpdateHostIptablesRules updates the nat table iptables rules on the host
	UpdateHostIptablesRules(vpcCIDRs []string, primaryMAC string, primaryAddr *net.IP, v6Enabled bool) error
	CleanUpStaleAWSChains(v4Enabled, v6Enabled bool) error
	UseExternalSNAT() bool
	GetExcludeSNATCIDRs() []string
	GetExternalServiceCIDRs() []string
	GetRuleList(v6enabled bool) ([]netlink.Rule, error)
	GetRuleListBySrc(ruleList []netlink.Rule, src net.IPNet) ([]netlink.Rule, error)
	UpdateRuleListBySrc(ruleList []netlink.Rule, src net.IPNet) error
	UpdateExternalServiceIpRules(ruleList []netlink.Rule, externalIPs []string) error
	GetLinkByMac(mac string, retryInterval time.Duration) (netlink.Link, error)
	DeleteRulesBySrc(eniIP string, v6enabled bool) error
	GetRouteTableNumberForENI(networkCard int, eniIP string, deviceNumber int, maxENIsPerNetworkCard int, isV6 bool) (int, bool, error)
}

type linuxNetwork struct {
	useExternalSNAT        bool
	ipv6EgressEnabled      bool
	excludeSNATCIDRs       []string
	externalServiceCIDRs   []string
	typeOfSNAT             snatType
	nodePortSupportEnabled bool
	mtu                    int
	vethPrefix             string
	podSGEnforcingMode     sgpp.EnforcingMode

	netLink     netlinkwrapper.NetLink
	ns          nswrapper.NS
	newIptables func(IPProtocol iptables.Protocol) (iptableswrapper.IPTablesIface, error)
	mainENIMark uint32
}

type snatType uint32

const (
	sequentialSNAT snatType = iota
	randomHashSNAT
	randomPRNGSNAT
)

// New creates a linuxNetwork object
func New() NetworkAPIs { _ = "STUB: not implemented"; return *new(NetworkAPIs) }

// find out the primary interface name
func findPrimaryInterfaceName(primaryMAC string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (n *linuxNetwork) enableIPv6() (err error) { _ = "STUB: not implemented"; return nil }

func (n *linuxNetwork) SetupRuleToBlockNodeLocalV4Access() error {
	_ = "STUB: not implemented"
	return nil
}

func (n *linuxNetwork) SetupRuleToBlockNodeLocalV6Access() error {
	_ = "STUB: not implemented"
	return nil
}

// Set up a rule to block traffic directed to v4/v6 egress interface of the Pod
func (n *linuxNetwork) setupRuleToBlockNodeLocalAccess(protocol iptables.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

//Let's add the rule. Rule is either missing (or) we're not able to validate its presence.

// SetupHostNetwork performs node level network configuration
func (n *linuxNetwork) SetupHostNetwork(vpcCIDRs []string, primaryMAC string, primaryAddr *net.IP, enablePodENI bool, v6Enabled bool) error {
	_ = "STUB: not implemented"
	return nil
}

// If node port support is enabled, add a rule that will force force marked traffic out of the main ENI.  We then
// add iptables rules below that will mark traffic that needs this special treatment.  In particular NodePort
// traffic always comes in via the main ENI but response traffic would go out of the pod's assigned ENI if we
// didn't handle it specially. This is because the routing decision is done before the NodePort's DNAT is
// reversed so, to the routing table, it looks like the traffic is pod traffic instead of NodePort traffic.
// Note: With v6 PD mode support, all the pods will be behind Primary ENI of the node and so we might not even need
// to mark the packets entering via Primary ENI for NodePort support.

// If this is a restart, cleanup previous rule first

// In strict mode, packets egressing pod veth interfaces must route via the trunk ENI in order for security group
// rules to be applied. Therefore, the rule to lookup the local routing table is moved to a lower priority than VLAN rules.

// Add new rule with higher priority

// Delete the priority 0 rule

// In IPv6 strict mode, ICMPv6 packets from the gateway must lookup in the local routing table so that branch interfaces can resolve their gateway.

// UpdateHostIptablesRules updates the NAT table rules based on the VPC CIDRs configuration
func (n *linuxNetwork) UpdateHostIptablesRules(vpcCIDRs []string, primaryMAC string, primaryAddr *net.IP,
	v6Enabled bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *linuxNetwork) CleanUpStaleAWSChains(v4Enabled, v6Enabled bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Chains 1 --> x (0 indexed) will be stale

// No need to clear the chain since computeStaleIptablesRules cleans up all rules already

func (n *linuxNetwork) updateHostIptablesRules(vpcCIDRs []string, primaryMAC string, primaryAddr *net.IP,
	v6Enabled bool) error {
	_ = "STUB: not implemented"
	return nil
}

// For v6, we don't add any SNAT or Connmark rules as traffic enters and exits from the ENI it came from

func (n *linuxNetwork) buildIptablesSNATRules(vpcCIDRs []string, primaryAddr *net.IP, primaryIntf string, ipt iptableswrapper.IPTablesIface) ([]iptablesRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// build IPTABLES chain for SNAT of non-VPC outbound traffic and excluded CIDRs

// build SNAT rules for outbound non-VPC traffic

// Exclude VPC traffic from SNAT rule

// Prepare the Desired Rule for SNAT Rule for non-pod ENIs

func (n *linuxNetwork) buildIptablesConnmarkRules(vpcCIDRs []string, ipt iptableswrapper.IPTablesIface) ([]iptablesRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Force delete legacy rule: the rule was matching on "-m state --state NEW", which is
// always true for packets traversing the nat table

// Force delete existing restore mark rule so that the subsequent rule gets added to the end

// Being in the nat table, this only applies to the first packet of the connection. The mark
// will be restored in the mangle table for subsequent packets.

func (n *linuxNetwork) updateIptablesRules(iptableRules []iptablesRule, ipt iptableswrapper.IPTablesIface) error {
	_ = "STUB: not implemented"
	return nil
}

// All CIDR rules must go before the SNAT/Mark rule

func listCurrentIptablesRules(ipt iptableswrapper.IPTablesIface, table, chainPrefix string) ([]iptablesRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// To trigger ipt.Delete for stale rules

//drop action and chain name

func computeStaleIptablesRules(ipt iptableswrapper.IPTablesIface, table, chainPrefix string, newRules []iptablesRule, chains []string) ([]iptablesRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func containChainExistErr(err error) bool { _ = "STUB: not implemented"; return false }

type iptablesRule struct {
	name         string
	shouldExist  bool
	table, chain string
	rule         []string
}

func (r iptablesRule) String() string { _ = "STUB: not implemented"; return "" }

// NetLinkRuleDelAll deletes all matching route rules (instead of only first instance).
func NetLinkRuleDelAll(nl netlinkwrapper.NetLink, rule *netlink.Rule) error {
	_ = "STUB: not implemented"
	return nil
}

func ContainsNoSuchRule(err error) bool { _ = "STUB: not implemented"; return false }

func containsNoSuchRule(err error) bool { _ = "STUB: not implemented"; return false }

func IsRuleExistsError(err error) bool { _ = "STUB: not implemented"; return false }

func isRuleExistsError(err error) bool { _ = "STUB: not implemented"; return false }

// GetConfigForDebug returns the active values of the configuration env vars (for debugging purposes).
func GetConfigForDebug() map[string]interface{} { _ = "STUB: not implemented"; return nil }

// UseExternalSNAT returns whether SNAT of secondary ENI IPs should be handled with an external
// NAT gateway rather than on node. Failure to parse the setting will result in a log and the
// setting will be disabled.
func (n *linuxNetwork) UseExternalSNAT() bool { _ = "STUB: not implemented"; return false }

func useExternalSNAT() bool { _ = "STUB: not implemented"; return false }

func (n *linuxNetwork) Ipv6EgressEnabled() bool { _ = "STUB: not implemented"; return false }

func ipV6EgressEnabled() bool { _ = "STUB: not implemented"; return false }

// GetExcludeSNATCIDRs returns a list of CIDRs that should be excluded from SNAT if UseExternalSNAT is false,
// otherwise it returns an empty list.
func (n *linuxNetwork) GetExcludeSNATCIDRs() []string { _ = "STUB: not implemented"; return nil }

// GetExternalServiceCIDRs return a list of CIDRs that should always be routed to via main routing table.
func (n *linuxNetwork) GetExternalServiceCIDRs() []string { _ = "STUB: not implemented"; return nil }

func parseCIDRString(envVar string) []string { _ = "STUB: not implemented"; return nil }

// Skip IPV6 CIDRs

func typeOfSNAT() snatType { _ = "STUB: not implemented"; return *new(snatType) }

// empty means default, which is --random-fully

// prng means to use --random-fully
// note: for old versions of iptables, this will fall back to --random

// none means to disable randomisation (no flag)

// hashrandom means to use --random

// if we get to this point, the environment variable has an invalid value

func nodePortSupportEnabled() bool { _ = "STUB: not implemented"; return false }

func getBoolEnvVar(name string, defaultValue bool) bool { _ = "STUB: not implemented"; return false }

func getConnmark() uint32 { _ = "STUB: not implemented"; return 0 }

// GetLinkByMac returns linux netlink based on interface MAC
func (n *linuxNetwork) GetLinkByMac(mac string, retryInterval time.Duration) (netlink.Link, error) {
	_ = "STUB: not implemented"
	return *new(netlink.Link), nil
}

// linkByMac returns linux netlink based on interface MAC
func linkByMac(mac string, netLink netlinkwrapper.NetLink, retryInterval time.Duration) (netlink.Link, error) {
	_ = "STUB: not implemented"
	// The adapter might not be immediately available, so we perform retries
	return *new(netlink.Link), nil
}

// On AWS/VPC, the subnet gateway can always be reached at FE80:EC2::1
// https://aws.amazon.com/about-aws/whats-new/2022/11/ipv6-subnet-default-gateway-router-multiple-addresses/
func GetIPv6Gateway() net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func GetIPv4Gateway(eniSubnetCIDR *net.IPNet) net.IP {
	_ = "STUB: not implemented"
	return *new(net.IP)
}

func (n *linuxNetwork) GetRouteTableNumberForENI(networkCard int, eniIP string, deviceNumber int, maxENIPerNIC int, isV6 bool) (int, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

// SetupENINetwork adds default route to route table (eni-<eni_table>), so it does not need to be called on the primary ENI
func (n *linuxNetwork) SetupENINetwork(eniIP string, eniMAC string, networkCard int, eniSubnetCIDR string, maxENIPerNIC int, isTrunkENI bool, routeTableID int, isRuleConfigured bool) error {
	_ = "STUB: not implemented"
	return nil
}

func setupENINetwork(eniIP string, eniMac string, networkCard int, eniSubnetCIDR string, netLink netlinkwrapper.NetLink,
	retryLinkByMacInterval time.Duration, retryRouteAddInterval time.Duration, mtu int, maxENIPerNIC int, isTrunkENI bool, routeTableID int, isRuleConfigured bool) error {
	_ = "STUB: not implemented"

	// routeTableID should only be unix.RT_TABLE_MAIN for primary ENI and should never be passed to this function
	return nil
}

// Get Networking defaults

// Explicitly delete IP addresses assigned to the device before assign ENI IP.
// For IPv6, do not delete the link-local address.

// Add a direct link route for the host's ENI IP only

// Route all other traffic via the host's ENI IP

// Remove the route that default out to ENI-x out of main route table

// eniSubnetIPNet was modified by GetIPv4Gateway, so the string must be parsed again

// Add IP rule for Primary IP of the ENI

func getRouteTableNumberForENI(networkCard int, eniIP string, mask int, deviceNumber int, maxENIsPerNetworkCard int, isV6 bool, netLink netlinkwrapper.NetLink) (int, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

// Calculate the route table number based on device number and network card

// Src IP has to be the /32 or /128 address of the ENI

// Reuse the rules present on the node. This happens
// 1. When AMI is old, CNI had previously setup the ENI (Route Table, IP Rules) or
// 2. When AMI sets up the ENI Route Table Number

// This will happen on following sequence
// 1. AMI has been updated which sets up the ENI. CNI is not updated, so it overrides the ENI setup
// 2. CNI is now updated to the new version on the same node, so now we have two rules

// len(srcRuleList) == 0 → No existing rules, so we will create a new one

// For IPv6 strict mode, ICMPv6 packets from the gateway must lookup in the local routing table so that branch interfaces can resolve their gateway.
func (n *linuxNetwork) createIPv6GatewayRule() error { _ = "STUB: not implemented"; return nil }

// Rule must be deleted when not in strict mode to support transitions.

// Increment the given net.IP by one. Incrementing the last IP in an IP space (IPv4, IPV6) is undefined.
func incrementIPAddr(ip net.IP) { _ = "STUB: not implemented"; return }

// only add to the next byte if we overflowed

func getRuleList(v6enabled bool, netLink netlinkwrapper.NetLink) ([]netlink.Rule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetRuleList returns IP rules
func (n *linuxNetwork) GetRuleList(v6enabled bool) ([]netlink.Rule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetRuleListBySrc returns IP rules with matching source IP
func getRuleListBySrc(ruleList []netlink.Rule, src net.IPNet) ([]netlink.Rule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetRuleListBySrc returns IP rules with matching source IP
func (n *linuxNetwork) GetRuleListBySrc(ruleList []netlink.Rule, src net.IPNet) ([]netlink.Rule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateRuleListBySrc modify IP rules that have a matching source IP
func (n *linuxNetwork) UpdateRuleListBySrc(ruleList []netlink.Rule, src net.IPNet) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateExternalServiceIpRules reconciles existing set of IP rules for external IPs with new set
func (n *linuxNetwork) UpdateExternalServiceIpRules(ruleList []netlink.Rule, externalServiceCidrs []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete all existing rules for external service CIDRs. Note that a bulk delete would be ideal here, but
// netlink does not support bulk delete by priority, so we must iterate over rule list.

// Program new rules

// GetEthernetMTU returns the MTU value to program for ENIs. Note that the value was already validated during container initialization.
func GetEthernetMTU() int { _ = "STUB: not implemented"; return 0 }

// GetPodMTU validates the pod MTU value. If an invalid value is passed, the default is used.
func GetPodMTU(podMTU string) int { _ = "STUB: not implemented"; return 0 }

// Only IPv4 bounds can be enforced, but note that the conflist value is already validated during container initialization.

// getVethPrefixName gets the name prefix of the veth devices based on the AWS_VPC_K8S_CNI_VETHPREFIX environment variable
func getVethPrefixName() string { _ = "STUB: not implemented"; return "" }

func (n *linuxNetwork) DeleteRulesBySrc(eniIP string, isV6 bool) error {
	_ = "STUB: not implemented"
	return nil
}
