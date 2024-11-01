// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: auth/v1/password_reset_flow.proto

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

type PasswordResetFlowIdentifierFieldType int32

const (
	PasswordResetFlowIdentifierFieldType_PASSWORD_RESET_FLOW_IDENTIFIER_FIELD_TYPE_UNSPECIFIED  PasswordResetFlowIdentifierFieldType = 0
	PasswordResetFlowIdentifierFieldType_PASSWORD_RESET_FLOW_IDENTIFIER_FIELD_TYPE_EMAIL        PasswordResetFlowIdentifierFieldType = 2
	PasswordResetFlowIdentifierFieldType_PASSWORD_RESET_FLOW_IDENTIFIER_FIELD_TYPE_PHONE_NUMBER PasswordResetFlowIdentifierFieldType = 3
)

// Enum value maps for PasswordResetFlowIdentifierFieldType.
var (
	PasswordResetFlowIdentifierFieldType_name = map[int32]string{
		0: "PASSWORD_RESET_FLOW_IDENTIFIER_FIELD_TYPE_UNSPECIFIED",
		2: "PASSWORD_RESET_FLOW_IDENTIFIER_FIELD_TYPE_EMAIL",
		3: "PASSWORD_RESET_FLOW_IDENTIFIER_FIELD_TYPE_PHONE_NUMBER",
	}
	PasswordResetFlowIdentifierFieldType_value = map[string]int32{
		"PASSWORD_RESET_FLOW_IDENTIFIER_FIELD_TYPE_UNSPECIFIED":  0,
		"PASSWORD_RESET_FLOW_IDENTIFIER_FIELD_TYPE_EMAIL":        2,
		"PASSWORD_RESET_FLOW_IDENTIFIER_FIELD_TYPE_PHONE_NUMBER": 3,
	}
)

func (x PasswordResetFlowIdentifierFieldType) Enum() *PasswordResetFlowIdentifierFieldType {
	p := new(PasswordResetFlowIdentifierFieldType)
	*p = x
	return p
}

func (x PasswordResetFlowIdentifierFieldType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PasswordResetFlowIdentifierFieldType) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_v1_password_reset_flow_proto_enumTypes[0].Descriptor()
}

func (PasswordResetFlowIdentifierFieldType) Type() protoreflect.EnumType {
	return &file_auth_v1_password_reset_flow_proto_enumTypes[0]
}

func (x PasswordResetFlowIdentifierFieldType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PasswordResetFlowIdentifierFieldType.Descriptor instead.
func (PasswordResetFlowIdentifierFieldType) EnumDescriptor() ([]byte, []int) {
	return file_auth_v1_password_reset_flow_proto_rawDescGZIP(), []int{0}
}

type PasswordResetFlowState int32

const (
	PasswordResetFlowState_PASSWORD_RESET_FLOW_STATE_UNSPECIFIED            PasswordResetFlowState = 0
	PasswordResetFlowState_PASSWORD_RESET_FLOW_STATE_INVALID_FIELDS         PasswordResetFlowState = 1
	PasswordResetFlowState_PASSWORD_RESET_FLOW_STATE_VERIFICATION_CODE_SENT PasswordResetFlowState = 2
	PasswordResetFlowState_PASSWORD_RESET_FLOW_STATE_MFA_REQUIRED           PasswordResetFlowState = 3
	PasswordResetFlowState_PASSWORD_RESET_FLOW_STATE_MFA_METHOD_SELECTED    PasswordResetFlowState = 4
	PasswordResetFlowState_PASSWORD_RESET_FLOW_STATE_AUTHENTICATED          PasswordResetFlowState = 5
	PasswordResetFlowState_PASSWORD_RESET_FLOW_STATE_COMPLETED              PasswordResetFlowState = 6
)

// Enum value maps for PasswordResetFlowState.
var (
	PasswordResetFlowState_name = map[int32]string{
		0: "PASSWORD_RESET_FLOW_STATE_UNSPECIFIED",
		1: "PASSWORD_RESET_FLOW_STATE_INVALID_FIELDS",
		2: "PASSWORD_RESET_FLOW_STATE_VERIFICATION_CODE_SENT",
		3: "PASSWORD_RESET_FLOW_STATE_MFA_REQUIRED",
		4: "PASSWORD_RESET_FLOW_STATE_MFA_METHOD_SELECTED",
		5: "PASSWORD_RESET_FLOW_STATE_AUTHENTICATED",
		6: "PASSWORD_RESET_FLOW_STATE_COMPLETED",
	}
	PasswordResetFlowState_value = map[string]int32{
		"PASSWORD_RESET_FLOW_STATE_UNSPECIFIED":            0,
		"PASSWORD_RESET_FLOW_STATE_INVALID_FIELDS":         1,
		"PASSWORD_RESET_FLOW_STATE_VERIFICATION_CODE_SENT": 2,
		"PASSWORD_RESET_FLOW_STATE_MFA_REQUIRED":           3,
		"PASSWORD_RESET_FLOW_STATE_MFA_METHOD_SELECTED":    4,
		"PASSWORD_RESET_FLOW_STATE_AUTHENTICATED":          5,
		"PASSWORD_RESET_FLOW_STATE_COMPLETED":              6,
	}
)

func (x PasswordResetFlowState) Enum() *PasswordResetFlowState {
	p := new(PasswordResetFlowState)
	*p = x
	return p
}

func (x PasswordResetFlowState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PasswordResetFlowState) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_v1_password_reset_flow_proto_enumTypes[1].Descriptor()
}

func (PasswordResetFlowState) Type() protoreflect.EnumType {
	return &file_auth_v1_password_reset_flow_proto_enumTypes[1]
}

func (x PasswordResetFlowState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PasswordResetFlowState.Descriptor instead.
func (PasswordResetFlowState) EnumDescriptor() ([]byte, []int) {
	return file_auth_v1_password_reset_flow_proto_rawDescGZIP(), []int{1}
}

type PasswordResetFlowFields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdentifierType  PasswordResetFlowIdentifierFieldType `protobuf:"varint,1,opt,name=identifier_type,json=identifierType,proto3,enum=auth.v1.PasswordResetFlowIdentifierFieldType" json:"identifier_type,omitempty"`
	Identifier      *ValueField                          `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	NewPassword     *PasswordField                       `protobuf:"bytes,3,opt,name=new_password,json=newPassword,proto3,oneof" json:"new_password,omitempty"`
	EmailCode       *VerificationCodeField               `protobuf:"bytes,4,opt,name=email_code,json=emailCode,proto3,oneof" json:"email_code,omitempty"`
	PhoneNumberCode *VerificationCodeField               `protobuf:"bytes,5,opt,name=phone_number_code,json=phoneNumberCode,proto3,oneof" json:"phone_number_code,omitempty"`
	Totp            *VerificationField                   `protobuf:"bytes,6,opt,name=totp,proto3,oneof" json:"totp,omitempty"`
	Webauthn        *VerificationField                   `protobuf:"bytes,7,opt,name=webauthn,proto3,oneof" json:"webauthn,omitempty"`
	RecoveryCode    *VerificationField                   `protobuf:"bytes,8,opt,name=recovery_code,json=recoveryCode,proto3,oneof" json:"recovery_code,omitempty"`
}

func (x *PasswordResetFlowFields) Reset() {
	*x = PasswordResetFlowFields{}
	mi := &file_auth_v1_password_reset_flow_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PasswordResetFlowFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordResetFlowFields) ProtoMessage() {}

func (x *PasswordResetFlowFields) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_password_reset_flow_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordResetFlowFields.ProtoReflect.Descriptor instead.
func (*PasswordResetFlowFields) Descriptor() ([]byte, []int) {
	return file_auth_v1_password_reset_flow_proto_rawDescGZIP(), []int{0}
}

func (x *PasswordResetFlowFields) GetIdentifierType() PasswordResetFlowIdentifierFieldType {
	if x != nil {
		return x.IdentifierType
	}
	return PasswordResetFlowIdentifierFieldType_PASSWORD_RESET_FLOW_IDENTIFIER_FIELD_TYPE_UNSPECIFIED
}

func (x *PasswordResetFlowFields) GetIdentifier() *ValueField {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (x *PasswordResetFlowFields) GetNewPassword() *PasswordField {
	if x != nil {
		return x.NewPassword
	}
	return nil
}

func (x *PasswordResetFlowFields) GetEmailCode() *VerificationCodeField {
	if x != nil {
		return x.EmailCode
	}
	return nil
}

func (x *PasswordResetFlowFields) GetPhoneNumberCode() *VerificationCodeField {
	if x != nil {
		return x.PhoneNumberCode
	}
	return nil
}

func (x *PasswordResetFlowFields) GetTotp() *VerificationField {
	if x != nil {
		return x.Totp
	}
	return nil
}

func (x *PasswordResetFlowFields) GetWebauthn() *VerificationField {
	if x != nil {
		return x.Webauthn
	}
	return nil
}

func (x *PasswordResetFlowFields) GetRecoveryCode() *VerificationField {
	if x != nil {
		return x.RecoveryCode
	}
	return nil
}

type PasswordResetFlow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  int64                    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt           *timestamppb.Timestamp   `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt           *timestamppb.Timestamp   `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ExpireAt            *timestamppb.Timestamp   `protobuf:"bytes,5,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
	CsrfToken           string                   `protobuf:"bytes,6,opt,name=csrf_token,json=csrfToken,proto3" json:"csrf_token,omitempty"`
	Fields              *PasswordResetFlowFields `protobuf:"bytes,7,opt,name=fields,proto3" json:"fields,omitempty"`
	State               PasswordResetFlowState   `protobuf:"varint,8,opt,name=state,proto3,enum=auth.v1.PasswordResetFlowState" json:"state,omitempty"`
	AvailableMfaMethods *MFAType                 `protobuf:"varint,9,opt,name=available_mfa_methods,json=availableMfaMethods,proto3,enum=auth.v1.MFAType,oneof" json:"available_mfa_methods,omitempty"`
	SelectedMfaMethod   *MFAType                 `protobuf:"varint,10,opt,name=selected_mfa_method,json=selectedMfaMethod,proto3,enum=auth.v1.MFAType,oneof" json:"selected_mfa_method,omitempty"`
	JsonWebauthnOptions *string                  `protobuf:"bytes,11,opt,name=json_webauthn_options,json=jsonWebauthnOptions,proto3,oneof" json:"json_webauthn_options,omitempty"`
}

func (x *PasswordResetFlow) Reset() {
	*x = PasswordResetFlow{}
	mi := &file_auth_v1_password_reset_flow_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PasswordResetFlow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordResetFlow) ProtoMessage() {}

func (x *PasswordResetFlow) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_password_reset_flow_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordResetFlow.ProtoReflect.Descriptor instead.
func (*PasswordResetFlow) Descriptor() ([]byte, []int) {
	return file_auth_v1_password_reset_flow_proto_rawDescGZIP(), []int{1}
}

func (x *PasswordResetFlow) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PasswordResetFlow) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PasswordResetFlow) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *PasswordResetFlow) GetExpireAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireAt
	}
	return nil
}

func (x *PasswordResetFlow) GetCsrfToken() string {
	if x != nil {
		return x.CsrfToken
	}
	return ""
}

func (x *PasswordResetFlow) GetFields() *PasswordResetFlowFields {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *PasswordResetFlow) GetState() PasswordResetFlowState {
	if x != nil {
		return x.State
	}
	return PasswordResetFlowState_PASSWORD_RESET_FLOW_STATE_UNSPECIFIED
}

func (x *PasswordResetFlow) GetAvailableMfaMethods() MFAType {
	if x != nil && x.AvailableMfaMethods != nil {
		return *x.AvailableMfaMethods
	}
	return MFAType_MFA_TYPE_UNSPECIFIED
}

func (x *PasswordResetFlow) GetSelectedMfaMethod() MFAType {
	if x != nil && x.SelectedMfaMethod != nil {
		return *x.SelectedMfaMethod
	}
	return MFAType_MFA_TYPE_UNSPECIFIED
}

func (x *PasswordResetFlow) GetJsonWebauthnOptions() string {
	if x != nil && x.JsonWebauthnOptions != nil {
		return *x.JsonWebauthnOptions
	}
	return ""
}

var File_auth_v1_password_reset_flow_proto protoreflect.FileDescriptor

var file_auth_v1_password_reset_flow_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x66, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x05, 0x0a, 0x17, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x12, 0x56, 0x0a, 0x0f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x3e, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x00, 0x52,
	0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x42, 0x0a, 0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x48, 0x01, 0x52, 0x09, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x4f, 0x0a, 0x11, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x02,
	0x52, 0x0f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x6f, 0x74, 0x70, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x03,
	0x52, 0x04, 0x74, 0x6f, 0x74, 0x70, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x08, 0x77, 0x65, 0x62,
	0x61, 0x75, 0x74, 0x68, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x04, 0x52, 0x08, 0x77, 0x65, 0x62, 0x61, 0x75,
	0x74, 0x68, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x44, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d,
	0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x14, 0x0a, 0x12,
	0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x6f, 0x74, 0x70, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x77, 0x65, 0x62, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xf9, 0x04, 0x0a, 0x11, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x73, 0x72, 0x66, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x73, 0x72, 0x66, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x38,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x35, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x46,
	0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x49, 0x0a, 0x15, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x66, 0x61,
	0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x46, 0x41, 0x54, 0x79, 0x70, 0x65,
	0x48, 0x00, 0x52, 0x13, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x66, 0x61,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x13, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x66, 0x61, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x46, 0x41, 0x54, 0x79, 0x70, 0x65, 0x48, 0x01, 0x52, 0x11, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x4d, 0x66, 0x61, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x37, 0x0a, 0x15, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x77, 0x65, 0x62, 0x61, 0x75, 0x74,
	0x68, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x13, 0x6a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x61, 0x75, 0x74, 0x68, 0x6e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x66, 0x61, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x6d, 0x66, 0x61, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x18, 0x0a, 0x16,
	0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x77, 0x65, 0x62, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0xd2, 0x01, 0x0a, 0x24, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x39, 0x0a, 0x35, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x45,
	0x54, 0x5f, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45,
	0x52, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x33, 0x0a, 0x2f, 0x50, 0x41,
	0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x46, 0x4c, 0x4f,
	0x57, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x46, 0x49, 0x45,
	0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x02, 0x12,
	0x3a, 0x0a, 0x36, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x45,
	0x54, 0x5f, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45,
	0x52, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x48, 0x4f,
	0x4e, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x03, 0x2a, 0xdc, 0x02, 0x0a, 0x16,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x46, 0x6c, 0x6f,
	0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x25, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f,
	0x52, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x2c, 0x0a, 0x28, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x52, 0x45,
	0x53, 0x45, 0x54, 0x5f, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x53, 0x10, 0x01, 0x12,
	0x34, 0x0a, 0x30, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x45,
	0x54, 0x5f, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x56, 0x45, 0x52,
	0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53,
	0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x2a, 0x0a, 0x26, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52,
	0x44, 0x5f, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x4d, 0x46, 0x41, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x31, 0x0a, 0x2d, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x52, 0x45,
	0x53, 0x45, 0x54, 0x5f, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4d,
	0x46, 0x41, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54,
	0x45, 0x44, 0x10, 0x04, 0x12, 0x2b, 0x0a, 0x27, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44,
	0x5f, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x05, 0x12, 0x27, 0x0a, 0x23, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x52, 0x45,
	0x53, 0x45, 0x54, 0x5f, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x06, 0x42, 0x9b, 0x01, 0x0a, 0x0b, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x6e, 0x75, 0x72, 0x61, 0x67, 0x6b, 0x75, 0x6d, 0x61, 0x72, 0x31, 0x39, 0x2f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x41, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07,
	0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08,
	0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_v1_password_reset_flow_proto_rawDescOnce sync.Once
	file_auth_v1_password_reset_flow_proto_rawDescData = file_auth_v1_password_reset_flow_proto_rawDesc
)

func file_auth_v1_password_reset_flow_proto_rawDescGZIP() []byte {
	file_auth_v1_password_reset_flow_proto_rawDescOnce.Do(func() {
		file_auth_v1_password_reset_flow_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_v1_password_reset_flow_proto_rawDescData)
	})
	return file_auth_v1_password_reset_flow_proto_rawDescData
}

var file_auth_v1_password_reset_flow_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_auth_v1_password_reset_flow_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_v1_password_reset_flow_proto_goTypes = []any{
	(PasswordResetFlowIdentifierFieldType)(0), // 0: auth.v1.PasswordResetFlowIdentifierFieldType
	(PasswordResetFlowState)(0),               // 1: auth.v1.PasswordResetFlowState
	(*PasswordResetFlowFields)(nil),           // 2: auth.v1.PasswordResetFlowFields
	(*PasswordResetFlow)(nil),                 // 3: auth.v1.PasswordResetFlow
	(*ValueField)(nil),                        // 4: auth.v1.ValueField
	(*PasswordField)(nil),                     // 5: auth.v1.PasswordField
	(*VerificationCodeField)(nil),             // 6: auth.v1.VerificationCodeField
	(*VerificationField)(nil),                 // 7: auth.v1.VerificationField
	(*timestamppb.Timestamp)(nil),             // 8: google.protobuf.Timestamp
	(MFAType)(0),                              // 9: auth.v1.MFAType
}
var file_auth_v1_password_reset_flow_proto_depIdxs = []int32{
	0,  // 0: auth.v1.PasswordResetFlowFields.identifier_type:type_name -> auth.v1.PasswordResetFlowIdentifierFieldType
	4,  // 1: auth.v1.PasswordResetFlowFields.identifier:type_name -> auth.v1.ValueField
	5,  // 2: auth.v1.PasswordResetFlowFields.new_password:type_name -> auth.v1.PasswordField
	6,  // 3: auth.v1.PasswordResetFlowFields.email_code:type_name -> auth.v1.VerificationCodeField
	6,  // 4: auth.v1.PasswordResetFlowFields.phone_number_code:type_name -> auth.v1.VerificationCodeField
	7,  // 5: auth.v1.PasswordResetFlowFields.totp:type_name -> auth.v1.VerificationField
	7,  // 6: auth.v1.PasswordResetFlowFields.webauthn:type_name -> auth.v1.VerificationField
	7,  // 7: auth.v1.PasswordResetFlowFields.recovery_code:type_name -> auth.v1.VerificationField
	8,  // 8: auth.v1.PasswordResetFlow.created_at:type_name -> google.protobuf.Timestamp
	8,  // 9: auth.v1.PasswordResetFlow.updated_at:type_name -> google.protobuf.Timestamp
	8,  // 10: auth.v1.PasswordResetFlow.expire_at:type_name -> google.protobuf.Timestamp
	2,  // 11: auth.v1.PasswordResetFlow.fields:type_name -> auth.v1.PasswordResetFlowFields
	1,  // 12: auth.v1.PasswordResetFlow.state:type_name -> auth.v1.PasswordResetFlowState
	9,  // 13: auth.v1.PasswordResetFlow.available_mfa_methods:type_name -> auth.v1.MFAType
	9,  // 14: auth.v1.PasswordResetFlow.selected_mfa_method:type_name -> auth.v1.MFAType
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_auth_v1_password_reset_flow_proto_init() }
func file_auth_v1_password_reset_flow_proto_init() {
	if File_auth_v1_password_reset_flow_proto != nil {
		return
	}
	file_auth_v1_fields_proto_init()
	file_auth_v1_mfa_proto_init()
	file_auth_v1_password_reset_flow_proto_msgTypes[0].OneofWrappers = []any{}
	file_auth_v1_password_reset_flow_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_v1_password_reset_flow_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_v1_password_reset_flow_proto_goTypes,
		DependencyIndexes: file_auth_v1_password_reset_flow_proto_depIdxs,
		EnumInfos:         file_auth_v1_password_reset_flow_proto_enumTypes,
		MessageInfos:      file_auth_v1_password_reset_flow_proto_msgTypes,
	}.Build()
	File_auth_v1_password_reset_flow_proto = out.File
	file_auth_v1_password_reset_flow_proto_rawDesc = nil
	file_auth_v1_password_reset_flow_proto_goTypes = nil
	file_auth_v1_password_reset_flow_proto_depIdxs = nil
}
