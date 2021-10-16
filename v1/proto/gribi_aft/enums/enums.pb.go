// gribi_aft.enums is generated by proto_generator as a protobuf
// representation of a YANG schema.
//
// Input schema modules:
//  - github.com/openconfig/gribi/v1/yang/gribi-aft.yang
// Include paths:
//   - github.com/openconfig/gribi/v1/yang/...
//   - github.com/openconfig/gribi/v1/yang/deps/...

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: v1/proto/gribi_aft/enums/enums.proto

package enums

import (
	_ "github.com/openconfig/ygot/proto/yext"
	_ "github.com/openconfig/ygot/proto/ywrapper"
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

// OpenconfigAftTypesEncapsulationHeaderType represents an enumerated type generated for the YANG enumerated type encapsulation-header-type.
type OpenconfigAftTypesEncapsulationHeaderType int32

const (
	OpenconfigAftTypesEncapsulationHeaderType_OPENCONFIGAFTTYPESENCAPSULATIONHEADERTYPE_UNSET OpenconfigAftTypesEncapsulationHeaderType = 0
	OpenconfigAftTypesEncapsulationHeaderType_OPENCONFIGAFTTYPESENCAPSULATIONHEADERTYPE_GRE   OpenconfigAftTypesEncapsulationHeaderType = 1
	OpenconfigAftTypesEncapsulationHeaderType_OPENCONFIGAFTTYPESENCAPSULATIONHEADERTYPE_IPV4  OpenconfigAftTypesEncapsulationHeaderType = 2
	OpenconfigAftTypesEncapsulationHeaderType_OPENCONFIGAFTTYPESENCAPSULATIONHEADERTYPE_IPV6  OpenconfigAftTypesEncapsulationHeaderType = 3
	OpenconfigAftTypesEncapsulationHeaderType_OPENCONFIGAFTTYPESENCAPSULATIONHEADERTYPE_MPLS  OpenconfigAftTypesEncapsulationHeaderType = 4
)

// Enum value maps for OpenconfigAftTypesEncapsulationHeaderType.
var (
	OpenconfigAftTypesEncapsulationHeaderType_name = map[int32]string{
		0: "OPENCONFIGAFTTYPESENCAPSULATIONHEADERTYPE_UNSET",
		1: "OPENCONFIGAFTTYPESENCAPSULATIONHEADERTYPE_GRE",
		2: "OPENCONFIGAFTTYPESENCAPSULATIONHEADERTYPE_IPV4",
		3: "OPENCONFIGAFTTYPESENCAPSULATIONHEADERTYPE_IPV6",
		4: "OPENCONFIGAFTTYPESENCAPSULATIONHEADERTYPE_MPLS",
	}
	OpenconfigAftTypesEncapsulationHeaderType_value = map[string]int32{
		"OPENCONFIGAFTTYPESENCAPSULATIONHEADERTYPE_UNSET": 0,
		"OPENCONFIGAFTTYPESENCAPSULATIONHEADERTYPE_GRE":   1,
		"OPENCONFIGAFTTYPESENCAPSULATIONHEADERTYPE_IPV4":  2,
		"OPENCONFIGAFTTYPESENCAPSULATIONHEADERTYPE_IPV6":  3,
		"OPENCONFIGAFTTYPESENCAPSULATIONHEADERTYPE_MPLS":  4,
	}
)

func (x OpenconfigAftTypesEncapsulationHeaderType) Enum() *OpenconfigAftTypesEncapsulationHeaderType {
	p := new(OpenconfigAftTypesEncapsulationHeaderType)
	*p = x
	return p
}

func (x OpenconfigAftTypesEncapsulationHeaderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpenconfigAftTypesEncapsulationHeaderType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_proto_gribi_aft_enums_enums_proto_enumTypes[0].Descriptor()
}

func (OpenconfigAftTypesEncapsulationHeaderType) Type() protoreflect.EnumType {
	return &file_v1_proto_gribi_aft_enums_enums_proto_enumTypes[0]
}

func (x OpenconfigAftTypesEncapsulationHeaderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpenconfigAftTypesEncapsulationHeaderType.Descriptor instead.
func (OpenconfigAftTypesEncapsulationHeaderType) EnumDescriptor() ([]byte, []int) {
	return file_v1_proto_gribi_aft_enums_enums_proto_rawDescGZIP(), []int{0}
}

// OpenconfigMplsTypesMplsLabelEnum represents an enumerated type generated for the YANG enumerated type mpls-label.
type OpenconfigMplsTypesMplsLabelEnum int32

const (
	OpenconfigMplsTypesMplsLabelEnum_OPENCONFIGMPLSTYPESMPLSLABELENUM_UNSET                   OpenconfigMplsTypesMplsLabelEnum = 0
	OpenconfigMplsTypesMplsLabelEnum_OPENCONFIGMPLSTYPESMPLSLABELENUM_IPV4_EXPLICIT_NULL      OpenconfigMplsTypesMplsLabelEnum = 1
	OpenconfigMplsTypesMplsLabelEnum_OPENCONFIGMPLSTYPESMPLSLABELENUM_ROUTER_ALERT            OpenconfigMplsTypesMplsLabelEnum = 2
	OpenconfigMplsTypesMplsLabelEnum_OPENCONFIGMPLSTYPESMPLSLABELENUM_IPV6_EXPLICIT_NULL      OpenconfigMplsTypesMplsLabelEnum = 3
	OpenconfigMplsTypesMplsLabelEnum_OPENCONFIGMPLSTYPESMPLSLABELENUM_IMPLICIT_NULL           OpenconfigMplsTypesMplsLabelEnum = 4
	OpenconfigMplsTypesMplsLabelEnum_OPENCONFIGMPLSTYPESMPLSLABELENUM_ENTROPY_LABEL_INDICATOR OpenconfigMplsTypesMplsLabelEnum = 8
	OpenconfigMplsTypesMplsLabelEnum_OPENCONFIGMPLSTYPESMPLSLABELENUM_NO_LABEL                OpenconfigMplsTypesMplsLabelEnum = 9
)

// Enum value maps for OpenconfigMplsTypesMplsLabelEnum.
var (
	OpenconfigMplsTypesMplsLabelEnum_name = map[int32]string{
		0: "OPENCONFIGMPLSTYPESMPLSLABELENUM_UNSET",
		1: "OPENCONFIGMPLSTYPESMPLSLABELENUM_IPV4_EXPLICIT_NULL",
		2: "OPENCONFIGMPLSTYPESMPLSLABELENUM_ROUTER_ALERT",
		3: "OPENCONFIGMPLSTYPESMPLSLABELENUM_IPV6_EXPLICIT_NULL",
		4: "OPENCONFIGMPLSTYPESMPLSLABELENUM_IMPLICIT_NULL",
		8: "OPENCONFIGMPLSTYPESMPLSLABELENUM_ENTROPY_LABEL_INDICATOR",
		9: "OPENCONFIGMPLSTYPESMPLSLABELENUM_NO_LABEL",
	}
	OpenconfigMplsTypesMplsLabelEnum_value = map[string]int32{
		"OPENCONFIGMPLSTYPESMPLSLABELENUM_UNSET":                   0,
		"OPENCONFIGMPLSTYPESMPLSLABELENUM_IPV4_EXPLICIT_NULL":      1,
		"OPENCONFIGMPLSTYPESMPLSLABELENUM_ROUTER_ALERT":            2,
		"OPENCONFIGMPLSTYPESMPLSLABELENUM_IPV6_EXPLICIT_NULL":      3,
		"OPENCONFIGMPLSTYPESMPLSLABELENUM_IMPLICIT_NULL":           4,
		"OPENCONFIGMPLSTYPESMPLSLABELENUM_ENTROPY_LABEL_INDICATOR": 8,
		"OPENCONFIGMPLSTYPESMPLSLABELENUM_NO_LABEL":                9,
	}
)

func (x OpenconfigMplsTypesMplsLabelEnum) Enum() *OpenconfigMplsTypesMplsLabelEnum {
	p := new(OpenconfigMplsTypesMplsLabelEnum)
	*p = x
	return p
}

func (x OpenconfigMplsTypesMplsLabelEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpenconfigMplsTypesMplsLabelEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_proto_gribi_aft_enums_enums_proto_enumTypes[1].Descriptor()
}

func (OpenconfigMplsTypesMplsLabelEnum) Type() protoreflect.EnumType {
	return &file_v1_proto_gribi_aft_enums_enums_proto_enumTypes[1]
}

func (x OpenconfigMplsTypesMplsLabelEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpenconfigMplsTypesMplsLabelEnum.Descriptor instead.
func (OpenconfigMplsTypesMplsLabelEnum) EnumDescriptor() ([]byte, []int) {
	return file_v1_proto_gribi_aft_enums_enums_proto_rawDescGZIP(), []int{1}
}

// OpenconfigPacketMatchTypesIPPROTOCOL represents an enumerated type generated for the YANG identity IP_PROTOCOL.
type OpenconfigPacketMatchTypesIPPROTOCOL int32

const (
	OpenconfigPacketMatchTypesIPPROTOCOL_OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_UNSET   OpenconfigPacketMatchTypesIPPROTOCOL = 0
	OpenconfigPacketMatchTypesIPPROTOCOL_OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_GRE  OpenconfigPacketMatchTypesIPPROTOCOL = 24050007
	OpenconfigPacketMatchTypesIPPROTOCOL_OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_TCP  OpenconfigPacketMatchTypesIPPROTOCOL = 38721802
	OpenconfigPacketMatchTypesIPPROTOCOL_OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_L2TP OpenconfigPacketMatchTypesIPPROTOCOL = 81903923
	OpenconfigPacketMatchTypesIPPROTOCOL_OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_AUTH OpenconfigPacketMatchTypesIPPROTOCOL = 203904199
	OpenconfigPacketMatchTypesIPPROTOCOL_OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_PIM  OpenconfigPacketMatchTypesIPPROTOCOL = 272904165
	OpenconfigPacketMatchTypesIPPROTOCOL_OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_IGMP OpenconfigPacketMatchTypesIPPROTOCOL = 512271866
	OpenconfigPacketMatchTypesIPPROTOCOL_OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_ICMP OpenconfigPacketMatchTypesIPPROTOCOL = 512860246
	OpenconfigPacketMatchTypesIPPROTOCOL_OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_UDP  OpenconfigPacketMatchTypesIPPROTOCOL = 525100026
	OpenconfigPacketMatchTypesIPPROTOCOL_OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_RSVP OpenconfigPacketMatchTypesIPPROTOCOL = 530070378
)

// Enum value maps for OpenconfigPacketMatchTypesIPPROTOCOL.
var (
	OpenconfigPacketMatchTypesIPPROTOCOL_name = map[int32]string{
		0:         "OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_UNSET",
		24050007:  "OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_GRE",
		38721802:  "OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_TCP",
		81903923:  "OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_L2TP",
		203904199: "OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_AUTH",
		272904165: "OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_PIM",
		512271866: "OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_IGMP",
		512860246: "OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_ICMP",
		525100026: "OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_UDP",
		530070378: "OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_RSVP",
	}
	OpenconfigPacketMatchTypesIPPROTOCOL_value = map[string]int32{
		"OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_UNSET":   0,
		"OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_GRE":  24050007,
		"OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_TCP":  38721802,
		"OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_L2TP": 81903923,
		"OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_AUTH": 203904199,
		"OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_PIM":  272904165,
		"OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_IGMP": 512271866,
		"OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_ICMP": 512860246,
		"OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_UDP":  525100026,
		"OPENCONFIGPACKETMATCHTYPESIPPROTOCOL_IP_RSVP": 530070378,
	}
)

func (x OpenconfigPacketMatchTypesIPPROTOCOL) Enum() *OpenconfigPacketMatchTypesIPPROTOCOL {
	p := new(OpenconfigPacketMatchTypesIPPROTOCOL)
	*p = x
	return p
}

func (x OpenconfigPacketMatchTypesIPPROTOCOL) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpenconfigPacketMatchTypesIPPROTOCOL) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_proto_gribi_aft_enums_enums_proto_enumTypes[2].Descriptor()
}

func (OpenconfigPacketMatchTypesIPPROTOCOL) Type() protoreflect.EnumType {
	return &file_v1_proto_gribi_aft_enums_enums_proto_enumTypes[2]
}

func (x OpenconfigPacketMatchTypesIPPROTOCOL) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpenconfigPacketMatchTypesIPPROTOCOL.Descriptor instead.
func (OpenconfigPacketMatchTypesIPPROTOCOL) EnumDescriptor() ([]byte, []int) {
	return file_v1_proto_gribi_aft_enums_enums_proto_rawDescGZIP(), []int{2}
}

var File_v1_proto_gribi_aft_enums_enums_proto protoreflect.FileDescriptor

var file_v1_proto_gribi_aft_enums_enums_proto_rawDesc = []byte{
	0x0a, 0x24, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x69, 0x62, 0x69,
	0x5f, 0x61, 0x66, 0x74, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x72, 0x69, 0x62, 0x69, 0x5f, 0x61, 0x66,
	0x74, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x79,
	0x67, 0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x2f, 0x79, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x79, 0x67, 0x6f, 0x74, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x65, 0x78, 0x74, 0x2f, 0x79, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0xd2, 0x02, 0x0a, 0x29, 0x4f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x41, 0x66, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x45, 0x6e, 0x63, 0x61, 0x70, 0x73,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x33, 0x0a, 0x2f, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x41,
	0x46, 0x54, 0x54, 0x59, 0x50, 0x45, 0x53, 0x45, 0x4e, 0x43, 0x41, 0x50, 0x53, 0x55, 0x4c, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x39, 0x0a, 0x2d, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f,
	0x4e, 0x46, 0x49, 0x47, 0x41, 0x46, 0x54, 0x54, 0x59, 0x50, 0x45, 0x53, 0x45, 0x4e, 0x43, 0x41,
	0x50, 0x53, 0x55, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x47, 0x52, 0x45, 0x10, 0x01, 0x1a, 0x06, 0x82, 0x41, 0x03, 0x47, 0x52,
	0x45, 0x12, 0x3b, 0x0a, 0x2e, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x41,
	0x46, 0x54, 0x54, 0x59, 0x50, 0x45, 0x53, 0x45, 0x4e, 0x43, 0x41, 0x50, 0x53, 0x55, 0x4c, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49,
	0x50, 0x56, 0x34, 0x10, 0x02, 0x1a, 0x07, 0x82, 0x41, 0x04, 0x49, 0x50, 0x56, 0x34, 0x12, 0x3b,
	0x0a, 0x2e, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x41, 0x46, 0x54, 0x54,
	0x59, 0x50, 0x45, 0x53, 0x45, 0x4e, 0x43, 0x41, 0x50, 0x53, 0x55, 0x4c, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x50, 0x56, 0x36,
	0x10, 0x03, 0x1a, 0x07, 0x82, 0x41, 0x04, 0x49, 0x50, 0x56, 0x36, 0x12, 0x3b, 0x0a, 0x2e, 0x4f,
	0x50, 0x45, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x41, 0x46, 0x54, 0x54, 0x59, 0x50, 0x45,
	0x53, 0x45, 0x4e, 0x43, 0x41, 0x50, 0x53, 0x55, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x48, 0x45,
	0x41, 0x44, 0x45, 0x52, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x50, 0x4c, 0x53, 0x10, 0x04, 0x1a,
	0x07, 0x82, 0x41, 0x04, 0x4d, 0x50, 0x4c, 0x53, 0x2a, 0x8e, 0x04, 0x0a, 0x20, 0x4f, 0x70, 0x65,
	0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x70, 0x6c, 0x73, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x4d, 0x70, 0x6c, 0x73, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x2a, 0x0a,
	0x26, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x4d, 0x50, 0x4c, 0x53, 0x54,
	0x59, 0x50, 0x45, 0x53, 0x4d, 0x50, 0x4c, 0x53, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x45, 0x4e, 0x55,
	0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x4e, 0x0a, 0x33, 0x4f, 0x50, 0x45,
	0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x4d, 0x50, 0x4c, 0x53, 0x54, 0x59, 0x50, 0x45, 0x53,
	0x4d, 0x50, 0x4c, 0x53, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x49, 0x50,
	0x56, 0x34, 0x5f, 0x45, 0x58, 0x50, 0x4c, 0x49, 0x43, 0x49, 0x54, 0x5f, 0x4e, 0x55, 0x4c, 0x4c,
	0x10, 0x01, 0x1a, 0x15, 0x82, 0x41, 0x12, 0x49, 0x50, 0x56, 0x34, 0x5f, 0x45, 0x58, 0x50, 0x4c,
	0x49, 0x43, 0x49, 0x54, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x12, 0x42, 0x0a, 0x2d, 0x4f, 0x50, 0x45,
	0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x4d, 0x50, 0x4c, 0x53, 0x54, 0x59, 0x50, 0x45, 0x53,
	0x4d, 0x50, 0x4c, 0x53, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x52, 0x4f,
	0x55, 0x54, 0x45, 0x52, 0x5f, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x10, 0x02, 0x1a, 0x0f, 0x82, 0x41,
	0x0c, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x52, 0x5f, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x12, 0x4e, 0x0a,
	0x33, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x4d, 0x50, 0x4c, 0x53, 0x54,
	0x59, 0x50, 0x45, 0x53, 0x4d, 0x50, 0x4c, 0x53, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x45, 0x4e, 0x55,
	0x4d, 0x5f, 0x49, 0x50, 0x56, 0x36, 0x5f, 0x45, 0x58, 0x50, 0x4c, 0x49, 0x43, 0x49, 0x54, 0x5f,
	0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x03, 0x1a, 0x15, 0x82, 0x41, 0x12, 0x49, 0x50, 0x56, 0x36, 0x5f,
	0x45, 0x58, 0x50, 0x4c, 0x49, 0x43, 0x49, 0x54, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x12, 0x44, 0x0a,
	0x2e, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x4d, 0x50, 0x4c, 0x53, 0x54,
	0x59, 0x50, 0x45, 0x53, 0x4d, 0x50, 0x4c, 0x53, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x45, 0x4e, 0x55,
	0x4d, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x49, 0x43, 0x49, 0x54, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x10,
	0x04, 0x1a, 0x10, 0x82, 0x41, 0x0d, 0x49, 0x4d, 0x50, 0x4c, 0x49, 0x43, 0x49, 0x54, 0x5f, 0x4e,
	0x55, 0x4c, 0x4c, 0x12, 0x58, 0x0a, 0x38, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49,
	0x47, 0x4d, 0x50, 0x4c, 0x53, 0x54, 0x59, 0x50, 0x45, 0x53, 0x4d, 0x50, 0x4c, 0x53, 0x4c, 0x41,
	0x42, 0x45, 0x4c, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x45, 0x4e, 0x54, 0x52, 0x4f, 0x50, 0x59, 0x5f,
	0x4c, 0x41, 0x42, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x44, 0x49, 0x43, 0x41, 0x54, 0x4f, 0x52, 0x10,
	0x08, 0x1a, 0x1a, 0x82, 0x41, 0x17, 0x45, 0x4e, 0x54, 0x52, 0x4f, 0x50, 0x59, 0x5f, 0x4c, 0x41,
	0x42, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x44, 0x49, 0x43, 0x41, 0x54, 0x4f, 0x52, 0x12, 0x3a, 0x0a,
	0x29, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x4d, 0x50, 0x4c, 0x53, 0x54,
	0x59, 0x50, 0x45, 0x53, 0x4d, 0x50, 0x4c, 0x53, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x45, 0x4e, 0x55,
	0x4d, 0x5f, 0x4e, 0x4f, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x10, 0x09, 0x1a, 0x0b, 0x82, 0x41,
	0x08, 0x4e, 0x4f, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x2a, 0x9c, 0x05, 0x0a, 0x24, 0x4f, 0x70,
	0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x73, 0x49, 0x50, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43,
	0x4f, 0x4c, 0x12, 0x2e, 0x0a, 0x2a, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47,
	0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x54, 0x59, 0x50, 0x45, 0x53,
	0x49, 0x50, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54,
	0x10, 0x00, 0x12, 0x3d, 0x0a, 0x2b, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47,
	0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x54, 0x59, 0x50, 0x45, 0x53,
	0x49, 0x50, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x49, 0x50, 0x5f, 0x47, 0x52,
	0x45, 0x10, 0xd7, 0xf2, 0xbb, 0x0b, 0x1a, 0x09, 0x82, 0x41, 0x06, 0x49, 0x50, 0x5f, 0x47, 0x52,
	0x45, 0x12, 0x3d, 0x0a, 0x2b, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x50,
	0x41, 0x43, 0x4b, 0x45, 0x54, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x54, 0x59, 0x50, 0x45, 0x53, 0x49,
	0x50, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x49, 0x50, 0x5f, 0x54, 0x43, 0x50,
	0x10, 0x8a, 0xb2, 0xbb, 0x12, 0x1a, 0x09, 0x82, 0x41, 0x06, 0x49, 0x50, 0x5f, 0x54, 0x43, 0x50,
	0x12, 0x3f, 0x0a, 0x2c, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x50, 0x41,
	0x43, 0x4b, 0x45, 0x54, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x54, 0x59, 0x50, 0x45, 0x53, 0x49, 0x50,
	0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x49, 0x50, 0x5f, 0x4c, 0x32, 0x54, 0x50,
	0x10, 0xb3, 0x82, 0x87, 0x27, 0x1a, 0x0a, 0x82, 0x41, 0x07, 0x49, 0x50, 0x5f, 0x4c, 0x32, 0x54,
	0x50, 0x12, 0x3f, 0x0a, 0x2c, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x50,
	0x41, 0x43, 0x4b, 0x45, 0x54, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x54, 0x59, 0x50, 0x45, 0x53, 0x49,
	0x50, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x49, 0x50, 0x5f, 0x41, 0x55, 0x54,
	0x48, 0x10, 0xc7, 0xa9, 0x9d, 0x61, 0x1a, 0x0a, 0x82, 0x41, 0x07, 0x49, 0x50, 0x5f, 0x41, 0x55,
	0x54, 0x48, 0x12, 0x3e, 0x0a, 0x2b, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47,
	0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x54, 0x59, 0x50, 0x45, 0x53,
	0x49, 0x50, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x49, 0x50, 0x5f, 0x50, 0x49,
	0x4d, 0x10, 0xe5, 0xdf, 0x90, 0x82, 0x01, 0x1a, 0x09, 0x82, 0x41, 0x06, 0x49, 0x50, 0x5f, 0x50,
	0x49, 0x4d, 0x12, 0x40, 0x0a, 0x2c, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47,
	0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x54, 0x59, 0x50, 0x45, 0x53,
	0x49, 0x50, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x49, 0x50, 0x5f, 0x49, 0x47,
	0x4d, 0x50, 0x10, 0xfa, 0xcb, 0xa2, 0xf4, 0x01, 0x1a, 0x0a, 0x82, 0x41, 0x07, 0x49, 0x50, 0x5f,
	0x49, 0x47, 0x4d, 0x50, 0x12, 0x40, 0x0a, 0x2c, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f, 0x4e, 0x46,
	0x49, 0x47, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x54, 0x59, 0x50,
	0x45, 0x53, 0x49, 0x50, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x49, 0x50, 0x5f,
	0x49, 0x43, 0x4d, 0x50, 0x10, 0xd6, 0xc0, 0xc6, 0xf4, 0x01, 0x1a, 0x0a, 0x82, 0x41, 0x07, 0x49,
	0x50, 0x5f, 0x49, 0x43, 0x4d, 0x50, 0x12, 0x3e, 0x0a, 0x2b, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f,
	0x4e, 0x46, 0x49, 0x47, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x54,
	0x59, 0x50, 0x45, 0x53, 0x49, 0x50, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x49,
	0x50, 0x5f, 0x55, 0x44, 0x50, 0x10, 0xfa, 0xc7, 0xb1, 0xfa, 0x01, 0x1a, 0x09, 0x82, 0x41, 0x06,
	0x49, 0x50, 0x5f, 0x55, 0x44, 0x50, 0x12, 0x40, 0x0a, 0x2c, 0x4f, 0x50, 0x45, 0x4e, 0x43, 0x4f,
	0x4e, 0x46, 0x49, 0x47, 0x50, 0x41, 0x43, 0x4b, 0x45, 0x54, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x54,
	0x59, 0x50, 0x45, 0x53, 0x49, 0x50, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x49,
	0x50, 0x5f, 0x52, 0x53, 0x56, 0x50, 0x10, 0xea, 0xf6, 0xe0, 0xfc, 0x01, 0x1a, 0x0a, 0x82, 0x41,
	0x07, 0x49, 0x50, 0x5f, 0x52, 0x53, 0x56, 0x50, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x67, 0x72, 0x69, 0x62, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x72, 0x69, 0x62, 0x69, 0x5f, 0x61, 0x66, 0x74, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_proto_gribi_aft_enums_enums_proto_rawDescOnce sync.Once
	file_v1_proto_gribi_aft_enums_enums_proto_rawDescData = file_v1_proto_gribi_aft_enums_enums_proto_rawDesc
)

func file_v1_proto_gribi_aft_enums_enums_proto_rawDescGZIP() []byte {
	file_v1_proto_gribi_aft_enums_enums_proto_rawDescOnce.Do(func() {
		file_v1_proto_gribi_aft_enums_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_proto_gribi_aft_enums_enums_proto_rawDescData)
	})
	return file_v1_proto_gribi_aft_enums_enums_proto_rawDescData
}

var file_v1_proto_gribi_aft_enums_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_v1_proto_gribi_aft_enums_enums_proto_goTypes = []interface{}{
	(OpenconfigAftTypesEncapsulationHeaderType)(0), // 0: gribi_aft.enums.OpenconfigAftTypesEncapsulationHeaderType
	(OpenconfigMplsTypesMplsLabelEnum)(0),          // 1: gribi_aft.enums.OpenconfigMplsTypesMplsLabelEnum
	(OpenconfigPacketMatchTypesIPPROTOCOL)(0),      // 2: gribi_aft.enums.OpenconfigPacketMatchTypesIPPROTOCOL
}
var file_v1_proto_gribi_aft_enums_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_proto_gribi_aft_enums_enums_proto_init() }
func file_v1_proto_gribi_aft_enums_enums_proto_init() {
	if File_v1_proto_gribi_aft_enums_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_proto_gribi_aft_enums_enums_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_proto_gribi_aft_enums_enums_proto_goTypes,
		DependencyIndexes: file_v1_proto_gribi_aft_enums_enums_proto_depIdxs,
		EnumInfos:         file_v1_proto_gribi_aft_enums_enums_proto_enumTypes,
	}.Build()
	File_v1_proto_gribi_aft_enums_enums_proto = out.File
	file_v1_proto_gribi_aft_enums_enums_proto_rawDesc = nil
	file_v1_proto_gribi_aft_enums_enums_proto_goTypes = nil
	file_v1_proto_gribi_aft_enums_enums_proto_depIdxs = nil
}
