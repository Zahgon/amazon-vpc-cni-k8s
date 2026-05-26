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

type NamespaceManager interface {
	CreateNamespace(namespace string) error
	CreateNamespaceWithLabels(namespace string, labels map[string]string) error
	DeleteAndWaitTillNamespaceDeleted(namespace string) error

	WaitUntilNamespaceDeleted(ctx context.Context, ns *v1.Namespace) error
}

type defaultNamespaceManager struct {
	k8sClient client.Client
}

func NewDefaultNamespaceManager(k8sClient client.Client) NamespaceManager {
	_ = "STUB: not implemented"
	return *new(NamespaceManager)
}

func (m *defaultNamespaceManager) CreateNamespace(namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *defaultNamespaceManager) CreateNamespaceWithLabels(namespace string, labels map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *defaultNamespaceManager) DeleteAndWaitTillNamespaceDeleted(namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *defaultNamespaceManager) WaitUntilNamespaceDeleted(ctx context.Context, ns *v1.Namespace) error {
	_ = "STUB: not implemented"
	return nil
}
