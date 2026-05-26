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

package manifest

import (
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type DaemonsetBuilder struct {
	namespace              string
	name                   string
	container              corev1.Container
	labels                 map[string]string
	nodeSelector           map[string]string
	terminationGracePeriod int
	hostNetwork            bool
	volume                 []corev1.Volume
	volumeMount            []corev1.VolumeMount
}

func NewDefaultDaemonsetBuilder() *DaemonsetBuilder { _ = "STUB: not implemented"; return nil }

func (d *DaemonsetBuilder) Labels(labels map[string]string) *DaemonsetBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DaemonsetBuilder) NodeSelector(labelKey string, labelVal string) *DaemonsetBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DaemonsetBuilder) Namespace(namespace string) *DaemonsetBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DaemonsetBuilder) TerminationGracePeriod(tg int) *DaemonsetBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DaemonsetBuilder) Name(name string) *DaemonsetBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DaemonsetBuilder) Container(container corev1.Container) *DaemonsetBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DaemonsetBuilder) PodLabel(labelKey string, labelValue string) *DaemonsetBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DaemonsetBuilder) HostNetwork(hostNetwork bool) *DaemonsetBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DaemonsetBuilder) MountVolume(volume []corev1.Volume, volumeMount []corev1.VolumeMount) *DaemonsetBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DaemonsetBuilder) Build() *v1.DaemonSet { _ = "STUB: not implemented"; return nil }
