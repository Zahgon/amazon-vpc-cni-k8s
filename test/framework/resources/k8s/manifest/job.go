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
	batchV1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
)

type JobBuilder struct {
	namespace              string
	name                   string
	parallelism            int
	container              v1.Container
	labels                 map[string]string
	terminationGracePeriod int
	nodeName               string
	hostNetwork            bool
	nodeSelector           map[string]string
}

func NewDefaultJobBuilder() *JobBuilder { _ = "STUB: not implemented"; return nil }

func (j *JobBuilder) Name(name string) *JobBuilder { _ = "STUB: not implemented"; return nil }

func (j *JobBuilder) NodeSelector(selectorKey string, selectorVal string) *JobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobBuilder) Namespace(namespace string) *JobBuilder { _ = "STUB: not implemented"; return nil }

func (j *JobBuilder) Container(container v1.Container) *JobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobBuilder) PodLabels(labelKey string, labelVal string) *JobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobBuilder) TerminationGracePeriod(terminationGracePeriod int) *JobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobBuilder) NodeName(nodeName string) *JobBuilder { _ = "STUB: not implemented"; return nil }

func (j *JobBuilder) Parallelism(parallelism int) *JobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobBuilder) HostNetwork(hostNetwork bool) *JobBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (j *JobBuilder) Build() *batchV1.Job { _ = "STUB: not implemented"; return nil }
