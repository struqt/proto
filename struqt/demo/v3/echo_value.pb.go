// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: struqt/demo/v3/echo_value.proto

package demo

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EchoValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *structpb.Value `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EchoValueRequest) Reset() {
	*x = EchoValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_struqt_demo_v3_echo_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoValueRequest) ProtoMessage() {}

func (x *EchoValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_struqt_demo_v3_echo_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoValueRequest.ProtoReflect.Descriptor instead.
func (*EchoValueRequest) Descriptor() ([]byte, []int) {
	return file_struqt_demo_v3_echo_value_proto_rawDescGZIP(), []int{0}
}

func (x *EchoValueRequest) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type EchoValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value      *structpb.Value        `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Pb3Message *Proto3Message         `protobuf:"bytes,3,opt,name=pb3_message,json=pb3Message,proto3" json:"pb3_message,omitempty"`
}

func (x *EchoValueResponse) Reset() {
	*x = EchoValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_struqt_demo_v3_echo_value_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoValueResponse) ProtoMessage() {}

func (x *EchoValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_struqt_demo_v3_echo_value_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoValueResponse.ProtoReflect.Descriptor instead.
func (*EchoValueResponse) Descriptor() ([]byte, []int) {
	return file_struqt_demo_v3_echo_value_proto_rawDescGZIP(), []int{1}
}

func (x *EchoValueResponse) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *EchoValueResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *EchoValueResponse) GetPb3Message() *Proto3Message {
	if x != nil {
		return x.Pb3Message
	}
	return nil
}

var File_struqt_demo_v3_echo_value_proto protoreflect.FileDescriptor

var file_struqt_demo_v3_echo_value_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x74, 0x72, 0x75, 0x71, 0x74, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x76, 0x33,
	0x2f, 0x65, 0x63, 0x68, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x73, 0x74, 0x72, 0x75, 0x71, 0x74, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76,
	0x33, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x73, 0x74, 0x72, 0x75, 0x71, 0x74, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x76, 0x33,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a,
	0x10, 0x45, 0x63, 0x68, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0xbb, 0x01, 0x0a, 0x11, 0x45, 0x63, 0x68, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3e, 0x0a,
	0x0b, 0x70, 0x62, 0x33, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x74, 0x72, 0x75, 0x71, 0x74, 0x2e, 0x64, 0x65, 0x6d, 0x6f,
	0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x0a, 0x70, 0x62, 0x33, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0xab, 0x01,
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x72, 0x75, 0x71, 0x74, 0x2e, 0x64, 0x65, 0x6d,
	0x6f, 0x2e, 0x76, 0x33, 0x42, 0x0e, 0x45, 0x63, 0x68, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x71, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x71, 0x74, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x76, 0x33, 0x3b, 0x64,
	0x65, 0x6d, 0x6f, 0xa2, 0x02, 0x03, 0x53, 0x44, 0x58, 0xaa, 0x02, 0x0e, 0x53, 0x74, 0x72, 0x75,
	0x71, 0x74, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x2e, 0x56, 0x33, 0xca, 0x02, 0x0e, 0x53, 0x74, 0x72,
	0x75, 0x71, 0x74, 0x5c, 0x44, 0x65, 0x6d, 0x6f, 0x5c, 0x56, 0x33, 0xe2, 0x02, 0x1a, 0x53, 0x74,
	0x72, 0x75, 0x71, 0x74, 0x5c, 0x44, 0x65, 0x6d, 0x6f, 0x5c, 0x56, 0x33, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x53, 0x74, 0x72, 0x75, 0x71,
	0x74, 0x3a, 0x3a, 0x44, 0x65, 0x6d, 0x6f, 0x3a, 0x3a, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_struqt_demo_v3_echo_value_proto_rawDescOnce sync.Once
	file_struqt_demo_v3_echo_value_proto_rawDescData = file_struqt_demo_v3_echo_value_proto_rawDesc
)

func file_struqt_demo_v3_echo_value_proto_rawDescGZIP() []byte {
	file_struqt_demo_v3_echo_value_proto_rawDescOnce.Do(func() {
		file_struqt_demo_v3_echo_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_struqt_demo_v3_echo_value_proto_rawDescData)
	})
	return file_struqt_demo_v3_echo_value_proto_rawDescData
}

var file_struqt_demo_v3_echo_value_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_struqt_demo_v3_echo_value_proto_goTypes = []interface{}{
	(*EchoValueRequest)(nil),      // 0: struqt.demo.v3.EchoValueRequest
	(*EchoValueResponse)(nil),     // 1: struqt.demo.v3.EchoValueResponse
	(*structpb.Value)(nil),        // 2: google.protobuf.Value
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*Proto3Message)(nil),         // 4: struqt.demo.v3.Proto3Message
}
var file_struqt_demo_v3_echo_value_proto_depIdxs = []int32{
	2, // 0: struqt.demo.v3.EchoValueRequest.value:type_name -> google.protobuf.Value
	2, // 1: struqt.demo.v3.EchoValueResponse.value:type_name -> google.protobuf.Value
	3, // 2: struqt.demo.v3.EchoValueResponse.timestamp:type_name -> google.protobuf.Timestamp
	4, // 3: struqt.demo.v3.EchoValueResponse.pb3_message:type_name -> struqt.demo.v3.Proto3Message
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_struqt_demo_v3_echo_value_proto_init() }
func file_struqt_demo_v3_echo_value_proto_init() {
	if File_struqt_demo_v3_echo_value_proto != nil {
		return
	}
	file_struqt_demo_v3_proto3_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_struqt_demo_v3_echo_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoValueRequest); i {
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
		file_struqt_demo_v3_echo_value_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoValueResponse); i {
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
			RawDescriptor: file_struqt_demo_v3_echo_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_struqt_demo_v3_echo_value_proto_goTypes,
		DependencyIndexes: file_struqt_demo_v3_echo_value_proto_depIdxs,
		MessageInfos:      file_struqt_demo_v3_echo_value_proto_msgTypes,
	}.Build()
	File_struqt_demo_v3_echo_value_proto = out.File
	file_struqt_demo_v3_echo_value_proto_rawDesc = nil
	file_struqt_demo_v3_echo_value_proto_goTypes = nil
	file_struqt_demo_v3_echo_value_proto_depIdxs = nil
}