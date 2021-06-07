//
// Copyright 2021 Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// package main - augmented network accounting server simulator
package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/golang/glog"
	anpb "github.com/magma/augmented-networks/accounting/protos"
	"google.golang.org/grpc"
)

var (
	// Certificate files.
	ca       = flag.String("ca", "", "CA certificate file.")
	cert     = flag.String("cert", "", "Certificate file.")
	key      = flag.String("key", "", "Private key file.")
	insecure = flag.Bool("insecure", false, "Skip TLS validation.")
	notls    = flag.Bool("notls", false, "Disable TLS validation. If true, no need to specify TLS related options.")

	address string
)

func usage() {
	fmt.Printf(
		"\nUsage:\t%s [OPTIONS] <[address]:port>\n\nAvailable options:\n",
		os.Args[0], os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if f := flag.Lookup("logtostderr"); f != nil {
		f.Value.Set("true")
	}
	if flag.NArg() != 1 || len(flag.Arg(0)) < 2 {
		fmt.Print("\nPlease, specify server [address]:port\n")
		flag.Usage()
		os.Exit(1)
	}
	address = flag.Arg(0)

	anscServer := &anaServer{}
	lis, err := net.Listen("tcp", address)
	if err != nil {
		glog.Exitf("Failed to listen on %s: %v", address, err)
	}
	grpcServer := grpc.NewServer(ServerCredentials()...)
	anpb.RegisterAccountingServer(grpcServer, anscServer)
	glog.Infof("Starting AccountingServer on %s", lis.Addr().String())

	grpcServer.Serve(lis)
}
