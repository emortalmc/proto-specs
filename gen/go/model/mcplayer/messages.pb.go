// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: mc_player/messages.proto

package mcplayer

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

type PlayerLevelUpMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// player_id of type UUID
	PlayerId      string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	PreviousLevel int32  `protobuf:"varint,2,opt,name=previous_level,json=previousLevel,proto3" json:"previous_level,omitempty"`
	NewLevel      int32  `protobuf:"varint,3,opt,name=new_level,json=newLevel,proto3" json:"new_level,omitempty"`
	// experience is the XP they are into the current level - NOT total
	Experience int32 `protobuf:"varint,4,opt,name=experience,proto3" json:"experience,omitempty"`
}

func (x *PlayerLevelUpMessage) Reset() {
	*x = PlayerLevelUpMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mc_player_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerLevelUpMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerLevelUpMessage) ProtoMessage() {}

func (x *PlayerLevelUpMessage) ProtoReflect() protoreflect.Message {
	mi := &file_mc_player_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerLevelUpMessage.ProtoReflect.Descriptor instead.
func (*PlayerLevelUpMessage) Descriptor() ([]byte, []int) {
	return file_mc_player_messages_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerLevelUpMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerLevelUpMessage) GetPreviousLevel() int32 {
	if x != nil {
		return x.PreviousLevel
	}
	return 0
}

func (x *PlayerLevelUpMessage) GetNewLevel() int32 {
	if x != nil {
		return x.NewLevel
	}
	return 0
}

func (x *PlayerLevelUpMessage) GetExperience() int32 {
	if x != nil {
		return x.Experience
	}
	return 0
}

type PlayerExperienceMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// player_id of type UUID
	PlayerId           string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	PreviousExperience int32  `protobuf:"varint,2,opt,name=previous_experience,json=previousExperience,proto3" json:"previous_experience,omitempty"`
	NewExperience      int32  `protobuf:"varint,3,opt,name=new_experience,json=newExperience,proto3" json:"new_experience,omitempty"`
}

func (x *PlayerExperienceMessage) Reset() {
	*x = PlayerExperienceMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mc_player_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerExperienceMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerExperienceMessage) ProtoMessage() {}

func (x *PlayerExperienceMessage) ProtoReflect() protoreflect.Message {
	mi := &file_mc_player_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerExperienceMessage.ProtoReflect.Descriptor instead.
func (*PlayerExperienceMessage) Descriptor() ([]byte, []int) {
	return file_mc_player_messages_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerExperienceMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerExperienceMessage) GetPreviousExperience() int32 {
	if x != nil {
		return x.PreviousExperience
	}
	return 0
}

func (x *PlayerExperienceMessage) GetNewExperience() int32 {
	if x != nil {
		return x.NewExperience
	}
	return 0
}

var File_mc_player_messages_proto protoreflect.FileDescriptor

var file_mc_player_messages_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x65, 0x6d, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6d, 0x63, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x22, 0x97, 0x01, 0x0a, 0x14, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x55, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x6f, 0x75, 0x73, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x8e, 0x01, 0x0a,
	0x17, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75,
	0x73, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x12, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x45, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x5f, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x6e, 0x65, 0x77, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x5a, 0x0a,
	0x1e, 0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6d, 0x63, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50,
	0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6d, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70,
	0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x6d, 0x63, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mc_player_messages_proto_rawDescOnce sync.Once
	file_mc_player_messages_proto_rawDescData = file_mc_player_messages_proto_rawDesc
)

func file_mc_player_messages_proto_rawDescGZIP() []byte {
	file_mc_player_messages_proto_rawDescOnce.Do(func() {
		file_mc_player_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_mc_player_messages_proto_rawDescData)
	})
	return file_mc_player_messages_proto_rawDescData
}

var file_mc_player_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mc_player_messages_proto_goTypes = []interface{}{
	(*PlayerLevelUpMessage)(nil),    // 0: emortal.model.mcplayer.PlayerLevelUpMessage
	(*PlayerExperienceMessage)(nil), // 1: emortal.model.mcplayer.PlayerExperienceMessage
}
var file_mc_player_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mc_player_messages_proto_init() }
func file_mc_player_messages_proto_init() {
	if File_mc_player_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mc_player_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerLevelUpMessage); i {
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
		file_mc_player_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerExperienceMessage); i {
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
			RawDescriptor: file_mc_player_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mc_player_messages_proto_goTypes,
		DependencyIndexes: file_mc_player_messages_proto_depIdxs,
		MessageInfos:      file_mc_player_messages_proto_msgTypes,
	}.Build()
	File_mc_player_messages_proto = out.File
	file_mc_player_messages_proto_rawDesc = nil
	file_mc_player_messages_proto_goTypes = nil
	file_mc_player_messages_proto_depIdxs = nil
}
