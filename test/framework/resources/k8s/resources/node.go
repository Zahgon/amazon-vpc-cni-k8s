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

type NodeManager interface {
	GetNodes(nodeLabelKey string, nodeLabelVal string) (v1.NodeList, error)
	GetAllNodes() (v1.NodeList, error)
	UpdateNode(oldNode *v1.Node, newNode *v1.Node) error
	DeleteNode(node *v1.Node, opts ...client.DeleteOption) error
	WaitTillNodesReady(nodeLabelKey string, nodeLabelVal string, asgSize int) error
}

type defaultNodeManager struct {
	k8sClient client.Client
}

func NewDefaultNodeManager(k8sClient client.Client) NodeManager {
	_ = "STUB: not implemented"
	return *new(NodeManager)
}

func (d *defaultNodeManager) GetNodes(nodeLabelKey string, nodeLabelVal string) (v1.NodeList, error) {
	_ = "STUB: not implemented"
	return *new(v1.NodeList), nil
}

// Filtering control plane nodes from the list of nodes. kOps creates control plane nodes in the
// same subnet as worker nodes. Control plane nodes have the label `node-role.kubernetes.io/control-plane`
// defined, which can be used to filter out the control plane nodes

func (d *defaultNodeManager) GetAllNodes() (v1.NodeList, error) {
	_ = "STUB: not implemented"
	return *new(v1.NodeList), nil
}

func (d *defaultNodeManager) UpdateNode(oldNode *v1.Node, newNode *v1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultNodeManager) DeleteNode(node *v1.Node, opts ...client.DeleteOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultNodeManager) WaitTillNodesReady(nodeLabelKey string, nodeLabelVal string, asgSize int) error {
	_ = "STUB: not implemented"
	return nil
}
