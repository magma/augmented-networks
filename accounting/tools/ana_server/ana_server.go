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
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/golang/glog"
	anpb "github.com/magma/augmented-networks/accounting/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/keepalive"
)

type IdMap map[string]string

var (
	// Certificate files.
	ca       = flag.String("ca", "", "CA certificate file.")
	cert     = flag.String("cert", "", "Certificate file.")
	key      = flag.String("key", "", "Private key file.")
	insecure = flag.Bool("insecure", false, "Skip TLS validation.")
	notls    = flag.Bool("notls", false, "Disable TLS validation. If true, no need to specify TLS related options.")

	proxyTo = flag.String("proxy_to", "", "Address of accounting service to forward requests to.")
	pcMap   = IdMap{}

	address     string
	proxyClient anpb.AccountingClient
)

func usage() {
	fmt.Printf(
		"\nUsage:\t%s [OPTIONS] <[address]:port>\n\nAvailable options:\n",
		os.Args[0], os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Var(
		&pcMap,
		"remap",
		"Replace consumer/provider ID with mapped ID; incoming_id:outgoing_id. Works only with valid proxy_to flag")
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

	if len(*proxyTo) > 2 {
		glog.Infof("Proxying requests to %s with ID remapping: %q", *proxyTo, pcMap)
		opts := []grpc.DialOption{
			grpc.WithConnectParams(grpc.ConnectParams{
				Backoff:           backoff.DefaultConfig,
				MinConnectTimeout: time.Second * 20,
			}),
			grpc.WithBlock(),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                30 * time.Second,
				Timeout:             20 * time.Second,
				PermitWithoutStream: true,
			}),
			grpc.WithInsecure(),
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		conn, err := grpc.DialContext(ctx, *proxyTo, opts...)
		if err != nil {
			glog.Fatalf("failed to dial %s: %v", *proxyTo, err)
		}
		cancel()
		proxyClient = anpb.NewAccountingClient(conn)
	}
	grpcServer.Serve(lis)
}

func (v *IdMap) Set(s string) error {
	if v != nil && len(s) > 0 {
		parts := strings.Split(s, ":")
		if len(parts) != 2 || len(parts[0]) == 0 || len(parts[1]) == 0 {
			return fmt.Errorf("invalid key:value specified: %s", s)
		}
		k, val := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		fmt.Printf("remapping '%s' to '%s'\n", k, val)
		(*v)[k] = val
	}
	return nil
}

func (v *IdMap) String() string {
	return ""
}
