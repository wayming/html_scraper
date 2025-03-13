// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: scrape.proto

package HtmlScraper

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

// 请求消息
type Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HtmlText      string                 `protobuf:"bytes,1,opt,name=html_text,json=htmlText,proto3" json:"html_text,omitempty"` // 客户端发送的 HTML 内容
	PageType      string                 `protobuf:"bytes,2,opt,name=page_type,json=pageType,proto3" json:"page_type,omitempty"` // 数据类型指定，例如 'income_statement', 'balance_sheet', 'cash_flow' 等
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_scrape_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_scrape_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_scrape_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetHtmlText() string {
	if x != nil {
		return x.HtmlText
	}
	return ""
}

func (x *Request) GetPageType() string {
	if x != nil {
		return x.PageType
	}
	return ""
}

// 响应消息
type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`                     // 请求状态，成功或失败
	JsonData      string                 `protobuf:"bytes,2,opt,name=json_data,json=jsonData,proto3" json:"json_data,omitempty"` // 返回的 JSON 数据
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_scrape_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_scrape_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_scrape_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Response) GetJsonData() string {
	if x != nil {
		return x.JsonData
	}
	return ""
}

var File_scrape_proto protoreflect.FileDescriptor

var file_scrape_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x48, 0x74, 0x6d, 0x6c, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x22, 0x43, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x74, 0x6d, 0x6c, 0x5f, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x74, 0x6d, 0x6c, 0x54,
	0x65, 0x78, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x3f, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x32, 0x49, 0x0a, 0x0b, 0x48, 0x74, 0x6d, 0x6c, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72,
	0x12, 0x3a, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x2e, 0x48, 0x74, 0x6d, 0x6c, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x48, 0x74, 0x6d, 0x6c, 0x53, 0x63, 0x72, 0x61,
	0x70, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a, 0x0e,
	0x2e, 0x2f, 0x3b, 0x48, 0x74, 0x6d, 0x6c, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_scrape_proto_rawDescOnce sync.Once
	file_scrape_proto_rawDescData []byte
)

func file_scrape_proto_rawDescGZIP() []byte {
	file_scrape_proto_rawDescOnce.Do(func() {
		file_scrape_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_scrape_proto_rawDesc), len(file_scrape_proto_rawDesc)))
	})
	return file_scrape_proto_rawDescData
}

var file_scrape_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_scrape_proto_goTypes = []any{
	(*Request)(nil),  // 0: HtmlScraper.Request
	(*Response)(nil), // 1: HtmlScraper.Response
}
var file_scrape_proto_depIdxs = []int32{
	0, // 0: HtmlScraper.HtmlScraper.ProcessPage:input_type -> HtmlScraper.Request
	1, // 1: HtmlScraper.HtmlScraper.ProcessPage:output_type -> HtmlScraper.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_scrape_proto_init() }
func file_scrape_proto_init() {
	if File_scrape_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_scrape_proto_rawDesc), len(file_scrape_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scrape_proto_goTypes,
		DependencyIndexes: file_scrape_proto_depIdxs,
		MessageInfos:      file_scrape_proto_msgTypes,
	}.Build()
	File_scrape_proto = out.File
	file_scrape_proto_goTypes = nil
	file_scrape_proto_depIdxs = nil
}
