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

package agent

import (
	"github.com/aws/amazon-vpc-cni-k8s/test/agent/pkg/input"
	"github.com/aws/amazon-vpc-cni-k8s/test/framework"
	"github.com/aws/amazon-vpc-cni-k8s/test/framework/resources/k8s/manifest"

	. "github.com/onsi/ginkgo/v2"
	appsV1 "k8s.io/api/apps/v1"
	batchV1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
)

// TrafficTest is used to execute a traffic test for TCP/UDP traffic
type TrafficTest struct {
	Framework *framework.Framework
	// The deployment builder is needed instead of deployment, in order
	// to inject the server container to the deployment
	TrafficServerDeploymentBuilder *manifest.DeploymentBuilder
	// Port on which the server should listen for traffic
	ServerPort int
	// TCP/UPD are the supported protocols
	ServerProtocol string
	// The number of client pods to create for testing connection to
	// each server Pod
	ClientCount int
	// Server count is the number of server pods to be created
	ServerCount int
	// Server Pod Label Key/Val is required in order to get the list of
	// pods belonging to the server deployment
	ServerPodLabelKey string
	ServerPodLabelVal string
	// Client Pod Label Key/Val is required in order to get the list of
	// pods belonging to the client job
	ClientPodLabelKey string
	ClientPodLabelVal string
	// If supplied the function will be used to validate the pods are as expected
	// For instance, to validate server/client pods are using Branch ENI
	ValidateServerPods func(list v1.PodList) error
	ValidateClientPods func(list v1.PodList) error
	// Boolean that indicates if IPv6 mode is enabled
	IsV6Enabled bool
}

// Tests traffic by creating multiple server pods using a deployment and multiple client pods
// using a Job. Each client Pod tests connectivity to each Server Pod.
func (t *TrafficTest) TestTraffic() (float64, error) {
	_ = "STUB: not implemented"
	// Server listens on a given TCP/UDP Port
	return 0, nil
}

// The Metric Server Aggregates all metrics from all the client Pod so we
// don't have to query each client Pod to get the metric

// Get the list of Server Pod in order to get the IP Address of the Servers

// Using this Validation Injector you can validate the Server Pods

// To the Client Job pass the list of Server IPs, so each client Pod tests connectivity to each
// server

// Get List of client Pods for validation

// Get the aggregated response from the metric server for calculating the connection success rate

// Clean up all the resources

func (t *TrafficTest) startTrafficServer() (*appsV1.Deployment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TrafficTest) startTrafficClient(serverAddList string, metricServerIP string) (*batchV1.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TrafficTest) startMetricServerPod() (*v1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TrafficTest) getTestStatusFromMetricServer(metricPodIP string) ([]input.TestStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TrafficTest) calculateSuccessRate(testStatuses []input.TestStatus) float64 {
	_ = "STUB: not implemented"
	return 0
}
