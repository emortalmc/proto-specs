// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: account_conn_manager/grpc.proto

package accountconnmanager

import (
	accountconnmanager "github.com/emortalmc/proto-specs/gen/go/model/accountconnmanager"
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

type CreateOAuthLinkErrorResponse_Reason int32

const (
	CreateOAuthLinkErrorResponse_INVALID_SOURCE_ID CreateOAuthLinkErrorResponse_Reason = 0
	// CONNECTION_ALREADY_EXISTS note this is only that this user already has this connection, we don't know the target conn ID yet.
	CreateOAuthLinkErrorResponse_CONNECTION_ALREADY_EXISTS CreateOAuthLinkErrorResponse_Reason = 1
)

// Enum value maps for CreateOAuthLinkErrorResponse_Reason.
var (
	CreateOAuthLinkErrorResponse_Reason_name = map[int32]string{
		0: "INVALID_SOURCE_ID",
		1: "CONNECTION_ALREADY_EXISTS",
	}
	CreateOAuthLinkErrorResponse_Reason_value = map[string]int32{
		"INVALID_SOURCE_ID":         0,
		"CONNECTION_ALREADY_EXISTS": 1,
	}
)

func (x CreateOAuthLinkErrorResponse_Reason) Enum() *CreateOAuthLinkErrorResponse_Reason {
	p := new(CreateOAuthLinkErrorResponse_Reason)
	*p = x
	return p
}

func (x CreateOAuthLinkErrorResponse_Reason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateOAuthLinkErrorResponse_Reason) Descriptor() protoreflect.EnumDescriptor {
	return file_account_conn_manager_grpc_proto_enumTypes[0].Descriptor()
}

func (CreateOAuthLinkErrorResponse_Reason) Type() protoreflect.EnumType {
	return &file_account_conn_manager_grpc_proto_enumTypes[0]
}

func (x CreateOAuthLinkErrorResponse_Reason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateOAuthLinkErrorResponse_Reason.Descriptor instead.
func (CreateOAuthLinkErrorResponse_Reason) EnumDescriptor() ([]byte, []int) {
	return file_account_conn_manager_grpc_proto_rawDescGZIP(), []int{4, 0}
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Id:
	//
	//	*GetUserRequest_MinecraftId
	//	*GetUserRequest_DiscordId
	//	*GetUserRequest_GithubId
	Id isGetUserRequest_Id `protobuf_oneof:"id"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_conn_manager_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_conn_manager_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_account_conn_manager_grpc_proto_rawDescGZIP(), []int{0}
}

func (m *GetUserRequest) GetId() isGetUserRequest_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (x *GetUserRequest) GetMinecraftId() string {
	if x, ok := x.GetId().(*GetUserRequest_MinecraftId); ok {
		return x.MinecraftId
	}
	return ""
}

func (x *GetUserRequest) GetDiscordId() uint64 {
	if x, ok := x.GetId().(*GetUserRequest_DiscordId); ok {
		return x.DiscordId
	}
	return 0
}

func (x *GetUserRequest) GetGithubId() uint64 {
	if x, ok := x.GetId().(*GetUserRequest_GithubId); ok {
		return x.GithubId
	}
	return 0
}

type isGetUserRequest_Id interface {
	isGetUserRequest_Id()
}

type GetUserRequest_MinecraftId struct {
	MinecraftId string `protobuf:"bytes,1,opt,name=minecraft_id,json=minecraftId,proto3,oneof"`
}

type GetUserRequest_DiscordId struct {
	DiscordId uint64 `protobuf:"varint,2,opt,name=discord_id,json=discordId,proto3,oneof"`
}

type GetUserRequest_GithubId struct {
	GithubId uint64 `protobuf:"varint,3,opt,name=github_id,json=githubId,proto3,oneof"`
}

func (*GetUserRequest_MinecraftId) isGetUserRequest_Id() {}

func (*GetUserRequest_DiscordId) isGetUserRequest_Id() {}

func (*GetUserRequest_GithubId) isGetUserRequest_Id() {}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *accountconnmanager.ConnectionUser `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_conn_manager_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_conn_manager_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_account_conn_manager_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserResponse) GetUser() *accountconnmanager.ConnectionUser {
	if x != nil {
		return x.User
	}
	return nil
}

// CreateOAuthLinkRequest requests an OAuth URL to link an account to another.
type CreateOAuthLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to SourceId:
	//
	//	*CreateOAuthLinkRequest_SourceMinecraftId
	//	*CreateOAuthLinkRequest_SourceDiscordId
	//	*CreateOAuthLinkRequest_SourceGithubId
	SourceId         isCreateOAuthLinkRequest_SourceId `protobuf_oneof:"source_id"`
	TargetConnection accountconnmanager.ConnectionType `protobuf:"varint,4,opt,name=target_connection,json=targetConnection,proto3,enum=emortal.model.account_conn_manager.ConnectionType" json:"target_connection,omitempty"`
}

func (x *CreateOAuthLinkRequest) Reset() {
	*x = CreateOAuthLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_conn_manager_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOAuthLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOAuthLinkRequest) ProtoMessage() {}

func (x *CreateOAuthLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_conn_manager_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOAuthLinkRequest.ProtoReflect.Descriptor instead.
func (*CreateOAuthLinkRequest) Descriptor() ([]byte, []int) {
	return file_account_conn_manager_grpc_proto_rawDescGZIP(), []int{2}
}

func (m *CreateOAuthLinkRequest) GetSourceId() isCreateOAuthLinkRequest_SourceId {
	if m != nil {
		return m.SourceId
	}
	return nil
}

func (x *CreateOAuthLinkRequest) GetSourceMinecraftId() string {
	if x, ok := x.GetSourceId().(*CreateOAuthLinkRequest_SourceMinecraftId); ok {
		return x.SourceMinecraftId
	}
	return ""
}

func (x *CreateOAuthLinkRequest) GetSourceDiscordId() uint64 {
	if x, ok := x.GetSourceId().(*CreateOAuthLinkRequest_SourceDiscordId); ok {
		return x.SourceDiscordId
	}
	return 0
}

func (x *CreateOAuthLinkRequest) GetSourceGithubId() uint64 {
	if x, ok := x.GetSourceId().(*CreateOAuthLinkRequest_SourceGithubId); ok {
		return x.SourceGithubId
	}
	return 0
}

func (x *CreateOAuthLinkRequest) GetTargetConnection() accountconnmanager.ConnectionType {
	if x != nil {
		return x.TargetConnection
	}
	return accountconnmanager.ConnectionType(0)
}

type isCreateOAuthLinkRequest_SourceId interface {
	isCreateOAuthLinkRequest_SourceId()
}

type CreateOAuthLinkRequest_SourceMinecraftId struct {
	SourceMinecraftId string `protobuf:"bytes,1,opt,name=source_minecraft_id,json=sourceMinecraftId,proto3,oneof"`
}

type CreateOAuthLinkRequest_SourceDiscordId struct {
	SourceDiscordId uint64 `protobuf:"varint,2,opt,name=source_discord_id,json=sourceDiscordId,proto3,oneof"`
}

type CreateOAuthLinkRequest_SourceGithubId struct {
	SourceGithubId uint64 `protobuf:"varint,3,opt,name=source_github_id,json=sourceGithubId,proto3,oneof"`
}

func (*CreateOAuthLinkRequest_SourceMinecraftId) isCreateOAuthLinkRequest_SourceId() {}

func (*CreateOAuthLinkRequest_SourceDiscordId) isCreateOAuthLinkRequest_SourceId() {}

func (*CreateOAuthLinkRequest_SourceGithubId) isCreateOAuthLinkRequest_SourceId() {}

type CreateOAuthLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateOAuthLinkResponse) Reset() {
	*x = CreateOAuthLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_conn_manager_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOAuthLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOAuthLinkResponse) ProtoMessage() {}

func (x *CreateOAuthLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_conn_manager_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOAuthLinkResponse.ProtoReflect.Descriptor instead.
func (*CreateOAuthLinkResponse) Descriptor() ([]byte, []int) {
	return file_account_conn_manager_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOAuthLinkResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type CreateOAuthLinkErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateOAuthLinkErrorResponse) Reset() {
	*x = CreateOAuthLinkErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_conn_manager_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOAuthLinkErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOAuthLinkErrorResponse) ProtoMessage() {}

func (x *CreateOAuthLinkErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_conn_manager_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOAuthLinkErrorResponse.ProtoReflect.Descriptor instead.
func (*CreateOAuthLinkErrorResponse) Descriptor() ([]byte, []int) {
	return file_account_conn_manager_grpc_proto_rawDescGZIP(), []int{4}
}

var File_account_conn_manager_grpc_proto protoreflect.FileDescriptor

var file_account_conn_manager_grpc_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x21, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x1a, 0x21, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0c, 0x6d, 0x69, 0x6e,
	0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x48, 0x00, 0x52, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x09, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x48, 0x00, 0x52, 0x08, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x49, 0x64, 0x42, 0x04,
	0x0a, 0x02, 0x69, 0x64, 0x22, 0x59, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x92, 0x02, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x11, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4d, 0x69, 0x6e, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x49, 0x64, 0x12, 0x5f, 0x0a, 0x11, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x32, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x5e, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68,
	0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3e, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x11, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10,
	0x01, 0x32, 0x97, 0x02, 0x0a, 0x18, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x70,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x31, 0x2e, 0x65, 0x6d, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x65,
	0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x88, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68,
	0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x39, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x41, 0x75, 0x74, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3a, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x83, 0x01, 0x0a, 0x27,
	0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x63, 0x6f, 0x6e, 0x6e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x42, 0x17, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x6d, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70, 0x65,
	0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x63, 0x6f, 0x6e, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_conn_manager_grpc_proto_rawDescOnce sync.Once
	file_account_conn_manager_grpc_proto_rawDescData = file_account_conn_manager_grpc_proto_rawDesc
)

func file_account_conn_manager_grpc_proto_rawDescGZIP() []byte {
	file_account_conn_manager_grpc_proto_rawDescOnce.Do(func() {
		file_account_conn_manager_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_conn_manager_grpc_proto_rawDescData)
	})
	return file_account_conn_manager_grpc_proto_rawDescData
}

var file_account_conn_manager_grpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_account_conn_manager_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_account_conn_manager_grpc_proto_goTypes = []interface{}{
	(CreateOAuthLinkErrorResponse_Reason)(0),  // 0: emortal.grpc.account_conn_manager.CreateOAuthLinkErrorResponse.Reason
	(*GetUserRequest)(nil),                    // 1: emortal.grpc.account_conn_manager.GetUserRequest
	(*GetUserResponse)(nil),                   // 2: emortal.grpc.account_conn_manager.GetUserResponse
	(*CreateOAuthLinkRequest)(nil),            // 3: emortal.grpc.account_conn_manager.CreateOAuthLinkRequest
	(*CreateOAuthLinkResponse)(nil),           // 4: emortal.grpc.account_conn_manager.CreateOAuthLinkResponse
	(*CreateOAuthLinkErrorResponse)(nil),      // 5: emortal.grpc.account_conn_manager.CreateOAuthLinkErrorResponse
	(*accountconnmanager.ConnectionUser)(nil), // 6: emortal.model.account_conn_manager.ConnectionUser
	(accountconnmanager.ConnectionType)(0),    // 7: emortal.model.account_conn_manager.ConnectionType
}
var file_account_conn_manager_grpc_proto_depIdxs = []int32{
	6, // 0: emortal.grpc.account_conn_manager.GetUserResponse.user:type_name -> emortal.model.account_conn_manager.ConnectionUser
	7, // 1: emortal.grpc.account_conn_manager.CreateOAuthLinkRequest.target_connection:type_name -> emortal.model.account_conn_manager.ConnectionType
	1, // 2: emortal.grpc.account_conn_manager.AccountConnectionManager.GetUser:input_type -> emortal.grpc.account_conn_manager.GetUserRequest
	3, // 3: emortal.grpc.account_conn_manager.AccountConnectionManager.CreateOAuthLink:input_type -> emortal.grpc.account_conn_manager.CreateOAuthLinkRequest
	2, // 4: emortal.grpc.account_conn_manager.AccountConnectionManager.GetUser:output_type -> emortal.grpc.account_conn_manager.GetUserResponse
	4, // 5: emortal.grpc.account_conn_manager.AccountConnectionManager.CreateOAuthLink:output_type -> emortal.grpc.account_conn_manager.CreateOAuthLinkResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_account_conn_manager_grpc_proto_init() }
func file_account_conn_manager_grpc_proto_init() {
	if File_account_conn_manager_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_conn_manager_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_account_conn_manager_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
		file_account_conn_manager_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOAuthLinkRequest); i {
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
		file_account_conn_manager_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOAuthLinkResponse); i {
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
		file_account_conn_manager_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOAuthLinkErrorResponse); i {
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
	file_account_conn_manager_grpc_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GetUserRequest_MinecraftId)(nil),
		(*GetUserRequest_DiscordId)(nil),
		(*GetUserRequest_GithubId)(nil),
	}
	file_account_conn_manager_grpc_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CreateOAuthLinkRequest_SourceMinecraftId)(nil),
		(*CreateOAuthLinkRequest_SourceDiscordId)(nil),
		(*CreateOAuthLinkRequest_SourceGithubId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_account_conn_manager_grpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_conn_manager_grpc_proto_goTypes,
		DependencyIndexes: file_account_conn_manager_grpc_proto_depIdxs,
		EnumInfos:         file_account_conn_manager_grpc_proto_enumTypes,
		MessageInfos:      file_account_conn_manager_grpc_proto_msgTypes,
	}.Build()
	File_account_conn_manager_grpc_proto = out.File
	file_account_conn_manager_grpc_proto_rawDesc = nil
	file_account_conn_manager_grpc_proto_goTypes = nil
	file_account_conn_manager_grpc_proto_depIdxs = nil
}