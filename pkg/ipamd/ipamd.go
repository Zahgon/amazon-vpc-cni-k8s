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

package ipamd

import (
	"context"
	"net"
	"sync"
	"time"

	"sigs.k8s.io/controller-runtime/pkg/client"

	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	corev1 "k8s.io/api/core/v1"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/awsutils"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/ipamd/datastore"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/networkutils"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/utils/logger"
	rcv1alpha1 "github.com/aws/amazon-vpc-resource-controller-k8s/apis/vpcresources/v1alpha1"
)

// The package ipamd is a long running daemon which manages a warm pool of available IP addresses.
// It also monitors the size of the pool, dynamically allocates more ENIs when the pool size goes below
// the minimum threshold and frees them back when the pool size goes above max threshold.

const (
	ipPoolMonitorInterval       = 5 * time.Second
	maxRetryCheckENI            = 5
	eniAttachTime               = 10 * time.Second
	nodeIPPoolReconcileInterval = 60 * time.Second
	decreaseIPPoolInterval      = 30 * time.Second

	// ipReconcileCooldown is the amount of time that an IP address must wait until it can be added to the data store
	// during reconciliation after being discovered on the EC2 instance metadata.
	ipReconcileCooldown = 60 * time.Second

	// This environment variable is used to specify the desired number of free IPs always available in the "warm pool".
	// When it is not set, ipamd defaults to use all available IPs per ENI for that instance type.
	// For example, for a m4.4xlarge node,
	//     If WARM_IP_TARGET is set to 1, and there are 9 pods running on the node, ipamd will try
	//     to make the "warm pool" have 10 IP addresses with 9 being assigned to pods and 1 free IP.
	//
	//     If "WARM_IP_TARGET is not set, it will default to 30 (which the maximum number of IPs per ENI).
	//     If there are 9 pods running on the node, ipamd will try to make the "warm pool" have 39 IPs with 9 being
	//     assigned to pods and 30 free IPs.
	envWarmIPTarget = "WARM_IP_TARGET"
	noWarmIPTarget  = 0

	// This environment variable is used to specify the desired minimum number of total IPs.
	// When it is not set, ipamd defaults to 0.
	// For example, for a m4.4xlarge node,
	//     If WARM_IP_TARGET is set to 1 and MINIMUM_IP_TARGET is set to 12, and there are 9 pods running on the node,
	//     ipamd will make the "warm pool" have 12 IP addresses with 9 being assigned to pods and 3 free IPs.
	//
	//     If "MINIMUM_IP_TARGET is not set, it will default to 0, which causes WARM_IP_TARGET settings to be the
	//	   only settings considered.
	envMinimumIPTarget = "MINIMUM_IP_TARGET"
	noMinimumIPTarget  = 0

	// This environment is used to specify the desired number of free ENIs along with all of its IP addresses
	// always available in "warm pool".
	// When it is not set, it is default to 1.
	//
	// when "WARM_IP_TARGET" is defined, ipamd will use behavior defined for "WARM_IP_TARGET".
	//
	// For example, for a m4.4xlarge node
	//     If WARM_ENI_TARGET is set to 2, and there are 9 pods running on the node, ipamd will try to
	//     make the "warm pool" to have 2 extra ENIs and its IP addresses, in other words, 90 IP addresses
	//     with 9 IPs assigned to pods and 81 free IPs.
	//
	//     If "WARM_ENI_TARGET" is not set, it defaults to 1, so if there are 9 pods running on the node,
	//     ipamd will try to make the "warm pool" have 1 extra ENI, in other words, 60 IPs with 9 already
	//     being assigned to pods and 51 free IPs.
	envWarmENITarget     = "WARM_ENI_TARGET"
	defaultWarmENITarget = 1

	// This environment variable is used to specify the maximum number of ENIs that will be allocated.
	// When it is not set or less than 1, the default is to use the maximum available for the instance type.
	//
	// The maximum number of ENIs is in any case limited to the amount allowed for the instance type.
	envMaxENI     = "MAX_ENI"
	defaultMaxENI = -1

	// This environment is used to specify whether Pods need to use a security group and subnet defined in an ENIConfig CRD.
	// When it is NOT set or set to false, ipamd will use primary interface security group and subnet for Pod network.
	envCustomNetworkCfg = "AWS_VPC_K8S_CNI_CUSTOM_NETWORK_CFG"

	// This environment variable specifies whether IPAMD should allocate or deallocate ENIs on a non-schedulable node (default false).
	envManageENIsNonSchedulable = "AWS_MANAGE_ENIS_NON_SCHEDULABLE"

	// This environment is used to specify whether we should use enhanced subnet selection or not when creating ENIs (default true).
	envSubnetDiscovery = "ENABLE_SUBNET_DISCOVERY"

	// eniNoManageTagKey is the tag that may be set on an ENI to indicate ipamd
	// should not manage it in any form.
	eniNoManageTagKey = "node.k8s.amazonaws.com/no_manage"

	// disableENIProvisioning is used to specify that ENIs do not need to be synced during initializing a pod.
	envDisableENIProvisioning = "DISABLE_NETWORK_RESOURCE_PROVISIONING"

	// disableLeakedENICleanup is used to specify that the task checking and cleaning up leaked ENIs should not be run.
	envDisableLeakedENICleanup = "DISABLE_LEAKED_ENI_CLEANUP"

	// Specify where ipam should persist its current IP<->container allocations.
	envBackingStorePath     = "AWS_VPC_K8S_CNI_BACKING_STORE"
	defaultBackingStorePath = "/var/run/aws-node/ipam.json"

	// envEnablePodENI is used to attach a Trunk ENI to every node. Required in order to give Branch ENIs to pods.
	envEnablePodENI = "ENABLE_POD_ENI"

	// envNodeName will be used to store Node name
	envNodeName = "MY_NODE_NAME"

	//envEnableIpv4PrefixDelegation is used to allocate /28 prefix instead of secondary IP for an ENI.
	envEnableIpv4PrefixDelegation = "ENABLE_PREFIX_DELEGATION"

	// envWarmPrefixTarget is used to keep a /28 prefix in warm pool.
	envWarmPrefixTarget     = "WARM_PREFIX_TARGET"
	defaultWarmPrefixTarget = 0

	// envEnableIPv4 - Env variable to enable/disable IPv4 mode
	envEnableIPv4 = "ENABLE_IPv4"

	// envEnableIPv6 - Env variable to enable/disable IPv6 mode
	envEnableIPv6 = "ENABLE_IPv6"

	ipV4AddrFamily = "4"
	ipV6AddrFamily = "6"

	// insufficientCidrErrorCooldown is the amount of time reconciler will wait before trying to fetch
	// more IPs/prefixes for an ENI. With InsufficientCidr we know the subnet doesn't have enough IPs so
	// instead of retrying every 5s which would lead to increase in EC2 AllocIPAddress calls, we wait for
	// 120 seconds for a retry.
	insufficientCidrErrorCooldown = 120 * time.Second

	// envManageUntaggedENI is used to determine if untagged ENIs should be managed or unmanaged
	envManageUntaggedENI = "MANAGE_UNTAGGED_ENI"

	eniNodeTagKey = "node.k8s.amazonaws.com/instance_id"

	// envAnnotatePodIP is used to annotate[vpc.amazonaws.com/pod-ips] pod's with IPs
	// Ref : https://github.com/projectcalico/calico/issues/3530
	// not present; in which case we fall back to the k8s podIP
	// Present and set to an IP; in which case we use it
	// Present and set to the empty string, which we use to mean "CNI DEL had occurred; networking has been removed from this pod"
	// The empty string one helps close a trace at pod shutdown where it looks like the pod still has its IP when the IP has been released
	envAnnotatePodIP = "ANNOTATE_POD_IP"

	// aws error codes for insufficient IP address scenario
	INSUFFICIENT_CIDR_BLOCKS    = "InsufficientCidrBlocks"
	INSUFFICIENT_FREE_IP_SUBNET = "InsufficientFreeAddressesInSubnet"

	// envEnableNetworkPolicy is used to enable IPAMD/CNI to send pod create events to network policy agent.
	envNetworkPolicyMode = "NETWORK_POLICY_ENFORCING_MODE"

	defaultMaxPodsFromKubelet = 110
	kubeletConfigPath         = "/host/etc/kubernetes/kubelet/kubelet-config.json"
	eniMaxPodsFilePath        = "/app/eni-max-pods.txt"

	// Application name for k8s client
	appName = "aws-node"

	defaultNetworkPolicyMode = "standard"

	// Network Card index of primary ENI
	DefaultNetworkCardIndex = 0

	// Enable Multi NIC support in CNI
	// This configures the ENIs on Network Card > 0 which is be used by pods that require multi-nic attachments
	envEnableMultiNICSupport = "ENABLE_MULTI_NIC"

	// Scale config for network cards > 0
	DefaultWarmIPTarget    = 1
	DefaultMinimumIPTarget = 1
)

var log = logger.Get()

var prometheusRegistered = false

// IPAMContext contains node level control information
type IPAMContext struct {
	awsClient                 awsutils.APIs
	dataStoreAccess           *datastore.DataStoreAccess
	k8sClient                 client.Client
	enableIPv4                bool
	enableIPv6                bool
	useCustomNetworking       bool
	manageENIsNonScheduleable bool
	useSubnetDiscovery        bool
	networkClient             networkutils.NetworkAPIs
	maxIPsPerENI              int
	maxENI                    int
	maxPrefixesPerENI         int
	unmanagedENI              []int
	numNetworkCards           int

	warmENITarget        int
	warmIPTarget         int
	minimumIPTarget      int
	warmPrefixTarget     int
	primaryIP            map[string]string // primaryIP is a map from ENI ID to primary IP of that ENI
	lastNodeIPPoolAction time.Time
	lastDecreaseIPPool   time.Time
	// reconcileCooldownCache keeps timestamps of the last time an IP address was unassigned from an ENI,
	// so that we don't reconcile and add it back too quickly if IMDS lags behind reality.
	reconcileCooldownCache    ReconcileCooldownCache
	terminating               int32 // Flag to warn that the pod is about to shut down.
	disableENIProvisioning    bool
	enablePodENI              bool
	myNodeName                string
	enablePrefixDelegation    bool
	lastInsufficientCidrError time.Time
	enableManageUntaggedMode  bool
	enablePodIPAnnotation     bool
	maxPods                   int // maximum number of pods that can be scheduled on the node
	networkPolicyMode         string
	enableMultiNICSupport     bool
	withApiServer             bool
}

type kubeletConfig struct {
	MaxPods *int64 `json:"maxPods"`
}

// setUnmanagedENIs will rebuild the set of ENI IDs for ENIs tagged as "no_manage"
func (c *IPAMContext) setUnmanagedENIs(tagMap map[string]awsutils.TagMap) {
	_ = "STUB: not implemented"
	return
}

// if "no_manage" tag is present and is true - ENI is unmanaged
// if "no_manage" tag is present and is "not true" - ENI is managed
// if "instance_id" tag is present and is set to instanceID - ENI is managed since this was created by IPAMD
// if "no_manage" tag is not present or not IPAMD created ENI, check if we are in Manage Untagged Mode, default is true.
// if enableManageUntaggedMode is false (defaults to true), then consider all untagged ENIs as unmanaged.

// setUnmanagedNetworkCards will mark the Network Cards which are not managed by CNI
// When ENABLE_MULTI_NIC is false, all Network Cards > 0 are marked as unmanaged
// When ENABLE_MULTI_NIC is true, Network Cards which ONLY includes an EFA-only device are marked as unmanaged.
// If there is a ENA device on a Network Card along with EFA-only device, CNI manages the Network Card but excludes the EFA-only device
func (c *IPAMContext) markUnmanagedNetworkCards(efaOnlyENINetworkCards []string, enisByNetworkCard [][]string) []bool {
	_ = "STUB: not implemented"
	return nil
}

// Network card does not have an EFA-only ENI, so can be managed

// Skip the network card by default, unless we find a ENA device on this network card which is managed

// ReconcileCooldownCache keep track of recently freed CIDRs to avoid reading stale EC2 metadata
type ReconcileCooldownCache struct {
	sync.RWMutex
	cache map[string]time.Time
}

// Add sets a timestamp for the CIDR added that says how long they are not to be put back in the data store.
func (r *ReconcileCooldownCache) Add(cidr string) { _ = "STUB: not implemented"; return }

// Remove removes a CIDR from the cooldown cache.
func (r *ReconcileCooldownCache) Remove(cidr string) { _ = "STUB: not implemented"; return }

// RecentlyFreed checks if this CIDR was recently freed.
func (r *ReconcileCooldownCache) RecentlyFreed(cidr string) (found, recentlyFreed bool) {
	_ = "STUB: not implemented"
	return false, false
}

func prometheusRegister() { _ = "STUB: not implemented"; return }

// containsInsufficientCIDRsOrSubnetIPs returns whether a CIDR cannot be carved in the subnet or subnet is running out of IP addresses
func containsInsufficientCIDRsOrSubnetIPs(err error) bool { _ = "STUB: not implemented"; return false }

// IP exhaustion can be due to Insufficient Cidr blocks or Insufficient Free Address in a Subnet
// In these 2 cases we will back off for 2 minutes before retrying

// containsPrivateIPAddressLimitExceededError returns whether exceeds ENI's IP address limit
func containsPrivateIPAddressLimitExceededError(err error) bool {
	_ = "STUB: not implemented"
	return false
}

// inInsufficientCidrCoolingPeriod checks whether IPAMD is in insufficientCidrErrorCooldown
func (c *IPAMContext) inInsufficientCidrCoolingPeriod() bool {
	_ = "STUB: not implemented"
	return false
}

// New retrieves IP address usage information from Instance MetaData service and Kubelet
// then initializes IP address pool data store
func New(ctx context.Context, k8sClient client.Client, withApiServer bool) (*IPAMContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WARM and Min IP/Prefix targets are ignored in IPv6 mode

// Validate if the configured combination of env variables is supported before proceeding further

func (c *IPAMContext) nodeInit(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// For multi-card IPv4, non-VPC traffic is still routed out of Primary ENI

// We should not error if clean up fails since these chains don't affect the rules

// Queries IMDS for all attached ENIs and then compares it against EC2.
// Groups the ENIs into different types

// check if primary ENI is excluded for ipv6 cluster

// Retry ENI sync

// If we can't find the matching link for this MAC address, there is no point in retrying for this ENI.

// Read from all backing datastores corresponding to the NICs

// Check all ENIs (primary and secondary) for subnet exclusion

// During upgrade or if prefix delgation knob is disabled to enabled then we
// might have secondary IPs attached to ENIs so doing a cleanup if not used before moving on

// When prefix delegation knob is enabled to disabled then we might
// have unused prefixes attached to the ENIs so need to cleanup

// Spawning updateCIDRsRulesOnChange go-routine

// RefreshSGIDs populates the ENI cache with ENI -> security group ID mappings, and so it must be called:
// 1. after managed/unmanaged ENIs have been determined
// 2. before any new ENIs are attached

// Also refresh custom security groups for secondary subnets
// Custom security groups are only relevant when subnet discovery is enabled and custom networking is disabled.
// When custom networking is enabled, ENIConfig defines the security groups for secondary ENIs,
// and auto-discovered SGs should not overwrite them.

// Refresh security groups and VPC CIDR blocks in the background
// Ignoring errors since we will retry in 30s

// Also refresh custom security groups for secondary subnets

// if apiserver is connected, get the maxPods from node

// When custom networking is enabled and a valid ENIConfig is found, IPAMD patches the CNINode
// resource for this instance. The operation is safe as enabling/disabling custom networking
// requires terminating the previous instance.

// If Security Groups for Pods is enabled, the VPC Resource Controller must also know that Custom Networking is enabled

// Now that Custom Networking is (potentially) enabled, Security Groups for Pods can be enabled for IPv4 nodes.

// We do not return an error here as create ENI failure shouldn not cause the node to go in not ready state.
// However that does restrict multi-nic pods to get failed

func (c *IPAMContext) handlePreScaling(ctx context.Context) error {
	_ = "STUB: not implemented"
	// On node init, check if datastore pool needs to be increased. If so, attach CIDRs from existing ENIs and attach new ENIs.
	return nil
}

// Note that the only error currently returned by increaseDatastorePool is an error attaching CIDRs (other than insufficient IPs)

// If custom networking is enabled and the pool is empty, return an error, as there is a misconfiguration and
// the node should not become ready.

func (c *IPAMContext) configureIPRulesForPods() error { _ = "STUB: not implemented"; return nil }

// TODO(gus): This should really be done via CNI CHECK calls, rather than in ipam (requires upstream k8s changes).

// Update ip rules in case there is a change in VPC CIDRs, AWS_VPC_K8S_CNI_EXTERNALSNAT setting

// Program IP rules for external service CIDRs and cleanup stale rules.
// Note that we can reuse rule list despite it being modified by UpdateRuleListBySrc, as the
// modifications touched rules that this function ignores.

func (c *IPAMContext) updateCIDRsRulesOnChange(oldVPCCIDRs []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *IPAMContext) updateIPStats(unmanaged int) { _ = "STUB: not implemented"; return }

// StartNodeIPPoolManager monitors the IP pool, add or del them when it is required.
func (c *IPAMContext) StartNodeIPPoolManager(ctx context.Context) {
	_ = "STUB: not implemented"
	// For IPv6, if Security Groups for Pods is enabled, wait until trunk ENI is attached and add it to the datastore.
	return
}

// Outside of Security Groups for Pods, no additional ENIs are attached in IPv6 mode.
// The prefix used for the primary ENI is more than enough for all pods.

func (c *IPAMContext) updateIPPoolIfRequired(ctx context.Context) {
	_ = "STUB: not implemented"
	// When IPv4 Security Groups for Pods is configured, do not write to CNINode until there is room for a trunk ENI
	return
}

// Each iteration, log the current datastore IP stats

// Store the last update time so that we only scale down every

func (c *IPAMContext) shouldSkipDataStorePoolDecrease() bool {
	_ = "STUB: not implemented"
	return false
}

// decreaseDatastorePool runs every `interval` and attempts to return unused ENIs and IPs
func (c *IPAMContext) decreaseDatastorePool(ctx context.Context, networkCard int) {
	_ = "STUB: not implemented"
	return
}

// tryFreeENI always tries to free one ENI
func (c *IPAMContext) tryFreeENI(ctx context.Context, networkCard int) {
	_ = "STUB: not implemented"
	return
}

// If the ENI attachment not found or cannot be detached we return without deleting the primary IP rules

// When warm IP/prefix targets are defined, free extra IPs
func (c *IPAMContext) tryUnassignCidrsFromAll(ctx context.Context, networkCard int) {
	_ = "STUB: not implemented"
	return
}

// If WARM IP targets are not defined, check if WARM_PREFIX_TARGET is defined.

// Either returns prefixes or IPs [Cidrs]

// Free the number of Cidrs `over` the warm IP target, unless `over` is greater than the number of available Cidrs on
// this ENI. In that case we should only free the number of available Cidrs.

// Delete IPs from datastore

// Do not force the delete, since a freeable Cidr might have been assigned to a pod
// before we get around to deleting it.
/* force */

// Deallocate Cidrs from the instance if they are not used by pods.

// reduce the deallocation target, if the deallocation target is achieved, we can exit

func (c *IPAMContext) initNetworkConfig() ([]string, net.IP, error) {
	_ = "STUB: not implemented"
	return nil, *new(net.IP), nil
}

func (c *IPAMContext) isENIAttachmentAllowed() bool { _ = "STUB: not implemented"; return false }

// PRECONDITION: isDatastorePoolTooLow returned true
func (c *IPAMContext) increaseDatastorePool(ctx context.Context, networkCard int) error {
	_ = "STUB: not implemented"
	return nil
}

// Try to add more Cidrs to existing ENIs first.

// If we did not add any IPs, try to allocate an ENI.

// Note that no error is returned if ENI allocation fails. This is because ENI allocation failure should not cause node to be "NotReady".

func (c *IPAMContext) createSecondaryIPv6ENIs(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// This means we have not added an ENI before as it would have been detected ENI detection

// Note that no error is returned if ENI allocation fails. This is because ENI allocation failure should not cause node to be "NotReady".

func (c *IPAMContext) updateLastNodeIPPoolAction(networkCard int) {
	_ = "STUB: not implemented"
	return
}

func (c *IPAMContext) tryAllocateENI(ctx context.Context, networkCard int) error {
	_ = "STUB: not implemented"
	return nil
}

// The CNI does not create trunk or EFA ENIs, so they will always be false here

// For an ENI, fill in missing IPs or prefixes.
// PRECONDITION: isDatastorePoolTooLow returned true
func (c *IPAMContext) tryAssignCidrs(ctx context.Context, networkCard int) (increasedPool bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// For an ENI, try to fill in missing IPs on an existing ENI.
// PRECONDITION: isDatastorePoolTooLow returned true
func (c *IPAMContext) tryAssignIPs(ctx context.Context, networkCard int) (increasedPool bool, err error) {
	_ = "STUB: not implemented"
	// If WARM_IP_TARGET is set, only proceed if we are short of target
	return false, nil
}

// If WARM_IP_TARGET is set we only want to allocate up to that target to avoid overallocating and releasing

// Find an ENI where we can add more IPs

// Try to allocate all available IPs for this ENI

// Try to just get one more IP

// This call to EC2 is needed to verify which IPs got attached to this ENI.

func (c *IPAMContext) assignIPv6Prefix(ctx context.Context, eniID string, networkCard int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Let's make an EC2 API call to get a list of IPv6 prefixes (if any) that are already attached to the
// current ENI. We will make this call only once during boot up/init and doing so will shield us from any
// IMDS out of sync issues. We only need one v6 prefix per ENI/Node.

// Note: If we find more than one v6 prefix attached to the ENI, VPC CNI will not attempt to free it. VPC CNI
// will only attach a single v6 prefix and it will not attempt to free the additional Prefixes.
// We will add all the prefixes to our datastore. TODO - Should we instead pick one of them. If we do, how to track
// that across restarts?

// Check if we already have v6 Prefix(es) attached

// Allocate and attach a v6 Prefix to Primary ENI

// Found more than one v6 prefix attached to the ENI. VPC CNI will only attach a single v6 prefix
// and it will not attempt to free any additional Prefixes that are already attached.
// Will use the first IPv6 Prefix attached for IP address allocation.

// PRECONDITION: isDatastorePoolTooLow returned true
func (c *IPAMContext) tryAssignPrefixes(ctx context.Context, networkCard int) (increasedPool bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Returns an ENI which has space for more prefixes to be attached, but this
// ENI might not suffice the WARM_IP_TARGET/WARM_PREFIX_TARGET

// Try to just get one more prefix

// This call to EC2 is needed to verify which IPs got attached to this ENI.

// setupENI does following:
// 1) get route table ID for the ENI
// 2) add ENI to datastore
// 3) set up linux ENI related networking stack.
// 4) add all ENI's secondary IP addresses to datastore
func (c *IPAMContext) setupENI(ctx context.Context, eni string, eniMetadata awsutils.ENIMetadata, isTrunkENI, isEFAENI bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the route table ID for the ENI

// Add the ENI to the datastore

// Check if this ENI (primary or secondary) is in an excluded subnet and mark it for exclusion

// Store the addressable IP for the ENI

// In v6 PD mode, VPC CNI will manage the primary ENI, ENIs on NC > 0 and trunk ENI. Once we start supporting secondary
// IP and custom networking modes for IPv6, this restriction can be relaxed.

// For other ENIs, set up the network

// Failed to set up the ENI

// Either case add the IPs and prefixes to datastore.

func (c *IPAMContext) addENIsecondaryIPsToDataStore(ec2PrivateIpAddrs []ec2types.NetworkInterfacePrivateIpAddress, eni string, networkCard int) {
	_ = "STUB: not implemented"
	// Add all the secondary IPs
	return
}

// continue to add next address

func (c *IPAMContext) addENIv4prefixesToDataStore(ec2PrefixAddrs []ec2types.Ipv4PrefixSpecification, eni string, networkCard int) {
	_ = "STUB: not implemented"
	// Walk thru all prefixes
	return
}

// Parsing failed, get next prefix

// continue to add next address

func (c *IPAMContext) addENIv6prefixesToDataStore(ec2PrefixAddrs []ec2types.Ipv6PrefixSpecification, eni string, networkCard int) {
	_ = "STUB: not implemented"
	return
}

// Walk through all prefixes

// Parsing failed, get next prefix

// continue to add next address

// getMaxENI returns the maximum number of ENIs to attach to this instance. This is calculated as the lesser of
// the limit for the instance type and the value configured via the MAX_ENI environment variable. If the value of
// the environment variable is 0 or less, it will be ignored and the maximum for the instance is returned.
func (c *IPAMContext) getMaxENI() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func getWarmENITarget() int { _ = "STUB: not implemented"; return 0 }

func getWarmPrefixTarget() int { _ = "STUB: not implemented"; return 0 }

// logPoolStats logs usage information for allocated addresses/prefixes.
func (c *IPAMContext) logPoolStats(dataStoreStats *datastore.DataStoreStats, networkCard int) {
	_ = "STUB: not implemented"
	return
}

func (c *IPAMContext) tryEnableSecurityGroupsForPods(ctx context.Context) {
	_ = "STUB: not implemented"
	// For IPv4, check that there is room for a trunk ENI before patching CNINode CRD. We only check on the Default Network Card
	return
}

// Signal to the VPC Resource Controller that Security Groups for Pods is enabled

// shouldRemoveExtraENIs returns true if we should attempt to find an ENI to free
// PD enabled: If the WARM_PREFIX_TARGET is spread across ENIs and we have more than needed, this function will return true.
// If the number of prefixes are on just one ENI, and there are more than available, it returns true so getDeletableENI will
// recheck if we need the ENI for prefix target.
func (c *IPAMContext) shouldRemoveExtraENIs(stats *datastore.DataStoreStats, networkCard int) bool {
	_ = "STUB: not implemented"
	// When WARM_IP_TARGET is set, return true as verification is always done in getDeletableENI()
	return false
}

// We need the +1 to make sure we are not going below the WARM_ENI_TARGET/WARM_PREFIX_TARGET

// When prefix target count is reduced, datastore would have deleted extra prefixes over the warm prefix target.
// Hence available will be less than (warmTarget)*c.maxIPsPerENI, but there can be some extra ENIs which are not used hence see if we can clean it up.

func (c *IPAMContext) computeExtraPrefixesOverWarmTarget(networkCard int) int {
	_ = "STUB: not implemented"
	return 0
}

func ipamdErrInc(fn string) { _ = "STUB: not implemented"; return }

func podENIErrInc(fn string) { _ = "STUB: not implemented"; return }

// Used in IPv6 mode to check if trunk ENI has been successfully attached
func (c *IPAMContext) checkForTrunkENI(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *IPAMContext) getENIsByNetworkCard(allENIs []awsutils.ENIMetadata) map[int][]awsutils.ENIMetadata {
	_ = "STUB: not implemented"
	return nil
}

// nodeIPPoolReconcile reconcile ENI and IP info from metadata service and IP addresses in datastore
func (c *IPAMContext) nodeIPPoolReconcile(ctx context.Context, interval time.Duration) {
	_ = "STUB: not implemented"
	// To reduce the number of EC2 API calls, skip reconciliation if IPs were recently added to the datastore.
	return
}

// Make an exception if node needs a trunk ENI and one is not currently attached.

// We must always have at least the primary ENI of the instance

// Initialize the set with the known EFA interfaces

// Check if a new ENI was added, if so we need to update the tags.

// Call describeENIs only once per nodeIPPoolReconcile by checking if metadataResult.ENIMetadata is nil

// Update trunk ENI

// Just copy values of the EFA set

// Mark phase

// If the attached ENI is in the data store

// Reconcile IP pool

// If the attached ENI is in the data store

// Reconcile IP pool

// Mark action, remove this ENI from currentENIs map

// Add new ENI

// Continue if having trouble with ONLY 1 ENI, instead of bailout here?

// Sweep phase: since the marked ENI have been removed, the remaining ones needs to be sweeped

// Force the delete, since aws local metadata has told us that this ENI is no longer
// attached, so any IPs assigned from this ENI will no longer work.
/* force */

func (c *IPAMContext) eniIPPoolReconcile(ctx context.Context, ipPool []string, attachedENI awsutils.ENIMetadata, eni string, networkCard int) {
	_ = "STUB: not implemented"
	return
}

// Here we can't trust attachedENI since the IMDS metadata can be stale. We need to check with EC2 API.
// IPsSimilar will exclude primary IP of the ENI that is not added to the ipPool and not available for pods to use.

// Call EC2 to verify IPs on this ENI

// Add all known attached IPs to the datastore

// Sweep phase, delete remaining IPs since they should not remain in the datastore

// Force the delete, since we have verified with EC2 that these secondary IPs are no longer assigned to this ENI

/* force */

// continue instead of bailout due to one ip

func (c *IPAMContext) eniPrefixPoolReconcile(ctx context.Context, prefixPool []string, attachedENI awsutils.ENIMetadata, eni string, networkCard int) {
	_ = "STUB: not implemented"
	return
}

// Here we can't trust attachedENI since the IMDS metadata can be stale. We need to check with EC2 API.

// Call EC2 to verify IPs on this ENI

// Add all known attached IPs to the datastore

// Sweep phase, delete remaining Prefixes since they should not remain in the datastore

// Force the delete, since we have verified with EC2 that these secondary IPs are no longer assigned to this ENI

/* force */

// continue instead of bailout due to one ip

// verifyAndAddIPsToDatastore updates the datastore with the known secondary IPs. IPs who are out of cooldown gets added
// back to the datastore after being verified against EC2.
func (c *IPAMContext) verifyAndAddIPsToDatastore(ctx context.Context, eni string, attachedENIIPs []ec2types.NetworkInterfacePrivateIpAddress, needEC2Reconcile bool, networkCard int) map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

// Check if this IP was recently freed

// IMDS data might be stale

// Only call EC2 once for this ENI

// Call EC2 to verify IPs on this ENI

// Do not delete this IP from the datastore or cooldown until we have confirmed with EC2

// Verify that the IP really belongs to this ENI

// The IP can be removed from the cooldown cache
// TODO: Here we could check if the IP is still used by a pod stuck in Terminating state. (Issue #1091)

// Try to add the IP

// Continue to check the other IPs instead of bailout due to one wrong IP

// Mark action

// verifyAndAddPrefixesToDatastore updates the datastore with the known Prefixes. Prefixes who are out of cooldown gets added
// back to the datastore after being verified against EC2.
func (c *IPAMContext) verifyAndAddPrefixesToDatastore(ctx context.Context, eni string, attachedENIPrefixes []ec2types.Ipv4PrefixSpecification, needEC2Reconcile bool, networkCard int) map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

// Check if this Prefix was recently freed

// IMDS data might be stale

// Only call EC2 once for this ENI and post GA fix this logic for both prefixes
// and secondary IPs as per "split the loop" comment

// Call EC2 to verify Prefixes on this ENI

// Do not delete this Prefix from the datastore or cooldown until we have confirmed with EC2

// Verify that the Prefix really belongs to this ENI

// The IP can be removed from the cooldown cache
// TODO: Here we could check if the Prefix is still used by a pod stuck in Terminating state. (Issue #1091)

// Continue to check the other Prefixs instead of bailout due to one wrong IP

// Mark action

// return true when WARM_IP_TARGET or MINIMUM_IP_TARGET is defined
func (c *IPAMContext) warmIPTargetsDefined() bool { _ = "STUB: not implemented"; return false }

// max pods from instance type mapping file
type instanceTypeMaxPodsMapping map[string]int64

// getMaxPodsFromFile reads the max pods value from the eni-max-pods.txt file
// based on the instance type
func (c *IPAMContext) getMaxPodsFromFile() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// parseMaxPodsFile parses the eni-max-pods.txt file content and returns a mapping
// of instance type to max pods
func parseMaxPodsForInstanceFromFile(content string, instanceType string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Skip comments and empty lines

// Split the line into instance type and max pods

// UseCustomNetworkCfg returns whether Pods needs to use pod specific configuration or not.
func UseCustomNetworkCfg() bool { _ = "STUB: not implemented"; return false }

// ManageENIsOnNonSchedulableNode returns whether IPAMd should manage ENIs on the node or not.
func ManageENIsOnNonSchedulableNode() bool { _ = "STUB: not implemented"; return false }

// UseSubnetDiscovery returns whether we should use enhanced subnet selection or not when creating ENIs.
func UseSubnetDiscovery() bool { _ = "STUB: not implemented"; return false }

func parseBoolEnvVar(envVariableName string, defaultVal bool) bool {
	_ = "STUB: not implemented"
	return false
}

func dsBackingStorePath() string { _ = "STUB: not implemented"; return "" }

func getWarmIPTarget() int { _ = "STUB: not implemented"; return 0 }

func getMinimumIPTarget() int { _ = "STUB: not implemented"; return 0 }

func disableENIProvisioning() bool { _ = "STUB: not implemented"; return false }

func disableLeakedENICleanup() bool {
	_ = "STUB: not implemented"
	// Cases where leaked ENI cleanup is disabled:
	// 1. IPv6 is enabled, so no ENIs are attached
	// 2. ENI provisioning is disabled, so ENIs are not managed by IPAMD
	// 3. CNI is operating in IMDS only mode
	// 4. Environment var explicitly disabling task is set
	return false
}

func enableImdsOnlyMode() bool { _ = "STUB: not implemented"; return false }

func EnablePodENI() bool { _ = "STUB: not implemented"; return false }

func enableMultiNICSupport() bool { _ = "STUB: not implemented"; return false }

func getNetworkPolicyMode() (string, error) { _ = "STUB: not implemented"; return "", nil }

func usePrefixDelegation() bool { _ = "STUB: not implemented"; return false }

func isIPv4Enabled() bool { _ = "STUB: not implemented"; return false }

func isIPv6Enabled() bool { _ = "STUB: not implemented"; return false }

func enableManageUntaggedMode() bool { _ = "STUB: not implemented"; return false }

func EnablePodIPAnnotation() bool { _ = "STUB: not implemented"; return false }

// filterUnmanagedENIs filters out ENIs marked with the "node.k8s.amazonaws.com/no_manage" tag
func (c *IPAMContext) filterUnmanagedENIs(enis []awsutils.ENIMetadata) []awsutils.ENIMetadata {
	_ = "STUB: not implemented"
	return nil
}

// Filter out any Unmanaged ENIs. VPC CNI will only work with Primary ENI in IPv6 Prefix Delegation mode until
// we open up IPv6 support in Secondary IP and Custom networking modes. Filtering out the ENIs here will
// help us avoid myriad of if/else loops elsewhere in the code.
// We shouldn't need the IsPrimaryENI check as ENIs not created by vpc-cni will be marked unmanaged (including trunk ENI)

// datastoreTargetState determines the number of IPs `short` or `over` our WARM_IP_TARGET, accounting for the MINIMUM_IP_TARGET.
// With prefix delegation, this function determines the number of Prefixes `short` or `over`
func (c *IPAMContext) datastoreTargetState(stats *datastore.DataStoreStats, networkCard int) (short int, over int, enabled bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

// multi card ENIs will use WARM_IP_TARGET=1 and MINIMUM_IP_TARGET=1 by default

// there is no WARM_IP_TARGET defined and no MINIMUM_IP_TARGET, fallback to use all IP addresses on ENI

// Calculating DataStore stats can be expensive, so allow the caller to optionally pass stats it already calculated

// short is greater than 0 when we have fewer available IPs than the warm IP target

// short is greater than the warm IP target alone when we have fewer total IPs than the minimum target

// over is the number of available IPs we have beyond the warm IP target

// over is less than the warm IP target alone if it would imply reducing total IPs below the minimum target

// short : number of IPs short to reach warm targets
// over : number of IPs over the warm targets

// Number of prefixes IPAMD is short of to achieve warm targets

// Over will have number of IPs more than needed but with PD we would have allocated in chunks of /28
// Say assigned = 1, warm ip target = 16, this will need 2 prefixes. But over will return 15.
// Hence we need to check if 'over' number of IPs are needed to maintain the warm targets

// over will be number of prefixes over than needed but could be spread across used prefixes,
// say, after couple of pod churns, 3 prefixes are allocated with 1 IP each assigned and warm ip target is 15
// (J : is this needed? since we have to walk thru the loop of prefixes)

// datastorePrefixTargetState determines the number of prefixes short to reach WARM_PREFIX_TARGET
func (c *IPAMContext) datastorePrefixTargetState(networkCard int) (short int, enabled bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// /28 will consume 16 IPs so let's not allocate if not needed.

// setTerminating atomically sets the terminating flag.
func (c *IPAMContext) setTerminating() { _ = "STUB: not implemented"; return }

func (c *IPAMContext) isTerminating() bool { _ = "STUB: not implemented"; return false }

func (c *IPAMContext) isNodeNonSchedulable() bool { _ = "STUB: not implemented"; return false }

// Find my node

// GetConfigForDebug returns the active values of the configuration env vars (for debugging purposes).
func GetConfigForDebug() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func max(x, y int) int { _ = "STUB: not implemented"; return 0 }

func min(x, y int) int { _ = "STUB: not implemented"; return 0 }

func (c *IPAMContext) getTrunkLinkIndex() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// GetPod returns the pod matching the name and namespace
func (c *IPAMContext) GetPod(podName, namespace string) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AnnotatePod annotates the pod with the provided key and value
func (c *IPAMContext) AnnotatePod(podName string, podNamespace string, key string, newVal string, releasedIP string) error {
	_ = "STUB: not implemented"
	return nil
}

// if pod is nil and err is nil for any reason, this is not retriable case, returning a nil error to not-retry

// since the GetPod() error has been decorated, we have to check key words
// releasedIP is not empty meaning del path

// On CNI ADD, always set new annotation

// Skip patch operation if new value is the same as existing value

// On CNI DEL, set annotation to empty string if IP is the one we are releasing

func (c *IPAMContext) tryUnassignIPsFromENIs(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// From all datastores, get ENIInfos and unassign IPs

func (c *IPAMContext) tryUnassignIPFromENI(ctx context.Context, eniID string, networkCard int) {
	_ = "STUB: not implemented"
	return
}

// Delete IPs from datastore

// Don't force the delete, since a freeable IP might have been assigned to a pod
// before we get around to deleting it.
/* force */

// Deallocate IPs from the instance if they aren't used by pods.

func (c *IPAMContext) tryUnassignPrefixesFromENIs(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// From all datastores get ENIs and remove prefixes

func (c *IPAMContext) tryUnassignPrefixFromENI(ctx context.Context, eniID string, networkCard int) {
	_ = "STUB: not implemented"
	return
}

// Delete Prefixes from datastore

// Don't force the delete, since a freeable Prefix might have been assigned to a pod
// before we get around to deleting it.

// Determine if this is an IPv4 or IPv6 CIDR and call the appropriate deletion function

// IPv4 CIDR
/* force */

// IPv6 CIDR
/* force */

// Deallocate prefixes from the instance if they aren't used by pods.
// DeallocPrefixAddresses handles both IPv4 and IPv6 prefixes

// cleanupExcludedENI cleans up unassigned secondary IPs and prefixes from excluded ENI
func (c *IPAMContext) cleanupExcludedENI(ctx context.Context, eniID string) {
	_ = "STUB: not implemented"
	return
}

// Clean up based on the current IP mode

// In prefix delegation mode, cleanup unassigned IPv4 prefixes

// In secondary IP mode, cleanup unassigned IPv4 secondary IPs

// Clean up IPv6 prefixes if IPv6 is enabled (IPv6 only uses prefix delegation)

// checkAndHandleAllENIExclusion checks all existing ENIs (primary and secondary) across all network cards for subnet exclusion
func (c *IPAMContext) checkAndHandleAllENIExclusion(ctx context.Context) {
	_ = "STUB: not implemented"
	// Iterate over all datastores (one per network card)
	return
}

// Get ENI information using the public method

// Only process ENIs that are already marked as excluded

// Clean up unassigned resources from this excluded ENI

func (c *IPAMContext) GetENIResourcesToAllocate(networkCard int) int {
	_ = "STUB: not implemented"
	return 0
}

// IPv6 always uses PD, so it should always be 1

func (c *IPAMContext) GetIPv4Limit() (int, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

// Single PD - allocate one prefix per ENI and new add will be new ENI + prefix
// Multi - allocate one prefix per ENI and new add will be new prefix or new ENI + prefix

func (c *IPAMContext) isDatastorePoolEmpty(networkCard int) bool {
	_ = "STUB: not implemented"
	return false
}

// Return whether the maximum number of ENIs that can be attached to the node has already been reached
func (c *IPAMContext) hasRoomForEni(networkCard int) bool { _ = "STUB: not implemented"; return false }

// Count excluded ENIs to account for them in our calculations

// Subtract excluded ENIs from current count to match the exclusion from max count

// Check if we have room considering the excluded ENI

func (c *IPAMContext) isDatastorePoolTooLow() map[int]Decisions {
	_ = "STUB: not implemented"
	return nil
}

// If max pods has been reached, pool is not too low

func (c *IPAMContext) isDatastorePoolTooHigh(stats *datastore.DataStoreStats, networkCard int) bool {
	_ = "STUB: not implemented"
	// NOTE: IPs may be allocated in chunks (full ENIs of prefixes), so the "too-high" condition does not check max pods. The limit is enforced on the allocation side.
	return false
}

// For the existing ENIs check if we can cleanup prefixes

// We only ever report the pool being too high if WARM_IP_TARGET or WARM_PREFIX_TARGET is set

func (c *IPAMContext) warmPrefixTargetDefined() bool { _ = "STUB: not implemented"; return false }

// DeallocCidrs frees IPs and Prefixes from EC2
func (c *IPAMContext) DeallocCidrs(ctx context.Context, eniID string, deletableCidrs []datastore.CidrInfo) {
	_ = "STUB: not implemented"
	return
}

// Track the last time we unassigned Cidrs from an ENI. We won't reconcile any Cidrs in this cache
// for at least ipReconcileCooldown

// Track the last time we unassigned IPs from an ENI. We won't reconcile any IPs in this cache
// for at least ipReconcileCooldown

// getPrefixesNeeded returns the number of prefixes need to be allocated to the ENI
func (c *IPAMContext) getPrefixesNeeded(networkCard int) int {
	_ = "STUB: not implemented"
	// By default allocate 1 prefix at a time
	return 0
}

// TODO - post GA we can evaluate to see if these two calls can be merged.
// datastoreTargetState already has complex math so adding Prefix target will make it even more complex.

// WARM_IP_TARGET takes precendence over WARM_PREFIX_TARGET

func (c *IPAMContext) initENIAndIPLimits() (err error) { _ = "STUB: not implemented"; return nil }

// WARM and MAX ENI & IP/Prefix counts are no-op in IPv6 Prefix delegation mode. Once we start supporting IPv6 in
// Secondary IP mode, these variables will play the same role as they currently do in IPv4 mode. So, for now we
// leave them at their default values.

func (c *IPAMContext) isConfigValid() bool {
	_ = "STUB: not implemented"
	// Validate that only one among v4 and v6 is enabled.
	return false
}

// Validate PD mode is enabled if VPC CNI is operating in IPv6 mode. Custom networking is not supported in IPv6 mode.

// Validate Prefix Delegation against v4 and v6 modes.

func (c *IPAMContext) AddFeatureToCNINode(ctx context.Context, featureName rcv1alpha1.FeatureName, featureValue string) error {
	_ = "STUB: not implemented"
	return nil
}

// this should happen when user updated eniConfig name
// aws-node restarted and CNINode need to be updated if node wasn't terminated
// we need groom old features here

// SetAPIServerConnectivity updates the API server connectivity status and reconfigures
// components that depend on API server access
func (c *IPAMContext) SetAPIServerConnectivity(connected bool) { _ = "STUB: not implemented"; return }

// Status didn't change

// API server is now available - update maxPods from node object
// First, try to recreate the client with caching enabled

// Now get the node to update maxPods

// Update maxPods with the value from the node

// API server connection was lost
// No action needed here as we already have maxPods from file or kubelet
// and we want to keep working with that value

// excludedENIBasedOnSubnetTags excludes ENIs in datastore if subnets have valid tags
func (c *IPAMContext) excludedENIBasedOnSubnetTags(ctx context.Context, eni string, eniMetadata awsutils.ENIMetadata) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Check if this ENI (primary or secondary) is in an excluded subnet and mark it for exclusion

type Decisions struct {
	Stats *datastore.DataStoreStats
	IsLow bool
}
