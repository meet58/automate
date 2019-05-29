// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/common/rules.proto

package common // import "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/common"

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

type RuleType int32

const (
	RuleType_RULE_TYPE_UNSET RuleType = 0
	RuleType_NODE            RuleType = 1
	RuleType_EVENT           RuleType = 2
)

var RuleType_name = map[int32]string{
	0: "RULE_TYPE_UNSET",
	1: "NODE",
	2: "EVENT",
}
var RuleType_value = map[string]int32{
	"RULE_TYPE_UNSET": 0,
	"NODE":            1,
	"EVENT":           2,
}

func (x RuleType) String() string {
	return proto.EnumName(RuleType_name, int32(x))
}
func (RuleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rules_fb90a59a4ef611f7, []int{0}
}

type ConditionType int32

const (
	ConditionType_CONDITION_TYPE_UNSET ConditionType = 0
	ConditionType_CHEF_SERVERS         ConditionType = 1
	ConditionType_CHEF_ORGS            ConditionType = 2
	ConditionType_CHEF_ENVIRONMENTS    ConditionType = 3
	ConditionType_ROLES                ConditionType = 4
	ConditionType_CHEF_TAGS            ConditionType = 5
	ConditionType_POLICY_GROUP         ConditionType = 6
	ConditionType_POLICY_NAME          ConditionType = 7
)

var ConditionType_name = map[int32]string{
	0: "CONDITION_TYPE_UNSET",
	1: "CHEF_SERVERS",
	2: "CHEF_ORGS",
	3: "CHEF_ENVIRONMENTS",
	4: "ROLES",
	5: "CHEF_TAGS",
	6: "POLICY_GROUP",
	7: "POLICY_NAME",
}
var ConditionType_value = map[string]int32{
	"CONDITION_TYPE_UNSET": 0,
	"CHEF_SERVERS":         1,
	"CHEF_ORGS":            2,
	"CHEF_ENVIRONMENTS":    3,
	"ROLES":                4,
	"CHEF_TAGS":            5,
	"POLICY_GROUP":         6,
	"POLICY_NAME":          7,
}

func (x ConditionType) String() string {
	return proto.EnumName(ConditionType_name, int32(x))
}
func (ConditionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rules_fb90a59a4ef611f7, []int{1}
}

type Rule struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectId            string       `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Name                 string       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type                 RuleType     `protobuf:"varint,4,opt,name=type,proto3,enum=chef.automate.api.iam.v2beta.RuleType" json:"type,omitempty"`
	Conditions           []*Condition `protobuf:"bytes,5,rep,name=conditions,proto3" json:"conditions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Rule) Reset()         { *m = Rule{} }
func (m *Rule) String() string { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()    {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_rules_fb90a59a4ef611f7, []int{0}
}
func (m *Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rule.Unmarshal(m, b)
}
func (m *Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rule.Marshal(b, m, deterministic)
}
func (dst *Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rule.Merge(dst, src)
}
func (m *Rule) XXX_Size() int {
	return xxx_messageInfo_Rule.Size(m)
}
func (m *Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Rule proto.InternalMessageInfo

func (m *Rule) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Rule) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *Rule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Rule) GetType() RuleType {
	if m != nil {
		return m.Type
	}
	return RuleType_RULE_TYPE_UNSET
}

func (m *Rule) GetConditions() []*Condition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

type Condition struct {
	Type                 ConditionType `protobuf:"varint,1,opt,name=type,proto3,enum=chef.automate.api.iam.v2beta.ConditionType" json:"type,omitempty"`
	Values               []string      `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Condition) Reset()         { *m = Condition{} }
func (m *Condition) String() string { return proto.CompactTextString(m) }
func (*Condition) ProtoMessage()    {}
func (*Condition) Descriptor() ([]byte, []int) {
	return fileDescriptor_rules_fb90a59a4ef611f7, []int{1}
}
func (m *Condition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Condition.Unmarshal(m, b)
}
func (m *Condition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Condition.Marshal(b, m, deterministic)
}
func (dst *Condition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Condition.Merge(dst, src)
}
func (m *Condition) XXX_Size() int {
	return xxx_messageInfo_Condition.Size(m)
}
func (m *Condition) XXX_DiscardUnknown() {
	xxx_messageInfo_Condition.DiscardUnknown(m)
}

var xxx_messageInfo_Condition proto.InternalMessageInfo

func (m *Condition) GetType() ConditionType {
	if m != nil {
		return m.Type
	}
	return ConditionType_CONDITION_TYPE_UNSET
}

func (m *Condition) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Rule)(nil), "chef.automate.api.iam.v2beta.Rule")
	proto.RegisterType((*Condition)(nil), "chef.automate.api.iam.v2beta.Condition")
	proto.RegisterEnum("chef.automate.api.iam.v2beta.RuleType", RuleType_name, RuleType_value)
	proto.RegisterEnum("chef.automate.api.iam.v2beta.ConditionType", ConditionType_name, ConditionType_value)
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2beta/common/rules.proto", fileDescriptor_rules_fb90a59a4ef611f7)
}

var fileDescriptor_rules_fb90a59a4ef611f7 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0x26, 0x69, 0x5a, 0x96, 0x77, 0x6c, 0x33, 0xe6, 0x43, 0x39, 0x80, 0x54, 0xed, 0x00, 0xd5,
	0x10, 0x89, 0x54, 0x38, 0x21, 0x21, 0x34, 0x3a, 0x53, 0x82, 0x3a, 0xa7, 0x72, 0xd2, 0x4a, 0xe3,
	0x12, 0xb9, 0x89, 0xd9, 0x8c, 0xea, 0x38, 0x6a, 0x9d, 0xa1, 0xfe, 0x19, 0x7e, 0x12, 0xbf, 0x09,
	0xd5, 0xb4, 0xd5, 0x76, 0x99, 0xd8, 0xcd, 0xef, 0xf3, 0xfa, 0xf9, 0xb0, 0xf5, 0xc0, 0xc7, 0x42,
	0xab, 0x5a, 0x57, 0xa2, 0x32, 0xcb, 0x88, 0x37, 0x46, 0x2b, 0x6e, 0xc4, 0xdb, 0x4b, 0x6e, 0xc4,
	0x2f, 0xbe, 0x8a, 0x78, 0x2d, 0x23, 0xc9, 0x55, 0x74, 0xdd, 0x9f, 0x09, 0xc3, 0xa3, 0x42, 0x2b,
	0xa5, 0xab, 0x68, 0xd1, 0xcc, 0xc5, 0x32, 0xac, 0x17, 0xda, 0x68, 0xfc, 0xa2, 0xb8, 0x12, 0x3f,
	0xc2, 0x2d, 0x31, 0xe4, 0xb5, 0x0c, 0x25, 0x57, 0xe1, 0x3f, 0xc2, 0xf1, 0x1f, 0x07, 0x3c, 0xd6,
	0xcc, 0x05, 0x3e, 0x04, 0x57, 0x96, 0x81, 0xd3, 0x75, 0x7a, 0x3e, 0x73, 0x65, 0x89, 0x5f, 0x02,
	0xd4, 0x0b, 0xfd, 0x53, 0x14, 0x26, 0x97, 0x65, 0xe0, 0x5a, 0xdc, 0xdf, 0x20, 0x71, 0x89, 0x31,
	0x78, 0x15, 0x57, 0x22, 0x68, 0xd9, 0x85, 0x3d, 0xe3, 0x0f, 0xe0, 0x99, 0x55, 0x2d, 0x02, 0xaf,
	0xeb, 0xf4, 0x0e, 0xfb, 0xaf, 0xc2, 0xbb, 0x8c, 0xc3, 0xb5, 0x69, 0xb6, 0xaa, 0x05, 0xb3, 0x1c,
	0x3c, 0x04, 0x28, 0x74, 0x55, 0x4a, 0x23, 0x75, 0xb5, 0x0c, 0xda, 0xdd, 0x56, 0x6f, 0xbf, 0xff,
	0xfa, 0x6e, 0x85, 0xc1, 0xf6, 0x3e, 0xbb, 0x41, 0x3d, 0x2e, 0xc1, 0xdf, 0x2d, 0xf0, 0xa7, 0x4d,
	0x22, 0xc7, 0x26, 0x7a, 0xf3, 0x9f, 0x7a, 0x37, 0x62, 0x3d, 0x87, 0xce, 0x35, 0x9f, 0x37, 0x62,
	0x19, 0xb8, 0xdd, 0x56, 0xcf, 0x67, 0x9b, 0xe9, 0xe4, 0x3d, 0xec, 0x6d, 0x1f, 0x80, 0x9f, 0xc0,
	0x11, 0x9b, 0x8c, 0x48, 0x9e, 0x5d, 0x8c, 0x49, 0x3e, 0xa1, 0x29, 0xc9, 0xd0, 0x03, 0xbc, 0x07,
	0x1e, 0x4d, 0xce, 0x08, 0x72, 0xb0, 0x0f, 0x6d, 0x32, 0x25, 0x34, 0x43, 0xee, 0xc9, 0x6f, 0x07,
	0x0e, 0x6e, 0xb9, 0xe0, 0x00, 0x9e, 0x0e, 0x12, 0x7a, 0x16, 0x67, 0x71, 0x42, 0x6f, 0x0b, 0x20,
	0x78, 0x34, 0xf8, 0x4a, 0xbe, 0xe4, 0x29, 0x61, 0x53, 0xc2, 0x52, 0xe4, 0xe0, 0x03, 0xf0, 0x2d,
	0x92, 0xb0, 0x61, 0x8a, 0x5c, 0xfc, 0x0c, 0x1e, 0xdb, 0x91, 0xd0, 0x69, 0xcc, 0x12, 0x7a, 0x4e,
	0x68, 0x96, 0xa2, 0xd6, 0xda, 0x8e, 0x25, 0x23, 0x92, 0x22, 0x6f, 0x47, 0xc8, 0x4e, 0x87, 0x29,
	0x6a, 0xaf, 0x15, 0xc7, 0xc9, 0x28, 0x1e, 0x5c, 0xe4, 0x43, 0x96, 0x4c, 0xc6, 0xa8, 0x83, 0x8f,
	0x60, 0x7f, 0x83, 0xd0, 0xd3, 0x73, 0x82, 0x1e, 0x7e, 0x1e, 0x7d, 0xff, 0x76, 0x29, 0xcd, 0x55,
	0x33, 0x0b, 0x0b, 0xad, 0xa2, 0xf5, 0x6f, 0xed, 0x1a, 0x17, 0xdd, 0xbb, 0x85, 0xb3, 0x8e, 0x2d,
	0xe0, 0xbb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xda, 0xbb, 0x50, 0xd3, 0xc1, 0x02, 0x00, 0x00,
}