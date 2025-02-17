// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.24.4
// source: proto/internal_bff.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	InternalBffAPI_BulkCreateInventoryItems_FullMethodName = "/proto.InternalBffAPI/BulkCreateInventoryItems"
)

// InternalBffAPIClient is the client API for InternalBffAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InternalBffAPIClient interface {
	BulkCreateInventoryItems(ctx context.Context, opts ...grpc.CallOption) (InternalBffAPI_BulkCreateInventoryItemsClient, error)
}

type internalBffAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalBffAPIClient(cc grpc.ClientConnInterface) InternalBffAPIClient {
	return &internalBffAPIClient{cc}
}

func (c *internalBffAPIClient) BulkCreateInventoryItems(ctx context.Context, opts ...grpc.CallOption) (InternalBffAPI_BulkCreateInventoryItemsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &InternalBffAPI_ServiceDesc.Streams[0], InternalBffAPI_BulkCreateInventoryItems_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &internalBffAPIBulkCreateInventoryItemsClient{ClientStream: stream}
	return x, nil
}

type InternalBffAPI_BulkCreateInventoryItemsClient interface {
	Send(*NewInventoryItem) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type internalBffAPIBulkCreateInventoryItemsClient struct {
	grpc.ClientStream
}

func (x *internalBffAPIBulkCreateInventoryItemsClient) Send(m *NewInventoryItem) error {
	return x.ClientStream.SendMsg(m)
}

func (x *internalBffAPIBulkCreateInventoryItemsClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InternalBffAPIServer is the server API for InternalBffAPI service.
// All implementations should embed UnimplementedInternalBffAPIServer
// for forward compatibility
type InternalBffAPIServer interface {
	BulkCreateInventoryItems(InternalBffAPI_BulkCreateInventoryItemsServer) error
}

// UnimplementedInternalBffAPIServer should be embedded to have forward compatible implementations.
type UnimplementedInternalBffAPIServer struct {
}

func (UnimplementedInternalBffAPIServer) BulkCreateInventoryItems(InternalBffAPI_BulkCreateInventoryItemsServer) error {
	return status.Errorf(codes.Unimplemented, "method BulkCreateInventoryItems not implemented")
}

// UnsafeInternalBffAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InternalBffAPIServer will
// result in compilation errors.
type UnsafeInternalBffAPIServer interface {
	mustEmbedUnimplementedInternalBffAPIServer()
}

func RegisterInternalBffAPIServer(s grpc.ServiceRegistrar, srv InternalBffAPIServer) {
	s.RegisterService(&InternalBffAPI_ServiceDesc, srv)
}

func _InternalBffAPI_BulkCreateInventoryItems_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InternalBffAPIServer).BulkCreateInventoryItems(&internalBffAPIBulkCreateInventoryItemsServer{ServerStream: stream})
}

type InternalBffAPI_BulkCreateInventoryItemsServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*NewInventoryItem, error)
	grpc.ServerStream
}

type internalBffAPIBulkCreateInventoryItemsServer struct {
	grpc.ServerStream
}

func (x *internalBffAPIBulkCreateInventoryItemsServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *internalBffAPIBulkCreateInventoryItemsServer) Recv() (*NewInventoryItem, error) {
	m := new(NewInventoryItem)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InternalBffAPI_ServiceDesc is the grpc.ServiceDesc for InternalBffAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InternalBffAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.InternalBffAPI",
	HandlerType: (*InternalBffAPIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BulkCreateInventoryItems",
			Handler:       _InternalBffAPI_BulkCreateInventoryItems_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/internal_bff.proto",
}
