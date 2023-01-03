// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: server_discovery.proto

package server_discovery

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerCount      int64 `protobuf:"varint,1,opt,name=player_count,json=playerCount,proto3" json:"player_count,omitempty"`
	EnsureSameServer *bool `protobuf:"varint,2,opt,name=ensure_same_server,json=ensureSameServer,proto3,oneof" json:"ensure_same_server,omitempty"`
}

func (x *ServerRequest) Reset() {
	*x = ServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_discovery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerRequest) ProtoMessage() {}

func (x *ServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_discovery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerRequest.ProtoReflect.Descriptor instead.
func (*ServerRequest) Descriptor() ([]byte, []int) {
	return file_server_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *ServerRequest) GetPlayerCount() int64 {
	if x != nil {
		return x.PlayerCount
	}
	return 0
}

func (x *ServerRequest) GetEnsureSameServer() bool {
	if x != nil && x.EnsureSameServer != nil {
		return *x.EnsureSameServer
	}
	return false
}

type TowerDefenceServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerCount int64 `protobuf:"varint,1,opt,name=player_count,json=playerCount,proto3" json:"player_count,omitempty"`
	InProgress  bool  `protobuf:"varint,2,opt,name=in_progress,json=inProgress,proto3" json:"in_progress,omitempty"`
}

func (x *TowerDefenceServerRequest) Reset() {
	*x = TowerDefenceServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_discovery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TowerDefenceServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TowerDefenceServerRequest) ProtoMessage() {}

func (x *TowerDefenceServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_discovery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TowerDefenceServerRequest.ProtoReflect.Descriptor instead.
func (*TowerDefenceServerRequest) Descriptor() ([]byte, []int) {
	return file_server_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *TowerDefenceServerRequest) GetPlayerCount() int64 {
	if x != nil {
		return x.PlayerCount
	}
	return 0
}

func (x *TowerDefenceServerRequest) GetInProgress() bool {
	if x != nil {
		return x.InProgress
	}
	return false
}

type ConnectableServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Port    uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ConnectableServer) Reset() {
	*x = ConnectableServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_discovery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectableServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectableServer) ProtoMessage() {}

func (x *ConnectableServer) ProtoReflect() protoreflect.Message {
	mi := &file_server_discovery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectableServer.ProtoReflect.Descriptor instead.
func (*ConnectableServer) Descriptor() ([]byte, []int) {
	return file_server_discovery_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectableServer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConnectableServer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ConnectableServer) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type LobbyServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectableServer *ConnectableServer `protobuf:"bytes,1,opt,name=connectable_server,json=connectableServer,proto3" json:"connectable_server,omitempty"` //  uint32 number = 2; // Lobby #x - currently not assigned
}

func (x *LobbyServer) Reset() {
	*x = LobbyServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_discovery_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LobbyServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LobbyServer) ProtoMessage() {}

func (x *LobbyServer) ProtoReflect() protoreflect.Message {
	mi := &file_server_discovery_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LobbyServer.ProtoReflect.Descriptor instead.
func (*LobbyServer) Descriptor() ([]byte, []int) {
	return file_server_discovery_proto_rawDescGZIP(), []int{3}
}

func (x *LobbyServer) GetConnectableServer() *ConnectableServer {
	if x != nil {
		return x.ConnectableServer
	}
	return nil
}

var File_server_discovery_proto protoreflect.FileDescriptor

var file_server_discovery_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x64,
	0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7c, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x12, 0x65, 0x6e, 0x73, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x61,
	0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x00, 0x52, 0x10, 0x65, 0x6e, 0x73, 0x75, 0x72, 0x65, 0x53, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x65, 0x6e, 0x73, 0x75, 0x72,
	0x65, 0x5f, 0x73, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x5f, 0x0a,
	0x19, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x44, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x51,
	0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x22, 0x79, 0x0a, 0x0b, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x6a, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x74,
	0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x32, 0xb0, 0x03, 0x0a,
	0x0f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x12, 0x89, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x37, 0x2e, 0x74,
	0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x6c, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4f, 0x74, 0x70, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x3b, 0x2e,
	0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0xa2, 0x01, 0x0a, 0x1e, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x6f, 0x77, 0x65, 0x72,
	0x44, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x43, 0x2e,
	0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x63, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x44, 0x65,
	0x66, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x63, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42,
	0x74, 0x0a, 0x17, 0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x14, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x4d, 0x43, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2d,
	0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_discovery_proto_rawDescOnce sync.Once
	file_server_discovery_proto_rawDescData = file_server_discovery_proto_rawDesc
)

func file_server_discovery_proto_rawDescGZIP() []byte {
	file_server_discovery_proto_rawDescOnce.Do(func() {
		file_server_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_discovery_proto_rawDescData)
	})
	return file_server_discovery_proto_rawDescData
}

var file_server_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_server_discovery_proto_goTypes = []interface{}{
	(*ServerRequest)(nil),             // 0: towerdefence.cc.service.server_discovery.ServerRequest
	(*TowerDefenceServerRequest)(nil), // 1: towerdefence.cc.service.server_discovery.TowerDefenceServerRequest
	(*ConnectableServer)(nil),         // 2: towerdefence.cc.service.server_discovery.ConnectableServer
	(*LobbyServer)(nil),               // 3: towerdefence.cc.service.server_discovery.LobbyServer
	(*emptypb.Empty)(nil),             // 4: google.protobuf.Empty
}
var file_server_discovery_proto_depIdxs = []int32{
	2, // 0: towerdefence.cc.service.server_discovery.LobbyServer.connectable_server:type_name -> towerdefence.cc.service.server_discovery.ConnectableServer
	0, // 1: towerdefence.cc.service.server_discovery.ServerDiscovery.GetSuggestedLobbyServer:input_type -> towerdefence.cc.service.server_discovery.ServerRequest
	4, // 2: towerdefence.cc.service.server_discovery.ServerDiscovery.GetSuggestedOtpServer:input_type -> google.protobuf.Empty
	1, // 3: towerdefence.cc.service.server_discovery.ServerDiscovery.GetSuggestedTowerDefenceServer:input_type -> towerdefence.cc.service.server_discovery.TowerDefenceServerRequest
	3, // 4: towerdefence.cc.service.server_discovery.ServerDiscovery.GetSuggestedLobbyServer:output_type -> towerdefence.cc.service.server_discovery.LobbyServer
	2, // 5: towerdefence.cc.service.server_discovery.ServerDiscovery.GetSuggestedOtpServer:output_type -> towerdefence.cc.service.server_discovery.ConnectableServer
	2, // 6: towerdefence.cc.service.server_discovery.ServerDiscovery.GetSuggestedTowerDefenceServer:output_type -> towerdefence.cc.service.server_discovery.ConnectableServer
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_server_discovery_proto_init() }
func file_server_discovery_proto_init() {
	if File_server_discovery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_discovery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerRequest); i {
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
		file_server_discovery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TowerDefenceServerRequest); i {
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
		file_server_discovery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectableServer); i {
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
		file_server_discovery_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LobbyServer); i {
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
	file_server_discovery_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_discovery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_discovery_proto_goTypes,
		DependencyIndexes: file_server_discovery_proto_depIdxs,
		MessageInfos:      file_server_discovery_proto_msgTypes,
	}.Build()
	File_server_discovery_proto = out.File
	file_server_discovery_proto_rawDesc = nil
	file_server_discovery_proto_goTypes = nil
	file_server_discovery_proto_depIdxs = nil
}
