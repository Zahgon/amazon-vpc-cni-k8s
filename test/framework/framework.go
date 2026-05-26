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

package framework

import (
	"github.com/go-logr/logr"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/aws/amazon-vpc-cni-k8s/test/framework/controller"
	"github.com/aws/amazon-vpc-cni-k8s/test/framework/resources/aws"
	"github.com/aws/amazon-vpc-cni-k8s/test/framework/resources/k8s"
)

type Framework struct {
	Options             Options
	K8sClient           client.Client
	CloudServices       aws.Cloud
	K8sResourceManagers k8s.ResourceManagers
	InstallationManager controller.InstallationManager
	Logger              logr.Logger
}

func New(options Options) *Framework { _ = "STUB: not implemented"; return nil }

// Create config for clients that need to access subresources

// For integration tests, the schema contains all Kubernetes resources for simplicity.

// For the IPAMD events test, the cache must be able to index on Event reasons.

// Start cache and wait for initial sync

// The cache will start a WATCH for all GVKs in the scheme. CNINode objects should not
// be cached, so their GVK is added only for the client.
