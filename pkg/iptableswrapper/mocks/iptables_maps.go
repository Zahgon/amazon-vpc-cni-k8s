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

package mock_iptableswrapper

type MockIptables struct {
	// DataplaneState is a map from table name to chain name to slice of rulespecs
	DataplaneState map[string]map[string][][]string
}

func NewMockIptables() *MockIptables { _ = "STUB: not implemented"; return nil }

func (ipt *MockIptables) Exists(table, chainName string, rulespec ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ipt *MockIptables) Insert(table, chain string, pos int, rulespec ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ipt *MockIptables) Append(table, chain string, rulespec ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ipt *MockIptables) AppendUnique(table, chain string, rulespec ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ipt *MockIptables) Delete(table, chainName string, rulespec ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ipt *MockIptables) List(table, chain string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ipt *MockIptables) NewChain(table, chain string) error { _ = "STUB: not implemented"; return nil }

// Creating a new chain adds a -N chain rule to iptables

func (ipt *MockIptables) ClearChain(table, chain string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ipt *MockIptables) DeleteChain(table, chain string) error {
	_ = "STUB: not implemented"
	// More than just the create chain rule
	return nil
}

func (ipt *MockIptables) ListChains(table string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ipt *MockIptables) ChainExists(table, chain string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ipt *MockIptables) HasRandomFully() bool {
	_ = "STUB: not implemented"
	// TODO: Work out how to write a test case for this
	return false
}
