// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.26.1
// source: operate/app.proto

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

// 数据格式
type Content struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 商品唯一ID
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 商品标题
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// 商品描述
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// 图片url
	PictureUrl string `protobuf:"bytes,4,opt,name=picture_url,json=pictureUrl,proto3" json:"picture_url,omitempty"`
	// 商品价格
	Price float32 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	// 商品库存>=0
	Quantity uint32 `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// 商品分类（可多个分类）
	Categories    []string `protobuf:"bytes,7,rep,name=categories,proto3" json:"categories,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Content) Reset() {
	*x = Content{}
	mi := &file_operate_app_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{0}
}

func (x *Content) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Content) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Content) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Content) GetPictureUrl() string {
	if x != nil {
		return x.PictureUrl
	}
	return ""
}

func (x *Content) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Content) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Content) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

// 精确查找，单个内容
type GetContentReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 商品id
	Id            int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetContentReq) Reset() {
	*x = GetContentReq{}
	mi := &file_operate_app_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContentReq) ProtoMessage() {}

func (x *GetContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContentReq.ProtoReflect.Descriptor instead.
func (*GetContentReq) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{1}
}

func (x *GetContentReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetContentRsp struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 内容
	Contents      *Content `protobuf:"bytes,1,opt,name=contents,proto3" json:"contents,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetContentRsp) Reset() {
	*x = GetContentRsp{}
	mi := &file_operate_app_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetContentRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContentRsp) ProtoMessage() {}

func (x *GetContentRsp) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContentRsp.ProtoReflect.Descriptor instead.
func (*GetContentRsp) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{2}
}

func (x *GetContentRsp) GetContents() *Content {
	if x != nil {
		return x.Contents
	}
	return nil
}

var File_operate_app_proto protoreflect.FileDescriptor

var file_operate_app_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x22, 0xc4, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x08, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x4b, 0x0a, 0x03, 0x41,
	0x70, 0x70, 0x12, 0x44, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x42, 0x30, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x01, 0x5a, 0x1f, 0x43, 0x61, 0x72, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x3b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_operate_app_proto_rawDescOnce sync.Once
	file_operate_app_proto_rawDescData []byte
)

func file_operate_app_proto_rawDescGZIP() []byte {
	file_operate_app_proto_rawDescOnce.Do(func() {
		file_operate_app_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_operate_app_proto_rawDesc), len(file_operate_app_proto_rawDesc)))
	})
	return file_operate_app_proto_rawDescData
}

var file_operate_app_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_operate_app_proto_goTypes = []any{
	(*Content)(nil),       // 0: api.operate.Content
	(*GetContentReq)(nil), // 1: api.operate.GetContentReq
	(*GetContentRsp)(nil), // 2: api.operate.GetContentRsp
}
var file_operate_app_proto_depIdxs = []int32{
	0, // 0: api.operate.GetContentRsp.contents:type_name -> api.operate.Content
	1, // 1: api.operate.App.GetContent:input_type -> api.operate.GetContentReq
	2, // 2: api.operate.App.GetContent:output_type -> api.operate.GetContentRsp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_operate_app_proto_init() }
func file_operate_app_proto_init() {
	if File_operate_app_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_operate_app_proto_rawDesc), len(file_operate_app_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_operate_app_proto_goTypes,
		DependencyIndexes: file_operate_app_proto_depIdxs,
		MessageInfos:      file_operate_app_proto_msgTypes,
	}.Build()
	File_operate_app_proto = out.File
	file_operate_app_proto_goTypes = nil
	file_operate_app_proto_depIdxs = nil
}
