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

// Package awsutils is a utility package for calling EC2 or IMDS
package awsutils

import (
	"context"
	"net"
	"regexp"
	"sync"
	"time"

	"github.com/aws/smithy-go"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/ipamd/datastore"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/ec2wrapper"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/utils/logger"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/vpc"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/sets"
)

const (
	maxENIEC2APIRetries  = 12
	maxENIBackoffDelay   = time.Minute
	eniDescriptionPrefix = "aws-K8S-"

	// AllocENI need to choose a first free device number between 0 and maxENI
	// 100 is a hard limit because we use vlanID + 100 for pod networking table names
	maxENIs           = 100
	clusterNameEnvVar = "CLUSTER_NAME"

	// clusterTagKeyPrefix is the prefix for the cluster-specific subnet tags
	clusterTagKeyPrefix = "cni.networking.k8s.aws/cluster/"

	eniNodeTagKey                   = "node.k8s.amazonaws.com/instance_id"
	eniCreatedAtTagKey              = "node.k8s.amazonaws.com/createdAt"
	eniClusterTagKey                = "cluster.k8s.amazonaws.com/name"
	eniOwnerTagKey                  = "eks:eni:owner"
	eniOwnerTagValue                = "amazon-vpc-cni"
	additionalEniTagsEnvVar         = "ADDITIONAL_ENI_TAGS"
	reservedTagKeyPrefix            = "k8s.amazonaws.com"
	subnetDiscoveryTagKey           = "kubernetes.io/role/cni"
	subnetDiscoveryTagValueIncluded = "1"
	subnetDiscoveryTagValueExcluded = "0"
	envVpcCniVersion                = "VPC_CNI_VERSION"

	// UnknownInstanceType indicates that the instance type is not yet supported
	UnknownInstanceType = "vpc ip resource(eni ip limit): unknown instance type"

	// Stagger cleanup start time to avoid calling EC2 too much. Time in seconds.
	eniCleanupStartupDelayMax = 300
	eniDeleteCooldownTime     = 5 * time.Minute

	// the default page size when paginating the DescribeNetworkInterfaces call
	describeENIPageSize = 1000
)

var (
	awsAPIError        smithy.APIError
	awsGenericAPIError *smithy.GenericAPIError

	// ErrENINotFound is an error when ENI is not found.
	ErrENINotFound = errors.New("ENI is not found")
	// ErrAllSecondaryIPsNotFound is returned when not all secondary IPs on an ENI have been assigned
	ErrAllSecondaryIPsNotFound = errors.New("All secondary IPs not found")
	// ErrNoSecondaryIPsFound is returned when not all secondary IPs on an ENI have been assigned
	ErrNoSecondaryIPsFound = errors.New("No secondary IPs have been assigned to this ENI")
	// ErrNoNetworkInterfaces occurs when DescribeNetworkInterfaces(eniID) returns no network interfaces
	ErrNoNetworkInterfaces = errors.New("No network interfaces found for ENI")
	// ErrUnableToDetachENI is returned when the ENI cannot be detached from the instance
	ErrUnableToDetachENI = errors.New("unable to detach ENI from EC2 instance, giving up")
	// ErrENIAttachmentIdNotFound is returned when the ENI attachment ID is not found
	ErrENIAttachmentIdNotFound = errors.New("ENI attachment ID not found")
)

var log = logger.Get()

// APIs defines interfaces calls for adding/getting/deleting ENIs/secondary IPs. The APIs are not thread-safe.
type APIs interface {
	// AllocENI creates an ENI and attaches it to the instance
	AllocENI(ctx context.Context, sg []*string, eniCfgSubnet string, numIPs int, networkCard int) (eni string, err error)

	// FreeENI detaches ENI interface and deletes it
	FreeENI(ctx context.Context, eniName string) error

	// TagENI Tags ENI with current tags to contain expected tags.
	TagENI(ctx context.Context, eniID string, currentTags map[string]string) error

	// GetAttachedENIs retrieves eni information from instance metadata service
	GetAttachedENIs() (eniList []ENIMetadata, err error)

	// GetIPv4sFromEC2 returns the IPv4 addresses for a given ENI
	GetIPv4sFromEC2(ctx context.Context, eniID string) (addrList []ec2types.NetworkInterfacePrivateIpAddress, err error)

	// GetIPv4PrefixesFromEC2 returns the IPv4 prefixes for a given ENI
	GetIPv4PrefixesFromEC2(ctx context.Context, eniID string) (addrList []ec2types.Ipv4PrefixSpecification, err error)

	// GetIPv6PrefixesFromEC2 returns the IPv6 prefixes for a given ENI
	GetIPv6PrefixesFromEC2(ctx context.Context, eniID string) (addrList []ec2types.Ipv6PrefixSpecification, err error)

	// DescribeAllENIs calls EC2 and returns a fully populated DescribeAllENIsResult struct and an error
	DescribeAllENIs(ctx context.Context) (DescribeAllENIsResult, error)

	// AllocIPAddress allocates an IP address for an ENI
	AllocIPAddress(ctx context.Context, eniID string) error

	// AllocIPAddresses allocates numIPs IP addresses on a ENI
	AllocIPAddresses(ctx context.Context, eniID string, numIPs int) (*ec2.AssignPrivateIpAddressesOutput, error)

	// DeallocIPAddresses deallocates the list of IP addresses from a ENI
	DeallocIPAddresses(ctx context.Context, eniID string, ips []string) error

	// DeallocPrefixAddresses deallocates the list of IP addresses from a ENI
	DeallocPrefixAddresses(ctx context.Context, eniID string, ips []string) error

	// AllocIPv6Prefixes allocates IPv6 prefixes to the ENI passed in
	AllocIPv6Prefixes(ctx context.Context, eniID string) ([]*string, error)

	// GetVPCIPv4CIDRs returns VPC's IPv4 CIDRs from instance metadata
	GetVPCIPv4CIDRs() ([]string, error)

	// GetLocalIPv4 returns the primary IPv4 address on the primary ENI interface
	GetLocalIPv4() net.IP

	// GetLocalIPv6 returns the primary IPv6 address on the primary ENI interface
	GetLocalIPv6() net.IP

	// GetVPCIPv6CIDRs returns VPC's IPv6 CIDRs from instance metadata
	GetVPCIPv6CIDRs() ([]string, error)

	// GetPrimaryENI returns the primary ENI
	GetPrimaryENI() string

	// GetENIIPv4Limit return IP address limit per ENI based on EC2 instance type
	GetENIIPv4Limit() int

	// GetENILimit returns the number of ENIs that can be attached to an instance
	GetENILimit() int

	// GetNetworkCards returns the network cards the instance has
	GetNetworkCards() []vpc.NetworkCard

	// GetPrimaryENImac returns the mac address of the primary ENI
	GetPrimaryENImac() string

	// SetUnmanagedENIs sets the list of unmanaged ENI IDs
	SetUnmanagedENIs(eniIDs []string)

	// SetUnmanagedNetworkCards sets the list of unmanaged Network Cards
	SetUnmanagedNetworkCards(skipNetworkCards []bool)

	// Set EFAOnlyENIs
	SetEFAOnlyENIs(efaOnlyENIByNetworkCard []string)

	// IsUnmanagedENI checks if an ENI is unmanaged
	IsUnmanagedENI(eniID string) bool

	// IsUnmanagedNIC checks if an Network Card is unmanaged
	IsUnmanagedNIC(networkCard int) bool

	// IsEfaOnlyENI checks if an ENI is efa-only
	IsEfaOnlyENI(networkCard int, eni string) bool

	// WaitForENIAndIPsAttached waits until the ENI has been attached and the secondary IPs have been added
	WaitForENIAndIPsAttached(eni string, wantedSecondaryIPs int) (ENIMetadata, error)

	// IsPrimaryENI
	IsPrimaryENI(eniID string) bool

	// RefreshSGIDs
	RefreshSGIDs(ctx context.Context, mac string, ds *datastore.DataStoreAccess) error

	// RefreshCustomSGIDs discovers and refreshes security groups tagged with kubernetes.io/role/cni=1
	RefreshCustomSGIDs(ctx context.Context, dsAccess *datastore.DataStoreAccess) error

	// GetInstanceHypervisorFamily returns the hypervisor family for the instance
	GetInstanceHypervisorFamily() string

	// GetInstanceType returns the EC2 instance type
	GetInstanceType() string

	// Update cached prefix delegation flag
	InitCachedPrefixDelegation(bool)

	// GetInstanceID returns the instance ID
	GetInstanceID() string

	// FetchInstanceTypeLimits Verify if the InstanceNetworkingLimits has the ENI limits else make EC2 call to fill cache.
	FetchInstanceTypeLimits(ctx context.Context) error

	IsPrefixDelegationSupported() bool

	IsTrunkingCompatible() bool

	// GetENISubnetID gets the subnet ID for an ENI from AWS
	GetENISubnetID(ctx context.Context, eniID string) (string, error)

	// GetVpcSubnets returns all subnets in the VPC
	GetVpcSubnets(ctx context.Context) ([]ec2types.Subnet, error)

	// IsSubnetExcluded returns if a subnet is excluded for pod IPs based on its tags
	IsSubnetExcluded(ctx context.Context, subnetID string) (bool, error)
}

// EC2InstanceMetadataCache caches instance metadata
type EC2InstanceMetadataCache struct {
	// metadata info
	securityGroups           StringSet
	customSecurityGroups     StringSet
	subnetID                 string
	localIPv4                net.IP
	v4Enabled                bool
	v6Enabled                bool
	instanceID               string
	instanceType             string
	primaryENI               string
	primaryENImac            string
	availabilityZone         string
	region                   string
	vpcID                    string
	unmanagedENIs            StringSet
	useCustomNetworking      bool
	unmanagedNICs            []bool
	efaOnlyENIsByNetworkCard []string
	useSubnetDiscovery       bool
	enablePrefixDelegation   bool
	clusterName              string
	additionalENITags        map[string]string
	imds                     TypedIMDS
	ec2SVC                   ec2wrapper.EC2
	connectionTrackingSpec   *ec2types.ConnectionTrackingSpecificationRequest
}

// ENIMetadata contains information about an ENI
type ENIMetadata struct {
	// ENIID is the id of network interface
	ENIID string

	// MAC is the mac address of network interface
	MAC string

	// DeviceNumber is the  device number of network interface
	DeviceNumber int // 0 means it is primary interface

	// SubnetIPv4CIDR is the IPv4 CIDR of network interface
	SubnetIPv4CIDR string

	// SubnetIPv6CIDR is the IPv6 CIDR of network interface
	SubnetIPv6CIDR string

	// The ip addresses allocated for the network interface
	IPv4Addresses []ec2types.NetworkInterfacePrivateIpAddress

	// IPv4 Prefixes allocated for the network interface
	IPv4Prefixes []ec2types.Ipv4PrefixSpecification

	// IPv6 addresses allocated for the network interface
	IPv6Addresses []ec2types.NetworkInterfaceIpv6Address

	// IPv6 Prefixes allocated for the network interface
	IPv6Prefixes []ec2types.Ipv6PrefixSpecification

	// Network card the ENI is attached on
	NetworkCard int

	// SubnetID the ENI is created from
	SubnetID string
}

// PrimaryIPv4Address returns the primary IPv4 address of this node
func (eni ENIMetadata) PrimaryIPv4Address() string { _ = "STUB: not implemented"; return "" }

// PrimaryIPv6Address returns the primary IPv6 address of this node
func (eni ENIMetadata) PrimaryIPv6Address() string { _ = "STUB: not implemented"; return "" }

// TagMap keeps track of the EC2 tags on each ENI
type TagMap map[string]string

// DescribeAllENIsResult contains the fully
type DescribeAllENIsResult struct {
	ENIMetadata             []ENIMetadata
	TagMap                  map[string]TagMap
	TrunkENI                string
	EFAENIs                 map[string]bool
	EFAOnlyENIByNetworkCard []string
	ENIsByNetworkCard       [][]string
}

// msSince returns milliseconds since start.
func msSince(start time.Time) float64 { _ = "STUB: not implemented"; return 0 }

// StringSet is a set of strings
type StringSet struct {
	sync.RWMutex
	data sets.String
}

// SortedList returns a sorted string slice from this set
func (ss *StringSet) SortedList() []string { _ = "STUB: not implemented"; return nil }

// sets.String.List() returns a sorted list

// Set sets the string set
func (ss *StringSet) Set(items []string) { _ = "STUB: not implemented"; return }

// Difference compares this StringSet with another
func (ss *StringSet) Difference(other *StringSet) *StringSet { _ = "STUB: not implemented"; return nil }

// example: s1 = {a1, a2, a3} s2 = {a1, a2, a4, a5} s1.Difference(s2) = {a3} s2.Difference(s1) = {a4, a5}

// Has returns true if the StringSet contains the string
func (ss *StringSet) Has(item string) bool { _ = "STUB: not implemented"; return false }

type instrumentedIMDS struct {
	EC2MetadataIface
}

func awsReqStatus(err error) string { _ = "STUB: not implemented"; return "" }

// Unknown HTTP status code

func (i instrumentedIMDS) GetMetadataWithContext(ctx context.Context, p string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// New creates an EC2InstanceMetadataCache
func New(ctx context.Context, useSubnetDiscovery, useCustomNetworking, disableLeakedENICleanup, v4Enabled, v6Enabled bool) (*EC2InstanceMetadataCache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Clean up leaked ENIs in the background

func (cache *EC2InstanceMetadataCache) InitCachedPrefixDelegation(enablePrefixDelegation bool) {
	_ = "STUB: not implemented"
	return
}

// InitWithEC2metadata initializes the EC2InstanceMetadataCache with the data retrieved from EC2 metadata service
func (cache *EC2InstanceMetadataCache) initWithEC2Metadata(ctx context.Context) error {
	_ = "STUB: not implemented"

	// retrieve availability-zone
	return nil
}

// retrieve primary interface local-ipv4

// retrieve instance-id

// retrieve instance-type

// retrieve primary interface's mac

// retrieve subnet-id

// retrieve vpc-id

// We use the ctx here for testing, since we spawn go-routines above which will run forever.

// discoverCustomSecurityGroups discovers security groups with the cni-role tag
func (cache *EC2InstanceMetadataCache) discoverCustomSecurityGroups(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetENISubnetID gets the subnet ID for an ENI from AWS
func (cache *EC2InstanceMetadataCache) GetENISubnetID(ctx context.Context, eniID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Helper function to get ENIs that match specific criteria
func (cache *EC2InstanceMetadataCache) getFilteredENIs(store *datastore.DataStore, onlySecondarySubnets bool) []string {
	_ = "STUB: not implemented"
	return nil
}

// Skip primary ENI for secondary subnet operations

// Filter based on subnet type

// Apply standard filters (unmanaged and multi-card ENIs)

// Helper function to apply security groups to a list of ENIs
func (cache *EC2InstanceMetadataCache) applySecurityGroupsToENIs(ctx context.Context, eniIDs []string, sgIDs []string, logPrefix string) {
	_ = "STUB: not implemented"
	return
}

// Helper function to detect and log security group changes
func (cache *EC2InstanceMetadataCache) detectSecurityGroupChanges(newSGs []string, currentSGs *StringSet, sgType string) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// RefreshCustomSGIDs discovers and refreshes security groups tagged for use with the CNI
func (cache *EC2InstanceMetadataCache) RefreshCustomSGIDs(ctx context.Context, dsAccess *datastore.DataStoreAccess) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if no custom security groups were found (empty list)

// Clear custom security groups cache

// Apply primary security groups to ENIs in secondary subnets as fallback

// If there are changes, update ENIs in secondary subnets

// only secondary subnet ENIs

// applyPrimarySGsToSecondarySubnetENIs applies primary security groups to ENIs in secondary subnets across all datastores
func (cache *EC2InstanceMetadataCache) applyPrimarySGsToSecondarySubnetENIs(ctx context.Context, dsAccess *datastore.DataStoreAccess) {
	_ = "STUB: not implemented"
	return
}

// only secondary subnet ENIs

// RefreshSGIDs retrieves security groups
func (cache *EC2InstanceMetadataCache) RefreshSGIDs(ctx context.Context, mac string, dsAccess *datastore.DataStoreAccess) error {
	_ = "STUB: not implemented"
	return nil
}

// When subnet discovery is enabled, only apply primary SGs to primary subnet ENIs

// Get only primary subnet ENIs (onlySecondarySubnets=false)

// Filter out unmanaged ENIs

// Original behavior: apply to all managed ENIs when subnet discovery is disabled

// Apply security groups to the filtered ENIs

// GetAttachedENIs retrieves ENI information from meta data service
func (cache *EC2InstanceMetadataCache) GetAttachedENIs() (eniList []ENIMetadata, err error) {
	_ = "STUB: not implemented"
	return nil,

		// retrieve number of interfaces
		nil
}

// retrieve the attached ENIs

func (cache *EC2InstanceMetadataCache) getENIMetadata(eniMAC string) (ENIMetadata, error) {
	_ = "STUB: not implemented"
	return *new(ENIMetadata), nil
}

// Can this even happen? To be backwards compatible, we will always use 0 here and log an error.

// Get IMDS fields for the interface

// Efa-only interfaces do not have any ipv4s or ipv6s associated with it. If we don't find any local-ipv4 or ipv6 info in imds we assume it to be efa-only interface and validate this later via ec2 call

// Get IPv4 and IPv6 addresses assigned to interface

// For IPv6 ENIs, we have to return the error if Subnet is not discovered

// Handle the case where GetSubnetIPv6CIDRBlocks returns empty IPNet for IPv4-only subnets
// IMPORTANT: This scenario includes cross-VPC IPv4 ENIs attached to IPv6 nodes
// where the ENI subnet is IPv4-only but the node is configured for IPv6

// If IPv6 is enabled, get attached v6 prefixes.

// Get prefix on primary ENI when custom networking is enabled is not needed.
// If primary ENI has prefixes attached and then we move to custom networking, we don't need to fetch
// the prefix since recommendation is to terminate the nodes and that would have deleted the prefix on the
// primary ENI.

// awsGetFreeDeviceNumber calls EC2 API DescribeInstances to get the next free device index
func (cache *EC2InstanceMetadataCache) awsGetFreeDeviceNumber(ctx context.Context, networkCard int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// AllocENI creates an ENI and attaches it to the instance
// returns: newly created ENI ID
func (cache *EC2InstanceMetadataCache) AllocENI(ctx context.Context, sg []*string, eniCfgSubnet string, numIPs int, networkCard int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Also change the ENI's attribute so that the ENI will be deleted when the instance is deleted.

// attachENI calls EC2 API to attach the ENI and returns the attachment id
func (cache *EC2InstanceMetadataCache) attachENI(ctx context.Context, eniID string, networkCard int) (string, error) {
	_ = "STUB: not implemented"
	// attach to instance
	return "", nil
}

// createENITags creates all the tags required to be added to the ENI
//
// Returns:
// - []ec2types.TagSpecification: Returns the tags by converting it into AWS SDK class
func (cache *EC2InstanceMetadataCache) createENITags() []ec2types.TagSpecification {
	_ = "STUB: not implemented"
	return nil
}

func (cache *EC2InstanceMetadataCache) createENIInput(eniDescription string, tags []ec2types.TagSpecification, needIPs int) *ec2.CreateNetworkInterfaceInput {
	_ = "STUB: not implemented"
	return nil
}

// Even though IPv6 PD is enabled, we require a Primary IP for the ENI.
// This always creates an ENI which has 1 Primary IPv6 address
// We use assignIPv6Prefix to assign a prefix during setupENI

// setConnectionTrackingSettings applies connection tracking settings only if the primary ENI has it configured.
// Only non-nil values from the primary ENI configuration are stored.
func (cache *EC2InstanceMetadataCache) setConnectionTrackingSettings(config *ec2types.ConnectionTrackingConfiguration) {
	_ = "STUB: not implemented"
	return
}

// return ENI id, error
func (cache *EC2InstanceMetadataCache) createENI(ctx context.Context, sg []*string, eniCfgSubnet string, numIPs int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Even in fallback, check if primary subnet is excluded

// Primary subnet is explicitly excluded

// Check tag for all subnets including primary

// Log when primary subnet is excluded

// preset security groups for ENI with primary SGs

// If this is a secondary subnet and we have custom security groups, use those instead
// We already determined isPrimarySubnet above, just reuse the variable

// overring SGs if using secondary subnets and sgs

// Secondary subnet but no custom security groups available - use primary SGs as fallback

// If no valid subnets found, return appropriate error

// When subnet discovery is disabled, check if primary subnet is excluded

// If we can't determine exclusion status, log warning and proceed

// Primary subnet is explicitly excluded

func (cache *EC2InstanceMetadataCache) GetVpcSubnets(ctx context.Context) ([]ec2types.Subnet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sort the subnet by available IP address counter (desc order) before determining subnet to use

// isSubnetValidForENICreation checks if subnet should be used for ENI/IP allocation
// For primary subnet: include by default (no tag), exclude only if tag value is "0"
// For secondary subnets: exclude by default (no tag), include only if tag exists with non-"0" value
// If the subnet has cluster-specific tags, it will only be used by the matching cluster
func isSubnetValidForENICreation(subnet ec2types.Subnet, isPrimarySubnet bool) bool {
	_ = "STUB: not implemented"
	// Parse subnet tags
	return false
}

// Rule 1: CNI tag with value "0" always excludes the subnet

// Rule 2: Check CNI tag requirements based on subnet type

// Primary subnets are included by default (backwards compatibility)

// Secondary subnets require explicit opt-in via CNI tag

// Rule 3: Check cluster-specific tags

// Subnet has cluster tags but not for this cluster

// ValidSubnetForCluster checks if a subnet is valid for use by this cluster
// For secondary subnets, they must either have no cluster tags or have a matching cluster tag
func ValidSubnetTagsMatchingClusterName(subnet ec2types.Subnet) bool {
	_ = "STUB: not implemented"
	// Get cluster name for cluster-specific tag checks
	return false
}

// getTagValue returns the value of a specific tag key, or empty string if not found
func getTagValue(tags []ec2types.Tag, key string) string { _ = "STUB: not implemented"; return "" }

// checkClusterTags checks if subnet has cluster-specific tags and if it belongs to the current cluster
func checkClusterTags(tags []ec2types.Tag, localClusterTagKey string) (hasClusterTags bool, belongsToThisCluster bool) {
	_ = "STUB: not implemented"
	return false, false
}

func createENIUsingCustomCfg(sg []*string, eniCfgSubnet string, input *ec2.CreateNetworkInterfaceInput) *ec2.CreateNetworkInterfaceInput {
	_ = "STUB: not implemented"
	return nil
}

func (cache *EC2InstanceMetadataCache) tryCreateNetworkInterface(ctx context.Context, input *ec2.CreateNetworkInterfaceInput) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// buildENITags computes the desired AWS Tags for eni
func (cache *EC2InstanceMetadataCache) buildENITags() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// If clusterName is provided,
// tag the ENI with "cluster.k8s.amazonaws.com/name=<cluster_name>"

func (cache *EC2InstanceMetadataCache) TagENI(ctx context.Context, eniID string, currentTags map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func awsAPIErrInc(api string, err error) { _ = "STUB: not implemented"; return }

func awsUtilsErrInc(fn string, err error) { _ = "STUB: not implemented"; return }

// FreeENI detaches and deletes the ENI interface
func (cache *EC2InstanceMetadataCache) FreeENI(ctx context.Context, eniName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (cache *EC2InstanceMetadataCache) freeENI(ctx context.Context, eniName string, sleepDelayAfterDetach time.Duration, maxBackoffDelay time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Find out attachment

// Retry detaching the ENI from the instance

// It does take awhile for EC2 to detach ENI from instance, so we wait 2s before trying the delete.

// getENIAttachmentID calls EC2 to fetch the attachmentID of a given ENI
func (cache *EC2InstanceMetadataCache) getENIAttachmentID(ctx context.Context, eniID string) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Shouldn't happen, but let's be safe

// We cannot assume that the NetworkInterface.Attachment field is a non-nil
// pointer to a NetworkInterfaceAttachment struct.
// Ref: https://github.com/aws/amazon-vpc-cni-k8s/issues/914

func (cache *EC2InstanceMetadataCache) deleteENI(ctx context.Context, eniName string, maxBackoffDelay time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// If already deleted, we are good

// GetIPv4sFromEC2 calls EC2 and returns a list of all addresses on the ENI
func (cache *EC2InstanceMetadataCache) GetIPv4sFromEC2(ctx context.Context, eniID string) (addrList []ec2types.NetworkInterfacePrivateIpAddress, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Shouldn't happen, but let's be safe

// GetIPv4PrefixesFromEC2 calls EC2 and returns a list of all addresses on the ENI
func (cache *EC2InstanceMetadataCache) GetIPv4PrefixesFromEC2(ctx context.Context, eniID string) (addrList []ec2types.Ipv4PrefixSpecification, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Shouldn't happen, but let's be safe

// GetIPv6PrefixesFromEC2 calls EC2 and returns a list of all addresses on the ENI
func (cache *EC2InstanceMetadataCache) GetIPv6PrefixesFromEC2(ctx context.Context, eniID string) (addrList []ec2types.Ipv6PrefixSpecification, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DescribeAllENIs calls EC2 to refresh the ENIMetadata and tags for all attached ENIs
func (cache *EC2InstanceMetadataCache) DescribeAllENIs(ctx context.Context) (DescribeAllENIsResult, error) {
	_ = "STUB: not implemented"
	// Fetch all local ENI info from metadata
	return *new(DescribeAllENIsResult), nil
}

// If ENABLE_IMDS_ONLY_MODE is enabled, skip EC2 API call and return IMDS metadata only

// Collect the verified ENIs, adding multicards information from IMDS cache as well

// Return the result with empty tag map, trunk ENI and EFA ENIs as those cannot get from IMDS metadata

// Try calling EC2 to describe the interfaces.

// No error, exit the loop

// Remove this ENI from the map

// Remove the failing ENI ID from the EC2 API request and try again

// For other errors sleep a short while before the next retry

// Collect the verified ENIs

// Collect ENI response into ENI metadata and tags.

// Validate that Attachment is populated by EC2 response before logging

// Check if DeleteOnTermination is set for Primary ENI

// Set Connection Tracking settings from Primary ENI

// Network Card where EFA-only ENI is attached

// This assumes we only have one trunk attached to the node..

// Check IPv4 addresses

// convertTagsToSDKTags converts tags in stringMap format to AWS SDK format
func convertTagsToSDKTags(tagsMap map[string]string) []ec2types.Tag {
	_ = "STUB: not implemented"
	return nil
}

// convertSDKTagsToTags converts tags in AWS SDKs format to stringMap format
func convertSDKTagsToTags(sdkTags []ec2types.Tag) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// loadAdditionalENITags will load the additional ENI Tags from environment variables.
func loadAdditionalENITags() map[string]string { _ = "STUB: not implemented"; return nil }

// TODO: ideally we should fail in CNI init phase if the validation fails instead of warn.
// currently we only warn to be backwards-compatible and keep changes minimal in this version.

// If duplicate keys exist, the value of the key will be the value of latter key.

var eniErrorMessageRegex = regexp.MustCompile("'([a-zA-Z0-9-]+)'")

func badENIID(errMsg string) string { _ = "STUB: not implemented"; return "" }

// logOutOfSyncState compares the IP and metadata returned by IMDS and the EC2 API DescribeNetworkInterfaces calls
func logOutOfSyncState(eniID string, imdsIPv4s, ec2IPv4s []ec2types.NetworkInterfacePrivateIpAddress) {
	_ = "STUB: not implemented"
	// Comparing the IMDS IPv4 addresses attached to the ENI with the DescribeNetworkInterfaces AWS API call, which
	// technically should be the source of truth and contain the freshest information. Let's just do a quick scan here
	// and output some diagnostic messages if we find stale info in the IMDS result.
	return
}

// AllocIPAddress allocates an IP address for an ENI
func (cache *EC2InstanceMetadataCache) AllocIPAddress(ctx context.Context, eniID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (cache *EC2InstanceMetadataCache) FetchInstanceTypeLimits(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore any missing values

// Not checking for empty hypervisorType since have seen certain instances not getting this filled.

// GetENIIPv4Limit return IP address limit per ENI based on EC2 instance type
func (cache *EC2InstanceMetadataCache) GetENIIPv4Limit() int { _ = "STUB: not implemented"; return 0 }

// Subtract one from the IPv4Limit since we don't use the primary IP on each ENI for pods.

// GetENILimit returns the number of ENIs can be attached to an instance
func (cache *EC2InstanceMetadataCache) GetENILimit() int { _ = "STUB: not implemented"; return 0 }

// GetNetworkCards returns the network cards the instance has
func (cache *EC2InstanceMetadataCache) GetNetworkCards() []vpc.NetworkCard {
	_ = "STUB: not implemented"
	return nil
}

// fallback to default for network card index 0 as all instances have at least one network card
// this needs be changed when an instance can have multiple network cards each with different maxENI limits

// GetInstanceHypervisorFamily returns hypervisor of EC2 instance type
func (cache *EC2InstanceMetadataCache) GetInstanceHypervisorFamily() string {
	_ = "STUB: not implemented"
	return ""
}

// IsInstanceBareMetal derives bare metal value of the instance
func (cache *EC2InstanceMetadataCache) IsInstanceBareMetal() bool {
	_ = "STUB: not implemented"
	return false
}

// GetInstanceType return EC2 instance type
func (cache *EC2InstanceMetadataCache) GetInstanceType() string {
	_ = "STUB: not implemented"
	return ""

	// IsPrefixDelegationSupported return true if the instance type supports Prefix Assignment/Delegation
}

func (cache *EC2InstanceMetadataCache) IsPrefixDelegationSupported() bool {
	_ = "STUB: not implemented"
	return false
}

// IsTrunkingCompatible return true if the instance type supports ENI trunking or not exist in the list
func (cache *EC2InstanceMetadataCache) IsTrunkingCompatible() bool {
	_ = "STUB: not implemented"
	return false
}

// AllocIPAddresses allocates numIPs of IP address on an ENI
func (cache *EC2InstanceMetadataCache) AllocIPAddresses(ctx context.Context, eniID string, numIPs int) (*ec2.AssignPrivateIpAddressesOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If we don't need any more IPs, exit

func (cache *EC2InstanceMetadataCache) AllocIPv6Prefixes(ctx context.Context, eniID string) ([]*string, error) {
	_ = "STUB: not implemented"
	// We only need to allocate one IPv6 prefix per ENI.
	return nil, nil
}

// WaitForENIAndIPsAttached waits until the ENI has been attached and the secondary IPs have been added
func (cache *EC2InstanceMetadataCache) WaitForENIAndIPsAttached(eni string, wantedCidrs int) (eniMetadata ENIMetadata, err error) {
	_ = "STUB: not implemented"
	return *new(ENIMetadata), nil
}

func (cache *EC2InstanceMetadataCache) waitForENIAndIPsAttached(eni string, wantedCidrs int, maxBackoffDelay time.Duration) (eniMetadata ENIMetadata, err error) {
	_ = "STUB: not implemented"
	return *new(ENIMetadata), nil
}

// Wait until the ENI shows up in the instance metadata service and has at least some secondary IPs

// Verify that the ENI we are waiting for is in the returned list

// Check how many Secondary IPs or Prefixes have been attached

// We look for IPv6Address instead if prefix here

// Ignore primary IP of the ENI
// wantedCidrs will be at most 1 less then the IP limit for the ENI because of the primary IP in secondary pod

// At least some are attached

// If we have at least 1 Secondary IP, by now return what we have without an error

// We have some Secondary IPs, return the ones we have

// We have some prefixes, return the ones we have

// DeallocIPAddresses frees IP address on an ENI
func (cache *EC2InstanceMetadataCache) DeallocIPAddresses(ctx context.Context, eniID string, ips []string) error {
	_ = "STUB: not implemented"
	return nil
}

// DeallocPrefixAddresses frees Prefixes on an ENI (supports both IPv4 and IPv6)
func (cache *EC2InstanceMetadataCache) DeallocPrefixAddresses(ctx context.Context, eniID string, prefixes []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Separate IPv4 and IPv6 prefixes

// Parse the CIDR to determine if it's IPv4 or IPv6

// Handle IPv4 prefixes using UnassignPrivateIpAddresses API

// Handle IPv6 prefixes using UnassignIpv6Addresses API

func (cache *EC2InstanceMetadataCache) cleanUpLeakedENIs(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (cache *EC2InstanceMetadataCache) cleanUpLeakedENIsInternal(ctx context.Context, startupDelay time.Duration) {
	_ = "STUB: not implemented"
	return
}

// Clean up all the leaked ones we found

func (cache *EC2InstanceMetadataCache) tagENIcreateTS(ctx context.Context, eniID string, maxBackoffDelay time.Duration) {
	_ = "STUB: not implemented"
	// Tag the ENI with "node.k8s.amazonaws.com/createdAt=currentTime"
	return
}

// getLeakedENIs calls DescribeNetworkInterfaces to get all available ENIs that were allocated by
// the AWS CNI plugin, but were not deleted.
func (cache *EC2InstanceMetadataCache) getLeakedENIs(ctx context.Context) ([]ec2types.NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verify the description starts with "aws-K8S-"

// Check that it's not a newly created ENI

/* Set a time if we didn't find one. This is to prevent accidentally deleting ENIs that are in the
 * process of being attached by CNI versions v1.5.x or earlier.
 */

// GetVPCIPv4CIDRs returns VPC CIDRs
func (cache *EC2InstanceMetadataCache) GetVPCIPv4CIDRs() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: keep as net.IPNet and remove this round-trip to/from string

// GetLocalIPv4 returns the primary IP address on the primary interface
func (cache *EC2InstanceMetadataCache) GetLocalIPv4() net.IP {
	_ = "STUB: not implemented"
	return *

	// GetLocalIPv4 returns the primary IP address on the primary interface
	new(net.IP)
}

func (cache *EC2InstanceMetadataCache) GetLocalIPv6() net.IP {
	_ = "STUB: not implemented"
	return *new(net.IP)
}

// GetVPCIPv6CIDRs returns VPC CIDRs
func (cache *EC2InstanceMetadataCache) GetVPCIPv6CIDRs() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPrimaryENI returns the primary ENI
func (cache *EC2InstanceMetadataCache) GetPrimaryENI() string { _ = "STUB: not implemented"; return "" }

// GetPrimaryENImac returns the mac address of primary eni
func (cache *EC2InstanceMetadataCache) GetPrimaryENImac() string {
	_ = "STUB: not implemented"
	return ""
}

// SetUnmanagedENIs Set unmanaged ENI set
func (cache *EC2InstanceMetadataCache) SetUnmanagedENIs(eniIDs []string) {
	_ = "STUB: not implemented"
	return
}

// SetUnmanagedENIs Set unmanaged ENI set
func (cache *EC2InstanceMetadataCache) SetUnmanagedNetworkCards(skipNetworkCards []bool) {
	_ = "STUB: not implemented"
	return
}

// SetEfaOnlyENIsByNetworkCards
func (cache *EC2InstanceMetadataCache) SetEFAOnlyENIs(efaOnlyENIByNetworkCard []string) {
	_ = "STUB: not implemented"
	return
}

// GetInstanceID returns the instance ID
func (cache *EC2InstanceMetadataCache) GetInstanceID() string { _ = "STUB: not implemented"; return "" }

// IsUnmanagedENI returns if the eni is unmanaged
func (cache *EC2InstanceMetadataCache) IsUnmanagedENI(eniID string) bool {
	_ = "STUB: not implemented"
	return false
}

// IsUnmanagedENI returns if the eni is unmanaged
func (cache *EC2InstanceMetadataCache) IsUnmanagedNIC(networkCardIndex int) bool {
	_ = "STUB: not implemented"
	return false
}

// IsEfaOnlyENI the efaOnlyENI
func (cache *EC2InstanceMetadataCache) IsEfaOnlyENI(networkCardIndex int, eniID string) bool {
	_ = "STUB: not implemented"
	return false
}

func (cache *EC2InstanceMetadataCache) getENIsFromPaginatedDescribeNetworkInterfaces(input *ec2.DescribeNetworkInterfacesInput, filterFn func(networkInterface ec2types.NetworkInterface) error) error {
	_ = "STUB: not implemented"
	return nil
}

// IsPrimaryENI returns if the eni is unmanaged
func (cache *EC2InstanceMetadataCache) IsPrimaryENI(eniID string) bool {
	_ = "STUB: not implemented"
	return false
}

func checkAPIErrorAndBroadcastEvent(err error, api string) { _ = "STUB: not implemented"; return }

// IsSubnetExcluded checks if a subnet is excluded by examining its kubernetes.io/role/cni tag
func (cache *EC2InstanceMetadataCache) IsSubnetExcluded(ctx context.Context, subnetID string) (bool, error) {
	_ = "STUB: not implemented"
	// Get all VPC subnets with their tags
	return false, nil
}

// Find the specific subnet and check its tags

// cni=1, now check cluster tags

// Subnet not found in VPC
