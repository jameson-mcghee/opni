// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/plugins/metrics/pkg/apis/remotewrite/remotewrite.proto

package remotewrite

import (
	context "context"
	cortexpb "github.com/cortexproject/cortex/pkg/cortexpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RemoteWrite_Push_FullMethodName      = "/remotewrite.RemoteWrite/Push"
	RemoteWrite_SyncRules_FullMethodName = "/remotewrite.RemoteWrite/SyncRules"
)

// RemoteWriteClient is the client API for RemoteWrite service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoteWriteClient interface {
	Push(ctx context.Context, in *cortexpb.WriteRequest, opts ...grpc.CallOption) (*cortexpb.WriteResponse, error)
	// rpc SyncRules(rules.RuleDesc) returns (google.protobuf.Empty);
	SyncRules(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type remoteWriteClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteWriteClient(cc grpc.ClientConnInterface) RemoteWriteClient {
	return &remoteWriteClient{cc}
}

func (c *remoteWriteClient) Push(ctx context.Context, in *cortexpb.WriteRequest, opts ...grpc.CallOption) (*cortexpb.WriteResponse, error) {
	out := new(cortexpb.WriteResponse)
	err := c.cc.Invoke(ctx, RemoteWrite_Push_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteWriteClient) SyncRules(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RemoteWrite_SyncRules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteWriteServer is the server API for RemoteWrite service.
// All implementations must embed UnimplementedRemoteWriteServer
// for forward compatibility
type RemoteWriteServer interface {
	Push(context.Context, *cortexpb.WriteRequest) (*cortexpb.WriteResponse, error)
	// rpc SyncRules(rules.RuleDesc) returns (google.protobuf.Empty);
	SyncRules(context.Context, *Payload) (*emptypb.Empty, error)
	mustEmbedUnimplementedRemoteWriteServer()
}

// UnimplementedRemoteWriteServer must be embedded to have forward compatible implementations.
type UnimplementedRemoteWriteServer struct {
}

func (UnimplementedRemoteWriteServer) Push(context.Context, *cortexpb.WriteRequest) (*cortexpb.WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedRemoteWriteServer) SyncRules(context.Context, *Payload) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncRules not implemented")
}
func (UnimplementedRemoteWriteServer) mustEmbedUnimplementedRemoteWriteServer() {}

// UnsafeRemoteWriteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteWriteServer will
// result in compilation errors.
type UnsafeRemoteWriteServer interface {
	mustEmbedUnimplementedRemoteWriteServer()
}

func RegisterRemoteWriteServer(s grpc.ServiceRegistrar, srv RemoteWriteServer) {
	s.RegisterService(&RemoteWrite_ServiceDesc, srv)
}

func _RemoteWrite_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cortexpb.WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteWriteServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteWrite_Push_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteWriteServer).Push(ctx, req.(*cortexpb.WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteWrite_SyncRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteWriteServer).SyncRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemoteWrite_SyncRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteWriteServer).SyncRules(ctx, req.(*Payload))
	}
	return interceptor(ctx, in, info, handler)
}

// RemoteWrite_ServiceDesc is the grpc.ServiceDesc for RemoteWrite service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemoteWrite_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "remotewrite.RemoteWrite",
	HandlerType: (*RemoteWriteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _RemoteWrite_Push_Handler,
		},
		{
			MethodName: "SyncRules",
			Handler:    _RemoteWrite_SyncRules_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/plugins/metrics/pkg/apis/remotewrite/remotewrite.proto",
}
