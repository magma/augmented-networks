package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"time"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
)

// createConnection creates a new gRPC connection
func createConnection(addr string) (*grpc.ClientConn, error) {
	if len(addr) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty target address")
	}
	opts := getDialOptions(*clientCertFlag, *clientKeyFlag, *caCertFlag, *notlsFlag, *insecureFlag)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*grpcConnectTimeoutSec)*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "failed to dial %s: %v", addr, err)
	}
	return conn, err
}

var (
	keepaliveParams = keepalive.ClientParameters{
		Time:                59 * time.Second,
		Timeout:             20 * time.Second,
		PermitWithoutStream: true,
	}
)

func getDialOptions(clientCrt, clientCrtKey, rootCa string, notls, insecure bool) []grpc.DialOption {
	bckoff := backoff.DefaultConfig
	bckoff.MaxDelay = time.Second * time.Duration(*grpcMaxDelaySec)
	var opts = []grpc.DialOption{
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff:           bckoff,
			MinConnectTimeout: time.Duration(*grpcRemoteTimeoutSec) * time.Second,
		}),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepaliveParams),
	}
	if notls {
		return append(opts, grpc.WithInsecure())
	}
	tlsCfg := &tls.Config{}
	if insecure {
		tlsCfg.InsecureSkipVerify = true
	} else {
		// always try to add OS certs
		certPool, err := x509.SystemCertPool()
		if err != nil {
			glog.Warningf("OS Cert Pool initialization error: %v", err)
			certPool = x509.NewCertPool()
		}
		if len(rootCa) > 0 {
			// Add server RootCA
			if rootCa, err := ioutil.ReadFile(rootCa); err == nil {
				if !certPool.AppendCertsFromPEM(rootCa) {
					glog.Errorf("Failed to append certificates from %s", rootCa)
				}
			} else {
				glog.Errorf("Cannot load Root CA from '%s': %v", rootCa, err)
			}
		}
		if len(certPool.Subjects()) > 0 {
			tlsCfg.RootCAs = certPool
		} else {
			glog.Warning("Empty server certificate pool, using TLS InsecureSkipVerify")
			tlsCfg.InsecureSkipVerify = true
		}
		if len(clientCrt) > 0 {
			if len(clientCrtKey) > 0 {
				clientCert, err := tls.LoadX509KeyPair(clientCrt, clientCrtKey)
				if err == nil {
					tlsCfg.Certificates = []tls.Certificate{clientCert}
				} else {
					glog.Errorf("failed to load Client Certificate & Key from '%s', '%s': %v",
						clientCrt, clientCrtKey, err)
				}
			} else {
				glog.Error("empty certificate key location")
			}
		} else {
			glog.Errorf("empty certificate location")
		}
	}
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tlsCfg)))
	return opts
}
