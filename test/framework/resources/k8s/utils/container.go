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
	v1 "k8s.io/api/core/v1"
)

// AddOrUpdateEnvironmentVariable adds or updates existing Environment variable to the
// specified container name
func AddOrUpdateEnvironmentVariable(containers []v1.Container, containerName string,
	envVars map[string]string) error {
	_ = "STUB: not implemented"
	return nil

	// Update existing environment variable first
}

// Delete, so we don't add the environment variable multiple times

// Add the environment variables that were not already present

// RemoveEnvironmentVariables removes the environment variable from the specified container
func RemoveEnvironmentVariables(containers []v1.Container, containerName string,
	envVars map[string]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}
