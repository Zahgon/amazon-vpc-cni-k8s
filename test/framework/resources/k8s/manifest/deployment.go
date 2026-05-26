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

type DeploymentBuilder struct {
	namespace              string
	name                   string
	replicas               int
	container              corev1.Container
	labels                 map[string]string
	annotations            map[string]string
	nodeSelector           map[string]string
	terminationGracePeriod int
	nodeName               string
	hostNetwork            bool
	volume                 []corev1.Volume
	volumeMount            []corev1.VolumeMount
}

func NewBusyBoxDeploymentBuilder(testImageRegistry string) *DeploymentBuilder {
	_ = "STUB: not implemented"
	return nil
}

func NewDefaultDeploymentBuilder() *DeploymentBuilder { _ = "STUB: not implemented"; return nil }

func NewCalicoStarDeploymentBuilder() *DeploymentBuilder { _ = "STUB: not implemented"; return nil }

func (d *DeploymentBuilder) Labels(labels map[string]string) *DeploymentBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentBuilder) NodeSelector(labelKey string, labelVal string) *DeploymentBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentBuilder) Namespace(namespace string) *DeploymentBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentBuilder) TerminationGracePeriod(tg int) *DeploymentBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentBuilder) Name(name string) *DeploymentBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentBuilder) NodeName(nodeName string) *DeploymentBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentBuilder) Replicas(replicas int) *DeploymentBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentBuilder) Container(container corev1.Container) *DeploymentBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentBuilder) PodLabel(labelKey string, labelValue string) *DeploymentBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentBuilder) PodAnnotation(key string, value string) *DeploymentBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentBuilder) HostNetwork(hostNetwork bool) *DeploymentBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentBuilder) MountVolume(volume []corev1.Volume, volumeMount []corev1.VolumeMount) *DeploymentBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentBuilder) Build() *v1.Deployment { _ = "STUB: not implemented"; return nil }
