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
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/aws/amazon-vpc-cni-k8s/rpc"
)

const (
	ipamdgRPCaddress      = "127.0.0.1:50051"
	grpcHealthServiceName = "grpc.health.v1.aws-node"

	vpccniPodIPKey = "vpc.amazonaws.com/pod-ips"

	defaultIpPerPodRequired = 1
)

// server controls RPC service responses.
type server struct {
	rpc.UnimplementedCNIBackendServer
	rpc.UnimplementedConfigServerBackendServer
	rpc.UnimplementedNPBackendServer
	version     string
	ipamContext *IPAMContext
}

// PodENIData is used to parse the list of ENIs in the branch ENI pod annotation
type PodENIData struct {
	ENIID        string `json:"eniId"`
	IfAddress    string `json:"ifAddress"`
	PrivateIP    string `json:"privateIp"`
	IPV6Addr     string `json:"ipv6Addr"`
	VlanID       int    `json:"vlanID"`
	SubnetCIDR   string `json:"subnetCidr"`
	SubnetV6CIDR string `json:"subnetV6Cidr"`
}

// AddNetwork processes CNI add network request and return an IP address for container
func (s *server) AddNetwork(ctx context.Context, in *rpc.AddNetworkRequest) (*rpc.AddNetworkReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Do this early, but after logging trace

// This will be a list of IPs. For now it's just one

// var interfacesCount int

// Check pod spec for Branch ENI

// Check that we have a trunk

// Parse JSON data

// Get pod IPv4 or IPv6 address based on mode

// For IPv6, the gateway is derived from the RA route on the primary ENI. The primary ENI is always in the same subnet as the trunk and branch ENI.
// For IPv4, the gateway is always the .1 address for the subnet CIDR.

// Not needed for branch ENI, they depend on trunkENIDeviceIndex

// This will always be device number + 1 as SGP won't run on multi-NIC

// continue to look through other datastores till you are unable to find an IP address when ONLY 1 ip is required
// if the last datastore also return ErrNoAvailableIPInDataStore, return an error

// We are only adding the pods primary IP to annotation

func (s *server) validateVersion(clientVersion string) error { _ = "STUB: not implemented"; return nil }

func (s *server) DelNetwork(ctx context.Context, in *rpc.DelNetworkRequest) (*rpc.DelNetworkReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Do this early, but after logging trace

// All datastores will store this count in its datastore

// ipsAllocated will always be same in all datastores for a Pod, so this will not change ever between datastores

// cidrStr will be pod IP i.e, IP/32 for v4 (or) IP/128 for v6.
// Case 1: PD is enabled but IP/32 key in AvailableIPv4Cidrs[cidrStr] exists, this means it is a secondary IP. Added IsPrefix check just for sanity.
// So this IP should be released immediately.
// Case 2: PD is disabled then IP/32 key in AvailableIPv4Cidrs[cidrStr] will not exists since key to AvailableIPv4Cidrs will be either /28 prefix or /32
// secondary IP. Hence now see if we need free up a prefix is no other pods are using it.

// Parse JSON data

// If ip address is still empty, continue till we reach the end of datastore
// This happens when pod is not in on network card 0, DS will return datastore.ErrUnknownPod

// On DEL, we pass IP being released

func (s *server) GetNetworkPolicyConfigs(ctx context.Context, e *emptypb.Empty) (*rpc.NetworkPolicyAgentConfigReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunRPCHandler handles request from gRPC
func (c *IPAMContext) RunRPCHandler(version string) error { _ = "STUB: not implemented"; return nil }

// If ipamd can talk to the API server and to the EC2 API, the pod is healthy.
// No need to ever change this to HealthCheckResponse_NOT_SERVING since it's a local service only

// Register reflection service on gRPC server.

// Add shutdown hook

// shutdownListener - Listen to signals and set ipamd to be in status "terminating"
func (c *IPAMContext) shutdownListener() { _ = "STUB: not implemented"; return }

// Interrupt signal sent from terminal

// Terminate signal sent from Kubernetes

// We received an interrupt signal, shut down.
