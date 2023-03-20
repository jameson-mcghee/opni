// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/plugins/aiops/pkg/apis/modeltraining/modeltraining.proto

package modeltraining

import (
	context "context"
	v1 "github.com/rancher/opni/pkg/apis/core/v1"
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
	ModelTraining_TrainModel_FullMethodName                 = "/modeltraining.ModelTraining/TrainModel"
	ModelTraining_PutModelTrainingStatus_FullMethodName     = "/modeltraining.ModelTraining/PutModelTrainingStatus"
	ModelTraining_ClusterWorkloadAggregation_FullMethodName = "/modeltraining.ModelTraining/ClusterWorkloadAggregation"
	ModelTraining_GetModelStatus_FullMethodName             = "/modeltraining.ModelTraining/GetModelStatus"
	ModelTraining_GetModelTrainingParameters_FullMethodName = "/modeltraining.ModelTraining/GetModelTrainingParameters"
	ModelTraining_GPUInfo_FullMethodName                    = "/modeltraining.ModelTraining/GPUInfo"
)

// ModelTrainingClient is the client API for ModelTraining service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelTrainingClient interface {
	TrainModel(ctx context.Context, in *ModelTrainingParametersList, opts ...grpc.CallOption) (*ModelTrainingResponse, error)
	PutModelTrainingStatus(ctx context.Context, in *ModelTrainingStatistics, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ClusterWorkloadAggregation(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*WorkloadAggregationList, error)
	GetModelStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ModelStatus, error)
	GetModelTrainingParameters(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ModelTrainingParametersList, error)
	GPUInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GPUInfoList, error)
}

type modelTrainingClient struct {
	cc grpc.ClientConnInterface
}

func NewModelTrainingClient(cc grpc.ClientConnInterface) ModelTrainingClient {
	return &modelTrainingClient{cc}
}

func (c *modelTrainingClient) TrainModel(ctx context.Context, in *ModelTrainingParametersList, opts ...grpc.CallOption) (*ModelTrainingResponse, error) {
	out := new(ModelTrainingResponse)
	err := c.cc.Invoke(ctx, ModelTraining_TrainModel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelTrainingClient) PutModelTrainingStatus(ctx context.Context, in *ModelTrainingStatistics, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ModelTraining_PutModelTrainingStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelTrainingClient) ClusterWorkloadAggregation(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*WorkloadAggregationList, error) {
	out := new(WorkloadAggregationList)
	err := c.cc.Invoke(ctx, ModelTraining_ClusterWorkloadAggregation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelTrainingClient) GetModelStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ModelStatus, error) {
	out := new(ModelStatus)
	err := c.cc.Invoke(ctx, ModelTraining_GetModelStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelTrainingClient) GetModelTrainingParameters(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ModelTrainingParametersList, error) {
	out := new(ModelTrainingParametersList)
	err := c.cc.Invoke(ctx, ModelTraining_GetModelTrainingParameters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelTrainingClient) GPUInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GPUInfoList, error) {
	out := new(GPUInfoList)
	err := c.cc.Invoke(ctx, ModelTraining_GPUInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelTrainingServer is the server API for ModelTraining service.
// All implementations must embed UnimplementedModelTrainingServer
// for forward compatibility
type ModelTrainingServer interface {
	TrainModel(context.Context, *ModelTrainingParametersList) (*ModelTrainingResponse, error)
	PutModelTrainingStatus(context.Context, *ModelTrainingStatistics) (*emptypb.Empty, error)
	ClusterWorkloadAggregation(context.Context, *v1.Reference) (*WorkloadAggregationList, error)
	GetModelStatus(context.Context, *emptypb.Empty) (*ModelStatus, error)
	GetModelTrainingParameters(context.Context, *emptypb.Empty) (*ModelTrainingParametersList, error)
	GPUInfo(context.Context, *emptypb.Empty) (*GPUInfoList, error)
	mustEmbedUnimplementedModelTrainingServer()
}

// UnimplementedModelTrainingServer must be embedded to have forward compatible implementations.
type UnimplementedModelTrainingServer struct {
}

func (UnimplementedModelTrainingServer) TrainModel(context.Context, *ModelTrainingParametersList) (*ModelTrainingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrainModel not implemented")
}
func (UnimplementedModelTrainingServer) PutModelTrainingStatus(context.Context, *ModelTrainingStatistics) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutModelTrainingStatus not implemented")
}
func (UnimplementedModelTrainingServer) ClusterWorkloadAggregation(context.Context, *v1.Reference) (*WorkloadAggregationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterWorkloadAggregation not implemented")
}
func (UnimplementedModelTrainingServer) GetModelStatus(context.Context, *emptypb.Empty) (*ModelStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModelStatus not implemented")
}
func (UnimplementedModelTrainingServer) GetModelTrainingParameters(context.Context, *emptypb.Empty) (*ModelTrainingParametersList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModelTrainingParameters not implemented")
}
func (UnimplementedModelTrainingServer) GPUInfo(context.Context, *emptypb.Empty) (*GPUInfoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GPUInfo not implemented")
}
func (UnimplementedModelTrainingServer) mustEmbedUnimplementedModelTrainingServer() {}

// UnsafeModelTrainingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelTrainingServer will
// result in compilation errors.
type UnsafeModelTrainingServer interface {
	mustEmbedUnimplementedModelTrainingServer()
}

func RegisterModelTrainingServer(s grpc.ServiceRegistrar, srv ModelTrainingServer) {
	s.RegisterService(&ModelTraining_ServiceDesc, srv)
}

func _ModelTraining_TrainModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelTrainingParametersList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelTrainingServer).TrainModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelTraining_TrainModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelTrainingServer).TrainModel(ctx, req.(*ModelTrainingParametersList))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelTraining_PutModelTrainingStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelTrainingStatistics)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelTrainingServer).PutModelTrainingStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelTraining_PutModelTrainingStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelTrainingServer).PutModelTrainingStatus(ctx, req.(*ModelTrainingStatistics))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelTraining_ClusterWorkloadAggregation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelTrainingServer).ClusterWorkloadAggregation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelTraining_ClusterWorkloadAggregation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelTrainingServer).ClusterWorkloadAggregation(ctx, req.(*v1.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelTraining_GetModelStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelTrainingServer).GetModelStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelTraining_GetModelStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelTrainingServer).GetModelStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelTraining_GetModelTrainingParameters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelTrainingServer).GetModelTrainingParameters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelTraining_GetModelTrainingParameters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelTrainingServer).GetModelTrainingParameters(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelTraining_GPUInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelTrainingServer).GPUInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelTraining_GPUInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelTrainingServer).GPUInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ModelTraining_ServiceDesc is the grpc.ServiceDesc for ModelTraining service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelTraining_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "modeltraining.ModelTraining",
	HandlerType: (*ModelTrainingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TrainModel",
			Handler:    _ModelTraining_TrainModel_Handler,
		},
		{
			MethodName: "PutModelTrainingStatus",
			Handler:    _ModelTraining_PutModelTrainingStatus_Handler,
		},
		{
			MethodName: "ClusterWorkloadAggregation",
			Handler:    _ModelTraining_ClusterWorkloadAggregation_Handler,
		},
		{
			MethodName: "GetModelStatus",
			Handler:    _ModelTraining_GetModelStatus_Handler,
		},
		{
			MethodName: "GetModelTrainingParameters",
			Handler:    _ModelTraining_GetModelTrainingParameters_Handler,
		},
		{
			MethodName: "GPUInfo",
			Handler:    _ModelTraining_GPUInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/plugins/aiops/pkg/apis/modeltraining/modeltraining.proto",
}
