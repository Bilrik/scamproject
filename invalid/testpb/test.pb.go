// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: Test/testpb/test.proto

package testpb

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

type IsValid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string `protobuf:"bytes,1,opt,name=Number,proto3" json:"Number,omitempty"`
}

func (x *IsValid) Reset() {
	*x = IsValid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Test_testpb_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsValid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsValid) ProtoMessage() {}

func (x *IsValid) ProtoReflect() protoreflect.Message {
	mi := &file_Test_testpb_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsValid.ProtoReflect.Descriptor instead.
func (*IsValid) Descriptor() ([]byte, []int) {
	return file_Test_testpb_test_proto_rawDescGZIP(), []int{0}
}

func (x *IsValid) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

type IsValidrequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid *IsValid `protobuf:"bytes,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *IsValidrequest) Reset() {
	*x = IsValidrequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Test_testpb_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsValidrequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsValidrequest) ProtoMessage() {}

func (x *IsValidrequest) ProtoReflect() protoreflect.Message {
	mi := &file_Test_testpb_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsValidrequest.ProtoReflect.Descriptor instead.
func (*IsValidrequest) Descriptor() ([]byte, []int) {
	return file_Test_testpb_test_proto_rawDescGZIP(), []int{1}
}

func (x *IsValidrequest) GetValid() *IsValid {
	if x != nil {
		return x.Valid
	}
	return nil
}

type IsValidresponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *IsValidresponse) Reset() {
	*x = IsValidresponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Test_testpb_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsValidresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsValidresponse) ProtoMessage() {}

func (x *IsValidresponse) ProtoReflect() protoreflect.Message {
	mi := &file_Test_testpb_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsValidresponse.ProtoReflect.Descriptor instead.
func (*IsValidresponse) Descriptor() ([]byte, []int) {
	return file_Test_testpb_test_proto_rawDescGZIP(), []int{2}
}

func (x *IsValidresponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type IsValidListrequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid *IsValid `protobuf:"bytes,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *IsValidListrequest) Reset() {
	*x = IsValidListrequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Test_testpb_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsValidListrequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsValidListrequest) ProtoMessage() {}

func (x *IsValidListrequest) ProtoReflect() protoreflect.Message {
	mi := &file_Test_testpb_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsValidListrequest.ProtoReflect.Descriptor instead.
func (*IsValidListrequest) Descriptor() ([]byte, []int) {
	return file_Test_testpb_test_proto_rawDescGZIP(), []int{3}
}

func (x *IsValidListrequest) GetValid() *IsValid {
	if x != nil {
		return x.Valid
	}
	return nil
}

type IsValidListresponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *IsValidListresponse) Reset() {
	*x = IsValidListresponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Test_testpb_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsValidListresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsValidListresponse) ProtoMessage() {}

func (x *IsValidListresponse) ProtoReflect() protoreflect.Message {
	mi := &file_Test_testpb_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsValidListresponse.ProtoReflect.Descriptor instead.
func (*IsValidListresponse) Descriptor() ([]byte, []int) {
	return file_Test_testpb_test_proto_rawDescGZIP(), []int{4}
}

func (x *IsValidListresponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

var File_Test_testpb_test_proto protoreflect.FileDescriptor

var file_Test_testpb_test_proto_rawDesc = []byte{
	0x0a, 0x16, 0x54, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x54, 0x65, 0x73, 0x74, 0x22, 0x21,
	0x0a, 0x07, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x35, 0x0a, 0x0e, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x0f, 0x49, 0x73, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x39, 0x0a, 0x12, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e,
	0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x2b,
	0x0a, 0x13, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x32, 0x8b, 0x01, 0x0a, 0x0b,
	0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x12, 0x14, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x73, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x2e, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x18, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x2e, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Test_testpb_test_proto_rawDescOnce sync.Once
	file_Test_testpb_test_proto_rawDescData = file_Test_testpb_test_proto_rawDesc
)

func file_Test_testpb_test_proto_rawDescGZIP() []byte {
	file_Test_testpb_test_proto_rawDescOnce.Do(func() {
		file_Test_testpb_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_Test_testpb_test_proto_rawDescData)
	})
	return file_Test_testpb_test_proto_rawDescData
}

var file_Test_testpb_test_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Test_testpb_test_proto_goTypes = []interface{}{
	(*IsValid)(nil),             // 0: Test.IsValid
	(*IsValidrequest)(nil),      // 1: Test.IsValidrequest
	(*IsValidresponse)(nil),     // 2: Test.IsValidresponse
	(*IsValidListrequest)(nil),  // 3: Test.IsValidListrequest
	(*IsValidListresponse)(nil), // 4: Test.IsValidListresponse
}
var file_Test_testpb_test_proto_depIdxs = []int32{
	0, // 0: Test.IsValidrequest.valid:type_name -> Test.IsValid
	0, // 1: Test.IsValidListrequest.valid:type_name -> Test.IsValid
	1, // 2: Test.TestService.Valid:input_type -> Test.IsValidrequest
	3, // 3: Test.TestService.ValidList:input_type -> Test.IsValidListrequest
	2, // 4: Test.TestService.Valid:output_type -> Test.IsValidresponse
	4, // 5: Test.TestService.ValidList:output_type -> Test.IsValidListresponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Test_testpb_test_proto_init() }
func file_Test_testpb_test_proto_init() {
	if File_Test_testpb_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Test_testpb_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsValid); i {
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
		file_Test_testpb_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsValidrequest); i {
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
		file_Test_testpb_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsValidresponse); i {
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
		file_Test_testpb_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsValidListrequest); i {
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
		file_Test_testpb_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsValidListresponse); i {
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
			RawDescriptor: file_Test_testpb_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Test_testpb_test_proto_goTypes,
		DependencyIndexes: file_Test_testpb_test_proto_depIdxs,
		MessageInfos:      file_Test_testpb_test_proto_msgTypes,
	}.Build()
	File_Test_testpb_test_proto = out.File
	file_Test_testpb_test_proto_rawDesc = nil
	file_Test_testpb_test_proto_goTypes = nil
	file_Test_testpb_test_proto_depIdxs = nil
}
