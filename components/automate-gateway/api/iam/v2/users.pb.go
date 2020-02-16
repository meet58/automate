// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/users.proto

package v2

import (
	context "context"
	fmt "fmt"
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2/request"
	response "github.com/chef/automate/components/automate-gateway/api/iam/v2/response"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/api"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/users.proto", fileDescriptor_8e7276f80918e43e)
}

var fileDescriptor_8e7276f80918e43e = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0x4d, 0x6e, 0xd3, 0x40,
	0x14, 0xc7, 0x65, 0xa0, 0xa1, 0x1d, 0x28, 0x6d, 0x1f, 0xa5, 0xb2, 0x0c, 0x2b, 0xab, 0x50, 0x14,
	0xa9, 0xb6, 0x14, 0x58, 0x85, 0x05, 0x9f, 0x12, 0x1b, 0x56, 0xa0, 0x6e, 0xb2, 0x9b, 0x3a, 0x2f,
	0x8e, 0x25, 0xdb, 0x33, 0xf1, 0x8c, 0x8b, 0x22, 0x84, 0x90, 0x22, 0xc1, 0x22, 0x5b, 0x2e, 0xc2,
	0xc6, 0x47, 0xe0, 0x00, 0xc0, 0x05, 0x58, 0x20, 0xc1, 0x31, 0xd0, 0x8c, 0x13, 0x7f, 0x28, 0x31,
	0x31, 0x4b, 0xcf, 0xfb, 0x3d, 0xcf, 0xff, 0xe7, 0xcc, 0x9b, 0x90, 0x9e, 0xc7, 0x22, 0xce, 0x62,
	0x8c, 0xa5, 0x70, 0x69, 0x2a, 0x59, 0x44, 0x25, 0x9e, 0xfa, 0x54, 0xe2, 0x5b, 0x3a, 0x75, 0x29,
	0x0f, 0xdc, 0x80, 0x46, 0xee, 0x45, 0xcf, 0x4d, 0x05, 0x26, 0xc2, 0xe1, 0x09, 0x93, 0x0c, 0x4c,
	0x6f, 0x8c, 0x23, 0x67, 0x49, 0x3b, 0x94, 0x07, 0x4e, 0x40, 0x23, 0xe7, 0xa2, 0x67, 0xdd, 0xf1,
	0x19, 0xf3, 0x43, 0xd4, 0x8d, 0x34, 0x8e, 0x99, 0xa4, 0x32, 0x60, 0xf1, 0xa2, 0xcf, 0xea, 0xb7,
	0xdc, 0x2b, 0xc1, 0x49, 0x8a, 0x42, 0x56, 0xf7, 0xb4, 0x1e, 0xb5, 0xee, 0x15, 0x9c, 0xc5, 0x02,
	0x6b, 0xcd, 0x4f, 0xd6, 0x36, 0x27, 0xdc, 0x73, 0x75, 0xdd, 0x3b, 0xf5, 0x31, 0x3e, 0xe5, 0x2c,
	0x0c, 0xbc, 0x69, 0x43, 0xf4, 0xff, 0x79, 0x83, 0x4a, 0xb2, 0xf2, 0x86, 0xde, 0x17, 0x42, 0xb6,
	0xce, 0x54, 0x26, 0xf8, 0x6e, 0x10, 0xf2, 0x3c, 0x41, 0x2a, 0x51, 0x3d, 0xc3, 0x89, 0xd3, 0xf4,
	0x39, 0x9d, 0x92, 0x7a, 0x8d, 0x13, 0xeb, 0x7e, 0x3b, 0x50, 0x70, 0x7b, 0x32, 0xcb, 0xcc, 0xeb,
	0x84, 0xd0, 0x54, 0x8e, 0xfb, 0xfa, 0x3b, 0xcc, 0x32, 0x73, 0x1b, 0x3a, 0x9e, 0xa6, 0xe6, 0x99,
	0x79, 0x8d, 0xec, 0x04, 0x34, 0xca, 0x4b, 0xf3, 0xcc, 0x04, 0xd8, 0x2f, 0x1e, 0xfb, 0x39, 0x34,
	0xfb, 0xf1, 0xeb, 0xf3, 0xa5, 0xae, 0xbd, 0x5b, 0xfb, 0xf9, 0xfb, 0x46, 0x77, 0x70, 0x64, 0x1f,
	0x2c, 0xd6, 0xce, 0x51, 0xd2, 0x62, 0x1d, 0xbe, 0x1a, 0x64, 0xe7, 0x55, 0x20, 0x64, 0xae, 0x78,
	0xaf, 0x39, 0x6a, 0x01, 0x29, 0xa5, 0x93, 0x56, 0x9c, 0xe0, 0xb6, 0xbf, 0xc6, 0xa8, 0x03, 0x57,
	0x12, 0xa4, 0xc3, 0x55, 0x9f, 0x7d, 0xb8, 0x51, 0xfa, 0x84, 0x81, 0x90, 0xda, 0xe6, 0x18, 0xea,
	0x36, 0x83, 0x9b, 0xb0, 0xaa, 0x02, 0xdf, 0x0c, 0x72, 0xf5, 0x25, 0xea, 0x9d, 0xe1, 0xb8, 0x39,
	0xdd, 0x02, 0x51, 0x0e, 0x77, 0x5b, 0x50, 0x82, 0xdb, 0xd3, 0x59, 0x66, 0x1e, 0x90, 0xbd, 0xd2,
	0xa0, 0xff, 0x2e, 0x18, 0xbe, 0x9f, 0x65, 0xe6, 0x16, 0x5c, 0xf6, 0x51, 0xaa, 0xd8, 0xa4, 0x12,
	0x5b, 0x15, 0xe7, 0x99, 0xb9, 0x07, 0xbb, 0xe5, 0x9a, 0x8f, 0xb9, 0x89, 0x03, 0x50, 0x33, 0x71,
	0x15, 0x3c, 0x30, 0xe1, 0x68, 0x45, 0x47, 0x57, 0xe0, 0xa7, 0x41, 0xc8, 0x0b, 0x0c, 0x71, 0xf3,
	0x81, 0x2b, 0xa9, 0x0d, 0x07, 0xae, 0x0a, 0x0a, 0x6e, 0x7f, 0x68, 0x92, 0xdb, 0x86, 0xce, 0x50,
	0xa3, 0x0d, 0x7e, 0xb5, 0xa3, 0x97, 0x93, 0xb9, 0x62, 0x77, 0xad, 0x62, 0xb7, 0x49, 0xf1, 0xb7,
	0x41, 0xc8, 0x19, 0x1f, 0xb6, 0x98, 0xa9, 0x92, 0xda, 0xa0, 0x58, 0x05, 0x05, 0xb7, 0x3f, 0x1a,
	0xff, 0x70, 0x4c, 0x35, 0xdb, 0xc6, 0x31, 0x27, 0xb5, 0xe3, 0x43, 0x6b, 0x8d, 0xa3, 0x9a, 0xb1,
	0xdb, 0x56, 0x83, 0xa6, 0x1a, 0xb4, 0x3f, 0x85, 0xe9, 0x1b, 0x0c, 0x47, 0x9b, 0x4d, 0x15, 0xd5,
	0xca, 0x34, 0x07, 0x05, 0xb7, 0x3f, 0x19, 0xf9, 0xb0, 0x35, 0x4a, 0x1e, 0x12, 0x28, 0x84, 0x54,
	0xd7, 0x52, 0xf4, 0x08, 0x0e, 0xeb, 0xeb, 0x15, 0xd9, 0x9e, 0xb5, 0x1c, 0x36, 0x57, 0x60, 0x38,
	0x2a, 0x5c, 0x2d, 0xeb, 0x56, 0xd5, 0xb5, 0x5a, 0x7b, 0xf6, 0x74, 0xf0, 0xd8, 0x0f, 0xe4, 0x38,
	0x3d, 0x77, 0x3c, 0x16, 0xb9, 0x2a, 0x7e, 0x71, 0xf7, 0xba, 0xed, 0xfe, 0x0e, 0xce, 0x3b, 0xfa,
	0xf2, 0x7d, 0xf0, 0x37, 0x00, 0x00, 0xff, 0xff, 0xf1, 0xcd, 0x95, 0x94, 0xe7, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	CreateUser(ctx context.Context, in *request.CreateUserReq, opts ...grpc.CallOption) (*response.CreateUserResp, error)
	ListUsers(ctx context.Context, in *request.ListUsersReq, opts ...grpc.CallOption) (*response.ListUsersResp, error)
	GetUser(ctx context.Context, in *request.GetUserReq, opts ...grpc.CallOption) (*response.GetUserResp, error)
	DeleteUser(ctx context.Context, in *request.DeleteUserReq, opts ...grpc.CallOption) (*response.DeleteUserResp, error)
	UpdateUser(ctx context.Context, in *request.UpdateUserReq, opts ...grpc.CallOption) (*response.UpdateUserResp, error)
	UpdateSelf(ctx context.Context, in *request.UpdateSelfReq, opts ...grpc.CallOption) (*response.UpdateSelfResp, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) CreateUser(ctx context.Context, in *request.CreateUserReq, opts ...grpc.CallOption) (*response.CreateUserResp, error) {
	out := new(response.CreateUserResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Users/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ListUsers(ctx context.Context, in *request.ListUsersReq, opts ...grpc.CallOption) (*response.ListUsersResp, error) {
	out := new(response.ListUsersResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Users/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUser(ctx context.Context, in *request.GetUserReq, opts ...grpc.CallOption) (*response.GetUserResp, error) {
	out := new(response.GetUserResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Users/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) DeleteUser(ctx context.Context, in *request.DeleteUserReq, opts ...grpc.CallOption) (*response.DeleteUserResp, error) {
	out := new(response.DeleteUserResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Users/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdateUser(ctx context.Context, in *request.UpdateUserReq, opts ...grpc.CallOption) (*response.UpdateUserResp, error) {
	out := new(response.UpdateUserResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Users/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdateSelf(ctx context.Context, in *request.UpdateSelfReq, opts ...grpc.CallOption) (*response.UpdateSelfResp, error) {
	out := new(response.UpdateSelfResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Users/UpdateSelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	CreateUser(context.Context, *request.CreateUserReq) (*response.CreateUserResp, error)
	ListUsers(context.Context, *request.ListUsersReq) (*response.ListUsersResp, error)
	GetUser(context.Context, *request.GetUserReq) (*response.GetUserResp, error)
	DeleteUser(context.Context, *request.DeleteUserReq) (*response.DeleteUserResp, error)
	UpdateUser(context.Context, *request.UpdateUserReq) (*response.UpdateUserResp, error)
	UpdateSelf(context.Context, *request.UpdateSelfReq) (*response.UpdateSelfResp, error)
}

// UnimplementedUsersServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (*UnimplementedUsersServer) CreateUser(ctx context.Context, req *request.CreateUserReq) (*response.CreateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUsersServer) ListUsers(ctx context.Context, req *request.ListUsersReq) (*response.ListUsersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (*UnimplementedUsersServer) GetUser(ctx context.Context, req *request.GetUserReq) (*response.GetUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUsersServer) DeleteUser(ctx context.Context, req *request.DeleteUserReq) (*response.DeleteUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (*UnimplementedUsersServer) UpdateUser(ctx context.Context, req *request.UpdateUserReq) (*response.UpdateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUsersServer) UpdateSelf(ctx context.Context, req *request.UpdateSelfReq) (*response.UpdateSelfResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSelf not implemented")
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Users/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CreateUser(ctx, req.(*request.CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Users/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ListUsers(ctx, req.(*request.ListUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Users/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUser(ctx, req.(*request.GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Users/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).DeleteUser(ctx, req.(*request.DeleteUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Users/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdateUser(ctx, req.(*request.UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdateSelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateSelfReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdateSelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Users/UpdateSelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdateSelf(ctx, req.(*request.UpdateSelfReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.iam.v2.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Users_CreateUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Users_ListUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Users_GetUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Users_DeleteUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Users_UpdateUser_Handler,
		},
		{
			MethodName: "UpdateSelf",
			Handler:    _Users_UpdateSelf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/automate-gateway/api/iam/v2/users.proto",
}
