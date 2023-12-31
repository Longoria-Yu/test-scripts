// Code generated by protoc-gen-go. DO NOT EDIT.
// source: content-feed/cgproduct.proto

package ContentFeedService_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type SearchCGProductListingsRequestV1_Sort int32

const (
	// SORT_TYPE_UNSPECIFIED will be SORT_TYPE_PRICE_LOW_TO_HIGH for now
	SearchCGProductListingsRequestV1_SORT_TYPE_UNSPECIFIED          SearchCGProductListingsRequestV1_Sort = 0
	SearchCGProductListingsRequestV1_SORT_TYPE_PRICE_LOW_TO_HIGH    SearchCGProductListingsRequestV1_Sort = 1
	SearchCGProductListingsRequestV1_SORT_TYPE_RECOMMENDED          SearchCGProductListingsRequestV1_Sort = 2
	SearchCGProductListingsRequestV1_SORT_TYPE_RECOMMENDED_SAMPLING SearchCGProductListingsRequestV1_Sort = 3
)

var SearchCGProductListingsRequestV1_Sort_name = map[int32]string{
	0: "SORT_TYPE_UNSPECIFIED",
	1: "SORT_TYPE_PRICE_LOW_TO_HIGH",
	2: "SORT_TYPE_RECOMMENDED",
	3: "SORT_TYPE_RECOMMENDED_SAMPLING",
}

var SearchCGProductListingsRequestV1_Sort_value = map[string]int32{
	"SORT_TYPE_UNSPECIFIED":          0,
	"SORT_TYPE_PRICE_LOW_TO_HIGH":    1,
	"SORT_TYPE_RECOMMENDED":          2,
	"SORT_TYPE_RECOMMENDED_SAMPLING": 3,
}

func (x SearchCGProductListingsRequestV1_Sort) String() string {
	return proto.EnumName(SearchCGProductListingsRequestV1_Sort_name, int32(x))
}

func (SearchCGProductListingsRequestV1_Sort) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7612c057c03ad3a7, []int{2, 0}
}

type GetCGProductVariantCountsRequestV1 struct {
	CgproductId string `protobuf:"bytes,1,opt,name=cgproduct_id,json=cgproductId,proto3" json:"cgproduct_id,omitempty"`
	CountryId   int64  `protobuf:"varint,2,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	// used to determine which grouping to do; also includes condition
	VariationAttributes  []string           `protobuf:"bytes,3,rep,name=variation_attributes,json=variationAttributes,proto3" json:"variation_attributes,omitempty"`
	AuthUserContainer    *AuthUserContainer `protobuf:"bytes,4,opt,name=auth_user_container,json=authUserContainer,proto3" json:"auth_user_container,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetCGProductVariantCountsRequestV1) Reset()         { *m = GetCGProductVariantCountsRequestV1{} }
func (m *GetCGProductVariantCountsRequestV1) String() string { return proto.CompactTextString(m) }
func (*GetCGProductVariantCountsRequestV1) ProtoMessage()    {}
func (*GetCGProductVariantCountsRequestV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_7612c057c03ad3a7, []int{0}
}

func (m *GetCGProductVariantCountsRequestV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCGProductVariantCountsRequestV1.Unmarshal(m, b)
}
func (m *GetCGProductVariantCountsRequestV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCGProductVariantCountsRequestV1.Marshal(b, m, deterministic)
}
func (m *GetCGProductVariantCountsRequestV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCGProductVariantCountsRequestV1.Merge(m, src)
}
func (m *GetCGProductVariantCountsRequestV1) XXX_Size() int {
	return xxx_messageInfo_GetCGProductVariantCountsRequestV1.Size(m)
}
func (m *GetCGProductVariantCountsRequestV1) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCGProductVariantCountsRequestV1.DiscardUnknown(m)
}

var xxx_messageInfo_GetCGProductVariantCountsRequestV1 proto.InternalMessageInfo

func (m *GetCGProductVariantCountsRequestV1) GetCgproductId() string {
	if m != nil {
		return m.CgproductId
	}
	return ""
}

func (m *GetCGProductVariantCountsRequestV1) GetCountryId() int64 {
	if m != nil {
		return m.CountryId
	}
	return 0
}

func (m *GetCGProductVariantCountsRequestV1) GetVariationAttributes() []string {
	if m != nil {
		return m.VariationAttributes
	}
	return nil
}

func (m *GetCGProductVariantCountsRequestV1) GetAuthUserContainer() *AuthUserContainer {
	if m != nil {
		return m.AuthUserContainer
	}
	return nil
}

type GetCGProductVariantCountsResponseV1 struct {
	VariantCounts        []*GetCGProductVariantCountsResponseV1_VariantCount `protobuf:"bytes,1,rep,name=variant_counts,json=variantCounts,proto3" json:"variant_counts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                            `json:"-"`
	XXX_unrecognized     []byte                                              `json:"-"`
	XXX_sizecache        int32                                               `json:"-"`
}

func (m *GetCGProductVariantCountsResponseV1) Reset()         { *m = GetCGProductVariantCountsResponseV1{} }
func (m *GetCGProductVariantCountsResponseV1) String() string { return proto.CompactTextString(m) }
func (*GetCGProductVariantCountsResponseV1) ProtoMessage()    {}
func (*GetCGProductVariantCountsResponseV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_7612c057c03ad3a7, []int{1}
}

func (m *GetCGProductVariantCountsResponseV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCGProductVariantCountsResponseV1.Unmarshal(m, b)
}
func (m *GetCGProductVariantCountsResponseV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCGProductVariantCountsResponseV1.Marshal(b, m, deterministic)
}
func (m *GetCGProductVariantCountsResponseV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCGProductVariantCountsResponseV1.Merge(m, src)
}
func (m *GetCGProductVariantCountsResponseV1) XXX_Size() int {
	return xxx_messageInfo_GetCGProductVariantCountsResponseV1.Size(m)
}
func (m *GetCGProductVariantCountsResponseV1) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCGProductVariantCountsResponseV1.DiscardUnknown(m)
}

var xxx_messageInfo_GetCGProductVariantCountsResponseV1 proto.InternalMessageInfo

func (m *GetCGProductVariantCountsResponseV1) GetVariantCounts() []*GetCGProductVariantCountsResponseV1_VariantCount {
	if m != nil {
		return m.VariantCounts
	}
	return nil
}

type GetCGProductVariantCountsResponseV1_VariantCount struct {
	// should match the order of the variation_attributes
	VariationAttributeValues []string `protobuf:"bytes,1,rep,name=variation_attribute_values,json=variationAttributeValues,proto3" json:"variation_attribute_values,omitempty"`
	Count                    int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *GetCGProductVariantCountsResponseV1_VariantCount) Reset() {
	*m = GetCGProductVariantCountsResponseV1_VariantCount{}
}
func (m *GetCGProductVariantCountsResponseV1_VariantCount) String() string {
	return proto.CompactTextString(m)
}
func (*GetCGProductVariantCountsResponseV1_VariantCount) ProtoMessage() {}
func (*GetCGProductVariantCountsResponseV1_VariantCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_7612c057c03ad3a7, []int{1, 0}
}

func (m *GetCGProductVariantCountsResponseV1_VariantCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCGProductVariantCountsResponseV1_VariantCount.Unmarshal(m, b)
}
func (m *GetCGProductVariantCountsResponseV1_VariantCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCGProductVariantCountsResponseV1_VariantCount.Marshal(b, m, deterministic)
}
func (m *GetCGProductVariantCountsResponseV1_VariantCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCGProductVariantCountsResponseV1_VariantCount.Merge(m, src)
}
func (m *GetCGProductVariantCountsResponseV1_VariantCount) XXX_Size() int {
	return xxx_messageInfo_GetCGProductVariantCountsResponseV1_VariantCount.Size(m)
}
func (m *GetCGProductVariantCountsResponseV1_VariantCount) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCGProductVariantCountsResponseV1_VariantCount.DiscardUnknown(m)
}

var xxx_messageInfo_GetCGProductVariantCountsResponseV1_VariantCount proto.InternalMessageInfo

func (m *GetCGProductVariantCountsResponseV1_VariantCount) GetVariationAttributeValues() []string {
	if m != nil {
		return m.VariationAttributeValues
	}
	return nil
}

func (m *GetCGProductVariantCountsResponseV1_VariantCount) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type SearchCGProductListingsRequestV1 struct {
	CountryId         int64                                 `protobuf:"varint,1,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	Count             int32                                 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Filters           []*FilterParamV4                      `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	Sort              SearchCGProductListingsRequestV1_Sort `protobuf:"varint,4,opt,name=sort,proto3,enum=ContentFeedService_proto.SearchCGProductListingsRequestV1_Sort" json:"sort,omitempty"`
	AuthUserContainer *AuthUserContainer                    `protobuf:"bytes,5,opt,name=auth_user_container,json=authUserContainer,proto3" json:"auth_user_container,omitempty"`
	// Internal use
	IncludeDebuggingInfo bool     `protobuf:"varint,99,opt,name=include_debugging_info,json=includeDebuggingInfo,proto3" json:"include_debugging_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchCGProductListingsRequestV1) Reset()         { *m = SearchCGProductListingsRequestV1{} }
func (m *SearchCGProductListingsRequestV1) String() string { return proto.CompactTextString(m) }
func (*SearchCGProductListingsRequestV1) ProtoMessage()    {}
func (*SearchCGProductListingsRequestV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_7612c057c03ad3a7, []int{2}
}

func (m *SearchCGProductListingsRequestV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchCGProductListingsRequestV1.Unmarshal(m, b)
}
func (m *SearchCGProductListingsRequestV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchCGProductListingsRequestV1.Marshal(b, m, deterministic)
}
func (m *SearchCGProductListingsRequestV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchCGProductListingsRequestV1.Merge(m, src)
}
func (m *SearchCGProductListingsRequestV1) XXX_Size() int {
	return xxx_messageInfo_SearchCGProductListingsRequestV1.Size(m)
}
func (m *SearchCGProductListingsRequestV1) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchCGProductListingsRequestV1.DiscardUnknown(m)
}

var xxx_messageInfo_SearchCGProductListingsRequestV1 proto.InternalMessageInfo

func (m *SearchCGProductListingsRequestV1) GetCountryId() int64 {
	if m != nil {
		return m.CountryId
	}
	return 0
}

func (m *SearchCGProductListingsRequestV1) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *SearchCGProductListingsRequestV1) GetFilters() []*FilterParamV4 {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *SearchCGProductListingsRequestV1) GetSort() SearchCGProductListingsRequestV1_Sort {
	if m != nil {
		return m.Sort
	}
	return SearchCGProductListingsRequestV1_SORT_TYPE_UNSPECIFIED
}

func (m *SearchCGProductListingsRequestV1) GetAuthUserContainer() *AuthUserContainer {
	if m != nil {
		return m.AuthUserContainer
	}
	return nil
}

func (m *SearchCGProductListingsRequestV1) GetIncludeDebuggingInfo() bool {
	if m != nil {
		return m.IncludeDebuggingInfo
	}
	return false
}

type SearchCGProductListingsResponseV1 struct {
	CgproductListing     []*SearchCGProductListingsResponseV1_CGProductListing `protobuf:"bytes,1,rep,name=cgproduct_listing,json=cgproductListing,proto3" json:"cgproduct_listing,omitempty"`
	DebuggingInfo        *SearchCGProductListingsResponseV1_DebuggingInfo      `protobuf:"bytes,99,opt,name=debugging_info,json=debuggingInfo,proto3" json:"debugging_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                              `json:"-"`
	XXX_unrecognized     []byte                                                `json:"-"`
	XXX_sizecache        int32                                                 `json:"-"`
}

func (m *SearchCGProductListingsResponseV1) Reset()         { *m = SearchCGProductListingsResponseV1{} }
func (m *SearchCGProductListingsResponseV1) String() string { return proto.CompactTextString(m) }
func (*SearchCGProductListingsResponseV1) ProtoMessage()    {}
func (*SearchCGProductListingsResponseV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_7612c057c03ad3a7, []int{3}
}

func (m *SearchCGProductListingsResponseV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchCGProductListingsResponseV1.Unmarshal(m, b)
}
func (m *SearchCGProductListingsResponseV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchCGProductListingsResponseV1.Marshal(b, m, deterministic)
}
func (m *SearchCGProductListingsResponseV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchCGProductListingsResponseV1.Merge(m, src)
}
func (m *SearchCGProductListingsResponseV1) XXX_Size() int {
	return xxx_messageInfo_SearchCGProductListingsResponseV1.Size(m)
}
func (m *SearchCGProductListingsResponseV1) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchCGProductListingsResponseV1.DiscardUnknown(m)
}

var xxx_messageInfo_SearchCGProductListingsResponseV1 proto.InternalMessageInfo

func (m *SearchCGProductListingsResponseV1) GetCgproductListing() []*SearchCGProductListingsResponseV1_CGProductListing {
	if m != nil {
		return m.CgproductListing
	}
	return nil
}

func (m *SearchCGProductListingsResponseV1) GetDebuggingInfo() *SearchCGProductListingsResponseV1_DebuggingInfo {
	if m != nil {
		return m.DebuggingInfo
	}
	return nil
}

type SearchCGProductListingsResponseV1_CGProductListing struct {
	LegacyId   int64  `protobuf:"varint,1,opt,name=legacy_id,json=legacyId,proto3" json:"legacy_id,omitempty"`
	CategoryId string `protobuf:"bytes,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	//// Internal use
	// Debugging information
	Explain              string   `protobuf:"bytes,99,opt,name=explain,proto3" json:"explain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchCGProductListingsResponseV1_CGProductListing) Reset() {
	*m = SearchCGProductListingsResponseV1_CGProductListing{}
}
func (m *SearchCGProductListingsResponseV1_CGProductListing) String() string {
	return proto.CompactTextString(m)
}
func (*SearchCGProductListingsResponseV1_CGProductListing) ProtoMessage() {}
func (*SearchCGProductListingsResponseV1_CGProductListing) Descriptor() ([]byte, []int) {
	return fileDescriptor_7612c057c03ad3a7, []int{3, 0}
}

func (m *SearchCGProductListingsResponseV1_CGProductListing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchCGProductListingsResponseV1_CGProductListing.Unmarshal(m, b)
}
func (m *SearchCGProductListingsResponseV1_CGProductListing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchCGProductListingsResponseV1_CGProductListing.Marshal(b, m, deterministic)
}
func (m *SearchCGProductListingsResponseV1_CGProductListing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchCGProductListingsResponseV1_CGProductListing.Merge(m, src)
}
func (m *SearchCGProductListingsResponseV1_CGProductListing) XXX_Size() int {
	return xxx_messageInfo_SearchCGProductListingsResponseV1_CGProductListing.Size(m)
}
func (m *SearchCGProductListingsResponseV1_CGProductListing) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchCGProductListingsResponseV1_CGProductListing.DiscardUnknown(m)
}

var xxx_messageInfo_SearchCGProductListingsResponseV1_CGProductListing proto.InternalMessageInfo

func (m *SearchCGProductListingsResponseV1_CGProductListing) GetLegacyId() int64 {
	if m != nil {
		return m.LegacyId
	}
	return 0
}

func (m *SearchCGProductListingsResponseV1_CGProductListing) GetCategoryId() string {
	if m != nil {
		return m.CategoryId
	}
	return ""
}

func (m *SearchCGProductListingsResponseV1_CGProductListing) GetExplain() string {
	if m != nil {
		return m.Explain
	}
	return ""
}

//// Internal use
// Debugging information for us to get info of each search query.
type SearchCGProductListingsResponseV1_DebuggingInfo struct {
	SearchQueries        *SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery `protobuf:"bytes,3,opt,name=search_queries,json=searchQueries,proto3" json:"search_queries,omitempty"`
	Total                *wrappers.Int64Value                                         `protobuf:"bytes,4,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                     `json:"-"`
	XXX_unrecognized     []byte                                                       `json:"-"`
	XXX_sizecache        int32                                                        `json:"-"`
}

func (m *SearchCGProductListingsResponseV1_DebuggingInfo) Reset() {
	*m = SearchCGProductListingsResponseV1_DebuggingInfo{}
}
func (m *SearchCGProductListingsResponseV1_DebuggingInfo) String() string {
	return proto.CompactTextString(m)
}
func (*SearchCGProductListingsResponseV1_DebuggingInfo) ProtoMessage() {}
func (*SearchCGProductListingsResponseV1_DebuggingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7612c057c03ad3a7, []int{3, 1}
}

func (m *SearchCGProductListingsResponseV1_DebuggingInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchCGProductListingsResponseV1_DebuggingInfo.Unmarshal(m, b)
}
func (m *SearchCGProductListingsResponseV1_DebuggingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchCGProductListingsResponseV1_DebuggingInfo.Marshal(b, m, deterministic)
}
func (m *SearchCGProductListingsResponseV1_DebuggingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchCGProductListingsResponseV1_DebuggingInfo.Merge(m, src)
}
func (m *SearchCGProductListingsResponseV1_DebuggingInfo) XXX_Size() int {
	return xxx_messageInfo_SearchCGProductListingsResponseV1_DebuggingInfo.Size(m)
}
func (m *SearchCGProductListingsResponseV1_DebuggingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchCGProductListingsResponseV1_DebuggingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SearchCGProductListingsResponseV1_DebuggingInfo proto.InternalMessageInfo

func (m *SearchCGProductListingsResponseV1_DebuggingInfo) GetSearchQueries() *SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery {
	if m != nil {
		return m.SearchQueries
	}
	return nil
}

func (m *SearchCGProductListingsResponseV1_DebuggingInfo) GetTotal() *wrappers.Int64Value {
	if m != nil {
		return m.Total
	}
	return nil
}

type SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	EsQuery              string   `protobuf:"bytes,2,opt,name=es_query,json=esQuery,proto3" json:"es_query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery) Reset() {
	*m = SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery{}
}
func (m *SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery) String() string {
	return proto.CompactTextString(m)
}
func (*SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery) ProtoMessage() {}
func (*SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_7612c057c03ad3a7, []int{3, 1, 0}
}

func (m *SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery.Unmarshal(m, b)
}
func (m *SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery.Marshal(b, m, deterministic)
}
func (m *SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery.Merge(m, src)
}
func (m *SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery) XXX_Size() int {
	return xxx_messageInfo_SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery.Size(m)
}
func (m *SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery.DiscardUnknown(m)
}

var xxx_messageInfo_SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery proto.InternalMessageInfo

func (m *SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery) GetEsQuery() string {
	if m != nil {
		return m.EsQuery
	}
	return ""
}

func init() {
	proto.RegisterEnum("ContentFeedService_proto.SearchCGProductListingsRequestV1_Sort", SearchCGProductListingsRequestV1_Sort_name, SearchCGProductListingsRequestV1_Sort_value)
	proto.RegisterType((*GetCGProductVariantCountsRequestV1)(nil), "ContentFeedService_proto.GetCGProductVariantCountsRequestV1")
	proto.RegisterType((*GetCGProductVariantCountsResponseV1)(nil), "ContentFeedService_proto.GetCGProductVariantCountsResponseV1")
	proto.RegisterType((*GetCGProductVariantCountsResponseV1_VariantCount)(nil), "ContentFeedService_proto.GetCGProductVariantCountsResponseV1.VariantCount")
	proto.RegisterType((*SearchCGProductListingsRequestV1)(nil), "ContentFeedService_proto.SearchCGProductListingsRequestV1")
	proto.RegisterType((*SearchCGProductListingsResponseV1)(nil), "ContentFeedService_proto.SearchCGProductListingsResponseV1")
	proto.RegisterType((*SearchCGProductListingsResponseV1_CGProductListing)(nil), "ContentFeedService_proto.SearchCGProductListingsResponseV1.CGProductListing")
	proto.RegisterType((*SearchCGProductListingsResponseV1_DebuggingInfo)(nil), "ContentFeedService_proto.SearchCGProductListingsResponseV1.DebuggingInfo")
	proto.RegisterType((*SearchCGProductListingsResponseV1_DebuggingInfo_SearchQuery)(nil), "ContentFeedService_proto.SearchCGProductListingsResponseV1.DebuggingInfo.SearchQuery")
}

func init() { proto.RegisterFile("content-feed/cgproduct.proto", fileDescriptor_7612c057c03ad3a7) }

var fileDescriptor_7612c057c03ad3a7 = []byte{
	// 809 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6f, 0xe3, 0x44,
	0x10, 0xc7, 0x4d, 0x4a, 0x9b, 0xc9, 0xb5, 0xca, 0xed, 0x15, 0x94, 0xa6, 0x70, 0x97, 0x0b, 0x0f,
	0x44, 0x42, 0xe7, 0xa8, 0xa1, 0xe2, 0x85, 0x93, 0x50, 0x49, 0xd3, 0x9e, 0x51, 0x3f, 0xc2, 0xa6,
	0x0d, 0xe2, 0x43, 0xb2, 0x36, 0xf6, 0xc4, 0x31, 0x72, 0x77, 0xd3, 0xdd, 0x75, 0x20, 0x12, 0x2f,
	0xbc, 0xf0, 0x47, 0xf1, 0x87, 0xf1, 0x04, 0x08, 0x65, 0xed, 0xb8, 0x4e, 0x68, 0x00, 0x71, 0xf7,
	0xe6, 0xf9, 0xfd, 0xe6, 0x7b, 0x3c, 0xb3, 0xf0, 0x9e, 0x27, 0xb8, 0x46, 0xae, 0x5f, 0x8c, 0x10,
	0xfd, 0x96, 0x17, 0x4c, 0xa4, 0xf0, 0x63, 0x4f, 0xdb, 0x13, 0x29, 0xb4, 0x20, 0xd5, 0x4e, 0xc2,
	0x9e, 0x22, 0xfa, 0x7d, 0x94, 0xd3, 0xd0, 0x43, 0xd7, 0x30, 0xb5, 0xa7, 0x81, 0x10, 0x41, 0x84,
	0x2d, 0x23, 0x0d, 0xe3, 0x51, 0xeb, 0x07, 0xc9, 0x26, 0x13, 0x94, 0x2a, 0xb1, 0xac, 0xed, 0x2f,
	0xfb, 0x15, 0xb7, 0xb7, 0x82, 0xa7, 0xd4, 0x72, 0x48, 0x85, 0x4c, 0x7a, 0x63, 0x77, 0x7a, 0x94,
	0xb0, 0x8d, 0xdf, 0x2c, 0x68, 0x9c, 0xa1, 0xee, 0x9c, 0xf5, 0x92, 0x4c, 0x06, 0x4c, 0x86, 0x8c,
	0xeb, 0x8e, 0x88, 0xb9, 0x56, 0x14, 0xef, 0x62, 0x54, 0x7a, 0x70, 0x48, 0x9e, 0xc3, 0xa3, 0x2c,
	0x59, 0x37, 0xf4, 0xab, 0x56, 0xdd, 0x6a, 0x96, 0x68, 0x39, 0xc3, 0x1c, 0x9f, 0xbc, 0x0f, 0xe0,
	0xcd, 0xad, 0xe4, 0x6c, 0xae, 0xb0, 0x51, 0xb7, 0x9a, 0x05, 0x5a, 0x4a, 0x11, 0xc7, 0x27, 0x87,
	0xb0, 0x37, 0x9d, 0xfb, 0xd6, 0xa1, 0xe0, 0x2e, 0xd3, 0x5a, 0x86, 0xc3, 0x58, 0xa3, 0xaa, 0x16,
	0xea, 0x85, 0x66, 0x89, 0x3e, 0xc9, 0xb8, 0xe3, 0x8c, 0x22, 0xdf, 0xc2, 0x13, 0x16, 0xeb, 0xb1,
	0x1b, 0x2b, 0x94, 0xee, 0xbc, 0x0a, 0x16, 0x72, 0x94, 0xd5, 0x62, 0xdd, 0x6a, 0x96, 0xdb, 0x1f,
	0xd9, 0xeb, 0x9a, 0x65, 0x1f, 0xc7, 0x7a, 0x7c, 0xa3, 0x50, 0x76, 0x16, 0x26, 0xf4, 0x31, 0x5b,
	0x85, 0x1a, 0xbf, 0x5b, 0xf0, 0xc1, 0x3f, 0x14, 0xae, 0x26, 0x82, 0x2b, 0x1c, 0x1c, 0x92, 0x3b,
	0xd8, 0x9d, 0x26, 0x94, 0x6b, 0x8a, 0x51, 0x55, 0xab, 0x5e, 0x68, 0x96, 0xdb, 0x5f, 0xac, 0x8f,
	0xff, 0x1f, 0xdc, 0xda, 0x79, 0x9c, 0xee, 0x4c, 0xf3, 0x5a, 0xb5, 0x21, 0x3c, 0xca, 0xd3, 0xe4,
	0x25, 0xd4, 0x1e, 0x68, 0x9d, 0x3b, 0x65, 0x51, 0x8c, 0x49, 0x3a, 0x25, 0x5a, 0xfd, 0x7b, 0x03,
	0x07, 0x86, 0x27, 0x7b, 0xb0, 0x69, 0x12, 0x4f, 0x47, 0x92, 0x08, 0x8d, 0x5f, 0x8a, 0x50, 0xef,
	0x9b, 0x7f, 0x21, 0x4b, 0xf5, 0x3c, 0x54, 0x3a, 0xe4, 0x41, 0x6e, 0xea, 0xcb, 0x23, 0xb5, 0x56,
	0x47, 0xba, 0xe4, 0x79, 0x33, 0xf5, 0x4c, 0x8e, 0x61, 0x6b, 0x14, 0x46, 0x1a, 0x65, 0x32, 0xdb,
	0x72, 0xfb, 0xc3, 0xf5, 0x9d, 0x3a, 0x35, 0x8a, 0x3d, 0x26, 0xd9, 0xed, 0xe0, 0x88, 0x2e, 0xec,
	0x48, 0x1f, 0x8a, 0x4a, 0x48, 0x6d, 0x26, 0xbd, 0xdb, 0xfe, 0x6c, 0xbd, 0xfd, 0xbf, 0x55, 0x60,
	0xf7, 0x85, 0xd4, 0xd4, 0x38, 0x5b, 0xf7, 0x37, 0x6d, 0xbe, 0x89, 0xbf, 0x89, 0x1c, 0xc1, 0xbb,
	0x21, 0xf7, 0xa2, 0xd8, 0x47, 0xd7, 0xc7, 0x61, 0x1c, 0x04, 0x21, 0x0f, 0xdc, 0x90, 0x8f, 0x44,
	0xd5, 0xab, 0x5b, 0xcd, 0x6d, 0xba, 0x97, 0xb2, 0x27, 0x0b, 0xd2, 0xe1, 0x23, 0xd1, 0xf8, 0xd9,
	0x82, 0xe2, 0x3c, 0x43, 0xb2, 0x0f, 0xef, 0xf4, 0xaf, 0xe8, 0xb5, 0x7b, 0xfd, 0x75, 0xaf, 0xeb,
	0xde, 0x5c, 0xf6, 0x7b, 0xdd, 0x8e, 0x73, 0xea, 0x74, 0x4f, 0x2a, 0x6f, 0x91, 0x67, 0x70, 0x70,
	0x4f, 0xf5, 0xa8, 0xd3, 0xe9, 0xba, 0xe7, 0x57, 0x5f, 0xb9, 0xd7, 0x57, 0xee, 0x2b, 0xe7, 0xec,
	0x55, 0xc5, 0x5a, 0xb6, 0xa5, 0xdd, 0xce, 0xd5, 0xc5, 0x45, 0xf7, 0xf2, 0xa4, 0x7b, 0x52, 0xd9,
	0x20, 0x0d, 0x78, 0xfa, 0x20, 0xe5, 0xf6, 0x8f, 0x2f, 0x7a, 0xe7, 0xce, 0xe5, 0x59, 0xa5, 0xd0,
	0xf8, 0xa3, 0x08, 0xcf, 0xd7, 0xb6, 0x31, 0xdb, 0x82, 0x19, 0x3c, 0xbe, 0xdf, 0xff, 0x28, 0xe1,
	0xd3, 0x45, 0x38, 0xff, 0x1f, 0xe3, 0xc9, 0xd6, 0x60, 0x95, 0xa3, 0x95, 0x2c, 0x4c, 0x8a, 0x90,
	0x09, 0xec, 0x3e, 0xd0, 0xd2, 0x72, 0xdb, 0x79, 0x9d, 0xb8, 0x4b, 0x73, 0xa0, 0x3b, 0x7e, 0x5e,
	0xac, 0x7d, 0x0f, 0x95, 0x55, 0x5b, 0x72, 0x00, 0xa5, 0x08, 0x03, 0xe6, 0xe5, 0x36, 0x61, 0x3b,
	0x01, 0x1c, 0x9f, 0x3c, 0x83, 0xb2, 0xc7, 0x34, 0x06, 0xe2, 0xfe, 0xf6, 0x95, 0x28, 0x2c, 0x20,
	0xc7, 0x27, 0x55, 0xd8, 0xc2, 0x1f, 0x27, 0x11, 0x0b, 0xb9, 0x49, 0xbe, 0x44, 0x17, 0x62, 0xed,
	0x4f, 0x0b, 0x76, 0x96, 0x92, 0x21, 0x3f, 0xc1, 0x6e, 0x7a, 0xa4, 0xef, 0x62, 0x94, 0xa1, 0x39,
	0x91, 0xf3, 0x7a, 0x6f, 0xde, 0x58, 0xbd, 0xa9, 0xfe, 0x97, 0x31, 0xca, 0x19, 0xdd, 0x51, 0x99,
	0x10, 0xa2, 0x22, 0x87, 0xb0, 0xa9, 0x85, 0x66, 0x51, 0x7a, 0x65, 0x0f, 0xec, 0xe4, 0xe1, 0xb1,
	0x17, 0x0f, 0x8f, 0xed, 0x70, 0xfd, 0xc9, 0x91, 0x39, 0x2d, 0x34, 0xd1, 0xac, 0xbd, 0x84, 0x72,
	0xce, 0x21, 0x21, 0x50, 0xd4, 0xb3, 0x09, 0xa6, 0x4f, 0x84, 0xf9, 0x26, 0xfb, 0xb0, 0x8d, 0xca,
	0xd4, 0x33, 0x4b, 0xbb, 0xb3, 0x85, 0xca, 0xa8, 0x7f, 0xfe, 0xdd, 0x37, 0x97, 0x41, 0xa8, 0xc7,
	0xf1, 0xd0, 0xf6, 0xc4, 0x6d, 0xcb, 0x63, 0x52, 0xc4, 0x0a, 0xa3, 0xa8, 0xa5, 0xc6, 0x4c, 0xa2,
	0xff, 0xc2, 0xc4, 0x6d, 0x05, 0xc8, 0x5b, 0x81, 0x68, 0xe5, 0x5f, 0xb2, 0x4f, 0xd7, 0x75, 0xe3,
	0xd7, 0x8d, 0x52, 0xe7, 0x34, 0x45, 0x86, 0x6f, 0x1b, 0xe8, 0xe3, 0xbf, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x21, 0x23, 0xa6, 0xb5, 0x78, 0x07, 0x00, 0x00,
}
