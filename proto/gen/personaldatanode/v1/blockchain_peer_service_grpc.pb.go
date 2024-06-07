// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: personaldatanode/v1/blockchain_peer_service.proto

package v1

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

// BlockchainPeerServiceClient is the client API for BlockchainPeerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockchainPeerServiceClient interface {
}

type blockchainPeerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockchainPeerServiceClient(cc grpc.ClientConnInterface) BlockchainPeerServiceClient {
	return &blockchainPeerServiceClient{cc}
}

// BlockchainPeerServiceServer is the server API for BlockchainPeerService service.
// All implementations must embed UnimplementedBlockchainPeerServiceServer
// for forward compatibility
type BlockchainPeerServiceServer interface {
	mustEmbedUnimplementedBlockchainPeerServiceServer()
}

// UnimplementedBlockchainPeerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlockchainPeerServiceServer struct {
}

func (UnimplementedBlockchainPeerServiceServer) mustEmbedUnimplementedBlockchainPeerServiceServer() {}

// UnsafeBlockchainPeerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockchainPeerServiceServer will
// result in compilation errors.
type UnsafeBlockchainPeerServiceServer interface {
	mustEmbedUnimplementedBlockchainPeerServiceServer()
}

func RegisterBlockchainPeerServiceServer(s grpc.ServiceRegistrar, srv BlockchainPeerServiceServer) {
	s.RegisterService(&BlockchainPeerService_ServiceDesc, srv)
}

// BlockchainPeerService_ServiceDesc is the grpc.ServiceDesc for BlockchainPeerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockchainPeerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "personaldatanode.v1.BlockchainPeerService",
	HandlerType: (*BlockchainPeerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "personaldatanode/v1/blockchain_peer_service.proto",
}
