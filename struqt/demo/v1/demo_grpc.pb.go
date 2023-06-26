// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: struqt/demo/v1/demo.proto

package demo

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
	BasicService_JsonRpc_FullMethodName = "/struqt.demo.v1.BasicService/JsonRpc"
)

// BasicServiceClient is the client API for BasicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BasicServiceClient interface {
	JsonRpc(ctx context.Context, in *JsonRpcRequest, opts ...grpc.CallOption) (*JsonRpcResponse, error)
}

type basicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBasicServiceClient(cc grpc.ClientConnInterface) BasicServiceClient {
	return &basicServiceClient{cc}
}

func (c *basicServiceClient) JsonRpc(ctx context.Context, in *JsonRpcRequest, opts ...grpc.CallOption) (*JsonRpcResponse, error) {
	out := new(JsonRpcResponse)
	err := c.cc.Invoke(ctx, BasicService_JsonRpc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasicServiceServer is the server API for BasicService service.
// All implementations must embed UnimplementedBasicServiceServer
// for forward compatibility
type BasicServiceServer interface {
	JsonRpc(context.Context, *JsonRpcRequest) (*JsonRpcResponse, error)
	mustEmbedUnimplementedBasicServiceServer()
}

// UnimplementedBasicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBasicServiceServer struct {
}

func (UnimplementedBasicServiceServer) JsonRpc(context.Context, *JsonRpcRequest) (*JsonRpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JsonRpc not implemented")
}
func (UnimplementedBasicServiceServer) mustEmbedUnimplementedBasicServiceServer() {}

// UnsafeBasicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BasicServiceServer will
// result in compilation errors.
type UnsafeBasicServiceServer interface {
	mustEmbedUnimplementedBasicServiceServer()
}

func RegisterBasicServiceServer(s grpc.ServiceRegistrar, srv BasicServiceServer) {
	s.RegisterService(&BasicService_ServiceDesc, srv)
}

func _BasicService_JsonRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JsonRpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicServiceServer).JsonRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BasicService_JsonRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicServiceServer).JsonRpc(ctx, req.(*JsonRpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BasicService_ServiceDesc is the grpc.ServiceDesc for BasicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BasicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "struqt.demo.v1.BasicService",
	HandlerType: (*BasicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JsonRpc",
			Handler:    _BasicService_JsonRpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "struqt/demo/v1/demo.proto",
}

const (
	DemoService_Echo_FullMethodName = "/struqt.demo.v1.DemoService/Echo"
)

// DemoServiceClient is the client API for DemoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DemoServiceClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type demoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDemoServiceClient(cc grpc.ClientConnInterface) DemoServiceClient {
	return &demoServiceClient{cc}
}

func (c *demoServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, DemoService_Echo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoServiceServer is the server API for DemoService service.
// All implementations must embed UnimplementedDemoServiceServer
// for forward compatibility
type DemoServiceServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	mustEmbedUnimplementedDemoServiceServer()
}

// UnimplementedDemoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDemoServiceServer struct {
}

func (UnimplementedDemoServiceServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedDemoServiceServer) mustEmbedUnimplementedDemoServiceServer() {}

// UnsafeDemoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DemoServiceServer will
// result in compilation errors.
type UnsafeDemoServiceServer interface {
	mustEmbedUnimplementedDemoServiceServer()
}

func RegisterDemoServiceServer(s grpc.ServiceRegistrar, srv DemoServiceServer) {
	s.RegisterService(&DemoService_ServiceDesc, srv)
}

func _DemoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DemoService_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DemoService_ServiceDesc is the grpc.ServiceDesc for DemoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DemoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "struqt.demo.v1.DemoService",
	HandlerType: (*DemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _DemoService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "struqt/demo/v1/demo.proto",
}
