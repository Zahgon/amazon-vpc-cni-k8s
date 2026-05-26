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

// The aws-node initialization
package main

import (
	"os"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/procsyswrapper"
)

const (
	defaultHostCNIBinPath           = "/host/opt/cni/bin"
	metadataLocalIP                 = "local-ipv4"
	metadataMAC                     = "mac"
	defaultDisableIPv4TcpEarlyDemux = false
	defaultEnableIPv6               = false
	defaultEnableIPv6Egress         = false

	envDisableIPv4TcpEarlyDemux = "DISABLE_TCP_EARLY_DEMUX"
	envEnableIPv6               = "ENABLE_IPv6"
	envHostCniBinPath           = "HOST_CNI_BIN_PATH"
	envEgressV6                 = "ENABLE_V6_EGRESS"
)

func getNodePrimaryIF() (string, error) { _ = "STUB: not implemented"; return "", nil }

func configureSystemParams(procSys procsyswrapper.ProcSys, primaryIF string) error {
	_ = "STUB: not implemented"

	// Configure rp_filter in loose mode
	return nil
}

// Enable or disable TCP early demux based on environment variable
// Note that older kernels may not support tcp_early_demux, so we must first check that it exists.

func configureIPv6Settings(procSys procsyswrapper.ProcSys, primaryIF string) error {
	_ = "STUB: not implemented"

	// Enable IPv6 when environment variable is set
	// Note that IPv6 is not disabled when environment variable is unset. This is omitted to preserve default host semantics.
	return nil
}

// Check if IPv6 egress support is enabled in IPv4 cluster.

// Enable IPv6 forwarding on all interfaces by default

// For the primary ENI in IPv6, sysctls are set to:
// 1. forwarding=1
// 2. accept_ra=2
// 3. accept_redirects=1

func main() {
	os.Exit(_main())
}

func _main() int { _ = "STUB: not implemented"; return 0 }

// Copy all binaries from workdir to host bin dir except container init binary
