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

package main

import (
	"net"
	"time"

	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/coreos/go-iptables/iptables"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/hostipamwrapper"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/iptableswrapper"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/netlinkwrapper"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/nswrapper"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/procsyswrapper"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/utils/logger"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/vethwrapper"
)

const (
	ipv4MulticastRange = "224.0.0.0/4"
	ipv6MulticastRange = "ff00::/8"
	// WaitInterval Time duration CNI waits before next check for an IPv6 address assigned to an interface
	// to move to stable state.
	WaitInterval = 50 * time.Millisecond
	// DadTimeout Time duration CNI waits for an IPv6 address assigned to an interface
	// to move to stable state before error'ing out.
	DadTimeout = 10 * time.Second
)

// egressContext includes all info to run container ADD or DEL action
type egressContext struct {
	Procsys       procsyswrapper.ProcSys
	Ipam          hostipamwrapper.HostIpam
	Link          netlinkwrapper.NetLink
	Ns            nswrapper.NS
	NsPath        string
	ArgsIfName    string
	Veth          vethwrapper.Veth
	IPTablesIface iptableswrapper.IPTablesIface
	IptCreator    func(iptables.Protocol) (iptableswrapper.IPTablesIface, error)

	NetConf   *NetConf
	Result    *current.Result
	TmpResult *current.Result
	Log       logger.Logger

	Mtu int
	// SnatChain is the chain name for iptables rules
	SnatChain string
	// SnatComment is the comment for iptables rules
	SnatComment string
}

// NewEgressAddContext create a context for container egress traffic
func NewEgressAddContext(nsPath, ifName string) egressContext {
	_ = "STUB: not implemented"
	return *new(egressContext)
}

// NewEgressDelContext create a context for container egress traffic
func NewEgressDelContext(nsPath string) egressContext {
	_ = "STUB: not implemented"
	return *new(egressContext)
}

func (ec *egressContext) setupContainerVethV4() (*current.Interface, *current.Interface, error) {
	_ = "STUB: not implemented"
	// The IPAM result will be something like IP=192.168.3.5/24, GW=192.168.3.1.
	// What we want is really a point-to-point link but veth does not support IFF_POINTTOPOINT.
	// Next best thing would be to let it ARP but set interface to 192.168.3.5/32 and
	// add a route like "192.168.3.0/24 via 192.168.3.1 dev $ifName".
	// Unfortunately that won't work as the GW will be outside the interface's subnet.
	return nil, nil, nil
}

// Our solution is to configure the interface with 192.168.3.5/24, then delete the
// "192.168.3.0/24 dev $ifName" route that was automatically added. Then we add
// "192.168.3.1/32 dev $ifName" and "192.168.3.0/24 via 192.168.3.1 dev $ifName".
// In other words we force all traffic to ARP via the gateway except for GW itself.

// Empty veth MAC is passed

// All addresses apply to the container veth interface

// Delete the route that was automatically added

func (ec *egressContext) setupHostVethV4(vethName string) error {
	_ = "STUB: not implemented"
	// hostVeth moved namespaces and may have a new ifindex
	return nil
}

// NB: this is modified from standard ptp plugin.

// <- ptp uses SCOPE_UNIVERSE here

// <- ptp uses SCOPE_HOST here

// cmdAddEgressV4 exec necessary settings to support IPv4 egress traffic in EKS IPv6 cluster
func (ec *egressContext) cmdAddEgressV4() (err error) { _ = "STUB: not implemented"; return nil }

// NB: This uses netConf.IfName NOT args.IfName.

// add SNAT chain/rules necessary for the container IPv6 egress traffic

// Copy interfaces over to result, but not IPs.

// Pass through the previous result

// cmdDelEgressV4 exec clear the setting to support IPv4 egress traffic in EKS IPv6 cluster
func (ec *egressContext) cmdDelEgress(ipv4 bool) (err error) { _ = "STUB: not implemented"; return nil }

// without iptables ir ip6tables, chain/rules could not be removed

// DelLinkByNameAddr function deletes a link and returns IPs assigned to it, but it
// excludes IPs that are not global unicast addresses (or) private IPs. Will not work for
// our scenario as we use 169.254.0.0/16 range for v4 IPs.

//Retrieve IP addresses assigned to the link

// for IPv4 egress, IP address is a link-local IPv4 address
// for IPv6 egress, IP address is a unique-local IPv6 address
// NOTE: IsGlobalUnicast returns true for unique-local IPv6 address

// cmdAddEgressV6 exec necessary settings to support IPv6 egress traffic in EKS IPv4 cluster
func (ec *egressContext) cmdAddEgressV6() (err error) {
	_ = "STUB: not implemented"
	// Per best practice, a new veth pair is created between container ns and node ns
	// this newly created veth pair is used for container's egress IPv6 traffic
	// NOTE:
	// 1. link-local IPv6 addresses are automatically assigned to veth's both ends.
	// 2. unique-local IPv6 address allocated from host-local IPAM plugin is assigned to veth's container end only
	// 3. veth node end has no unique-local IPv6 address assigned, only link-local IPv6 address
	// 4. container IPv6 egress traffic go through node primary interface (eth0) which has an IPv6 global unicast address
	// 5. IPv6 egress traffic of all containers in a node shares node primary interface (eth0) through SNAT
	return nil
}

// first disable IPv6 on container's primary interface (eth0)

// set up SNAT in host for container IPv6 egress traffic
// following line adds an ip6tables entries to NAT for IPv6 traffic between container v6if0 and node primary ENI (eth0)

// Copy interfaces over to result, but not IPs.

// Pass through the previous result

func (ec *egressContext) disableContainerInterfaceIPv6(ifName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ec *egressContext) setupContainerIPv6Route(hostInterface, containerInterface *current.Interface) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// search for interface's link-local IPv6 address

// set up from container off-cluster IPv6 route (egress)
// all from container IPv6 traffic via host veth interface's link-local IPv6 address

// setupHostIPv6Route adds a IPv6 route for traffic destined to container/pod from external/off-cluster
func (ec *egressContext) setupHostIPv6Route(hostInterface *current.Interface, containerIPv6 net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

// set up to container return traffic route in host

func (ec *egressContext) setupContainerVethV6() (hostInterface, containerInterface *current.Interface, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Empty veth MAC is passed

// Address (IPv6 ULA address) apply to the container veth interface - v6if0

func (ec *egressContext) hostLocalIpamAdd(stdinData []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}
