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

package manifest

import (
	v1 "k8s.io/api/core/v1"
)

type ServiceBuilder struct {
	name        string
	namespace   string
	port        int32
	nodePort    int32
	protocol    v1.Protocol
	selector    map[string]string
	annotation  map[string]string
	serviceType v1.ServiceType
}

func NewHTTPService() *ServiceBuilder { _ = "STUB: not implemented"; return nil }

func (s *ServiceBuilder) Name(name string) *ServiceBuilder { _ = "STUB: not implemented"; return nil }

func (s *ServiceBuilder) Namespace(namespace string) *ServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServiceBuilder) Port(port int32) *ServiceBuilder { _ = "STUB: not implemented"; return nil }

func (s *ServiceBuilder) NodePort(nodePort int32) *ServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServiceBuilder) Protocol(protocol v1.Protocol) *ServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServiceBuilder) Selector(labelKey string, labelVal string) *ServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServiceBuilder) ServiceType(serviceType v1.ServiceType) *ServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServiceBuilder) Annotations(annotations map[string]string) *ServiceBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServiceBuilder) Build() *v1.Service { _ = "STUB: not implemented"; return nil }
