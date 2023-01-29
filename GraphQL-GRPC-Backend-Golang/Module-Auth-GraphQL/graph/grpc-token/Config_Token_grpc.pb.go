// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: Config_Token.proto

package grpc_token

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

// RouteGuideClient is the client API for RouteGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteGuideClient interface {
	CreateToken(ctx context.Context, in *DataTokenRequest, opts ...grpc.CallOption) (*Data, error)
	ValidateToken(ctx context.Context, in *DataTokenRequest, opts ...grpc.CallOption) (*DataTokenReply, error)
}

type routeGuideClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteGuideClient(cc grpc.ClientConnInterface) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) CreateToken(ctx context.Context, in *DataTokenRequest, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/grpc_token.RouteGuide/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) ValidateToken(ctx context.Context, in *DataTokenRequest, opts ...grpc.CallOption) (*DataTokenReply, error) {
	out := new(DataTokenReply)
	err := c.cc.Invoke(ctx, "/grpc_token.RouteGuide/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteGuideServer is the server API for RouteGuide service.
// All implementations must embed UnimplementedRouteGuideServer
// for forward compatibility
type RouteGuideServer interface {
	CreateToken(context.Context, *DataTokenRequest) (*Data, error)
	ValidateToken(context.Context, *DataTokenRequest) (*DataTokenReply, error)
	mustEmbedUnimplementedRouteGuideServer()
}

// UnimplementedRouteGuideServer must be embedded to have forward compatible implementations.
type UnimplementedRouteGuideServer struct {
}

func (UnimplementedRouteGuideServer) CreateToken(context.Context, *DataTokenRequest) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (UnimplementedRouteGuideServer) ValidateToken(context.Context, *DataTokenRequest) (*DataTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedRouteGuideServer) mustEmbedUnimplementedRouteGuideServer() {}

// UnsafeRouteGuideServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteGuideServer will
// result in compilation errors.
type UnsafeRouteGuideServer interface {
	mustEmbedUnimplementedRouteGuideServer()
}

func RegisterRouteGuideServer(s grpc.ServiceRegistrar, srv RouteGuideServer) {
	s.RegisterService(&RouteGuide_ServiceDesc, srv)
}

func _RouteGuide_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_token.RouteGuide/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).CreateToken(ctx, req.(*DataTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_token.RouteGuide/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).ValidateToken(ctx, req.(*DataTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RouteGuide_ServiceDesc is the grpc.ServiceDesc for RouteGuide service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteGuide_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_token.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _RouteGuide_CreateToken_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _RouteGuide_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Config_Token.proto",
}