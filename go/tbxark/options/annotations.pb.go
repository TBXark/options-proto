// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: tbxark/options/annotations.proto

package options

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KeyValuePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are assignable to Value:
	//
	//	*KeyValuePair_Flag
	//	*KeyValuePair_Text
	//	*KeyValuePair_Number
	Value isKeyValuePair_Value `protobuf_oneof:"value"`
	Extra map[string]string    `protobuf:"bytes,5,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *KeyValuePair) Reset() {
	*x = KeyValuePair{}
	mi := &file_tbxark_options_annotations_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeyValuePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValuePair) ProtoMessage() {}

func (x *KeyValuePair) ProtoReflect() protoreflect.Message {
	mi := &file_tbxark_options_annotations_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValuePair.ProtoReflect.Descriptor instead.
func (*KeyValuePair) Descriptor() ([]byte, []int) {
	return file_tbxark_options_annotations_proto_rawDescGZIP(), []int{0}
}

func (x *KeyValuePair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (m *KeyValuePair) GetValue() isKeyValuePair_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *KeyValuePair) GetFlag() bool {
	if x, ok := x.GetValue().(*KeyValuePair_Flag); ok {
		return x.Flag
	}
	return false
}

func (x *KeyValuePair) GetText() string {
	if x, ok := x.GetValue().(*KeyValuePair_Text); ok {
		return x.Text
	}
	return ""
}

func (x *KeyValuePair) GetNumber() int64 {
	if x, ok := x.GetValue().(*KeyValuePair_Number); ok {
		return x.Number
	}
	return 0
}

func (x *KeyValuePair) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type isKeyValuePair_Value interface {
	isKeyValuePair_Value()
}

type KeyValuePair_Flag struct {
	Flag bool `protobuf:"varint,2,opt,name=flag,proto3,oneof"`
}

type KeyValuePair_Text struct {
	Text string `protobuf:"bytes,3,opt,name=text,proto3,oneof"`
}

type KeyValuePair_Number struct {
	Number int64 `protobuf:"varint,4,opt,name=number,proto3,oneof"`
}

func (*KeyValuePair_Flag) isKeyValuePair_Value() {}

func (*KeyValuePair_Text) isKeyValuePair_Value() {}

func (*KeyValuePair_Number) isKeyValuePair_Value() {}

var file_tbxark_options_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: ([]*KeyValuePair)(nil),
		Field:         114514,
		Name:          "tbxark.options.options",
		Tag:           "bytes,114514,rep,name=options",
		Filename:      "tbxark/options/annotations.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// repeated tbxark.options.KeyValuePair options = 114514;
	E_Options = &file_tbxark_options_annotations_proto_extTypes[0]
)

var File_tbxark_options_annotations_proto protoreflect.FileDescriptor

var file_tbxark_options_annotations_proto_rawDesc = []byte{
	0x0a, 0x20, 0x74, 0x62, 0x78, 0x61, 0x72, 0x6b, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x74, 0x62, 0x78, 0x61, 0x72, 0x6b, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x0c, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x50, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x14, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3d, 0x0a,
	0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74,
	0x62, 0x78, 0x61, 0x72, 0x6b, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4b, 0x65,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x58, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd2, 0xfe, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x62, 0x78, 0x61, 0x72, 0x6b, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72,
	0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0xca, 0x01, 0x0a, 0x12, 0x63, 0x6f,
	0x6d, 0x2e, 0x74, 0x62, 0x78, 0x61, 0x72, 0x6b, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x42, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x62, 0x75, 0x66, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x62, 0x78, 0x61, 0x72, 0x6b, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x62, 0x78, 0x61, 0x72, 0x6b, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xa2,
	0x02, 0x03, 0x54, 0x4f, 0x58, 0xaa, 0x02, 0x0e, 0x54, 0x62, 0x78, 0x61, 0x72, 0x6b, 0x2e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x0e, 0x54, 0x62, 0x78, 0x61, 0x72, 0x6b, 0x5c,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x1a, 0x54, 0x62, 0x78, 0x61, 0x72, 0x6b,
	0x5c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x54, 0x62, 0x78, 0x61, 0x72, 0x6b, 0x3a, 0x3a, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tbxark_options_annotations_proto_rawDescOnce sync.Once
	file_tbxark_options_annotations_proto_rawDescData = file_tbxark_options_annotations_proto_rawDesc
)

func file_tbxark_options_annotations_proto_rawDescGZIP() []byte {
	file_tbxark_options_annotations_proto_rawDescOnce.Do(func() {
		file_tbxark_options_annotations_proto_rawDescData = protoimpl.X.CompressGZIP(file_tbxark_options_annotations_proto_rawDescData)
	})
	return file_tbxark_options_annotations_proto_rawDescData
}

var file_tbxark_options_annotations_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tbxark_options_annotations_proto_goTypes = []any{
	(*KeyValuePair)(nil),               // 0: tbxark.options.KeyValuePair
	nil,                                // 1: tbxark.options.KeyValuePair.ExtraEntry
	(*descriptorpb.MethodOptions)(nil), // 2: google.protobuf.MethodOptions
}
var file_tbxark_options_annotations_proto_depIdxs = []int32{
	1, // 0: tbxark.options.KeyValuePair.extra:type_name -> tbxark.options.KeyValuePair.ExtraEntry
	2, // 1: tbxark.options.options:extendee -> google.protobuf.MethodOptions
	0, // 2: tbxark.options.options:type_name -> tbxark.options.KeyValuePair
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tbxark_options_annotations_proto_init() }
func file_tbxark_options_annotations_proto_init() {
	if File_tbxark_options_annotations_proto != nil {
		return
	}
	file_tbxark_options_annotations_proto_msgTypes[0].OneofWrappers = []any{
		(*KeyValuePair_Flag)(nil),
		(*KeyValuePair_Text)(nil),
		(*KeyValuePair_Number)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tbxark_options_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_tbxark_options_annotations_proto_goTypes,
		DependencyIndexes: file_tbxark_options_annotations_proto_depIdxs,
		MessageInfos:      file_tbxark_options_annotations_proto_msgTypes,
		ExtensionInfos:    file_tbxark_options_annotations_proto_extTypes,
	}.Build()
	File_tbxark_options_annotations_proto = out.File
	file_tbxark_options_annotations_proto_rawDesc = nil
	file_tbxark_options_annotations_proto_goTypes = nil
	file_tbxark_options_annotations_proto_depIdxs = nil
}