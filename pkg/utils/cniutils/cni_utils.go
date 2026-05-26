package cniutils

import (
	"time"

	current "github.com/containernetworking/cni/pkg/types/100"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/netlinkwrapper"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/procsyswrapper"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const (
	ipv4ForwardKey = "net/ipv4/ip_forward"
	ipv6ForwardKey = "net/ipv6/conf/all/forwarding"
)

func FindInterfaceByName(ifaceList []*current.Interface, ifaceName string) (ifaceIndex int, iface *current.Interface, found bool) {
	_ = "STUB: not implemented"
	return 0, nil, false
}

func FindIPConfigsByIfaceIndex(ipConfigs []*current.IPConfig, ifaceIndex int) []*current.IPConfig {
	_ = "STUB: not implemented"
	return nil
}

// WaitForAddressesToBeStable Implements `SettleAddresses` functionality of the `ip` package.
// waitForAddressesToBeStable waits for all addresses on a link to leave tentative state.
// Will be particularly useful for ipv6, where all addresses need to do DAD.
// If any addresses are still tentative after timeout seconds, then error.
func WaitForAddressesToBeStable(netLink netlinkwrapper.NetLink, ifName string, timeout, waitInterval time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// GetNodeMetadata calling node local imds metadata service using provided key
// return either a non-empty value or an error
func GetNodeMetadata(key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// EnableIpForwarding sets forwarding to 1 for both IPv4 and IPv6 if applicable.
// This func is to have a unit testable version of ip.EnableForward in ipforward_linux.go file
// link: https://github.com/containernetworking/plugins/blob/main/pkg/ip/ipforward_linux.go#L34
func EnableIpForwarding(procSys procsyswrapper.ProcSys, ips []*current.IPConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// IsLinkNotFoundError return true if err contains "Link not found"
func IsLinkNotFoundError(err error) bool { _ = "STUB: not implemented"; return false }

// IsIptableTargetNotExist returns true if the error is from iptables indicating
// that the target does not exist.
func IsIptableTargetNotExist(err error) bool { _ = "STUB: not implemented"; return false }

// PrefixSimilar checks if prefix pool and eni prefix are equivalent.
func PrefixSimilar(prefixPool []string, eniPrefixes []ec2types.Ipv4PrefixSpecification) bool {
	_ = "STUB: not implemented"
	return false
}

// IPsSimilar checks if ipPool and eniIPs are equivalent.
func IPsSimilar(ipPool []string, eniIPs []ec2types.NetworkInterfacePrivateIpAddress) bool {
	_ = "STUB: not implemented"
	// Here we do +1 in ipPool because eniIPs will also have primary IP which is not used by pods.
	return false
}
