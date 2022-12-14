// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.11
// source: user.proto

package __

import (
	context "context"
	"github.com/DudzinDzmitry/CourseProj/AuthService/internal/server"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CRUDClient is the client API for CRUD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CRUDClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error)
	LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error)
	LogOut(ctx context.Context, in *LogOutRequest, opts ...grpc.CallOption) (*LogOutResponse, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountResponse, error)
}

type cRUDClient struct {
	cc grpc.ClientConnInterface
}

func NewCRUDClient(cc grpc.ClientConnInterface) CRUDClient {
	return &cRUDClient{cc}
}

func (c *cRUDClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/CRUD/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error) {
	out := new(DeleteAccountResponse)
	err := c.cc.Invoke(ctx, "/CRUD/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error) {
	out := new(LogInResponse)
	err := c.cc.Invoke(ctx, "/CRUD/LogIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) LogOut(ctx context.Context, in *LogOutRequest, opts ...grpc.CallOption) (*LogOutResponse, error) {
	out := new(LogOutResponse)
	err := c.cc.Invoke(ctx, "/CRUD/LogOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountResponse, error) {
	out := new(UpdateAccountResponse)
	err := c.cc.Invoke(ctx, "/CRUD/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CRUDServer is the server API for CRUD service.
// All implementations must embed UnimplementedCRUDServer
// for forward compatibility
type CRUDServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error)
	LogIn(context.Context, *LogInRequest) (*LogInResponse, error)
	LogOut(context.Context, *LogOutRequest) (*LogOutResponse, error)
	UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error)
	mustEmbedUnimplementedCRUDServer()
}

// UnimplementedCRUDServer must be embedded to have forward compatible implementations.
type UnimplementedCRUDServer struct {
}

func (UnimplementedCRUDServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedCRUDServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedCRUDServer) LogIn(context.Context, *LogInRequest) (*LogInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogIn not implemented")
}
func (UnimplementedCRUDServer) LogOut(context.Context, *LogOutRequest) (*LogOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogOut not implemented")
}
func (UnimplementedCRUDServer) UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedCRUDServer) mustEmbedUnimplementedCRUDServer() {}

// UnsafeCRUDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CRUDServer will
// result in compilation errors.
type UnsafeCRUDServer interface {
	mustEmbedUnimplementedCRUDServer()
}

func RegisterCRUDServer(s grpc.ServiceRegistrar, srv *server.Server) {
	s.RegisterService(&CRUD_ServiceDesc, srv)
}

func _CRUD_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CRUD/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CRUD/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CRUD/LogIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).LogIn(ctx, req.(*LogInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_LogOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).LogOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CRUD/LogOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).LogOut(ctx, req.(*LogOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CRUD/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).UpdateAccount(ctx, req.(*UpdateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CRUD_ServiceDesc is the grpc.ServiceDesc for CRUD service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CRUD_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CRUD",
	HandlerType: (*CRUDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _CRUD_CreateAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _CRUD_DeleteAccount_Handler,
		},
		{
			MethodName: "LogIn",
			Handler:    _CRUD_LogIn_Handler,
		},
		{
			MethodName: "LogOut",
			Handler:    _CRUD_LogOut_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _CRUD_UpdateAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
