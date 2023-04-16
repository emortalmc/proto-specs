// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: message_handler/grpc.proto

package messagehandler

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessageHandlerClient is the client API for MessageHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageHandlerClient interface {
	SendPrivateMessage(ctx context.Context, in *PrivateMessageRequest, opts ...grpc.CallOption) (*PrivateMessageResponse, error)
}

type messageHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageHandlerClient(cc grpc.ClientConnInterface) MessageHandlerClient {
	return &messageHandlerClient{cc}
}

func (c *messageHandlerClient) SendPrivateMessage(ctx context.Context, in *PrivateMessageRequest, opts ...grpc.CallOption) (*PrivateMessageResponse, error) {
	out := new(PrivateMessageResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.messagehandler.MessageHandler/SendPrivateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageHandlerServer is the server API for MessageHandler service.
// All implementations must embed UnimplementedMessageHandlerServer
// for forward compatibility
type MessageHandlerServer interface {
	SendPrivateMessage(context.Context, *PrivateMessageRequest) (*PrivateMessageResponse, error)
	mustEmbedUnimplementedMessageHandlerServer()
}

// UnimplementedMessageHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedMessageHandlerServer struct {
}

func (UnimplementedMessageHandlerServer) SendPrivateMessage(context.Context, *PrivateMessageRequest) (*PrivateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPrivateMessage not implemented")
}
func (UnimplementedMessageHandlerServer) mustEmbedUnimplementedMessageHandlerServer() {}

// UnsafeMessageHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageHandlerServer will
// result in compilation errors.
type UnsafeMessageHandlerServer interface {
	mustEmbedUnimplementedMessageHandlerServer()
}

func RegisterMessageHandlerServer(s grpc.ServiceRegistrar, srv MessageHandlerServer) {
	s.RegisterService(&MessageHandler_ServiceDesc, srv)
}

func _MessageHandler_SendPrivateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageHandlerServer).SendPrivateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.messagehandler.MessageHandler/SendPrivateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageHandlerServer).SendPrivateMessage(ctx, req.(*PrivateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageHandler_ServiceDesc is the grpc.ServiceDesc for MessageHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emortal.grpc.messagehandler.MessageHandler",
	HandlerType: (*MessageHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPrivateMessage",
			Handler:    _MessageHandler_SendPrivateMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message_handler/grpc.proto",
}