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

type Container struct {
	name            string
	image           string
	imagePullPolicy v1.PullPolicy
	command         []string
	args            []string
	probe           *v1.Probe
	ports           []v1.ContainerPort
	securityContext *v1.SecurityContext
	Env             []v1.EnvVar
}

func NewBusyBoxContainerBuilder(testImageRegistry string) *Container {
	_ = "STUB: not implemented"
	return nil
}

func NewCurlContainer(testImageRegistry string) *Container { _ = "STUB: not implemented"; return nil }

// See test/agent/README.md in this repository for more details
func NewTestHelperContainer(testImageRegistry string) *Container {
	_ = "STUB: not implemented"
	return nil
}

func NewNetCatAlpineContainer(testImageRegistry string) *Container {
	_ = "STUB: not implemented"
	return nil
}

// simple netcat OpenBSD version with alpine as the base image
// compatible with arm64 and amd64

func NewBaseContainer() *Container { _ = "STUB: not implemented"; return nil }

func (w *Container) CapabilitiesForSecurityContext(add []v1.Capability, drop []v1.Capability) *Container {
	_ = "STUB: not implemented"
	return nil
}

func (w *Container) Name(name string) *Container { _ = "STUB: not implemented"; return nil }

func (w *Container) Image(image string) *Container { _ = "STUB: not implemented"; return nil }

func (w *Container) ImagePullPolicy(policy v1.PullPolicy) *Container {
	_ = "STUB: not implemented"
	return nil
}

func (w *Container) Command(cmd []string) *Container { _ = "STUB: not implemented"; return nil }

func (w *Container) EnvVar(env []v1.EnvVar) *Container { _ = "STUB: not implemented"; return nil }

func (w *Container) Args(arg []string) *Container { _ = "STUB: not implemented"; return nil }

func (w *Container) LivenessProbe(probe *v1.Probe) *Container {
	_ = "STUB: not implemented"
	return nil
}

func (w *Container) Port(port v1.ContainerPort) *Container { _ = "STUB: not implemented"; return nil }

func (w *Container) Build() v1.Container { _ = "STUB: not implemented"; return *new(v1.Container) }
