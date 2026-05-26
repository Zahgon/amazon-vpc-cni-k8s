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
	"github.com/aws/amazon-vpc-cni-k8s/test/framework"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/apps/v1"
)

func AddEnvVarToDaemonSetAndWaitTillUpdated(f *framework.Framework, dsName string, dsNamespace string,
	containerName string, envVars map[string]string) {
	_ = "STUB: not implemented"
	return
}

func RemoveVarFromDaemonSetAndWaitTillUpdated(f *framework.Framework, dsName string, dsNamespace string,
	containerName string, envVars map[string]struct{}) {
	_ = "STUB: not implemented"
	return
}

func UpdateEnvVarOnDaemonSetAndWaitUntilReady(f *framework.Framework, dsName string, dsNamespace string,
	containerName string, addOrUpdateEnv map[string]string, removeEnv map[string]struct{}) {
	_ = "STUB: not implemented"
	return
}

func updateDaemonsetEnvVarsAndWait(f *framework.Framework, dsName string, dsNamespace string,
	containerName string, addOrUpdateEnv map[string]string, removeEnv map[string]struct{}) {
	_ = "STUB: not implemented"
	return
}

// Check for init containers if the container is not found in list of containers

// Check for init containers if the container is not found in list of containers

// update multus daemonset if it exists
// to avoid being stuck in recursive loop, we need below check

func getDaemonSet(f *framework.Framework, dsName string, dsNamespace string) *v1.DaemonSet {
	_ = "STUB: not implemented"
	return nil
}

func waitTillDaemonSetUpdated(f *framework.Framework, oldDs *v1.DaemonSet, updatedDs *v1.DaemonSet) *v1.DaemonSet {
	_ = "STUB: not implemented"
	return nil
}
