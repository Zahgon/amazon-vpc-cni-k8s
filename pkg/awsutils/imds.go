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

package awsutils

import (
	"context"
	"net"

	"github.com/aws/smithy-go"

	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
)

// EC2MetadataIface is a subset of the EC2Metadata API.
type EC2MetadataIface interface {
	GetMetadata(ctx context.Context, params *imds.GetMetadataInput, optFns ...func(*imds.Options)) (*imds.GetMetadataOutput, error)
}

// TypedIMDS is a typed wrapper around raw untyped IMDS SDK API.
type TypedIMDS struct {
	EC2MetadataIface
}

// imdsRequestError to provide the caller on the request status
type imdsRequestError struct {
	requestKey string
	err        error
	code       string            // Added to support SDK V2 APIError interface
	fault      smithy.ErrorFault // Added to support SDK V2 APIError interface
}

var _ error = &imdsRequestError{}

func newIMDSRequestError(requestKey string, err error) *imdsRequestError {
	_ = "STUB: not implemented"
	return nil
}

// default code
// default fault

func (e *imdsRequestError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *imdsRequestError) Unwrap() error {
	_ = "STUB: not implemented"

	// Implement smithy.APIError interface
	return nil
}

func (e *imdsRequestError) ErrorCode() string {
	_ = "STUB: not implemented"
	// If wrapped error is an APIError, delegate to it
	return ""
}

func (e *imdsRequestError) ErrorMessage() string { _ = "STUB: not implemented"; return "" }

func (e *imdsRequestError) ErrorFault() smithy.ErrorFault {
	_ = "STUB: not implemented"
	// If wrapped error is an APIError, delegate to it
	return *new(smithy.ErrorFault)
}

func (e *imdsRequestError) HTTPStatusCode() int { _ = "STUB: not implemented"; return 0 }

func (e *imdsRequestError) RequestID() string { _ = "STUB: not implemented"; return "" }

func (typedimds TypedIMDS) getList(ctx context.Context, key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAZ returns the Availability Zone in which the instance launched.
func (typedimds TypedIMDS) GetAZ(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetInstanceType returns the type of this instance.
func (typedimds TypedIMDS) GetInstanceType(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetLocalIPv4 returns the private (primary) IPv4 address of the instance.
func (typedimds TypedIMDS) GetLocalIPv4(ctx context.Context) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

// GetInstanceID returns the ID of this instance.
func (typedimds TypedIMDS) GetInstanceID(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetMAC returns the first/primary network interface mac address.
func (typedimds TypedIMDS) GetMAC(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetMACs returns the interface addresses attached to the instance.
func (typedimds TypedIMDS) GetMACs(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Remove trailing /

// GetMACImdsFields returns the imds fields present for a MAC
func (typedimds TypedIMDS) GetMACImdsFields(ctx context.Context, mac string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Remove trailing /

// GetInterfaceID returns the ID of the network interface.
func (typedimds TypedIMDS) GetInterfaceID(ctx context.Context, mac string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (typedimds TypedIMDS) getInt(ctx context.Context, key string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetDeviceNumber returns the unique device number associated with an interface.  The primary interface is 0.
func (typedimds TypedIMDS) GetDeviceNumber(ctx context.Context, mac string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetSubnetID returns the ID of the subnet in which the interface resides.
func (typedimds TypedIMDS) GetSubnetID(ctx context.Context, mac string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Read the content first, even if there's an error

// Now handle any errors, but return subnetID if it was read

func (typedimds TypedIMDS) GetVpcID(ctx context.Context, mac string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Read the content first, even if there's an error

// Handle errors but preserve any partial vpcID data

// GetSecurityGroupIDs returns the IDs of the security groups to which the network interface belongs.
func (typedimds TypedIMDS) GetSecurityGroupIDs(ctx context.Context, mac string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (typedimds TypedIMDS) getIP(ctx context.Context, key string) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

func (typedimds TypedIMDS) getIPs(ctx context.Context, key string) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (typedimds TypedIMDS) getCIDR(ctx context.Context, key string) (*net.IPNet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No CIDR found. Not an error for IPv4-only subnets when requesting IPv6 CIDRs

// Why doesn't net.ParseCIDR just return values in this form?

func (typedimds TypedIMDS) getCIDRs(ctx context.Context, key string) ([]net.IPNet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Why doesn't net.ParseCIDR just return values in this form?

// GetLocalIPv4s returns the private IPv4 addresses associated with the interface.  First returned address is the primary address.
func (typedimds TypedIMDS) GetLocalIPv4s(ctx context.Context, mac string) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No IPv4 address on the interface, not an error

// GetLocalIPv4s returns the private IPv6 addresses associated with the interface.  First returned address is the primary address.
func (typedimds TypedIMDS) GetLocalIPv6s(ctx context.Context, mac string) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No IPv6 address on the interface, not an error

// GetIPv4Prefixes returns the IPv4 prefixes delegated to this interface
func (typedimds TypedIMDS) GetIPv4Prefixes(ctx context.Context, mac string) ([]net.IPNet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetIPv6Prefixes returns the IPv6 prefixes delegated to this interface
func (typedimds TypedIMDS) GetIPv6Prefixes(ctx context.Context, mac string) ([]net.IPNet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetLocalIPv6 returns the IPv6 addresses associated with the primary interface.
func (typedimds TypedIMDS) GetLocalIPv6(ctx context.Context) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

// No IPv6.  Not an error, just a disappointment :(

// GetIPv6s returns the IPv6 addresses associated with the interface.
func (typedimds TypedIMDS) GetIPv6s(ctx context.Context, mac string) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No IPv6.  Not an error, just a disappointment :(

// GetSubnetIPv4CIDRBlock returns the IPv4 CIDR block for the subnet in which the interface resides.
func (typedimds TypedIMDS) GetSubnetIPv4CIDRBlock(ctx context.Context, mac string) (*net.IPNet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetVPCIPv4CIDRBlocks returns the IPv4 CIDR blocks for the VPC.
func (typedimds TypedIMDS) GetVPCIPv4CIDRBlocks(ctx context.Context, mac string) ([]net.IPNet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetVPCIPv6CIDRBlocks returns the IPv6 CIDR blocks for the VPC.
func (typedimds TypedIMDS) GetVPCIPv6CIDRBlocks(ctx context.Context, mac string) ([]net.IPNet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No IPv6.  Not an error, just a disappointment :(

// GetSubnetIPv6CIDRBlocks returns the IPv6 CIDR block for the subnet in which the interface resides.
func (typedimds TypedIMDS) GetSubnetIPv6CIDRBlocks(ctx context.Context, mac string) (*net.IPNet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetNetworkCard returns the Network card the interface is attached on
func (typedimds TypedIMDS) GetNetworkCard(ctx context.Context, mac string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// If no network card field, it is connected to Network card 0

// IsNotFound returns true if the error was caused by an AWS API 404 response.
// We implement a Custom IMDS Error, so need to use APIError instead of HTTP Response Error
func IsNotFound(err error) bool { _ = "STUB: not implemented"; return false }

// Check for AWS ResponseError first

// Check if the error message contains status code 404

// Check for any APIError (including imdsRequestError)

// If it's our custom error with a wrapped ResponseError, check that

// Otherwise check if the error code indicates NotFound

// FakeIMDS is a trivial implementation of EC2MetadataIface using an in-memory map - for testing.
type FakeIMDS map[string]interface{}

func (f FakeIMDS) GetMetadata(ctx context.Context, params *imds.GetMetadataInput, optFns ...func(*imds.Options)) (*imds.GetMetadataOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Metadata API treats foo/ as foo

// Custom error type
type CustomRequestFailure struct {
	code       string
	message    string
	fault      smithy.ErrorFault
	statusCode int
	requestID  string
}

func (e *CustomRequestFailure) Error() string { _ = "STUB: not implemented"; return "" }

func (e *CustomRequestFailure) ErrorCode() string { _ = "STUB: not implemented"; return "" }

func (e *CustomRequestFailure) ErrorMessage() string { _ = "STUB: not implemented"; return "" }

func (e *CustomRequestFailure) ErrorFault() smithy.ErrorFault {
	_ = "STUB: not implemented"
	return *new(smithy.ErrorFault)
}

func (e *CustomRequestFailure) HTTPStatusCode() int { _ = "STUB: not implemented"; return 0 }

func (e *CustomRequestFailure) RequestID() string {
	_ = "STUB: not implemented"

	// GetMetadataWithContext implements the EC2MetadataIface interface.
	return ""
}

func (f FakeIMDS) GetMetadataWithContext(ctx context.Context, p string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Metadata API treats foo/ as foo
