// Code generated by protoc-gen-go. DO NOT EDIT.
// source: search.proto

package search_proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LanguageType int32

const (
	LanguageType_UNKNOWN LanguageType = 0
	LanguageType_EN      LanguageType = 1
	LanguageType_ZH_CN   LanguageType = 2
)

var LanguageType_name = map[int32]string{
	0: "UNKNOWN",
	1: "EN",
	2: "ZH_CN",
}

var LanguageType_value = map[string]int32{
	"UNKNOWN": 0,
	"EN":      1,
	"ZH_CN":   2,
}

func (x LanguageType) String() string {
	return proto.EnumName(LanguageType_name, int32(x))
}

func (LanguageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{0}
}

type TermInfo struct {
	RawWord              string   `protobuf:"bytes,1,opt,name=rawWord,proto3" json:"rawWord,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TermInfo) Reset()         { *m = TermInfo{} }
func (m *TermInfo) String() string { return proto.CompactTextString(m) }
func (*TermInfo) ProtoMessage()    {}
func (*TermInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{0}
}

func (m *TermInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TermInfo.Unmarshal(m, b)
}
func (m *TermInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TermInfo.Marshal(b, m, deterministic)
}
func (m *TermInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TermInfo.Merge(m, src)
}
func (m *TermInfo) XXX_Size() int {
	return xxx_messageInfo_TermInfo.Size(m)
}
func (m *TermInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TermInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TermInfo proto.InternalMessageInfo

func (m *TermInfo) GetRawWord() string {
	if m != nil {
		return m.RawWord
	}
	return ""
}

type SearchRequest struct {
	Query                string       `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Terms                []*TermInfo  `protobuf:"bytes,2,rep,name=terms,proto3" json:"terms,omitempty"`
	TraceId              string       `protobuf:"bytes,3,opt,name=traceId,proto3" json:"traceId,omitempty"`
	UserIp               string       `protobuf:"bytes,4,opt,name=userIp,proto3" json:"userIp,omitempty"`
	Timestamp            int64        `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Limit                int32        `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Lang                 LanguageType `protobuf:"varint,7,opt,name=lang,proto3,enum=search_proto.LanguageType" json:"lang,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{1}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchRequest) GetTerms() []*TermInfo {
	if m != nil {
		return m.Terms
	}
	return nil
}

func (m *SearchRequest) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *SearchRequest) GetUserIp() string {
	if m != nil {
		return m.UserIp
	}
	return ""
}

func (m *SearchRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *SearchRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchRequest) GetLang() LanguageType {
	if m != nil {
		return m.Lang
	}
	return LanguageType_UNKNOWN
}

type SearchResponseItem struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	SubContext           string   `protobuf:"bytes,2,opt,name=subContext,proto3" json:"subContext,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	RankScore            float32  `protobuf:"fixed32,4,opt,name=rankScore,proto3" json:"rankScore,omitempty"`
	RerankScore          float32  `protobuf:"fixed32,5,opt,name=rerankScore,proto3" json:"rerankScore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchResponseItem) Reset()         { *m = SearchResponseItem{} }
func (m *SearchResponseItem) String() string { return proto.CompactTextString(m) }
func (*SearchResponseItem) ProtoMessage()    {}
func (*SearchResponseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{2}
}

func (m *SearchResponseItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponseItem.Unmarshal(m, b)
}
func (m *SearchResponseItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponseItem.Marshal(b, m, deterministic)
}
func (m *SearchResponseItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponseItem.Merge(m, src)
}
func (m *SearchResponseItem) XXX_Size() int {
	return xxx_messageInfo_SearchResponseItem.Size(m)
}
func (m *SearchResponseItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponseItem.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponseItem proto.InternalMessageInfo

func (m *SearchResponseItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SearchResponseItem) GetSubContext() string {
	if m != nil {
		return m.SubContext
	}
	return ""
}

func (m *SearchResponseItem) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SearchResponseItem) GetRankScore() float32 {
	if m != nil {
		return m.RankScore
	}
	return 0
}

func (m *SearchResponseItem) GetRerankScore() float32 {
	if m != nil {
		return m.RerankScore
	}
	return 0
}

type SearchResponse struct {
	Count                int32                 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	SearchItems          []*SearchResponseItem `protobuf:"bytes,2,rep,name=searchItems,proto3" json:"searchItems,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{3}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *SearchResponse) GetSearchItems() []*SearchResponseItem {
	if m != nil {
		return m.SearchItems
	}
	return nil
}

func init() {
	proto.RegisterEnum("search_proto.LanguageType", LanguageType_name, LanguageType_value)
	proto.RegisterType((*TermInfo)(nil), "search_proto.TermInfo")
	proto.RegisterType((*SearchRequest)(nil), "search_proto.SearchRequest")
	proto.RegisterType((*SearchResponseItem)(nil), "search_proto.SearchResponseItem")
	proto.RegisterType((*SearchResponse)(nil), "search_proto.SearchResponse")
}

func init() { proto.RegisterFile("search.proto", fileDescriptor_453745cff914010e) }

var fileDescriptor_453745cff914010e = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x6a, 0x13, 0x41,
	0x14, 0xc6, 0xdd, 0x4d, 0x37, 0x69, 0x4e, 0x6a, 0x0d, 0x07, 0x29, 0x43, 0x2d, 0xb2, 0x04, 0x2f,
	0x16, 0x91, 0xbd, 0x88, 0x4f, 0xa0, 0x41, 0x30, 0x58, 0xb6, 0x30, 0xad, 0x54, 0xbc, 0x29, 0xd3,
	0xf4, 0x18, 0x57, 0x77, 0x67, 0xb6, 0xf3, 0x07, 0xed, 0xbb, 0xf8, 0x82, 0xbe, 0x85, 0xec, 0xcc,
	0x2e, 0xd9, 0xa0, 0x78, 0x37, 0xbf, 0xef, 0x9c, 0x99, 0xef, 0xfb, 0x92, 0x85, 0x23, 0x43, 0x42,
	0x6f, 0xbe, 0xe6, 0x8d, 0x56, 0x56, 0x61, 0x47, 0x37, 0x9e, 0x16, 0x2f, 0xe0, 0xf0, 0x8a, 0x74,
	0xbd, 0x96, 0x5f, 0x14, 0x32, 0x98, 0x68, 0xf1, 0xe3, 0x5a, 0xe9, 0x3b, 0x16, 0xa5, 0x51, 0x36,
	0xe5, 0x3d, 0x2e, 0x7e, 0x47, 0xf0, 0xf8, 0xd2, 0x5f, 0xe3, 0x74, 0xef, 0xc8, 0x58, 0x7c, 0x0a,
	0xc9, 0xbd, 0x23, 0xfd, 0xd0, 0x6d, 0x06, 0xc0, 0x57, 0x90, 0x58, 0xd2, 0xb5, 0x61, 0x71, 0x3a,
	0xca, 0x66, 0xcb, 0x93, 0x7c, 0xe8, 0x95, 0xf7, 0x46, 0x3c, 0x2c, 0xb5, 0x7e, 0x56, 0x8b, 0x0d,
	0xad, 0xef, 0xd8, 0x28, 0xf8, 0x75, 0x88, 0x27, 0x30, 0x76, 0x86, 0xf4, 0xba, 0x61, 0x07, 0x7e,
	0xd0, 0x11, 0x9e, 0xc1, 0xd4, 0x96, 0x35, 0x19, 0x2b, 0xea, 0x86, 0x25, 0x69, 0x94, 0x8d, 0xf8,
	0x4e, 0x68, 0x33, 0x55, 0x65, 0x5d, 0x5a, 0x36, 0x4e, 0xa3, 0x2c, 0xe1, 0x01, 0x30, 0x87, 0x83,
	0x4a, 0xc8, 0x2d, 0x9b, 0xa4, 0x51, 0x76, 0xbc, 0x3c, 0xdd, 0x8f, 0x74, 0x2e, 0xe4, 0xd6, 0x89,
	0x2d, 0x5d, 0x3d, 0x34, 0xc4, 0xfd, 0xde, 0xe2, 0x57, 0x04, 0xd8, 0x77, 0x35, 0x8d, 0x92, 0x86,
	0xd6, 0x96, 0xea, 0xf6, 0x71, 0x5b, 0xda, 0x8a, 0xfa, 0xc2, 0x1e, 0xf0, 0x39, 0x80, 0x71, 0xb7,
	0x2b, 0x25, 0x2d, 0xfd, 0xb4, 0x2c, 0xf6, 0xa3, 0x81, 0x82, 0x73, 0x18, 0x39, 0x5d, 0x75, 0xf5,
	0xda, 0x63, 0x5b, 0x41, 0x0b, 0xf9, 0xfd, 0x72, 0xa3, 0x34, 0xf9, 0x76, 0x31, 0xdf, 0x09, 0x98,
	0xc2, 0x4c, 0xd3, 0x6e, 0x9e, 0xf8, 0xf9, 0x50, 0x5a, 0x7c, 0x83, 0xe3, 0xfd, 0x74, 0x6d, 0xb2,
	0x8d, 0x72, 0xd2, 0xfa, 0x64, 0x09, 0x0f, 0x80, 0x6f, 0x61, 0x16, 0x9a, 0xb6, 0xe9, 0xfb, 0x3f,
	0x24, 0xdd, 0x6f, 0xff, 0x77, 0x4d, 0x3e, 0xbc, 0xf4, 0x32, 0x87, 0xa3, 0xe1, 0x0f, 0x84, 0x33,
	0x98, 0x7c, 0x2c, 0x3e, 0x14, 0x17, 0xd7, 0xc5, 0xfc, 0x11, 0x8e, 0x21, 0x7e, 0x57, 0xcc, 0x23,
	0x9c, 0x42, 0xf2, 0xf9, 0xfd, 0xcd, 0xaa, 0x98, 0xc7, 0xcb, 0x4f, 0x70, 0xf8, 0x46, 0x87, 0x47,
	0xf1, 0x1c, 0x9e, 0x84, 0xd3, 0x85, 0x5c, 0x55, 0xce, 0x58, 0xd2, 0xf8, 0xec, 0xdf, 0xee, 0xfe,
	0x83, 0x3a, 0x3d, 0xfb, 0x5f, 0xb4, 0xdb, 0xb1, 0x57, 0x5f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xc9, 0xbb, 0xd6, 0x2d, 0xcb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArSearchClient is the client API for ArSearch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArSearchClient interface {
	SearchOnCluster(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type arSearchClient struct {
	cc *grpc.ClientConn
}

func NewArSearchClient(cc *grpc.ClientConn) ArSearchClient {
	return &arSearchClient{cc}
}

func (c *arSearchClient) SearchOnCluster(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/search_proto.ArSearch/SearchOnCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArSearchServer is the server API for ArSearch service.
type ArSearchServer interface {
	SearchOnCluster(context.Context, *SearchRequest) (*SearchResponse, error)
}

// UnimplementedArSearchServer can be embedded to have forward compatible implementations.
type UnimplementedArSearchServer struct {
}

func (*UnimplementedArSearchServer) SearchOnCluster(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchOnCluster not implemented")
}

func RegisterArSearchServer(s *grpc.Server, srv ArSearchServer) {
	s.RegisterService(&_ArSearch_serviceDesc, srv)
}

func _ArSearch_SearchOnCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArSearchServer).SearchOnCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search_proto.ArSearch/SearchOnCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArSearchServer).SearchOnCluster(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArSearch_serviceDesc = grpc.ServiceDesc{
	ServiceName: "search_proto.ArSearch",
	HandlerType: (*ArSearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchOnCluster",
			Handler:    _ArSearch_SearchOnCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search.proto",
}
