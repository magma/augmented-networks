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
	"github.com/golang/glog"
	anpb "github.com/magma/augmented-networks/accounting/protos"
)

type anaServer struct {
	anpb.UnimplementedAccountingServer
}

func (_ *anaServer) Start(_ context.Context, in *anpb.Session) (*anpb.SessionResp, error) {
	glog.Infof("ANA Server::Start; request: %s", in.String())
	return &anpb.SessionResp{}, nil
}

func (_ *anaServer) Update(_ context.Context, in *anpb.UpdateReq) (*anpb.SessionResp, error) {
	glog.Infof("ANA Server::Update; request: %s", in.String())
	return &anpb.SessionResp{}, nil
}

func (_ *anaServer) Stop(ctx context.Context, in *anpb.UpdateReq) (*anpb.StopResp, error) {
	glog.Infof("ANA Server::Stop; request: %s", in.String())
	return &anpb.StopResp{}, nil
}
