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

// IPv4 network packet verifier
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/netlinkwrapper"
	"github.com/vishvananda/netlink"
)

const (
	shortDescription = "packet-verifier"
	longDescription  = "Packet verifier fetches corresponding interfaces and validates the packets."
)

var (
	// version string populated during build.
	version = "unknown"

	// ip to monitor on the interfaces
	ipAddress  string
	receiverIP string
	device     string

	// vlan ID to monitor on the interfaces
	vlanIDToMonitor int

	// pcap parameters (requires libpcap-devel to be installed on the host)
	snapshotLen int32 = 1024
	promiscuous       = false
	timeout           = 30 * time.Second
)

// eniConfig details regarding ENIs
type eniConfig struct {
	name                string
	shouldCheckSrc      bool
	shouldVerifyVlanTag bool
}

func main() {
	// list of enis to monitor
	var enis []eniConfig

	fmt.Print("Verifying packet flow...\n")

	helpFlag := flag.Bool("help", false, "displays usage information")
	versionFlag := flag.Bool("version", false, "displays version information")
	flag.StringVar(&ipAddress, "ip-to-monitor", "", "pod ip to monitor.")
	flag.StringVar(&receiverIP, "receiver-ip", "", "other IP that interacts with the pod.")
	flag.IntVar(&vlanIDToMonitor, "vlanid-to-monitor", 0, "pod vlan id to monitor.")
	flag.StringVar(&device, "host-device", "eth0", "host device of the node.")

	flag.Usage = printUsage

	// Parse command line flags.
	flag.Parse()

	if *helpFlag {
		printUsage()
		os.Exit(0)
	}

	if *versionFlag {
		printVersion()
		os.Exit(0)
	}

	if ipAddress == "" {
		fmt.Println("ip-to-monitor can't be empty")
		os.Exit(1)
	}
	ipToMonitor := net.ParseIP(ipAddress)

	hostName, err := os.Hostname()
	if err != nil {
		fmt.Printf("unable to retrieve the host name due to %+v", err)
		os.Exit(1)
	}

	// if its host ip then just use eth0 and skip rest of the operation
	if hostName == ipToMonitor.String() {
		hostENI := eniConfig{name: device}
		enis = append(enis, hostENI)
	} else {
		nl := netlinkwrapper.NewNetLink()

		// read route tables to find the hostveth
		routeFilter := &netlink.Route{
			Table: vlanIDToMonitor + 100,
		}
		routes, err := netlink.RouteListFiltered(netlink.FAMILY_V4, routeFilter, netlink.RT_FILTER_TABLE)
		if err != nil {
			fmt.Printf("unable to get routes for table using vlanID %d. Error: %+v", vlanIDToMonitor, err)
			os.Exit(1)
		}
		for _, route := range routes {
			if route.Dst != nil && ipToMonitor.Equal(route.Dst.IP) {
				linkIndex := route.LinkIndex
				link, err := nl.LinkByIndex(linkIndex)
				if err != nil {
					fmt.Printf("unable to find index %d error %+v", linkIndex, err)
					os.Exit(1)
				}
				hostVethToMonitor := eniConfig{name: link.Attrs().Name, shouldCheckSrc: true}
				enis = append(enis, hostVethToMonitor)
				break
			}
		}

		// get vlan devices
		if vlanIDToMonitor != 0 {
			link, err := nl.LinkByName(fmt.Sprintf("vlan.eth.%d", vlanIDToMonitor))
			if err != nil {
				fmt.Printf("unable to get vlan device, error: %+v", err)
				os.Exit(1)
			}
			vlanDevToMonitor := eniConfig{name: link.Attrs().Name, shouldCheckSrc: true}
			enis = append(enis, vlanDevToMonitor)

			// find the trunk dev
			parentLink, err := nl.LinkByIndex(link.Attrs().ParentIndex)
			if err != nil {
				fmt.Printf("unable to get parent link %d. Error %+v", link.Attrs().ParentIndex, err)
				os.Exit(1)
			}
			trunkDevToMonitor := eniConfig{name: parentLink.Attrs().Name, shouldVerifyVlanTag: true}
			enis = append(enis, trunkDevToMonitor)
		} else {
			// find the eni to monitor associated with the pod
			rules, err := nl.RuleList(netlink.FAMILY_V4)
			if err != nil {
				fmt.Printf("unable to get ip rules due to %+v", err)
				os.Exit(1)
			}
			for _, rule := range rules {
				// Find the ENI in the route table associated with the pod
				if rule.Src != nil && ipToMonitor.Equal(rule.Src.IP) {
					routeFilter := &netlink.Route{
						Table: rule.Table,
					}
					routes, err := netlink.RouteListFiltered(netlink.FAMILY_V4, routeFilter, netlink.RT_FILTER_TABLE)
					if err != nil {
						fmt.Printf("unable to get routes based on iprules %+v", err)
						os.Exit(1)
					}

					if len(routes) > 0 {
						parentLink, err := nl.LinkByIndex(routes[0].LinkIndex)
						if err != nil {
							fmt.Printf("unable to get parent ENI for the link %+v", err)
							os.Exit(1)
						}
						eniToMonitor := eniConfig{name: parentLink.Attrs().Name}
						enis = append(enis, eniToMonitor)
					}
				}
			}
			// if explicit route is not round then pod has to be associated with $device
			if len(enis) == 0 {
				eniToMonitor := eniConfig{name: device}
				enis = append(enis, eniToMonitor)
			}
		}
	}

	fmt.Printf("ENIs to monitor: %+v\n", enis)

	err = monitorPacketOnInterfaces(ipToMonitor, vlanIDToMonitor, enis)
	if err != nil {
		fmt.Printf("unable to verify packets on the interface %+v ", err)
		os.Exit(1)
	}

	fmt.Println("Successfully verified all the interfaces.")
}

// monitorPacketOnInterfaces invokes monitorPackets for each interface
func monitorPacketOnInterfaces(ipToMonitor net.IP, vlanIDToMonitor int, enis []eniConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// monitorPackets monitors the packets on the interfaces
func monitorPackets(ipToMonitor net.IP, vlanIDToMonitor int, iface eniConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Use the handle as a packet source to process all packets

// Verify vlan tag (on ENIs we could see other IP pkts as well)

/*icmpPkt := packet.Layer(layers.LayerTypeICMPv4)
if icmpPkt != nil {
	icmpData, _ := icmpPkt.(*layers.ICMPv4)
	log.Infof("Icmp packet sequence: %d", icmpData.Seq)
}*/

// printVersion prints the binary version to stderr.
func printVersion() { _ = "STUB: not implemented"; return }

// printUsage prints usage information to stderr.
func printUsage() { _ = "STUB: not implemented"; return }
