// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PortDomainServiceClient is the client API for PortDomainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortDomainServiceClient interface {
	UpdatePorts(ctx context.Context, opts ...grpc.CallOption) (PortDomainService_UpdatePortsClient, error)
	GetPort(ctx context.Context, in *GetPortRequest, opts ...grpc.CallOption) (*GetPortResponse, error)
}

type portDomainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortDomainServiceClient(cc grpc.ClientConnInterface) PortDomainServiceClient {
	return &portDomainServiceClient{cc}
}

func (c *portDomainServiceClient) UpdatePorts(ctx context.Context, opts ...grpc.CallOption) (PortDomainService_UpdatePortsClient, error) {
	stream, err := c.cc.NewStream(ctx, &PortDomainService_ServiceDesc.Streams[0], "/port.PortDomainService/UpdatePorts", opts...)
	if err != nil {
		return nil, err
	}
	x := &portDomainServiceUpdatePortsClient{stream}
	return x, nil
}

type PortDomainService_UpdatePortsClient interface {
	Send(*PortWithID) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type portDomainServiceUpdatePortsClient struct {
	grpc.ClientStream
}

func (x *portDomainServiceUpdatePortsClient) Send(m *PortWithID) error {
	return x.ClientStream.SendMsg(m)
}

func (x *portDomainServiceUpdatePortsClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *portDomainServiceClient) GetPort(ctx context.Context, in *GetPortRequest, opts ...grpc.CallOption) (*GetPortResponse, error) {
	out := new(GetPortResponse)
	err := c.cc.Invoke(ctx, "/port.PortDomainService/GetPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortDomainServiceServer is the server API for PortDomainService service.
// All implementations must embed UnimplementedPortDomainServiceServer
// for forward compatibility
type PortDomainServiceServer interface {
	UpdatePorts(PortDomainService_UpdatePortsServer) error
	GetPort(context.Context, *GetPortRequest) (*GetPortResponse, error)
	mustEmbedUnimplementedPortDomainServiceServer()
}

// UnimplementedPortDomainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPortDomainServiceServer struct {
}

func (UnimplementedPortDomainServiceServer) UpdatePorts(PortDomainService_UpdatePortsServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdatePorts not implemented")
}
func (UnimplementedPortDomainServiceServer) GetPort(context.Context, *GetPortRequest) (*GetPortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPort not implemented")
}
func (UnimplementedPortDomainServiceServer) mustEmbedUnimplementedPortDomainServiceServer() {}

// UnsafePortDomainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortDomainServiceServer will
// result in compilation errors.
type UnsafePortDomainServiceServer interface {
	mustEmbedUnimplementedPortDomainServiceServer()
}

func RegisterPortDomainServiceServer(s grpc.ServiceRegistrar, srv PortDomainServiceServer) {
	s.RegisterService(&PortDomainService_ServiceDesc, srv)
}

func _PortDomainService_UpdatePorts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PortDomainServiceServer).UpdatePorts(&portDomainServiceUpdatePortsServer{stream})
}

type PortDomainService_UpdatePortsServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*PortWithID, error)
	grpc.ServerStream
}

type portDomainServiceUpdatePortsServer struct {
	grpc.ServerStream
}

func (x *portDomainServiceUpdatePortsServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *portDomainServiceUpdatePortsServer) Recv() (*PortWithID, error) {
	m := new(PortWithID)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PortDomainService_GetPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortDomainServiceServer).GetPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/port.PortDomainService/GetPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortDomainServiceServer).GetPort(ctx, req.(*GetPortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PortDomainService_ServiceDesc is the grpc.ServiceDesc for PortDomainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortDomainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "port.PortDomainService",
	HandlerType: (*PortDomainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPort",
			Handler:    _PortDomainService_GetPort_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdatePorts",
			Handler:       _PortDomainService_UpdatePorts_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "port.proto",
}
