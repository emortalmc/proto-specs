// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: badges/grpc.proto

package badge

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

// BadgeManagerClient is the client API for BadgeManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BadgeManagerClient interface {
	GetPlayerBadges(ctx context.Context, in *GetPlayerBadgesRequest, opts ...grpc.CallOption) (*GetPlayerBadgesResponse, error)
	GetActivePlayerBadge(ctx context.Context, in *GetActivePlayerBadgeRequest, opts ...grpc.CallOption) (*GetActivePlayerBadgeResponse, error)
	SetActivePlayerBadge(ctx context.Context, in *SetActivePlayerBadgeRequest, opts ...grpc.CallOption) (*SetActivePlayerBadgeResponse, error)
	AddBadgeToPlayer(ctx context.Context, in *AddBadgeToPlayerRequest, opts ...grpc.CallOption) (*AddBadgeToPlayerResponse, error)
	RemoveBadgeFromPlayer(ctx context.Context, in *RemoveBadgeFromPlayerRequest, opts ...grpc.CallOption) (*RemoveBadgeFromPlayerResponse, error)
	// GetBadges gets all the registered badges
	GetBadges(ctx context.Context, in *GetBadgesRequest, opts ...grpc.CallOption) (*GetBadgesResponse, error)
}

type badgeManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewBadgeManagerClient(cc grpc.ClientConnInterface) BadgeManagerClient {
	return &badgeManagerClient{cc}
}

func (c *badgeManagerClient) GetPlayerBadges(ctx context.Context, in *GetPlayerBadgesRequest, opts ...grpc.CallOption) (*GetPlayerBadgesResponse, error) {
	out := new(GetPlayerBadgesResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.badge.BadgeManager/GetPlayerBadges", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *badgeManagerClient) GetActivePlayerBadge(ctx context.Context, in *GetActivePlayerBadgeRequest, opts ...grpc.CallOption) (*GetActivePlayerBadgeResponse, error) {
	out := new(GetActivePlayerBadgeResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.badge.BadgeManager/GetActivePlayerBadge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *badgeManagerClient) SetActivePlayerBadge(ctx context.Context, in *SetActivePlayerBadgeRequest, opts ...grpc.CallOption) (*SetActivePlayerBadgeResponse, error) {
	out := new(SetActivePlayerBadgeResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.badge.BadgeManager/SetActivePlayerBadge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *badgeManagerClient) AddBadgeToPlayer(ctx context.Context, in *AddBadgeToPlayerRequest, opts ...grpc.CallOption) (*AddBadgeToPlayerResponse, error) {
	out := new(AddBadgeToPlayerResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.badge.BadgeManager/AddBadgeToPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *badgeManagerClient) RemoveBadgeFromPlayer(ctx context.Context, in *RemoveBadgeFromPlayerRequest, opts ...grpc.CallOption) (*RemoveBadgeFromPlayerResponse, error) {
	out := new(RemoveBadgeFromPlayerResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.badge.BadgeManager/RemoveBadgeFromPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *badgeManagerClient) GetBadges(ctx context.Context, in *GetBadgesRequest, opts ...grpc.CallOption) (*GetBadgesResponse, error) {
	out := new(GetBadgesResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.badge.BadgeManager/GetBadges", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BadgeManagerServer is the server API for BadgeManager service.
// All implementations must embed UnimplementedBadgeManagerServer
// for forward compatibility
type BadgeManagerServer interface {
	GetPlayerBadges(context.Context, *GetPlayerBadgesRequest) (*GetPlayerBadgesResponse, error)
	GetActivePlayerBadge(context.Context, *GetActivePlayerBadgeRequest) (*GetActivePlayerBadgeResponse, error)
	SetActivePlayerBadge(context.Context, *SetActivePlayerBadgeRequest) (*SetActivePlayerBadgeResponse, error)
	AddBadgeToPlayer(context.Context, *AddBadgeToPlayerRequest) (*AddBadgeToPlayerResponse, error)
	RemoveBadgeFromPlayer(context.Context, *RemoveBadgeFromPlayerRequest) (*RemoveBadgeFromPlayerResponse, error)
	// GetBadges gets all the registered badges
	GetBadges(context.Context, *GetBadgesRequest) (*GetBadgesResponse, error)
	mustEmbedUnimplementedBadgeManagerServer()
}

// UnimplementedBadgeManagerServer must be embedded to have forward compatible implementations.
type UnimplementedBadgeManagerServer struct {
}

func (UnimplementedBadgeManagerServer) GetPlayerBadges(context.Context, *GetPlayerBadgesRequest) (*GetPlayerBadgesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerBadges not implemented")
}
func (UnimplementedBadgeManagerServer) GetActivePlayerBadge(context.Context, *GetActivePlayerBadgeRequest) (*GetActivePlayerBadgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivePlayerBadge not implemented")
}
func (UnimplementedBadgeManagerServer) SetActivePlayerBadge(context.Context, *SetActivePlayerBadgeRequest) (*SetActivePlayerBadgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetActivePlayerBadge not implemented")
}
func (UnimplementedBadgeManagerServer) AddBadgeToPlayer(context.Context, *AddBadgeToPlayerRequest) (*AddBadgeToPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBadgeToPlayer not implemented")
}
func (UnimplementedBadgeManagerServer) RemoveBadgeFromPlayer(context.Context, *RemoveBadgeFromPlayerRequest) (*RemoveBadgeFromPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBadgeFromPlayer not implemented")
}
func (UnimplementedBadgeManagerServer) GetBadges(context.Context, *GetBadgesRequest) (*GetBadgesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBadges not implemented")
}
func (UnimplementedBadgeManagerServer) mustEmbedUnimplementedBadgeManagerServer() {}

// UnsafeBadgeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BadgeManagerServer will
// result in compilation errors.
type UnsafeBadgeManagerServer interface {
	mustEmbedUnimplementedBadgeManagerServer()
}

func RegisterBadgeManagerServer(s grpc.ServiceRegistrar, srv BadgeManagerServer) {
	s.RegisterService(&BadgeManager_ServiceDesc, srv)
}

func _BadgeManager_GetPlayerBadges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerBadgesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgeManagerServer).GetPlayerBadges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.badge.BadgeManager/GetPlayerBadges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgeManagerServer).GetPlayerBadges(ctx, req.(*GetPlayerBadgesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BadgeManager_GetActivePlayerBadge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivePlayerBadgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgeManagerServer).GetActivePlayerBadge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.badge.BadgeManager/GetActivePlayerBadge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgeManagerServer).GetActivePlayerBadge(ctx, req.(*GetActivePlayerBadgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BadgeManager_SetActivePlayerBadge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetActivePlayerBadgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgeManagerServer).SetActivePlayerBadge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.badge.BadgeManager/SetActivePlayerBadge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgeManagerServer).SetActivePlayerBadge(ctx, req.(*SetActivePlayerBadgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BadgeManager_AddBadgeToPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBadgeToPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgeManagerServer).AddBadgeToPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.badge.BadgeManager/AddBadgeToPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgeManagerServer).AddBadgeToPlayer(ctx, req.(*AddBadgeToPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BadgeManager_RemoveBadgeFromPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBadgeFromPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgeManagerServer).RemoveBadgeFromPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.badge.BadgeManager/RemoveBadgeFromPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgeManagerServer).RemoveBadgeFromPlayer(ctx, req.(*RemoveBadgeFromPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BadgeManager_GetBadges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBadgesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgeManagerServer).GetBadges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.badge.BadgeManager/GetBadges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgeManagerServer).GetBadges(ctx, req.(*GetBadgesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BadgeManager_ServiceDesc is the grpc.ServiceDesc for BadgeManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BadgeManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emortal.grpc.badge.BadgeManager",
	HandlerType: (*BadgeManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayerBadges",
			Handler:    _BadgeManager_GetPlayerBadges_Handler,
		},
		{
			MethodName: "GetActivePlayerBadge",
			Handler:    _BadgeManager_GetActivePlayerBadge_Handler,
		},
		{
			MethodName: "SetActivePlayerBadge",
			Handler:    _BadgeManager_SetActivePlayerBadge_Handler,
		},
		{
			MethodName: "AddBadgeToPlayer",
			Handler:    _BadgeManager_AddBadgeToPlayer_Handler,
		},
		{
			MethodName: "RemoveBadgeFromPlayer",
			Handler:    _BadgeManager_RemoveBadgeFromPlayer_Handler,
		},
		{
			MethodName: "GetBadges",
			Handler:    _BadgeManager_GetBadges_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "badges/grpc.proto",
}
