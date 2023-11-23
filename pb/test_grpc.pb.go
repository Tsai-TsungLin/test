// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: test.proto

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
	YourService_Echo_FullMethodName = "/test.YourService/Echo"
)

// YourServiceClient is the client API for YourService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YourServiceClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type yourServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYourServiceClient(cc grpc.ClientConnInterface) YourServiceClient {
	return &yourServiceClient{cc}
}

func (c *yourServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, YourService_Echo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YourServiceServer is the server API for YourService service.
// All implementations must embed UnimplementedYourServiceServer
// for forward compatibility
type YourServiceServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	mustEmbedUnimplementedYourServiceServer()
}

// UnimplementedYourServiceServer must be embedded to have forward compatible implementations.
type UnimplementedYourServiceServer struct {
}

func (UnimplementedYourServiceServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedYourServiceServer) mustEmbedUnimplementedYourServiceServer() {}

// UnsafeYourServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YourServiceServer will
// result in compilation errors.
type UnsafeYourServiceServer interface {
	mustEmbedUnimplementedYourServiceServer()
}

func RegisterYourServiceServer(s grpc.ServiceRegistrar, srv YourServiceServer) {
	s.RegisterService(&YourService_ServiceDesc, srv)
}

func _YourService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YourServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YourService_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YourServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// YourService_ServiceDesc is the grpc.ServiceDesc for YourService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YourService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.YourService",
	HandlerType: (*YourServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _YourService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}