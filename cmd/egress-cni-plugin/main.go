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

import (
	"fmt"
	"runtime"

	"github.com/containernetworking/cni/pkg/skel"
	cniversion "github.com/containernetworking/cni/pkg/version"
)

var version string

func init() {
	// this ensures that main runs only on main thread (thread group leader).
	// since namespace ops (unshare, setns) are done for a single thread, we
	// must ensure that the goroutine does not jump from OS thread to thread
	runtime.LockOSThread()
}

func main() {
	skel.PluginMain(cmdAdd, nil, cmdDel, cniversion.All, fmt.Sprintf("egress CNI plugin %s", version))
}

func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

func add(args *skel.CmdArgs, ec *egressContext) (err error) { _ = "STUB: not implemented"; return nil }

// uncomment following lines to debug inputs
//stdin := args.StdinData
//args.StdinData = nil
//ec.Log.Debugf("args: %+v, stdinData: %+v", *args, string(stdin))
//args.StdinData = stdin

// Convert MTU from string to int

// We will not be vending out this as a separate plugin by itself, and it is only intended to be used as a
// chained plugin to VPC CNI. We only need this plugin to kick in if egress is enabled in VPC CNI. So, the
// value of an env variable in VPC CNI determines whether this plugin should be enabled and this is an attempt to
// pass through the variable configured in VPC CNI.

// Invoke ipam del if err to avoid ip leak

// NodeIP is not IPv4 address, pod IPv6 egress for eks IPv4 cluster

// NodeIP is IPv4 address, pod IPv4 egress for eks IPv6 cluster

func cmdDel(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

func del(args *skel.CmdArgs, ec *egressContext) (err error) { _ = "STUB: not implemented"; return nil }

// We only need this plugin to kick in if egress is enabled

// NodeIP is not IPv4 address

// IPv6 egress

// IPv4 egress
