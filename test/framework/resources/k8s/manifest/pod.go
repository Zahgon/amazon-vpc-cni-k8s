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
	v1 "k8s.io/api/core/v1"
)

type PodBuilder struct {
	name                   string
	namespace              string
	hostNetwork            bool
	container              v1.Container
	labels                 map[string]string
	terminationGracePeriod int
	nodeName               string
	restartPolicy          v1.RestartPolicy
	nodeSelector           map[string]string
	volume                 []v1.Volume
	volumeMount            []v1.VolumeMount
}

func NewDefaultPodBuilder() *PodBuilder { _ = "STUB: not implemented"; return nil }

func (p *PodBuilder) Name(name string) *PodBuilder { _ = "STUB: not implemented"; return nil }

func (p *PodBuilder) Namespace(namespace string) *PodBuilder { _ = "STUB: not implemented"; return nil }

func (p *PodBuilder) HostNetwork(hostNetwork bool) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (p *PodBuilder) Container(container v1.Container) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (p *PodBuilder) PodLabel(labelKey string, labelVal string) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (p *PodBuilder) NodeName(nodeName string) *PodBuilder { _ = "STUB: not implemented"; return nil }

func (p *PodBuilder) NodeSelector(nodeLabelKey string, nodeLabelVal string) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (p *PodBuilder) TerminationGracePeriod(period int) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (p *PodBuilder) RestartPolicy(policy v1.RestartPolicy) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (p *PodBuilder) MountVolume(volume []v1.Volume, volumeMount []v1.VolumeMount) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (p *PodBuilder) Build() *v1.Pod { _ = "STUB: not implemented"; return nil }
