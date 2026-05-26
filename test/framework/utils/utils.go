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

package utils

import (
	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// NamespacedName returns the namespaced name for k8s objects
func NamespacedName(obj v1.Object) types.NamespacedName {
	_ = "STUB: not implemented"
	return *new(types.NamespacedName)
}

func GetEnvValueForKeyFromDaemonSet(key string, ds *appsV1.DaemonSet) string {
	_ = "STUB: not implemented"
	return ""
}

func GetProjectRoot() string { _ = "STUB: not implemented"; return "" }

// in prow tests, the repository name is "vpc-cni"
