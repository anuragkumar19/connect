// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: mailer/v1/mailer.proto

package mailerv1

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

type EmailType int32

const (
	EmailType_EMAIL_TYPE_UNSPECIFIED     EmailType = 0
	EmailType_EMAIL_TYPE_ACCOUNT_RELATED EmailType = 1
	EmailType_EMAIL_TYPE_PROMOTIONAL     EmailType = 2
	EmailType_EMAIL_TYPE_NOTIFICATION    EmailType = 3
)

// Enum value maps for EmailType.
var (
	EmailType_name = map[int32]string{
		0: "EMAIL_TYPE_UNSPECIFIED",
		1: "EMAIL_TYPE_ACCOUNT_RELATED",
		2: "EMAIL_TYPE_PROMOTIONAL",
		3: "EMAIL_TYPE_NOTIFICATION",
	}
	EmailType_value = map[string]int32{
		"EMAIL_TYPE_UNSPECIFIED":     0,
		"EMAIL_TYPE_ACCOUNT_RELATED": 1,
		"EMAIL_TYPE_PROMOTIONAL":     2,
		"EMAIL_TYPE_NOTIFICATION":    3,
	}
)

func (x EmailType) Enum() *EmailType {
	p := new(EmailType)
	*p = x
	return p
}

func (x EmailType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmailType) Descriptor() protoreflect.EnumDescriptor {
	return file_mailer_v1_mailer_proto_enumTypes[0].Descriptor()
}

func (EmailType) Type() protoreflect.EnumType {
	return &file_mailer_v1_mailer_proto_enumTypes[0]
}

func (x EmailType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmailType.Descriptor instead.
func (EmailType) EnumDescriptor() ([]byte, []int) {
	return file_mailer_v1_mailer_proto_rawDescGZIP(), []int{0}
}

type SendEmail struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          EmailType              `protobuf:"varint,1,opt,name=type,proto3,enum=mailer.v1.EmailType" json:"type,omitempty"`
	To            string                 `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Subject       string                 `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Html          string                 `protobuf:"bytes,4,opt,name=html,proto3" json:"html,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendEmail) Reset() {
	*x = SendEmail{}
	mi := &file_mailer_v1_mailer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendEmail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmail) ProtoMessage() {}

func (x *SendEmail) ProtoReflect() protoreflect.Message {
	mi := &file_mailer_v1_mailer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmail.ProtoReflect.Descriptor instead.
func (*SendEmail) Descriptor() ([]byte, []int) {
	return file_mailer_v1_mailer_proto_rawDescGZIP(), []int{0}
}

func (x *SendEmail) GetType() EmailType {
	if x != nil {
		return x.Type
	}
	return EmailType_EMAIL_TYPE_UNSPECIFIED
}

func (x *SendEmail) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SendEmail) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendEmail) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

var File_mailer_v1_mailer_proto protoreflect.FileDescriptor

var file_mailer_v1_mailer_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x69, 0x6c,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x22, 0x73, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x2a, 0x80, 0x01, 0x0a, 0x09, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x1b,
	0x0a, 0x17, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54,
	0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x42, 0x9e, 0x01, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x4d,
	0x61, 0x69, 0x6c, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x75, 0x72, 0x61, 0x67, 0x6b,
	0x75, 0x6d, 0x61, 0x72, 0x31, 0x39, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x3b, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa,
	0x02, 0x09, 0x4d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x4d, 0x61,
	0x69, 0x6c, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x4d, 0x61, 0x69, 0x6c, 0x65, 0x72,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0a, 0x4d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_mailer_v1_mailer_proto_rawDescOnce sync.Once
	file_mailer_v1_mailer_proto_rawDescData []byte
)

func file_mailer_v1_mailer_proto_rawDescGZIP() []byte {
	file_mailer_v1_mailer_proto_rawDescOnce.Do(func() {
		file_mailer_v1_mailer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_mailer_v1_mailer_proto_rawDesc), len(file_mailer_v1_mailer_proto_rawDesc)))
	})
	return file_mailer_v1_mailer_proto_rawDescData
}

var file_mailer_v1_mailer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mailer_v1_mailer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mailer_v1_mailer_proto_goTypes = []any{
	(EmailType)(0),    // 0: mailer.v1.EmailType
	(*SendEmail)(nil), // 1: mailer.v1.SendEmail
}
var file_mailer_v1_mailer_proto_depIdxs = []int32{
	0, // 0: mailer.v1.SendEmail.type:type_name -> mailer.v1.EmailType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mailer_v1_mailer_proto_init() }
func file_mailer_v1_mailer_proto_init() {
	if File_mailer_v1_mailer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_mailer_v1_mailer_proto_rawDesc), len(file_mailer_v1_mailer_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mailer_v1_mailer_proto_goTypes,
		DependencyIndexes: file_mailer_v1_mailer_proto_depIdxs,
		EnumInfos:         file_mailer_v1_mailer_proto_enumTypes,
		MessageInfos:      file_mailer_v1_mailer_proto_msgTypes,
	}.Build()
	File_mailer_v1_mailer_proto = out.File
	file_mailer_v1_mailer_proto_goTypes = nil
	file_mailer_v1_mailer_proto_depIdxs = nil
}
