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

// Package metrics provide common data structure and routines for converting/aggregating prometheus metrics to cloudwatch metrics
package metrics

import (
	"context"

	dto "github.com/prometheus/client_model/go"
	"k8s.io/client-go/kubernetes"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/publisher"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/utils/logger"
)

type metricMatcher func(metric *dto.Metric) bool
type actionFuncType func(aggregatedValue *float64, sampleValue float64)

type metricsTarget interface {
	grabMetricsFromTarget(ctx context.Context, target string) ([]byte, error)
	getInterestingMetrics() map[string]metricsConvert
	getCWMetricsPublisher() publisher.Publisher
	getTargetList(ctx context.Context) ([]string, error)
	submitCloudWatch() bool
	submitPrometheus() bool
	getLogger() logger.Logger
}

type metricsConvert struct {
	actions []metricsAction
}

type metricsAction struct {
	cwMetricName string
	matchFunc    metricMatcher
	actionFunc   actionFuncType
	data         *dataPoints
	bucket       *bucketPoints
	logToFile    bool
}

type dataPoints struct {
	lastSingleDataPoint float64
	curSingleDataPoint  float64
}

type bucketPoint struct {
	CumulativeCount *float64
	UpperBound      *float64
}

type bucketPoints struct {
	lastBucket []*bucketPoint
	curBucket  []*bucketPoint
}

func matchAny(metric *dto.Metric) bool { _ = "STUB: not implemented"; return false }

func metricsAdd(aggregatedValue *float64, sampleValue float64) { _ = "STUB: not implemented"; return }

func metricsMax(aggregatedValue *float64, sampleValue float64) { _ = "STUB: not implemented"; return }

func getMetricsFromPod(ctx context.Context, k8sClient kubernetes.Interface, podName string, namespace string, port int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processGauge(metric *dto.Metric, act *metricsAction) { _ = "STUB: not implemented"; return }

func processCounter(metric *dto.Metric, act *metricsAction) { _ = "STUB: not implemented"; return }

func processPercentile(metric *dto.Metric, act *metricsAction) { _ = "STUB: not implemented"; return }

func processHistogram(metric *dto.Metric, act *metricsAction, log logger.Logger) {
	_ = "STUB: not implemented"
	return
}

// found the matching bucket

func postProcessingCounter(convert metricsConvert, log logger.Logger) bool {
	_ = "STUB: not implemented"
	return false
}

// Only do delta if metric target did NOT restart

func postProcessingHistogram(convert metricsConvert, log logger.Logger) bool {
	_ = "STUB: not implemented"
	return false
}

// Delta against the previous bucket value
// e.g. diff between bucket LE250000 and previous bucket LE125000

// Delta against the previous value

// Only do delta if there is no restart for metric target

// Only do delta if there is no restart for metric target

func processMetric(family *dto.MetricFamily, convert metricsConvert, log logger.Logger) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// no addition work needs for GAUGE

// no addition work needs for PERCENTILE

func produceHistogram(act metricsAction, cw publisher.Publisher) { _ = "STUB: not implemented"; return }

func filterMetrics(originalMetrics map[string]*dto.MetricFamily,
	interestingMetrics map[string]metricsConvert,
) (map[string]*dto.MetricFamily, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func produceCloudWatchMetrics(t metricsTarget, families map[string]*dto.MetricFamily, convertDef map[string]metricsConvert, cw publisher.Publisher) error {
	_ = "STUB: not implemented"
	return nil
}

// Prometheus export supports only gauge metrics for now.

func producePrometheusMetrics(t metricsTarget, families map[string]*dto.MetricFamily, convertDef map[string]metricsConvert) error {
	_ = "STUB: not implemented"
	return nil
}

func resetMetrics(interestingMetrics map[string]metricsConvert) { _ = "STUB: not implemented"; return }

func metricsListGrabAggregateConvert(ctx context.Context, t metricsTarget) (map[string]*dto.MetricFamily, map[string]metricsConvert, bool, error) {
	_ = "STUB: not implemented"
	return nil, nil, false, nil
}

// it may take times to remove some metric targets

// TODO resetDetected is NOT right for cniMetrics, so force it for now

// Handler grabs metrics from target, aggregates the metrics and convert them into cloudwatch metrics
func Handler(ctx context.Context, t metricsTarget) { _ = "STUB: not implemented"; return }
