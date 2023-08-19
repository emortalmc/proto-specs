// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.1
// source: party/grpc.proto

package party

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

// PartyServiceClient is the client API for PartyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PartyServiceClient interface {
	// EmptyParty empties a party, removing all members and invites.
	EmptyParty(ctx context.Context, in *EmptyPartyRequest, opts ...grpc.CallOption) (*EmptyPartyResponse, error)
	GetParty(ctx context.Context, in *GetPartyRequest, opts ...grpc.CallOption) (*GetPartyResponse, error)
	SetOpenParty(ctx context.Context, in *SetOpenPartyRequest, opts ...grpc.CallOption) (*SetOpenPartyResponse, error)
	GetPartyInvites(ctx context.Context, in *GetPartyInvitesRequest, opts ...grpc.CallOption) (*GetPartyInvitesResponse, error)
	InvitePlayer(ctx context.Context, in *InvitePlayerRequest, opts ...grpc.CallOption) (*InvitePlayerResponse, error)
	JoinParty(ctx context.Context, in *JoinPartyRequest, opts ...grpc.CallOption) (*JoinPartyResponse, error)
	LeaveParty(ctx context.Context, in *LeavePartyRequest, opts ...grpc.CallOption) (*LeavePartyResponse, error)
	KickPlayer(ctx context.Context, in *KickPlayerRequest, opts ...grpc.CallOption) (*KickPlayerResponse, error)
	SetPartyLeader(ctx context.Context, in *SetPartyLeaderRequest, opts ...grpc.CallOption) (*SetPartyLeaderResponse, error)
}

type partyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartyServiceClient(cc grpc.ClientConnInterface) PartyServiceClient {
	return &partyServiceClient{cc}
}

func (c *partyServiceClient) EmptyParty(ctx context.Context, in *EmptyPartyRequest, opts ...grpc.CallOption) (*EmptyPartyResponse, error) {
	out := new(EmptyPartyResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.party.PartyService/EmptyParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) GetParty(ctx context.Context, in *GetPartyRequest, opts ...grpc.CallOption) (*GetPartyResponse, error) {
	out := new(GetPartyResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.party.PartyService/GetParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) SetOpenParty(ctx context.Context, in *SetOpenPartyRequest, opts ...grpc.CallOption) (*SetOpenPartyResponse, error) {
	out := new(SetOpenPartyResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.party.PartyService/SetOpenParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) GetPartyInvites(ctx context.Context, in *GetPartyInvitesRequest, opts ...grpc.CallOption) (*GetPartyInvitesResponse, error) {
	out := new(GetPartyInvitesResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.party.PartyService/GetPartyInvites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) InvitePlayer(ctx context.Context, in *InvitePlayerRequest, opts ...grpc.CallOption) (*InvitePlayerResponse, error) {
	out := new(InvitePlayerResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.party.PartyService/InvitePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) JoinParty(ctx context.Context, in *JoinPartyRequest, opts ...grpc.CallOption) (*JoinPartyResponse, error) {
	out := new(JoinPartyResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.party.PartyService/JoinParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) LeaveParty(ctx context.Context, in *LeavePartyRequest, opts ...grpc.CallOption) (*LeavePartyResponse, error) {
	out := new(LeavePartyResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.party.PartyService/LeaveParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) KickPlayer(ctx context.Context, in *KickPlayerRequest, opts ...grpc.CallOption) (*KickPlayerResponse, error) {
	out := new(KickPlayerResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.party.PartyService/KickPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) SetPartyLeader(ctx context.Context, in *SetPartyLeaderRequest, opts ...grpc.CallOption) (*SetPartyLeaderResponse, error) {
	out := new(SetPartyLeaderResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.party.PartyService/SetPartyLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartyServiceServer is the server API for PartyService service.
// All implementations must embed UnimplementedPartyServiceServer
// for forward compatibility
type PartyServiceServer interface {
	// EmptyParty empties a party, removing all members and invites.
	EmptyParty(context.Context, *EmptyPartyRequest) (*EmptyPartyResponse, error)
	GetParty(context.Context, *GetPartyRequest) (*GetPartyResponse, error)
	SetOpenParty(context.Context, *SetOpenPartyRequest) (*SetOpenPartyResponse, error)
	GetPartyInvites(context.Context, *GetPartyInvitesRequest) (*GetPartyInvitesResponse, error)
	InvitePlayer(context.Context, *InvitePlayerRequest) (*InvitePlayerResponse, error)
	JoinParty(context.Context, *JoinPartyRequest) (*JoinPartyResponse, error)
	LeaveParty(context.Context, *LeavePartyRequest) (*LeavePartyResponse, error)
	KickPlayer(context.Context, *KickPlayerRequest) (*KickPlayerResponse, error)
	SetPartyLeader(context.Context, *SetPartyLeaderRequest) (*SetPartyLeaderResponse, error)
	mustEmbedUnimplementedPartyServiceServer()
}

// UnimplementedPartyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPartyServiceServer struct {
}

func (UnimplementedPartyServiceServer) EmptyParty(context.Context, *EmptyPartyRequest) (*EmptyPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmptyParty not implemented")
}
func (UnimplementedPartyServiceServer) GetParty(context.Context, *GetPartyRequest) (*GetPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParty not implemented")
}
func (UnimplementedPartyServiceServer) SetOpenParty(context.Context, *SetOpenPartyRequest) (*SetOpenPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetOpenParty not implemented")
}
func (UnimplementedPartyServiceServer) GetPartyInvites(context.Context, *GetPartyInvitesRequest) (*GetPartyInvitesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartyInvites not implemented")
}
func (UnimplementedPartyServiceServer) InvitePlayer(context.Context, *InvitePlayerRequest) (*InvitePlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvitePlayer not implemented")
}
func (UnimplementedPartyServiceServer) JoinParty(context.Context, *JoinPartyRequest) (*JoinPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinParty not implemented")
}
func (UnimplementedPartyServiceServer) LeaveParty(context.Context, *LeavePartyRequest) (*LeavePartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveParty not implemented")
}
func (UnimplementedPartyServiceServer) KickPlayer(context.Context, *KickPlayerRequest) (*KickPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KickPlayer not implemented")
}
func (UnimplementedPartyServiceServer) SetPartyLeader(context.Context, *SetPartyLeaderRequest) (*SetPartyLeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPartyLeader not implemented")
}
func (UnimplementedPartyServiceServer) mustEmbedUnimplementedPartyServiceServer() {}

// UnsafePartyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PartyServiceServer will
// result in compilation errors.
type UnsafePartyServiceServer interface {
	mustEmbedUnimplementedPartyServiceServer()
}

func RegisterPartyServiceServer(s grpc.ServiceRegistrar, srv PartyServiceServer) {
	s.RegisterService(&PartyService_ServiceDesc, srv)
}

func _PartyService_EmptyParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).EmptyParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.party.PartyService/EmptyParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).EmptyParty(ctx, req.(*EmptyPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_GetParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).GetParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.party.PartyService/GetParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).GetParty(ctx, req.(*GetPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_SetOpenParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetOpenPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).SetOpenParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.party.PartyService/SetOpenParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).SetOpenParty(ctx, req.(*SetOpenPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_GetPartyInvites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartyInvitesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).GetPartyInvites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.party.PartyService/GetPartyInvites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).GetPartyInvites(ctx, req.(*GetPartyInvitesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_InvitePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvitePlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).InvitePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.party.PartyService/InvitePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).InvitePlayer(ctx, req.(*InvitePlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_JoinParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).JoinParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.party.PartyService/JoinParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).JoinParty(ctx, req.(*JoinPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_LeaveParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeavePartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).LeaveParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.party.PartyService/LeaveParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).LeaveParty(ctx, req.(*LeavePartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_KickPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KickPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).KickPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.party.PartyService/KickPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).KickPlayer(ctx, req.(*KickPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_SetPartyLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPartyLeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).SetPartyLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.party.PartyService/SetPartyLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).SetPartyLeader(ctx, req.(*SetPartyLeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PartyService_ServiceDesc is the grpc.ServiceDesc for PartyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PartyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emortal.grpc.party.PartyService",
	HandlerType: (*PartyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmptyParty",
			Handler:    _PartyService_EmptyParty_Handler,
		},
		{
			MethodName: "GetParty",
			Handler:    _PartyService_GetParty_Handler,
		},
		{
			MethodName: "SetOpenParty",
			Handler:    _PartyService_SetOpenParty_Handler,
		},
		{
			MethodName: "GetPartyInvites",
			Handler:    _PartyService_GetPartyInvites_Handler,
		},
		{
			MethodName: "InvitePlayer",
			Handler:    _PartyService_InvitePlayer_Handler,
		},
		{
			MethodName: "JoinParty",
			Handler:    _PartyService_JoinParty_Handler,
		},
		{
			MethodName: "LeaveParty",
			Handler:    _PartyService_LeaveParty_Handler,
		},
		{
			MethodName: "KickPlayer",
			Handler:    _PartyService_KickPlayer_Handler,
		},
		{
			MethodName: "SetPartyLeader",
			Handler:    _PartyService_SetPartyLeader_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "party/grpc.proto",
}

// PartySettingsServiceClient is the client API for PartySettingsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PartySettingsServiceClient interface {
	GetPartySettings(ctx context.Context, in *GetPartySettingsRequest, opts ...grpc.CallOption) (*GetPartySettingsResponse, error)
	UpdatePartySettings(ctx context.Context, in *UpdatePartySettingsRequest, opts ...grpc.CallOption) (*UpdatePartySettingsResponse, error)
}

type partySettingsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartySettingsServiceClient(cc grpc.ClientConnInterface) PartySettingsServiceClient {
	return &partySettingsServiceClient{cc}
}

func (c *partySettingsServiceClient) GetPartySettings(ctx context.Context, in *GetPartySettingsRequest, opts ...grpc.CallOption) (*GetPartySettingsResponse, error) {
	out := new(GetPartySettingsResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.party.PartySettingsService/GetPartySettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partySettingsServiceClient) UpdatePartySettings(ctx context.Context, in *UpdatePartySettingsRequest, opts ...grpc.CallOption) (*UpdatePartySettingsResponse, error) {
	out := new(UpdatePartySettingsResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.party.PartySettingsService/UpdatePartySettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartySettingsServiceServer is the server API for PartySettingsService service.
// All implementations must embed UnimplementedPartySettingsServiceServer
// for forward compatibility
type PartySettingsServiceServer interface {
	GetPartySettings(context.Context, *GetPartySettingsRequest) (*GetPartySettingsResponse, error)
	UpdatePartySettings(context.Context, *UpdatePartySettingsRequest) (*UpdatePartySettingsResponse, error)
	mustEmbedUnimplementedPartySettingsServiceServer()
}

// UnimplementedPartySettingsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPartySettingsServiceServer struct {
}

func (UnimplementedPartySettingsServiceServer) GetPartySettings(context.Context, *GetPartySettingsRequest) (*GetPartySettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartySettings not implemented")
}
func (UnimplementedPartySettingsServiceServer) UpdatePartySettings(context.Context, *UpdatePartySettingsRequest) (*UpdatePartySettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePartySettings not implemented")
}
func (UnimplementedPartySettingsServiceServer) mustEmbedUnimplementedPartySettingsServiceServer() {}

// UnsafePartySettingsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PartySettingsServiceServer will
// result in compilation errors.
type UnsafePartySettingsServiceServer interface {
	mustEmbedUnimplementedPartySettingsServiceServer()
}

func RegisterPartySettingsServiceServer(s grpc.ServiceRegistrar, srv PartySettingsServiceServer) {
	s.RegisterService(&PartySettingsService_ServiceDesc, srv)
}

func _PartySettingsService_GetPartySettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartySettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartySettingsServiceServer).GetPartySettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.party.PartySettingsService/GetPartySettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartySettingsServiceServer).GetPartySettings(ctx, req.(*GetPartySettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartySettingsService_UpdatePartySettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePartySettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartySettingsServiceServer).UpdatePartySettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.party.PartySettingsService/UpdatePartySettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartySettingsServiceServer).UpdatePartySettings(ctx, req.(*UpdatePartySettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PartySettingsService_ServiceDesc is the grpc.ServiceDesc for PartySettingsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PartySettingsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emortal.grpc.party.PartySettingsService",
	HandlerType: (*PartySettingsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPartySettings",
			Handler:    _PartySettingsService_GetPartySettings_Handler,
		},
		{
			MethodName: "UpdatePartySettings",
			Handler:    _PartySettingsService_UpdatePartySettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "party/grpc.proto",
}
