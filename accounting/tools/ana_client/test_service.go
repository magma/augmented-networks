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
package main

import (
	"context"
	"net"

	"github.com/golang/glog"
	"github.com/magma/augmented-networks/accounting/protos"
	"google.golang.org/grpc"
)

func startTestService() string {
	lis, err := net.Listen("tcp", "localhost:")
	if err != nil {
		glog.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	protos.RegisterAccountingServer(grpcServer, testService{})
	go grpcServer.Serve(lis)
	return lis.Addr().String()
}

type testService struct{}

func (testService) Start(_ context.Context, in *protos.Session) (*protos.SessionResp, error) {
	glog.V(1).Infof("Received Start: %+v", in)
	return &protos.SessionResp{}, nil
}

func (testService) Update(_ context.Context, in *protos.UpdateReq) (*protos.SessionResp, error) {
	glog.V(1).Infof("Received Update: %+v", in)
	return &protos.SessionResp{}, nil
}

func (testService) Stop(ctx context.Context, in *protos.UpdateReq) (*protos.StopResp, error) {
	glog.V(1).Infof("Received Stop: %+v", in)
	return &protos.StopResp{}, nil
}
