// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: pb/what.proto

package pb

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
	What_UserSinup_FullMethodName = "/pb.What/UserSinup"
	What_UserLogin_FullMethodName = "/pb.What/UserLogin"
	What_GetBook_FullMethodName   = "/pb.What/GetBook"
)

// WhatClient is the client API for What service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WhatClient interface {
	UserSinup(ctx context.Context, in *SinupRequest, opts ...grpc.CallOption) (*SinupResponse, error)
	UserLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error)
}

type whatClient struct {
	cc grpc.ClientConnInterface
}

func NewWhatClient(cc grpc.ClientConnInterface) WhatClient {
	return &whatClient{cc}
}

func (c *whatClient) UserSinup(ctx context.Context, in *SinupRequest, opts ...grpc.CallOption) (*SinupResponse, error) {
	out := new(SinupResponse)
	err := c.cc.Invoke(ctx, What_UserSinup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *whatClient) UserLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, What_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *whatClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error) {
	out := new(GetBookResponse)
	err := c.cc.Invoke(ctx, What_GetBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WhatServer is the server API for What service.
// All implementations must embed UnimplementedWhatServer
// for forward compatibility
type WhatServer interface {
	UserSinup(context.Context, *SinupRequest) (*SinupResponse, error)
	UserLogin(context.Context, *LoginRequest) (*LoginResponse, error)
	GetBook(context.Context, *GetBookRequest) (*GetBookResponse, error)
	mustEmbedUnimplementedWhatServer()
}

// UnimplementedWhatServer must be embedded to have forward compatible implementations.
type UnimplementedWhatServer struct {
}

func (UnimplementedWhatServer) UserSinup(context.Context, *SinupRequest) (*SinupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSinup not implemented")
}
func (UnimplementedWhatServer) UserLogin(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedWhatServer) GetBook(context.Context, *GetBookRequest) (*GetBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedWhatServer) mustEmbedUnimplementedWhatServer() {}

// UnsafeWhatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WhatServer will
// result in compilation errors.
type UnsafeWhatServer interface {
	mustEmbedUnimplementedWhatServer()
}

func RegisterWhatServer(s grpc.ServiceRegistrar, srv WhatServer) {
	s.RegisterService(&What_ServiceDesc, srv)
}

func _What_UserSinup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SinupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhatServer).UserSinup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: What_UserSinup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhatServer).UserSinup(ctx, req.(*SinupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _What_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhatServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: What_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhatServer).UserLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _What_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhatServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: What_GetBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhatServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// What_ServiceDesc is the grpc.ServiceDesc for What service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var What_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.What",
	HandlerType: (*WhatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserSinup",
			Handler:    _What_UserSinup_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _What_UserLogin_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _What_GetBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/what.proto",
}
