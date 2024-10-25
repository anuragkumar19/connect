// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: auth/v1/update_registration_flow.proto

package authv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegistrationFlowUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email                       *string `protobuf:"bytes,1,opt,name=email,proto3,oneof" json:"email,omitempty"`
	PhoneNumber                 *string `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3,oneof" json:"phone_number,omitempty"`
	Name                        *string `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Username                    *string `protobuf:"bytes,4,opt,name=username,proto3,oneof" json:"username,omitempty"`
	Password                    *string `protobuf:"bytes,5,opt,name=password,proto3,oneof" json:"password,omitempty"`
	EmailVerificationCode       *string `protobuf:"bytes,6,opt,name=email_verification_code,json=emailVerificationCode,proto3,oneof" json:"email_verification_code,omitempty"`
	PhoneNumberVerificationCode *string `protobuf:"bytes,7,opt,name=phone_number_verification_code,json=phoneNumberVerificationCode,proto3,oneof" json:"phone_number_verification_code,omitempty"`
}

func (x *RegistrationFlowUpdate) Reset() {
	*x = RegistrationFlowUpdate{}
	mi := &file_auth_v1_update_registration_flow_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegistrationFlowUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationFlowUpdate) ProtoMessage() {}

func (x *RegistrationFlowUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_update_registration_flow_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationFlowUpdate.ProtoReflect.Descriptor instead.
func (*RegistrationFlowUpdate) Descriptor() ([]byte, []int) {
	return file_auth_v1_update_registration_flow_proto_rawDescGZIP(), []int{0}
}

func (x *RegistrationFlowUpdate) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *RegistrationFlowUpdate) GetPhoneNumber() string {
	if x != nil && x.PhoneNumber != nil {
		return *x.PhoneNumber
	}
	return ""
}

func (x *RegistrationFlowUpdate) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *RegistrationFlowUpdate) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *RegistrationFlowUpdate) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *RegistrationFlowUpdate) GetEmailVerificationCode() string {
	if x != nil && x.EmailVerificationCode != nil {
		return *x.EmailVerificationCode
	}
	return ""
}

func (x *RegistrationFlowUpdate) GetPhoneNumberVerificationCode() string {
	if x != nil && x.PhoneNumberVerificationCode != nil {
		return *x.PhoneNumberVerificationCode
	}
	return ""
}

type UpdateRegistrationFlowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Update     *RegistrationFlowUpdate `protobuf:"bytes,2,opt,name=update,proto3" json:"update,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask  `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateRegistrationFlowRequest) Reset() {
	*x = UpdateRegistrationFlowRequest{}
	mi := &file_auth_v1_update_registration_flow_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRegistrationFlowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRegistrationFlowRequest) ProtoMessage() {}

func (x *UpdateRegistrationFlowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_update_registration_flow_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRegistrationFlowRequest.ProtoReflect.Descriptor instead.
func (*UpdateRegistrationFlowRequest) Descriptor() ([]byte, []int) {
	return file_auth_v1_update_registration_flow_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRegistrationFlowRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateRegistrationFlowRequest) GetUpdate() *RegistrationFlowUpdate {
	if x != nil {
		return x.Update
	}
	return nil
}

func (x *UpdateRegistrationFlowRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type UpdateRegistrationFlowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flow *RegistrationFlow `protobuf:"bytes,1,opt,name=flow,proto3" json:"flow,omitempty"`
}

func (x *UpdateRegistrationFlowResponse) Reset() {
	*x = UpdateRegistrationFlowResponse{}
	mi := &file_auth_v1_update_registration_flow_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRegistrationFlowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRegistrationFlowResponse) ProtoMessage() {}

func (x *UpdateRegistrationFlowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_update_registration_flow_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRegistrationFlowResponse.ProtoReflect.Descriptor instead.
func (*UpdateRegistrationFlowResponse) Descriptor() ([]byte, []int) {
	return file_auth_v1_update_registration_flow_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRegistrationFlowResponse) GetFlow() *RegistrationFlow {
	if x != nil {
		return x.Flow
	}
	return nil
}

var File_auth_v1_update_registration_flow_proto protoreflect.FileDescriptor

var file_auth_v1_update_registration_flow_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x03, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6c, 0x6f, 0x77, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88,
	0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a,
	0x17, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05,
	0x52, 0x15, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x48, 0x0a, 0x1e, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x06, 0x52, 0x1b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x21,
	0x0a, 0x1f, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0xa5, 0x01, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6c, 0x6f, 0x77, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x4f, 0x0a, 0x1e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x66,
	0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x6c, 0x6f, 0x77, 0x52, 0x04, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0xa0, 0x01, 0x0a, 0x0b, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x1b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6c,
	0x6f, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x75, 0x72, 0x61, 0x67, 0x6b, 0x75, 0x6d, 0x61,
	0x72, 0x31, 0x39, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x41, 0x75, 0x74, 0x68, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x07, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x41,
	0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x08, 0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_v1_update_registration_flow_proto_rawDescOnce sync.Once
	file_auth_v1_update_registration_flow_proto_rawDescData = file_auth_v1_update_registration_flow_proto_rawDesc
)

func file_auth_v1_update_registration_flow_proto_rawDescGZIP() []byte {
	file_auth_v1_update_registration_flow_proto_rawDescOnce.Do(func() {
		file_auth_v1_update_registration_flow_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_v1_update_registration_flow_proto_rawDescData)
	})
	return file_auth_v1_update_registration_flow_proto_rawDescData
}

var file_auth_v1_update_registration_flow_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_auth_v1_update_registration_flow_proto_goTypes = []any{
	(*RegistrationFlowUpdate)(nil),         // 0: auth.v1.RegistrationFlowUpdate
	(*UpdateRegistrationFlowRequest)(nil),  // 1: auth.v1.UpdateRegistrationFlowRequest
	(*UpdateRegistrationFlowResponse)(nil), // 2: auth.v1.UpdateRegistrationFlowResponse
	(*fieldmaskpb.FieldMask)(nil),          // 3: google.protobuf.FieldMask
	(*RegistrationFlow)(nil),               // 4: auth.v1.RegistrationFlow
}
var file_auth_v1_update_registration_flow_proto_depIdxs = []int32{
	0, // 0: auth.v1.UpdateRegistrationFlowRequest.update:type_name -> auth.v1.RegistrationFlowUpdate
	3, // 1: auth.v1.UpdateRegistrationFlowRequest.update_mask:type_name -> google.protobuf.FieldMask
	4, // 2: auth.v1.UpdateRegistrationFlowResponse.flow:type_name -> auth.v1.RegistrationFlow
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_auth_v1_update_registration_flow_proto_init() }
func file_auth_v1_update_registration_flow_proto_init() {
	if File_auth_v1_update_registration_flow_proto != nil {
		return
	}
	file_auth_v1_registration_flow_proto_init()
	file_auth_v1_update_registration_flow_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_v1_update_registration_flow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_v1_update_registration_flow_proto_goTypes,
		DependencyIndexes: file_auth_v1_update_registration_flow_proto_depIdxs,
		MessageInfos:      file_auth_v1_update_registration_flow_proto_msgTypes,
	}.Build()
	File_auth_v1_update_registration_flow_proto = out.File
	file_auth_v1_update_registration_flow_proto_rawDesc = nil
	file_auth_v1_update_registration_flow_proto_goTypes = nil
	file_auth_v1_update_registration_flow_proto_depIdxs = nil
}
