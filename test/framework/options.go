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

package framework

var GlobalOptions Options

func init() {
	GlobalOptions.BindFlags()
}

type Options struct {
	KubeConfig         string
	ClusterName        string
	AWSRegion          string
	AWSVPCID           string
	NgNameLabelKey     string
	NgNameLabelVal     string
	EKSEndpoint        string
	CalicoVersion      string
	ContainerRuntime   string
	InstanceType       string
	InitialAddon       string
	TargetAddon        string
	InitialManifest    string
	TargetManifest     string
	InstallCalico      bool
	PublicSubnets      string
	PrivateSubnets     string
	AvailabilityZones  string
	PublicRouteTableID string
	NgK8SVersion       string
	TestImageRegistry  string
	PublishCWMetrics   bool
}

func (options *Options) BindFlags() { _ = "STUB: not implemented"; return }

func (options *Options) Validate() error { _ = "STUB: not implemented"; return nil }
