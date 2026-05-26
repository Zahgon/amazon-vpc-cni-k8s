// Package ec2wrapper is used to wrap around the ec2 service APIs
package ec2wrapper

import (
	"github.com/aws/amazon-vpc-cni-k8s/pkg/utils/logger"
	ec2metadata "github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

const (
	resourceID   = "resource-id"
	resourceKey  = "key"
	clusterIDTag = "CLUSTER_ID"
)

var log = logger.Get()

// EC2Wrapper is used to wrap around EC2 service APIs to obtain ClusterID from
// the ec2 instance tags
type EC2Wrapper struct {
	ec2ServiceClient         ec2.DescribeTagsAPIClient
	instanceIdentityDocument ec2metadata.InstanceIdentityDocument
}

// NewMetricsClient returns an instance of the EC2 wrapper
func NewMetricsClient() (*EC2Wrapper, error) { _ = "STUB: not implemented"; return nil, nil }

// GetClusterTag is used to retrieve a tag from the ec2 instance
func (e *EC2Wrapper) GetClusterTag(tagKey string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
