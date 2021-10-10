// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/recommendationengine/v1beta1/catalog.proto

package recommendationengine

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Optional. Item stock state. If this field is unspecified, the item is
// assumed to be in stock.
type ProductCatalogItem_StockState int32

const (
	// Default item stock status. Should never be used.
	ProductCatalogItem_STOCK_STATE_UNSPECIFIED ProductCatalogItem_StockState = 0
	// Item in stock.
	ProductCatalogItem_IN_STOCK ProductCatalogItem_StockState = 0
	// Item out of stock.
	ProductCatalogItem_OUT_OF_STOCK ProductCatalogItem_StockState = 1
	// Item that is in pre-order state.
	ProductCatalogItem_PREORDER ProductCatalogItem_StockState = 2
	// Item that is back-ordered (i.e. temporarily out of stock).
	ProductCatalogItem_BACKORDER ProductCatalogItem_StockState = 3
)

var ProductCatalogItem_StockState_name = map[int32]string{
	0: "STOCK_STATE_UNSPECIFIED",
	// Duplicate value: 0: "IN_STOCK",
	1: "OUT_OF_STOCK",
	2: "PREORDER",
	3: "BACKORDER",
}

var ProductCatalogItem_StockState_value = map[string]int32{
	"STOCK_STATE_UNSPECIFIED": 0,
	"IN_STOCK":                0,
	"OUT_OF_STOCK":            1,
	"PREORDER":                2,
	"BACKORDER":               3,
}

func (x ProductCatalogItem_StockState) String() string {
	return proto.EnumName(ProductCatalogItem_StockState_name, int32(x))
}

func (ProductCatalogItem_StockState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_988daa8a4f3967d9, []int{1, 0}
}

// CatalogItem captures all metadata information of items to be recommended.
type CatalogItem struct {
	// Required. Catalog item identifier. UTF-8 encoded string with a length limit
	// of 128 characters.
	//
	// This id must be unique among all catalog items within the same catalog. It
	// should also be used when logging user events in order for the user events
	// to be joined with the Catalog.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Required. Catalog item categories. This field is repeated for supporting
	// one catalog item belonging to several parallel category hierarchies.
	//
	// For example, if a shoes product belongs to both
	// ["Shoes & Accessories" -> "Shoes"] and
	// ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be
	// represented as:
	//
	//      "categoryHierarchies": [
	//        { "categories": ["Shoes & Accessories", "Shoes"]},
	//        { "categories": ["Sports & Fitness", "Athletic Clothing", "Shoes"] }
	//      ]
	CategoryHierarchies []*CatalogItem_CategoryHierarchy `protobuf:"bytes,2,rep,name=category_hierarchies,json=categoryHierarchies,proto3" json:"category_hierarchies,omitempty"`
	// Required. Catalog item title. UTF-8 encoded string with a length limit of
	// 1250 characters.
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// Optional. Catalog item description. UTF-8 encoded string with a length
	// limit of 1250 characters.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Optional (but highly encouraged). Extra catalog item attributes to be
	// included in the recommendation model. For example, for retail products,
	// this could include the store name, vendor, style, color, etc. These are
	// very strong signals for recommendation model, thus we highly recommend
	// providing the item attributes here.
	ItemAttributes *FeatureMap `protobuf:"bytes,5,opt,name=item_attributes,json=itemAttributes,proto3" json:"item_attributes,omitempty"`
	// Optional. Language of the title/description/item_attributes. Use language
	// tags defined by BCP 47. https://www.rfc-editor.org/rfc/bcp/bcp47.txt. Our
	// supported language codes include 'en', 'es', 'fr', 'de', 'ar', 'fa', 'zh',
	// 'ja', 'ko', 'sv', 'ro', 'nl'. For other languages, contact
	// your Google account manager.
	LanguageCode string `protobuf:"bytes,6,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Optional. Filtering tags associated with the catalog item. This tag can be
	// used for filtering recommendation results by passing the tag as part of the
	// predict request filter. The tags have to satisfy the following
	// restrictions:
	//
	// * Only contain alphanumeric characters (`a-z`, `A-Z`, `0-9`), underscores
	//   (`_`) and dashes (`-`).
	Tags []string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	// Optional. Variant group identifier for prediction results. UTF-8 encoded
	// string with a length limit of 128 characters.
	//
	// This field must be enabled before it can be used. [Learn
	// more](/recommendations-ai/docs/catalog#item-group-id).
	ItemGroupId string `protobuf:"bytes,9,opt,name=item_group_id,json=itemGroupId,proto3" json:"item_group_id,omitempty"`
	// Extra catalog item metadata for different recommendation types.
	//
	// Types that are valid to be assigned to RecommendationType:
	//	*CatalogItem_ProductMetadata
	RecommendationType   isCatalogItem_RecommendationType `protobuf_oneof:"recommendation_type"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *CatalogItem) Reset()         { *m = CatalogItem{} }
func (m *CatalogItem) String() string { return proto.CompactTextString(m) }
func (*CatalogItem) ProtoMessage()    {}
func (*CatalogItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_988daa8a4f3967d9, []int{0}
}

func (m *CatalogItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CatalogItem.Unmarshal(m, b)
}
func (m *CatalogItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CatalogItem.Marshal(b, m, deterministic)
}
func (m *CatalogItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatalogItem.Merge(m, src)
}
func (m *CatalogItem) XXX_Size() int {
	return xxx_messageInfo_CatalogItem.Size(m)
}
func (m *CatalogItem) XXX_DiscardUnknown() {
	xxx_messageInfo_CatalogItem.DiscardUnknown(m)
}

var xxx_messageInfo_CatalogItem proto.InternalMessageInfo

func (m *CatalogItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CatalogItem) GetCategoryHierarchies() []*CatalogItem_CategoryHierarchy {
	if m != nil {
		return m.CategoryHierarchies
	}
	return nil
}

func (m *CatalogItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CatalogItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CatalogItem) GetItemAttributes() *FeatureMap {
	if m != nil {
		return m.ItemAttributes
	}
	return nil
}

func (m *CatalogItem) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *CatalogItem) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *CatalogItem) GetItemGroupId() string {
	if m != nil {
		return m.ItemGroupId
	}
	return ""
}

type isCatalogItem_RecommendationType interface {
	isCatalogItem_RecommendationType()
}

type CatalogItem_ProductMetadata struct {
	ProductMetadata *ProductCatalogItem `protobuf:"bytes,10,opt,name=product_metadata,json=productMetadata,proto3,oneof"`
}

func (*CatalogItem_ProductMetadata) isCatalogItem_RecommendationType() {}

func (m *CatalogItem) GetRecommendationType() isCatalogItem_RecommendationType {
	if m != nil {
		return m.RecommendationType
	}
	return nil
}

func (m *CatalogItem) GetProductMetadata() *ProductCatalogItem {
	if x, ok := m.GetRecommendationType().(*CatalogItem_ProductMetadata); ok {
		return x.ProductMetadata
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CatalogItem) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CatalogItem_ProductMetadata)(nil),
	}
}

// Category represents catalog item category hierarchy.
type CatalogItem_CategoryHierarchy struct {
	// Catalog item categories. Note that the order in the list denotes the
	// specificity (from least to most specific). Required.
	Categories           []string `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CatalogItem_CategoryHierarchy) Reset()         { *m = CatalogItem_CategoryHierarchy{} }
func (m *CatalogItem_CategoryHierarchy) String() string { return proto.CompactTextString(m) }
func (*CatalogItem_CategoryHierarchy) ProtoMessage()    {}
func (*CatalogItem_CategoryHierarchy) Descriptor() ([]byte, []int) {
	return fileDescriptor_988daa8a4f3967d9, []int{0, 0}
}

func (m *CatalogItem_CategoryHierarchy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CatalogItem_CategoryHierarchy.Unmarshal(m, b)
}
func (m *CatalogItem_CategoryHierarchy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CatalogItem_CategoryHierarchy.Marshal(b, m, deterministic)
}
func (m *CatalogItem_CategoryHierarchy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatalogItem_CategoryHierarchy.Merge(m, src)
}
func (m *CatalogItem_CategoryHierarchy) XXX_Size() int {
	return xxx_messageInfo_CatalogItem_CategoryHierarchy.Size(m)
}
func (m *CatalogItem_CategoryHierarchy) XXX_DiscardUnknown() {
	xxx_messageInfo_CatalogItem_CategoryHierarchy.DiscardUnknown(m)
}

var xxx_messageInfo_CatalogItem_CategoryHierarchy proto.InternalMessageInfo

func (m *CatalogItem_CategoryHierarchy) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

// ProductCatalogItem captures item metadata specific to retail products.
type ProductCatalogItem struct {
	// Product price. Only one of 'exactPrice'/'priceRange' can be provided.
	//
	// Types that are valid to be assigned to Price:
	//	*ProductCatalogItem_ExactPrice_
	//	*ProductCatalogItem_PriceRange_
	Price isProductCatalogItem_Price `protobuf_oneof:"price"`
	// Optional. A map to pass the costs associated with the product.
	//
	// For example:
	// {"manufacturing": 45.5} The profit of selling this item is computed like
	// so:
	//
	// * If 'exactPrice' is provided, profit = displayPrice - sum(costs)
	// * If 'priceRange' is provided, profit = minPrice - sum(costs)
	Costs map[string]float32 `protobuf:"bytes,3,rep,name=costs,proto3" json:"costs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	// Required if the price is set. Currency code for price/costs. Use
	// three-character ISO-4217 code.
	CurrencyCode string `protobuf:"bytes,4,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// Optional. Stock state of the catalog item. Default is `IN_STOCK`.
	StockState ProductCatalogItem_StockState `protobuf:"varint,5,opt,name=stock_state,json=stockState,proto3,enum=google.cloud.recommendationengine.v1beta1.ProductCatalogItem_StockState" json:"stock_state,omitempty"`
	// Optional. The available quantity of the item.
	AvailableQuantity int64 `protobuf:"varint,6,opt,name=available_quantity,json=availableQuantity,proto3" json:"available_quantity,omitempty"`
	// Optional. Canonical URL directly linking to the item detail page.
	CanonicalProductUri string `protobuf:"bytes,7,opt,name=canonical_product_uri,json=canonicalProductUri,proto3" json:"canonical_product_uri,omitempty"`
	// Optional. Product images for the catalog item.
	Images               []*Image `protobuf:"bytes,8,rep,name=images,proto3" json:"images,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductCatalogItem) Reset()         { *m = ProductCatalogItem{} }
func (m *ProductCatalogItem) String() string { return proto.CompactTextString(m) }
func (*ProductCatalogItem) ProtoMessage()    {}
func (*ProductCatalogItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_988daa8a4f3967d9, []int{1}
}

func (m *ProductCatalogItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductCatalogItem.Unmarshal(m, b)
}
func (m *ProductCatalogItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductCatalogItem.Marshal(b, m, deterministic)
}
func (m *ProductCatalogItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductCatalogItem.Merge(m, src)
}
func (m *ProductCatalogItem) XXX_Size() int {
	return xxx_messageInfo_ProductCatalogItem.Size(m)
}
func (m *ProductCatalogItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductCatalogItem.DiscardUnknown(m)
}

var xxx_messageInfo_ProductCatalogItem proto.InternalMessageInfo

type isProductCatalogItem_Price interface {
	isProductCatalogItem_Price()
}

type ProductCatalogItem_ExactPrice_ struct {
	ExactPrice *ProductCatalogItem_ExactPrice `protobuf:"bytes,1,opt,name=exact_price,json=exactPrice,proto3,oneof"`
}

type ProductCatalogItem_PriceRange_ struct {
	PriceRange *ProductCatalogItem_PriceRange `protobuf:"bytes,2,opt,name=price_range,json=priceRange,proto3,oneof"`
}

func (*ProductCatalogItem_ExactPrice_) isProductCatalogItem_Price() {}

func (*ProductCatalogItem_PriceRange_) isProductCatalogItem_Price() {}

func (m *ProductCatalogItem) GetPrice() isProductCatalogItem_Price {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *ProductCatalogItem) GetExactPrice() *ProductCatalogItem_ExactPrice {
	if x, ok := m.GetPrice().(*ProductCatalogItem_ExactPrice_); ok {
		return x.ExactPrice
	}
	return nil
}

func (m *ProductCatalogItem) GetPriceRange() *ProductCatalogItem_PriceRange {
	if x, ok := m.GetPrice().(*ProductCatalogItem_PriceRange_); ok {
		return x.PriceRange
	}
	return nil
}

func (m *ProductCatalogItem) GetCosts() map[string]float32 {
	if m != nil {
		return m.Costs
	}
	return nil
}

func (m *ProductCatalogItem) GetCurrencyCode() string {
	if m != nil {
		return m.CurrencyCode
	}
	return ""
}

func (m *ProductCatalogItem) GetStockState() ProductCatalogItem_StockState {
	if m != nil {
		return m.StockState
	}
	return ProductCatalogItem_STOCK_STATE_UNSPECIFIED
}

func (m *ProductCatalogItem) GetAvailableQuantity() int64 {
	if m != nil {
		return m.AvailableQuantity
	}
	return 0
}

func (m *ProductCatalogItem) GetCanonicalProductUri() string {
	if m != nil {
		return m.CanonicalProductUri
	}
	return ""
}

func (m *ProductCatalogItem) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ProductCatalogItem) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ProductCatalogItem_ExactPrice_)(nil),
		(*ProductCatalogItem_PriceRange_)(nil),
	}
}

// Exact product price.
type ProductCatalogItem_ExactPrice struct {
	// Optional. Display price of the product.
	DisplayPrice float32 `protobuf:"fixed32,1,opt,name=display_price,json=displayPrice,proto3" json:"display_price,omitempty"`
	// Optional. Price of the product without any discount. If zero, by default
	// set to be the 'displayPrice'.
	OriginalPrice        float32  `protobuf:"fixed32,2,opt,name=original_price,json=originalPrice,proto3" json:"original_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductCatalogItem_ExactPrice) Reset()         { *m = ProductCatalogItem_ExactPrice{} }
func (m *ProductCatalogItem_ExactPrice) String() string { return proto.CompactTextString(m) }
func (*ProductCatalogItem_ExactPrice) ProtoMessage()    {}
func (*ProductCatalogItem_ExactPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_988daa8a4f3967d9, []int{1, 0}
}

func (m *ProductCatalogItem_ExactPrice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductCatalogItem_ExactPrice.Unmarshal(m, b)
}
func (m *ProductCatalogItem_ExactPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductCatalogItem_ExactPrice.Marshal(b, m, deterministic)
}
func (m *ProductCatalogItem_ExactPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductCatalogItem_ExactPrice.Merge(m, src)
}
func (m *ProductCatalogItem_ExactPrice) XXX_Size() int {
	return xxx_messageInfo_ProductCatalogItem_ExactPrice.Size(m)
}
func (m *ProductCatalogItem_ExactPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductCatalogItem_ExactPrice.DiscardUnknown(m)
}

var xxx_messageInfo_ProductCatalogItem_ExactPrice proto.InternalMessageInfo

func (m *ProductCatalogItem_ExactPrice) GetDisplayPrice() float32 {
	if m != nil {
		return m.DisplayPrice
	}
	return 0
}

func (m *ProductCatalogItem_ExactPrice) GetOriginalPrice() float32 {
	if m != nil {
		return m.OriginalPrice
	}
	return 0
}

// Product price range when there are a range of prices for different
// variations of the same product.
type ProductCatalogItem_PriceRange struct {
	// Required. The minimum product price.
	Min float32 `protobuf:"fixed32,1,opt,name=min,proto3" json:"min,omitempty"`
	// Required. The maximum product price.
	Max                  float32  `protobuf:"fixed32,2,opt,name=max,proto3" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductCatalogItem_PriceRange) Reset()         { *m = ProductCatalogItem_PriceRange{} }
func (m *ProductCatalogItem_PriceRange) String() string { return proto.CompactTextString(m) }
func (*ProductCatalogItem_PriceRange) ProtoMessage()    {}
func (*ProductCatalogItem_PriceRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_988daa8a4f3967d9, []int{1, 1}
}

func (m *ProductCatalogItem_PriceRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductCatalogItem_PriceRange.Unmarshal(m, b)
}
func (m *ProductCatalogItem_PriceRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductCatalogItem_PriceRange.Marshal(b, m, deterministic)
}
func (m *ProductCatalogItem_PriceRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductCatalogItem_PriceRange.Merge(m, src)
}
func (m *ProductCatalogItem_PriceRange) XXX_Size() int {
	return xxx_messageInfo_ProductCatalogItem_PriceRange.Size(m)
}
func (m *ProductCatalogItem_PriceRange) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductCatalogItem_PriceRange.DiscardUnknown(m)
}

var xxx_messageInfo_ProductCatalogItem_PriceRange proto.InternalMessageInfo

func (m *ProductCatalogItem_PriceRange) GetMin() float32 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *ProductCatalogItem_PriceRange) GetMax() float32 {
	if m != nil {
		return m.Max
	}
	return 0
}

// Catalog item thumbnail/detail image.
type Image struct {
	// Required. URL of the image.
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Optional. Height of the image in number of pixels.
	Height int32 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	// Optional. Width of the image in number of pixels.
	Width                int32    `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_988daa8a4f3967d9, []int{2}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Image) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Image) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func init() {
	proto.RegisterEnum("google.cloud.recommendationengine.v1beta1.ProductCatalogItem_StockState", ProductCatalogItem_StockState_name, ProductCatalogItem_StockState_value)
	proto.RegisterType((*CatalogItem)(nil), "google.cloud.recommendationengine.v1beta1.CatalogItem")
	proto.RegisterType((*CatalogItem_CategoryHierarchy)(nil), "google.cloud.recommendationengine.v1beta1.CatalogItem.CategoryHierarchy")
	proto.RegisterType((*ProductCatalogItem)(nil), "google.cloud.recommendationengine.v1beta1.ProductCatalogItem")
	proto.RegisterMapType((map[string]float32)(nil), "google.cloud.recommendationengine.v1beta1.ProductCatalogItem.CostsEntry")
	proto.RegisterType((*ProductCatalogItem_ExactPrice)(nil), "google.cloud.recommendationengine.v1beta1.ProductCatalogItem.ExactPrice")
	proto.RegisterType((*ProductCatalogItem_PriceRange)(nil), "google.cloud.recommendationengine.v1beta1.ProductCatalogItem.PriceRange")
	proto.RegisterType((*Image)(nil), "google.cloud.recommendationengine.v1beta1.Image")
}

func init() {
	proto.RegisterFile("google/cloud/recommendationengine/v1beta1/catalog.proto", fileDescriptor_988daa8a4f3967d9)
}

var fileDescriptor_988daa8a4f3967d9 = []byte{
	// 837 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x51, 0x6f, 0xdb, 0x36,
	0x10, 0x8e, 0xe5, 0x38, 0x69, 0xce, 0x4e, 0xea, 0x32, 0xed, 0x26, 0x78, 0xc3, 0x60, 0x78, 0x18,
	0x90, 0x3d, 0xcc, 0x6e, 0x53, 0x6c, 0x2b, 0x36, 0xec, 0x21, 0x71, 0x9d, 0xda, 0x28, 0xda, 0x64,
	0x74, 0x02, 0x0c, 0x03, 0x56, 0x81, 0xa1, 0x0e, 0x32, 0x1b, 0x89, 0xd4, 0x28, 0x2a, 0xab, 0xb1,
	0xff, 0xb7, 0xfd, 0xad, 0x81, 0x14, 0xed, 0x78, 0x4b, 0x1e, 0x62, 0xac, 0x6f, 0xbc, 0x3b, 0xde,
	0xf7, 0x7d, 0xe4, 0x7d, 0xa2, 0xe0, 0xfb, 0x44, 0xa9, 0x24, 0xc5, 0x01, 0x4f, 0x55, 0x19, 0x0f,
	0x34, 0x72, 0x95, 0x65, 0x28, 0x63, 0x66, 0x84, 0x92, 0x28, 0x13, 0x21, 0x71, 0x70, 0xfd, 0xec,
	0x12, 0x0d, 0x7b, 0x36, 0xe0, 0xcc, 0xb0, 0x54, 0x25, 0xfd, 0x5c, 0x2b, 0xa3, 0xc8, 0xd7, 0x55,
	0x63, 0xdf, 0x35, 0xf6, 0xef, 0x6a, 0xec, 0xfb, 0xc6, 0xce, 0xe7, 0x9e, 0x83, 0xe5, 0x62, 0xc0,
	0xa4, 0x54, 0xc6, 0x6d, 0x2a, 0x2a, 0xa0, 0xce, 0x77, 0x6b, 0x28, 0x50, 0x59, 0xa6, 0x64, 0xd5,
	0xd7, 0xfb, 0x7b, 0x13, 0x9a, 0xc3, 0x4a, 0xd2, 0xc4, 0x60, 0x46, 0xf6, 0x20, 0x10, 0x71, 0x58,
	0xeb, 0xd6, 0x0e, 0x76, 0x68, 0x20, 0x62, 0xf2, 0x27, 0x3c, 0xe6, 0xcc, 0x60, 0xa2, 0xf4, 0x3c,
	0x9a, 0x09, 0xd4, 0x4c, 0xf3, 0x99, 0xc0, 0x22, 0x0c, 0xba, 0xf5, 0x83, 0xe6, 0xe1, 0xb8, 0x7f,
	0x6f, 0xfd, 0xfd, 0x15, 0x16, 0xbb, 0x76, 0x90, 0x63, 0x8f, 0x38, 0xa7, 0xfb, 0xfc, 0x3f, 0x29,
	0x81, 0x05, 0x79, 0x0c, 0x0d, 0x23, 0x4c, 0x8a, 0x61, 0xdd, 0xe9, 0xa9, 0x02, 0xd2, 0x85, 0x66,
	0x8c, 0x05, 0xd7, 0x22, 0xb7, 0x2c, 0xe1, 0xa6, 0xab, 0xad, 0xa6, 0xc8, 0x3b, 0x78, 0x28, 0x0c,
	0x66, 0x11, 0x33, 0x46, 0x8b, 0xcb, 0xd2, 0x60, 0x11, 0x36, 0xba, 0xb5, 0x83, 0xe6, 0xe1, 0xb7,
	0x6b, 0xe8, 0x3d, 0x41, 0x66, 0x4a, 0x8d, 0x6f, 0x58, 0x4e, 0xf7, 0x2c, 0xda, 0xd1, 0x12, 0x8c,
	0x7c, 0x09, 0xbb, 0x29, 0x93, 0x49, 0xc9, 0x12, 0x8c, 0xb8, 0x8a, 0x31, 0xdc, 0x72, 0x1a, 0x5a,
	0x8b, 0xe4, 0x50, 0xc5, 0x48, 0x08, 0x6c, 0x1a, 0x96, 0x14, 0xe1, 0x83, 0x6e, 0xfd, 0x60, 0x87,
	0xba, 0x35, 0xe9, 0xc1, 0xae, 0x13, 0x96, 0x68, 0x55, 0xe6, 0x91, 0x88, 0xc3, 0x9d, 0x4a, 0xbc,
	0x4d, 0xbe, 0xb2, 0xb9, 0x49, 0x4c, 0xde, 0x43, 0x3b, 0xd7, 0x2a, 0x2e, 0xb9, 0x89, 0x32, 0x34,
	0x2c, 0x66, 0x86, 0x85, 0xe0, 0xd4, 0xff, 0xb4, 0x86, 0xfa, 0xb3, 0x0a, 0x62, 0xe5, 0xd2, 0xc7,
	0x1b, 0xf4, 0xa1, 0x07, 0x7e, 0xe3, 0x71, 0x3b, 0xcf, 0xe1, 0xd1, 0xad, 0x51, 0x90, 0x2f, 0x00,
	0xfc, 0x30, 0xec, 0xa0, 0x6b, 0x4e, 0xfe, 0x4a, 0xe6, 0xf8, 0x09, 0xec, 0xff, 0x9b, 0x3a, 0x32,
	0xf3, 0x1c, 0x7b, 0x7f, 0x6d, 0x03, 0xb9, 0xcd, 0x4a, 0xae, 0xa0, 0x89, 0x1f, 0x18, 0x37, 0x51,
	0xae, 0x05, 0x47, 0xe7, 0xac, 0xf5, 0x7c, 0x73, 0x1b, 0xb3, 0x3f, 0xb2, 0x80, 0x67, 0x16, 0x6f,
	0xbc, 0x41, 0x01, 0x97, 0x91, 0x25, 0x73, 0x34, 0x91, 0x66, 0x32, 0xc1, 0x30, 0xf8, 0x18, 0x64,
	0x0e, 0x99, 0x5a, 0x3c, 0x4b, 0x96, 0x2f, 0x23, 0xf2, 0x0e, 0x1a, 0x5c, 0x15, 0xa6, 0x08, 0xeb,
	0x6b, 0x7f, 0x0b, 0x77, 0xd0, 0x0c, 0x2d, 0xd4, 0x48, 0x1a, 0x3d, 0xa7, 0x15, 0xac, 0x75, 0x19,
	0x2f, 0xb5, 0x46, 0xc9, 0xe7, 0x95, 0xcb, 0x2a, 0xa7, 0xb7, 0x16, 0x49, 0xe7, 0x32, 0x01, 0xcd,
	0xc2, 0x28, 0x7e, 0x15, 0x15, 0x86, 0x19, 0x74, 0x36, 0xdf, 0xfb, 0xbf, 0x52, 0xa6, 0x16, 0x70,
	0x6a, 0xf1, 0x28, 0x14, 0xcb, 0x35, 0xf9, 0x06, 0x08, 0xbb, 0x66, 0x22, 0x65, 0x97, 0x29, 0x46,
	0xbf, 0x97, 0x4c, 0x1a, 0x61, 0xe6, 0xce, 0xfa, 0x75, 0xfa, 0x68, 0x59, 0xf9, 0xd9, 0x17, 0xc8,
	0x21, 0x3c, 0xe1, 0x4c, 0x2a, 0x29, 0x38, 0x4b, 0xa3, 0x85, 0xa3, 0x4b, 0x2d, 0xc2, 0x6d, 0x77,
	0x8c, 0xfd, 0x65, 0xd1, 0x2b, 0xb8, 0xd0, 0x82, 0x8c, 0x61, 0x4b, 0x64, 0x2c, 0xc1, 0xea, 0xab,
	0x69, 0x1e, 0x3e, 0x5d, 0xe3, 0x20, 0x13, 0xdb, 0x48, 0x7d, 0x7f, 0xe7, 0x17, 0x80, 0x1b, 0x97,
	0xd8, 0xab, 0x8c, 0x45, 0x91, 0xa7, 0x6c, 0xbe, 0x62, 0xc3, 0x80, 0xb6, 0x7c, 0xb2, 0xda, 0xf4,
	0x15, 0xec, 0x29, 0x2d, 0x12, 0x21, 0x9d, 0x5e, 0xbb, 0x2b, 0x70, 0xbb, 0x76, 0x17, 0x59, 0xb7,
	0xad, 0xf3, 0x14, 0xe0, 0xc6, 0x12, 0xa4, 0x0d, 0xf5, 0x4c, 0x48, 0x8f, 0x67, 0x97, 0x2e, 0xc3,
	0x3e, 0xf8, 0x5e, 0xbb, 0xec, 0xbc, 0x00, 0xb8, 0x99, 0xae, 0xad, 0x5f, 0xe1, 0xdc, 0x3f, 0xb1,
	0x76, 0x69, 0x9f, 0xb9, 0x6b, 0x96, 0x96, 0x0b, 0xbe, 0x2a, 0xf8, 0x21, 0x78, 0x51, 0xeb, 0xbd,
	0x07, 0xb8, 0x19, 0x06, 0xf9, 0x0c, 0x3e, 0x9d, 0x9e, 0x9f, 0x0e, 0x5f, 0x47, 0xd3, 0xf3, 0xa3,
	0xf3, 0x51, 0x74, 0xf1, 0x76, 0x7a, 0x36, 0x1a, 0x4e, 0x4e, 0x26, 0xa3, 0x97, 0xed, 0x0d, 0xd2,
	0x82, 0x07, 0x93, 0xb7, 0x91, 0xab, 0xb7, 0x37, 0x48, 0x1b, 0x5a, 0xa7, 0x17, 0xe7, 0xd1, 0xe9,
	0x89, 0xcf, 0xd4, 0x6c, 0xfd, 0x8c, 0x8e, 0x4e, 0xe9, 0xcb, 0x11, 0x6d, 0x07, 0x64, 0x17, 0x76,
	0x8e, 0x8f, 0x86, 0xaf, 0xab, 0xb0, 0xde, 0x09, 0xda, 0xb5, 0xe3, 0x6d, 0x68, 0xb8, 0x53, 0xf7,
	0x5e, 0x41, 0xc3, 0xdd, 0xa5, 0x55, 0x6a, 0xe7, 0xe5, 0x95, 0x96, 0x5a, 0x90, 0x4f, 0x60, 0x6b,
	0x86, 0x22, 0x99, 0x19, 0x27, 0xb5, 0x41, 0x7d, 0x64, 0x4f, 0xf0, 0x87, 0x88, 0xcd, 0xcc, 0x3d,
	0xd4, 0x0d, 0x5a, 0x05, 0xc7, 0xd1, 0xaf, 0xbf, 0xf9, 0xf1, 0x25, 0xca, 0x3e, 0x8d, 0x7d, 0xa5,
	0x93, 0x41, 0x82, 0xd2, 0xfd, 0x79, 0x06, 0x55, 0x89, 0xe5, 0xa2, 0xb8, 0xc7, 0x4f, 0xeb, 0xc7,
	0xbb, 0x8a, 0x97, 0x5b, 0x0e, 0xe9, 0xf9, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xff, 0xaf, 0x77,
	0xb8, 0x7f, 0x07, 0x00, 0x00,
}
