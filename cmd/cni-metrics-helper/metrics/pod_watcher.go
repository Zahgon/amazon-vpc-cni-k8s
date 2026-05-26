package metrics

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/utils/logger"
)

type PodWatcher interface {
	GetCNIPods(ctx context.Context) ([]string, error)
}

type defaultPodWatcher struct {
	k8sClient client.Client
	log       logger.Logger
}

// NewDefaultPodWatcher creates a new podWatcher
func NewDefaultPodWatcher(k8sClient client.Client, log logger.Logger) *defaultPodWatcher {
	_ = "STUB: not implemented"
	return nil
}

// Returns aws-node pod info. Below function assumes CNI pods follow aws-node* naming format
// and so the function has to be updated if the CNI pod name format changes.
func (d *defaultPodWatcher) GetCNIPods(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
