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

// package main - augmented network accounting client simulator tool
package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/magma/augmented-networks/accounting/protos"
)

var (
	// Certificate files.
	caCertFlag     = flag.String("ca_crt", "", "CA certificate file. Used to verify server TLS certificate.")
	clientCertFlag = flag.String("client_crt", "", "Client certificate file. Used for client certificate-based authentication.")
	clientKeyFlag  = flag.String("client_key", "", "Client private key file. Used for client certificate-based authentication.")

	// notls flag, use to disable TLS
	notlsFlag = flag.Bool("notls", false, `Disable TLS when set`)

	// insecure flag, use to disable server TLS cert verification
	insecureFlag = flag.Bool("insecure", false, `Disable TLS server certificate verification`)

	// gRPC connection timeouts
	grpcConnectTimeoutSec = flag.Int64("grpc_connect_tout_sec", 30, "Max gRPC timeout seconds")
	grpcRemoteTimeoutSec  = flag.Int64("remote_tout_sec", 20, "Max gRPC remote connect timeout seconds")
	grpcMaxDelaySec       = flag.Int64("max_delay_sec", 20, "Max gRPC retry delay seconds")

	concurentClients = flag.Int("extra_clients", 0, "Number of additional concurent clients")
	averageInUsage   = flag.Uint("average_in_usage", 1048576, "Average octets IN per report")
	averageOutUsage  = flag.Uint("average_out_usage", 102400, "Average octets Out per report")
	updateReportTime = flag.Uint("report_time", 10, "Max time between reports in seconds")

	updatesNumber = flag.Uint("updates", 10, "Number of concequitive updates to send")

	providerId = flag.String("provider_id", "provider_12345", "Provider Network ID")
	consumerId = flag.String("consumer_id", "consumer_67890", "Consumer Network ID")

	selfTest = flag.Bool("selftest", false, `Run internal test service and selftest`)

	address    string
	remoteTout time.Duration
)

func usage() {
	fmt.Printf(
		"\nUsage:\t%s [OPTIONS] <[address]:port>\n\t%s -selftest -v=1\n\nAvailable options:\n",
		os.Args[0], os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if f := flag.Lookup("logtostderr"); f != nil {
		f.Value.Set("true")
	}
	if *selfTest {
		glog.Info("Running self test")
		*notlsFlag = true
		address = startTestService()
	} else {
		if flag.NArg() != 1 || len(flag.Arg(0)) < 2 {
			flag.Usage()
			os.Exit(1)
		}
		address = flag.Arg(0)
	}
	remoteTout = time.Duration(*grpcRemoteTimeoutSec) * time.Second
	rand.Seed(time.Now().UnixNano())

	threads := *concurentClients
	if threads < 0 {
		threads = 0
	}
	done := make(chan struct{})
	for i := 0; i <= threads; i++ {
		go run(done)
	}
	for i := 0; i <= threads; i++ {
		<-done
	}
}

func run(done chan struct{}) {
	var (
		octetsIn, octetsOut uint64
		session             = &protos.Session{
			User:              &protos.Session_IMSI{IMSI: fmt.Sprintf("%.15d", rand.Int63())[:15]},
			ConsumerId:        *consumerId,
			SessionId:         "session_" + rndStr(8),
			ProviderId:        *providerId,
			ProviderApn:       rndStr(6),
			ProviderGatewayId: "magma_gw_" + rndStr(8),
		}
		startSeconds = time.Now().Unix()
	)

	conn, err := createConnection(address)
	if err != nil {
		glog.Fatalf("Failed to create gRPC connection to '%s': %v", address, err)
	}
	client := protos.NewAccountingClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), remoteTout)
	glog.V(1).Infof("  Sending Start: %+v", session)
	resp, err := client.Start(ctx, session)
	if err != nil {
		glog.Fatalf("Failed Start for session '%+v': %v", session, err)
	}
	cancel()
	glog.V(1).Infof("Start resp: %+v", resp)

	inDelta := *averageInUsage / 10
	inBase := *averageInUsage - inDelta
	inDelta *= 2
	outDelta := *averageOutUsage / 10
	outBase := *averageOutUsage - outDelta
	outDelta *= 2
	reportDuration := time.Duration(*updateReportTime) * time.Second
	for i := uint(0); i < *updatesNumber; i++ {
		time.Sleep(time.Duration(rand.Int63()) % reportDuration)
		octetsIn += uint64(inBase + uint(rand.Uint32())%inDelta)
		octetsOut += uint64(outBase + uint(rand.Uint32())%outDelta)
		ctx, cancel := context.WithTimeout(context.Background(), remoteTout)
		ureq := &protos.UpdateReq{
			Session:     session,
			OctetsIn:    octetsIn,
			OctetsOut:   octetsOut,
			SessionTime: uint32(time.Now().Unix() - startSeconds),
		}
		glog.V(1).Infof("  Sending Update: %+v", ureq)
		uresp, err := client.Update(ctx, ureq)
		if err != nil {
			glog.Fatalf("Failed Update for session %+v: %v", session, err)
		}
		cancel()
		glog.V(1).Infof("Update resp: %+v", uresp)
	}
	ctx, cancel = context.WithTimeout(context.Background(), remoteTout)
	ureq := &protos.UpdateReq{
		Session:     session,
		OctetsIn:    octetsIn + uint64(inDelta),
		OctetsOut:   octetsOut + uint64(outDelta),
		SessionTime: uint32(time.Now().Unix() - startSeconds),
	}
	glog.V(1).Infof("  Sending Stop: %+v", ureq)
	sresp, err := client.Stop(ctx, ureq)
	if err != nil {
		glog.Fatalf("Failed Stop for session '%+v': %v", session, err)
	}
	cancel()
	glog.V(1).Infof("Stop resp: %+v", sresp)
	if done != nil {
		done <- struct{}{}
	}
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

var lettersLen int64 = int64(len(letters))

func rndStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%lettersLen]
	}
	return string(b)
}
