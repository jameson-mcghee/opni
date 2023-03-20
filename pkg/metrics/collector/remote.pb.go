// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v1.0.0
// source: github.com/rancher/opni/pkg/metrics/collector/remote.proto

package collector

import (
	_go "github.com/prometheus/client_model/go"
	desc "github.com/rancher/opni/pkg/metrics/desc"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DescriptorList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Descriptors []*desc.Desc `protobuf:"bytes,1,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
}

func (x *DescriptorList) Reset() {
	*x = DescriptorList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescriptorList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescriptorList) ProtoMessage() {}

func (x *DescriptorList) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescriptorList.ProtoReflect.Descriptor instead.
func (*DescriptorList) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_rawDescGZIP(), []int{0}
}

func (x *DescriptorList) GetDescriptors() []*desc.Desc {
	if x != nil {
		return x.Descriptors
	}
	return nil
}

type MetricList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *MetricList) Reset() {
	*x = MetricList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricList) ProtoMessage() {}

func (x *MetricList) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricList.ProtoReflect.Descriptor instead.
func (*MetricList) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_rawDescGZIP(), []int{1}
}

func (x *MetricList) GetMetrics() []*Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Desc   *desc.Desc  `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`
	Metric *_go.Metric `protobuf:"bytes,2,opt,name=metric,proto3" json:"metric,omitempty"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_rawDescGZIP(), []int{2}
}

func (x *Metric) GetDesc() *desc.Desc {
	if x != nil {
		return x.Desc
	}
	return nil
}

func (x *Metric) GetMetric() *_go.Metric {
	if x != nil {
		return x.Metric
	}
	return nil
}

var File_github_com_rancher_opni_pkg_metrics_collector_remote_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6d,
	0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f,
	0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3e, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x64, 0x65, 0x73, 0x63, 0x2e, 0x44,
	0x65, 0x73, 0x63, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73,
	0x22, 0x39, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b,
	0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x5e, 0x0a, 0x06, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x64, 0x65, 0x73, 0x63, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x52,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x34, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x65,
	0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x32, 0x8a, 0x01, 0x0a, 0x0f,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x3d, 0x0a, 0x08, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38,
	0x0a, 0x07, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f,
	0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_rawDescData = file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_rawDesc
)

func file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_rawDescData)
	})
	return file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_rawDescData
}

var file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_goTypes = []interface{}{
	(*DescriptorList)(nil), // 0: collector.DescriptorList
	(*MetricList)(nil),     // 1: collector.MetricList
	(*Metric)(nil),         // 2: collector.Metric
	(*desc.Desc)(nil),      // 3: desc.Desc
	(*_go.Metric)(nil),     // 4: io.prometheus.client.Metric
	(*emptypb.Empty)(nil),  // 5: google.protobuf.Empty
}
var file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_depIdxs = []int32{
	3, // 0: collector.DescriptorList.descriptors:type_name -> desc.Desc
	2, // 1: collector.MetricList.metrics:type_name -> collector.Metric
	3, // 2: collector.Metric.desc:type_name -> desc.Desc
	4, // 3: collector.Metric.metric:type_name -> io.prometheus.client.Metric
	5, // 4: collector.RemoteCollector.Describe:input_type -> google.protobuf.Empty
	5, // 5: collector.RemoteCollector.Collect:input_type -> google.protobuf.Empty
	0, // 6: collector.RemoteCollector.Describe:output_type -> collector.DescriptorList
	1, // 7: collector.RemoteCollector.Collect:output_type -> collector.MetricList
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_init() }
func file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_init() {
	if File_github_com_rancher_opni_pkg_metrics_collector_remote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescriptorList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_depIdxs,
		MessageInfos:      file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_msgTypes,
	}.Build()
	File_github_com_rancher_opni_pkg_metrics_collector_remote_proto = out.File
	file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_rawDesc = nil
	file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_goTypes = nil
	file_github_com_rancher_opni_pkg_metrics_collector_remote_proto_depIdxs = nil
}
