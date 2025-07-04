// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: pbf/metrics/common.proto

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

// Action is the request body for any given action.
//
//	{
//	    "labels": {
//	        "key": "value"
//	    },
//	    "metric": "teams_bridge_error_total",
//	    "number": 3.141592
//	}
type Action struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// labels is the optional set of key-value pairs whitelisted per metric.
	// Providing labels that are not whitelisted causes the entire RPC to be
	// rejected on a request level.
	Labels map[string]string `protobuf:"bytes,100,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// metric is the required metric name in camel case format. Providing metric
	// names that are not whitelisted causes the entire RPC to be rejected on a
	// request level. Metric names must not be longer than 255 characters.
	Metric string `protobuf:"bytes,200,opt,name=metric,proto3" json:"metric,omitempty"`
	// number is the required floating point value that should be applied to the
	// timeseries specified by the labels and metric properties. Negative values
	// may cause the entire RPC to be rejected on a per metric basis.
	Number        float64 `protobuf:"fixed64,300,opt,name=number,proto3" json:"number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Action) Reset() {
	*x = Action{}
	mi := &file_pbf_metrics_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_metrics_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_pbf_metrics_common_proto_rawDescGZIP(), []int{0}
}

func (x *Action) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Action) GetMetric() string {
	if x != nil {
		return x.Metric
	}
	return ""
}

func (x *Action) GetNumber() float64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type Reason struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Info          string                 `protobuf:"bytes,100,opt,name=info,proto3" json:"info,omitempty"`
	Kind          string                 `protobuf:"bytes,200,opt,name=kind,proto3" json:"kind,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Reason) Reset() {
	*x = Reason{}
	mi := &file_pbf_metrics_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Reason) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reason) ProtoMessage() {}

func (x *Reason) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_metrics_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reason.ProtoReflect.Descriptor instead.
func (*Reason) Descriptor() ([]byte, []int) {
	return file_pbf_metrics_common_proto_rawDescGZIP(), []int{1}
}

func (x *Reason) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *Reason) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

// Result is the response body for any given action.
//
//	{
//	    "reason": {
//	        "desc": "This is why the request was rejected.",
//	        "kind": "someErrorCode"
//	    },
//	    "status": "failure"
//	}
type Result struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// reason is the optional feedback that the backend provides in case of
	// action failure.
	Reason *Reason `protobuf:"bytes,100,opt,name=reason,proto3,oneof" json:"reason,omitempty"`
	// status is the final state of the executed action.
	//
	//	success, the given actions were executed successfully
	//
	//	failure, the given actions were rejected for reasons
	Status        string `protobuf:"bytes,200,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Result) Reset() {
	*x = Result{}
	mi := &file_pbf_metrics_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_pbf_metrics_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_pbf_metrics_common_proto_rawDescGZIP(), []int{2}
}

func (x *Result) GetReason() *Reason {
	if x != nil {
		return x.Reason
	}
	return nil
}

func (x *Result) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_pbf_metrics_common_proto protoreflect.FileDescriptor

const file_pbf_metrics_common_proto_rawDesc = "" +
	"\n" +
	"\x18pbf/metrics/common.proto\x12\ametrics\"\xaa\x01\n" +
	"\x06Action\x123\n" +
	"\x06labels\x18d \x03(\v2\x1b.metrics.Action.LabelsEntryR\x06labels\x12\x17\n" +
	"\x06metric\x18\xc8\x01 \x01(\tR\x06metric\x12\x17\n" +
	"\x06number\x18\xac\x02 \x01(\x01R\x06number\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"1\n" +
	"\x06Reason\x12\x12\n" +
	"\x04info\x18d \x01(\tR\x04info\x12\x13\n" +
	"\x04kind\x18\xc8\x01 \x01(\tR\x04kind\"Z\n" +
	"\x06Result\x12,\n" +
	"\x06reason\x18d \x01(\v2\x0f.metrics.ReasonH\x00R\x06reason\x88\x01\x01\x12\x17\n" +
	"\x06status\x18\xc8\x01 \x01(\tR\x06statusB\t\n" +
	"\a_reasonB\fZ\n" +
	"./;metricsb\x06proto3"

var (
	file_pbf_metrics_common_proto_rawDescOnce sync.Once
	file_pbf_metrics_common_proto_rawDescData []byte
)

func file_pbf_metrics_common_proto_rawDescGZIP() []byte {
	file_pbf_metrics_common_proto_rawDescOnce.Do(func() {
		file_pbf_metrics_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pbf_metrics_common_proto_rawDesc), len(file_pbf_metrics_common_proto_rawDesc)))
	})
	return file_pbf_metrics_common_proto_rawDescData
}

var file_pbf_metrics_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pbf_metrics_common_proto_goTypes = []any{
	(*Action)(nil), // 0: metrics.Action
	(*Reason)(nil), // 1: metrics.Reason
	(*Result)(nil), // 2: metrics.Result
	nil,            // 3: metrics.Action.LabelsEntry
}
var file_pbf_metrics_common_proto_depIdxs = []int32{
	3, // 0: metrics.Action.labels:type_name -> metrics.Action.LabelsEntry
	1, // 1: metrics.Result.reason:type_name -> metrics.Reason
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pbf_metrics_common_proto_init() }
func file_pbf_metrics_common_proto_init() {
	if File_pbf_metrics_common_proto != nil {
		return
	}
	file_pbf_metrics_common_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pbf_metrics_common_proto_rawDesc), len(file_pbf_metrics_common_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbf_metrics_common_proto_goTypes,
		DependencyIndexes: file_pbf_metrics_common_proto_depIdxs,
		MessageInfos:      file_pbf_metrics_common_proto_msgTypes,
	}.Build()
	File_pbf_metrics_common_proto = out.File
	file_pbf_metrics_common_proto_goTypes = nil
	file_pbf_metrics_common_proto_depIdxs = nil
}
