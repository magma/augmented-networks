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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: auth.proto

package protos

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RAT_Type int32

const (
	RAT_Type_WLAN           RAT_Type = 0
	RAT_Type_VIRTUAL        RAT_Type = 1
	RAT_Type_UTRAN          RAT_Type = 1000
	RAT_Type_GERAN          RAT_Type = 1001
	RAT_Type_GAN            RAT_Type = 1002
	RAT_Type_HSPA_EVOLUTION RAT_Type = 1003
	RAT_Type_EUTRAN         RAT_Type = 1004
	RAT_Type_CDMA2000_1X    RAT_Type = 2000
	RAT_Type_HRPD           RAT_Type = 2001
	RAT_Type_UMB            RAT_Type = 2002
	RAT_Type_EHRPD          RAT_Type = 2003
)

// Enum value maps for RAT_Type.
var (
	RAT_Type_name = map[int32]string{
		0:    "WLAN",
		1:    "VIRTUAL",
		1000: "UTRAN",
		1001: "GERAN",
		1002: "GAN",
		1003: "HSPA_EVOLUTION",
		1004: "EUTRAN",
		2000: "CDMA2000_1X",
		2001: "HRPD",
		2002: "UMB",
		2003: "EHRPD",
	}
	RAT_Type_value = map[string]int32{
		"WLAN":           0,
		"VIRTUAL":        1,
		"UTRAN":          1000,
		"GERAN":          1001,
		"GAN":            1002,
		"HSPA_EVOLUTION": 1003,
		"EUTRAN":         1004,
		"CDMA2000_1X":    2000,
		"HRPD":           2001,
		"UMB":            2002,
		"EHRPD":          2003,
	}
)

func (x RAT_Type) Enum() *RAT_Type {
	p := new(RAT_Type)
	*p = x
	return p
}

func (x RAT_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RAT_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_proto_enumTypes[0].Descriptor()
}

func (RAT_Type) Type() protoreflect.EnumType {
	return &file_auth_proto_enumTypes[0]
}

func (x RAT_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RAT_Type.Descriptor instead.
func (RAT_Type) EnumDescriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

// Authentication Information Request
type AuthenticationInformationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Subscriber identifier (IMSI)
	UserName string `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	// RAT Type - optional parameter
	RatType RAT_Type `protobuf:"varint,2,opt,name=RatType,proto3,enum=protos.RAT_Type" json:"RatType,omitempty"`
	// AuthenticationScheme UTF8Stringn: "EAP-AKA, EAP-AKA', etc." - optional parameter
	AuthenticationScheme string `protobuf:"bytes,3,opt,name=AuthenticationScheme,proto3" json:"AuthenticationScheme,omitempty"`
	// Number of EUTRAN vectors to request in response
	RequestedEutranVectors uint32 `protobuf:"varint,4,opt,name=RequestedEutranVectors,proto3" json:"RequestedEutranVectors,omitempty"`
	// Concatenation of RAND and AUTS in the case of a resync attach case (optional)
	EutranResyncInfo []byte `protobuf:"bytes,5,opt,name=EutranResyncInfo,proto3" json:"EutranResyncInfo,omitempty"`
	// Number of UTRAN/GERAN vectors to request in response (optional)
	RequestedUtranGeranVectors uint32 `protobuf:"varint,6,opt,name=RequestedUtranGeranVectors,proto3" json:"RequestedUtranGeranVectors,omitempty"`
	// UTRAN/GERAN Resync Info (optional)
	UtranGeranResyncInfo []byte `protobuf:"bytes,7,opt,name=UtranGeranResyncInfo,proto3" json:"UtranGeranResyncInfo,omitempty"`
}

func (x *AuthenticationInformationRequest) Reset() {
	*x = AuthenticationInformationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationInformationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationInformationRequest) ProtoMessage() {}

func (x *AuthenticationInformationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationInformationRequest.ProtoReflect.Descriptor instead.
func (*AuthenticationInformationRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticationInformationRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *AuthenticationInformationRequest) GetRatType() RAT_Type {
	if x != nil {
		return x.RatType
	}
	return RAT_Type_WLAN
}

func (x *AuthenticationInformationRequest) GetAuthenticationScheme() string {
	if x != nil {
		return x.AuthenticationScheme
	}
	return ""
}

func (x *AuthenticationInformationRequest) GetRequestedEutranVectors() uint32 {
	if x != nil {
		return x.RequestedEutranVectors
	}
	return 0
}

func (x *AuthenticationInformationRequest) GetEutranResyncInfo() []byte {
	if x != nil {
		return x.EutranResyncInfo
	}
	return nil
}

func (x *AuthenticationInformationRequest) GetRequestedUtranGeranVectors() uint32 {
	if x != nil {
		return x.RequestedUtranGeranVectors
	}
	return 0
}

func (x *AuthenticationInformationRequest) GetUtranGeranResyncInfo() []byte {
	if x != nil {
		return x.UtranGeranResyncInfo
	}
	return nil
}

// Authentication Information Answer (Section 7.2.6)
type AuthenticationInformationAnswer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Authentication vectors matching the requested number
	EutranVectors []*AuthenticationInformationAnswer_EUTRANVector `protobuf:"bytes,2,rep,name=EutranVectors,proto3" json:"EutranVectors,omitempty"`
	UtranVectors  []*AuthenticationInformationAnswer_UTRANVector  `protobuf:"bytes,3,rep,name=UtranVectors,proto3" json:"UtranVectors,omitempty"`
	GeranVectors  []*AuthenticationInformationAnswer_GERANVector  `protobuf:"bytes,4,rep,name=GeranVectors,proto3" json:"GeranVectors,omitempty"`
	Profile       *AuthenticationInformationAnswer_UserProfile    `protobuf:"bytes,5,opt,name=Profile,proto3" json:"Profile,omitempty"`
}

func (x *AuthenticationInformationAnswer) Reset() {
	*x = AuthenticationInformationAnswer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationInformationAnswer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationInformationAnswer) ProtoMessage() {}

func (x *AuthenticationInformationAnswer) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationInformationAnswer.ProtoReflect.Descriptor instead.
func (*AuthenticationInformationAnswer) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticationInformationAnswer) GetEutranVectors() []*AuthenticationInformationAnswer_EUTRANVector {
	if x != nil {
		return x.EutranVectors
	}
	return nil
}

func (x *AuthenticationInformationAnswer) GetUtranVectors() []*AuthenticationInformationAnswer_UTRANVector {
	if x != nil {
		return x.UtranVectors
	}
	return nil
}

func (x *AuthenticationInformationAnswer) GetGeranVectors() []*AuthenticationInformationAnswer_GERANVector {
	if x != nil {
		return x.GeranVectors
	}
	return nil
}

func (x *AuthenticationInformationAnswer) GetProfile() *AuthenticationInformationAnswer_UserProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

// 3GPP TS 29.272, 7.3.18 E-UTRAN-Vector
// For details about fields read 3GPP 33.401
type AuthenticationInformationAnswer_EUTRANVector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rand  []byte `protobuf:"bytes,1,opt,name=rand,proto3" json:"rand,omitempty"`
	Xres  []byte `protobuf:"bytes,2,opt,name=xres,proto3" json:"xres,omitempty"`
	Autn  []byte `protobuf:"bytes,3,opt,name=autn,proto3" json:"autn,omitempty"`
	Kasme []byte `protobuf:"bytes,4,opt,name=kasme,proto3" json:"kasme,omitempty"`
}

func (x *AuthenticationInformationAnswer_EUTRANVector) Reset() {
	*x = AuthenticationInformationAnswer_EUTRANVector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationInformationAnswer_EUTRANVector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationInformationAnswer_EUTRANVector) ProtoMessage() {}

func (x *AuthenticationInformationAnswer_EUTRANVector) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationInformationAnswer_EUTRANVector.ProtoReflect.Descriptor instead.
func (*AuthenticationInformationAnswer_EUTRANVector) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1, 0}
}

func (x *AuthenticationInformationAnswer_EUTRANVector) GetRand() []byte {
	if x != nil {
		return x.Rand
	}
	return nil
}

func (x *AuthenticationInformationAnswer_EUTRANVector) GetXres() []byte {
	if x != nil {
		return x.Xres
	}
	return nil
}

func (x *AuthenticationInformationAnswer_EUTRANVector) GetAutn() []byte {
	if x != nil {
		return x.Autn
	}
	return nil
}

func (x *AuthenticationInformationAnswer_EUTRANVector) GetKasme() []byte {
	if x != nil {
		return x.Kasme
	}
	return nil
}

// 3GPP TS 29.272, Section 7.3.19 UTRAN-Vector
type AuthenticationInformationAnswer_UTRANVector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rand               []byte `protobuf:"bytes,1,opt,name=rand,proto3" json:"rand,omitempty"`
	Xres               []byte `protobuf:"bytes,2,opt,name=xres,proto3" json:"xres,omitempty"`
	Autn               []byte `protobuf:"bytes,3,opt,name=autn,proto3" json:"autn,omitempty"`
	ConfidentialityKey []byte `protobuf:"bytes,4,opt,name=confidentiality_key,json=confidentialityKey,proto3" json:"confidentiality_key,omitempty"`
	IntegrityKey       []byte `protobuf:"bytes,5,opt,name=integrity_key,json=integrityKey,proto3" json:"integrity_key,omitempty"`
}

func (x *AuthenticationInformationAnswer_UTRANVector) Reset() {
	*x = AuthenticationInformationAnswer_UTRANVector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationInformationAnswer_UTRANVector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationInformationAnswer_UTRANVector) ProtoMessage() {}

func (x *AuthenticationInformationAnswer_UTRANVector) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationInformationAnswer_UTRANVector.ProtoReflect.Descriptor instead.
func (*AuthenticationInformationAnswer_UTRANVector) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1, 1}
}

func (x *AuthenticationInformationAnswer_UTRANVector) GetRand() []byte {
	if x != nil {
		return x.Rand
	}
	return nil
}

func (x *AuthenticationInformationAnswer_UTRANVector) GetXres() []byte {
	if x != nil {
		return x.Xres
	}
	return nil
}

func (x *AuthenticationInformationAnswer_UTRANVector) GetAutn() []byte {
	if x != nil {
		return x.Autn
	}
	return nil
}

func (x *AuthenticationInformationAnswer_UTRANVector) GetConfidentialityKey() []byte {
	if x != nil {
		return x.ConfidentialityKey
	}
	return nil
}

func (x *AuthenticationInformationAnswer_UTRANVector) GetIntegrityKey() []byte {
	if x != nil {
		return x.IntegrityKey
	}
	return nil
}

// 3GPP TS 29.272, 7.3.20 GERAN-Vector
type AuthenticationInformationAnswer_GERANVector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rand []byte `protobuf:"bytes,1,opt,name=rand,proto3" json:"rand,omitempty"`
	Sres []byte `protobuf:"bytes,2,opt,name=sres,proto3" json:"sres,omitempty"`
	Kc   []byte `protobuf:"bytes,3,opt,name=Kc,proto3" json:"Kc,omitempty"`
}

func (x *AuthenticationInformationAnswer_GERANVector) Reset() {
	*x = AuthenticationInformationAnswer_GERANVector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationInformationAnswer_GERANVector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationInformationAnswer_GERANVector) ProtoMessage() {}

func (x *AuthenticationInformationAnswer_GERANVector) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationInformationAnswer_GERANVector.ProtoReflect.Descriptor instead.
func (*AuthenticationInformationAnswer_GERANVector) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1, 2}
}

func (x *AuthenticationInformationAnswer_GERANVector) GetRand() []byte {
	if x != nil {
		return x.Rand
	}
	return nil
}

func (x *AuthenticationInformationAnswer_GERANVector) GetSres() []byte {
	if x != nil {
		return x.Sres
	}
	return nil
}

func (x *AuthenticationInformationAnswer_GERANVector) GetKc() []byte {
	if x != nil {
		return x.Kc
	}
	return nil
}

// Optional User Profile (
type AuthenticationInformationAnswer_UserProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// MSISDN
	Msisdn string `protobuf:"bytes,1,opt,name=Msisdn,proto3" json:"Msisdn,omitempty"`
}

func (x *AuthenticationInformationAnswer_UserProfile) Reset() {
	*x = AuthenticationInformationAnswer_UserProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationInformationAnswer_UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationInformationAnswer_UserProfile) ProtoMessage() {}

func (x *AuthenticationInformationAnswer_UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationInformationAnswer_UserProfile.ProtoReflect.Descriptor instead.
func (*AuthenticationInformationAnswer_UserProfile) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1, 3}
}

func (x *AuthenticationInformationAnswer_UserProfile) GetMsisdn() string {
	if x != nil {
		return x.Msisdn
	}
	return ""
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x22, 0xf6, 0x02, 0x0a, 0x20, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x52, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x52, 0x41, 0x54, 0x5f, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x52, 0x61, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x32, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x45, 0x75, 0x74, 0x72, 0x61, 0x6e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x45, 0x75, 0x74, 0x72, 0x61, 0x6e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x2a, 0x0a,
	0x10, 0x45, 0x75, 0x74, 0x72, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x45, 0x75, 0x74, 0x72, 0x61, 0x6e, 0x52,
	0x65, 0x73, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3e, 0x0a, 0x1a, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x55, 0x74, 0x72, 0x61, 0x6e, 0x47, 0x65, 0x72, 0x61, 0x6e,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1a, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x55, 0x74, 0x72, 0x61, 0x6e, 0x47, 0x65, 0x72,
	0x61, 0x6e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x32, 0x0a, 0x14, 0x55, 0x74, 0x72,
	0x61, 0x6e, 0x47, 0x65, 0x72, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x14, 0x55, 0x74, 0x72, 0x61, 0x6e, 0x47, 0x65,
	0x72, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xf0, 0x05,
	0x0a, 0x1f, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x12, 0x5a, 0x0a, 0x0d, 0x45, 0x75, 0x74, 0x72, 0x61, 0x6e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x2e, 0x45, 0x55, 0x54, 0x52, 0x41, 0x4e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0d,
	0x45, 0x75, 0x74, 0x72, 0x61, 0x6e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x57, 0x0a,
	0x0c, 0x55, 0x74, 0x72, 0x61, 0x6e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x55, 0x54, 0x52,
	0x41, 0x4e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0c, 0x55, 0x74, 0x72, 0x61, 0x6e, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x57, 0x0a, 0x0c, 0x47, 0x65, 0x72, 0x61, 0x6e, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x47, 0x45, 0x52, 0x41, 0x4e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x0c, 0x47, 0x65, 0x72, 0x61, 0x6e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12,
	0x4d, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x60,
	0x0a, 0x0c, 0x45, 0x55, 0x54, 0x52, 0x41, 0x4e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x72, 0x61,
	0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x78, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x78, 0x72, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x61, 0x75, 0x74, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6b, 0x61,
	0x73, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6b, 0x61, 0x73, 0x6d, 0x65,
	0x1a, 0x9f, 0x01, 0x0a, 0x0b, 0x55, 0x54, 0x52, 0x41, 0x4e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x72, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x78, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x78, 0x72, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x61, 0x75, 0x74, 0x6e, 0x12, 0x2f, 0x0a, 0x13,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a,
	0x0d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x69, 0x74, 0x79, 0x4b,
	0x65, 0x79, 0x1a, 0x45, 0x0a, 0x0b, 0x47, 0x45, 0x52, 0x41, 0x4e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x72, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x4b, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x4b, 0x63, 0x1a, 0x25, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x73, 0x69, 0x73,
	0x64, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x73, 0x69, 0x73, 0x64, 0x6e,
	0x2a, 0x98, 0x01, 0x0a, 0x08, 0x52, 0x41, 0x54, 0x5f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a,
	0x04, 0x57, 0x4c, 0x41, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x56, 0x49, 0x52, 0x54, 0x55,
	0x41, 0x4c, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x05, 0x55, 0x54, 0x52, 0x41, 0x4e, 0x10, 0xe8, 0x07,
	0x12, 0x0a, 0x0a, 0x05, 0x47, 0x45, 0x52, 0x41, 0x4e, 0x10, 0xe9, 0x07, 0x12, 0x08, 0x0a, 0x03,
	0x47, 0x41, 0x4e, 0x10, 0xea, 0x07, 0x12, 0x13, 0x0a, 0x0e, 0x48, 0x53, 0x50, 0x41, 0x5f, 0x45,
	0x56, 0x4f, 0x4c, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0xeb, 0x07, 0x12, 0x0b, 0x0a, 0x06, 0x45,
	0x55, 0x54, 0x52, 0x41, 0x4e, 0x10, 0xec, 0x07, 0x12, 0x10, 0x0a, 0x0b, 0x43, 0x44, 0x4d, 0x41,
	0x32, 0x30, 0x30, 0x30, 0x5f, 0x31, 0x58, 0x10, 0xd0, 0x0f, 0x12, 0x09, 0x0a, 0x04, 0x48, 0x52,
	0x50, 0x44, 0x10, 0xd1, 0x0f, 0x12, 0x08, 0x0a, 0x03, 0x55, 0x4d, 0x42, 0x10, 0xd2, 0x0f, 0x12,
	0x0a, 0x0a, 0x05, 0x45, 0x48, 0x52, 0x50, 0x44, 0x10, 0xd3, 0x0f, 0x32, 0x71, 0x0a, 0x10, 0x53,
	0x49, 0x4d, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x5d, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x42, 0x3a,
	0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x67,
	0x6d, 0x61, 0x2f, 0x61, 0x75, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x2d, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_auth_proto_goTypes = []interface{}{
	(RAT_Type)(0),                                        // 0: protos.RAT_Type
	(*AuthenticationInformationRequest)(nil),             // 1: protos.AuthenticationInformationRequest
	(*AuthenticationInformationAnswer)(nil),              // 2: protos.AuthenticationInformationAnswer
	(*AuthenticationInformationAnswer_EUTRANVector)(nil), // 3: protos.AuthenticationInformationAnswer.EUTRANVector
	(*AuthenticationInformationAnswer_UTRANVector)(nil),  // 4: protos.AuthenticationInformationAnswer.UTRANVector
	(*AuthenticationInformationAnswer_GERANVector)(nil),  // 5: protos.AuthenticationInformationAnswer.GERANVector
	(*AuthenticationInformationAnswer_UserProfile)(nil),  // 6: protos.AuthenticationInformationAnswer.UserProfile
}
var file_auth_proto_depIdxs = []int32{
	0, // 0: protos.AuthenticationInformationRequest.RatType:type_name -> protos.RAT_Type
	3, // 1: protos.AuthenticationInformationAnswer.EutranVectors:type_name -> protos.AuthenticationInformationAnswer.EUTRANVector
	4, // 2: protos.AuthenticationInformationAnswer.UtranVectors:type_name -> protos.AuthenticationInformationAnswer.UTRANVector
	5, // 3: protos.AuthenticationInformationAnswer.GeranVectors:type_name -> protos.AuthenticationInformationAnswer.GERANVector
	6, // 4: protos.AuthenticationInformationAnswer.Profile:type_name -> protos.AuthenticationInformationAnswer.UserProfile
	1, // 5: protos.SIMAuthenticator.AuthInfo:input_type -> protos.AuthenticationInformationRequest
	2, // 6: protos.SIMAuthenticator.AuthInfo:output_type -> protos.AuthenticationInformationAnswer
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationInformationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationInformationAnswer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationInformationAnswer_EUTRANVector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationInformationAnswer_UTRANVector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationInformationAnswer_GERANVector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationInformationAnswer_UserProfile); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		EnumInfos:         file_auth_proto_enumTypes,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
