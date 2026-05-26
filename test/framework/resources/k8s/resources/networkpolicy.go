package resources

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type NetworkPolicyManager interface {
	CreateNetworkPolicy(networkPolicy client.Object) error
	DeleteNetworkPolicy(networkPolicy client.Object) error
}

type defaultNetworkPolicyManager struct {
	networkPolicyClient client.Client
}

func NewNetworkPolicyManager(client client.Client) NetworkPolicyManager {
	_ = "STUB: not implemented"
	return *new(NetworkPolicyManager)
}

func (d *defaultNetworkPolicyManager) CreateNetworkPolicy(networkPolicy client.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultNetworkPolicyManager) DeleteNetworkPolicy(networkPolicy client.Object) error {
	_ = "STUB: not implemented"
	return nil
}
