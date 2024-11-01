// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: auth/v1/fields.proto

package authv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type VerificationCodeField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoOfTimesCodeSent             int64                  `protobuf:"varint,1,opt,name=no_of_times_code_sent,json=noOfTimesCodeSent,proto3" json:"no_of_times_code_sent,omitempty"`
	LastCodeSentAt                *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_code_sent_at,json=lastCodeSentAt,proto3" json:"last_code_sent_at,omitempty"`
	CanResendAfter                *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=can_resend_after,json=canResendAfter,proto3" json:"can_resend_after,omitempty"`
	IncorrectVerificationAttempts int64                  `protobuf:"varint,4,opt,name=incorrect_verification_attempts,json=incorrectVerificationAttempts,proto3" json:"incorrect_verification_attempts,omitempty"`
	RemainingVerificationAttempts int64                  `protobuf:"varint,5,opt,name=remaining_verification_attempts,json=remainingVerificationAttempts,proto3" json:"remaining_verification_attempts,omitempty"`
	IsVerified                    bool                   `protobuf:"varint,6,opt,name=is_verified,json=isVerified,proto3" json:"is_verified,omitempty"`
	ErrorMessage                  *string                `protobuf:"bytes,7,opt,name=error_message,json=errorMessage,proto3,oneof" json:"error_message,omitempty"`
}

func (x *VerificationCodeField) Reset() {
	*x = VerificationCodeField{}
	mi := &file_auth_v1_fields_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerificationCodeField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerificationCodeField) ProtoMessage() {}

func (x *VerificationCodeField) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_fields_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerificationCodeField.ProtoReflect.Descriptor instead.
func (*VerificationCodeField) Descriptor() ([]byte, []int) {
	return file_auth_v1_fields_proto_rawDescGZIP(), []int{0}
}

func (x *VerificationCodeField) GetNoOfTimesCodeSent() int64 {
	if x != nil {
		return x.NoOfTimesCodeSent
	}
	return 0
}

func (x *VerificationCodeField) GetLastCodeSentAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastCodeSentAt
	}
	return nil
}

func (x *VerificationCodeField) GetCanResendAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.CanResendAfter
	}
	return nil
}

func (x *VerificationCodeField) GetIncorrectVerificationAttempts() int64 {
	if x != nil {
		return x.IncorrectVerificationAttempts
	}
	return 0
}

func (x *VerificationCodeField) GetRemainingVerificationAttempts() int64 {
	if x != nil {
		return x.RemainingVerificationAttempts
	}
	return 0
}

func (x *VerificationCodeField) GetIsVerified() bool {
	if x != nil {
		return x.IsVerified
	}
	return false
}

func (x *VerificationCodeField) GetErrorMessage() string {
	if x != nil && x.ErrorMessage != nil {
		return *x.ErrorMessage
	}
	return ""
}

type VerificationField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsVerified                    bool    `protobuf:"varint,1,opt,name=is_verified,json=isVerified,proto3" json:"is_verified,omitempty"`
	ErrorMessage                  *string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3,oneof" json:"error_message,omitempty"`
	RemainingVerificationAttempts int64   `protobuf:"varint,3,opt,name=remaining_verification_attempts,json=remainingVerificationAttempts,proto3" json:"remaining_verification_attempts,omitempty"`
	IncorrectVerificationAttempts int64   `protobuf:"varint,4,opt,name=incorrect_verification_attempts,json=incorrectVerificationAttempts,proto3" json:"incorrect_verification_attempts,omitempty"`
}

func (x *VerificationField) Reset() {
	*x = VerificationField{}
	mi := &file_auth_v1_fields_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerificationField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerificationField) ProtoMessage() {}

func (x *VerificationField) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_fields_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerificationField.ProtoReflect.Descriptor instead.
func (*VerificationField) Descriptor() ([]byte, []int) {
	return file_auth_v1_fields_proto_rawDescGZIP(), []int{1}
}

func (x *VerificationField) GetIsVerified() bool {
	if x != nil {
		return x.IsVerified
	}
	return false
}

func (x *VerificationField) GetErrorMessage() string {
	if x != nil && x.ErrorMessage != nil {
		return *x.ErrorMessage
	}
	return ""
}

func (x *VerificationField) GetRemainingVerificationAttempts() int64 {
	if x != nil {
		return x.RemainingVerificationAttempts
	}
	return 0
}

func (x *VerificationField) GetIncorrectVerificationAttempts() int64 {
	if x != nil {
		return x.IncorrectVerificationAttempts
	}
	return 0
}

type ValueField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value        string  `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	IsValid      bool    `protobuf:"varint,2,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
	ErrorMessage *string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3,oneof" json:"error_message,omitempty"`
}

func (x *ValueField) Reset() {
	*x = ValueField{}
	mi := &file_auth_v1_fields_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValueField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueField) ProtoMessage() {}

func (x *ValueField) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_fields_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueField.ProtoReflect.Descriptor instead.
func (*ValueField) Descriptor() ([]byte, []int) {
	return file_auth_v1_fields_proto_rawDescGZIP(), []int{2}
}

func (x *ValueField) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ValueField) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

func (x *ValueField) GetErrorMessage() string {
	if x != nil && x.ErrorMessage != nil {
		return *x.ErrorMessage
	}
	return ""
}

type PasswordField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValid      bool    `protobuf:"varint,1,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
	ErrorMessage *string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3,oneof" json:"error_message,omitempty"`
}

func (x *PasswordField) Reset() {
	*x = PasswordField{}
	mi := &file_auth_v1_fields_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PasswordField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordField) ProtoMessage() {}

func (x *PasswordField) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_fields_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordField.ProtoReflect.Descriptor instead.
func (*PasswordField) Descriptor() ([]byte, []int) {
	return file_auth_v1_fields_proto_rawDescGZIP(), []int{3}
}

func (x *PasswordField) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

func (x *PasswordField) GetErrorMessage() string {
	if x != nil && x.ErrorMessage != nil {
		return *x.ErrorMessage
	}
	return ""
}

var File_auth_v1_fields_proto protoreflect.FileDescriptor

var file_auth_v1_fields_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc3, 0x03, 0x0a, 0x15, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x30, 0x0a, 0x15, 0x6e, 0x6f,
	0x5f, 0x6f, 0x66, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6e, 0x6f, 0x4f, 0x66, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x11,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x6e,
	0x74, 0x41, 0x74, 0x12, 0x44, 0x0a, 0x10, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x64, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x63, 0x61, 0x6e, 0x52, 0x65,
	0x73, 0x65, 0x6e, 0x64, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x1f, 0x69, 0x6e, 0x63,
	0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x1d, 0x69, 0x6e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x73, 0x12, 0x46, 0x0a, 0x1f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x74, 0x65,
	0x6d, 0x70, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1d, 0x72, 0x65, 0x6d, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x69, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x0d, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x80, 0x02, 0x0a, 0x11, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x69, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x69, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x28, 0x0a,
	0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x46, 0x0a, 0x1f, 0x72, 0x65, 0x6d, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x1d, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x12,
	0x46, 0x0a, 0x1f, 0x69, 0x6e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1d, 0x69, 0x6e, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x63, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x79, 0x0a, 0x0a, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x66, 0x0a, 0x0d, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x12, 0x28, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x90, 0x01, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x75, 0x72, 0x61, 0x67, 0x6b, 0x75,
	0x6d, 0x61, 0x72, 0x31, 0x39, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75,
	0x74, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x41, 0x75, 0x74,
	0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x13, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_v1_fields_proto_rawDescOnce sync.Once
	file_auth_v1_fields_proto_rawDescData = file_auth_v1_fields_proto_rawDesc
)

func file_auth_v1_fields_proto_rawDescGZIP() []byte {
	file_auth_v1_fields_proto_rawDescOnce.Do(func() {
		file_auth_v1_fields_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_v1_fields_proto_rawDescData)
	})
	return file_auth_v1_fields_proto_rawDescData
}

var file_auth_v1_fields_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_auth_v1_fields_proto_goTypes = []any{
	(*VerificationCodeField)(nil), // 0: auth.v1.VerificationCodeField
	(*VerificationField)(nil),     // 1: auth.v1.VerificationField
	(*ValueField)(nil),            // 2: auth.v1.ValueField
	(*PasswordField)(nil),         // 3: auth.v1.PasswordField
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_auth_v1_fields_proto_depIdxs = []int32{
	4, // 0: auth.v1.VerificationCodeField.last_code_sent_at:type_name -> google.protobuf.Timestamp
	4, // 1: auth.v1.VerificationCodeField.can_resend_after:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_auth_v1_fields_proto_init() }
func file_auth_v1_fields_proto_init() {
	if File_auth_v1_fields_proto != nil {
		return
	}
	file_auth_v1_fields_proto_msgTypes[0].OneofWrappers = []any{}
	file_auth_v1_fields_proto_msgTypes[1].OneofWrappers = []any{}
	file_auth_v1_fields_proto_msgTypes[2].OneofWrappers = []any{}
	file_auth_v1_fields_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_v1_fields_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_v1_fields_proto_goTypes,
		DependencyIndexes: file_auth_v1_fields_proto_depIdxs,
		MessageInfos:      file_auth_v1_fields_proto_msgTypes,
	}.Build()
	File_auth_v1_fields_proto = out.File
	file_auth_v1_fields_proto_rawDesc = nil
	file_auth_v1_fields_proto_goTypes = nil
	file_auth_v1_fields_proto_depIdxs = nil
}
