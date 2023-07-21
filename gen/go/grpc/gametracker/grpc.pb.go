// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: game_tracker/grpc.proto

package gametracker

import (
	common "github.com/emortalmc/proto-specs/gen/go/model/common"
	_ "github.com/emortalmc/proto-specs/gen/go/model/gametracker"
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

type GetPlayerActiveGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// player_id of type UUID
	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *GetPlayerActiveGameRequest) Reset() {
	*x = GetPlayerActiveGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerActiveGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerActiveGameRequest) ProtoMessage() {}

func (x *GetPlayerActiveGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerActiveGameRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerActiveGameRequest) Descriptor() ([]byte, []int) {
	return file_game_tracker_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *GetPlayerActiveGameRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type GetPlayerActiveGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPlayerActiveGameResponse) Reset() {
	*x = GetPlayerActiveGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerActiveGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerActiveGameResponse) ProtoMessage() {}

func (x *GetPlayerActiveGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerActiveGameResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerActiveGameResponse) Descriptor() ([]byte, []int) {
	return file_game_tracker_grpc_proto_rawDescGZIP(), []int{1}
}

type ListGamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filters  *ListGamesRequest_Filters `protobuf:"bytes,1,opt,name=filters,proto3" json:"filters,omitempty"`
	Pageable *common.Pageable          `protobuf:"bytes,2,opt,name=pageable,proto3" json:"pageable,omitempty"`
}

func (x *ListGamesRequest) Reset() {
	*x = ListGamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGamesRequest) ProtoMessage() {}

func (x *ListGamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGamesRequest.ProtoReflect.Descriptor instead.
func (*ListGamesRequest) Descriptor() ([]byte, []int) {
	return file_game_tracker_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *ListGamesRequest) GetFilters() *ListGamesRequest_Filters {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *ListGamesRequest) GetPageable() *common.Pageable {
	if x != nil {
		return x.Pageable
	}
	return nil
}

type ListGamesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListGamesResponse) Reset() {
	*x = ListGamesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGamesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGamesResponse) ProtoMessage() {}

func (x *ListGamesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGamesResponse.ProtoReflect.Descriptor instead.
func (*ListGamesResponse) Descriptor() ([]byte, []int) {
	return file_game_tracker_grpc_proto_rawDescGZIP(), []int{3}
}

type ListGamesRequest_Filters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId   *string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3,oneof" json:"player_id,omitempty"`
	GameModeId *string `protobuf:"bytes,2,opt,name=game_mode_id,json=gameModeId,proto3,oneof" json:"game_mode_id,omitempty"`
	MapId      *string `protobuf:"bytes,3,opt,name=map_id,json=mapId,proto3,oneof" json:"map_id,omitempty"`
}

func (x *ListGamesRequest_Filters) Reset() {
	*x = ListGamesRequest_Filters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGamesRequest_Filters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGamesRequest_Filters) ProtoMessage() {}

func (x *ListGamesRequest_Filters) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGamesRequest_Filters.ProtoReflect.Descriptor instead.
func (*ListGamesRequest_Filters) Descriptor() ([]byte, []int) {
	return file_game_tracker_grpc_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ListGamesRequest_Filters) GetPlayerId() string {
	if x != nil && x.PlayerId != nil {
		return *x.PlayerId
	}
	return ""
}

func (x *ListGamesRequest_Filters) GetGameModeId() string {
	if x != nil && x.GameModeId != nil {
		return *x.GameModeId
	}
	return ""
}

func (x *ListGamesRequest_Filters) GetMapId() string {
	if x != nil && x.MapId != nil {
		return *x.MapId
	}
	return ""
}

var File_game_tracker_grpc_proto protoreflect.FileDescriptor

var file_game_tracker_grpc_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x65, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x1a, 0x19, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x1d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb1,
	0x02, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x1a, 0x98, 0x01, 0x0a, 0x07, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x67,
	0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x06,
	0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05,
	0x6d, 0x61, 0x70, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6d, 0x61, 0x70, 0x5f,
	0x69, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xfc, 0x01, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x84, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x12,
	0x35, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x2b, 0x2e, 0x65, 0x6d,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6e, 0x0a, 0x20, 0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x42, 0x10, 0x47, 0x61, 0x6d, 0x65,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x38, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c,
	0x6d, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_tracker_grpc_proto_rawDescOnce sync.Once
	file_game_tracker_grpc_proto_rawDescData = file_game_tracker_grpc_proto_rawDesc
)

func file_game_tracker_grpc_proto_rawDescGZIP() []byte {
	file_game_tracker_grpc_proto_rawDescOnce.Do(func() {
		file_game_tracker_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_tracker_grpc_proto_rawDescData)
	})
	return file_game_tracker_grpc_proto_rawDescData
}

var file_game_tracker_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_game_tracker_grpc_proto_goTypes = []interface{}{
	(*GetPlayerActiveGameRequest)(nil),  // 0: emortal.grpc.game_tracker.GetPlayerActiveGameRequest
	(*GetPlayerActiveGameResponse)(nil), // 1: emortal.grpc.game_tracker.GetPlayerActiveGameResponse
	(*ListGamesRequest)(nil),            // 2: emortal.grpc.game_tracker.ListGamesRequest
	(*ListGamesResponse)(nil),           // 3: emortal.grpc.game_tracker.ListGamesResponse
	(*ListGamesRequest_Filters)(nil),    // 4: emortal.grpc.game_tracker.ListGamesRequest.Filters
	(*common.Pageable)(nil),             // 5: emortal.model.Pageable
}
var file_game_tracker_grpc_proto_depIdxs = []int32{
	4, // 0: emortal.grpc.game_tracker.ListGamesRequest.filters:type_name -> emortal.grpc.game_tracker.ListGamesRequest.Filters
	5, // 1: emortal.grpc.game_tracker.ListGamesRequest.pageable:type_name -> emortal.model.Pageable
	0, // 2: emortal.grpc.game_tracker.GameTracker.GetPlayerActiveGame:input_type -> emortal.grpc.game_tracker.GetPlayerActiveGameRequest
	2, // 3: emortal.grpc.game_tracker.GameTracker.ListGames:input_type -> emortal.grpc.game_tracker.ListGamesRequest
	1, // 4: emortal.grpc.game_tracker.GameTracker.GetPlayerActiveGame:output_type -> emortal.grpc.game_tracker.GetPlayerActiveGameResponse
	3, // 5: emortal.grpc.game_tracker.GameTracker.ListGames:output_type -> emortal.grpc.game_tracker.ListGamesResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_game_tracker_grpc_proto_init() }
func file_game_tracker_grpc_proto_init() {
	if File_game_tracker_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_tracker_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerActiveGameRequest); i {
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
		file_game_tracker_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerActiveGameResponse); i {
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
		file_game_tracker_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGamesRequest); i {
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
		file_game_tracker_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGamesResponse); i {
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
		file_game_tracker_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGamesRequest_Filters); i {
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
	file_game_tracker_grpc_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_tracker_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_tracker_grpc_proto_goTypes,
		DependencyIndexes: file_game_tracker_grpc_proto_depIdxs,
		MessageInfos:      file_game_tracker_grpc_proto_msgTypes,
	}.Build()
	File_game_tracker_grpc_proto = out.File
	file_game_tracker_grpc_proto_rawDesc = nil
	file_game_tracker_grpc_proto_goTypes = nil
	file_game_tracker_grpc_proto_depIdxs = nil
}