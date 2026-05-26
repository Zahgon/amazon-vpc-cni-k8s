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

package datastore

import (
	"net"
	"sync"
	"time"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/netlinkwrapper"
	"github.com/vishvananda/netlink"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/utils/logger"
	"github.com/pkg/errors"
)

const (
	// minENILifeTime is the shortest time before we consider deleting a newly created ENI
	minENILifeTime = 1 * time.Minute

	// envIPCooldownPeriod (default 30 seconds) specifies the time after pod deletion before an IP can be assigned to a new pod
	envIPCooldownPeriod = "IP_COOLDOWN_PERIOD"

	// DuplicatedENIError is an error when caller tries to add an duplicate ENI to data store
	DuplicatedENIError = "data store: duplicate ENI"

	// IPAlreadyInStoreError is an error when caller tries to add an duplicate IP address to data store
	IPAlreadyInStoreError = "datastore: IP already in data store"

	// UnknownIPError is an error when caller tries to delete an IP which is unknown to data store
	UnknownIPError = "datastore: unknown IP"

	// IPInUseError is an error when caller tries to delete an IP where IP is still assigned to a Pod
	IPInUseError = "datastore: IP is used and can not be deleted"

	// ENIInUseError is an error when caller tries to delete an ENI where there are IP still assigned to a pod
	ENIInUseError = "datastore: ENI is used and can not be deleted"

	// UnknownENIError is an error when caller tries to access an ENI which is unknown to datastore
	UnknownENIError = "datastore: unknown ENI"
)

// On IPAMD/datastore restarts, we need to determine which pod IPs are already allocated. In VPC CNI <= 1.6,
// this was done by querying kubelet's CRI. This required scary access to the CRI socket, which could result
// in race conditions with the CRI's internal logic. Therefore, we transitioned to storing IP allocations
// ourselves in a persistent file (similar to host-ipam CNI plugin).
//
// The migration process took place over a series of phases:
//   Phase 0 (<= CNI 1.6): Read/write from CRI only
//   Phase 1 (CNI 1.7): Read from CRI. Write to CRI+file
//   Phase 2 (CNI 1.8): Read from file. Write to CRI+file
//   Phase 3 (CNI 1.12+): Read/write from file only
//
// Now that phase 3 has completed, IPAMD has no CRI dependency.

// Placeholders used for unknown values
const backfillNetworkName = "_migrated-from-cri"
const backfillNetworkIface = "unknown"

// ErrUnknownPod is an error when there is no pod in data store matching pod name, namespace, sandbox id
var ErrUnknownPod = errors.New("datastore: unknown pod")

// ErrNoAvailableIPInDataStore is an error when IPAM cannot assign an IP address from the datastore
var ErrNoAvailableIPInDataStore = errors.New("AssignPodIPAddress: no available IP/Prefix addresses")

// IPAMKey is the IPAM primary key.  Quoting CNI spec:
//
//	Plugins that store state should do so using a primary key of
//	(network name, CNI_CONTAINERID, CNI_IFNAME).
type IPAMKey struct {
	NetworkName string `json:"networkName"`
	ContainerID string `json:"containerID"`
	IfName      string `json:"ifName"`
}

// IsZero returns true if object is equal to the golang zero/null value.
func (k IPAMKey) IsZero() bool { _ = "STUB: not implemented"; return false }

// String() implements the fmt.Stringer interface.
func (k IPAMKey) String() string { _ = "STUB: not implemented"; return "" }

// IPAMMetadata is the metadata associated with IP allocations.
type IPAMMetadata struct {
	K8SPodNamespace string `json:"k8sPodNamespace,omitempty"`
	K8SPodName      string `json:"k8sPodName,omitempty"`
	InterfacesCount int    `json:"interfacesCount,omitempty"`
}

// ENI represents a single ENI. Exported fields will be marshaled for introspection.
type ENI struct {
	// AWS ENI ID
	ID         string
	createTime time.Time
	// IsPrimary indicates whether ENI is a primary ENI
	IsPrimary bool
	// IsTrunk indicates whether this ENI is used to provide pods with dedicated ENIs
	IsTrunk bool
	// IsEFA indicates whether this ENI is tagged as an EFA
	IsEFA bool
	// DeviceNumber is the device number of ENI (0 means the primary ENI)
	DeviceNumber int
	// IPv4Addresses shows whether each address is assigned, the key is IP address, which must
	// be in dot-decimal notation with no leading zeros and no whitespace(eg: "10.1.0.253")
	// Key is the IP address - PD: "IP/28" and SIP: "IP/32"
	AvailableIPv4Cidrs map[string]*CidrInfo
	//IPv6CIDRs contains information tied to IPv6 Prefixes attached to the ENI
	IPv6Cidrs map[string]*CidrInfo
	// IsExcludedForPodIPs indicates whether this ENI should be excluded from pod IP allocation
	IsExcludedForPodIPs bool
	// RouteTableID is the route table ID associated with the ENI on the host
	RouteTableID int
	// SubnetID is the subnet which the ENI was created
	SubnetID string
}

// AddressInfo contains information about an IP, Exported fields will be marshaled for introspection.
type AddressInfo struct {
	Address string

	IPAMKey        IPAMKey
	IPAMMetadata   IPAMMetadata
	AssignedTime   time.Time
	UnassignedTime time.Time
}

// CidrInfo
type CidrInfo struct {
	// Either v4/v6 Host or LPM Prefix
	Cidr net.IPNet
	// Key is individual IP addresses from the Prefix - /32 (v4) or /128 (v6)
	IPAddresses map[string]*AddressInfo
	// true if Cidr here is an LPM prefix
	IsPrefix bool
	// IP Address Family of the Cidr
	AddressFamily string
}

func (cidr *CidrInfo) Size() int { _ = "STUB: not implemented"; return 0 }

func (e *ENI) findAddressForSandbox(ipamKey IPAMKey) (*CidrInfo, *AddressInfo) {
	_ = "STUB: not implemented"
	// Either v4 or v6 for now.
	// Check in V4 prefixes
	return nil, nil
}

// Check in V6 prefixes

// AssignedIPv4Addresses is the number of IP addresses already assigned
func (e *ENI) AssignedIPv4Addresses() int { _ = "STUB: not implemented"; return 0 }

// AssignedIPv6Addresses is the number of IPv6 addresses already assigned
func (e *ENI) AssignedIPv6Addresses() int { _ = "STUB: not implemented"; return 0 }

// AssignedIPAddressesInCidr is the number of IP addresses already assigned in the IPv4 CIDR
func (cidr *CidrInfo) AssignedIPAddressesInCidr() int {
	_ = "STUB: not implemented"

	// SIP : This will run just once and count will be 0 if addr is not assigned or addr is not allocated yet(unused IP)
	// PD : This will return count of number /32 assigned in /28 CIDR.
	return 0
}

type CidrStats struct {
	AssignedIPs int
	CooldownIPs int
}

// Gets number of assigned IPs and the IPs in cooldown from a given CIDR
func (cidr *CidrInfo) GetIPStatsFromCidr(ipCooldownPeriod time.Duration) CidrStats {
	_ = "STUB: not implemented"
	return *new(CidrStats)
}

// Assigned returns true iff the address is allocated to a pod/sandbox.
func (addr AddressInfo) Assigned() bool { _ = "STUB: not implemented"; return false }

// getCooldownPeriod returns the time duration in seconds configured by the IP_COOLDOWN_PERIOD env variable
func getCooldownPeriod() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// InCoolingPeriod checks whether an addr is in ipCooldownPeriod
func (addr AddressInfo) inCoolingPeriod(ipCooldownPeriod time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

// ENIPool is a collection of ENI, keyed by ENI ID
type ENIPool map[string]*ENI

// AssignedIPv4Addresses is the number of IP addresses already assigned
func (p *ENIPool) AssignedIPv4Addresses() int { _ = "STUB: not implemented"; return 0 }

// FindAddressForSandbox returns ENI and AddressInfo or (nil, nil) if not found
func (p *ENIPool) FindAddressForSandbox(ipamKey IPAMKey) (*ENI, *CidrInfo, *AddressInfo) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// PodIPInfo contains pod's IP and the device number of the ENI
type PodIPInfo struct {
	IPAMKey IPAMKey
	// IP is the IPv4 address of pod
	IP string
	// DeviceNumber is the device number of the ENI
	DeviceNumber int
}

// DataStore contains node level ENI/IP
type DataStore struct {
	total            int
	assigned         int
	allocatedPrefix  int
	eniPool          ENIPool
	lock             sync.Mutex
	log              logger.Logger
	backingStore     Checkpointer
	netLink          netlinkwrapper.NetLink
	isPDEnabled      bool
	ipCooldownPeriod time.Duration
	networkCard      int
}

// ENIInfos contains ENI IP information
type ENIInfos struct {
	// TotalIPs is the total number of IP addresses
	TotalIPs int
	// assigned is the number of IP addresses that has been assigned
	AssignedIPs int
	// ENIs contains ENI IP pool information
	ENIs map[string]ENI
}

// NewDataStore returns DataStore structure
func NewDataStore(log logger.Logger, backingStore Checkpointer, isPDEnabled bool, networkCard int) *DataStore {
	_ = "STUB: not implemented"
	return nil
}

// CheckpointFormatVersion is the version stamp used on stored checkpoints.
const CheckpointFormatVersion = "vpc-cni-ipam/1"

// CheckpointData is the format of stored checkpoints. Note this is
// deliberately a "dumb" format since efficiency is less important
// than version stability here.
type CheckpointData struct {
	Version     string            `json:"version"`
	Allocations []CheckpointEntry `json:"allocations"`
}

// CheckpointEntry is a "row" in the conceptual IPAM datastore, as stored
// in checkpoints.
type CheckpointEntry struct {
	IPAMKey
	IPv4                string       `json:"ipv4,omitempty"`
	IPv6                string       `json:"ipv6,omitempty"`
	AllocationTimestamp int64        `json:"allocationTimestamp"`
	Metadata            IPAMMetadata `json:"metadata"`
}

// ReadBackingStore initializes the IP allocation state from the
// configured backing store. Should be called before using data store.
func (ds *DataStore) ReadBackingStore(isv6Enabled bool) error {
	_ = "STUB: not implemented"
	return nil

	// Read from checkpoint file
}

// Assume that no file == no containers are currently in use, e.g. a fresh reboot just cleared everything out.
// This is ok, and no-op.

// Found!

// Increment ENI IP usage upon finding assigned ips

// Update prometheus for ips per cidr
// Secondary IP mode will have /32:1 and Prefix mode will have /28:<number of /32s>

// Some entries may have been purged during recovery, so write to backing store

func (ds *DataStore) writeBackingStoreUnsafe() error { _ = "STUB: not implemented"; return nil }

// Loop through ENI's v4 prefixes

// Loop through ENI's v6 prefixes

// AddENI add ENI to data store
func (ds *DataStore) AddENI(eniID string, deviceNumber int, isPrimary, isTrunk, isEFA bool, routeTableID int, subnetID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Initialize ENI IPs In Use to 0 when an ENI is created

// SetENIExcludedForPodIPs marks an ENI as excluded from pod IP allocation
func (ds *DataStore) SetENIExcludedForPodIPs(eniID string, excluded bool) error {
	_ = "STUB: not implemented"
	return nil
}

// IsENIExcludedForPodIPs returns whether an ENI is excluded from pod IP allocation
func (ds *DataStore) IsENIExcludedForPodIPs(eniID string) bool {
	_ = "STUB: not implemented"
	return false
}

// AddIPv4AddressToStore adds IPv4 CIDR of an ENI to data store
func (ds *DataStore) AddIPv4CidrToStore(eniID string, ipv4Cidr net.IPNet, isPrefix bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Already there

func (ds *DataStore) DelIPv4CidrFromStore(eniID string, cidr net.IPNet, force bool) error {
	_ = "STUB: not implemented"
	return nil
}

// SIP case : This runs just once
// PD case : if (force is false) then if there are any unassigned IPs, those will get freed but the first assigned IP will
//  break the loop, should be fine since freed IPs will be reused for new pods.

// Continuing because 'force'

// DelIPv6CidrFromStore deletes IPv6 CIDR from the datastore
func (ds *DataStore) DelIPv6CidrFromStore(eniID string, cidr net.IPNet, force bool) error {
	_ = "STUB: not implemented"
	return nil
}

// IPv6 only uses prefix delegation, check for any assigned IPs

// Continuing because 'force'

// AddIPv6AddressToStore adds IPv6 CIDR of an ENI to data store
func (ds *DataStore) AddIPv6CidrToStore(eniID string, ipv6Cidr net.IPNet, isPrefix bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if already present in datastore.

func (ds *DataStore) AssignPodIPAddress(ipamKey IPAMKey, ipamMetadata IPAMMetadata, isIPv4Enabled bool, isIPv6Enabled bool) (ipv4Address string,
	ipv6Address string, deviceNumber int, routeTableId int, err error) {
	_ = "STUB: not implemented"
	//Currently it's either v4 or v6. Dual Stack mode isn't supported.
	return "", "", 0, 0, nil
}

// AssignPodIPv6Address assigns an IPv6 address to a pod.
// Returns the assigned IPv6 address, device number, route table ID, and an error.
func (ds *DataStore) AssignPodIPv6Address(ipamKey IPAMKey, ipamMetadata IPAMMetadata) (ipv6Address string, deviceNumber int, routeTableId int, err error) {
	_ = "STUB: not implemented"
	return "", 0, 0, nil
}

// In IPv6 Prefix Delegation mode, eniPool will only have Primary ENI.

// Skip ENIs that are excluded from pod IP allocation

//In v6 mode, we (should) only have one CIDR/Prefix. So, we can bail out but we will let the loop
//exit instead.

// Important! Unwind assignment

//Remove the IP from eni DB

// Increment ENI IP usage on pod IPv6 allocation

// AssignPodIPv4Address assigns an IPv4 address to pod
// It returns the assigned IPv4 address, device number, route table ID, and error
func (ds *DataStore) AssignPodIPv4Address(ipamKey IPAMKey, ipamMetadata IPAMMetadata) (ipv4address string, deviceNumber int, routeTableId int, err error) {
	_ = "STUB: not implemented"
	return "", 0, 0, nil
}

// Skip ENIs that are excluded from pod IP allocation

// Check in next CIDR

// Update prometheus for ips per cidr
// Secondary IP mode will have /32:1 and Prefix mode will have /28:<number of /32s>

// This can happen during upgrade or PD enable/disable knob toggle
// ENI can have prefixes attached and no space for SIPs or vice versa

// addr is nil when we are using a new IP from prefix or SIP pool
// if addr is out of cooldown or not assigned, we can reuse addr

// IP was previously used and is being reassigned

// Important! Unwind assignment

// Remove the IP from eni DB

// Update prometheus for ips per cidr

// Increment ENI IP usage on pod IPv4 allocation

// assignPodIPAddressUnsafe mark Address as assigned.
func (ds *DataStore) assignPodIPAddressUnsafe(addr *AddressInfo, ipamKey IPAMKey, ipamMetadata IPAMMetadata, assignedTime time.Time) {
	_ = "STUB: not implemented"
	return
}

// This marks the addr as assigned

// Prometheus gauge

// unassignPodIPAddressUnsafe mark Address as unassigned.
func (ds *DataStore) unassignPodIPAddressUnsafe(addr *AddressInfo) {
	_ = "STUB: not implemented"
	return

	// Already unassigned
}

// unassign the addr

// Prometheus gauge

type DataStoreStats struct {
	// Total number of addresses allocated
	TotalIPs int
	// Total number of prefixes allocated
	TotalPrefixes int

	// Number of assigned addresses
	AssignedIPs int
	// Number of addresses in cooldown
	CooldownIPs int
}

func (stats *DataStoreStats) String() string { _ = "STUB: not implemented"; return "" }

func (stats *DataStoreStats) AvailableAddresses() int { _ = "STUB: not implemented"; return 0 }

// GetIPStats returns DataStoreStats for addressFamily
func (ds *DataStore) GetIPStats(addressFamily string) *DataStoreStats {
	_ = "STUB: not implemented"
	return nil
}

// Skip excluded ENIs when calculating available IPs for pod allocation

// GetTrunkENI returns the trunk ENI ID or an empty string
func (ds *DataStore) GetTrunkENI() string { _ = "STUB: not implemented"; return "" }

// GetEFAENIs returns the a map containing all attached EFA ENIs
func (ds *DataStore) GetEFAENIs() map[string]bool { _ = "STUB: not implemented"; return nil }

// IsRequiredForWarmIPTarget determines if this ENI has warm IPs that are required to fulfill whatever WARM_IP_TARGET is set to.
func (ds *DataStore) isRequiredForWarmIPTarget(warmIPTarget int, eni *ENI) bool {
	_ = "STUB: not implemented"
	return false
}

// Skip excluded ENIs when calculating warm IPs for pod allocation

// IsRequiredForMinimumIPTarget determines if this ENI is necessary to fulfill whatever MINIMUM_IP_TARGET is set to.
func (ds *DataStore) isRequiredForMinimumIPTarget(minimumIPTarget int, eni *ENI) bool {
	_ = "STUB: not implemented"
	return false
}

// Skip excluded ENIs when calculating total IPs for pod allocation

// IsRequiredForWarmPrefixTarget determines if this ENI is necessary to fulfill whatever WARM_PREFIX_TARGET is set to.
func (ds *DataStore) isRequiredForWarmPrefixTarget(warmPrefixTarget int, eni *ENI) bool {
	_ = "STUB: not implemented"
	return false
}

// Skip excluded ENIs when calculating free prefixes for pod allocation

func (ds *DataStore) getDeletableENI(warmIPTarget, minimumIPTarget, warmPrefixTarget int) *ENI {
	_ = "STUB: not implemented"
	return nil
}

// IsTooYoung returns true if the ENI hasn't been around long enough to be deleted.
func (e *ENI) isTooYoung() bool { _ = "STUB: not implemented"; return false }

// HasIPInCooling returns true if an IP address was unassigned recently.
func (e *ENI) hasIPInCooling(ipCooldownPeriod time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

// HasPods returns true if the ENI has pods assigned to it.
func (e *ENI) hasPods() bool { _ = "STUB: not implemented"; return false }

// GetAllocatableENIs finds ENIs in the datastore that needs more IP addresses allocated
func (ds *DataStore) GetAllocatableENIs(maxIPperENI int, skipPrimary bool) []*ENI {
	_ = "STUB: not implemented"
	return nil
}

// Skip ENIs that are excluded from pod IP allocation

// RemoveUnusedENIFromStore removes a deletable ENI from the data store.
// It returns the name of the ENI which has been removed from the data store and needs to be deleted,
// or empty string if no ENI could be removed.
func (ds *DataStore) RemoveUnusedENIFromStore(warmIPTarget, minimumIPTarget, warmPrefixTarget int) string {
	_ = "STUB: not implemented"
	return ""
}

// Prometheus update

// Delete ENI IPs In Use when ENI is removed

// RemoveENIFromDataStore removes an ENI from the datastore. It returns nil on success, or an error.
func (ds *DataStore) RemoveENIFromDataStore(eniID string, force bool) error {
	_ = "STUB: not implemented"
	return nil
}

// This scenario can occur if the reconciliation process discovered this ENI was detached
// from the EC2 instance outside of the control of ipamd. If this happens, there's nothing
// we can do other than force all pods to be unassigned from the IPs on this ENI.

// Continuing, because 'force'

// Prometheus gauge

// Delete ENI IPs In Use when ENI is removed

// UnassignPodIPAddress:
// a) Finds the IP address based on PodName and PodNamespace.
// b) Marks the IP address as unassigned.
// Returns:
//   - *ENI: the ENI object associated with the IP address
//   - ip string: the IP address being unassigned
//   - deviceNumber int: the ENI's device number
//   - interfaces int: the number of interfaces associated with the pod
//   - routeTableId int: the ENI's route table ID
//   - err error: error if any occurred during unassignment
func (ds *DataStore) UnassignPodIPAddress(ipamKey IPAMKey) (e *ENI, ip string, deviceNumber int, interfaces int, routeTableId int, err error) {
	_ = "STUB: not implemented"
	return nil, "", 0, 0, 0, nil
}

// If the entry is not present in state file, check if it is present under placeholder value.
// This scenario could happen if the pod was created by an older CNI version back when CRI read was done.

// If entry is still not found, IPAMD has no knowledge of this pod, so there is nothing to do.

// Unwind un-assignment

// Debug log for IP entering cooldown

// Interfaces Count 0 means this property did not exist in the datastore when we restored. A Pod entry always have atleast one interface

//Update prometheus for ips per cidr

// Decrement ENI IP usage when a pod is deallocated

// Check if ENI is excluded and CIDR is now empty - cleanup if needed

// Schedule async cleanup of the empty CIDR

// AllocatedIPs returns a recent snapshot of allocated sandbox<->IPs.
// Note result may already be stale by the time you look at it.
func (ds *DataStore) AllocatedIPs() []PodIPInfo { _ = "STUB: not implemented"; return nil }

// FreeableIPs returns a list of unused and potentially freeable IPs.
// Note result may already be stale by the time you look at it.
func (ds *DataStore) FreeableIPs(eniID string) []net.IPNet { _ = "STUB: not implemented"; return nil }

// Can't free any IPs from an ENI we don't know about...

// FreeablePrefixes returns a list of unused and potentially freeable prefixes (both IPv4 and IPv6).
// Note result may already be stale by the time you look at it.
func (ds *DataStore) FreeablePrefixes(eniID string) []net.IPNet {
	_ = "STUB: not implemented"
	return nil
}

// Can't free any Prefixes from an ENI we don't know about...

// Check IPv4 prefixes

// Check IPv6 prefixes (IPv6 only uses prefix delegation mode)

// GetENIInfos provides ENI and IP information about the datastore
func (ds *DataStore) GetENIInfos() *ENIInfos { _ = "STUB: not implemented"; return nil }

// Since IP Addresses might get removed, we need to make a deep copy here.

// Since IP Addresses might get removed, we need to make a deep copy here.

// GetENIs provides the number of ENI in the datastore
func (ds *DataStore) GetENIs() int { _ = "STUB: not implemented"; return 0 }

// GetENICIDRs returns the known (allocated & unallocated) ENI secondary IPs and Prefixes
func (ds *DataStore) GetENICIDRs(eniID string) ([]string, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetFreePrefixes return free prefixes
func (ds *DataStore) GetFreePrefixes() int { _ = "STUB: not implemented"; return 0 }

// getFreeIPv4AddrfromCidr returs a free IP/32 address from CIDR
func (ds *DataStore) getFreeIPv4AddrfromCidr(availableCidr *CidrInfo) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ds *DataStore) getFreeIPv6AddrFromCidr(IPv6Cidr *CidrInfo) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

/*
  /28 -> 10.1.1.1/32[out of cooldown], 10.1.1.2/32[out of cooldown], 10.1.1.3/32[assigned]
  cached ip = 10.1.1.1/32
  /28 ->  10.1.1.3/32[assigned]
  return 10.1.1.1/32
*/

func (ds *DataStore) getUnusedIP(availableCidr *CidrInfo) (string, error) {
	_ = "STUB: not implemented"
	//Check if there is any IP out of cooldown
	return "", nil
}

//if the IP is out of cooldown and not assigned then cache the first available IP
//continue cleaning up the DB, this is to avoid stale entries and a new thread :)

// Debug log for IP released from cooldown

//availableCidr.IPAddresses[addr.Address] = nil //Avoid mem leak - TODO

//If not in cooldown then generate next IP

func (ds *DataStore) GetNetworkCard() int { _ = "STUB: not implemented"; return 0 }

func getNextIPAddr(ip net.IP) { _ = "STUB: not implemented"; return }

// Function to return PD defaults supported by VPC
func GetPrefixDelegationDefaults() (int, int, int) { _ = "STUB: not implemented"; return 0, 0, 0 }

// FindFreeableCidrs finds and returns Cidrs that are not assigned to Pods but are attached
// to ENIs on the node.
func (ds *DataStore) FindFreeableCidrs(eniID string) []CidrInfo {
	_ = "STUB: not implemented"
	return nil
}

// Can't free any Cidrs from an ENI we don't know about...

func DivCeil(x, y int) int { _ = "STUB: not implemented"; return 0 }

// CheckFreeableENIexists will return true if there is an ENI which is unused.
// Could have just called getDeletaleENI, this is just to optimize a bit.
func (ds *DataStore) CheckFreeableENIexists() bool { _ = "STUB: not implemented"; return false }

// NormalizeCheckpointDataByPodVethExistence will normalize checkpoint data by removing allocations that do not have a corresponding pod veth.
// This can happen if pods are deleted while IPAMD is inactive.
func (ds *DataStore) normalizeCheckpointDataByPodVethExistence(checkpoint CheckpointData) (CheckpointData, error) {
	_ = "STUB: not implemented"
	return *new(CheckpointData), nil
}

// Stale allocations may have dangling IP rules that need cleanup

func (ds *DataStore) validateAllocationByPodVethExistence(allocation CheckpointEntry, hostNSLinks []netlink.Link) error {
	_ = "STUB: not implemented"
	// for backwards compatibility, we skip the validation when metadata contains empty namespace/name.
	return nil
}

// For each stale allocation, cleanup leaked IP rules if they exist
func (ds *DataStore) PruneStaleAllocations(staleAllocations []CheckpointEntry) {
	_ = "STUB: not implemented"
	return
}

func (ds *DataStore) DeleteToContainerRule(entry *CheckpointEntry) {
	_ = "STUB: not implemented"
	return
}

// Remove toContainer rule, if it exists. Note that toContainer rule will always be in main routing table.

// Continue to prune, even on deletion error

func (ds *DataStore) DeleteFromContainerRule(entry *CheckpointEntry) {
	_ = "STUB: not implemented"
	return
}

// Remove fromContainer rule, if it exists. Note that fromContainer rule can be in any routing table,
// so no table is set.

// Continue to prune, even on deletion error

type DataStoreAccess struct {
	DataStores []*DataStore
}

func InitializeDataStores(skipNetworkCards []bool, defaultDataStorePath string, enablePD bool, log logger.Logger) *DataStoreAccess {
	_ = "STUB: not implemented"
	return nil
}

func (ds *DataStoreAccess) GetDataStore(networkCard int) *DataStore {
	_ = "STUB: not implemented"
	return nil
}

func (ds *DataStoreAccess) ReadAllDataStores(enableIPv6 bool) error {
	_ = "STUB: not implemented"
	return nil
}

// deallocateEmptyCIDR asynchronously deallocates an empty CIDR from an excluded ENI
func (ds *DataStore) deallocateEmptyCIDR(eniID string, cidrToCleanup *CidrInfo) {
	_ = "STUB: not implemented"
	return
}

// Add delay to avoid race conditions with pod cleanup

// Double-check that the CIDR is still empty and ENI is still excluded

// Find the CIDR in the appropriate map using AddressFamily

// Check if CIDR is still empty

// Remove the empty CIDR from the ENI structure

// Remove the CIDR from the appropriate ENI map

// Update backing store

// Note: We continue since the CIDR removal from local state was successful
