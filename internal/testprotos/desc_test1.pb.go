// Code generated by protoc-gen-go. DO NOT EDIT.
// source: desc_test1.proto

/*
Package testprotos is a generated protocol buffer package.

It is generated from these files:
	desc_test1.proto
	desc_test2.proto
	desc_test_defaults.proto
	desc_test_field_types.proto
	desc_test_options.proto
	desc_test_proto3.proto
	desc_test_wellknowntypes.proto

It has these top-level messages:
	TestMessage
	AnotherTestMessage
	Frobnitz
	Whatchamacallit
	Whatzit
	GroupX
	PrimitiveDefaults
	StringAndBytesDefaults
	EnumDefaults
	UnaryFields
	RepeatedFields
	RepeatedPackedFields
	MapKeyFields
	MapValFields
	ReallySimpleMessage
	TestRequest
	TestResponse
	TestWellKnownTypes
*/
package testprotos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Comment for NestedEnum
type TestMessage_NestedEnum int32

const (
	// Comment for VALUE1
	TestMessage_VALUE1 TestMessage_NestedEnum = 1
	// Comment for VALUE2
	TestMessage_VALUE2 TestMessage_NestedEnum = 2
)

var TestMessage_NestedEnum_name = map[int32]string{
	1: "VALUE1",
	2: "VALUE2",
}
var TestMessage_NestedEnum_value = map[string]int32{
	"VALUE1": 1,
	"VALUE2": 2,
}

func (x TestMessage_NestedEnum) Enum() *TestMessage_NestedEnum {
	p := new(TestMessage_NestedEnum)
	*p = x
	return p
}
func (x TestMessage_NestedEnum) String() string {
	return proto.EnumName(TestMessage_NestedEnum_name, int32(x))
}
func (x *TestMessage_NestedEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TestMessage_NestedEnum_value, data, "TestMessage_NestedEnum")
	if err != nil {
		return err
	}
	*x = TestMessage_NestedEnum(value)
	return nil
}
func (TestMessage_NestedEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Comment for DeeplyNestedEnum
type TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum int32

const (
	// Comment for VALUE1
	TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_VALUE1 TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum = 1
	// Comment for VALUE2
	TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_VALUE2 TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum = 2
)

var TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum_name = map[int32]string{
	1: "VALUE1",
	2: "VALUE2",
}
var TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum_value = map[string]int32{
	"VALUE1": 1,
	"VALUE2": 2,
}

func (x TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum) Enum() *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum {
	p := new(TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum)
	*p = x
	return p
}
func (x TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum) String() string {
	return proto.EnumName(TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum_name, int32(x))
}
func (x *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum_value, data, "TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum")
	if err != nil {
		return err
	}
	*x = TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum(value)
	return nil
}
func (TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0, 0, 0}
}

// Comment for TestMessage
type TestMessage struct {
	// Comment for nm
	Nm *TestMessage_NestedMessage `protobuf:"bytes,1,opt,name=nm" json:"nm,omitempty"`
	// Comment for anm
	Anm *TestMessage_NestedMessage_AnotherNestedMessage `protobuf:"bytes,2,opt,name=anm" json:"anm,omitempty"`
	// Comment for yanm
	Yanm *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage `protobuf:"bytes,3,opt,name=yanm" json:"yanm,omitempty"`
	// Comment for ne
	Ne               []TestMessage_NestedEnum `protobuf:"varint,4,rep,name=ne,enum=testprotos.TestMessage_NestedEnum" json:"ne,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *TestMessage) Reset()                    { *m = TestMessage{} }
func (m *TestMessage) String() string            { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()               {}
func (*TestMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TestMessage) GetNm() *TestMessage_NestedMessage {
	if m != nil {
		return m.Nm
	}
	return nil
}

func (m *TestMessage) GetAnm() *TestMessage_NestedMessage_AnotherNestedMessage {
	if m != nil {
		return m.Anm
	}
	return nil
}

func (m *TestMessage) GetYanm() *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage {
	if m != nil {
		return m.Yanm
	}
	return nil
}

func (m *TestMessage) GetNe() []TestMessage_NestedEnum {
	if m != nil {
		return m.Ne
	}
	return nil
}

// Comment for NestedMessage
type TestMessage_NestedMessage struct {
	// Comment for anm
	Anm *TestMessage_NestedMessage_AnotherNestedMessage `protobuf:"bytes,1,opt,name=anm" json:"anm,omitempty"`
	// Comment for yanm
	Yanm             *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage `protobuf:"bytes,2,opt,name=yanm" json:"yanm,omitempty"`
	XXX_unrecognized []byte                                                                  `json:"-"`
}

func (m *TestMessage_NestedMessage) Reset()                    { *m = TestMessage_NestedMessage{} }
func (m *TestMessage_NestedMessage) String() string            { return proto.CompactTextString(m) }
func (*TestMessage_NestedMessage) ProtoMessage()               {}
func (*TestMessage_NestedMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *TestMessage_NestedMessage) GetAnm() *TestMessage_NestedMessage_AnotherNestedMessage {
	if m != nil {
		return m.Anm
	}
	return nil
}

func (m *TestMessage_NestedMessage) GetYanm() *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage {
	if m != nil {
		return m.Yanm
	}
	return nil
}

// Comment for AnotherNestedMessage
type TestMessage_NestedMessage_AnotherNestedMessage struct {
	// Comment for yanm
	Yanm             []*TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage `protobuf:"bytes,1,rep,name=yanm" json:"yanm,omitempty"`
	XXX_unrecognized []byte                                                                    `json:"-"`
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage) Reset() {
	*m = TestMessage_NestedMessage_AnotherNestedMessage{}
}
func (m *TestMessage_NestedMessage_AnotherNestedMessage) String() string {
	return proto.CompactTextString(m)
}
func (*TestMessage_NestedMessage_AnotherNestedMessage) ProtoMessage() {}
func (*TestMessage_NestedMessage_AnotherNestedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage) GetYanm() []*TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage {
	if m != nil {
		return m.Yanm
	}
	return nil
}

var E_TestMessage_NestedMessage_AnotherNestedMessage_Flags = &proto.ExtensionDesc{
	ExtendedType:  (*AnotherTestMessage)(nil),
	ExtensionType: ([]bool)(nil),
	Field:         200,
	Name:          "testprotos.TestMessage.NestedMessage.AnotherNestedMessage.flags",
	Tag:           "varint,200,rep,packed,name=flags",
	Filename:      "desc_test1.proto",
}

// Comment for YetAnotherNestedMessage
type TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage struct {
	// Comment for foo
	Foo *string `protobuf:"bytes,1,opt,name=foo" json:"foo,omitempty"`
	// Comment for bar
	Bar *int32 `protobuf:"varint,2,opt,name=bar" json:"bar,omitempty"`
	// Comment for baz
	Baz []byte `protobuf:"bytes,3,opt,name=baz" json:"baz,omitempty"`
	// Comment for dne
	Dne *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum `protobuf:"varint,4,opt,name=dne,enum=testprotos.TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum" json:"dne,omitempty"`
	// Comment for anm
	Anm *TestMessage_NestedMessage_AnotherNestedMessage `protobuf:"bytes,5,opt,name=anm" json:"anm,omitempty"`
	// Comment for nm
	Nm *TestMessage_NestedMessage `protobuf:"bytes,6,opt,name=nm" json:"nm,omitempty"`
	// Comment for tm
	Tm               *TestMessage `protobuf:"bytes,7,opt,name=tm" json:"tm,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) Reset() {
	*m = TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage{}
}
func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) String() string {
	return proto.CompactTextString(m)
}
func (*TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) ProtoMessage() {}
func (*TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0, 0}
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetFoo() string {
	if m != nil && m.Foo != nil {
		return *m.Foo
	}
	return ""
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetBar() int32 {
	if m != nil && m.Bar != nil {
		return *m.Bar
	}
	return 0
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetBaz() []byte {
	if m != nil {
		return m.Baz
	}
	return nil
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetDne() TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum {
	if m != nil && m.Dne != nil {
		return *m.Dne
	}
	return TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_VALUE1
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetAnm() *TestMessage_NestedMessage_AnotherNestedMessage {
	if m != nil {
		return m.Anm
	}
	return nil
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetNm() *TestMessage_NestedMessage {
	if m != nil {
		return m.Nm
	}
	return nil
}

func (m *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage) GetTm() *TestMessage {
	if m != nil {
		return m.Tm
	}
	return nil
}

// Comment for AnotherTestMessage
type AnotherTestMessage struct {
	// Comment for dne
	Dne *TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum `protobuf:"varint,1,opt,name=dne,enum=testprotos.TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum" json:"dne,omitempty"`
	// Comment for map_field1
	MapField1 map[int32]string `protobuf:"bytes,2,rep,name=map_field1,json=mapField1" json:"map_field1,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Comment for map_field2
	MapField2 map[int64]float32 `protobuf:"bytes,3,rep,name=map_field2,json=mapField2" json:"map_field2,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"fixed32,2,opt,name=value"`
	// Comment for map_field3
	MapField3 map[uint32]bool `protobuf:"bytes,4,rep,name=map_field3,json=mapField3" json:"map_field3,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// Comment for map_field4
	MapField4                    map[string]*AnotherTestMessage `protobuf:"bytes,5,rep,name=map_field4,json=mapField4" json:"map_field4,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Rocknroll                    *AnotherTestMessage_RockNRoll  `protobuf:"group,6,opt,name=RockNRoll,json=rocknroll" json:"rocknroll,omitempty"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *AnotherTestMessage) Reset()                    { *m = AnotherTestMessage{} }
func (m *AnotherTestMessage) String() string            { return proto.CompactTextString(m) }
func (*AnotherTestMessage) ProtoMessage()               {}
func (*AnotherTestMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

var extRange_AnotherTestMessage = []proto.ExtensionRange{
	{100, 200},
}

func (*AnotherTestMessage) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_AnotherTestMessage
}

func (m *AnotherTestMessage) GetDne() TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum {
	if m != nil && m.Dne != nil {
		return *m.Dne
	}
	return TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_VALUE1
}

func (m *AnotherTestMessage) GetMapField1() map[int32]string {
	if m != nil {
		return m.MapField1
	}
	return nil
}

func (m *AnotherTestMessage) GetMapField2() map[int64]float32 {
	if m != nil {
		return m.MapField2
	}
	return nil
}

func (m *AnotherTestMessage) GetMapField3() map[uint32]bool {
	if m != nil {
		return m.MapField3
	}
	return nil
}

func (m *AnotherTestMessage) GetMapField4() map[string]*AnotherTestMessage {
	if m != nil {
		return m.MapField4
	}
	return nil
}

func (m *AnotherTestMessage) GetRocknroll() *AnotherTestMessage_RockNRoll {
	if m != nil {
		return m.Rocknroll
	}
	return nil
}

// Comment for RockNRoll
type AnotherTestMessage_RockNRoll struct {
	// Comment for beatles
	Beatles *string `protobuf:"bytes,7,opt,name=beatles" json:"beatles,omitempty"`
	// Comment for stones
	Stones *string `protobuf:"bytes,8,opt,name=stones" json:"stones,omitempty"`
	// Comment for doors
	Doors            *string `protobuf:"bytes,9,opt,name=doors" json:"doors,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AnotherTestMessage_RockNRoll) Reset()                    { *m = AnotherTestMessage_RockNRoll{} }
func (m *AnotherTestMessage_RockNRoll) String() string            { return proto.CompactTextString(m) }
func (*AnotherTestMessage_RockNRoll) ProtoMessage()               {}
func (*AnotherTestMessage_RockNRoll) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 4} }

func (m *AnotherTestMessage_RockNRoll) GetBeatles() string {
	if m != nil && m.Beatles != nil {
		return *m.Beatles
	}
	return ""
}

func (m *AnotherTestMessage_RockNRoll) GetStones() string {
	if m != nil && m.Stones != nil {
		return *m.Stones
	}
	return ""
}

func (m *AnotherTestMessage_RockNRoll) GetDoors() string {
	if m != nil && m.Doors != nil {
		return *m.Doors
	}
	return ""
}

var E_Xtm = &proto.ExtensionDesc{
	ExtendedType:  (*AnotherTestMessage)(nil),
	ExtensionType: (*TestMessage)(nil),
	Field:         100,
	Name:          "testprotos.xtm",
	Tag:           "bytes,100,opt,name=xtm",
	Filename:      "desc_test1.proto",
}

var E_Xs = &proto.ExtensionDesc{
	ExtendedType:  (*AnotherTestMessage)(nil),
	ExtensionType: (*string)(nil),
	Field:         101,
	Name:          "testprotos.xs",
	Tag:           "bytes,101,opt,name=xs",
	Filename:      "desc_test1.proto",
}

var E_Xi = &proto.ExtensionDesc{
	ExtendedType:  (*AnotherTestMessage)(nil),
	ExtensionType: (*int32)(nil),
	Field:         102,
	Name:          "testprotos.xi",
	Tag:           "varint,102,opt,name=xi",
	Filename:      "desc_test1.proto",
}

var E_Xui = &proto.ExtensionDesc{
	ExtendedType:  (*AnotherTestMessage)(nil),
	ExtensionType: (*uint64)(nil),
	Field:         103,
	Name:          "testprotos.xui",
	Tag:           "varint,103,opt,name=xui",
	Filename:      "desc_test1.proto",
}

func init() {
	proto.RegisterType((*TestMessage)(nil), "testprotos.TestMessage")
	proto.RegisterType((*TestMessage_NestedMessage)(nil), "testprotos.TestMessage.NestedMessage")
	proto.RegisterType((*TestMessage_NestedMessage_AnotherNestedMessage)(nil), "testprotos.TestMessage.NestedMessage.AnotherNestedMessage")
	proto.RegisterType((*TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage)(nil), "testprotos.TestMessage.NestedMessage.AnotherNestedMessage.YetAnotherNestedMessage")
	proto.RegisterType((*AnotherTestMessage)(nil), "testprotos.AnotherTestMessage")
	proto.RegisterType((*AnotherTestMessage_RockNRoll)(nil), "testprotos.AnotherTestMessage.RockNRoll")
	proto.RegisterEnum("testprotos.TestMessage_NestedEnum", TestMessage_NestedEnum_name, TestMessage_NestedEnum_value)
	proto.RegisterEnum("testprotos.TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum", TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum_name, TestMessage_NestedMessage_AnotherNestedMessage_YetAnotherNestedMessage_DeeplyNestedEnum_value)
	proto.RegisterExtension(E_TestMessage_NestedMessage_AnotherNestedMessage_Flags)
	proto.RegisterExtension(E_Xtm)
	proto.RegisterExtension(E_Xs)
	proto.RegisterExtension(E_Xi)
	proto.RegisterExtension(E_Xui)
}

func init() { proto.RegisterFile("desc_test1.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 664 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xdd, 0x6a, 0xdb, 0x3c,
	0x18, 0xc7, 0xb1, 0x1c, 0xa7, 0xf1, 0xd3, 0xb7, 0xc5, 0x88, 0xf2, 0xd6, 0xf4, 0x60, 0x84, 0xb0,
	0xb1, 0x50, 0x98, 0xb3, 0x28, 0x19, 0x6c, 0x61, 0x27, 0x2d, 0x6b, 0x61, 0xd0, 0xf6, 0x40, 0xfb,
	0x80, 0x8d, 0x41, 0x71, 0x63, 0x25, 0xcd, 0x6a, 0x4b, 0xc1, 0x52, 0x46, 0xd2, 0xab, 0xd9, 0xc1,
	0x2e, 0x66, 0xbb, 0x87, 0xdd, 0xc9, 0x0e, 0x36, 0xa4, 0x34, 0x8d, 0xf3, 0xd1, 0xba, 0x1d, 0xed,
	0xce, 0xf4, 0xd8, 0xcf, 0xff, 0xa7, 0xe7, 0x53, 0xe0, 0x45, 0x4c, 0xb6, 0x8f, 0x15, 0x93, 0xaa,
	0x1e, 0xf4, 0x53, 0xa1, 0x04, 0x06, 0x6d, 0x98, 0xa3, 0xac, 0xfc, 0x2a, 0xc1, 0xea, 0x5b, 0x26,
	0xd5, 0x21, 0x93, 0x32, 0xec, 0x32, 0xfc, 0x0c, 0x10, 0x4f, 0x7c, 0xab, 0x6c, 0x55, 0x57, 0xc9,
	0xa3, 0x60, 0xea, 0x18, 0x64, 0x9c, 0x82, 0x23, 0x26, 0x15, 0x8b, 0x2e, 0x2c, 0x8a, 0x78, 0x82,
	0x0f, 0xc0, 0x0e, 0x79, 0xe2, 0x23, 0xa3, 0x6b, 0xdd, 0x48, 0x17, 0xec, 0x70, 0xa1, 0x4e, 0x59,
	0x3a, 0x0b, 0xd3, 0x18, 0xdc, 0x81, 0xc2, 0x48, 0xe3, 0x6c, 0x83, 0xa3, 0x7f, 0x8f, 0x0b, 0x3e,
	0x30, 0xb5, 0xf4, 0x1a, 0xc3, 0xc7, 0x04, 0x10, 0x67, 0x7e, 0xa1, 0x6c, 0x57, 0xd7, 0x49, 0xe5,
	0xfa, 0x5b, 0xf6, 0xf8, 0x20, 0xa1, 0x88, 0xb3, 0xad, 0x6f, 0x45, 0x58, 0x9b, 0x61, 0x4d, 0x72,
	0xb7, 0xee, 0x36, 0x77, 0x74, 0xbf, 0xb9, 0x6f, 0xfd, 0x2c, 0xc0, 0xc6, 0xb2, 0xdf, 0x97, 0x01,
	0x58, 0x65, 0xfb, 0x5e, 0x03, 0xf8, 0x6a, 0xc3, 0xe6, 0x15, 0x1e, 0xd8, 0x03, 0xbb, 0x23, 0x84,
	0x29, 0xa9, 0x4b, 0xf5, 0x51, 0x7f, 0x39, 0x09, 0x53, 0x53, 0x15, 0x87, 0xea, 0xe3, 0xf8, 0xcb,
	0xb9, 0x99, 0x91, 0xff, 0xf4, 0x97, 0x73, 0x3c, 0x00, 0x3b, 0x32, 0xfd, 0xb4, 0xaa, 0xeb, 0xa4,
	0x7d, 0xf7, 0x81, 0x07, 0xaf, 0x18, 0xeb, 0xc7, 0xa3, 0xcc, 0x40, 0xe8, 0xfb, 0x26, 0xfd, 0x77,
	0xee, 0xa6, 0xff, 0xe3, 0x05, 0x2c, 0xde, 0x76, 0x01, 0x1f, 0x03, 0x52, 0x89, 0xbf, 0x62, 0x64,
	0x9b, 0x57, 0xc8, 0x28, 0x52, 0x49, 0x65, 0x1b, 0xbc, 0xf9, 0x34, 0x30, 0x40, 0xf1, 0xfd, 0xce,
	0xc1, 0xbb, 0xbd, 0xba, 0x67, 0x5d, 0x9e, 0x89, 0x87, 0xc8, 0x0b, 0x70, 0x3a, 0x71, 0xd8, 0x95,
	0xf8, 0x41, 0x96, 0x78, 0x11, 0x7b, 0x06, 0xec, 0x7f, 0xd7, 0xc3, 0x52, 0xda, 0x45, 0x9e, 0x45,
	0xc7, 0x8a, 0xca, 0x43, 0x80, 0xfc, 0x0b, 0x2a, 0xbf, 0x8b, 0x80, 0x17, 0x71, 0x93, 0x46, 0x5a,
	0xff, 0xbc, 0x91, 0x90, 0x84, 0xfd, 0xe3, 0x4e, 0x8f, 0xc5, 0x51, 0xdd, 0x47, 0x66, 0xfe, 0x9f,
	0x5c, 0x9f, 0x79, 0x70, 0x18, 0xf6, 0xf7, 0x8d, 0xff, 0x1e, 0x57, 0xe9, 0x88, 0xba, 0xc9, 0xc4,
	0x9e, 0xa1, 0x11, 0xdf, 0xbe, 0x15, 0x8d, 0xcc, 0xd1, 0xc8, 0x0c, 0xad, 0x61, 0x9e, 0xac, 0x9b,
	0xd3, 0x1a, 0x73, 0xb4, 0xc6, 0x0c, 0xad, 0xe9, 0x3b, 0xb7, 0xa2, 0x35, 0xe7, 0x68, 0x4d, 0xbc,
	0x0f, 0x6e, 0x2a, 0xda, 0x67, 0x3c, 0x15, 0x71, 0x6c, 0x26, 0x17, 0x48, 0x35, 0x07, 0x46, 0x45,
	0xfb, 0xec, 0x88, 0x8a, 0x38, 0xa6, 0x53, 0xe9, 0xd6, 0x4b, 0x58, 0x9f, 0x2d, 0xa7, 0xde, 0xf1,
	0x33, 0x36, 0x32, 0x83, 0xe0, 0x50, 0x7d, 0xc4, 0x1b, 0xe0, 0x7c, 0x09, 0xe3, 0x01, 0x33, 0x2f,
	0x81, 0x4b, 0xc7, 0x46, 0x0b, 0x3d, 0xb7, 0xb2, 0x6a, 0xb2, 0xa0, 0xb6, 0x97, 0xa8, 0xd1, 0x15,
	0xea, 0xc6, 0x82, 0x7a, 0x6d, 0x89, 0xba, 0x94, 0x55, 0x7f, 0x9a, 0xaa, 0x9b, 0x0b, 0x6a, 0x77,
	0xac, 0x6e, 0x66, 0xd5, 0xab, 0x24, 0x67, 0xa5, 0xb2, 0xf4, 0x37, 0xe0, 0x5e, 0xd6, 0x0b, 0xfb,
	0xb0, 0x72, 0xc2, 0x42, 0x15, 0x33, 0x69, 0xb6, 0xdd, 0xa5, 0x13, 0x13, 0xff, 0x0f, 0x45, 0xa9,
	0x04, 0x67, 0xd2, 0x2f, 0x99, 0x1f, 0x17, 0x96, 0x0e, 0x3b, 0x12, 0x22, 0x95, 0xbe, 0x3b, 0x2e,
	0x99, 0x31, 0xb6, 0x9d, 0x52, 0xe4, 0xfd, 0xb0, 0x5a, 0xaf, 0xc1, 0x1e, 0xaa, 0x24, 0x77, 0xc1,
	0xa3, 0xeb, 0x1f, 0x16, 0xcd, 0x68, 0x05, 0x80, 0x86, 0xf9, 0x4f, 0x05, 0x33, 0x41, 0xa0, 0xa1,
	0x34, 0xfe, 0xbd, 0x5c, 0xff, 0x8e, 0xe9, 0x3d, 0x1a, 0xf6, 0x5a, 0x4f, 0xc1, 0x1e, 0x0e, 0xf2,
	0x05, 0xdd, 0xb2, 0x55, 0x2d, 0x50, 0xed, 0xba, 0xdb, 0xf8, 0x58, 0xef, 0xf6, 0xd4, 0xe9, 0xe0,
	0x24, 0x68, 0x8b, 0xa4, 0xf6, 0xf9, 0x74, 0x90, 0xf4, 0x6b, 0x46, 0x98, 0xb2, 0x4e, 0xcc, 0xda,
	0xaa, 0xd6, 0xe3, 0x8a, 0xa5, 0x3c, 0x8c, 0x6b, 0x53, 0xe4, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x08, 0x1b, 0xc2, 0x40, 0x30, 0x09, 0x00, 0x00,
}
