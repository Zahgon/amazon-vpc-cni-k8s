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

// Package iptableswrapper is a wrapper interface for the iptables package
package iptableswrapper

import "github.com/coreos/go-iptables/iptables"

// IPTablesIface is an interface created to make code unit testable.
// Both the iptables package version and mocked version implement the same interface
type IPTablesIface interface {
	Exists(table, chain string, rulespec ...string) (bool, error)
	Insert(table, chain string, pos int, rulespec ...string) error
	Append(table, chain string, rulespec ...string) error
	AppendUnique(table, chain string, rulespec ...string) error
	Delete(table, chain string, rulespec ...string) error
	List(table, chain string) ([]string, error)
	NewChain(table, chain string) error
	ClearChain(table, chain string) error
	DeleteChain(table, chain string) error
	ListChains(table string) ([]string, error)
	ChainExists(table, chain string) (bool, error)
	HasRandomFully() bool
}

// ipTables is a struct that implements IPTablesIface using iptables package.
type ipTables struct {
	ipt *iptables.IPTables
}

// NewIPTables return a ipTables struct that implements IPTablesIface
func NewIPTables(protocol iptables.Protocol) (IPTablesIface, error) {
	_ = "STUB: not implemented"
	return *new(IPTablesIface), nil
}

// Exists implements IPTablesIface interface by calling iptables package
func (i ipTables) Exists(table, chain string, rulespec ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Insert implements IPTablesIface interface by calling iptables package
func (i ipTables) Insert(table, chain string, pos int, rulespec ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// Append implements IPTablesIface interface by calling iptables package
func (i ipTables) Append(table, chain string, rulespec ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// AppendUnique implements IPTablesIface interface by calling iptables package
func (i ipTables) AppendUnique(table, chain string, rulespec ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete implements IPTablesIface interface by calling iptables package
func (i ipTables) Delete(table, chain string, rulespec ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// List implements IPTablesIface interface by calling iptables package
func (i ipTables) List(table, chain string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewChain implements IPTablesIface interface by calling iptables package
func (i ipTables) NewChain(table, chain string) error { _ = "STUB: not implemented"; return nil }

// ClearChain implements IPTablesIface interface by calling iptables package
func (i ipTables) ClearChain(table, chain string) error { _ = "STUB: not implemented"; return nil }

// DeleteChain implements IPTablesIface interface by calling iptables package
func (i ipTables) DeleteChain(table, chain string) error { _ = "STUB: not implemented"; return nil }

// ListChains implements IPTablesIface interface by calling iptables package
func (i ipTables) ListChains(table string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// ChainExists implements IPTablesIface interface by calling iptables package
}

func (i ipTables) ChainExists(table, chain string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// HasRandomFully implements IPTablesIface interface by calling iptables package
func (i ipTables) HasRandomFully() bool { _ = "STUB: not implemented"; return false }
