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

// Package netlinkwrapper is a wrapper methods for the netlink package
package netlinkwrapper

import (
	"github.com/aws/amazon-vpc-cni-k8s/pkg/utils/logger"
	"github.com/vishvananda/netlink"
)

var log = logger.Get()

// NetLink wraps methods used from the vishvananda/netlink package
type NetLink interface {
	// LinkByName gets a link object given the device name
	LinkByName(name string) (netlink.Link, error)
	// LinkSetNsFd is equivalent to `ip link set $link netns $ns`
	LinkSetNsFd(link netlink.Link, fd int) error
	// ParseAddr parses an address string
	ParseAddr(s string) (*netlink.Addr, error)
	// AddrAdd is equivalent to `ip addr add $addr dev $link`
	AddrAdd(link netlink.Link, addr *netlink.Addr) error
	// AddrDel is equivalent to `ip addr del $addr dev $link`
	AddrDel(link netlink.Link, addr *netlink.Addr) error
	// AddrList is equivalent to `ip addr show `
	AddrList(link netlink.Link, family int) ([]netlink.Addr, error)
	// LinkAdd is equivalent to `ip link add`
	LinkAdd(link netlink.Link) error
	// LinkSetUp is equivalent to `ip link set $link up`
	LinkSetUp(link netlink.Link) error
	// LinkList is equivalent to: `ip link show`
	LinkList() ([]netlink.Link, error)
	// LinkSetDown is equivalent to: `ip link set $link down`
	LinkSetDown(link netlink.Link) error
	// RouteList gets a list of routes in the system.
	RouteList(link netlink.Link, family int) ([]netlink.Route, error)
	// RouteAdd will add a route to the route table
	RouteAdd(route *netlink.Route) error
	// RouteReplace will replace the route in the route table
	RouteReplace(route *netlink.Route) error
	// RouteDel is equivalent to `ip route del`
	RouteDel(route *netlink.Route) error
	// NeighAdd equivalent to: `ip neigh add ....`
	NeighAdd(neigh *netlink.Neigh) error
	// LinkDel equivalent to: `ip link del $link`
	LinkDel(link netlink.Link) error
	// NewRule creates a new empty rule
	NewRule() *netlink.Rule
	// RuleAdd is equivalent to: ip rule add
	RuleAdd(rule *netlink.Rule) error
	// RuleDel is equivalent to: ip rule del
	RuleDel(rule *netlink.Rule) error
	// RuleList is equivalent to: ip rule list
	RuleList(family int) ([]netlink.Rule, error)
	// LinkSetMTU is equivalent to `ip link set dev $link mtu $mtu`
	LinkSetMTU(link netlink.Link, mtu int) error
}

type netLink struct {
}

func retryOnErrDumpInterrupted(f func() error) error { _ = "STUB: not implemented"; return nil }

// Add small delay after first failed attempt to avoid overwhelming the kernel

// NewNetLink creates a new NetLink object
func NewNetLink() NetLink { _ = "STUB: not implemented"; return *new(NetLink) }

func (*netLink) LinkAdd(link netlink.Link) error { _ = "STUB: not implemented"; return nil }

func (*netLink) LinkByName(name string) (netlink.Link, error) {
	_ = "STUB: not implemented"
	return *new(netlink.Link), nil
}

func (*netLink) LinkSetNsFd(link netlink.Link, fd int) error { _ = "STUB: not implemented"; return nil }

func (*netLink) ParseAddr(s string) (*netlink.Addr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*netLink) AddrAdd(link netlink.Link, addr *netlink.Addr) error {
	_ = "STUB: not implemented"
	return nil
}

func (*netLink) AddrDel(link netlink.Link, addr *netlink.Addr) error {
	_ = "STUB: not implemented"
	return nil
}

func (*netLink) LinkSetUp(link netlink.Link) error { _ = "STUB: not implemented"; return nil }

func (*netLink) LinkList() ([]netlink.Link, error) { _ = "STUB: not implemented"; return nil, nil }

func (*netLink) LinkSetDown(link netlink.Link) error { _ = "STUB: not implemented"; return nil }

func (*netLink) RouteList(link netlink.Link, family int) ([]netlink.Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*netLink) RouteAdd(route *netlink.Route) error { _ = "STUB: not implemented"; return nil }

func (*netLink) RouteReplace(route *netlink.Route) error { _ = "STUB: not implemented"; return nil }

func (*netLink) RouteDel(route *netlink.Route) error { _ = "STUB: not implemented"; return nil }

func (*netLink) AddrList(link netlink.Link, family int) ([]netlink.Addr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*netLink) NeighAdd(neigh *netlink.Neigh) error { _ = "STUB: not implemented"; return nil }

func (*netLink) LinkDel(link netlink.Link) error { _ = "STUB: not implemented"; return nil }

func (*netLink) NewRule() *netlink.Rule { _ = "STUB: not implemented"; return nil }

func (*netLink) RuleAdd(rule *netlink.Rule) error { _ = "STUB: not implemented"; return nil }

func (*netLink) RuleDel(rule *netlink.Rule) error { _ = "STUB: not implemented"; return nil }

func (*netLink) RuleList(family int) ([]netlink.Rule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*netLink) LinkSetMTU(link netlink.Link, mtu int) error { _ = "STUB: not implemented"; return nil }

// IsNotExistsError returns true if the error type is syscall.ESRCH
// This helps us determine if we should ignore this error as the route
// that we want to cleanup has been deleted already routing table
func IsNotExistsError(err error) bool { _ = "STUB: not implemented"; return false }
