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

	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ServiceManager interface {
	GetService(ctx context.Context, namespace string, name string) (*v1.Service, error)
	CreateService(ctx context.Context, service *v1.Service) (*v1.Service, error)
	DeleteAndWaitTillServiceDeleted(ctx context.Context, service *v1.Service) error
}

type defaultServiceManager struct {
	k8sClient client.Client
}

func NewDefaultServiceManager(k8sClient client.Client) ServiceManager {
	_ = "STUB: not implemented"
	return *new(ServiceManager)
}

func (s *defaultServiceManager) GetService(ctx context.Context, namespace string,
	name string) (*v1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *defaultServiceManager) CreateService(ctx context.Context, service *v1.Service) (*v1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Wait till the cache is refreshed

func (s *defaultServiceManager) DeleteAndWaitTillServiceDeleted(ctx context.Context, service *v1.Service) error {
	_ = "STUB: not implemented"
	return nil
}
