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

// Package publisher is used to batch and send metric data to CloudWatch
package publisher

import (
	"context"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	types "github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/ec2wrapper"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/utils/logger"
)

const (
	// cloudwatchMetricNamespace for custom metrics
	cloudwatchMetricNamespace = "Kubernetes"

	// Metric dimension constants
	clusterIDDimension = "CLUSTER_ID"

	// localMetricData is the default size for the local queue(slice)
	localMetricDataSize = 100

	// maxDataPoints is the maximum number of data points per PutMetricData API request
	maxDataPoints = 20

	// Default cluster id if unable to detect something more suitable
	defaultClusterID = "k8s-cluster"
)

var (
	// List of EC2 tags (in priority order) to use as the CLUSTER_ID metric dimension
	clusterIDTags = []string{
		"eks:cluster-name",
		"CLUSTER_ID",
		"Name",
	}
)

// cloudWatchAPI defines the interface with methods required from CloudWatch Service
type cloudWatchAPI interface {
	PutMetricData(ctx context.Context, params *cloudwatch.PutMetricDataInput, optFns ...func(*cloudwatch.Options)) (*cloudwatch.PutMetricDataOutput, error)
}

// Publisher defines the interface to publish one or more data points
type Publisher interface {
	// Publish publishes one or more metric data points
	Publish(metricsDataPoints ...types.MetricDatum)

	// Start is to initiate the batch and publish operation
	Start(publishInterval int)

	// Stop is to terminate the batch and publish operation
	Stop()
}

// cloudWatchPublisher implements the `Publisher` interface for batching and publishing
// metric data to the CloudWatch metrics backend
type cloudWatchPublisher struct {
	ctx                  context.Context
	cancel               context.CancelFunc
	updateIntervalTicker *time.Ticker
	clusterID            string
	cloudwatchClient     cloudWatchAPI
	localMetricData      []types.MetricDatum
	lock                 sync.RWMutex
	log                  logger.Logger
}

// Logic to fetch Region and CLUSTER_ID
// Case 1: Cx not using IRSA, we need to get region and clusterID using IMDS
// Case 2: Cx using IRSA but not specified clusterID, we can still get this info if IMDS is not blocked
// Case 3: Cx blocked IMDS access and not using IRSA (which means region == "") AND
// not specified clusterID then its a Cx error
// New returns a new instance of `Publisher`
func New(ctx context.Context, region string, clusterID string, log logger.Logger) (Publisher, error) {
	_ = "STUB: not implemented"
	return *new(Publisher), nil
}

// If Customers have explicitly specified clusterID then skip generating it

// Try to fetch region if not available

// Get ec2metadata client

// Build derived context

// Start is used to set up the monitor loop
func (p *cloudWatchPublisher) Start(publishInterval int) { _ = "STUB: not implemented"; return }

// Stop is used to cancel the monitor loop
func (p *cloudWatchPublisher) Stop() { _ = "STUB: not implemented"; return }

// Publish is a variadic function to publish one or more metric data points
func (p *cloudWatchPublisher) Publish(metricDataPoints ...types.MetricDatum) {
	_ = "STUB: not implemented"
	// Fetch dimensions for override
	return
}

// Grab lock

// NOTE: Iteration is used to override the metric dimensions

func (p *cloudWatchPublisher) pushLocal() { _ = "STUB: not implemented"; return }

func (p *cloudWatchPublisher) push(metricData []types.MetricDatum) {
	_ = "STUB: not implemented"
	return
}

// Setup input

// Publish data

// Mutate slice

// Reset Input

// Why is there a *cloudwatch.PutMetricDataInput and cloudwatch.PutMetricDataInput?
func (p *cloudWatchPublisher) send(input *cloudwatch.PutMetricDataInput) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudWatchPublisher) monitor(interval time.Duration) { _ = "STUB: not implemented"; return }

func (p *cloudWatchPublisher) getCloudWatchMetricNamespace() *string {
	_ = "STUB: not implemented"
	return nil
}

func getClusterID(ec2Client *ec2wrapper.EC2Wrapper, log logger.Logger) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *cloudWatchPublisher) getCloudWatchMetricDatumDimensions() []types.Dimension {
	_ = "STUB: not implemented"
	return nil
}

// min is a helper to compute the min of two integers
func min(x, y int) int { _ = "STUB: not implemented"; return 0 }
