// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.17.3
// source: echo/v1/echo.proto

package echov1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/lyft/clutch/backend/api/api/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SayHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SayHelloRequest) Reset() {
	*x = SayHelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_v1_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloRequest) ProtoMessage() {}

func (x *SayHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_echo_v1_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloRequest.ProtoReflect.Descriptor instead.
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return file_echo_v1_echo_proto_rawDescGZIP(), []int{0}
}

func (x *SayHelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SayHelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SayHelloResponse) Reset() {
	*x = SayHelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_v1_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloResponse) ProtoMessage() {}

func (x *SayHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_echo_v1_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloResponse.ProtoReflect.Descriptor instead.
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return file_echo_v1_echo_proto_rawDescGZIP(), []int{1}
}

func (x *SayHelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_echo_v1_echo_proto protoreflect.FileDescriptor

var file_echo_v1_echo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x44, 0x61, 0x6e, 0x69, 0x65, 0x6c, 0x59, 0x65, 0x76, 0x65,
	0x6c, 0x6b, 0x69, 0x6e, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a,
	0x0f, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a,
	0x10, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x8d, 0x01, 0x0a, 0x07,
	0x45, 0x63, 0x68, 0x6f, 0x41, 0x50, 0x49, 0x12, 0x81, 0x01, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x27, 0x2e, 0x44, 0x61, 0x6e, 0x69, 0x65, 0x6c, 0x59, 0x65, 0x76,
	0x65, 0x6c, 0x6b, 0x69, 0x6e, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61,
	0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x44, 0x61, 0x6e, 0x69, 0x65, 0x6c, 0x59, 0x65, 0x76, 0x65, 0x6c, 0x6b, 0x69, 0x6e, 0x2e, 0x65,
	0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0xaa, 0xe1, 0x1c, 0x02, 0x08, 0x02, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x63,
	0x68, 0x6f, 0x2f, 0x73, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x42, 0x4c, 0x5a, 0x4a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x6e, 0x69, 0x65, 0x6c,
	0x59, 0x65, 0x76, 0x65, 0x6c, 0x6b, 0x69, 0x6e, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2d,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2f,
	0x76, 0x31, 0x3b, 0x65, 0x63, 0x68, 0x6f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_echo_v1_echo_proto_rawDescOnce sync.Once
	file_echo_v1_echo_proto_rawDescData = file_echo_v1_echo_proto_rawDesc
)

func file_echo_v1_echo_proto_rawDescGZIP() []byte {
	file_echo_v1_echo_proto_rawDescOnce.Do(func() {
		file_echo_v1_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_echo_v1_echo_proto_rawDescData)
	})
	return file_echo_v1_echo_proto_rawDescData
}

var file_echo_v1_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_echo_v1_echo_proto_goTypes = []interface{}{
	(*SayHelloRequest)(nil),  // 0: DanielYevelkin.echo.v1.SayHelloRequest
	(*SayHelloResponse)(nil), // 1: DanielYevelkin.echo.v1.SayHelloResponse
}
var file_echo_v1_echo_proto_depIdxs = []int32{
	0, // 0: DanielYevelkin.echo.v1.EchoAPI.SayHello:input_type -> DanielYevelkin.echo.v1.SayHelloRequest
	1, // 1: DanielYevelkin.echo.v1.EchoAPI.SayHello:output_type -> DanielYevelkin.echo.v1.SayHelloResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_echo_v1_echo_proto_init() }
func file_echo_v1_echo_proto_init() {
	if File_echo_v1_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_echo_v1_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloRequest); i {
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
		file_echo_v1_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloResponse); i {
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
			RawDescriptor: file_echo_v1_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_echo_v1_echo_proto_goTypes,
		DependencyIndexes: file_echo_v1_echo_proto_depIdxs,
		MessageInfos:      file_echo_v1_echo_proto_msgTypes,
	}.Build()
	File_echo_v1_echo_proto = out.File
	file_echo_v1_echo_proto_rawDesc = nil
	file_echo_v1_echo_proto_goTypes = nil
	file_echo_v1_echo_proto_depIdxs = nil
}
