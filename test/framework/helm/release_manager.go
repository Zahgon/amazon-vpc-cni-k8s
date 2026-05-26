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

package helm

import (
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
)

type ReleaseManager interface {
	InstallUnPackagedRelease(chart string, releaseName string, namespace string,
		values map[string]interface{}) (*release.Release, error)
	UninstallRelease(namespace string, releaseName string) (*release.UninstallReleaseResponse, error)
	InstallPackagedRelease(chart string, releaseName string, version string, namespace string,
		values map[string]interface{}) (*release.Release, error)
}

type defaultReleaseManager struct {
	kubeConfig string
}

func NewDefaultReleaseManager(kubeConfig string) ReleaseManager {
	_ = "STUB: not implemented"
	return *new(ReleaseManager)
}

func (d *defaultReleaseManager) InstallUnPackagedRelease(chart string, releaseName string, namespace string,
	values map[string]interface{}) (*release.Release, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultReleaseManager) InstallPackagedRelease(chart string, releaseName string, version string, namespace string,
	values map[string]interface{}) (*release.Release, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func installCharts(installAction *action.Install, chart string, values map[string]interface{}) (*release.Release, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultReleaseManager) UninstallRelease(namespace string, releaseName string) (*release.UninstallReleaseResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultReleaseManager) obtainActionConfig(namespace string) *action.Configuration {
	_ = "STUB: not implemented"
	return nil
}
