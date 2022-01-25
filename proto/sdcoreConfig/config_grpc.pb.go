// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: config.proto

package sdcoreConfig

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

// ConfigServiceClient is the client API for ConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigServiceClient interface {
	GetNetworkSlice(ctx context.Context, in *NetworkSliceRequest, opts ...grpc.CallOption) (*NetworkSliceResponse, error)
	NetworkSliceSubscribe(ctx context.Context, in *NetworkSliceRequest, opts ...grpc.CallOption) (ConfigService_NetworkSliceSubscribeClient, error)
}

type configServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigServiceClient(cc grpc.ClientConnInterface) ConfigServiceClient {
	return &configServiceClient{cc}
}

func (c *configServiceClient) GetNetworkSlice(ctx context.Context, in *NetworkSliceRequest, opts ...grpc.CallOption) (*NetworkSliceResponse, error) {
	out := new(NetworkSliceResponse)
	err := c.cc.Invoke(ctx, "/sdcoreConfig.ConfigService/GetNetworkSlice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) NetworkSliceSubscribe(ctx context.Context, in *NetworkSliceRequest, opts ...grpc.CallOption) (ConfigService_NetworkSliceSubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConfigService_ServiceDesc.Streams[0], "/sdcoreConfig.ConfigService/NetworkSliceSubscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &configServiceNetworkSliceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ConfigService_NetworkSliceSubscribeClient interface {
	Recv() (*NetworkSliceResponse, error)
	grpc.ClientStream
}

type configServiceNetworkSliceSubscribeClient struct {
	grpc.ClientStream
}

func (x *configServiceNetworkSliceSubscribeClient) Recv() (*NetworkSliceResponse, error) {
	m := new(NetworkSliceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConfigServiceServer is the server API for ConfigService service.
// All implementations must embed UnimplementedConfigServiceServer
// for forward compatibility
type ConfigServiceServer interface {
	GetNetworkSlice(context.Context, *NetworkSliceRequest) (*NetworkSliceResponse, error)
	NetworkSliceSubscribe(*NetworkSliceRequest, ConfigService_NetworkSliceSubscribeServer) error
	mustEmbedUnimplementedConfigServiceServer()
}

// UnimplementedConfigServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConfigServiceServer struct {
}

func (UnimplementedConfigServiceServer) GetNetworkSlice(context.Context, *NetworkSliceRequest) (*NetworkSliceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkSlice not implemented")
}
func (UnimplementedConfigServiceServer) NetworkSliceSubscribe(*NetworkSliceRequest, ConfigService_NetworkSliceSubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method NetworkSliceSubscribe not implemented")
}
func (UnimplementedConfigServiceServer) mustEmbedUnimplementedConfigServiceServer() {}

// UnsafeConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigServiceServer will
// result in compilation errors.
type UnsafeConfigServiceServer interface {
	mustEmbedUnimplementedConfigServiceServer()
}

func RegisterConfigServiceServer(s grpc.ServiceRegistrar, srv ConfigServiceServer) {
	s.RegisterService(&ConfigService_ServiceDesc, srv)
}

func _ConfigService_GetNetworkSlice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkSliceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetNetworkSlice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdcoreConfig.ConfigService/GetNetworkSlice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetNetworkSlice(ctx, req.(*NetworkSliceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_NetworkSliceSubscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NetworkSliceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConfigServiceServer).NetworkSliceSubscribe(m, &configServiceNetworkSliceSubscribeServer{stream})
}

type ConfigService_NetworkSliceSubscribeServer interface {
	Send(*NetworkSliceResponse) error
	grpc.ServerStream
}

type configServiceNetworkSliceSubscribeServer struct {
	grpc.ServerStream
}

func (x *configServiceNetworkSliceSubscribeServer) Send(m *NetworkSliceResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ConfigService_ServiceDesc is the grpc.ServiceDesc for ConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sdcoreConfig.ConfigService",
	HandlerType: (*ConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNetworkSlice",
			Handler:    _ConfigService_GetNetworkSlice_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NetworkSliceSubscribe",
			Handler:       _ConfigService_NetworkSliceSubscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "config.proto",
}
