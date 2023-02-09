// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.12
// source: product.proto

package models

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

// Symbols defined in public import of google/protobuf/timestamp.proto.

type Timestamp = timestamppb.Timestamp

type GImg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link  string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Des   string `protobuf:"bytes,3,opt,name=des,proto3" json:"des,omitempty"`
}

func (x *GImg) Reset() {
	*x = GImg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GImg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GImg) ProtoMessage() {}

func (x *GImg) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GImg.ProtoReflect.Descriptor instead.
func (*GImg) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *GImg) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *GImg) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GImg) GetDes() string {
	if x != nil {
		return x.Des
	}
	return ""
}

type YahooDsc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Des   string `protobuf:"bytes,2,opt,name=des,proto3" json:"des,omitempty"`
}

func (x *YahooDsc) Reset() {
	*x = YahooDsc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YahooDsc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YahooDsc) ProtoMessage() {}

func (x *YahooDsc) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YahooDsc.ProtoReflect.Descriptor instead.
func (*YahooDsc) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *YahooDsc) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *YahooDsc) GetDes() string {
	if x != nil {
		return x.Des
	}
	return ""
}

type VersionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateAt *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
	Name     string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	FileName string                 `protobuf:"bytes,3,opt,name=FileName,proto3" json:"FileName,omitempty"`
	Count    uint64                 `protobuf:"varint,4,opt,name=Count,proto3" json:"Count,omitempty"`
	DownPic  bool                   `protobuf:"varint,5,opt,name=DownPic,proto3" json:"DownPic,omitempty"`
}

func (x *VersionInfo) Reset() {
	*x = VersionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionInfo) ProtoMessage() {}

func (x *VersionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionInfo.ProtoReflect.Descriptor instead.
func (*VersionInfo) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{2}
}

func (x *VersionInfo) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

func (x *VersionInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VersionInfo) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *VersionInfo) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *VersionInfo) GetDownPic() bool {
	if x != nil {
		return x.DownPic
	}
	return false
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`         // @gotags: storm:"id,increment"
	CateId   uint64   `protobuf:"varint,2,opt,name=CateId,proto3" json:"CateId,omitempty"` // @gotags: storm:"index"
	Name     string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Image    string   `protobuf:"bytes,4,opt,name=Image,proto3" json:"Image,omitempty"`
	DomainID uint64   `protobuf:"varint,5,opt,name=DomainID,proto3" json:"DomainID,omitempty"`
	Brand    string   `protobuf:"bytes,6,opt,name=Brand,proto3" json:"Brand,omitempty"`
	Price    string   `protobuf:"bytes,7,opt,name=Price,proto3" json:"Price,omitempty"`
	Specials string   `protobuf:"bytes,8,opt,name=Specials,proto3" json:"Specials,omitempty"`
	Keywords []string `protobuf:"bytes,9,rep,name=Keywords,proto3" json:"Keywords,omitempty"`
	Link     string   `protobuf:"bytes,10,opt,name=Link,proto3" json:"Link,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{3}
}

func (x *Product) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Product) GetCateId() uint64 {
	if x != nil {
		return x.CateId
	}
	return 0
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Product) GetDomainID() uint64 {
	if x != nil {
		return x.DomainID
	}
	return 0
}

func (x *Product) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Product) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Product) GetSpecials() string {
	if x != nil {
		return x.Specials
	}
	return ""
}

func (x *Product) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *Product) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type ProductInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid          string      `protobuf:"bytes,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	Model        string      `protobuf:"bytes,2,opt,name=Model,proto3" json:"Model,omitempty"`
	Images       []string    `protobuf:"bytes,3,rep,name=Images,proto3" json:"Images,omitempty"`
	Description  string      `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	MTitle       string      `protobuf:"bytes,5,opt,name=MTitle,proto3" json:"MTitle,omitempty"`
	MKeywords    string      `protobuf:"bytes,6,opt,name=MKeywords,proto3" json:"MKeywords,omitempty"`
	MDescription string      `protobuf:"bytes,7,opt,name=MDescription,proto3" json:"MDescription,omitempty"`
	GoogleImgs   []*GImg     `protobuf:"bytes,8,rep,name=GoogleImgs,proto3" json:"GoogleImgs,omitempty"`
	YahooDesc    []*YahooDsc `protobuf:"bytes,9,rep,name=YahooDesc,proto3" json:"YahooDesc,omitempty"`
	BingDesc     []*YahooDsc `protobuf:"bytes,10,rep,name=BingDesc,proto3" json:"BingDesc,omitempty"`
	Youtube      []*GImg     `protobuf:"bytes,11,rep,name=Youtube,proto3" json:"Youtube,omitempty"`
	Categories   []string    `protobuf:"bytes,12,rep,name=Categories,proto3" json:"Categories,omitempty"`
}

func (x *ProductInfo) Reset() {
	*x = ProductInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductInfo) ProtoMessage() {}

func (x *ProductInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductInfo.ProtoReflect.Descriptor instead.
func (*ProductInfo) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{4}
}

func (x *ProductInfo) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

func (x *ProductInfo) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *ProductInfo) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *ProductInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductInfo) GetMTitle() string {
	if x != nil {
		return x.MTitle
	}
	return ""
}

func (x *ProductInfo) GetMKeywords() string {
	if x != nil {
		return x.MKeywords
	}
	return ""
}

func (x *ProductInfo) GetMDescription() string {
	if x != nil {
		return x.MDescription
	}
	return ""
}

func (x *ProductInfo) GetGoogleImgs() []*GImg {
	if x != nil {
		return x.GoogleImgs
	}
	return nil
}

func (x *ProductInfo) GetYahooDesc() []*YahooDsc {
	if x != nil {
		return x.YahooDesc
	}
	return nil
}

func (x *ProductInfo) GetBingDesc() []*YahooDsc {
	if x != nil {
		return x.BingDesc
	}
	return nil
}

func (x *ProductInfo) GetYoutube() []*GImg {
	if x != nil {
		return x.Youtube
	}
	return nil
}

func (x *ProductInfo) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x04, 0x47, 0x49, 0x6d, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x65, 0x73, 0x22, 0x32, 0x0a, 0x08,
	0x59, 0x61, 0x68, 0x6f, 0x6f, 0x44, 0x73, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x65, 0x73,
	0x22, 0xa5, 0x01, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x36, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x44, 0x6f, 0x77, 0x6e, 0x50, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x44, 0x6f, 0x77, 0x6e, 0x50, 0x69, 0x63, 0x22, 0xef, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x43, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x4b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x4b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x9d, 0x03, 0x0a, 0x0b, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x50, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x4d, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4d, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x0a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x49, 0x6d, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x47, 0x49, 0x6d, 0x67, 0x52, 0x0a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x49, 0x6d, 0x67, 0x73, 0x12, 0x2e, 0x0a, 0x09, 0x59, 0x61, 0x68, 0x6f, 0x6f, 0x44, 0x65, 0x73,
	0x63, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x59, 0x61, 0x68, 0x6f, 0x6f, 0x44, 0x73, 0x63, 0x52, 0x09, 0x59, 0x61, 0x68, 0x6f, 0x6f,
	0x44, 0x65, 0x73, 0x63, 0x12, 0x2c, 0x0a, 0x08, 0x42, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x73, 0x63,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x59, 0x61, 0x68, 0x6f, 0x6f, 0x44, 0x73, 0x63, 0x52, 0x08, 0x42, 0x69, 0x6e, 0x67, 0x44, 0x65,
	0x73, 0x63, 0x12, 0x26, 0x0a, 0x07, 0x59, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x47, 0x49, 0x6d,
	0x67, 0x52, 0x07, 0x59, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_product_proto_goTypes = []interface{}{
	(*GImg)(nil),                  // 0: models.GImg
	(*YahooDsc)(nil),              // 1: models.YahooDsc
	(*VersionInfo)(nil),           // 2: models.VersionInfo
	(*Product)(nil),               // 3: models.Product
	(*ProductInfo)(nil),           // 4: models.ProductInfo
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_product_proto_depIdxs = []int32{
	5, // 0: models.VersionInfo.CreateAt:type_name -> google.protobuf.Timestamp
	0, // 1: models.ProductInfo.GoogleImgs:type_name -> models.GImg
	1, // 2: models.ProductInfo.YahooDesc:type_name -> models.YahooDsc
	1, // 3: models.ProductInfo.BingDesc:type_name -> models.YahooDsc
	0, // 4: models.ProductInfo.Youtube:type_name -> models.GImg
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GImg); i {
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
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YahooDsc); i {
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
		file_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionInfo); i {
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
		file_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductInfo); i {
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
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}
