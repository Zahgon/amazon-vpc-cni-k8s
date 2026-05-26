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

// The aws-node ipam daemon binary
package main

import (
	"os"
	"time"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/ipamd"
)

const (
	appName = "aws-node"
	// metricsPort is the port for prometheus metrics
	metricsPort = 61678

	// Environment variable to disable the metrics endpoint on 61678
	envDisableMetrics = "DISABLE_METRICS"

	// Environment variable to disable the IPAMD introspection endpoint on 61679
	envDisableIntrospection = "DISABLE_INTROSPECTION"

	restCfgTimeout = 5 * time.Second
	pollInterval   = 5 * time.Second
	pollTimeout    = 30 * time.Second
)

func main() {
	os.Exit(_main())
}

// startBackgroundAPIServerCheck checks API connectivity in the background
func startBackgroundAPIServerCheck(ipamContext *ipamd.IPAMContext) {
	_ = "STUB: not implemented"
	return
}

// Create a new client for API server check

// Keep checking until connection is established

// Update IPAM context with new API server connectivity

// Exit the goroutine after successful connection

func _main() int {
	_ = "STUB: not implemented"
	// Start measuring full startup duration
	return 0
}

// Do not add anything before initializing logger

// Initialize controller-runtime logger

// Check API Server Connectivity

// Record failed startup

// Try a quick check first

// Create Kubernetes client for API server requests

// Create EventRecorder for use by IPAMD

// Measure node initialization duration

// Record failed IPAMD initialization and failed startup

// Record successful AWS initialization

// If not connected to API server yet, start background checks

// Pool manager

// Prometheus metrics

// CNI introspection endpoints

// Record successful startup duration before the blocking RPC handler call

// Start the RPC listener (this is a blocking call)
