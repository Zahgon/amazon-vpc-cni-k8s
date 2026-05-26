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

package main

const (
	// HostIfName is the interface name in host network namespace, it starts with veth
	HostIfName = "vethxxxx"
)

// SetupAddExpectV4 has all the mock EXPECT required when a container is added
func SetupAddExpectV4(ec egressContext, chain string, actualIptablesRules, actualRouteAdd, actualRouteDel *[]string) error {
	_ = "STUB: not implemented"
	return nil
}

// container route adding

// SetupDelExpectV4 has all the mock EXPECT required when a container is deleted
func SetupDelExpectV4(ec egressContext, actualIptablesDel *[]string) error {
	_ = "STUB: not implemented"
	return nil
}

// SetupAddExpectV6 has all the mock EXPECT required when a container is added
func SetupAddExpectV6(c egressContext, chain string, actualIptablesRules, actualRouteAdd, actualRouteReplace *[]string) error {
	_ = "STUB: not implemented"
	return nil
}

// container route adding

// SetupDelExpectV6 has all the mock EXPECT required when a container is deleted
func SetupDelExpectV6(c egressContext, chain string, actualIptablesDel *[]string) error {
	_ = "STUB: not implemented"
	return nil
}
