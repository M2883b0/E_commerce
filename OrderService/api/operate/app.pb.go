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
	Price uint32 `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
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

func (x *Content) GetPrice() uint32 {
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

// 增加的前端后端返回的数据格式
type CreateContentReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Content       *Content               `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateContentReq) Reset() {
	*x = CreateContentReq{}
	mi := &file_operate_app_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContentReq) ProtoMessage() {}

func (x *CreateContentReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateContentReq.ProtoReflect.Descriptor instead.
func (*CreateContentReq) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{1}
}

func (x *CreateContentReq) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type CreateContentRsp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateContentRsp) Reset() {
	*x = CreateContentRsp{}
	mi := &file_operate_app_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateContentRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContentRsp) ProtoMessage() {}

func (x *CreateContentRsp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateContentRsp.ProtoReflect.Descriptor instead.
func (*CreateContentRsp) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{2}
}

// 更新的前端后端返回的数据格式
type UpdateContentReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Content       *Content               `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateContentReq) Reset() {
	*x = UpdateContentReq{}
	mi := &file_operate_app_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateContentReq) ProtoMessage() {}

func (x *UpdateContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateContentReq.ProtoReflect.Descriptor instead.
func (*UpdateContentReq) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateContentReq) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type UpdateContentRsp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateContentRsp) Reset() {
	*x = UpdateContentRsp{}
	mi := &file_operate_app_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateContentRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateContentRsp) ProtoMessage() {}

func (x *UpdateContentRsp) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateContentRsp.ProtoReflect.Descriptor instead.
func (*UpdateContentRsp) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{4}
}

// 删除的前端后端返回的数据格式
type DeleteContentReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteContentReq) Reset() {
	*x = DeleteContentReq{}
	mi := &file_operate_app_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteContentReq) ProtoMessage() {}

func (x *DeleteContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteContentReq.ProtoReflect.Descriptor instead.
func (*DeleteContentReq) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteContentReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteContentRsp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteContentRsp) Reset() {
	*x = DeleteContentRsp{}
	mi := &file_operate_app_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteContentRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteContentRsp) ProtoMessage() {}

func (x *DeleteContentRsp) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteContentRsp.ProtoReflect.Descriptor instead.
func (*DeleteContentRsp) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{6}
}

// 批量查找的前端后端返回的数据格式
type FindContentReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 搜索内容
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// 页
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	// 页大小
	PageSize      int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindContentReq) Reset() {
	*x = FindContentReq{}
	mi := &file_operate_app_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindContentReq) ProtoMessage() {}

func (x *FindContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindContentReq.ProtoReflect.Descriptor instead.
func (*FindContentReq) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{7}
}

func (x *FindContentReq) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *FindContentReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindContentReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type FindContentRsp struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 内容总数
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// 内容列表
	Contents      []*Content `protobuf:"bytes,2,rep,name=contents,proto3" json:"contents,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindContentRsp) Reset() {
	*x = FindContentRsp{}
	mi := &file_operate_app_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindContentRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindContentRsp) ProtoMessage() {}

func (x *FindContentRsp) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindContentRsp.ProtoReflect.Descriptor instead.
func (*FindContentRsp) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{8}
}

func (x *FindContentRsp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *FindContentRsp) GetContents() []*Content {
	if x != nil {
		return x.Contents
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
	mi := &file_operate_app_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContentReq) ProtoMessage() {}

func (x *GetContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[9]
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
	return file_operate_app_proto_rawDescGZIP(), []int{9}
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
	mi := &file_operate_app_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetContentRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContentRsp) ProtoMessage() {}

func (x *GetContentRsp) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[10]
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
	return file_operate_app_proto_rawDescGZIP(), []int{10}
}

func (x *GetContentRsp) GetContents() *Content {
	if x != nil {
		return x.Contents
	}
	return nil
}

// 商品推送
type RecommendContentReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 用户id
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// 页
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	// 页大小
	PageSize      int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RecommendContentReq) Reset() {
	*x = RecommendContentReq{}
	mi := &file_operate_app_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecommendContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendContentReq) ProtoMessage() {}

func (x *RecommendContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendContentReq.ProtoReflect.Descriptor instead.
func (*RecommendContentReq) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{11}
}

func (x *RecommendContentReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RecommendContentReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *RecommendContentReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type RecommendContentRsp struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 内容总数
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// 内容列表
	Contents      []*Content `protobuf:"bytes,2,rep,name=contents,proto3" json:"contents,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RecommendContentRsp) Reset() {
	*x = RecommendContentRsp{}
	mi := &file_operate_app_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecommendContentRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendContentRsp) ProtoMessage() {}

func (x *RecommendContentRsp) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendContentRsp.ProtoReflect.Descriptor instead.
func (*RecommendContentRsp) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{12}
}

func (x *RecommendContentRsp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *RecommendContentRsp) GetContents() []*Content {
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
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x42, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2e, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x22,
	0x42, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x22,
	0x57, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x58, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x30, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x5f, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x5d, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x30, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xd9, 0x03, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x12, 0x4d,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x12, 0x4d, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x12, 0x4d, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x12, 0x47, 0x0a, 0x0b, 0x46,
	0x69, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x73, 0x70, 0x12, 0x44, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x12, 0x56, 0x0a, 0x10, 0x52, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x20,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x73, 0x70, 0x42, 0x31, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x50, 0x01, 0x5a, 0x20, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x3b, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_operate_app_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_operate_app_proto_goTypes = []any{
	(*Content)(nil),             // 0: api.operate.Content
	(*CreateContentReq)(nil),    // 1: api.operate.CreateContentReq
	(*CreateContentRsp)(nil),    // 2: api.operate.CreateContentRsp
	(*UpdateContentReq)(nil),    // 3: api.operate.UpdateContentReq
	(*UpdateContentRsp)(nil),    // 4: api.operate.UpdateContentRsp
	(*DeleteContentReq)(nil),    // 5: api.operate.DeleteContentReq
	(*DeleteContentRsp)(nil),    // 6: api.operate.DeleteContentRsp
	(*FindContentReq)(nil),      // 7: api.operate.FindContentReq
	(*FindContentRsp)(nil),      // 8: api.operate.FindContentRsp
	(*GetContentReq)(nil),       // 9: api.operate.GetContentReq
	(*GetContentRsp)(nil),       // 10: api.operate.GetContentRsp
	(*RecommendContentReq)(nil), // 11: api.operate.RecommendContentReq
	(*RecommendContentRsp)(nil), // 12: api.operate.RecommendContentRsp
}
var file_operate_app_proto_depIdxs = []int32{
	0,  // 0: api.operate.CreateContentReq.content:type_name -> api.operate.Content
	0,  // 1: api.operate.UpdateContentReq.content:type_name -> api.operate.Content
	0,  // 2: api.operate.FindContentRsp.contents:type_name -> api.operate.Content
	0,  // 3: api.operate.GetContentRsp.contents:type_name -> api.operate.Content
	0,  // 4: api.operate.RecommendContentRsp.contents:type_name -> api.operate.Content
	1,  // 5: api.operate.App.CreateContent:input_type -> api.operate.CreateContentReq
	3,  // 6: api.operate.App.UpdateContent:input_type -> api.operate.UpdateContentReq
	5,  // 7: api.operate.App.DeleteContent:input_type -> api.operate.DeleteContentReq
	7,  // 8: api.operate.App.FindContent:input_type -> api.operate.FindContentReq
	9,  // 9: api.operate.App.GetContent:input_type -> api.operate.GetContentReq
	11, // 10: api.operate.App.RecommendContent:input_type -> api.operate.RecommendContentReq
	2,  // 11: api.operate.App.CreateContent:output_type -> api.operate.CreateContentRsp
	4,  // 12: api.operate.App.UpdateContent:output_type -> api.operate.UpdateContentRsp
	6,  // 13: api.operate.App.DeleteContent:output_type -> api.operate.DeleteContentRsp
	8,  // 14: api.operate.App.FindContent:output_type -> api.operate.FindContentRsp
	10, // 15: api.operate.App.GetContent:output_type -> api.operate.GetContentRsp
	12, // 16: api.operate.App.RecommendContent:output_type -> api.operate.RecommendContentRsp
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
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
			NumMessages:   13,
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
