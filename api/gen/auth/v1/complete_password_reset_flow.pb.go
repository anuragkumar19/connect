// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: auth/v1/complete_password_reset_flow.proto

package authv1

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

type CompletePasswordResetFlowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CompletePasswordResetFlowRequest) Reset() {
	*x = CompletePasswordResetFlowRequest{}
	mi := &file_auth_v1_complete_password_reset_flow_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompletePasswordResetFlowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletePasswordResetFlowRequest) ProtoMessage() {}

func (x *CompletePasswordResetFlowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_complete_password_reset_flow_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletePasswordResetFlowRequest.ProtoReflect.Descriptor instead.
func (*CompletePasswordResetFlowRequest) Descriptor() ([]byte, []int) {
	return file_auth_v1_complete_password_reset_flow_proto_rawDescGZIP(), []int{0}
}

func (x *CompletePasswordResetFlowRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CompletePasswordResetFlowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CompletePasswordResetFlowResponse) Reset() {
	*x = CompletePasswordResetFlowResponse{}
	mi := &file_auth_v1_complete_password_reset_flow_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompletePasswordResetFlowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletePasswordResetFlowResponse) ProtoMessage() {}

func (x *CompletePasswordResetFlowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_complete_password_reset_flow_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletePasswordResetFlowResponse.ProtoReflect.Descriptor instead.
func (*CompletePasswordResetFlowResponse) Descriptor() ([]byte, []int) {
	return file_auth_v1_complete_password_reset_flow_proto_rawDescGZIP(), []int{1}
}

var File_auth_v1_complete_password_reset_flow_proto protoreflect.FileDescriptor

var file_auth_v1_complete_password_reset_flow_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x22, 0x32, 0x0a, 0x20, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x46, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x21, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa3,
	0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x1e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x75,
	0x72, 0x61, 0x67, 0x6b, 0x75, 0x6d, 0x61, 0x72, 0x31, 0x39, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa,
	0x02, 0x07, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x41, 0x75, 0x74, 0x68,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x41, 0x75, 0x74, 0x68,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_v1_complete_password_reset_flow_proto_rawDescOnce sync.Once
	file_auth_v1_complete_password_reset_flow_proto_rawDescData = file_auth_v1_complete_password_reset_flow_proto_rawDesc
)

func file_auth_v1_complete_password_reset_flow_proto_rawDescGZIP() []byte {
	file_auth_v1_complete_password_reset_flow_proto_rawDescOnce.Do(func() {
		file_auth_v1_complete_password_reset_flow_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_v1_complete_password_reset_flow_proto_rawDescData)
	})
	return file_auth_v1_complete_password_reset_flow_proto_rawDescData
}

var file_auth_v1_complete_password_reset_flow_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_v1_complete_password_reset_flow_proto_goTypes = []any{
	(*CompletePasswordResetFlowRequest)(nil),  // 0: auth.v1.CompletePasswordResetFlowRequest
	(*CompletePasswordResetFlowResponse)(nil), // 1: auth.v1.CompletePasswordResetFlowResponse
}
var file_auth_v1_complete_password_reset_flow_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_v1_complete_password_reset_flow_proto_init() }
func file_auth_v1_complete_password_reset_flow_proto_init() {
	if File_auth_v1_complete_password_reset_flow_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_v1_complete_password_reset_flow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_v1_complete_password_reset_flow_proto_goTypes,
		DependencyIndexes: file_auth_v1_complete_password_reset_flow_proto_depIdxs,
		MessageInfos:      file_auth_v1_complete_password_reset_flow_proto_msgTypes,
	}.Build()
	File_auth_v1_complete_password_reset_flow_proto = out.File
	file_auth_v1_complete_password_reset_flow_proto_rawDesc = nil
	file_auth_v1_complete_password_reset_flow_proto_goTypes = nil
	file_auth_v1_complete_password_reset_flow_proto_depIdxs = nil
}
