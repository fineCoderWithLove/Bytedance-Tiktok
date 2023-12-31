// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.14.0
// source: relation.proto

package rpb

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

const (
	UserService_UserRelationAction_FullMethodName       = "/rpb.UserService/UserRelationAction"
	UserService_UserRelationFollowList_FullMethodName   = "/rpb.UserService/UserRelationFollowList"
	UserService_UserRelationFollowerList_FullMethodName = "/rpb.UserService/UserRelationFollowerList"
	UserService_UserRelationFriendList_FullMethodName   = "/rpb.UserService/UserRelationFriendList"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	UserRelationAction(ctx context.Context, in *UserRelationActionReq, opts ...grpc.CallOption) (*UserRelationActionResp, error)
	UserRelationFollowList(ctx context.Context, in *UserRelationFollowListReq, opts ...grpc.CallOption) (*UserRelationFollowListResp, error)
	UserRelationFollowerList(ctx context.Context, in *UserRelationFollowerListReq, opts ...grpc.CallOption) (*UserRelationFollowerListResp, error)
	UserRelationFriendList(ctx context.Context, in *UserRelationFriendListReq, opts ...grpc.CallOption) (*UserRelationFriendListResp, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserRelationAction(ctx context.Context, in *UserRelationActionReq, opts ...grpc.CallOption) (*UserRelationActionResp, error) {
	out := new(UserRelationActionResp)
	err := c.cc.Invoke(ctx, UserService_UserRelationAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserRelationFollowList(ctx context.Context, in *UserRelationFollowListReq, opts ...grpc.CallOption) (*UserRelationFollowListResp, error) {
	out := new(UserRelationFollowListResp)
	err := c.cc.Invoke(ctx, UserService_UserRelationFollowList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserRelationFollowerList(ctx context.Context, in *UserRelationFollowerListReq, opts ...grpc.CallOption) (*UserRelationFollowerListResp, error) {
	out := new(UserRelationFollowerListResp)
	err := c.cc.Invoke(ctx, UserService_UserRelationFollowerList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserRelationFriendList(ctx context.Context, in *UserRelationFriendListReq, opts ...grpc.CallOption) (*UserRelationFriendListResp, error) {
	out := new(UserRelationFriendListResp)
	err := c.cc.Invoke(ctx, UserService_UserRelationFriendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	UserRelationAction(context.Context, *UserRelationActionReq) (*UserRelationActionResp, error)
	UserRelationFollowList(context.Context, *UserRelationFollowListReq) (*UserRelationFollowListResp, error)
	UserRelationFollowerList(context.Context, *UserRelationFollowerListReq) (*UserRelationFollowerListResp, error)
	UserRelationFriendList(context.Context, *UserRelationFriendListReq) (*UserRelationFriendListResp, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UserRelationAction(context.Context, *UserRelationActionReq) (*UserRelationActionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRelationAction not implemented")
}
func (UnimplementedUserServiceServer) UserRelationFollowList(context.Context, *UserRelationFollowListReq) (*UserRelationFollowListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRelationFollowList not implemented")
}
func (UnimplementedUserServiceServer) UserRelationFollowerList(context.Context, *UserRelationFollowerListReq) (*UserRelationFollowerListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRelationFollowerList not implemented")
}
func (UnimplementedUserServiceServer) UserRelationFriendList(context.Context, *UserRelationFriendListReq) (*UserRelationFriendListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRelationFriendList not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UserRelationAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRelationActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserRelationAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserRelationAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserRelationAction(ctx, req.(*UserRelationActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserRelationFollowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRelationFollowListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserRelationFollowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserRelationFollowList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserRelationFollowList(ctx, req.(*UserRelationFollowListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserRelationFollowerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRelationFollowerListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserRelationFollowerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserRelationFollowerList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserRelationFollowerList(ctx, req.(*UserRelationFollowerListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserRelationFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRelationFriendListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserRelationFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserRelationFriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserRelationFriendList(ctx, req.(*UserRelationFriendListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRelationAction",
			Handler:    _UserService_UserRelationAction_Handler,
		},
		{
			MethodName: "UserRelationFollowList",
			Handler:    _UserService_UserRelationFollowList_Handler,
		},
		{
			MethodName: "UserRelationFollowerList",
			Handler:    _UserService_UserRelationFollowerList_Handler,
		},
		{
			MethodName: "UserRelationFriendList",
			Handler:    _UserService_UserRelationFriendList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relation.proto",
}
