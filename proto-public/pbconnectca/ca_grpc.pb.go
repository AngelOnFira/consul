// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto-public/pbconnectca/ca.proto

package pbconnectca

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

// ConnectCAServiceClient is the client API for ConnectCAService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectCAServiceClient interface {
	// WatchRoots provides a stream on which you can receive the list of active
	// Connect CA roots. Current roots are sent immediately at the start of the
	// stream, and new lists will be sent whenever the roots are rotated.
	WatchRoots(ctx context.Context, in *WatchRootsRequest, opts ...grpc.CallOption) (ConnectCAService_WatchRootsClient, error)
	// Sign a leaf certificate for the service or agent identified by the SPIFFE
	// ID in the given CSR's SAN.
	Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error)
}

type connectCAServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectCAServiceClient(cc grpc.ClientConnInterface) ConnectCAServiceClient {
	return &connectCAServiceClient{cc}
}

func (c *connectCAServiceClient) WatchRoots(ctx context.Context, in *WatchRootsRequest, opts ...grpc.CallOption) (ConnectCAService_WatchRootsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConnectCAService_ServiceDesc.Streams[0], "/hashicorp.consul.connectca.ConnectCAService/WatchRoots", opts...)
	if err != nil {
		return nil, err
	}
	x := &connectCAServiceWatchRootsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ConnectCAService_WatchRootsClient interface {
	Recv() (*WatchRootsResponse, error)
	grpc.ClientStream
}

type connectCAServiceWatchRootsClient struct {
	grpc.ClientStream
}

func (x *connectCAServiceWatchRootsClient) Recv() (*WatchRootsResponse, error) {
	m := new(WatchRootsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *connectCAServiceClient) Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error) {
	out := new(SignResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.consul.connectca.ConnectCAService/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectCAServiceServer is the server API for ConnectCAService service.
// All implementations should embed UnimplementedConnectCAServiceServer
// for forward compatibility
type ConnectCAServiceServer interface {
	// WatchRoots provides a stream on which you can receive the list of active
	// Connect CA roots. Current roots are sent immediately at the start of the
	// stream, and new lists will be sent whenever the roots are rotated.
	WatchRoots(*WatchRootsRequest, ConnectCAService_WatchRootsServer) error
	// Sign a leaf certificate for the service or agent identified by the SPIFFE
	// ID in the given CSR's SAN.
	Sign(context.Context, *SignRequest) (*SignResponse, error)
}

// UnimplementedConnectCAServiceServer should be embedded to have forward compatible implementations.
type UnimplementedConnectCAServiceServer struct {
}

func (UnimplementedConnectCAServiceServer) WatchRoots(*WatchRootsRequest, ConnectCAService_WatchRootsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchRoots not implemented")
}
func (UnimplementedConnectCAServiceServer) Sign(context.Context, *SignRequest) (*SignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}

// UnsafeConnectCAServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectCAServiceServer will
// result in compilation errors.
type UnsafeConnectCAServiceServer interface {
	mustEmbedUnimplementedConnectCAServiceServer()
}

func RegisterConnectCAServiceServer(s grpc.ServiceRegistrar, srv ConnectCAServiceServer) {
	s.RegisterService(&ConnectCAService_ServiceDesc, srv)
}

func _ConnectCAService_WatchRoots_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRootsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConnectCAServiceServer).WatchRoots(m, &connectCAServiceWatchRootsServer{stream})
}

type ConnectCAService_WatchRootsServer interface {
	Send(*WatchRootsResponse) error
	grpc.ServerStream
}

type connectCAServiceWatchRootsServer struct {
	grpc.ServerStream
}

func (x *connectCAServiceWatchRootsServer) Send(m *WatchRootsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ConnectCAService_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectCAServiceServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.consul.connectca.ConnectCAService/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectCAServiceServer).Sign(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectCAService_ServiceDesc is the grpc.ServiceDesc for ConnectCAService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectCAService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.consul.connectca.ConnectCAService",
	HandlerType: (*ConnectCAServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sign",
			Handler:    _ConnectCAService_Sign_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchRoots",
			Handler:       _ConnectCAService_WatchRoots_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto-public/pbconnectca/ca.proto",
}