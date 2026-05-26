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

package awssession

import (
	"context"
	"time"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/utils/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
)

// Http client timeout env for sessions
const (
	httpTimeoutEnv = "HTTP_TIMEOUT"
	maxRetries     = 10

	// DefaultAWSSDKClientTimeout is the default timeout for individual HTTP requests made by AWS SDK clients.
	DefaultAWSSDKClientTimeout = 10 * time.Second
)

// NewAWSSDKHTTPClient returns a new HTTP client with the configured AWS SDK timeout.
// It returns *awshttp.BuildableClient (instead of *http.Client) so the SDK can
// inject custom CA bundles via WithTransportOptions in air-gapped regions.
func NewAWSSDKHTTPClient() *awshttp.BuildableClient { _ = "STUB: not implemented"; return nil }

var (
	log = logger.Get()
)

func getHTTPTimeout() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// New will return aws.Config to be used by Service Clients.
func New(ctx context.Context) (aws.Config, error) {
	_ = "STUB: not implemented"
	return *new(aws.Config), nil
}

// Fall back to default resolution
