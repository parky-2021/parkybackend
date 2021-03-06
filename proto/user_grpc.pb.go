// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package parkingpb

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

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationClient interface {
	AuthenticateUser(ctx context.Context, in *AuthenticateUserRequest, opts ...grpc.CallOption) (*AuthenticateUserResponse, error)
}

type authenticationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationClient(cc grpc.ClientConnInterface) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) AuthenticateUser(ctx context.Context, in *AuthenticateUserRequest, opts ...grpc.CallOption) (*AuthenticateUserResponse, error) {
	out := new(AuthenticateUserResponse)
	err := c.cc.Invoke(ctx, "/parkingpb.Authentication/AuthenticateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
// All implementations must embed UnimplementedAuthenticationServer
// for forward compatibility
type AuthenticationServer interface {
	AuthenticateUser(context.Context, *AuthenticateUserRequest) (*AuthenticateUserResponse, error)
	mustEmbedUnimplementedAuthenticationServer()
}

// UnimplementedAuthenticationServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServer struct {
}

func (UnimplementedAuthenticationServer) AuthenticateUser(context.Context, *AuthenticateUserRequest) (*AuthenticateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateUser not implemented")
}
func (UnimplementedAuthenticationServer) mustEmbedUnimplementedAuthenticationServer() {}

// UnsafeAuthenticationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServer will
// result in compilation errors.
type UnsafeAuthenticationServer interface {
	mustEmbedUnimplementedAuthenticationServer()
}

func RegisterAuthenticationServer(s grpc.ServiceRegistrar, srv AuthenticationServer) {
	s.RegisterService(&Authentication_ServiceDesc, srv)
}

func _Authentication_AuthenticateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).AuthenticateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parkingpb.Authentication/AuthenticateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).AuthenticateUser(ctx, req.(*AuthenticateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Authentication_ServiceDesc is the grpc.ServiceDesc for Authentication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authentication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "parkingpb.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthenticateUser",
			Handler:    _Authentication_AuthenticateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}
