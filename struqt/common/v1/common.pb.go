// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.3
// source: struqt/common/v1/common.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Dummy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Dummy) Reset() {
	*x = Dummy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_struqt_common_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dummy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dummy) ProtoMessage() {}

func (x *Dummy) ProtoReflect() protoreflect.Message {
	mi := &file_struqt_common_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dummy.ProtoReflect.Descriptor instead.
func (*Dummy) Descriptor() ([]byte, []int) {
	return file_struqt_common_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Dummy) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

var File_struqt_common_v1_common_proto protoreflect.FileDescriptor

var file_struqt_common_v1_common_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x74, 0x72, 0x75, 0x71, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x73, 0x74, 0x72, 0x75, 0x71, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x22, 0x21, 0x0a, 0x05, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x42, 0x60, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x72, 0x75,
	0x71, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x71, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x71, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0xaa, 0x02, 0x10, 0x53, 0x74, 0x72, 0x75, 0x71, 0x74, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_struqt_common_v1_common_proto_rawDescOnce sync.Once
	file_struqt_common_v1_common_proto_rawDescData = file_struqt_common_v1_common_proto_rawDesc
)

func file_struqt_common_v1_common_proto_rawDescGZIP() []byte {
	file_struqt_common_v1_common_proto_rawDescOnce.Do(func() {
		file_struqt_common_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_struqt_common_v1_common_proto_rawDescData)
	})
	return file_struqt_common_v1_common_proto_rawDescData
}

var file_struqt_common_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_struqt_common_v1_common_proto_goTypes = []interface{}{
	(*Dummy)(nil), // 0: struqt.common.v1.Dummy
}
var file_struqt_common_v1_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_struqt_common_v1_common_proto_init() }
func file_struqt_common_v1_common_proto_init() {
	if File_struqt_common_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_struqt_common_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dummy); i {
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
			RawDescriptor: file_struqt_common_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_struqt_common_v1_common_proto_goTypes,
		DependencyIndexes: file_struqt_common_v1_common_proto_depIdxs,
		MessageInfos:      file_struqt_common_v1_common_proto_msgTypes,
	}.Build()
	File_struqt_common_v1_common_proto = out.File
	file_struqt_common_v1_common_proto_rawDesc = nil
	file_struqt_common_v1_common_proto_goTypes = nil
	file_struqt_common_v1_common_proto_depIdxs = nil
}
