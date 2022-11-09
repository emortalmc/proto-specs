// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: service/mc_player_security.proto

package service

import (
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

type YubikeyResponse_Status int32

const (
	// returned by Yubico API
	YubikeyResponse_OK                    YubikeyResponse_Status = 0
	YubikeyResponse_BAD_OTP               YubikeyResponse_Status = 1
	YubikeyResponse_REPLAYED_OTP          YubikeyResponse_Status = 2
	YubikeyResponse_BAD_SIGNATURE         YubikeyResponse_Status = 3
	YubikeyResponse_MISSING_PARAMETER     YubikeyResponse_Status = 4
	YubikeyResponse_NO_SUCH_CLIENT        YubikeyResponse_Status = 5
	YubikeyResponse_OPERATION_NOT_ALLOWED YubikeyResponse_Status = 6
	YubikeyResponse_BACKEND_ERROR         YubikeyResponse_Status = 7
	YubikeyResponse_NOT_ENOUGH_ANSWERS    YubikeyResponse_Status = 8
	YubikeyResponse_REPLAYED_REQUEST      YubikeyResponse_Status = 9
	// returned by this service
	YubikeyResponse_KEY_NOT_LINKED YubikeyResponse_Status = 100
	YubikeyResponse_SECURITY_ERROR YubikeyResponse_Status = 101
)

// Enum value maps for YubikeyResponse_Status.
var (
	YubikeyResponse_Status_name = map[int32]string{
		0:   "OK",
		1:   "BAD_OTP",
		2:   "REPLAYED_OTP",
		3:   "BAD_SIGNATURE",
		4:   "MISSING_PARAMETER",
		5:   "NO_SUCH_CLIENT",
		6:   "OPERATION_NOT_ALLOWED",
		7:   "BACKEND_ERROR",
		8:   "NOT_ENOUGH_ANSWERS",
		9:   "REPLAYED_REQUEST",
		100: "KEY_NOT_LINKED",
		101: "SECURITY_ERROR",
	}
	YubikeyResponse_Status_value = map[string]int32{
		"OK":                    0,
		"BAD_OTP":               1,
		"REPLAYED_OTP":          2,
		"BAD_SIGNATURE":         3,
		"MISSING_PARAMETER":     4,
		"NO_SUCH_CLIENT":        5,
		"OPERATION_NOT_ALLOWED": 6,
		"BACKEND_ERROR":         7,
		"NOT_ENOUGH_ANSWERS":    8,
		"REPLAYED_REQUEST":      9,
		"KEY_NOT_LINKED":        100,
		"SECURITY_ERROR":        101,
	}
)

func (x YubikeyResponse_Status) Enum() *YubikeyResponse_Status {
	p := new(YubikeyResponse_Status)
	*p = x
	return p
}

func (x YubikeyResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (YubikeyResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_service_mc_player_security_proto_enumTypes[0].Descriptor()
}

func (YubikeyResponse_Status) Type() protoreflect.EnumType {
	return &file_service_mc_player_security_proto_enumTypes[0]
}

func (x YubikeyResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use YubikeyResponse_Status.Descriptor instead.
func (YubikeyResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_service_mc_player_security_proto_rawDescGZIP(), []int{1, 0}
}

type YubikeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IssuerId string `protobuf:"bytes,1,opt,name=issuer_id,json=issuerId,proto3" json:"issuer_id,omitempty"`
	Otp      string `protobuf:"bytes,2,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *YubikeyRequest) Reset() {
	*x = YubikeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_mc_player_security_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YubikeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YubikeyRequest) ProtoMessage() {}

func (x *YubikeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_mc_player_security_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YubikeyRequest.ProtoReflect.Descriptor instead.
func (*YubikeyRequest) Descriptor() ([]byte, []int) {
	return file_service_mc_player_security_proto_rawDescGZIP(), []int{0}
}

func (x *YubikeyRequest) GetIssuerId() string {
	if x != nil {
		return x.IssuerId
	}
	return ""
}

func (x *YubikeyRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

type YubikeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status YubikeyResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=service.mc_player_security.YubikeyResponse_Status" json:"status,omitempty"`
}

func (x *YubikeyResponse) Reset() {
	*x = YubikeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_mc_player_security_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YubikeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YubikeyResponse) ProtoMessage() {}

func (x *YubikeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_mc_player_security_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YubikeyResponse.ProtoReflect.Descriptor instead.
func (*YubikeyResponse) Descriptor() ([]byte, []int) {
	return file_service_mc_player_security_proto_rawDescGZIP(), []int{1}
}

func (x *YubikeyResponse) GetStatus() YubikeyResponse_Status {
	if x != nil {
		return x.Status
	}
	return YubikeyResponse_OK
}

var File_service_mc_player_security_proto protoreflect.FileDescriptor

var file_service_mc_player_security_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x63, 0x5f, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x63, 0x5f, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x22, 0x3f,
	0x0a, 0x0e, 0x59, 0x75, 0x62, 0x69, 0x6b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x6f, 0x74, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x22,
	0xd1, 0x02, 0x0a, 0x0f, 0x59, 0x75, 0x62, 0x69, 0x6b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x63,
	0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x59, 0x75, 0x62, 0x69, 0x6b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0xf1, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x41, 0x44, 0x5f, 0x4f, 0x54, 0x50, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x44, 0x5f, 0x4f, 0x54, 0x50, 0x10,
	0x02, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x41, 0x44, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55,
	0x52, 0x45, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f,
	0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x4e,
	0x4f, 0x5f, 0x53, 0x55, 0x43, 0x48, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x05, 0x12,
	0x19, 0x0a, 0x15, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x41,
	0x43, 0x4b, 0x45, 0x4e, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x07, 0x12, 0x16, 0x0a,
	0x12, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f, 0x55, 0x47, 0x48, 0x5f, 0x41, 0x4e, 0x53, 0x57,
	0x45, 0x52, 0x53, 0x10, 0x08, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x59, 0x45,
	0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x09, 0x12, 0x12, 0x0a, 0x0e, 0x4b,
	0x45, 0x59, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x45, 0x44, 0x10, 0x64, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x65, 0x32, 0xe7, 0x01, 0x0a, 0x10, 0x4d, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x67, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x59,
	0x75, 0x62, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x2a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x6d, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x2e, 0x59, 0x75, 0x62, 0x69, 0x6b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x63, 0x5f,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x59, 0x75, 0x62, 0x69, 0x6b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x6a, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x59, 0x75, 0x62, 0x69, 0x6b,
	0x65, 0x79, 0x12, 0x2a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x63, 0x5f,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x59, 0x75, 0x62, 0x69, 0x6b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x59, 0x75, 0x62, 0x69,
	0x6b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x74, 0x0a,
	0x1b, 0x63, 0x63, 0x2e, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x15, 0x4d, 0x63,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2d, 0x63, 0x63, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_mc_player_security_proto_rawDescOnce sync.Once
	file_service_mc_player_security_proto_rawDescData = file_service_mc_player_security_proto_rawDesc
)

func file_service_mc_player_security_proto_rawDescGZIP() []byte {
	file_service_mc_player_security_proto_rawDescOnce.Do(func() {
		file_service_mc_player_security_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_mc_player_security_proto_rawDescData)
	})
	return file_service_mc_player_security_proto_rawDescData
}

var file_service_mc_player_security_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_service_mc_player_security_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_mc_player_security_proto_goTypes = []interface{}{
	(YubikeyResponse_Status)(0), // 0: service.mc_player_security.YubikeyResponse.Status
	(*YubikeyRequest)(nil),      // 1: service.mc_player_security.YubikeyRequest
	(*YubikeyResponse)(nil),     // 2: service.mc_player_security.YubikeyResponse
}
var file_service_mc_player_security_proto_depIdxs = []int32{
	0, // 0: service.mc_player_security.YubikeyResponse.status:type_name -> service.mc_player_security.YubikeyResponse.Status
	1, // 1: service.mc_player_security.McPlayerSecurity.AddYubiKey:input_type -> service.mc_player_security.YubikeyRequest
	1, // 2: service.mc_player_security.McPlayerSecurity.VerifyYubikey:input_type -> service.mc_player_security.YubikeyRequest
	2, // 3: service.mc_player_security.McPlayerSecurity.AddYubiKey:output_type -> service.mc_player_security.YubikeyResponse
	2, // 4: service.mc_player_security.McPlayerSecurity.VerifyYubikey:output_type -> service.mc_player_security.YubikeyResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_mc_player_security_proto_init() }
func file_service_mc_player_security_proto_init() {
	if File_service_mc_player_security_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_mc_player_security_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YubikeyRequest); i {
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
		file_service_mc_player_security_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YubikeyResponse); i {
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
			RawDescriptor: file_service_mc_player_security_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_mc_player_security_proto_goTypes,
		DependencyIndexes: file_service_mc_player_security_proto_depIdxs,
		EnumInfos:         file_service_mc_player_security_proto_enumTypes,
		MessageInfos:      file_service_mc_player_security_proto_msgTypes,
	}.Build()
	File_service_mc_player_security_proto = out.File
	file_service_mc_player_security_proto_rawDesc = nil
	file_service_mc_player_security_proto_goTypes = nil
	file_service_mc_player_security_proto_depIdxs = nil
}
