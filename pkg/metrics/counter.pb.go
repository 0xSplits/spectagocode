// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: pbf/metrics/counter.proto

package metrics

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CounterI is the input for updating counters.
type CounterI struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Action        []*Action              `protobuf:"bytes,100,rep,name=action,proto3" json:"action,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CounterI) Reset() {
	*x = CounterI{}
	mi := &file_pbf_metrics_counter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CounterI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterI) ProtoMessage() {}

func (x *CounterI) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_metrics_counter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterI.ProtoReflect.Descriptor instead.
func (*CounterI) Descriptor() ([]byte, []int) {
	return file_pbf_metrics_counter_proto_rawDescGZIP(), []int{0}
}

func (x *CounterI) GetAction() []*Action {
	if x != nil {
		return x.Action
	}
	return nil
}

// CounterO is the output for updating counters.
type CounterO struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        []*Result              `protobuf:"bytes,100,rep,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CounterO) Reset() {
	*x = CounterO{}
	mi := &file_pbf_metrics_counter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CounterO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterO) ProtoMessage() {}

func (x *CounterO) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_metrics_counter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterO.ProtoReflect.Descriptor instead.
func (*CounterO) Descriptor() ([]byte, []int) {
	return file_pbf_metrics_counter_proto_rawDescGZIP(), []int{1}
}

func (x *CounterO) GetResult() []*Result {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_pbf_metrics_counter_proto protoreflect.FileDescriptor

const file_pbf_metrics_counter_proto_rawDesc = "" +
	"\n" +
	"\x19pbf/metrics/counter.proto\x12\ametrics\x1a\x18pbf/metrics/common.proto\"3\n" +
	"\bCounterI\x12'\n" +
	"\x06action\x18d \x03(\v2\x0f.metrics.ActionR\x06action\"3\n" +
	"\bCounterO\x12'\n" +
	"\x06result\x18d \x03(\v2\x0f.metrics.ResultR\x06resultB\fZ\n" +
	"./;metricsb\x06proto3"

var (
	file_pbf_metrics_counter_proto_rawDescOnce sync.Once
	file_pbf_metrics_counter_proto_rawDescData []byte
)

func file_pbf_metrics_counter_proto_rawDescGZIP() []byte {
	file_pbf_metrics_counter_proto_rawDescOnce.Do(func() {
		file_pbf_metrics_counter_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pbf_metrics_counter_proto_rawDesc), len(file_pbf_metrics_counter_proto_rawDesc)))
	})
	return file_pbf_metrics_counter_proto_rawDescData
}

var file_pbf_metrics_counter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pbf_metrics_counter_proto_goTypes = []any{
	(*CounterI)(nil), // 0: metrics.CounterI
	(*CounterO)(nil), // 1: metrics.CounterO
	(*Action)(nil),   // 2: metrics.Action
	(*Result)(nil),   // 3: metrics.Result
}
var file_pbf_metrics_counter_proto_depIdxs = []int32{
	2, // 0: metrics.CounterI.action:type_name -> metrics.Action
	3, // 1: metrics.CounterO.result:type_name -> metrics.Result
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pbf_metrics_counter_proto_init() }
func file_pbf_metrics_counter_proto_init() {
	if File_pbf_metrics_counter_proto != nil {
		return
	}
	file_pbf_metrics_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pbf_metrics_counter_proto_rawDesc), len(file_pbf_metrics_counter_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbf_metrics_counter_proto_goTypes,
		DependencyIndexes: file_pbf_metrics_counter_proto_depIdxs,
		MessageInfos:      file_pbf_metrics_counter_proto_msgTypes,
	}.Build()
	File_pbf_metrics_counter_proto = out.File
	file_pbf_metrics_counter_proto_goTypes = nil
	file_pbf_metrics_counter_proto_depIdxs = nil
}
