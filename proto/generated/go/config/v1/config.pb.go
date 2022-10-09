// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: config/v1/config.proto

package configv1

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

// buf:lint:ignore ENUM_VALUE_UPPER_SNAKE_CASE
// buf:lint:ignore ENUM_VALUE_PREFIX
type AppEnv int32

const (
	AppEnv_APP_ENV_UNSPECIFIED AppEnv = 0
	AppEnv_local               AppEnv = 1
)

// Enum value maps for AppEnv.
var (
	AppEnv_name = map[int32]string{
		0: "APP_ENV_UNSPECIFIED",
		1: "local",
	}
	AppEnv_value = map[string]int32{
		"APP_ENV_UNSPECIFIED": 0,
		"local":               1,
	}
)

func (x AppEnv) Enum() *AppEnv {
	p := new(AppEnv)
	*p = x
	return p
}

func (x AppEnv) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppEnv) Descriptor() protoreflect.EnumDescriptor {
	return file_config_v1_config_proto_enumTypes[0].Descriptor()
}

func (AppEnv) Type() protoreflect.EnumType {
	return &file_config_v1_config_proto_enumTypes[0]
}

func (x AppEnv) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppEnv.Descriptor instead.
func (AppEnv) EnumDescriptor() ([]byte, []int) {
	return file_config_v1_config_proto_rawDescGZIP(), []int{0}
}

var file_config_v1_config_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50000,
		Name:          "config.v1.enum_stringer",
		Tag:           "bytes,50000,opt,name=enum_stringer",
		Filename:      "config/v1/config.proto",
	},
}

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional string enum_stringer = 50000;
	E_EnumStringer = &file_config_v1_config_proto_extTypes[0]
)

var File_config_v1_config_proto protoreflect.FileDescriptor

var file_config_v1_config_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x2c, 0x0a, 0x06, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x76, 0x12,
	0x17, 0x0a, 0x13, 0x41, 0x50, 0x50, 0x5f, 0x45, 0x4e, 0x56, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x10, 0x01, 0x3a, 0x4b, 0x0a, 0x0d, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x6e, 0x75, 0x6d, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x75, 0x6e, 0x69, 0x74, 0x73, 0x75, 0x69, 0x6e, 0x63, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_config_v1_config_proto_rawDescOnce sync.Once
	file_config_v1_config_proto_rawDescData = file_config_v1_config_proto_rawDesc
)

func file_config_v1_config_proto_rawDescGZIP() []byte {
	file_config_v1_config_proto_rawDescOnce.Do(func() {
		file_config_v1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_v1_config_proto_rawDescData)
	})
	return file_config_v1_config_proto_rawDescData
}

var file_config_v1_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config_v1_config_proto_goTypes = []interface{}{
	(AppEnv)(0),                           // 0: config.v1.AppEnv
	(*descriptorpb.EnumValueOptions)(nil), // 1: google.protobuf.EnumValueOptions
}
var file_config_v1_config_proto_depIdxs = []int32{
	1, // 0: config.v1.enum_stringer:extendee -> google.protobuf.EnumValueOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_config_v1_config_proto_init() }
func file_config_v1_config_proto_init() {
	if File_config_v1_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_v1_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_config_v1_config_proto_goTypes,
		DependencyIndexes: file_config_v1_config_proto_depIdxs,
		EnumInfos:         file_config_v1_config_proto_enumTypes,
		ExtensionInfos:    file_config_v1_config_proto_extTypes,
	}.Build()
	File_config_v1_config_proto = out.File
	file_config_v1_config_proto_rawDesc = nil
	file_config_v1_config_proto_goTypes = nil
	file_config_v1_config_proto_depIdxs = nil
}