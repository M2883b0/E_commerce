// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: operate/ai.proto

package operate

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

type CreateAiRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAiRequest) Reset() {
	*x = CreateAiRequest{}
	mi := &file_operate_ai_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAiRequest) ProtoMessage() {}

func (x *CreateAiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operate_ai_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAiRequest.ProtoReflect.Descriptor instead.
func (*CreateAiRequest) Descriptor() ([]byte, []int) {
	return file_operate_ai_proto_rawDescGZIP(), []int{0}
}

type CreateAiReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAiReply) Reset() {
	*x = CreateAiReply{}
	mi := &file_operate_ai_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAiReply) ProtoMessage() {}

func (x *CreateAiReply) ProtoReflect() protoreflect.Message {
	mi := &file_operate_ai_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAiReply.ProtoReflect.Descriptor instead.
func (*CreateAiReply) Descriptor() ([]byte, []int) {
	return file_operate_ai_proto_rawDescGZIP(), []int{1}
}

type UpdateAiRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAiRequest) Reset() {
	*x = UpdateAiRequest{}
	mi := &file_operate_ai_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAiRequest) ProtoMessage() {}

func (x *UpdateAiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operate_ai_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAiRequest.ProtoReflect.Descriptor instead.
func (*UpdateAiRequest) Descriptor() ([]byte, []int) {
	return file_operate_ai_proto_rawDescGZIP(), []int{2}
}

type UpdateAiReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAiReply) Reset() {
	*x = UpdateAiReply{}
	mi := &file_operate_ai_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAiReply) ProtoMessage() {}

func (x *UpdateAiReply) ProtoReflect() protoreflect.Message {
	mi := &file_operate_ai_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAiReply.ProtoReflect.Descriptor instead.
func (*UpdateAiReply) Descriptor() ([]byte, []int) {
	return file_operate_ai_proto_rawDescGZIP(), []int{3}
}

type DeleteAiRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAiRequest) Reset() {
	*x = DeleteAiRequest{}
	mi := &file_operate_ai_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAiRequest) ProtoMessage() {}

func (x *DeleteAiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operate_ai_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAiRequest.ProtoReflect.Descriptor instead.
func (*DeleteAiRequest) Descriptor() ([]byte, []int) {
	return file_operate_ai_proto_rawDescGZIP(), []int{4}
}

type DeleteAiReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAiReply) Reset() {
	*x = DeleteAiReply{}
	mi := &file_operate_ai_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAiReply) ProtoMessage() {}

func (x *DeleteAiReply) ProtoReflect() protoreflect.Message {
	mi := &file_operate_ai_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAiReply.ProtoReflect.Descriptor instead.
func (*DeleteAiReply) Descriptor() ([]byte, []int) {
	return file_operate_ai_proto_rawDescGZIP(), []int{5}
}

type GetAiRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAiRequest) Reset() {
	*x = GetAiRequest{}
	mi := &file_operate_ai_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAiRequest) ProtoMessage() {}

func (x *GetAiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operate_ai_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAiRequest.ProtoReflect.Descriptor instead.
func (*GetAiRequest) Descriptor() ([]byte, []int) {
	return file_operate_ai_proto_rawDescGZIP(), []int{6}
}

type GetAiReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAiReply) Reset() {
	*x = GetAiReply{}
	mi := &file_operate_ai_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAiReply) ProtoMessage() {}

func (x *GetAiReply) ProtoReflect() protoreflect.Message {
	mi := &file_operate_ai_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAiReply.ProtoReflect.Descriptor instead.
func (*GetAiReply) Descriptor() ([]byte, []int) {
	return file_operate_ai_proto_rawDescGZIP(), []int{7}
}

type ListAiRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAiRequest) Reset() {
	*x = ListAiRequest{}
	mi := &file_operate_ai_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAiRequest) ProtoMessage() {}

func (x *ListAiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operate_ai_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAiRequest.ProtoReflect.Descriptor instead.
func (*ListAiRequest) Descriptor() ([]byte, []int) {
	return file_operate_ai_proto_rawDescGZIP(), []int{8}
}

type ListAiReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAiReply) Reset() {
	*x = ListAiReply{}
	mi := &file_operate_ai_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAiReply) ProtoMessage() {}

func (x *ListAiReply) ProtoReflect() protoreflect.Message {
	mi := &file_operate_ai_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAiReply.ProtoReflect.Descriptor instead.
func (*ListAiReply) Descriptor() ([]byte, []int) {
	return file_operate_ai_proto_rawDescGZIP(), []int{9}
}

var File_operate_ai_proto protoreflect.FileDescriptor

var file_operate_ai_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x22,
	0x11, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x69, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x69, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0e, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x41, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x41, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xd3, 0x02, 0x0a, 0x02, 0x41, 0x69,
	0x12, 0x44, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x69, 0x12, 0x1c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x44, 0x0a, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x69, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x44, 0x0a, 0x08,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x69, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x69, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x69, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x41, 0x69, 0x12, 0x19, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x69, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x3e, 0x0a, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x69, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x69, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x2e, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x01,
	0x5a, 0x1d, 0x61, 0x69, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x3b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_operate_ai_proto_rawDescOnce sync.Once
	file_operate_ai_proto_rawDescData []byte
)

func file_operate_ai_proto_rawDescGZIP() []byte {
	file_operate_ai_proto_rawDescOnce.Do(func() {
		file_operate_ai_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_operate_ai_proto_rawDesc), len(file_operate_ai_proto_rawDesc)))
	})
	return file_operate_ai_proto_rawDescData
}

var file_operate_ai_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_operate_ai_proto_goTypes = []any{
	(*CreateAiRequest)(nil), // 0: api.operate.CreateAiRequest
	(*CreateAiReply)(nil),   // 1: api.operate.CreateAiReply
	(*UpdateAiRequest)(nil), // 2: api.operate.UpdateAiRequest
	(*UpdateAiReply)(nil),   // 3: api.operate.UpdateAiReply
	(*DeleteAiRequest)(nil), // 4: api.operate.DeleteAiRequest
	(*DeleteAiReply)(nil),   // 5: api.operate.DeleteAiReply
	(*GetAiRequest)(nil),    // 6: api.operate.GetAiRequest
	(*GetAiReply)(nil),      // 7: api.operate.GetAiReply
	(*ListAiRequest)(nil),   // 8: api.operate.ListAiRequest
	(*ListAiReply)(nil),     // 9: api.operate.ListAiReply
}
var file_operate_ai_proto_depIdxs = []int32{
	0, // 0: api.operate.Ai.CreateAi:input_type -> api.operate.CreateAiRequest
	2, // 1: api.operate.Ai.UpdateAi:input_type -> api.operate.UpdateAiRequest
	4, // 2: api.operate.Ai.DeleteAi:input_type -> api.operate.DeleteAiRequest
	6, // 3: api.operate.Ai.GetAi:input_type -> api.operate.GetAiRequest
	8, // 4: api.operate.Ai.ListAi:input_type -> api.operate.ListAiRequest
	1, // 5: api.operate.Ai.CreateAi:output_type -> api.operate.CreateAiReply
	3, // 6: api.operate.Ai.UpdateAi:output_type -> api.operate.UpdateAiReply
	5, // 7: api.operate.Ai.DeleteAi:output_type -> api.operate.DeleteAiReply
	7, // 8: api.operate.Ai.GetAi:output_type -> api.operate.GetAiReply
	9, // 9: api.operate.Ai.ListAi:output_type -> api.operate.ListAiReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_operate_ai_proto_init() }
func file_operate_ai_proto_init() {
	if File_operate_ai_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_operate_ai_proto_rawDesc), len(file_operate_ai_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_operate_ai_proto_goTypes,
		DependencyIndexes: file_operate_ai_proto_depIdxs,
		MessageInfos:      file_operate_ai_proto_msgTypes,
	}.Build()
	File_operate_ai_proto = out.File
	file_operate_ai_proto_goTypes = nil
	file_operate_ai_proto_depIdxs = nil
}
