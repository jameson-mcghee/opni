// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/plugins/metrics/pkg/apis/cortexops/cortexops.proto

package cortexops

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

// CortexOpsClient is the client API for CortexOps service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CortexOpsClient interface {
	GetClusterConfiguration(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ClusterConfiguration, error)
	ConfigureCluster(ctx context.Context, in *ClusterConfiguration, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetClusterStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InstallStatus, error)
	UninstallCluster(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type cortexOpsClient struct {
	cc grpc.ClientConnInterface
}

func NewCortexOpsClient(cc grpc.ClientConnInterface) CortexOpsClient {
	return &cortexOpsClient{cc}
}

func (c *cortexOpsClient) GetClusterConfiguration(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ClusterConfiguration, error) {
	out := new(ClusterConfiguration)
	err := c.cc.Invoke(ctx, "/cortexops.CortexOps/GetClusterConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cortexOpsClient) ConfigureCluster(ctx context.Context, in *ClusterConfiguration, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cortexops.CortexOps/ConfigureCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cortexOpsClient) GetClusterStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InstallStatus, error) {
	out := new(InstallStatus)
	err := c.cc.Invoke(ctx, "/cortexops.CortexOps/GetClusterStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cortexOpsClient) UninstallCluster(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cortexops.CortexOps/UninstallCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CortexOpsServer is the server API for CortexOps service.
// All implementations must embed UnimplementedCortexOpsServer
// for forward compatibility
type CortexOpsServer interface {
	GetClusterConfiguration(context.Context, *emptypb.Empty) (*ClusterConfiguration, error)
	ConfigureCluster(context.Context, *ClusterConfiguration) (*emptypb.Empty, error)
	GetClusterStatus(context.Context, *emptypb.Empty) (*InstallStatus, error)
	UninstallCluster(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedCortexOpsServer()
}

// UnimplementedCortexOpsServer must be embedded to have forward compatible implementations.
type UnimplementedCortexOpsServer struct {
}

func (UnimplementedCortexOpsServer) GetClusterConfiguration(context.Context, *emptypb.Empty) (*ClusterConfiguration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterConfiguration not implemented")
}
func (UnimplementedCortexOpsServer) ConfigureCluster(context.Context, *ClusterConfiguration) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureCluster not implemented")
}
func (UnimplementedCortexOpsServer) GetClusterStatus(context.Context, *emptypb.Empty) (*InstallStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterStatus not implemented")
}
func (UnimplementedCortexOpsServer) UninstallCluster(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UninstallCluster not implemented")
}
func (UnimplementedCortexOpsServer) mustEmbedUnimplementedCortexOpsServer() {}

// UnsafeCortexOpsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CortexOpsServer will
// result in compilation errors.
type UnsafeCortexOpsServer interface {
	mustEmbedUnimplementedCortexOpsServer()
}

func RegisterCortexOpsServer(s grpc.ServiceRegistrar, srv CortexOpsServer) {
	s.RegisterService(&CortexOps_ServiceDesc, srv)
}

func _CortexOps_GetClusterConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CortexOpsServer).GetClusterConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cortexops.CortexOps/GetClusterConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CortexOpsServer).GetClusterConfiguration(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CortexOps_ConfigureCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterConfiguration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CortexOpsServer).ConfigureCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cortexops.CortexOps/ConfigureCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CortexOpsServer).ConfigureCluster(ctx, req.(*ClusterConfiguration))
	}
	return interceptor(ctx, in, info, handler)
}

func _CortexOps_GetClusterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CortexOpsServer).GetClusterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cortexops.CortexOps/GetClusterStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CortexOpsServer).GetClusterStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CortexOps_UninstallCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CortexOpsServer).UninstallCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cortexops.CortexOps/UninstallCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CortexOpsServer).UninstallCluster(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CortexOps_ServiceDesc is the grpc.ServiceDesc for CortexOps service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CortexOps_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cortexops.CortexOps",
	HandlerType: (*CortexOpsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClusterConfiguration",
			Handler:    _CortexOps_GetClusterConfiguration_Handler,
		},
		{
			MethodName: "ConfigureCluster",
			Handler:    _CortexOps_ConfigureCluster_Handler,
		},
		{
			MethodName: "GetClusterStatus",
			Handler:    _CortexOps_GetClusterStatus_Handler,
		},
		{
			MethodName: "UninstallCluster",
			Handler:    _CortexOps_UninstallCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/plugins/metrics/pkg/apis/cortexops/cortexops.proto",
}