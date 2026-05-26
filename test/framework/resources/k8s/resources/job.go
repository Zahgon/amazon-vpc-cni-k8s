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

package resources

import (
	v1 "k8s.io/api/batch/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type JobManager interface {
	CreateAndWaitTillJobCompleted(job *v1.Job) (*v1.Job, error)
	DeleteAndWaitTillJobIsDeleted(job *v1.Job) error
}

type defaultJobManager struct {
	k8sClient client.Client
}

func NewDefaultJobManager(k8sClient client.Client) JobManager {
	_ = "STUB: not implemented"
	return *new(JobManager)
}

func (d *defaultJobManager) CreateAndWaitTillJobCompleted(job *v1.Job) (*v1.Job, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultJobManager) DeleteAndWaitTillJobIsDeleted(job *v1.Job) error {
	_ = "STUB: not implemented"
	return nil
}
