// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: blocklist/blocklistpb/blocklist.proto

package blocklistpb

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

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocklist_blocklistpb_blocklist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_blocklist_blocklistpb_blocklist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_blocklist_blocklistpb_blocklist_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type ScoreEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId   string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	ANumber   string `protobuf:"bytes,2,opt,name=a_number,json=aNumber,proto3" json:"a_number,omitempty"`
	BNumber   string `protobuf:"bytes,3,opt,name=b_number,json=bNumber,proto3" json:"b_number,omitempty"`
	CallTime  int64  `protobuf:"varint,4,opt,name=call_time,json=callTime,proto3" json:"call_time,omitempty"`
	SipInvite string `protobuf:"bytes,5,opt,name=sip_invite,json=sipInvite,proto3" json:"sip_invite,omitempty"`
}

func (x *ScoreEventRequest) Reset() {
	*x = ScoreEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocklist_blocklistpb_blocklist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScoreEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoreEventRequest) ProtoMessage() {}

func (x *ScoreEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blocklist_blocklistpb_blocklist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScoreEventRequest.ProtoReflect.Descriptor instead.
func (*ScoreEventRequest) Descriptor() ([]byte, []int) {
	return file_blocklist_blocklistpb_blocklist_proto_rawDescGZIP(), []int{1}
}

func (x *ScoreEventRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *ScoreEventRequest) GetANumber() string {
	if x != nil {
		return x.ANumber
	}
	return ""
}

func (x *ScoreEventRequest) GetBNumber() string {
	if x != nil {
		return x.BNumber
	}
	return ""
}

func (x *ScoreEventRequest) GetCallTime() int64 {
	if x != nil {
		return x.CallTime
	}
	return 0
}

func (x *ScoreEventRequest) GetSipInvite() string {
	if x != nil {
		return x.SipInvite
	}
	return ""
}

type ScoreEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId   string    `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	BlockCall bool      `protobuf:"varint,2,opt,name=block_call,json=blockCall,proto3" json:"block_call,omitempty"`
	IsScam    bool      `protobuf:"varint,3,opt,name=is_scam,json=isScam,proto3" json:"is_scam,omitempty"`
	Score     int32     `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
	Metadata  *Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *ScoreEventResponse) Reset() {
	*x = ScoreEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocklist_blocklistpb_blocklist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScoreEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoreEventResponse) ProtoMessage() {}

func (x *ScoreEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blocklist_blocklistpb_blocklist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScoreEventResponse.ProtoReflect.Descriptor instead.
func (*ScoreEventResponse) Descriptor() ([]byte, []int) {
	return file_blocklist_blocklistpb_blocklist_proto_rawDescGZIP(), []int{2}
}

func (x *ScoreEventResponse) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *ScoreEventResponse) GetBlockCall() bool {
	if x != nil {
		return x.BlockCall
	}
	return false
}

func (x *ScoreEventResponse) GetIsScam() bool {
	if x != nil {
		return x.IsScam
	}
	return false
}

func (x *ScoreEventResponse) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *ScoreEventResponse) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_blocklist_blocklistpb_blocklist_proto protoreflect.FileDescriptor

var file_blocklist_blocklistpb_blocklist_proto_rawDesc = []byte{
	0x0a, 0x25, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x6c, 0x69, 0x73, 0x74, 0x70, 0x62, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0x26, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0xa0, 0x01, 0x0a, 0x11, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x69, 0x70, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x70, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x22, 0xae, 0x01,
	0x0a, 0x12, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x73, 0x5f, 0x73, 0x63, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x69, 0x73, 0x53, 0x63, 0x61, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2f, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0xad,
	0x01, 0x0a, 0x10, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x1c, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e,
	0x0a, 0x0b, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1c, 0x2e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0f,
	0x5a, 0x0d, 0x2e, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blocklist_blocklistpb_blocklist_proto_rawDescOnce sync.Once
	file_blocklist_blocklistpb_blocklist_proto_rawDescData = file_blocklist_blocklistpb_blocklist_proto_rawDesc
)

func file_blocklist_blocklistpb_blocklist_proto_rawDescGZIP() []byte {
	file_blocklist_blocklistpb_blocklist_proto_rawDescOnce.Do(func() {
		file_blocklist_blocklistpb_blocklist_proto_rawDescData = protoimpl.X.CompressGZIP(file_blocklist_blocklistpb_blocklist_proto_rawDescData)
	})
	return file_blocklist_blocklistpb_blocklist_proto_rawDescData
}

var file_blocklist_blocklistpb_blocklist_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_blocklist_blocklistpb_blocklist_proto_goTypes = []interface{}{
	(*Metadata)(nil),           // 0: blocklist.metadata
	(*ScoreEventRequest)(nil),  // 1: blocklist.ScoreEventRequest
	(*ScoreEventResponse)(nil), // 2: blocklist.ScoreEventResponse
}
var file_blocklist_blocklistpb_blocklist_proto_depIdxs = []int32{
	0, // 0: blocklist.ScoreEventResponse.metadata:type_name -> blocklist.metadata
	1, // 1: blocklist.BlocklistService.ScoreEvent:input_type -> blocklist.ScoreEventRequest
	1, // 2: blocklist.BlocklistService.ScoreEvents:input_type -> blocklist.ScoreEventRequest
	2, // 3: blocklist.BlocklistService.ScoreEvent:output_type -> blocklist.ScoreEventResponse
	2, // 4: blocklist.BlocklistService.ScoreEvents:output_type -> blocklist.ScoreEventResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_blocklist_blocklistpb_blocklist_proto_init() }
func file_blocklist_blocklistpb_blocklist_proto_init() {
	if File_blocklist_blocklistpb_blocklist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blocklist_blocklistpb_blocklist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_blocklist_blocklistpb_blocklist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScoreEventRequest); i {
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
		file_blocklist_blocklistpb_blocklist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScoreEventResponse); i {
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
			RawDescriptor: file_blocklist_blocklistpb_blocklist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blocklist_blocklistpb_blocklist_proto_goTypes,
		DependencyIndexes: file_blocklist_blocklistpb_blocklist_proto_depIdxs,
		MessageInfos:      file_blocklist_blocklistpb_blocklist_proto_msgTypes,
	}.Build()
	File_blocklist_blocklistpb_blocklist_proto = out.File
	file_blocklist_blocklistpb_blocklist_proto_rawDesc = nil
	file_blocklist_blocklistpb_blocklist_proto_goTypes = nil
	file_blocklist_blocklistpb_blocklist_proto_depIdxs = nil
}
