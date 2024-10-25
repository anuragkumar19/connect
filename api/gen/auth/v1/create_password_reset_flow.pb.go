// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: auth/v1/create_password_reset_flow.proto

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

type CreatePasswordResetFlowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdentifierType PasswordResetFlowIdentifierFieldType `protobuf:"varint,1,opt,name=identifier_type,json=identifierType,proto3,enum=auth.v1.PasswordResetFlowIdentifierFieldType" json:"identifier_type,omitempty"`
	Identifier     string                               `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *CreatePasswordResetFlowRequest) Reset() {
	*x = CreatePasswordResetFlowRequest{}
	mi := &file_auth_v1_create_password_reset_flow_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePasswordResetFlowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePasswordResetFlowRequest) ProtoMessage() {}

func (x *CreatePasswordResetFlowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_create_password_reset_flow_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePasswordResetFlowRequest.ProtoReflect.Descriptor instead.
func (*CreatePasswordResetFlowRequest) Descriptor() ([]byte, []int) {
	return file_auth_v1_create_password_reset_flow_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePasswordResetFlowRequest) GetIdentifierType() PasswordResetFlowIdentifierFieldType {
	if x != nil {
		return x.IdentifierType
	}
	return PasswordResetFlowIdentifierFieldType_PASSWORD_RESET_FLOW_IDENTIFIER_FIELD_TYPE_UNSPECIFIED
}

func (x *CreatePasswordResetFlowRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

type CreatePasswordResetFlowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flow  *PasswordResetFlow `protobuf:"bytes,1,opt,name=flow,proto3" json:"flow,omitempty"`
	Token string             `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CreatePasswordResetFlowResponse) Reset() {
	*x = CreatePasswordResetFlowResponse{}
	mi := &file_auth_v1_create_password_reset_flow_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePasswordResetFlowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePasswordResetFlowResponse) ProtoMessage() {}

func (x *CreatePasswordResetFlowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_create_password_reset_flow_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePasswordResetFlowResponse.ProtoReflect.Descriptor instead.
func (*CreatePasswordResetFlowResponse) Descriptor() ([]byte, []int) {
	return file_auth_v1_create_password_reset_flow_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePasswordResetFlowResponse) GetFlow() *PasswordResetFlow {
	if x != nil {
		return x.Flow
	}
	return nil
}

func (x *CreatePasswordResetFlowResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_auth_v1_create_password_reset_flow_proto protoreflect.FileDescriptor

var file_auth_v1_create_password_reset_flow_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x1a, 0x21, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x46, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x56, 0x0a, 0x0f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x22, 0x67, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x04,
	0x66, 0x6c, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0xa1, 0x01, 0x0a, 0x0b, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x1c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x46,
	0x6c, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x75, 0x72, 0x61, 0x67, 0x6b, 0x75, 0x6d,
	0x61, 0x72, 0x31, 0x39, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74,
	0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x41, 0x75, 0x74, 0x68,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13,
	0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_v1_create_password_reset_flow_proto_rawDescOnce sync.Once
	file_auth_v1_create_password_reset_flow_proto_rawDescData = file_auth_v1_create_password_reset_flow_proto_rawDesc
)

func file_auth_v1_create_password_reset_flow_proto_rawDescGZIP() []byte {
	file_auth_v1_create_password_reset_flow_proto_rawDescOnce.Do(func() {
		file_auth_v1_create_password_reset_flow_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_v1_create_password_reset_flow_proto_rawDescData)
	})
	return file_auth_v1_create_password_reset_flow_proto_rawDescData
}

var file_auth_v1_create_password_reset_flow_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_v1_create_password_reset_flow_proto_goTypes = []any{
	(*CreatePasswordResetFlowRequest)(nil),    // 0: auth.v1.CreatePasswordResetFlowRequest
	(*CreatePasswordResetFlowResponse)(nil),   // 1: auth.v1.CreatePasswordResetFlowResponse
	(PasswordResetFlowIdentifierFieldType)(0), // 2: auth.v1.PasswordResetFlowIdentifierFieldType
	(*PasswordResetFlow)(nil),                 // 3: auth.v1.PasswordResetFlow
}
var file_auth_v1_create_password_reset_flow_proto_depIdxs = []int32{
	2, // 0: auth.v1.CreatePasswordResetFlowRequest.identifier_type:type_name -> auth.v1.PasswordResetFlowIdentifierFieldType
	3, // 1: auth.v1.CreatePasswordResetFlowResponse.flow:type_name -> auth.v1.PasswordResetFlow
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_auth_v1_create_password_reset_flow_proto_init() }
func file_auth_v1_create_password_reset_flow_proto_init() {
	if File_auth_v1_create_password_reset_flow_proto != nil {
		return
	}
	file_auth_v1_password_reset_flow_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_v1_create_password_reset_flow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_v1_create_password_reset_flow_proto_goTypes,
		DependencyIndexes: file_auth_v1_create_password_reset_flow_proto_depIdxs,
		MessageInfos:      file_auth_v1_create_password_reset_flow_proto_msgTypes,
	}.Build()
	File_auth_v1_create_password_reset_flow_proto = out.File
	file_auth_v1_create_password_reset_flow_proto_rawDesc = nil
	file_auth_v1_create_password_reset_flow_proto_goTypes = nil
	file_auth_v1_create_password_reset_flow_proto_depIdxs = nil
}
