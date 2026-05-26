package networkutils

import (
	"net"
)

// BaseNumber is the base offset for multi-NIC route table IDs.
// This value is chosen to match the logic in Amazon EC2 net utils:
// https://github.com/amazonlinux/amazon-ec2-net-utils/blob/v2.7.1/lib/lib.sh#L301
const BaseNumber = 10000

func CalculateOldRouteTableId(deviceNumber int, networkCardIndex int, maxENIsPerNetworkCard int) int {
	_ = "STUB: not implemented"
	return 0
}

func CalculateRouteTableId(deviceNumber int, networkCardIndex int) int {
	_ = "STUB: not implemented"
	return 0
}

func CalculatePodIPv4GatewayIP(index int) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func CalculatePodIPv6GatewayIP(index int) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func IsIPv4(ip net.IP) bool { _ = "STUB: not implemented"; return false }
