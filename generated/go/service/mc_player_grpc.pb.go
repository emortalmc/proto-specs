// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: service/mc_player.proto

package service

import (
	context "context"
	model "github.com/towerdefence-cc/grpc-api-specs/generated/go/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// McPlayerClient is the client API for McPlayer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type McPlayerClient interface {
	GetPlayer(ctx context.Context, in *model.PlayerRequest, opts ...grpc.CallOption) (*PlayerResponse, error)
	GetPlayers(ctx context.Context, in *model.PlayersRequest, opts ...grpc.CallOption) (*PlayersResponse, error)
	GetPlayerByUsername(ctx context.Context, in *model.PlayerUsernameRequest, opts ...grpc.CallOption) (*PlayerResponse, error)
	SearchPlayerByUsername(ctx context.Context, in *PlayerSearchRequest, opts ...grpc.CallOption) (*PlayerSearchResponse, error)
	GetPlayerSessions(ctx context.Context, in *PageablePlayerRequest, opts ...grpc.CallOption) (*PlayerSessionsResponse, error)
	OnPlayerLogin(ctx context.Context, in *PlayerLoginRequest, opts ...grpc.CallOption) (*PlayerLoginResponse, error)
	OnPlayerDisconnect(ctx context.Context, in *PlayerDisconnectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type mcPlayerClient struct {
	cc grpc.ClientConnInterface
}

func NewMcPlayerClient(cc grpc.ClientConnInterface) McPlayerClient {
	return &mcPlayerClient{cc}
}

func (c *mcPlayerClient) GetPlayer(ctx context.Context, in *model.PlayerRequest, opts ...grpc.CallOption) (*PlayerResponse, error) {
	out := new(PlayerResponse)
	err := c.cc.Invoke(ctx, "/service.mc_player.McPlayer/GetPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mcPlayerClient) GetPlayers(ctx context.Context, in *model.PlayersRequest, opts ...grpc.CallOption) (*PlayersResponse, error) {
	out := new(PlayersResponse)
	err := c.cc.Invoke(ctx, "/service.mc_player.McPlayer/GetPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mcPlayerClient) GetPlayerByUsername(ctx context.Context, in *model.PlayerUsernameRequest, opts ...grpc.CallOption) (*PlayerResponse, error) {
	out := new(PlayerResponse)
	err := c.cc.Invoke(ctx, "/service.mc_player.McPlayer/GetPlayerByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mcPlayerClient) SearchPlayerByUsername(ctx context.Context, in *PlayerSearchRequest, opts ...grpc.CallOption) (*PlayerSearchResponse, error) {
	out := new(PlayerSearchResponse)
	err := c.cc.Invoke(ctx, "/service.mc_player.McPlayer/SearchPlayerByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mcPlayerClient) GetPlayerSessions(ctx context.Context, in *PageablePlayerRequest, opts ...grpc.CallOption) (*PlayerSessionsResponse, error) {
	out := new(PlayerSessionsResponse)
	err := c.cc.Invoke(ctx, "/service.mc_player.McPlayer/GetPlayerSessions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mcPlayerClient) OnPlayerLogin(ctx context.Context, in *PlayerLoginRequest, opts ...grpc.CallOption) (*PlayerLoginResponse, error) {
	out := new(PlayerLoginResponse)
	err := c.cc.Invoke(ctx, "/service.mc_player.McPlayer/OnPlayerLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mcPlayerClient) OnPlayerDisconnect(ctx context.Context, in *PlayerDisconnectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/service.mc_player.McPlayer/OnPlayerDisconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// McPlayerServer is the server API for McPlayer service.
// All implementations must embed UnimplementedMcPlayerServer
// for forward compatibility
type McPlayerServer interface {
	GetPlayer(context.Context, *model.PlayerRequest) (*PlayerResponse, error)
	GetPlayers(context.Context, *model.PlayersRequest) (*PlayersResponse, error)
	GetPlayerByUsername(context.Context, *model.PlayerUsernameRequest) (*PlayerResponse, error)
	SearchPlayerByUsername(context.Context, *PlayerSearchRequest) (*PlayerSearchResponse, error)
	GetPlayerSessions(context.Context, *PageablePlayerRequest) (*PlayerSessionsResponse, error)
	OnPlayerLogin(context.Context, *PlayerLoginRequest) (*PlayerLoginResponse, error)
	OnPlayerDisconnect(context.Context, *PlayerDisconnectRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMcPlayerServer()
}

// UnimplementedMcPlayerServer must be embedded to have forward compatible implementations.
type UnimplementedMcPlayerServer struct {
}

func (UnimplementedMcPlayerServer) GetPlayer(context.Context, *model.PlayerRequest) (*PlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayer not implemented")
}
func (UnimplementedMcPlayerServer) GetPlayers(context.Context, *model.PlayersRequest) (*PlayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayers not implemented")
}
func (UnimplementedMcPlayerServer) GetPlayerByUsername(context.Context, *model.PlayerUsernameRequest) (*PlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerByUsername not implemented")
}
func (UnimplementedMcPlayerServer) SearchPlayerByUsername(context.Context, *PlayerSearchRequest) (*PlayerSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPlayerByUsername not implemented")
}
func (UnimplementedMcPlayerServer) GetPlayerSessions(context.Context, *PageablePlayerRequest) (*PlayerSessionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerSessions not implemented")
}
func (UnimplementedMcPlayerServer) OnPlayerLogin(context.Context, *PlayerLoginRequest) (*PlayerLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnPlayerLogin not implemented")
}
func (UnimplementedMcPlayerServer) OnPlayerDisconnect(context.Context, *PlayerDisconnectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnPlayerDisconnect not implemented")
}
func (UnimplementedMcPlayerServer) mustEmbedUnimplementedMcPlayerServer() {}

// UnsafeMcPlayerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to McPlayerServer will
// result in compilation errors.
type UnsafeMcPlayerServer interface {
	mustEmbedUnimplementedMcPlayerServer()
}

func RegisterMcPlayerServer(s grpc.ServiceRegistrar, srv McPlayerServer) {
	s.RegisterService(&McPlayer_ServiceDesc, srv)
}

func _McPlayer_GetPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.PlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(McPlayerServer).GetPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.mc_player.McPlayer/GetPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(McPlayerServer).GetPlayer(ctx, req.(*model.PlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _McPlayer_GetPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.PlayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(McPlayerServer).GetPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.mc_player.McPlayer/GetPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(McPlayerServer).GetPlayers(ctx, req.(*model.PlayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _McPlayer_GetPlayerByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.PlayerUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(McPlayerServer).GetPlayerByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.mc_player.McPlayer/GetPlayerByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(McPlayerServer).GetPlayerByUsername(ctx, req.(*model.PlayerUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _McPlayer_SearchPlayerByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(McPlayerServer).SearchPlayerByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.mc_player.McPlayer/SearchPlayerByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(McPlayerServer).SearchPlayerByUsername(ctx, req.(*PlayerSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _McPlayer_GetPlayerSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageablePlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(McPlayerServer).GetPlayerSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.mc_player.McPlayer/GetPlayerSessions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(McPlayerServer).GetPlayerSessions(ctx, req.(*PageablePlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _McPlayer_OnPlayerLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(McPlayerServer).OnPlayerLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.mc_player.McPlayer/OnPlayerLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(McPlayerServer).OnPlayerLogin(ctx, req.(*PlayerLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _McPlayer_OnPlayerDisconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerDisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(McPlayerServer).OnPlayerDisconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.mc_player.McPlayer/OnPlayerDisconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(McPlayerServer).OnPlayerDisconnect(ctx, req.(*PlayerDisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// McPlayer_ServiceDesc is the grpc.ServiceDesc for McPlayer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var McPlayer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.mc_player.McPlayer",
	HandlerType: (*McPlayerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayer",
			Handler:    _McPlayer_GetPlayer_Handler,
		},
		{
			MethodName: "GetPlayers",
			Handler:    _McPlayer_GetPlayers_Handler,
		},
		{
			MethodName: "GetPlayerByUsername",
			Handler:    _McPlayer_GetPlayerByUsername_Handler,
		},
		{
			MethodName: "SearchPlayerByUsername",
			Handler:    _McPlayer_SearchPlayerByUsername_Handler,
		},
		{
			MethodName: "GetPlayerSessions",
			Handler:    _McPlayer_GetPlayerSessions_Handler,
		},
		{
			MethodName: "OnPlayerLogin",
			Handler:    _McPlayer_OnPlayerLogin_Handler,
		},
		{
			MethodName: "OnPlayerDisconnect",
			Handler:    _McPlayer_OnPlayerDisconnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/mc_player.proto",
}