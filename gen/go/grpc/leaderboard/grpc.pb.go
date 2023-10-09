// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: leaderboards/grpc.proto

package leaderboard

import (
	leaderboard "github.com/emortalmc/proto-specs/gen/go/model/leaderboard"
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

type GetLeaderboardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetLeaderboardRequest) Reset() {
	*x = GetLeaderboardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLeaderboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLeaderboardRequest) ProtoMessage() {}

func (x *GetLeaderboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLeaderboardRequest.ProtoReflect.Descriptor instead.
func (*GetLeaderboardRequest) Descriptor() ([]byte, []int) {
	return file_leaderboards_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *GetLeaderboardRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetLeaderboardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Leaderboard *leaderboard.Leaderboard `protobuf:"bytes,1,opt,name=leaderboard,proto3" json:"leaderboard,omitempty"`
}

func (x *GetLeaderboardResponse) Reset() {
	*x = GetLeaderboardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLeaderboardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLeaderboardResponse) ProtoMessage() {}

func (x *GetLeaderboardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLeaderboardResponse.ProtoReflect.Descriptor instead.
func (*GetLeaderboardResponse) Descriptor() ([]byte, []int) {
	return file_leaderboards_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *GetLeaderboardResponse) GetLeaderboard() *leaderboard.Leaderboard {
	if x != nil {
		return x.Leaderboard
	}
	return nil
}

type CreateLeaderboardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SortOrder leaderboard.SortOrder `protobuf:"varint,2,opt,name=sort_order,json=sortOrder,proto3,enum=emortal.model.leaderboard.SortOrder" json:"sort_order,omitempty"`
}

func (x *CreateLeaderboardRequest) Reset() {
	*x = CreateLeaderboardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLeaderboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLeaderboardRequest) ProtoMessage() {}

func (x *CreateLeaderboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLeaderboardRequest.ProtoReflect.Descriptor instead.
func (*CreateLeaderboardRequest) Descriptor() ([]byte, []int) {
	return file_leaderboards_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLeaderboardRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateLeaderboardRequest) GetSortOrder() leaderboard.SortOrder {
	if x != nil {
		return x.SortOrder
	}
	return leaderboard.SortOrder(0)
}

type CreateLeaderboardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Leaderboard *leaderboard.Leaderboard `protobuf:"bytes,1,opt,name=leaderboard,proto3" json:"leaderboard,omitempty"`
}

func (x *CreateLeaderboardResponse) Reset() {
	*x = CreateLeaderboardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLeaderboardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLeaderboardResponse) ProtoMessage() {}

func (x *CreateLeaderboardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLeaderboardResponse.ProtoReflect.Descriptor instead.
func (*CreateLeaderboardResponse) Descriptor() ([]byte, []int) {
	return file_leaderboards_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *CreateLeaderboardResponse) GetLeaderboard() *leaderboard.Leaderboard {
	if x != nil {
		return x.Leaderboard
	}
	return nil
}

type DeleteLeaderboardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteLeaderboardRequest) Reset() {
	*x = DeleteLeaderboardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLeaderboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLeaderboardRequest) ProtoMessage() {}

func (x *DeleteLeaderboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLeaderboardRequest.ProtoReflect.Descriptor instead.
func (*DeleteLeaderboardRequest) Descriptor() ([]byte, []int) {
	return file_leaderboards_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteLeaderboardRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteLeaderboardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteLeaderboardResponse) Reset() {
	*x = DeleteLeaderboardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLeaderboardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLeaderboardResponse) ProtoMessage() {}

func (x *DeleteLeaderboardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLeaderboardResponse.ProtoReflect.Descriptor instead.
func (*DeleteLeaderboardResponse) Descriptor() ([]byte, []int) {
	return file_leaderboards_grpc_proto_rawDescGZIP(), []int{5}
}

type CreateEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaderboardId string                        `protobuf:"bytes,1,opt,name=leaderboard_id,json=leaderboardId,proto3" json:"leaderboard_id,omitempty"`
	Entry         *leaderboard.LeaderboardEntry `protobuf:"bytes,2,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *CreateEntryRequest) Reset() {
	*x = CreateEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_grpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEntryRequest) ProtoMessage() {}

func (x *CreateEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_grpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEntryRequest.ProtoReflect.Descriptor instead.
func (*CreateEntryRequest) Descriptor() ([]byte, []int) {
	return file_leaderboards_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *CreateEntryRequest) GetLeaderboardId() string {
	if x != nil {
		return x.LeaderboardId
	}
	return ""
}

func (x *CreateEntryRequest) GetEntry() *leaderboard.LeaderboardEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type CreateEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateEntryResponse) Reset() {
	*x = CreateEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_grpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEntryResponse) ProtoMessage() {}

func (x *CreateEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_grpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEntryResponse.ProtoReflect.Descriptor instead.
func (*CreateEntryResponse) Descriptor() ([]byte, []int) {
	return file_leaderboards_grpc_proto_rawDescGZIP(), []int{7}
}

type DeleteEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaderboardId string `protobuf:"bytes,1,opt,name=leaderboard_id,json=leaderboardId,proto3" json:"leaderboard_id,omitempty"`
	Key           string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteEntryRequest) Reset() {
	*x = DeleteEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_grpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEntryRequest) ProtoMessage() {}

func (x *DeleteEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_grpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEntryRequest.ProtoReflect.Descriptor instead.
func (*DeleteEntryRequest) Descriptor() ([]byte, []int) {
	return file_leaderboards_grpc_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteEntryRequest) GetLeaderboardId() string {
	if x != nil {
		return x.LeaderboardId
	}
	return ""
}

func (x *DeleteEntryRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteEntryResponse) Reset() {
	*x = DeleteEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_grpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEntryResponse) ProtoMessage() {}

func (x *DeleteEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_grpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEntryResponse.ProtoReflect.Descriptor instead.
func (*DeleteEntryResponse) Descriptor() ([]byte, []int) {
	return file_leaderboards_grpc_proto_rawDescGZIP(), []int{9}
}

type UpdateScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaderboardId string  `protobuf:"bytes,1,opt,name=leaderboard_id,json=leaderboardId,proto3" json:"leaderboard_id,omitempty"`
	Key           string  `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Score         float32 `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *UpdateScoreRequest) Reset() {
	*x = UpdateScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_grpc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScoreRequest) ProtoMessage() {}

func (x *UpdateScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_grpc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScoreRequest.ProtoReflect.Descriptor instead.
func (*UpdateScoreRequest) Descriptor() ([]byte, []int) {
	return file_leaderboards_grpc_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateScoreRequest) GetLeaderboardId() string {
	if x != nil {
		return x.LeaderboardId
	}
	return ""
}

func (x *UpdateScoreRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UpdateScoreRequest) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type UpdateScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateScoreResponse) Reset() {
	*x = UpdateScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_grpc_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScoreResponse) ProtoMessage() {}

func (x *UpdateScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_grpc_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScoreResponse.ProtoReflect.Descriptor instead.
func (*UpdateScoreResponse) Descriptor() ([]byte, []int) {
	return file_leaderboards_grpc_proto_rawDescGZIP(), []int{11}
}

var File_leaderboards_grpc_proto protoreflect.FileDescriptor

var file_leaderboards_grpc_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x65, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x1a, 0x19, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x48, 0x0a, 0x0b, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x0b,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x22, 0x6f, 0x0a, 0x18, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x5f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x65, 0x6d,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x65, 0x0a, 0x19,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x0b, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x22, 0x2a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x1b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7e, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x05, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x15, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x4d, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x63, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x15,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc4, 0x04, 0x0a, 0x0b, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x73, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x2f, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7c, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12,
	0x32, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x2c, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x2c, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x6a, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x2c, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e,
	0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6e, 0x0a, 0x20,
	0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x42, 0x10, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6d, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73,
	0x70, 0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_leaderboards_grpc_proto_rawDescOnce sync.Once
	file_leaderboards_grpc_proto_rawDescData = file_leaderboards_grpc_proto_rawDesc
)

func file_leaderboards_grpc_proto_rawDescGZIP() []byte {
	file_leaderboards_grpc_proto_rawDescOnce.Do(func() {
		file_leaderboards_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_leaderboards_grpc_proto_rawDescData)
	})
	return file_leaderboards_grpc_proto_rawDescData
}

var file_leaderboards_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_leaderboards_grpc_proto_goTypes = []interface{}{
	(*GetLeaderboardRequest)(nil),        // 0: emortal.grpc.leaderboard.GetLeaderboardRequest
	(*GetLeaderboardResponse)(nil),       // 1: emortal.grpc.leaderboard.GetLeaderboardResponse
	(*CreateLeaderboardRequest)(nil),     // 2: emortal.grpc.leaderboard.CreateLeaderboardRequest
	(*CreateLeaderboardResponse)(nil),    // 3: emortal.grpc.leaderboard.CreateLeaderboardResponse
	(*DeleteLeaderboardRequest)(nil),     // 4: emortal.grpc.leaderboard.DeleteLeaderboardRequest
	(*DeleteLeaderboardResponse)(nil),    // 5: emortal.grpc.leaderboard.DeleteLeaderboardResponse
	(*CreateEntryRequest)(nil),           // 6: emortal.grpc.leaderboard.CreateEntryRequest
	(*CreateEntryResponse)(nil),          // 7: emortal.grpc.leaderboard.CreateEntryResponse
	(*DeleteEntryRequest)(nil),           // 8: emortal.grpc.leaderboard.DeleteEntryRequest
	(*DeleteEntryResponse)(nil),          // 9: emortal.grpc.leaderboard.DeleteEntryResponse
	(*UpdateScoreRequest)(nil),           // 10: emortal.grpc.leaderboard.UpdateScoreRequest
	(*UpdateScoreResponse)(nil),          // 11: emortal.grpc.leaderboard.UpdateScoreResponse
	(*leaderboard.Leaderboard)(nil),      // 12: emortal.model.leaderboard.Leaderboard
	(leaderboard.SortOrder)(0),           // 13: emortal.model.leaderboard.SortOrder
	(*leaderboard.LeaderboardEntry)(nil), // 14: emortal.model.leaderboard.LeaderboardEntry
}
var file_leaderboards_grpc_proto_depIdxs = []int32{
	12, // 0: emortal.grpc.leaderboard.GetLeaderboardResponse.leaderboard:type_name -> emortal.model.leaderboard.Leaderboard
	13, // 1: emortal.grpc.leaderboard.CreateLeaderboardRequest.sort_order:type_name -> emortal.model.leaderboard.SortOrder
	12, // 2: emortal.grpc.leaderboard.CreateLeaderboardResponse.leaderboard:type_name -> emortal.model.leaderboard.Leaderboard
	14, // 3: emortal.grpc.leaderboard.CreateEntryRequest.entry:type_name -> emortal.model.leaderboard.LeaderboardEntry
	0,  // 4: emortal.grpc.leaderboard.Leaderboard.GetLeaderboard:input_type -> emortal.grpc.leaderboard.GetLeaderboardRequest
	2,  // 5: emortal.grpc.leaderboard.Leaderboard.CreateLeaderboard:input_type -> emortal.grpc.leaderboard.CreateLeaderboardRequest
	6,  // 6: emortal.grpc.leaderboard.Leaderboard.CreateEntry:input_type -> emortal.grpc.leaderboard.CreateEntryRequest
	8,  // 7: emortal.grpc.leaderboard.Leaderboard.DeleteEntry:input_type -> emortal.grpc.leaderboard.DeleteEntryRequest
	10, // 8: emortal.grpc.leaderboard.Leaderboard.UpdateScore:input_type -> emortal.grpc.leaderboard.UpdateScoreRequest
	1,  // 9: emortal.grpc.leaderboard.Leaderboard.GetLeaderboard:output_type -> emortal.grpc.leaderboard.GetLeaderboardResponse
	3,  // 10: emortal.grpc.leaderboard.Leaderboard.CreateLeaderboard:output_type -> emortal.grpc.leaderboard.CreateLeaderboardResponse
	7,  // 11: emortal.grpc.leaderboard.Leaderboard.CreateEntry:output_type -> emortal.grpc.leaderboard.CreateEntryResponse
	9,  // 12: emortal.grpc.leaderboard.Leaderboard.DeleteEntry:output_type -> emortal.grpc.leaderboard.DeleteEntryResponse
	11, // 13: emortal.grpc.leaderboard.Leaderboard.UpdateScore:output_type -> emortal.grpc.leaderboard.UpdateScoreResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_leaderboards_grpc_proto_init() }
func file_leaderboards_grpc_proto_init() {
	if File_leaderboards_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_leaderboards_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLeaderboardRequest); i {
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
		file_leaderboards_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLeaderboardResponse); i {
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
		file_leaderboards_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLeaderboardRequest); i {
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
		file_leaderboards_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLeaderboardResponse); i {
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
		file_leaderboards_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLeaderboardRequest); i {
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
		file_leaderboards_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLeaderboardResponse); i {
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
		file_leaderboards_grpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEntryRequest); i {
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
		file_leaderboards_grpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEntryResponse); i {
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
		file_leaderboards_grpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEntryRequest); i {
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
		file_leaderboards_grpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEntryResponse); i {
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
		file_leaderboards_grpc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateScoreRequest); i {
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
		file_leaderboards_grpc_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateScoreResponse); i {
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
			RawDescriptor: file_leaderboards_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_leaderboards_grpc_proto_goTypes,
		DependencyIndexes: file_leaderboards_grpc_proto_depIdxs,
		MessageInfos:      file_leaderboards_grpc_proto_msgTypes,
	}.Build()
	File_leaderboards_grpc_proto = out.File
	file_leaderboards_grpc_proto_rawDesc = nil
	file_leaderboards_grpc_proto_goTypes = nil
	file_leaderboards_grpc_proto_depIdxs = nil
}