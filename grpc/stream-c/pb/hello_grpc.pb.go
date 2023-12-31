// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: pb/hello.proto

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
	Greeter_Login_FullMethodName      = "/pb.greeter/Login"
	Greeter_Hello_FullMethodName      = "/pb.greeter/Hello"
	Greeter_HandleHi_FullMethodName   = "/pb.greeter/HandleHi"
	Greeter_HandleChat_FullMethodName = "/pb.greeter/HandleChat"
)

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_HelloClient, error)
	HandleHi(ctx context.Context, opts ...grpc.CallOption) (Greeter_HandleHiClient, error)
	HandleChat(ctx context.Context, opts ...grpc.CallOption) (Greeter_HandleChatClient, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, Greeter_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_HelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[0], Greeter_Hello_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterHelloClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_HelloClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greeterHelloClient struct {
	grpc.ClientStream
}

func (x *greeterHelloClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) HandleHi(ctx context.Context, opts ...grpc.CallOption) (Greeter_HandleHiClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[1], Greeter_HandleHi_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterHandleHiClient{stream}
	return x, nil
}

type Greeter_HandleHiClient interface {
	Send(*HIRequest) error
	CloseAndRecv() (*HiResponse, error)
	grpc.ClientStream
}

type greeterHandleHiClient struct {
	grpc.ClientStream
}

func (x *greeterHandleHiClient) Send(m *HIRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterHandleHiClient) CloseAndRecv() (*HiResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HiResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) HandleChat(ctx context.Context, opts ...grpc.CallOption) (Greeter_HandleChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[2], Greeter_HandleChat_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterHandleChatClient{stream}
	return x, nil
}

type Greeter_HandleChatClient interface {
	Send(*ChatRequest) error
	Recv() (*ChatResponse, error)
	grpc.ClientStream
}

type greeterHandleChatClient struct {
	grpc.ClientStream
}

func (x *greeterHandleChatClient) Send(m *ChatRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterHandleChatClient) Recv() (*ChatResponse, error) {
	m := new(ChatResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Hello(*HelloRequest, Greeter_HelloServer) error
	HandleHi(Greeter_HandleHiServer) error
	HandleChat(Greeter_HandleChatServer) error
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedGreeterServer) Hello(*HelloRequest, Greeter_HelloServer) error {
	return status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedGreeterServer) HandleHi(Greeter_HandleHiServer) error {
	return status.Errorf(codes.Unimplemented, "method HandleHi not implemented")
}
func (UnimplementedGreeterServer) HandleChat(Greeter_HandleChatServer) error {
	return status.Errorf(codes.Unimplemented, "method HandleChat not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Hello_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).Hello(m, &greeterHelloServer{stream})
}

type Greeter_HelloServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type greeterHelloServer struct {
	grpc.ServerStream
}

func (x *greeterHelloServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_HandleHi_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).HandleHi(&greeterHandleHiServer{stream})
}

type Greeter_HandleHiServer interface {
	SendAndClose(*HiResponse) error
	Recv() (*HIRequest, error)
	grpc.ServerStream
}

type greeterHandleHiServer struct {
	grpc.ServerStream
}

func (x *greeterHandleHiServer) SendAndClose(m *HiResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterHandleHiServer) Recv() (*HIRequest, error) {
	m := new(HIRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_HandleChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).HandleChat(&greeterHandleChatServer{stream})
}

type Greeter_HandleChatServer interface {
	Send(*ChatResponse) error
	Recv() (*ChatRequest, error)
	grpc.ServerStream
}

type greeterHandleChatServer struct {
	grpc.ServerStream
}

func (x *greeterHandleChatServer) Send(m *ChatResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterHandleChatServer) Recv() (*ChatRequest, error) {
	m := new(ChatRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Greeter_Login_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Hello",
			Handler:       _Greeter_Hello_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HandleHi",
			Handler:       _Greeter_HandleHi_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "HandleChat",
			Handler:       _Greeter_HandleChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pb/hello.proto",
}
