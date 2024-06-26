// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: networknode/v1/api_gateway_service.proto

package v1

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

// ApiGatewayServiceClient is the client API for ApiGatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiGatewayServiceClient interface {
}

type apiGatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiGatewayServiceClient(cc grpc.ClientConnInterface) ApiGatewayServiceClient {
	return &apiGatewayServiceClient{cc}
}

// ApiGatewayServiceServer is the server API for ApiGatewayService service.
// All implementations must embed UnimplementedApiGatewayServiceServer
// for forward compatibility
type ApiGatewayServiceServer interface {
	mustEmbedUnimplementedApiGatewayServiceServer()
}

// UnimplementedApiGatewayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiGatewayServiceServer struct {
}

func (UnimplementedApiGatewayServiceServer) mustEmbedUnimplementedApiGatewayServiceServer() {}

// UnsafeApiGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiGatewayServiceServer will
// result in compilation errors.
type UnsafeApiGatewayServiceServer interface {
	mustEmbedUnimplementedApiGatewayServiceServer()
}

func RegisterApiGatewayServiceServer(s grpc.ServiceRegistrar, srv ApiGatewayServiceServer) {
	s.RegisterService(&ApiGatewayService_ServiceDesc, srv)
}

// ApiGatewayService_ServiceDesc is the grpc.ServiceDesc for ApiGatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiGatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "networknode.v1.ApiGatewayService",
	HandlerType: (*ApiGatewayServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "networknode/v1/api_gateway_service.proto",
}
