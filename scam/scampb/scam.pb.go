// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: scam/scampb/scam.proto

package scampb

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

type AddPhoneNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number            string   `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Category_Code     string   `protobuf:"bytes,2,opt,name=category_Code,json=categoryCode,proto3" json:"category_Code,omitempty"`
	Category_Name     string   `protobuf:"bytes,3,opt,name=category_Name,json=categoryName,proto3" json:"category_Name,omitempty"`
	BlockedANumbers   []string `protobuf:"bytes,4,rep,name=BlockedANumbers,proto3" json:"BlockedANumbers,omitempty"`
	BlockedCategories []string `protobuf:"bytes,5,rep,name=BlockedCategories,proto3" json:"BlockedCategories,omitempty"`
}

func (x *AddPhoneNumberRequest) Reset() {
	*x = AddPhoneNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scam_scampb_scam_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPhoneNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPhoneNumberRequest) ProtoMessage() {}

func (x *AddPhoneNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scam_scampb_scam_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPhoneNumberRequest.ProtoReflect.Descriptor instead.
func (*AddPhoneNumberRequest) Descriptor() ([]byte, []int) {
	return file_scam_scampb_scam_proto_rawDescGZIP(), []int{0}
}

func (x *AddPhoneNumberRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *AddPhoneNumberRequest) GetCategory_Code() string {
	if x != nil {
		return x.Category_Code
	}
	return ""
}

func (x *AddPhoneNumberRequest) GetCategory_Name() string {
	if x != nil {
		return x.Category_Name
	}
	return ""
}

func (x *AddPhoneNumberRequest) GetBlockedANumbers() []string {
	if x != nil {
		return x.BlockedANumbers
	}
	return nil
}

func (x *AddPhoneNumberRequest) GetBlockedCategories() []string {
	if x != nil {
		return x.BlockedCategories
	}
	return nil
}

type AddPhoneNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddPhoneNumberResponse) Reset() {
	*x = AddPhoneNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scam_scampb_scam_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPhoneNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPhoneNumberResponse) ProtoMessage() {}

func (x *AddPhoneNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scam_scampb_scam_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPhoneNumberResponse.ProtoReflect.Descriptor instead.
func (*AddPhoneNumberResponse) Descriptor() ([]byte, []int) {
	return file_scam_scampb_scam_proto_rawDescGZIP(), []int{1}
}

func (x *AddPhoneNumberResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetPhoneCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string `protobuf:"bytes,1,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
}

func (x *GetPhoneCategoryRequest) Reset() {
	*x = GetPhoneCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scam_scampb_scam_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPhoneCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPhoneCategoryRequest) ProtoMessage() {}

func (x *GetPhoneCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scam_scampb_scam_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPhoneCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetPhoneCategoryRequest) Descriptor() ([]byte, []int) {
	return file_scam_scampb_scam_proto_rawDescGZIP(), []int{2}
}

func (x *GetPhoneCategoryRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type GetPhoneCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryCode string `protobuf:"bytes,1,opt,name=CategoryCode,proto3" json:"CategoryCode,omitempty"`
	CategoryName string `protobuf:"bytes,2,opt,name=CategoryName,proto3" json:"CategoryName,omitempty"`
}

func (x *GetPhoneCategoryResponse) Reset() {
	*x = GetPhoneCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scam_scampb_scam_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPhoneCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPhoneCategoryResponse) ProtoMessage() {}

func (x *GetPhoneCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scam_scampb_scam_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPhoneCategoryResponse.ProtoReflect.Descriptor instead.
func (*GetPhoneCategoryResponse) Descriptor() ([]byte, []int) {
	return file_scam_scampb_scam_proto_rawDescGZIP(), []int{3}
}

func (x *GetPhoneCategoryResponse) GetCategoryCode() string {
	if x != nil {
		return x.CategoryCode
	}
	return ""
}

func (x *GetPhoneCategoryResponse) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

type GetPersonalBlocklistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string `protobuf:"bytes,1,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
}

func (x *GetPersonalBlocklistRequest) Reset() {
	*x = GetPersonalBlocklistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scam_scampb_scam_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonalBlocklistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonalBlocklistRequest) ProtoMessage() {}

func (x *GetPersonalBlocklistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scam_scampb_scam_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonalBlocklistRequest.ProtoReflect.Descriptor instead.
func (*GetPersonalBlocklistRequest) Descriptor() ([]byte, []int) {
	return file_scam_scampb_scam_proto_rawDescGZIP(), []int{4}
}

func (x *GetPersonalBlocklistRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type GetPersonalBlocklistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockedANumbers   []string `protobuf:"bytes,1,rep,name=BlockedANumbers,proto3" json:"BlockedANumbers,omitempty"`
	BlockedCategories []string `protobuf:"bytes,2,rep,name=BlockedCategories,proto3" json:"BlockedCategories,omitempty"`
}

func (x *GetPersonalBlocklistResponse) Reset() {
	*x = GetPersonalBlocklistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scam_scampb_scam_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonalBlocklistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonalBlocklistResponse) ProtoMessage() {}

func (x *GetPersonalBlocklistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scam_scampb_scam_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonalBlocklistResponse.ProtoReflect.Descriptor instead.
func (*GetPersonalBlocklistResponse) Descriptor() ([]byte, []int) {
	return file_scam_scampb_scam_proto_rawDescGZIP(), []int{5}
}

func (x *GetPersonalBlocklistResponse) GetBlockedANumbers() []string {
	if x != nil {
		return x.BlockedANumbers
	}
	return nil
}

func (x *GetPersonalBlocklistResponse) GetBlockedCategories() []string {
	if x != nil {
		return x.BlockedCategories
	}
	return nil
}

var File_scam_scampb_scam_proto protoreflect.FileDescriptor

var file_scam_scampb_scam_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x63, 0x61, 0x6d, 0x2f, 0x73, 0x63, 0x61, 0x6d, 0x70, 0x62, 0x2f, 0x73, 0x63,
	0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x63, 0x61, 0x6d, 0x22, 0xd1,
	0x01, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x41, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x41, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x11, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x22, 0x32, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x3b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0x62, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x76, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x41, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x41, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x32, 0x8d, 0x02, 0x0a, 0x0c, 0x53, 0x63, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x51, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x2e, 0x73, 0x63, 0x61, 0x6d, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x63, 0x61, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x73,
	0x63, 0x61, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x73, 0x63, 0x61, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x73, 0x63, 0x61, 0x6d, 0x2e, 0x41, 0x64, 0x64,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x63, 0x61, 0x6d, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x73, 0x63, 0x61, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scam_scampb_scam_proto_rawDescOnce sync.Once
	file_scam_scampb_scam_proto_rawDescData = file_scam_scampb_scam_proto_rawDesc
)

func file_scam_scampb_scam_proto_rawDescGZIP() []byte {
	file_scam_scampb_scam_proto_rawDescOnce.Do(func() {
		file_scam_scampb_scam_proto_rawDescData = protoimpl.X.CompressGZIP(file_scam_scampb_scam_proto_rawDescData)
	})
	return file_scam_scampb_scam_proto_rawDescData
}

var file_scam_scampb_scam_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_scam_scampb_scam_proto_goTypes = []interface{}{
	(*AddPhoneNumberRequest)(nil),        // 0: scam.AddPhoneNumberRequest
	(*AddPhoneNumberResponse)(nil),       // 1: scam.AddPhoneNumberResponse
	(*GetPhoneCategoryRequest)(nil),      // 2: scam.GetPhoneCategoryRequest
	(*GetPhoneCategoryResponse)(nil),     // 3: scam.GetPhoneCategoryResponse
	(*GetPersonalBlocklistRequest)(nil),  // 4: scam.GetPersonalBlocklistRequest
	(*GetPersonalBlocklistResponse)(nil), // 5: scam.GetPersonalBlocklistResponse
}
var file_scam_scampb_scam_proto_depIdxs = []int32{
	2, // 0: scam.ScamDatabase.getPhoneCategory:input_type -> scam.GetPhoneCategoryRequest
	4, // 1: scam.ScamDatabase.GetPersonalBlocklist:input_type -> scam.GetPersonalBlocklistRequest
	0, // 2: scam.ScamDatabase.AddPhoneNumber:input_type -> scam.AddPhoneNumberRequest
	3, // 3: scam.ScamDatabase.getPhoneCategory:output_type -> scam.GetPhoneCategoryResponse
	5, // 4: scam.ScamDatabase.GetPersonalBlocklist:output_type -> scam.GetPersonalBlocklistResponse
	1, // 5: scam.ScamDatabase.AddPhoneNumber:output_type -> scam.AddPhoneNumberResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_scam_scampb_scam_proto_init() }
func file_scam_scampb_scam_proto_init() {
	if File_scam_scampb_scam_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scam_scampb_scam_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPhoneNumberRequest); i {
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
		file_scam_scampb_scam_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPhoneNumberResponse); i {
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
		file_scam_scampb_scam_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPhoneCategoryRequest); i {
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
		file_scam_scampb_scam_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPhoneCategoryResponse); i {
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
		file_scam_scampb_scam_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonalBlocklistRequest); i {
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
		file_scam_scampb_scam_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonalBlocklistResponse); i {
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
			RawDescriptor: file_scam_scampb_scam_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scam_scampb_scam_proto_goTypes,
		DependencyIndexes: file_scam_scampb_scam_proto_depIdxs,
		MessageInfos:      file_scam_scampb_scam_proto_msgTypes,
	}.Build()
	File_scam_scampb_scam_proto = out.File
	file_scam_scampb_scam_proto_rawDesc = nil
	file_scam_scampb_scam_proto_goTypes = nil
	file_scam_scampb_scam_proto_depIdxs = nil
}
