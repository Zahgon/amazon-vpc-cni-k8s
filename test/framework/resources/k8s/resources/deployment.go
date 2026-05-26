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
	"context"
	"time"

	v1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DeploymentManager interface {
	CreateAndWaitTillDeploymentIsReady(deployment *v1.Deployment, timeout time.Duration) (*v1.Deployment, error)
	DeleteAndWaitTillDeploymentIsDeleted(deployment *v1.Deployment) error
	UpdateAndWaitTillDeploymentIsReady(deployment *v1.Deployment, timeout time.Duration) error
	GetDeployment(name, namespace string) (*v1.Deployment, error)
	WaitTillDeploymentReady(deployment *v1.Deployment, timeout time.Duration) (*v1.Deployment, error)

	WaitUntilDeploymentReady(ctx context.Context, dp *v1.Deployment) (*v1.Deployment, error)
	WaitUntilDeploymentDeleted(ctx context.Context, dp *v1.Deployment) error
}

type defaultDeploymentManager struct {
	k8sClient client.Client
}

// CreateAndWaitTillDeploymentIsReady creates and waits for deployment to become ready or timeout
// with error if deployment doesn't become ready.
func (d *defaultDeploymentManager) CreateAndWaitTillDeploymentIsReady(deployment *v1.Deployment, timeout time.Duration) (*v1.Deployment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Allow for the cache to sync

func (d *defaultDeploymentManager) DeleteAndWaitTillDeploymentIsDeleted(deployment *v1.Deployment) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultDeploymentManager) UpdateAndWaitTillDeploymentIsReady(deployment *v1.Deployment, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Retrieve the latest version of Deployment before attempting update

func (d *defaultDeploymentManager) GetDeployment(name, namespace string) (*v1.Deployment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultDeploymentManager) WaitTillDeploymentReady(deployment *v1.Deployment, timeout time.Duration) (*v1.Deployment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultDeploymentManager) WaitUntilDeploymentReady(ctx context.Context, dp *v1.Deployment) (*v1.Deployment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *defaultDeploymentManager) WaitUntilDeploymentDeleted(ctx context.Context, dp *v1.Deployment) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDefaultDeploymentManager(k8sClient client.Client) DeploymentManager {
	_ = "STUB: not implemented"
	return *new(DeploymentManager)
}
