// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: translate.proto

package translate

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

// TranslateClient is the client API for Translate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TranslateClient interface {
	Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateReply, error)
}

type translateClient struct {
	cc grpc.ClientConnInterface
}

func NewTranslateClient(cc grpc.ClientConnInterface) TranslateClient {
	return &translateClient{cc}
}

func (c *translateClient) Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateReply, error) {
	out := new(TranslateReply)
	err := c.cc.Invoke(ctx, "/translate.Translate/Translate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranslateServer is the server API for Translate service.
// All implementations must embed UnimplementedTranslateServer
// for forward compatibility
type TranslateServer interface {
	Translate(context.Context, *TranslateRequest) (*TranslateReply, error)
	mustEmbedUnimplementedTranslateServer()
}

// UnimplementedTranslateServer must be embedded to have forward compatible implementations.
type UnimplementedTranslateServer struct {
}

func (UnimplementedTranslateServer) Translate(context.Context, *TranslateRequest) (*TranslateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Translate not implemented")
}
func (UnimplementedTranslateServer) mustEmbedUnimplementedTranslateServer() {}

// UnsafeTranslateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TranslateServer will
// result in compilation errors.
type UnsafeTranslateServer interface {
	mustEmbedUnimplementedTranslateServer()
}

func RegisterTranslateServer(s grpc.ServiceRegistrar, srv TranslateServer) {
	s.RegisterService(&Translate_ServiceDesc, srv)
}

func _Translate_Translate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslateServer).Translate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/translate.Translate/Translate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslateServer).Translate(ctx, req.(*TranslateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Translate_ServiceDesc is the grpc.ServiceDesc for Translate service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Translate_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "translate.Translate",
	HandlerType: (*TranslateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Translate",
			Handler:    _Translate_Translate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "translate.proto",
}
