// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: api/abf.proto

package server

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ABruteforceClient is the client API for ABruteforce service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ABruteforceClient interface {
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AddToWhiteList(ctx context.Context, in *AddNetMaskRequest, opts ...grpc.CallOption) (*AddNetMaskResponse, error)
	RemoveFromWhiteList(ctx context.Context, in *RemoveNetMaskRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AddToBlackList(ctx context.Context, in *AddNetMaskRequest, opts ...grpc.CallOption) (*AddNetMaskResponse, error)
	RemoveFromBlackList(ctx context.Context, in *RemoveNetMaskRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type aBruteforceClient struct {
	cc grpc.ClientConnInterface
}

func NewABruteforceClient(cc grpc.ClientConnInterface) ABruteforceClient {
	return &aBruteforceClient{cc}
}

func (c *aBruteforceClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/abforce.ABruteforce/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBruteforceClient) Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/abforce.ABruteforce/Reset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBruteforceClient) AddToWhiteList(ctx context.Context, in *AddNetMaskRequest, opts ...grpc.CallOption) (*AddNetMaskResponse, error) {
	out := new(AddNetMaskResponse)
	err := c.cc.Invoke(ctx, "/abforce.ABruteforce/AddToWhiteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBruteforceClient) RemoveFromWhiteList(ctx context.Context, in *RemoveNetMaskRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/abforce.ABruteforce/RemoveFromWhiteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBruteforceClient) AddToBlackList(ctx context.Context, in *AddNetMaskRequest, opts ...grpc.CallOption) (*AddNetMaskResponse, error) {
	out := new(AddNetMaskResponse)
	err := c.cc.Invoke(ctx, "/abforce.ABruteforce/AddToBlackList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBruteforceClient) RemoveFromBlackList(ctx context.Context, in *RemoveNetMaskRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/abforce.ABruteforce/RemoveFromBlackList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ABruteforceServer is the server API for ABruteforce service.
// All implementations must embed UnimplementedABruteforceServer
// for forward compatibility
type ABruteforceServer interface {
	Auth(context.Context, *AuthRequest) (*AuthResponse, error)
	Reset(context.Context, *ResetRequest) (*empty.Empty, error)
	AddToWhiteList(context.Context, *AddNetMaskRequest) (*AddNetMaskResponse, error)
	RemoveFromWhiteList(context.Context, *RemoveNetMaskRequest) (*empty.Empty, error)
	AddToBlackList(context.Context, *AddNetMaskRequest) (*AddNetMaskResponse, error)
	RemoveFromBlackList(context.Context, *RemoveNetMaskRequest) (*empty.Empty, error)
	mustEmbedUnimplementedABruteforceServer()
}

// UnimplementedABruteforceServer must be embedded to have forward compatible implementations.
type UnimplementedABruteforceServer struct {
}

func (UnimplementedABruteforceServer) Auth(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedABruteforceServer) Reset(context.Context, *ResetRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reset not implemented")
}
func (UnimplementedABruteforceServer) AddToWhiteList(context.Context, *AddNetMaskRequest) (*AddNetMaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToWhiteList not implemented")
}
func (UnimplementedABruteforceServer) RemoveFromWhiteList(context.Context, *RemoveNetMaskRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromWhiteList not implemented")
}
func (UnimplementedABruteforceServer) AddToBlackList(context.Context, *AddNetMaskRequest) (*AddNetMaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToBlackList not implemented")
}
func (UnimplementedABruteforceServer) RemoveFromBlackList(context.Context, *RemoveNetMaskRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromBlackList not implemented")
}
func (UnimplementedABruteforceServer) mustEmbedUnimplementedABruteforceServer() {}

// UnsafeABruteforceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ABruteforceServer will
// result in compilation errors.
type UnsafeABruteforceServer interface {
	mustEmbedUnimplementedABruteforceServer()
}

func RegisterABruteforceServer(s grpc.ServiceRegistrar, srv ABruteforceServer) {
	s.RegisterService(&ABruteforce_ServiceDesc, srv)
}

func _ABruteforce_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABruteforceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abforce.ABruteforce/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABruteforceServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABruteforce_Reset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABruteforceServer).Reset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abforce.ABruteforce/Reset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABruteforceServer).Reset(ctx, req.(*ResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABruteforce_AddToWhiteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNetMaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABruteforceServer).AddToWhiteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abforce.ABruteforce/AddToWhiteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABruteforceServer).AddToWhiteList(ctx, req.(*AddNetMaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABruteforce_RemoveFromWhiteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveNetMaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABruteforceServer).RemoveFromWhiteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abforce.ABruteforce/RemoveFromWhiteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABruteforceServer).RemoveFromWhiteList(ctx, req.(*RemoveNetMaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABruteforce_AddToBlackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNetMaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABruteforceServer).AddToBlackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abforce.ABruteforce/AddToBlackList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABruteforceServer).AddToBlackList(ctx, req.(*AddNetMaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABruteforce_RemoveFromBlackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveNetMaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABruteforceServer).RemoveFromBlackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/abforce.ABruteforce/RemoveFromBlackList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABruteforceServer).RemoveFromBlackList(ctx, req.(*RemoveNetMaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ABruteforce_ServiceDesc is the grpc.ServiceDesc for ABruteforce service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ABruteforce_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "abforce.ABruteforce",
	HandlerType: (*ABruteforceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _ABruteforce_Auth_Handler,
		},
		{
			MethodName: "Reset",
			Handler:    _ABruteforce_Reset_Handler,
		},
		{
			MethodName: "AddToWhiteList",
			Handler:    _ABruteforce_AddToWhiteList_Handler,
		},
		{
			MethodName: "RemoveFromWhiteList",
			Handler:    _ABruteforce_RemoveFromWhiteList_Handler,
		},
		{
			MethodName: "AddToBlackList",
			Handler:    _ABruteforce_AddToBlackList_Handler,
		},
		{
			MethodName: "RemoveFromBlackList",
			Handler:    _ABruteforce_RemoveFromBlackList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/abf.proto",
}
