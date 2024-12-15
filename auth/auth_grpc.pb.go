// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: auth/auth.proto

package auth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Auth_UserToken_FullMethodName       = "/openim.auth.Auth/userToken"
	Auth_GetUserToken_FullMethodName    = "/openim.auth.Auth/getUserToken"
	Auth_ForceLogout_FullMethodName     = "/openim.auth.Auth/forceLogout"
	Auth_ParseToken_FullMethodName      = "/openim.auth.Auth/parseToken"
	Auth_InvalidateToken_FullMethodName = "/openim.auth.Auth/invalidateToken"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	// Generate token
	UserToken(ctx context.Context, in *UserTokenReq, opts ...grpc.CallOption) (*UserTokenResp, error)
	// Admin retrieves user token
	GetUserToken(ctx context.Context, in *GetUserTokenReq, opts ...grpc.CallOption) (*GetUserTokenResp, error)
	// Force logout
	ForceLogout(ctx context.Context, in *ForceLogoutReq, opts ...grpc.CallOption) (*ForceLogoutResp, error)
	// Parse token
	ParseToken(ctx context.Context, in *ParseTokenReq, opts ...grpc.CallOption) (*ParseTokenResp, error)
	// Invalidate or mark the token as kicked out
	InvalidateToken(ctx context.Context, in *InvalidateTokenReq, opts ...grpc.CallOption) (*InvalidateTokenResp, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) UserToken(ctx context.Context, in *UserTokenReq, opts ...grpc.CallOption) (*UserTokenResp, error) {
	out := new(UserTokenResp)
	err := c.cc.Invoke(ctx, Auth_UserToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetUserToken(ctx context.Context, in *GetUserTokenReq, opts ...grpc.CallOption) (*GetUserTokenResp, error) {
	out := new(GetUserTokenResp)
	err := c.cc.Invoke(ctx, Auth_GetUserToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ForceLogout(ctx context.Context, in *ForceLogoutReq, opts ...grpc.CallOption) (*ForceLogoutResp, error) {
	out := new(ForceLogoutResp)
	err := c.cc.Invoke(ctx, Auth_ForceLogout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ParseToken(ctx context.Context, in *ParseTokenReq, opts ...grpc.CallOption) (*ParseTokenResp, error) {
	out := new(ParseTokenResp)
	err := c.cc.Invoke(ctx, Auth_ParseToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) InvalidateToken(ctx context.Context, in *InvalidateTokenReq, opts ...grpc.CallOption) (*InvalidateTokenResp, error) {
	out := new(InvalidateTokenResp)
	err := c.cc.Invoke(ctx, Auth_InvalidateToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	// Generate token
	UserToken(context.Context, *UserTokenReq) (*UserTokenResp, error)
	// Admin retrieves user token
	GetUserToken(context.Context, *GetUserTokenReq) (*GetUserTokenResp, error)
	// Force logout
	ForceLogout(context.Context, *ForceLogoutReq) (*ForceLogoutResp, error)
	// Parse token
	ParseToken(context.Context, *ParseTokenReq) (*ParseTokenResp, error)
	// Invalidate or mark the token as kicked out
	InvalidateToken(context.Context, *InvalidateTokenReq) (*InvalidateTokenResp, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) UserToken(context.Context, *UserTokenReq) (*UserTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserToken not implemented")
}
func (UnimplementedAuthServer) GetUserToken(context.Context, *GetUserTokenReq) (*GetUserTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserToken not implemented")
}
func (UnimplementedAuthServer) ForceLogout(context.Context, *ForceLogoutReq) (*ForceLogoutResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForceLogout not implemented")
}
func (UnimplementedAuthServer) ParseToken(context.Context, *ParseTokenReq) (*ParseTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseToken not implemented")
}
func (UnimplementedAuthServer) InvalidateToken(context.Context, *InvalidateTokenReq) (*InvalidateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateToken not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_UserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_UserToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserToken(ctx, req.(*UserTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetUserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetUserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetUserToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetUserToken(ctx, req.(*GetUserTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ForceLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForceLogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ForceLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ForceLogout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ForceLogout(ctx, req.(*ForceLogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ParseToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ParseToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ParseToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ParseToken(ctx, req.(*ParseTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_InvalidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvalidateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).InvalidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_InvalidateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).InvalidateToken(ctx, req.(*InvalidateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openim.auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "userToken",
			Handler:    _Auth_UserToken_Handler,
		},
		{
			MethodName: "getUserToken",
			Handler:    _Auth_GetUserToken_Handler,
		},
		{
			MethodName: "forceLogout",
			Handler:    _Auth_ForceLogout_Handler,
		},
		{
			MethodName: "parseToken",
			Handler:    _Auth_ParseToken_Handler,
		},
		{
			MethodName: "invalidateToken",
			Handler:    _Auth_InvalidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}
