// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: model/player.proto

package model

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

type PlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *PlayerRequest) Reset() {
	*x = PlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerRequest) ProtoMessage() {}

func (x *PlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_model_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerRequest.ProtoReflect.Descriptor instead.
func (*PlayerRequest) Descriptor() ([]byte, []int) {
	return file_model_player_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type PlayersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerIds []string `protobuf:"bytes,1,rep,name=player_ids,json=playerIds,proto3" json:"player_ids,omitempty"`
}

func (x *PlayersRequest) Reset() {
	*x = PlayersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayersRequest) ProtoMessage() {}

func (x *PlayersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_model_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayersRequest.ProtoReflect.Descriptor instead.
func (*PlayersRequest) Descriptor() ([]byte, []int) {
	return file_model_player_proto_rawDescGZIP(), []int{1}
}

func (x *PlayersRequest) GetPlayerIds() []string {
	if x != nil {
		return x.PlayerIds
	}
	return nil
}

type PlayerUsernameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"` // ignore case when using the username
}

func (x *PlayerUsernameRequest) Reset() {
	*x = PlayerUsernameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerUsernameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerUsernameRequest) ProtoMessage() {}

func (x *PlayerUsernameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_model_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerUsernameRequest.ProtoReflect.Descriptor instead.
func (*PlayerUsernameRequest) Descriptor() ([]byte, []int) {
	return file_model_player_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerUsernameRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_model_player_proto protoreflect.FileDescriptor

var file_model_player_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x2c, 0x0a, 0x0d, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x0e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x33, 0x0a, 0x15, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x66, 0x0a, 0x19, 0x63, 0x63, 0x2e, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0b, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e,
	0x63, 0x65, 0x2d, 0x63, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73,
	0x70, 0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67,
	0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_player_proto_rawDescOnce sync.Once
	file_model_player_proto_rawDescData = file_model_player_proto_rawDesc
)

func file_model_player_proto_rawDescGZIP() []byte {
	file_model_player_proto_rawDescOnce.Do(func() {
		file_model_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_player_proto_rawDescData)
	})
	return file_model_player_proto_rawDescData
}

var file_model_player_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_model_player_proto_goTypes = []interface{}{
	(*PlayerRequest)(nil),         // 0: model.PlayerRequest
	(*PlayersRequest)(nil),        // 1: model.PlayersRequest
	(*PlayerUsernameRequest)(nil), // 2: model.PlayerUsernameRequest
}
var file_model_player_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_model_player_proto_init() }
func file_model_player_proto_init() {
	if File_model_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerRequest); i {
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
		file_model_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayersRequest); i {
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
		file_model_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerUsernameRequest); i {
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
			RawDescriptor: file_model_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_player_proto_goTypes,
		DependencyIndexes: file_model_player_proto_depIdxs,
		MessageInfos:      file_model_player_proto_msgTypes,
	}.Build()
	File_model_player_proto = out.File
	file_model_player_proto_rawDesc = nil
	file_model_player_proto_goTypes = nil
	file_model_player_proto_depIdxs = nil
}
