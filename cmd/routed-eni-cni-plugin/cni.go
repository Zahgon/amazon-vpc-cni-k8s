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

// AWS VPC CNI plugin binary
package main

import (
	"fmt"
	"net"
	"os"
	"runtime"

	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	cniSpecVersion "github.com/containernetworking/cni/pkg/version"

	"github.com/aws/amazon-vpc-cni-k8s/cmd/routed-eni-cni-plugin/driver"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/grpcwrapper"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/rpcwrapper"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/sgpp"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/typeswrapper"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/utils/logger"
	pb "github.com/aws/amazon-vpc-cni-k8s/rpc"
)

const (
	ipamdAddress            = "127.0.0.1:50051"
	dummyInterfacePrefix    = "dummy"
	npAgentConnTimeout      = 2
	npaSocketPath           = "/var/run/aws-node/npa.sock"
	multiNICPodAnnotation   = "k8s.amazonaws.com/nicConfig"
	multiNICAttachment      = "multi-nic-attachment"
	containerVethNamePrefix = "mNicIf"
)

var version string

// NetConf stores the common network config for the CNI plugin
type NetConf struct {
	types.NetConf

	// VethPrefix is the prefix to use when constructing the host-side
	// veth device name. It should be no more than four characters, and
	// defaults to 'eni'.
	VethPrefix string `json:"vethPrefix"`

	// MTU for eth0
	MTU string `json:"mtu"`

	// PodSGEnforcingMode is the enforcing mode for Security groups for pods feature
	PodSGEnforcingMode sgpp.EnforcingMode `json:"podSGEnforcingMode"`

	PluginLogFile string `json:"pluginLogFile"`

	PluginLogLevel string        `json:"pluginLogLevel"`
	RuntimeConfig  RuntimeConfig `json:"runtimeConfig"`
}

type RuntimeConfig struct {
	PodAnnotations map[string]string `json:"io.kubernetes.cri.pod-annotations"`
}

// K8sArgs is the valid CNI_ARGS used for Kubernetes
type K8sArgs struct {
	types.CommonArgs

	// K8S_POD_NAME is pod's name
	K8S_POD_NAME types.UnmarshallableString

	// K8S_POD_NAMESPACE is pod's namespace
	K8S_POD_NAMESPACE types.UnmarshallableString

	// K8S_POD_INFRA_CONTAINER_ID is pod's sandbox id
	K8S_POD_INFRA_CONTAINER_ID types.UnmarshallableString

	// K8S_POD_UID
	K8S_POD_UID types.UnmarshallableString
}

func init() {
	// This is to ensure that all the namespace operations are performed for
	// a single thread
	runtime.LockOSThread()
}

// LoadNetConf converts inputs (i.e. stdin) to NetConf
func LoadNetConf(bytes []byte) (*NetConf, logger.Logger, error) {
	_ = "STUB: not implemented"
	// Default config
	return nil, *new(logger.Logger), nil
}

func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

func add(args *skel.CmdArgs, cniTypes typeswrapper.CNITYPES, grpcClient grpcwrapper.GRPC,
	rpcClient rpcwrapper.RPC, driverClient driver.NetworkAPIs) error {
	_ = "STUB: not implemented"
	return nil
}

// Derive pod MTU. Note that the value has already been validated.

// Derive if Pod requires multiple NIC connected

// Set up a connection to the ipamD server.

// We will let the values in result struct guide us in terms of IP Address Family configured.

// AddNetwork guarantees that Gateway string is a valid IPNet

// CNI stores the route table ID in the MAC field of the interface. The Route table ID is on the host side
// and is used during cleanup to remove the ip rules when IPAMD is not reachable.

// This index always points to the container interface so that we get the IP address corresponding to the interface

// The dummy interface is purely virtual and is stored in the prevResult struct to assist in cleanup during the DEL command.

// Non-zero value means pods are using branch ENI

// SGP Pods will always get a single IP address, so it is safe to access the first element of vethMetadata

// For branch ENI mode, the pod VLAN ID is packed in Interface.Mac
// CNI currently uses dummyInterface to identify if a Pod is a SGP or a regular pod.
// To stop using dummy interface here, we have to let IPAMD store the SGP containers in a file, so the identification of such containers becomes easier during deletion flow and won't require API server call

// For non-branch ENI, the pod VLAN ID value of 0 is packed in Interface.Mac, while the interface device number is packed in Interface.Sandbox
// The use of this dummy interface can be removed for regular pods once we move to v1.20+ as CNI now stores the route table ID in the Mac field of the container interface.
// This is kept for supporting downgrade to v1.19 and below.

// return allocated IP back to IP pool

// dummy interface is appended to PrevResult for use during cleanup
// The interfaces field should only include host,container and dummy interfaces in the list.
// Revisit the prevResult cleanup logic if this changes

// Set up a connection to the network policy agent
// NP container might have been removed if network policies are not being used
// If NETWORK_POLICY_ENFORCING_MODE is not set, we will not configure anything related to NP

// Set timeout

//Make a GRPC call for network policy agent

// No need to cleanup IP and network, kubelet will send delete.

func cmdDel(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

func del(args *skel.CmdArgs, cniTypes typeswrapper.CNITYPES, grpcClient grpcwrapper.GRPC, rpcClient rpcwrapper.RPC,
	driverClient driver.NetworkAPIs) error {
	_ = "STUB: not implemented"
	return nil
}

// During del we cannot depend on the pod annotation for identifying a multi-NIC pod.
// Instead, we store the number of ips assigned to the pod in ipamd datastore so that it always returns what is assigned
// and for cases when IPAMD is down, we can use the previous result to identify the interfaces that were assigned to the pod.

// For pods using branch ENI, try to delete using previous result

// notify local IP address manager to free secondary IP
// Set up a connection to the server.

// When IPAMD is unreachable, try to teardown pod network using previous result. This action prevents rules from leaking while IPAMD is unreachable.
// Note that no error is returned to kubelet as there is no guarantee that kubelet will retry delete, and returning an error would prevent container runtime
// from cleaning up resources. When IPAMD again becomes responsive, it is responsible for reclaiming IP.

// Plugins should generally complete a DEL action without error even if some resources are missing. For example,
// an IPAM plugin should generally release an IP allocation and return success even if the container network
// namespace no longer exists, unless that network namespace is critical for IPAM management

// DelNetworkRequest may return a connection error, so try to delete using PrevResult whenever an error is returned. As with the case above, do
// not return error to kubelet, as there is no guarantee that delete is retried.

// vlanID != 0 means pod using security group

// Set up a connection to the network policy agent
// Set timeout

//Make a GRPC call for network policy agent

// NP agent will never return an error if its not able to delete ebpf probes

func getContainerNetworkMetadata(prevResult *current.Result, contVethName string) (net.IPNet, *current.Interface, error) {
	_ = "STUB: not implemented"
	return *new(net.IPNet), nil, nil
}

// tryDelWithPrevResult will try to process CNI delete request without IPAMD.
// returns true if the del request is handled.
func tryDelWithPrevResult(driverClient driver.NetworkAPIs, conf *NetConf, k8sArgs K8sArgs, contVethName string, netNS string, log logger.Logger) (bool, error) {
	_ = "STUB: not implemented"
	// prevResult might not be available, if we are still using older cni spec < 0.4.0.
	return false, nil
}

// For non-branch ENI pods, deletion requires handshake with IPAMD

// teardownPodNetworkWithPrevResult will try to process CNI delete for non-branch ENIs without IPAMD.
// Returns true if pod network is torn down
func teardownPodNetworkWithPrevResult(driverClient driver.NetworkAPIs, conf *NetConf, k8sArgs K8sArgs, contVethName string, log logger.Logger) bool {
	_ = "STUB: not implemented"
	// For non-branch ENI, prevResult is only available in v1.12.1+
	return false
}

// For non-branch ENI, VLAN ID of 0 is encoded in Mac and device number is encoded in Sandbox

// Since we are always using dummy interface to store the device number of network card 0 IP
// With v1.20+, route table id for NC-0 is stored in the container interface entry
// We have a path for migration where the dummy interface can be completely removed for regular network pods when we move completely to v1.20+
// This is currently done to support downgrade to v1.19 and below

// The number of interfaces attached to the pod is taken as the length of the interfaces array - 1 (for dummy interface) divided by 2 (for host and container interface)

// From v1.20+, this property of the container interface is used to store the route table ID
// Any pods that were created before v1.20+ will not have this property set. The pod will only have single interface as it only managed network card index 0

// Scope usage of this function to only SG pods scenario
// Don't process deletes when NetNS is empty
// as it implies that veth for this request is already deleted
// ref: https://github.com/kubernetes/kubernetes/issues/44100#issuecomment-329780382
func isNetnsEmpty(Netns string) bool { _ = "STUB: not implemented"; return false }

func main() {
	log := logger.DefaultLogger()
	about := fmt.Sprintf("AWS CNI %s", version)
	exitCode := 0
	if e := skel.PluginMainWithError(cmdAdd, nil, cmdDel, cniSpecVersion.All, about); e != nil {
		if err := e.Print(); err != nil {
			log.Errorf("Failed to write error to stdout: %v", err)
		}
		exitCode = 1
	}
	os.Exit(exitCode)
}

func getIPAddressFromIpAllocationMetadata(v *pb.IPAllocationMetadata) *net.IPNet {
	_ = "STUB: not implemented"
	return nil
}

func getDeviceNumberFromIpAllocationMetadata(v *pb.IPAllocationMetadata) int {
	_ = "STUB: not implemented"
	return 0
}

func getRouteTableIdFromIpAllocationMetadata(v *pb.IPAllocationMetadata) int {
	_ = "STUB: not implemented"
	return 0
}
