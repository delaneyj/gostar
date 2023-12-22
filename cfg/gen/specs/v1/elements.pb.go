// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: specs/v1/elements.proto

package pb

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

type Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string          `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Name        string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string          `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type        *Attribute_Type `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Attribute) Reset() {
	*x = Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_v1_elements_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute) ProtoMessage() {}

func (x *Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_specs_v1_elements_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute.ProtoReflect.Descriptor instead.
func (*Attribute) Descriptor() ([]byte, []int) {
	return file_specs_v1_elements_proto_rawDescGZIP(), []int{0}
}

func (x *Attribute) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Attribute) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Attribute) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Attribute) GetType() *Attribute_Type {
	if x != nil {
		return x.Type
	}
	return nil
}

type Element struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag         string       `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Name        string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	NoChildren  bool         `protobuf:"varint,4,opt,name=no_children,json=noChildren,proto3" json:"no_children,omitempty"`
	Attributes  []*Attribute `protobuf:"bytes,5,rep,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *Element) Reset() {
	*x = Element{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_v1_elements_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Element) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Element) ProtoMessage() {}

func (x *Element) ProtoReflect() protoreflect.Message {
	mi := &file_specs_v1_elements_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Element.ProtoReflect.Descriptor instead.
func (*Element) Descriptor() ([]byte, []int) {
	return file_specs_v1_elements_proto_rawDescGZIP(), []int{1}
}

func (x *Element) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Element) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Element) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Element) GetNoChildren() bool {
	if x != nil {
		return x.NoChildren
	}
	return false
}

func (x *Element) GetAttributes() []*Attribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type Namespace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description      string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	GlobalAttributes []*Attribute `protobuf:"bytes,3,rep,name=global_attributes,json=globalAttributes,proto3" json:"global_attributes,omitempty"`
	Elements         []*Element   `protobuf:"bytes,4,rep,name=elements,proto3" json:"elements,omitempty"`
}

func (x *Namespace) Reset() {
	*x = Namespace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_v1_elements_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Namespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Namespace) ProtoMessage() {}

func (x *Namespace) ProtoReflect() protoreflect.Message {
	mi := &file_specs_v1_elements_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Namespace.ProtoReflect.Descriptor instead.
func (*Namespace) Descriptor() ([]byte, []int) {
	return file_specs_v1_elements_proto_rawDescGZIP(), []int{2}
}

func (x *Namespace) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Namespace) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Namespace) GetGlobalAttributes() []*Attribute {
	if x != nil {
		return x.GlobalAttributes
	}
	return nil
}

func (x *Namespace) GetElements() []*Element {
	if x != nil {
		return x.Elements
	}
	return nil
}

type Attribute_KV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyValueDelimiter string `protobuf:"bytes,1,opt,name=key_value_delimiter,json=keyValueDelimiter,proto3" json:"key_value_delimiter,omitempty"`
	PairDelimiter     string `protobuf:"bytes,2,opt,name=pair_delimiter,json=pairDelimiter,proto3" json:"pair_delimiter,omitempty"`
}

func (x *Attribute_KV) Reset() {
	*x = Attribute_KV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_v1_elements_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute_KV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute_KV) ProtoMessage() {}

func (x *Attribute_KV) ProtoReflect() protoreflect.Message {
	mi := &file_specs_v1_elements_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute_KV.ProtoReflect.Descriptor instead.
func (*Attribute_KV) Descriptor() ([]byte, []int) {
	return file_specs_v1_elements_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Attribute_KV) GetKeyValueDelimiter() string {
	if x != nil {
		return x.KeyValueDelimiter
	}
	return ""
}

func (x *Attribute_KV) GetPairDelimiter() string {
	if x != nil {
		return x.PairDelimiter
	}
	return ""
}

type Attribute_Choice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Attribute_Choice) Reset() {
	*x = Attribute_Choice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_v1_elements_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute_Choice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute_Choice) ProtoMessage() {}

func (x *Attribute_Choice) ProtoReflect() protoreflect.Message {
	mi := &file_specs_v1_elements_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute_Choice.ProtoReflect.Descriptor instead.
func (*Attribute_Choice) Descriptor() ([]byte, []int) {
	return file_specs_v1_elements_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Attribute_Choice) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Attribute_Choice) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Attribute_Choices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Choices []*Attribute_Choice `protobuf:"bytes,1,rep,name=choices,proto3" json:"choices,omitempty"`
}

func (x *Attribute_Choices) Reset() {
	*x = Attribute_Choices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_v1_elements_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute_Choices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute_Choices) ProtoMessage() {}

func (x *Attribute_Choices) ProtoReflect() protoreflect.Message {
	mi := &file_specs_v1_elements_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute_Choices.ProtoReflect.Descriptor instead.
func (*Attribute_Choices) Descriptor() ([]byte, []int) {
	return file_specs_v1_elements_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Attribute_Choices) GetChoices() []*Attribute_Choice {
	if x != nil {
		return x.Choices
	}
	return nil
}

type Attribute_Type struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*Attribute_Type_String_
	//	*Attribute_Type_Delimited
	//	*Attribute_Type_Kv
	//	*Attribute_Type_Bool
	//	*Attribute_Type_Integer
	//	*Attribute_Type_Number
	//	*Attribute_Type_Choices
	//	*Attribute_Type_Rune
	Type isAttribute_Type_Type `protobuf_oneof:"type"`
}

func (x *Attribute_Type) Reset() {
	*x = Attribute_Type{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specs_v1_elements_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute_Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute_Type) ProtoMessage() {}

func (x *Attribute_Type) ProtoReflect() protoreflect.Message {
	mi := &file_specs_v1_elements_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute_Type.ProtoReflect.Descriptor instead.
func (*Attribute_Type) Descriptor() ([]byte, []int) {
	return file_specs_v1_elements_proto_rawDescGZIP(), []int{0, 3}
}

func (m *Attribute_Type) GetType() isAttribute_Type_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Attribute_Type) GetString_() bool {
	if x, ok := x.GetType().(*Attribute_Type_String_); ok {
		return x.String_
	}
	return false
}

func (x *Attribute_Type) GetDelimited() string {
	if x, ok := x.GetType().(*Attribute_Type_Delimited); ok {
		return x.Delimited
	}
	return ""
}

func (x *Attribute_Type) GetKv() *Attribute_KV {
	if x, ok := x.GetType().(*Attribute_Type_Kv); ok {
		return x.Kv
	}
	return nil
}

func (x *Attribute_Type) GetBool() bool {
	if x, ok := x.GetType().(*Attribute_Type_Bool); ok {
		return x.Bool
	}
	return false
}

func (x *Attribute_Type) GetInteger() bool {
	if x, ok := x.GetType().(*Attribute_Type_Integer); ok {
		return x.Integer
	}
	return false
}

func (x *Attribute_Type) GetNumber() bool {
	if x, ok := x.GetType().(*Attribute_Type_Number); ok {
		return x.Number
	}
	return false
}

func (x *Attribute_Type) GetChoices() *Attribute_Choices {
	if x, ok := x.GetType().(*Attribute_Type_Choices); ok {
		return x.Choices
	}
	return nil
}

func (x *Attribute_Type) GetRune() bool {
	if x, ok := x.GetType().(*Attribute_Type_Rune); ok {
		return x.Rune
	}
	return false
}

type isAttribute_Type_Type interface {
	isAttribute_Type_Type()
}

type Attribute_Type_String_ struct {
	String_ bool `protobuf:"varint,1,opt,name=string,proto3,oneof"`
}

type Attribute_Type_Delimited struct {
	Delimited string `protobuf:"bytes,2,opt,name=delimited,proto3,oneof"`
}

type Attribute_Type_Kv struct {
	Kv *Attribute_KV `protobuf:"bytes,3,opt,name=kv,proto3,oneof"`
}

type Attribute_Type_Bool struct {
	Bool bool `protobuf:"varint,4,opt,name=bool,proto3,oneof"`
}

type Attribute_Type_Integer struct {
	Integer bool `protobuf:"varint,5,opt,name=integer,proto3,oneof"`
}

type Attribute_Type_Number struct {
	Number bool `protobuf:"varint,6,opt,name=number,proto3,oneof"`
}

type Attribute_Type_Choices struct {
	Choices *Attribute_Choices `protobuf:"bytes,7,opt,name=choices,proto3,oneof"`
}

type Attribute_Type_Rune struct {
	Rune bool `protobuf:"varint,8,opt,name=rune,proto3,oneof"`
}

func (*Attribute_Type_String_) isAttribute_Type_Type() {}

func (*Attribute_Type_Delimited) isAttribute_Type_Type() {}

func (*Attribute_Type_Kv) isAttribute_Type_Type() {}

func (*Attribute_Type_Bool) isAttribute_Type_Type() {}

func (*Attribute_Type_Integer) isAttribute_Type_Type() {}

func (*Attribute_Type_Number) isAttribute_Type_Type() {}

func (*Attribute_Type_Choices) isAttribute_Type_Type() {}

func (*Attribute_Type_Rune) isAttribute_Type_Type() {}

var File_specs_v1_elements_proto protoreflect.FileDescriptor

var file_specs_v1_elements_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x70, 0x65, 0x63, 0x73,
	0x2e, 0x76, 0x31, 0x22, 0xef, 0x04, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x5b, 0x0a, 0x02, 0x4b, 0x56, 0x12, 0x2e, 0x0a,
	0x13, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x12, 0x25, 0x0a,
	0x0e, 0x70, 0x61, 0x69, 0x72, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x69, 0x72, 0x44, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x65, 0x72, 0x1a, 0x3e, 0x0a, 0x06, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x3f, 0x0a, 0x07, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x34, 0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x07, 0x63, 0x68,
	0x6f, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x8d, 0x02, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x64,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x02, 0x6b, 0x76, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x4b, 0x56, 0x48, 0x00, 0x52, 0x02,
	0x6b, 0x76, 0x12, 0x14, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x07, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x37,
	0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x2e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x48, 0x00, 0x52, 0x07,
	0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x04, 0x72, 0x75, 0x6e, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x65, 0x42, 0x06, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x07, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x74, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x6f, 0x5f,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x6e, 0x6f, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x33, 0x0a, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22,
	0xb2, 0x01, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x11, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x52, 0x10, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x42, 0x6b, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x70, 0x65, 0x63,
	0x73, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0b, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x70, 0x62, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x53, 0x70, 0x65, 0x63, 0x73,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x53, 0x70, 0x65, 0x63, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x14, 0x53, 0x70, 0x65, 0x63, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x53, 0x70, 0x65, 0x63, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_specs_v1_elements_proto_rawDescOnce sync.Once
	file_specs_v1_elements_proto_rawDescData = file_specs_v1_elements_proto_rawDesc
)

func file_specs_v1_elements_proto_rawDescGZIP() []byte {
	file_specs_v1_elements_proto_rawDescOnce.Do(func() {
		file_specs_v1_elements_proto_rawDescData = protoimpl.X.CompressGZIP(file_specs_v1_elements_proto_rawDescData)
	})
	return file_specs_v1_elements_proto_rawDescData
}

var file_specs_v1_elements_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_specs_v1_elements_proto_goTypes = []interface{}{
	(*Attribute)(nil),         // 0: specs.v1.Attribute
	(*Element)(nil),           // 1: specs.v1.Element
	(*Namespace)(nil),         // 2: specs.v1.Namespace
	(*Attribute_KV)(nil),      // 3: specs.v1.Attribute.KV
	(*Attribute_Choice)(nil),  // 4: specs.v1.Attribute.Choice
	(*Attribute_Choices)(nil), // 5: specs.v1.Attribute.Choices
	(*Attribute_Type)(nil),    // 6: specs.v1.Attribute.Type
}
var file_specs_v1_elements_proto_depIdxs = []int32{
	6, // 0: specs.v1.Attribute.type:type_name -> specs.v1.Attribute.Type
	0, // 1: specs.v1.Element.attributes:type_name -> specs.v1.Attribute
	0, // 2: specs.v1.Namespace.global_attributes:type_name -> specs.v1.Attribute
	1, // 3: specs.v1.Namespace.elements:type_name -> specs.v1.Element
	4, // 4: specs.v1.Attribute.Choices.choices:type_name -> specs.v1.Attribute.Choice
	3, // 5: specs.v1.Attribute.Type.kv:type_name -> specs.v1.Attribute.KV
	5, // 6: specs.v1.Attribute.Type.choices:type_name -> specs.v1.Attribute.Choices
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_specs_v1_elements_proto_init() }
func file_specs_v1_elements_proto_init() {
	if File_specs_v1_elements_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_specs_v1_elements_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute); i {
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
		file_specs_v1_elements_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Element); i {
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
		file_specs_v1_elements_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Namespace); i {
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
		file_specs_v1_elements_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute_KV); i {
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
		file_specs_v1_elements_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute_Choice); i {
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
		file_specs_v1_elements_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute_Choices); i {
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
		file_specs_v1_elements_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute_Type); i {
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
	file_specs_v1_elements_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*Attribute_Type_String_)(nil),
		(*Attribute_Type_Delimited)(nil),
		(*Attribute_Type_Kv)(nil),
		(*Attribute_Type_Bool)(nil),
		(*Attribute_Type_Integer)(nil),
		(*Attribute_Type_Number)(nil),
		(*Attribute_Type_Choices)(nil),
		(*Attribute_Type_Rune)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_specs_v1_elements_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_specs_v1_elements_proto_goTypes,
		DependencyIndexes: file_specs_v1_elements_proto_depIdxs,
		MessageInfos:      file_specs_v1_elements_proto_msgTypes,
	}.Build()
	File_specs_v1_elements_proto = out.File
	file_specs_v1_elements_proto_rawDesc = nil
	file_specs_v1_elements_proto_goTypes = nil
	file_specs_v1_elements_proto_depIdxs = nil
}