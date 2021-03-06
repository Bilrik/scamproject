// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: scam/scampb/scam.proto

package scampb

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

// ScamDatabaseClient is the client API for ScamDatabase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScamDatabaseClient interface {
	GetPhoneCategory(ctx context.Context, in *GetPhoneCategoryRequest, opts ...grpc.CallOption) (*GetPhoneCategoryResponse, error)
	GetPersonalBlocklist(ctx context.Context, in *GetPersonalBlocklistRequest, opts ...grpc.CallOption) (*GetPersonalBlocklistResponse, error)
	AddPhoneNumber(ctx context.Context, in *AddPhoneNumberRequest, opts ...grpc.CallOption) (*AddPhoneNumberResponse, error)
	GetScore(ctx context.Context, in *GetScoreRequest, opts ...grpc.CallOption) (*GetScoreResponse, error)
}

type scamDatabaseClient struct {
	cc grpc.ClientConnInterface
}

func NewScamDatabaseClient(cc grpc.ClientConnInterface) ScamDatabaseClient {
	return &scamDatabaseClient{cc}
}

func (c *scamDatabaseClient) GetPhoneCategory(ctx context.Context, in *GetPhoneCategoryRequest, opts ...grpc.CallOption) (*GetPhoneCategoryResponse, error) {
	out := new(GetPhoneCategoryResponse)
	err := c.cc.Invoke(ctx, "/scam.ScamDatabase/getPhoneCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scamDatabaseClient) GetPersonalBlocklist(ctx context.Context, in *GetPersonalBlocklistRequest, opts ...grpc.CallOption) (*GetPersonalBlocklistResponse, error) {
	out := new(GetPersonalBlocklistResponse)
	err := c.cc.Invoke(ctx, "/scam.ScamDatabase/GetPersonalBlocklist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scamDatabaseClient) AddPhoneNumber(ctx context.Context, in *AddPhoneNumberRequest, opts ...grpc.CallOption) (*AddPhoneNumberResponse, error) {
	out := new(AddPhoneNumberResponse)
	err := c.cc.Invoke(ctx, "/scam.ScamDatabase/AddPhoneNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scamDatabaseClient) GetScore(ctx context.Context, in *GetScoreRequest, opts ...grpc.CallOption) (*GetScoreResponse, error) {
	out := new(GetScoreResponse)
	err := c.cc.Invoke(ctx, "/scam.ScamDatabase/GetScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScamDatabaseServer is the server API for ScamDatabase service.
// All implementations must embed UnimplementedScamDatabaseServer
// for forward compatibility
type ScamDatabaseServer interface {
	GetPhoneCategory(context.Context, *GetPhoneCategoryRequest) (*GetPhoneCategoryResponse, error)
	GetPersonalBlocklist(context.Context, *GetPersonalBlocklistRequest) (*GetPersonalBlocklistResponse, error)
	AddPhoneNumber(context.Context, *AddPhoneNumberRequest) (*AddPhoneNumberResponse, error)
	GetScore(context.Context, *GetScoreRequest) (*GetScoreResponse, error)
	mustEmbedUnimplementedScamDatabaseServer()
}

// UnimplementedScamDatabaseServer must be embedded to have forward compatible implementations.
type UnimplementedScamDatabaseServer struct {
}

func (UnimplementedScamDatabaseServer) GetPhoneCategory(context.Context, *GetPhoneCategoryRequest) (*GetPhoneCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPhoneCategory not implemented")
}
func (UnimplementedScamDatabaseServer) GetPersonalBlocklist(context.Context, *GetPersonalBlocklistRequest) (*GetPersonalBlocklistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonalBlocklist not implemented")
}
func (UnimplementedScamDatabaseServer) AddPhoneNumber(context.Context, *AddPhoneNumberRequest) (*AddPhoneNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPhoneNumber not implemented")
}
func (UnimplementedScamDatabaseServer) GetScore(context.Context, *GetScoreRequest) (*GetScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScore not implemented")
}
func (UnimplementedScamDatabaseServer) mustEmbedUnimplementedScamDatabaseServer() {}

// UnsafeScamDatabaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScamDatabaseServer will
// result in compilation errors.
type UnsafeScamDatabaseServer interface {
	mustEmbedUnimplementedScamDatabaseServer()
}

func RegisterScamDatabaseServer(s grpc.ServiceRegistrar, srv ScamDatabaseServer) {
	s.RegisterService(&ScamDatabase_ServiceDesc, srv)
}

func _ScamDatabase_GetPhoneCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPhoneCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScamDatabaseServer).GetPhoneCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scam.ScamDatabase/getPhoneCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScamDatabaseServer).GetPhoneCategory(ctx, req.(*GetPhoneCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScamDatabase_GetPersonalBlocklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonalBlocklistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScamDatabaseServer).GetPersonalBlocklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scam.ScamDatabase/GetPersonalBlocklist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScamDatabaseServer).GetPersonalBlocklist(ctx, req.(*GetPersonalBlocklistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScamDatabase_AddPhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPhoneNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScamDatabaseServer).AddPhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scam.ScamDatabase/AddPhoneNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScamDatabaseServer).AddPhoneNumber(ctx, req.(*AddPhoneNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScamDatabase_GetScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScamDatabaseServer).GetScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scam.ScamDatabase/GetScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScamDatabaseServer).GetScore(ctx, req.(*GetScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScamDatabase_ServiceDesc is the grpc.ServiceDesc for ScamDatabase service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScamDatabase_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scam.ScamDatabase",
	HandlerType: (*ScamDatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getPhoneCategory",
			Handler:    _ScamDatabase_GetPhoneCategory_Handler,
		},
		{
			MethodName: "GetPersonalBlocklist",
			Handler:    _ScamDatabase_GetPersonalBlocklist_Handler,
		},
		{
			MethodName: "AddPhoneNumber",
			Handler:    _ScamDatabase_AddPhoneNumber_Handler,
		},
		{
			MethodName: "GetScore",
			Handler:    _ScamDatabase_GetScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scam/scampb/scam.proto",
}
