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
	"time"

	v1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DaemonSetManager interface {
	GetDaemonSet(namespace string, name string) (*v1.DaemonSet, error)

	CreateAndWaitTillDaemonSetIsReady(daemonSet *v1.DaemonSet, timeout time.Duration) (*v1.DaemonSet, error)

	UpdateAndWaitTillDaemonSetReady(old *v1.DaemonSet, new *v1.DaemonSet) (*v1.DaemonSet, error)
	CheckIfDaemonSetIsReady(namespace string, name string) error
	DeleteAndWaitTillDaemonSetIsDeleted(daemonSet *v1.DaemonSet, timeout time.Duration) error
}

type defaultDaemonSetManager struct {
	k8sClient client.Client
}

func NewDefaultDaemonSetManager(k8sClient client.Client) DaemonSetManager {
	_ = "STUB: not implemented"
	return *new(DaemonSetManager)
}

func (d *defaultDaemonSetManager) CreateAndWaitTillDaemonSetIsReady(daemonSet *v1.DaemonSet, timeout time.Duration) (*v1.DaemonSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Allow for the cache to sync

func (d *defaultDaemonSetManager) GetDaemonSet(namespace string, name string) (*v1.DaemonSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultDaemonSetManager) UpdateAndWaitTillDaemonSetReady(old *v1.DaemonSet, new *v1.DaemonSet) (*v1.DaemonSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultDaemonSetManager) CheckIfDaemonSetIsReady(namespace string, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// Need to ensure the DesiredNumberScheduled is not 0 as it may happen if the DS is still being deleted from previous run

func (d *defaultDaemonSetManager) DeleteAndWaitTillDaemonSetIsDeleted(daemonSet *v1.DaemonSet, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
