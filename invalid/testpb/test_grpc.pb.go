// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: Test/testpb/test.proto

package testpb

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

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	//Unary
	Valid(ctx context.Context, in *IsValidrequest, opts ...grpc.CallOption) (*IsValidresponse, error)
	//Client stream
	ValidList(ctx context.Context, opts ...grpc.CallOption) (TestService_ValidListClient, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) Valid(ctx context.Context, in *IsValidrequest, opts ...grpc.CallOption) (*IsValidresponse, error) {
	out := new(IsValidresponse)
	err := c.cc.Invoke(ctx, "/Test.TestService/Valid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) ValidList(ctx context.Context, opts ...grpc.CallOption) (TestService_ValidListClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[0], "/Test.TestService/ValidList", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceValidListClient{stream}
	return x, nil
}

type TestService_ValidListClient interface {
	Send(*IsValidListrequest) error
	CloseAndRecv() (*IsValidListresponse, error)
	grpc.ClientStream
}

type testServiceValidListClient struct {
	grpc.ClientStream
}

func (x *testServiceValidListClient) Send(m *IsValidListrequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceValidListClient) CloseAndRecv() (*IsValidListresponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(IsValidListresponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility
type TestServiceServer interface {
	//Unary
	Valid(context.Context, *IsValidrequest) (*IsValidresponse, error)
	//Client stream
	ValidList(TestService_ValidListServer) error
	mustEmbedUnimplementedTestServiceServer()
}

// UnimplementedTestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (UnimplementedTestServiceServer) Valid(context.Context, *IsValidrequest) (*IsValidresponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Valid not implemented")
}
func (UnimplementedTestServiceServer) ValidList(TestService_ValidListServer) error {
	return status.Errorf(codes.Unimplemented, "method ValidList not implemented")
}
func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_Valid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsValidrequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Valid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Test.TestService/Valid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Valid(ctx, req.(*IsValidrequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_ValidList_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).ValidList(&testServiceValidListServer{stream})
}

type TestService_ValidListServer interface {
	SendAndClose(*IsValidListresponse) error
	Recv() (*IsValidListrequest, error)
	grpc.ServerStream
}

type testServiceValidListServer struct {
	grpc.ServerStream
}

func (x *testServiceValidListServer) SendAndClose(m *IsValidListresponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceValidListServer) Recv() (*IsValidListrequest, error) {
	m := new(IsValidListrequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Test.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Valid",
			Handler:    _TestService_Valid_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ValidList",
			Handler:       _TestService_ValidList_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "Test/testpb/test.proto",
}
