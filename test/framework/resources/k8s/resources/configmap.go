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
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ConfigMapManager interface {
	GetConfigMap(namespace string, name string) (*v1.ConfigMap, error)
	UpdateConfigMap(oldConfigMap *v1.ConfigMap, newConfigMap *v1.ConfigMap) error
}

type defaultConfigMapManager struct {
	k8sClient client.Client
}

func (d defaultConfigMapManager) GetConfigMap(namespace string, name string) (*v1.ConfigMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d defaultConfigMapManager) UpdateConfigMap(oldConfigMap *v1.ConfigMap, newConfigMap *v1.ConfigMap) error {
	_ = "STUB: not implemented"
	return nil
}

func NewConfigMapManager(k8sClient client.Client) ConfigMapManager {
	_ = "STUB: not implemented"
	return *new(ConfigMapManager)
}
