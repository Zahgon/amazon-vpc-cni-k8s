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

package k8s

import (
	"github.com/aws/amazon-vpc-cni-k8s/test/framework/resources/k8s/resources"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ResourceManagers interface {
	JobManager() resources.JobManager
	DeploymentManager() resources.DeploymentManager
	CustomResourceManager() resources.CustomResourceManager
	NamespaceManager() resources.NamespaceManager
	ServiceManager() resources.ServiceManager
	NodeManager() resources.NodeManager
	PodManager() resources.PodManager
	DaemonSetManager() resources.DaemonSetManager
	ConfigMapManager() resources.ConfigMapManager
	NetworkPolicyManager() resources.NetworkPolicyManager
	EventManager() resources.EventManager
}

type defaultManager struct {
	jobManager            resources.JobManager
	deploymentManager     resources.DeploymentManager
	customResourceManager resources.CustomResourceManager
	namespaceManager      resources.NamespaceManager
	serviceManager        resources.ServiceManager
	nodeManager           resources.NodeManager
	podManager            resources.PodManager
	daemonSetManager      resources.DaemonSetManager
	configMapManager      resources.ConfigMapManager
	networkPolicyManager  resources.NetworkPolicyManager
	eventManager          resources.EventManager
}

func NewResourceManager(k8sClient client.Client, k8sClientset *kubernetes.Clientset, scheme *runtime.Scheme, config *rest.Config) ResourceManagers {
	_ = "STUB: not implemented"
	return *new(ResourceManagers)
}

func (m *defaultManager) JobManager() resources.JobManager {
	_ = "STUB: not implemented"
	return *new(resources.JobManager)
}

func (m *defaultManager) DeploymentManager() resources.DeploymentManager {
	_ = "STUB: not implemented"
	return *new(resources.DeploymentManager)
}

func (m *defaultManager) CustomResourceManager() resources.CustomResourceManager {
	_ = "STUB: not implemented"
	return *new(resources.CustomResourceManager)
}

func (m *defaultManager) NamespaceManager() resources.NamespaceManager {
	_ = "STUB: not implemented"
	return *new(resources.NamespaceManager)
}

func (m *defaultManager) ServiceManager() resources.ServiceManager {
	_ = "STUB: not implemented"
	return *new(resources.ServiceManager)
}

func (m *defaultManager) NodeManager() resources.NodeManager {
	_ = "STUB: not implemented"
	return *new(resources.NodeManager)
}

func (m *defaultManager) PodManager() resources.PodManager {
	_ = "STUB: not implemented"
	return *new(resources.PodManager)
}

func (m *defaultManager) DaemonSetManager() resources.DaemonSetManager {
	_ = "STUB: not implemented"
	return *new(resources.DaemonSetManager)
}

func (m *defaultManager) ConfigMapManager() resources.ConfigMapManager {
	_ = "STUB: not implemented"
	return *new(resources.ConfigMapManager)
}

func (m *defaultManager) NetworkPolicyManager() resources.NetworkPolicyManager {
	_ = "STUB: not implemented"
	return *new(resources.NetworkPolicyManager)
}

func (m defaultManager) EventManager() resources.EventManager {
	_ = "STUB: not implemented"
	return *new(resources.EventManager)
}
