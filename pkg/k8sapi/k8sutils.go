package k8sapi

import (
	"context"
	"time"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/utils/logger"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	cache "sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	awsNode         = "aws-node"
	envEnablePodENI = "ENABLE_POD_ENI"
	restCfgTimeout  = 5 * time.Second
)

var log = logger.Get()

// Get cache filters for IPAMD
func getIPAMDCacheFilters() map[client.Object]cache.ByObject { _ = "STUB: not implemented"; return nil }

// only cache CNINode when SGP is in use

// Get cache filters for CNI Metrics Helper
func getMetricsHelperCacheFilters() map[client.Object]cache.ByObject {
	_ = "STUB: not implemented"
	return nil
}

// Create cache reader for Kubernetes client
func CreateKubeClientCache(restCfg *rest.Config, scheme *runtime.Scheme, filterMap map[client.Object]cache.ByObject) (cache.Cache, error) {
	_ = "STUB: not implemented"
	// Get HTTP client and REST mapper for cache
	return *new(cache.Cache), nil
}

// Create a cache for the client to read from in order to decrease the number of API server calls.

func StartKubeClientCache(cache cache.Cache) { _ = "STUB: not implemented"; return }

// CreateKubeClient creates a k8s client
func CreateKubeClient(appName string) (client.Client, error) {
	_ = "STUB: not implemented"
	return *new(client.Client), nil
}

// The scheme should only contain GVKs that the client will access.

func GetKubeClientSet() (kubernetes.Interface, error) {
	_ = "STUB: not implemented"
	// creates the in-cluster config
	return *new(kubernetes.Interface), nil
}

// creates the clientset

func CheckAPIServerConnectivity() error { _ = "STUB: not implemented"; return nil }

// Reconcile the API server query after waiting for a second, as the request
// times out in one second if it fails to connect to the server

// When times out return no error, so the PollInfinite will retry with the given interval

func CheckAPIServerConnectivityWithTimeout(pollInterval time.Duration, pollTimeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// timeout for each connect try

// Retry

// GetRestConfig returns a Kubernetes REST config for API interactions
func GetRestConfig() (*rest.Config, error) { _ = "STUB: not implemented"; return nil, nil }

func GetNode(ctx context.Context, k8sClient client.Client) (corev1.Node, error) {
	_ = "STUB: not implemented"
	return *new(corev1.Node), nil
}

// If API server is unavailable, return immediately

// Create a context with timeout to avoid hanging indefinitely
// Set 3-second timeout
