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
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// loadFromFile loads a certificate key pair into a tls certificate and a CA certificate into a x509 certificate.
func loadFromFile() (*tls.Certificate, *x509.Certificate) {
	certificate, err := tls.LoadX509KeyPair(*cert, *key)
	if err != nil {
		glog.Exit("Could not load key/certificate pair from files:", err)
	}
	certificate.Leaf, err = x509.ParseCertificate(certificate.Certificate[0])
	if err != nil {
		glog.Exit("Could not parse x509 certificate from tls certificate:", err)
	}
	caFile, err := ioutil.ReadFile(*ca)
	if err != nil {
		glog.Errorf("could not read CA certificate: %s", err)
		return &certificate, nil
	}
	block, _ := pem.Decode(caFile)
	caCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		glog.Errorf("Error parsing CA certificate: %v", err)
	}
	return &certificate, caCert
}

// LoadCertificates loads certificates from files and exits if there's an error.
func LoadCertificates() ([]tls.Certificate, *x509.CertPool) {
	var (
		err      error
		certPool *x509.CertPool
	)
	certs, caBundle := loadFromFile()
	if certs == nil {
		glog.Exit("Please provide -cert & -key")
	}
	if *insecure {
		certPool, err = x509.SystemCertPool()
		if err != nil {
			certPool = x509.NewCertPool()
		}
	} else {
		if caBundle == nil {
			glog.Exit("Please provide -ca")
		}
		certPool = x509.NewCertPool()
	}
	if caBundle != nil {
		certPool.AddCert(caBundle)
	}
	return []tls.Certificate{*certs}, certPool
}

// ServerCredentials generates gRPC ServerOptions for existing credentials.
func ServerCredentials() []grpc.ServerOption {
	if *notls {
		return []grpc.ServerOption{}
	}
	certificates, certPool := LoadCertificates()
	if *insecure {
		return []grpc.ServerOption{grpc.Creds(credentials.NewTLS(&tls.Config{
			ClientAuth:   tls.VerifyClientCertIfGiven,
			Certificates: certificates,
			ClientCAs:    certPool,
		}))}
	}

	return []grpc.ServerOption{grpc.Creds(credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: certificates,
		ClientCAs:    certPool,
	}))}
}
