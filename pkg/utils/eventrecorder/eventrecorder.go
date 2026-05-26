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

// Package event recorder is used to raise events on aws-node pods
package eventrecorder

import (
	"os"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/sgpp"
	"github.com/aws/amazon-vpc-cni-k8s/pkg/utils/logger"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var log = logger.Get()
var MyNodeName = os.Getenv("MY_NODE_NAME")
var MyPodName = os.Getenv("MY_POD_NAME")

// Global variable for EventRecorder allows dependent packages to simply call Get
var eventRecorder *EventRecorder

const (
	EventReason = sgpp.VpcCNIEventReason
	appName     = "aws-node"
)

type EventRecorder struct {
	Recorder  events.EventRecorder
	K8sClient client.Client
	hostPod   corev1.Pod
}

func Init(k8sClient client.Client, withApiSever bool) error { _ = "STUB: not implemented"; return nil }

// EventRecorder is not considered critical, so no error is returned if host pod cannot be queried

func Get() *EventRecorder { _ = "STUB: not implemented"; return nil }

// SendPodEvent will raise event on aws-node with given type, reason, & message
func (e *EventRecorder) SendPodEvent(eventType, reason, action, message string) {
	_ = "STUB: not implemented"
	return
}

func findMyPod(k8sClient client.Client) (corev1.Pod, error) {
	_ = "STUB: not implemented"

	// Find my aws-node pod
	return *new(corev1.Pod), nil
}

// Functions used for mocking package
func InitMockEventRecorder() *events.FakeRecorder { _ = "STUB: not implemented"; return nil }
