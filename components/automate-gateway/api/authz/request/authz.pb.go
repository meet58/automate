// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/authz/request/authz.proto

package request

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type IsAuthorizedReq struct {
	Subject              string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Resource             string   `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	Action               string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAuthorizedReq) Reset()         { *m = IsAuthorizedReq{} }
func (m *IsAuthorizedReq) String() string { return proto.CompactTextString(m) }
func (*IsAuthorizedReq) ProtoMessage()    {}
func (*IsAuthorizedReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa62939b9d22bb69, []int{0}
}

func (m *IsAuthorizedReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorizedReq.Unmarshal(m, b)
}
func (m *IsAuthorizedReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorizedReq.Marshal(b, m, deterministic)
}
func (m *IsAuthorizedReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorizedReq.Merge(m, src)
}
func (m *IsAuthorizedReq) XXX_Size() int {
	return xxx_messageInfo_IsAuthorizedReq.Size(m)
}
func (m *IsAuthorizedReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorizedReq.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorizedReq proto.InternalMessageInfo

func (m *IsAuthorizedReq) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *IsAuthorizedReq) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *IsAuthorizedReq) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type CreatePolicyReq struct {
	Action               string   `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Subjects             []string `protobuf:"bytes,3,rep,name=subjects,proto3" json:"subjects,omitempty"`
	Resource             string   `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePolicyReq) Reset()         { *m = CreatePolicyReq{} }
func (m *CreatePolicyReq) String() string { return proto.CompactTextString(m) }
func (*CreatePolicyReq) ProtoMessage()    {}
func (*CreatePolicyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa62939b9d22bb69, []int{1}
}

func (m *CreatePolicyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePolicyReq.Unmarshal(m, b)
}
func (m *CreatePolicyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePolicyReq.Marshal(b, m, deterministic)
}
func (m *CreatePolicyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePolicyReq.Merge(m, src)
}
func (m *CreatePolicyReq) XXX_Size() int {
	return xxx_messageInfo_CreatePolicyReq.Size(m)
}
func (m *CreatePolicyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePolicyReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePolicyReq proto.InternalMessageInfo

func (m *CreatePolicyReq) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CreatePolicyReq) GetSubjects() []string {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *CreatePolicyReq) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

type ListPoliciesReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPoliciesReq) Reset()         { *m = ListPoliciesReq{} }
func (m *ListPoliciesReq) String() string { return proto.CompactTextString(m) }
func (*ListPoliciesReq) ProtoMessage()    {}
func (*ListPoliciesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa62939b9d22bb69, []int{2}
}

func (m *ListPoliciesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPoliciesReq.Unmarshal(m, b)
}
func (m *ListPoliciesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPoliciesReq.Marshal(b, m, deterministic)
}
func (m *ListPoliciesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPoliciesReq.Merge(m, src)
}
func (m *ListPoliciesReq) XXX_Size() int {
	return xxx_messageInfo_ListPoliciesReq.Size(m)
}
func (m *ListPoliciesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPoliciesReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListPoliciesReq proto.InternalMessageInfo

type DeletePolicyReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePolicyReq) Reset()         { *m = DeletePolicyReq{} }
func (m *DeletePolicyReq) String() string { return proto.CompactTextString(m) }
func (*DeletePolicyReq) ProtoMessage()    {}
func (*DeletePolicyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa62939b9d22bb69, []int{3}
}

func (m *DeletePolicyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePolicyReq.Unmarshal(m, b)
}
func (m *DeletePolicyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePolicyReq.Marshal(b, m, deterministic)
}
func (m *DeletePolicyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePolicyReq.Merge(m, src)
}
func (m *DeletePolicyReq) XXX_Size() int {
	return xxx_messageInfo_DeletePolicyReq.Size(m)
}
func (m *DeletePolicyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePolicyReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePolicyReq proto.InternalMessageInfo

func (m *DeletePolicyReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type IntrospectAllReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntrospectAllReq) Reset()         { *m = IntrospectAllReq{} }
func (m *IntrospectAllReq) String() string { return proto.CompactTextString(m) }
func (*IntrospectAllReq) ProtoMessage()    {}
func (*IntrospectAllReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa62939b9d22bb69, []int{4}
}

func (m *IntrospectAllReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntrospectAllReq.Unmarshal(m, b)
}
func (m *IntrospectAllReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntrospectAllReq.Marshal(b, m, deterministic)
}
func (m *IntrospectAllReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntrospectAllReq.Merge(m, src)
}
func (m *IntrospectAllReq) XXX_Size() int {
	return xxx_messageInfo_IntrospectAllReq.Size(m)
}
func (m *IntrospectAllReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IntrospectAllReq.DiscardUnknown(m)
}

var xxx_messageInfo_IntrospectAllReq proto.InternalMessageInfo

type IntrospectSomeReq struct {
	Paths                []string `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntrospectSomeReq) Reset()         { *m = IntrospectSomeReq{} }
func (m *IntrospectSomeReq) String() string { return proto.CompactTextString(m) }
func (*IntrospectSomeReq) ProtoMessage()    {}
func (*IntrospectSomeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa62939b9d22bb69, []int{5}
}

func (m *IntrospectSomeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntrospectSomeReq.Unmarshal(m, b)
}
func (m *IntrospectSomeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntrospectSomeReq.Marshal(b, m, deterministic)
}
func (m *IntrospectSomeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntrospectSomeReq.Merge(m, src)
}
func (m *IntrospectSomeReq) XXX_Size() int {
	return xxx_messageInfo_IntrospectSomeReq.Size(m)
}
func (m *IntrospectSomeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IntrospectSomeReq.DiscardUnknown(m)
}

var xxx_messageInfo_IntrospectSomeReq proto.InternalMessageInfo

func (m *IntrospectSomeReq) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

type IntrospectReq struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Parameters           []string `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntrospectReq) Reset()         { *m = IntrospectReq{} }
func (m *IntrospectReq) String() string { return proto.CompactTextString(m) }
func (*IntrospectReq) ProtoMessage()    {}
func (*IntrospectReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa62939b9d22bb69, []int{6}
}

func (m *IntrospectReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntrospectReq.Unmarshal(m, b)
}
func (m *IntrospectReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntrospectReq.Marshal(b, m, deterministic)
}
func (m *IntrospectReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntrospectReq.Merge(m, src)
}
func (m *IntrospectReq) XXX_Size() int {
	return xxx_messageInfo_IntrospectReq.Size(m)
}
func (m *IntrospectReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IntrospectReq.DiscardUnknown(m)
}

var xxx_messageInfo_IntrospectReq proto.InternalMessageInfo

func (m *IntrospectReq) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *IntrospectReq) GetParameters() []string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func init() {
	proto.RegisterType((*IsAuthorizedReq)(nil), "chef.automate.api.authz.request.IsAuthorizedReq")
	proto.RegisterType((*CreatePolicyReq)(nil), "chef.automate.api.authz.request.CreatePolicyReq")
	proto.RegisterType((*ListPoliciesReq)(nil), "chef.automate.api.authz.request.ListPoliciesReq")
	proto.RegisterType((*DeletePolicyReq)(nil), "chef.automate.api.authz.request.DeletePolicyReq")
	proto.RegisterType((*IntrospectAllReq)(nil), "chef.automate.api.authz.request.IntrospectAllReq")
	proto.RegisterType((*IntrospectSomeReq)(nil), "chef.automate.api.authz.request.IntrospectSomeReq")
	proto.RegisterType((*IntrospectReq)(nil), "chef.automate.api.authz.request.IntrospectReq")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/authz/request/authz.proto", fileDescriptor_fa62939b9d22bb69)
}

var fileDescriptor_fa62939b9d22bb69 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0xd5, 0xb4, 0x14, 0x38, 0x09, 0x42, 0x2d, 0x84, 0x2c, 0x06, 0x28, 0x99, 0x60, 0x20,
	0x19, 0x98, 0x18, 0x4b, 0x11, 0xa2, 0x12, 0x03, 0x2a, 0x1b, 0x0b, 0x72, 0xdd, 0xa3, 0x31, 0x4a,
	0x62, 0xd7, 0xbe, 0x08, 0xb5, 0x9f, 0x1e, 0xd9, 0x4d, 0xff, 0x8d, 0x6c, 0x79, 0x77, 0xef, 0xdd,
	0x2f, 0xd6, 0x83, 0x47, 0xa9, 0x4b, 0xa3, 0x2b, 0xac, 0xc8, 0x65, 0xa2, 0x26, 0x5d, 0x0a, 0xc2,
	0xfb, 0x99, 0x20, 0xfc, 0x15, 0x8b, 0x4c, 0x18, 0xe5, 0x87, 0xf9, 0x32, 0xb3, 0x38, 0xaf, 0xd1,
	0xd1, 0x4a, 0xa5, 0xc6, 0x6a, 0xd2, 0xec, 0x5a, 0xe6, 0xf8, 0x9d, 0xae, 0x43, 0xa9, 0x30, 0x2a,
	0x5d, 0xad, 0x1b, 0x73, 0xf2, 0x05, 0xf1, 0xc8, 0x0d, 0x6a, 0xca, 0xb5, 0x55, 0x4b, 0x9c, 0x8e,
	0x71, 0xce, 0x38, 0x1c, 0xba, 0x7a, 0xf2, 0x83, 0x92, 0x78, 0xab, 0xdf, 0xba, 0x3d, 0x1e, 0xaf,
	0x25, 0xbb, 0x84, 0x23, 0x8b, 0x4e, 0xd7, 0x56, 0x22, 0x8f, 0xc2, 0x6a, 0xa3, 0xd9, 0x05, 0x74,
	0x85, 0x24, 0xa5, 0x2b, 0xde, 0x0e, 0x9b, 0x46, 0x25, 0x02, 0xe2, 0xa1, 0x45, 0x41, 0xf8, 0xae,
	0x0b, 0x25, 0x17, 0x1e, 0xb0, 0xb5, 0x46, 0xbb, 0x56, 0x7f, 0xbe, 0x21, 0x39, 0xde, 0xee, 0xb7,
	0xfd, 0xf9, 0xb5, 0xde, 0x43, 0x77, 0xf6, 0xd1, 0x49, 0x0f, 0xe2, 0x37, 0xe5, 0x28, 0x00, 0x14,
	0xba, 0x31, 0xce, 0x93, 0x1b, 0x88, 0x9f, 0xb1, 0xc0, 0x5d, 0xea, 0x29, 0x44, 0x6a, 0xda, 0xbc,
	0x28, 0x52, 0xd3, 0x84, 0xc1, 0xd9, 0xa8, 0x22, 0xab, 0x9d, 0x41, 0x49, 0x83, 0xa2, 0xf0, 0xb1,
	0x3b, 0xe8, 0x6d, 0x67, 0x1f, 0xba, 0x44, 0x1f, 0x3c, 0x87, 0x03, 0x23, 0x28, 0x77, 0xbc, 0x15,
	0xfe, 0x69, 0x25, 0x92, 0x21, 0x9c, 0x6c, 0xad, 0xde, 0xc6, 0xa0, 0xe3, 0x37, 0x0d, 0x21, 0x7c,
	0xb3, 0x2b, 0x00, 0x23, 0xac, 0x28, 0x91, 0xd0, 0x3a, 0x1e, 0x85, 0xfc, 0xce, 0xe4, 0xe9, 0xf5,
	0xf3, 0x65, 0xa6, 0x28, 0xaf, 0x27, 0xa9, 0xd4, 0x65, 0xe6, 0xbb, 0xda, 0x14, 0x9c, 0xfd, 0xab,
	0xf4, 0x49, 0x37, 0xf4, 0xfd, 0xf0, 0x17, 0x00, 0x00, 0xff, 0xff, 0x50, 0xd0, 0xd9, 0x1a, 0x2c,
	0x02, 0x00, 0x00,
}
