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
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

const ServerResponse = "message from the server!"

// Server that listens to client connection and first reads data and then writes data to the client.
func main() {
	// Supported modes are tcp, udp, http
	var serverMode string
	// Port on which Server will listen for incoming connections
	var serverPort string

	flag.StringVar(&serverMode, "server-mode", "tcp", "server mode, accepts tcp or udp")
	flag.StringVar(&serverPort, "server-port", "2273", "port on which you want to start the server")

	flag.Parse()

	addr := fmt.Sprintf(":%s", serverPort)

	if serverMode == "tcp" {
		StartTCPServer(addr)
	} else if serverMode == "udp" {
		StartUDPServer(addr)
	} else if serverMode == "http" {
		StartHTTPServer()
	} else {
		log.Fatal("invalid server mode, can accept tcp/udp/http only")
	}
}

func StartTCPServer(serverAddr string) { _ = "STUB: not implemented"; return }

func readAndWriteFromConnection(conn net.Conn) error { _ = "STUB: not implemented"; return nil }

func StartUDPServer(serverAddr string) { _ = "STUB: not implemented"; return }

func h(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func StartHTTPServer() { _ = "STUB: not implemented"; return }
