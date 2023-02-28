// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: party/messages.proto

package party

import (
	party "github.com/emortalmc/proto-specs/gen/go/model/party"
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

type PartyCreatedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Party *party.Party `protobuf:"bytes,1,opt,name=party,proto3" json:"party,omitempty"`
}

func (x *PartyCreatedMessage) Reset() {
	*x = PartyCreatedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyCreatedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyCreatedMessage) ProtoMessage() {}

func (x *PartyCreatedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyCreatedMessage.ProtoReflect.Descriptor instead.
func (*PartyCreatedMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{0}
}

func (x *PartyCreatedMessage) GetParty() *party.Party {
	if x != nil {
		return x.Party
	}
	return nil
}

type PartyDisbandedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Party *party.Party `protobuf:"bytes,1,opt,name=party,proto3" json:"party,omitempty"`
}

func (x *PartyDisbandedMessage) Reset() {
	*x = PartyDisbandedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyDisbandedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyDisbandedMessage) ProtoMessage() {}

func (x *PartyDisbandedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyDisbandedMessage.ProtoReflect.Descriptor instead.
func (*PartyDisbandedMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{1}
}

func (x *PartyDisbandedMessage) GetParty() *party.Party {
	if x != nil {
		return x.Party
	}
	return nil
}

type PartyInviteCreatedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invite *party.PartyInvite `protobuf:"bytes,1,opt,name=invite,proto3" json:"invite,omitempty"`
}

func (x *PartyInviteCreatedMessage) Reset() {
	*x = PartyInviteCreatedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyInviteCreatedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyInviteCreatedMessage) ProtoMessage() {}

func (x *PartyInviteCreatedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyInviteCreatedMessage.ProtoReflect.Descriptor instead.
func (*PartyInviteCreatedMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{2}
}

func (x *PartyInviteCreatedMessage) GetInvite() *party.PartyInvite {
	if x != nil {
		return x.Invite
	}
	return nil
}

type PartyPlayerJoinedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// party_id of type mongo ObjectId
	PartyId string             `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	Member  *party.PartyMember `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *PartyPlayerJoinedMessage) Reset() {
	*x = PartyPlayerJoinedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyPlayerJoinedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyPlayerJoinedMessage) ProtoMessage() {}

func (x *PartyPlayerJoinedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyPlayerJoinedMessage.ProtoReflect.Descriptor instead.
func (*PartyPlayerJoinedMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{3}
}

func (x *PartyPlayerJoinedMessage) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyPlayerJoinedMessage) GetMember() *party.PartyMember {
	if x != nil {
		return x.Member
	}
	return nil
}

type PartyPlayerLeftMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// party_id of type mongo ObjectId
	PartyId string             `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	Member  *party.PartyMember `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *PartyPlayerLeftMessage) Reset() {
	*x = PartyPlayerLeftMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyPlayerLeftMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyPlayerLeftMessage) ProtoMessage() {}

func (x *PartyPlayerLeftMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyPlayerLeftMessage.ProtoReflect.Descriptor instead.
func (*PartyPlayerLeftMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{4}
}

func (x *PartyPlayerLeftMessage) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyPlayerLeftMessage) GetMember() *party.PartyMember {
	if x != nil {
		return x.Member
	}
	return nil
}

type PartyPlayerKickedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// party_id of type mongo ObjectId
	PartyId      string             `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	KickedMember *party.PartyMember `protobuf:"bytes,2,opt,name=kicked_member,json=kickedMember,proto3" json:"kicked_member,omitempty"`
	KickerMember *party.PartyMember `protobuf:"bytes,3,opt,name=kicker_member,json=kickerMember,proto3" json:"kicker_member,omitempty"`
}

func (x *PartyPlayerKickedMessage) Reset() {
	*x = PartyPlayerKickedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyPlayerKickedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyPlayerKickedMessage) ProtoMessage() {}

func (x *PartyPlayerKickedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyPlayerKickedMessage.ProtoReflect.Descriptor instead.
func (*PartyPlayerKickedMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{5}
}

func (x *PartyPlayerKickedMessage) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyPlayerKickedMessage) GetKickedMember() *party.PartyMember {
	if x != nil {
		return x.KickedMember
	}
	return nil
}

func (x *PartyPlayerKickedMessage) GetKickerMember() *party.PartyMember {
	if x != nil {
		return x.KickerMember
	}
	return nil
}

type PartyLeaderChangedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// party_id of type mongo ObjectId
	PartyId   string             `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	NewLeader *party.PartyMember `protobuf:"bytes,2,opt,name=new_leader,json=newLeader,proto3" json:"new_leader,omitempty"`
}

func (x *PartyLeaderChangedMessage) Reset() {
	*x = PartyLeaderChangedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyLeaderChangedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyLeaderChangedMessage) ProtoMessage() {}

func (x *PartyLeaderChangedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyLeaderChangedMessage.ProtoReflect.Descriptor instead.
func (*PartyLeaderChangedMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{6}
}

func (x *PartyLeaderChangedMessage) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyLeaderChangedMessage) GetNewLeader() *party.PartyMember {
	if x != nil {
		return x.NewLeader
	}
	return nil
}

type PartySettingsChangedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// player_id of type UUID
	PlayerId string               `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Settings *party.PartySettings `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *PartySettingsChangedMessage) Reset() {
	*x = PartySettingsChangedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartySettingsChangedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartySettingsChangedMessage) ProtoMessage() {}

func (x *PartySettingsChangedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_party_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartySettingsChangedMessage.ProtoReflect.Descriptor instead.
func (*PartySettingsChangedMessage) Descriptor() ([]byte, []int) {
	return file_party_messages_proto_rawDescGZIP(), []int{7}
}

func (x *PartySettingsChangedMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PartySettingsChangedMessage) GetSettings() *party.PartySettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

var File_party_messages_proto protoreflect.FileDescriptor

var file_party_messages_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x1a, 0x12, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x47, 0x0a, 0x13, 0x50, 0x61, 0x72, 0x74, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x79, 0x22, 0x49, 0x0a, 0x15, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x44, 0x69, 0x73, 0x62, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x05,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x22, 0x55, 0x0a, 0x19, 0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x52, 0x06, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x22, 0x6f, 0x0a, 0x18,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x65,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x6d, 0x0a,
	0x16, 0x50, 0x61, 0x72, 0x74, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x66, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x38, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xc3, 0x01, 0x0a,
	0x18, 0x50, 0x61, 0x72, 0x74, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4b, 0x69, 0x63, 0x6b,
	0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x0d, 0x6b, 0x69, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6d,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0c, 0x6b,
	0x69, 0x63, 0x6b, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x0d, 0x6b,
	0x69, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x0c, 0x6b, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x77, 0x0a, 0x19, 0x50, 0x61, 0x72, 0x74, 0x79, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0a, 0x6e, 0x65,
	0x77, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x09, 0x6e, 0x65, 0x77, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x7a, 0x0a, 0x1b, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6d, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x58, 0x0a, 0x1d, 0x64, 0x65, 0x76, 0x2e, 0x65,
	0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6d, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_party_messages_proto_rawDescOnce sync.Once
	file_party_messages_proto_rawDescData = file_party_messages_proto_rawDesc
)

func file_party_messages_proto_rawDescGZIP() []byte {
	file_party_messages_proto_rawDescOnce.Do(func() {
		file_party_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_party_messages_proto_rawDescData)
	})
	return file_party_messages_proto_rawDescData
}

var file_party_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_party_messages_proto_goTypes = []interface{}{
	(*PartyCreatedMessage)(nil),         // 0: emortal.message.party.PartyCreatedMessage
	(*PartyDisbandedMessage)(nil),       // 1: emortal.message.party.PartyDisbandedMessage
	(*PartyInviteCreatedMessage)(nil),   // 2: emortal.message.party.PartyInviteCreatedMessage
	(*PartyPlayerJoinedMessage)(nil),    // 3: emortal.message.party.PartyPlayerJoinedMessage
	(*PartyPlayerLeftMessage)(nil),      // 4: emortal.message.party.PartyPlayerLeftMessage
	(*PartyPlayerKickedMessage)(nil),    // 5: emortal.message.party.PartyPlayerKickedMessage
	(*PartyLeaderChangedMessage)(nil),   // 6: emortal.message.party.PartyLeaderChangedMessage
	(*PartySettingsChangedMessage)(nil), // 7: emortal.message.party.PartySettingsChangedMessage
	(*party.Party)(nil),                 // 8: emortal.model.party.Party
	(*party.PartyInvite)(nil),           // 9: emortal.model.party.PartyInvite
	(*party.PartyMember)(nil),           // 10: emortal.model.party.PartyMember
	(*party.PartySettings)(nil),         // 11: emortal.model.party.PartySettings
}
var file_party_messages_proto_depIdxs = []int32{
	8,  // 0: emortal.message.party.PartyCreatedMessage.party:type_name -> emortal.model.party.Party
	8,  // 1: emortal.message.party.PartyDisbandedMessage.party:type_name -> emortal.model.party.Party
	9,  // 2: emortal.message.party.PartyInviteCreatedMessage.invite:type_name -> emortal.model.party.PartyInvite
	10, // 3: emortal.message.party.PartyPlayerJoinedMessage.member:type_name -> emortal.model.party.PartyMember
	10, // 4: emortal.message.party.PartyPlayerLeftMessage.member:type_name -> emortal.model.party.PartyMember
	10, // 5: emortal.message.party.PartyPlayerKickedMessage.kicked_member:type_name -> emortal.model.party.PartyMember
	10, // 6: emortal.message.party.PartyPlayerKickedMessage.kicker_member:type_name -> emortal.model.party.PartyMember
	10, // 7: emortal.message.party.PartyLeaderChangedMessage.new_leader:type_name -> emortal.model.party.PartyMember
	11, // 8: emortal.message.party.PartySettingsChangedMessage.settings:type_name -> emortal.model.party.PartySettings
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_party_messages_proto_init() }
func file_party_messages_proto_init() {
	if File_party_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_party_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyCreatedMessage); i {
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
		file_party_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyDisbandedMessage); i {
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
		file_party_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyInviteCreatedMessage); i {
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
		file_party_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyPlayerJoinedMessage); i {
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
		file_party_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyPlayerLeftMessage); i {
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
		file_party_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyPlayerKickedMessage); i {
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
		file_party_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyLeaderChangedMessage); i {
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
		file_party_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartySettingsChangedMessage); i {
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
			RawDescriptor: file_party_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_party_messages_proto_goTypes,
		DependencyIndexes: file_party_messages_proto_depIdxs,
		MessageInfos:      file_party_messages_proto_msgTypes,
	}.Build()
	File_party_messages_proto = out.File
	file_party_messages_proto_rawDesc = nil
	file_party_messages_proto_goTypes = nil
	file_party_messages_proto_depIdxs = nil
}
