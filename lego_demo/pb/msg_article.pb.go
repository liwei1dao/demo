// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: msg_article.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateArticleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleId    uint32   `protobuf:"varint,1,opt,name=ArticleId,proto3" json:"ArticleId,omitempty"`
	Title        string   `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	ShortContent string   `protobuf:"bytes,3,opt,name=ShortContent,proto3" json:"ShortContent,omitempty"`
	Content      string   `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	Images       []string `protobuf:"bytes,5,rep,name=Images,proto3" json:"Images,omitempty"`
}

func (x *CreateArticleReq) Reset() {
	*x = CreateArticleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleReq) ProtoMessage() {}

func (x *CreateArticleReq) ProtoReflect() protoreflect.Message {
	mi := &file_msg_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleReq.ProtoReflect.Descriptor instead.
func (*CreateArticleReq) Descriptor() ([]byte, []int) {
	return file_msg_article_proto_rawDescGZIP(), []int{0}
}

func (x *CreateArticleReq) GetArticleId() uint32 {
	if x != nil {
		return x.ArticleId
	}
	return 0
}

func (x *CreateArticleReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateArticleReq) GetShortContent() string {
	if x != nil {
		return x.ShortContent
	}
	return ""
}

func (x *CreateArticleReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateArticleReq) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

type DeleteArticleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleId uint32 `protobuf:"varint,1,opt,name=ArticleId,proto3" json:"ArticleId,omitempty"`
}

func (x *DeleteArticleReq) Reset() {
	*x = DeleteArticleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArticleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArticleReq) ProtoMessage() {}

func (x *DeleteArticleReq) ProtoReflect() protoreflect.Message {
	mi := &file_msg_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArticleReq.ProtoReflect.Descriptor instead.
func (*DeleteArticleReq) Descriptor() ([]byte, []int) {
	return file_msg_article_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteArticleReq) GetArticleId() uint32 {
	if x != nil {
		return x.ArticleId
	}
	return 0
}

type GetAuthoIrdArticlesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAuthoIrdArticlesReq) Reset() {
	*x = GetAuthoIrdArticlesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_article_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthoIrdArticlesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthoIrdArticlesReq) ProtoMessage() {}

func (x *GetAuthoIrdArticlesReq) ProtoReflect() protoreflect.Message {
	mi := &file_msg_article_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthoIrdArticlesReq.ProtoReflect.Descriptor instead.
func (*GetAuthoIrdArticlesReq) Descriptor() ([]byte, []int) {
	return file_msg_article_proto_rawDescGZIP(), []int{2}
}

type GetArticleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArticleId uint32 `protobuf:"varint,1,opt,name=ArticleId,proto3" json:"ArticleId,omitempty"`
}

func (x *GetArticleReq) Reset() {
	*x = GetArticleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_article_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleReq) ProtoMessage() {}

func (x *GetArticleReq) ProtoReflect() protoreflect.Message {
	mi := &file_msg_article_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleReq.ProtoReflect.Descriptor instead.
func (*GetArticleReq) Descriptor() ([]byte, []int) {
	return file_msg_article_proto_rawDescGZIP(), []int{3}
}

func (x *GetArticleReq) GetArticleId() uint32 {
	if x != nil {
		return x.ArticleId
	}
	return 0
}

type GetLastArticlesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start int64 `protobuf:"varint,1,opt,name=Start,proto3" json:"Start,omitempty"`
	Num   int64 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (x *GetLastArticlesReq) Reset() {
	*x = GetLastArticlesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_article_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastArticlesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastArticlesReq) ProtoMessage() {}

func (x *GetLastArticlesReq) ProtoReflect() protoreflect.Message {
	mi := &file_msg_article_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastArticlesReq.ProtoReflect.Descriptor instead.
func (*GetLastArticlesReq) Descriptor() ([]byte, []int) {
	return file_msg_article_proto_rawDescGZIP(), []int{4}
}

func (x *GetLastArticlesReq) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *GetLastArticlesReq) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_msg_article_proto protoreflect.FileDescriptor

var file_msg_article_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x73, 0x67, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x22, 0x30, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x49, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x49, 0x72, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x22, 0x2d,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x1c, 0x0a, 0x09, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x3c, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x4e, 0x75, 0x6d, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_article_proto_rawDescOnce sync.Once
	file_msg_article_proto_rawDescData = file_msg_article_proto_rawDesc
)

func file_msg_article_proto_rawDescGZIP() []byte {
	file_msg_article_proto_rawDescOnce.Do(func() {
		file_msg_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_article_proto_rawDescData)
	})
	return file_msg_article_proto_rawDescData
}

var file_msg_article_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_msg_article_proto_goTypes = []interface{}{
	(*CreateArticleReq)(nil),       // 0: CreateArticleReq
	(*DeleteArticleReq)(nil),       // 1: DeleteArticleReq
	(*GetAuthoIrdArticlesReq)(nil), // 2: GetAuthoIrdArticlesReq
	(*GetArticleReq)(nil),          // 3: GetArticleReq
	(*GetLastArticlesReq)(nil),     // 4: GetLastArticlesReq
}
var file_msg_article_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msg_article_proto_init() }
func file_msg_article_proto_init() {
	if File_msg_article_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleReq); i {
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
		file_msg_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArticleReq); i {
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
		file_msg_article_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthoIrdArticlesReq); i {
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
		file_msg_article_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleReq); i {
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
		file_msg_article_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastArticlesReq); i {
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
			RawDescriptor: file_msg_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_article_proto_goTypes,
		DependencyIndexes: file_msg_article_proto_depIdxs,
		MessageInfos:      file_msg_article_proto_msgTypes,
	}.Build()
	File_msg_article_proto = out.File
	file_msg_article_proto_rawDesc = nil
	file_msg_article_proto_goTypes = nil
	file_msg_article_proto_depIdxs = nil
}