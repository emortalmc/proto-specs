// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: message_handler/models.proto

package messagehandler

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PrivateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId       string `protobuf:"bytes,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	SenderUsername string `protobuf:"bytes,2,opt,name=sender_username,json=senderUsername,proto3" json:"sender_username,omitempty"`
	RecipientId    string `protobuf:"bytes,3,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`
	// The message may have been modified from the original for censoring purposes
	Message string                 `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	SentAt  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
}

func (x *PrivateMessage) Reset() {
	*x = PrivateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_handler_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateMessage) ProtoMessage() {}

func (x *PrivateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_handler_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateMessage.ProtoReflect.Descriptor instead.
func (*PrivateMessage) Descriptor() ([]byte, []int) {
	return file_message_handler_models_proto_rawDescGZIP(), []int{0}
}

func (x *PrivateMessage) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *PrivateMessage) GetSenderUsername() string {
	if x != nil {
		return x.SenderUsername
	}
	return ""
}

func (x *PrivateMessage) GetRecipientId() string {
	if x != nil {
		return x.RecipientId
	}
	return ""
}

func (x *PrivateMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PrivateMessage) GetSentAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SentAt
	}
	return nil
}

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId string `protobuf:"bytes,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_handler_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_handler_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_message_handler_models_proto_rawDescGZIP(), []int{1}
}

func (x *ChatMessage) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *ChatMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_message_handler_models_proto protoreflect.FileDescriptor

var file_message_handler_models_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c,
	0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01,
	0x0a, 0x0e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a,
	0x0f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x22, 0x44, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x66,
	0x0a, 0x24, 0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6d, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_handler_models_proto_rawDescOnce sync.Once
	file_message_handler_models_proto_rawDescData = file_message_handler_models_proto_rawDesc
)

func file_message_handler_models_proto_rawDescGZIP() []byte {
	file_message_handler_models_proto_rawDescOnce.Do(func() {
		file_message_handler_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_handler_models_proto_rawDescData)
	})
	return file_message_handler_models_proto_rawDescData
}

var file_message_handler_models_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_message_handler_models_proto_goTypes = []interface{}{
	(*PrivateMessage)(nil),        // 0: emortal.model.messagehandler.PrivateMessage
	(*ChatMessage)(nil),           // 1: emortal.model.messagehandler.ChatMessage
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_message_handler_models_proto_depIdxs = []int32{
	2, // 0: emortal.model.messagehandler.PrivateMessage.sent_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_message_handler_models_proto_init() }
func file_message_handler_models_proto_init() {
	if File_message_handler_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_handler_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateMessage); i {
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
		file_message_handler_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
			RawDescriptor: file_message_handler_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_handler_models_proto_goTypes,
		DependencyIndexes: file_message_handler_models_proto_depIdxs,
		MessageInfos:      file_message_handler_models_proto_msgTypes,
	}.Build()
	File_message_handler_models_proto = out.File
	file_message_handler_models_proto_rawDesc = nil
	file_message_handler_models_proto_goTypes = nil
	file_message_handler_models_proto_depIdxs = nil
}
