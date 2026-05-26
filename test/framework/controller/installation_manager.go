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

package controller

import (
	"github.com/aws/amazon-vpc-cni-k8s/test/framework/helm"
)

type InstallationManager interface {
	InstallCNIMetricsHelper(image string, tag string, clusterId string) error
	UnInstallCNIMetricsHelper() error
	InstallTigeraOperator(version string) error
	UninstallTigeraOperator() error
}

func NewDefaultInstallationManager(manager helm.ReleaseManager) InstallationManager {
	_ = "STUB: not implemented"
	return *new(InstallationManager)
}

type defaultInstallationManager struct {
	releaseManager helm.ReleaseManager
}

func (d *defaultInstallationManager) InstallCNIMetricsHelper(image string, tag string, clusterId string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultInstallationManager) UnInstallCNIMetricsHelper() error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultInstallationManager) InstallTigeraOperator(version string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultInstallationManager) UninstallTigeraOperator() error {
	_ = "STUB: not implemented"
	return nil
}
