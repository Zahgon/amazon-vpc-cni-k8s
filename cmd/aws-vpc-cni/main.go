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

// NOTE(jaypipes): Normally, we would prefer *not* to have an entrypoint script
// and instead just start the agent daemon as the container's CMD. However, the
// design of CNI is such that Kubelet looks for the presence of binaries and CNI
// configuration files in specific directories, and the presence of those files
// is the trigger to Kubelet that that particular CNI plugin is "ready".
//
// In the case of the AWS VPC CNI plugin, we have two components to the plugin.
// The first component is the actual CNI binary that is execve'd from Kubelet
// when a container is started or destroyed. The second component is the
// aws-k8s-agent daemon which houses the IPAM controller.
//
// As mentioned above, Kubelet considers a CNI plugin "ready" when it sees the
// binary and configuration file for the plugin in a well-known directory. For
// the AWS VPC CNI plugin binary, we only want to copy the CNI plugin binary
// into that well-known directory AFTER we have successfully started the IPAM
// daemon and know that it can connect to Kubernetes and the local EC2 metadata
// service. This is why this entrypoint script exists; we start the IPAM daemon
// and wait until we know it is up and running successfully before copying the
// CNI plugin binary and its configuration file to the well-known directory that
// Kubelet looks in.

// AWS VPC CNI entrypoint binary
package main

import (
	"net"
	"os"

	"github.com/containernetworking/cni/pkg/types"
)

const (
	egressPluginIpamSubnetV4     = "169.254.172.0/22"
	egressPluginIpamSubnetV6     = "fd00::ac:00/118"
	egressPluginIpamDstV4        = "0.0.0.0/0"
	egressPluginIpamDstV6        = "::/0"
	egressPluginIpamDataDirV4    = "/run/cni/v6pd/egress-v4-ipam"
	egressPluginIpamDataDirV6    = "/run/cni/v4pd/egress-v6-ipam"
	defaultHostCniBinPath        = "/host/opt/cni/bin"
	defaultHostCniConfDirPath    = "/host/etc/cni/net.d"
	defaultAWSconflistFile       = "/app/10-aws.conflist"
	tmpAWSconflistFile           = "/tmp/10-aws.conflist"
	defaultVethPrefix            = "eni"
	defaultMTU                   = 9001
	minMTUv4                     = 576
	minMTUv6                     = 1280
	defaultEnablePodEni          = false
	defaultPodSGEnforcingMode    = "strict"
	defaultPluginLogFile         = "/var/log/aws-routed-eni/plugin.log"
	defaultEgressV4PluginLogFile = "/var/log/aws-routed-eni/egress-v4-plugin.log"
	defaultEgressV6PluginLogFile = "/var/log/aws-routed-eni/egress-v6-plugin.log"
	defaultPluginLogLevel        = "Debug"
	defaultEnableIPv6            = false
	defaultEnableIPv6Egress      = false
	defaultEnableIPv4Egress      = true
	defaultRandomizeSNAT         = "prng"
	awsConflistFile              = "/10-aws.conflist"
	vpcCniInitDonePath           = "/vpc-cni-init/done"
	defaultEnBandwidthPlugin     = false
	defaultEnPrefixDelegation    = false
	defaultIPCooldownPeriod      = 30
	defaultDisablePodV6          = false
	defaultEnableMultiNICSupport = false

	envHostCniBinPath        = "HOST_CNI_BIN_PATH"
	envHostCniConfDirPath    = "HOST_CNI_CONFDIR_PATH"
	envVethPrefix            = "AWS_VPC_K8S_CNI_VETHPREFIX"
	envEniMTU                = "AWS_VPC_ENI_MTU"
	envPodMTU                = "POD_MTU"
	envEnablePodEni          = "ENABLE_POD_ENI"
	envPodSGEnforcingMode    = "POD_SECURITY_GROUP_ENFORCING_MODE"
	envPluginLogFile         = "AWS_VPC_K8S_PLUGIN_LOG_FILE"
	envPluginLogLevel        = "AWS_VPC_K8S_PLUGIN_LOG_LEVEL"
	envEgressV4PluginLogFile = "AWS_VPC_K8S_EGRESS_V4_PLUGIN_LOG_FILE"
	envEgressV6PluginLogFile = "AWS_VPC_K8S_EGRESS_V6_PLUGIN_LOG_FILE"
	envEnPrefixDelegation    = "ENABLE_PREFIX_DELEGATION"
	envWarmIPTarget          = "WARM_IP_TARGET"
	envMinIPTarget           = "MINIMUM_IP_TARGET"
	envWarmPrefixTarget      = "WARM_PREFIX_TARGET"
	envEnBandwidthPlugin     = "ENABLE_BANDWIDTH_PLUGIN"
	envEnIPv6                = "ENABLE_IPv6"
	envEnIPv6Egress          = "ENABLE_V6_EGRESS"
	envEnIPv4Egress          = "ENABLE_V4_EGRESS"
	envRandomizeSNAT         = "AWS_VPC_K8S_CNI_RANDOMIZESNAT"
	envIPCooldownPeriod      = "IP_COOLDOWN_PERIOD"
	envDisablePodV6          = "DISABLE_POD_V6"
	envEnableMultiNICSupport = "ENABLE_MULTI_NIC"
)

// NetConfList describes an ordered list of networks.
type NetConfList struct {
	CNIVersion string `json:"cniVersion,omitempty"`

	Name         string     `json:"name,omitempty"`
	DisableCheck bool       `json:"disableCheck,omitempty"`
	Plugins      []*NetConf `json:"plugins,omitempty"`
}

// NetConf stores the common network config for the CNI plugin
type NetConf struct {
	CNIVersion string `json:"cniVersion,omitempty"`

	Name         string            `json:"name,omitempty"`
	Type         string            `json:"type,omitempty"`
	Capabilities map[string]bool   `json:"capabilities,omitempty"`
	IPAM         *IPAMConfig       `json:"ipam,omitempty"`
	DNS          *types.DNS        `json:"dns,omitempty"`
	Sysctl       map[string]string `json:"sysctl,omitempty"`

	RawPrevResult map[string]interface{} `json:"prevResult,omitempty"`
	PrevResult    types.Result           `json:"-"`

	// Interface inside container to create
	IfName string `json:"ifName,omitempty"`

	Enabled string `json:"enabled,,omitempty"`

	// IP to use as SNAT target
	NodeIP net.IP `json:"nodeIP,omitempty"`

	VethPrefix string `json:"vethPrefix,omitempty"`

	PodSGEnforcingMode string `json:"podSGEnforcingMode,omitempty"`

	RandomizeSNAT string `json:"randomizeSNAT,omitempty"`

	// MTU for eth0
	MTU string `json:"mtu,omitempty"`

	PluginLogFile string `json:"pluginLogFile,omitempty"`

	PluginLogLevel string `json:"pluginLogLevel,omitempty"`
}

// IPAMConfig references containernetworking structure defined at https://github.com/containernetworking/plugins/blob/main/plugins/ipam/host-local/backend/allocator/config.go
type IPAMConfig struct {
	*Range
	Name       string         `json:"name,omitempty"`
	Type       string         `json:"type,omitempty"`
	Routes     []*types.Route `json:"routes,omitempty"`
	DataDir    string         `json:"dataDir,omitempty"`
	ResolvConf string         `json:"resolvConf,omitempty"`
	Ranges     []RangeSet     `json:"ranges"`
	IPArgs     []net.IP       `json:"-"` // Requested IPs from CNI_ARGS and args
}

// RangeSet references containernetworking structure
type RangeSet []Range

// Range references containernetworking structure
type Range struct {
	RangeStart net.IP      `json:"rangeStart,omitempty"` // The first ip, inclusive
	RangeEnd   net.IP      `json:"rangeEnd,omitempty"`   // The last ip, inclusive
	Subnet     types.IPNet `json:"subnet"`
	Gateway    net.IP      `json:"gateway,omitempty"`
}

// Wait for IPAMD health check to pass. Note that if IPAMD fails to start, wait happens indefinitely until liveness probe kills pod
func waitForIPAM() bool { _ = "STUB: not implemented"; return false }

func getPrimaryIP(ipv4 bool) (string, error) { _ = "STUB: not implemented"; return "", nil }

func isValidJSON(inFile string) error { _ = "STUB: not implemented"; return nil }

func generateJSON(jsonFile string, outFile string, getPrimaryIP func(ipv4 bool) (string, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// enabledIPv6 is to determine if EKS cluster is IPv4 or IPv6 cluster
// if this EKS cluster is IPv6 cluster, egress-cni-plugin will enable IPv4 egress by default
// if this EKS cluster is IPv4 cluster, egress-cni-plugin will only enable IPv6 egress if env var "ENABLE_V6_EGRESS" is "true"

// EKS IPv6 cluster

// Enable IPv4 egress when "ENABLE_V4_EGRESS" is "true" (default)

// Node should have a IPv4 address even in IPv6 cluster

// EKS IPv4 cluster

// When ENABLE_V6_EGRESS is set, but the node is lacking an IPv6 address, log a warning and disable the egress-v6-cni plugin.
// This allows IPv4-only nodes to function while still alerting the customer to the possibility of a misconfiguration.

// Derive pod MTU from ENI MTU by default (note that values have already been validated)

// If pod MTU environment variable is set, overwrite ENI MTU.

// Chain any requested CNI plugins

// Unmarshall current conflist into data

// Chain the bandwidth plugin when enabled

// Chain the tuning plugin (configured to disable IPv6 in pod network namespace) when requested

// Marshall data back into byteValue

func validateEnvVars() bool { _ = "STUB: not implemented"; return false }

// Validate that veth prefix is less than or equal to four characters and not in reserved set: (eth, lo, vlan)

// When ENABLE_POD_ENI is set, validate security group enforcing mode

// Validate that IP_COOLDOWN_PERIOD is a valid integer

// Validate MTU value for ENIs and pods

// Note that these string values should probably be cast to integers, but the comparison for values greater than 0 works either way

func validateMTU(envVar string) bool {
	_ = "STUB: not implemented"
	// Validate MTU range based on IP address family
	return false
}

func main() {
	os.Exit(_main())
}

func _main() int { _ = "STUB: not implemented"; return 0 }

// Exec redirects stdout and stderr to /dev/null, redirecting to os.Stdout and os.Stderr is done explicitly.
// This enables the output of the aws-k8s-agent to be displayed in the kubectl logs for the aws-node container via stdout and stderr.
