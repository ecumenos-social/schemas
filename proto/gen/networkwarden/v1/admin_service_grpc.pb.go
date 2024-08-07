// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: networkwarden/v1/admin_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AdminService_LoginAdmin_FullMethodName                = "/networkwarden.v1.AdminService/LoginAdmin"
	AdminService_RefreshAdminToken_FullMethodName         = "/networkwarden.v1.AdminService/RefreshAdminToken"
	AdminService_LogoutAdmin_FullMethodName               = "/networkwarden.v1.AdminService/LogoutAdmin"
	AdminService_ChangeAdminPassword_FullMethodName       = "/networkwarden.v1.AdminService/ChangeAdminPassword"
	AdminService_GetPersonalDataNodesList_FullMethodName  = "/networkwarden.v1.AdminService/GetPersonalDataNodesList"
	AdminService_GetPersonalDataNodeByID_FullMethodName   = "/networkwarden.v1.AdminService/GetPersonalDataNodeByID"
	AdminService_SetPersonalDataNodeStatus_FullMethodName = "/networkwarden.v1.AdminService/SetPersonalDataNodeStatus"
	AdminService_GetNetworkNodesList_FullMethodName       = "/networkwarden.v1.AdminService/GetNetworkNodesList"
	AdminService_GetNetworkNodeByID_FullMethodName        = "/networkwarden.v1.AdminService/GetNetworkNodeByID"
	AdminService_SetNetworkNodeStatus_FullMethodName      = "/networkwarden.v1.AdminService/SetNetworkNodeStatus"
	AdminService_GetNetworkWardensList_FullMethodName     = "/networkwarden.v1.AdminService/GetNetworkWardensList"
	AdminService_GetNetworkWardenByID_FullMethodName      = "/networkwarden.v1.AdminService/GetNetworkWardenByID"
)

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	LoginAdmin(ctx context.Context, in *AdminServiceLoginAdminRequest, opts ...grpc.CallOption) (*AdminServiceLoginAdminResponse, error)
	RefreshAdminToken(ctx context.Context, in *AdminServiceRefreshAdminTokenRequest, opts ...grpc.CallOption) (*AdminServiceRefreshAdminTokenResponse, error)
	LogoutAdmin(ctx context.Context, in *AdminServiceLogoutAdminRequest, opts ...grpc.CallOption) (*AdminServiceLogoutAdminResponse, error)
	ChangeAdminPassword(ctx context.Context, in *AdminServiceChangeAdminPasswordRequest, opts ...grpc.CallOption) (*AdminServiceChangeAdminPasswordResponse, error)
	GetPersonalDataNodesList(ctx context.Context, in *AdminServiceGetPersonalDataNodesListRequest, opts ...grpc.CallOption) (*AdminServiceGetPersonalDataNodesListResponse, error)
	GetPersonalDataNodeByID(ctx context.Context, in *AdminServiceGetPersonalDataNodeByIDRequest, opts ...grpc.CallOption) (*AdminServiceGetPersonalDataNodeByIDResponse, error)
	SetPersonalDataNodeStatus(ctx context.Context, in *AdminServiceSetPersonalDataNodeStatusRequest, opts ...grpc.CallOption) (*AdminServiceSetPersonalDataNodeStatusResponse, error)
	GetNetworkNodesList(ctx context.Context, in *AdminServiceGetNetworkNodesListRequest, opts ...grpc.CallOption) (*AdminServiceGetNetworkNodesListResponse, error)
	GetNetworkNodeByID(ctx context.Context, in *AdminServiceGetNetworkNodeByIDRequest, opts ...grpc.CallOption) (*AdminServiceGetNetworkNodeByIDResponse, error)
	SetNetworkNodeStatus(ctx context.Context, in *AdminServiceSetNetworkNodeStatusRequest, opts ...grpc.CallOption) (*AdminServiceSetNetworkNodeStatusResponse, error)
	GetNetworkWardensList(ctx context.Context, in *AdminServiceGetNetworkWardensListRequest, opts ...grpc.CallOption) (*AdminServiceGetNetworkWardensListResponse, error)
	GetNetworkWardenByID(ctx context.Context, in *AdminServiceGetNetworkWardenByIDRequest, opts ...grpc.CallOption) (*AdminServiceGetNetworkWardenByIDResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) LoginAdmin(ctx context.Context, in *AdminServiceLoginAdminRequest, opts ...grpc.CallOption) (*AdminServiceLoginAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminServiceLoginAdminResponse)
	err := c.cc.Invoke(ctx, AdminService_LoginAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) RefreshAdminToken(ctx context.Context, in *AdminServiceRefreshAdminTokenRequest, opts ...grpc.CallOption) (*AdminServiceRefreshAdminTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminServiceRefreshAdminTokenResponse)
	err := c.cc.Invoke(ctx, AdminService_RefreshAdminToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) LogoutAdmin(ctx context.Context, in *AdminServiceLogoutAdminRequest, opts ...grpc.CallOption) (*AdminServiceLogoutAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminServiceLogoutAdminResponse)
	err := c.cc.Invoke(ctx, AdminService_LogoutAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ChangeAdminPassword(ctx context.Context, in *AdminServiceChangeAdminPasswordRequest, opts ...grpc.CallOption) (*AdminServiceChangeAdminPasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminServiceChangeAdminPasswordResponse)
	err := c.cc.Invoke(ctx, AdminService_ChangeAdminPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetPersonalDataNodesList(ctx context.Context, in *AdminServiceGetPersonalDataNodesListRequest, opts ...grpc.CallOption) (*AdminServiceGetPersonalDataNodesListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminServiceGetPersonalDataNodesListResponse)
	err := c.cc.Invoke(ctx, AdminService_GetPersonalDataNodesList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetPersonalDataNodeByID(ctx context.Context, in *AdminServiceGetPersonalDataNodeByIDRequest, opts ...grpc.CallOption) (*AdminServiceGetPersonalDataNodeByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminServiceGetPersonalDataNodeByIDResponse)
	err := c.cc.Invoke(ctx, AdminService_GetPersonalDataNodeByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) SetPersonalDataNodeStatus(ctx context.Context, in *AdminServiceSetPersonalDataNodeStatusRequest, opts ...grpc.CallOption) (*AdminServiceSetPersonalDataNodeStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminServiceSetPersonalDataNodeStatusResponse)
	err := c.cc.Invoke(ctx, AdminService_SetPersonalDataNodeStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetNetworkNodesList(ctx context.Context, in *AdminServiceGetNetworkNodesListRequest, opts ...grpc.CallOption) (*AdminServiceGetNetworkNodesListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminServiceGetNetworkNodesListResponse)
	err := c.cc.Invoke(ctx, AdminService_GetNetworkNodesList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetNetworkNodeByID(ctx context.Context, in *AdminServiceGetNetworkNodeByIDRequest, opts ...grpc.CallOption) (*AdminServiceGetNetworkNodeByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminServiceGetNetworkNodeByIDResponse)
	err := c.cc.Invoke(ctx, AdminService_GetNetworkNodeByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) SetNetworkNodeStatus(ctx context.Context, in *AdminServiceSetNetworkNodeStatusRequest, opts ...grpc.CallOption) (*AdminServiceSetNetworkNodeStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminServiceSetNetworkNodeStatusResponse)
	err := c.cc.Invoke(ctx, AdminService_SetNetworkNodeStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetNetworkWardensList(ctx context.Context, in *AdminServiceGetNetworkWardensListRequest, opts ...grpc.CallOption) (*AdminServiceGetNetworkWardensListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminServiceGetNetworkWardensListResponse)
	err := c.cc.Invoke(ctx, AdminService_GetNetworkWardensList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetNetworkWardenByID(ctx context.Context, in *AdminServiceGetNetworkWardenByIDRequest, opts ...grpc.CallOption) (*AdminServiceGetNetworkWardenByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminServiceGetNetworkWardenByIDResponse)
	err := c.cc.Invoke(ctx, AdminService_GetNetworkWardenByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility.
type AdminServiceServer interface {
	LoginAdmin(context.Context, *AdminServiceLoginAdminRequest) (*AdminServiceLoginAdminResponse, error)
	RefreshAdminToken(context.Context, *AdminServiceRefreshAdminTokenRequest) (*AdminServiceRefreshAdminTokenResponse, error)
	LogoutAdmin(context.Context, *AdminServiceLogoutAdminRequest) (*AdminServiceLogoutAdminResponse, error)
	ChangeAdminPassword(context.Context, *AdminServiceChangeAdminPasswordRequest) (*AdminServiceChangeAdminPasswordResponse, error)
	GetPersonalDataNodesList(context.Context, *AdminServiceGetPersonalDataNodesListRequest) (*AdminServiceGetPersonalDataNodesListResponse, error)
	GetPersonalDataNodeByID(context.Context, *AdminServiceGetPersonalDataNodeByIDRequest) (*AdminServiceGetPersonalDataNodeByIDResponse, error)
	SetPersonalDataNodeStatus(context.Context, *AdminServiceSetPersonalDataNodeStatusRequest) (*AdminServiceSetPersonalDataNodeStatusResponse, error)
	GetNetworkNodesList(context.Context, *AdminServiceGetNetworkNodesListRequest) (*AdminServiceGetNetworkNodesListResponse, error)
	GetNetworkNodeByID(context.Context, *AdminServiceGetNetworkNodeByIDRequest) (*AdminServiceGetNetworkNodeByIDResponse, error)
	SetNetworkNodeStatus(context.Context, *AdminServiceSetNetworkNodeStatusRequest) (*AdminServiceSetNetworkNodeStatusResponse, error)
	GetNetworkWardensList(context.Context, *AdminServiceGetNetworkWardensListRequest) (*AdminServiceGetNetworkWardensListResponse, error)
	GetNetworkWardenByID(context.Context, *AdminServiceGetNetworkWardenByIDRequest) (*AdminServiceGetNetworkWardenByIDResponse, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdminServiceServer struct{}

func (UnimplementedAdminServiceServer) LoginAdmin(context.Context, *AdminServiceLoginAdminRequest) (*AdminServiceLoginAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginAdmin not implemented")
}
func (UnimplementedAdminServiceServer) RefreshAdminToken(context.Context, *AdminServiceRefreshAdminTokenRequest) (*AdminServiceRefreshAdminTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshAdminToken not implemented")
}
func (UnimplementedAdminServiceServer) LogoutAdmin(context.Context, *AdminServiceLogoutAdminRequest) (*AdminServiceLogoutAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogoutAdmin not implemented")
}
func (UnimplementedAdminServiceServer) ChangeAdminPassword(context.Context, *AdminServiceChangeAdminPasswordRequest) (*AdminServiceChangeAdminPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAdminPassword not implemented")
}
func (UnimplementedAdminServiceServer) GetPersonalDataNodesList(context.Context, *AdminServiceGetPersonalDataNodesListRequest) (*AdminServiceGetPersonalDataNodesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonalDataNodesList not implemented")
}
func (UnimplementedAdminServiceServer) GetPersonalDataNodeByID(context.Context, *AdminServiceGetPersonalDataNodeByIDRequest) (*AdminServiceGetPersonalDataNodeByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonalDataNodeByID not implemented")
}
func (UnimplementedAdminServiceServer) SetPersonalDataNodeStatus(context.Context, *AdminServiceSetPersonalDataNodeStatusRequest) (*AdminServiceSetPersonalDataNodeStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPersonalDataNodeStatus not implemented")
}
func (UnimplementedAdminServiceServer) GetNetworkNodesList(context.Context, *AdminServiceGetNetworkNodesListRequest) (*AdminServiceGetNetworkNodesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkNodesList not implemented")
}
func (UnimplementedAdminServiceServer) GetNetworkNodeByID(context.Context, *AdminServiceGetNetworkNodeByIDRequest) (*AdminServiceGetNetworkNodeByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkNodeByID not implemented")
}
func (UnimplementedAdminServiceServer) SetNetworkNodeStatus(context.Context, *AdminServiceSetNetworkNodeStatusRequest) (*AdminServiceSetNetworkNodeStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNetworkNodeStatus not implemented")
}
func (UnimplementedAdminServiceServer) GetNetworkWardensList(context.Context, *AdminServiceGetNetworkWardensListRequest) (*AdminServiceGetNetworkWardensListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkWardensList not implemented")
}
func (UnimplementedAdminServiceServer) GetNetworkWardenByID(context.Context, *AdminServiceGetNetworkWardenByIDRequest) (*AdminServiceGetNetworkWardenByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkWardenByID not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}
func (UnimplementedAdminServiceServer) testEmbeddedByValue()                      {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	// If the following call pancis, it indicates UnimplementedAdminServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_LoginAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminServiceLoginAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).LoginAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_LoginAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).LoginAdmin(ctx, req.(*AdminServiceLoginAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_RefreshAdminToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminServiceRefreshAdminTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).RefreshAdminToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_RefreshAdminToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).RefreshAdminToken(ctx, req.(*AdminServiceRefreshAdminTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_LogoutAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminServiceLogoutAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).LogoutAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_LogoutAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).LogoutAdmin(ctx, req.(*AdminServiceLogoutAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ChangeAdminPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminServiceChangeAdminPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ChangeAdminPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_ChangeAdminPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ChangeAdminPassword(ctx, req.(*AdminServiceChangeAdminPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetPersonalDataNodesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminServiceGetPersonalDataNodesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetPersonalDataNodesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_GetPersonalDataNodesList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetPersonalDataNodesList(ctx, req.(*AdminServiceGetPersonalDataNodesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetPersonalDataNodeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminServiceGetPersonalDataNodeByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetPersonalDataNodeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_GetPersonalDataNodeByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetPersonalDataNodeByID(ctx, req.(*AdminServiceGetPersonalDataNodeByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_SetPersonalDataNodeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminServiceSetPersonalDataNodeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).SetPersonalDataNodeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_SetPersonalDataNodeStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).SetPersonalDataNodeStatus(ctx, req.(*AdminServiceSetPersonalDataNodeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetNetworkNodesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminServiceGetNetworkNodesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetNetworkNodesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_GetNetworkNodesList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetNetworkNodesList(ctx, req.(*AdminServiceGetNetworkNodesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetNetworkNodeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminServiceGetNetworkNodeByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetNetworkNodeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_GetNetworkNodeByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetNetworkNodeByID(ctx, req.(*AdminServiceGetNetworkNodeByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_SetNetworkNodeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminServiceSetNetworkNodeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).SetNetworkNodeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_SetNetworkNodeStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).SetNetworkNodeStatus(ctx, req.(*AdminServiceSetNetworkNodeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetNetworkWardensList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminServiceGetNetworkWardensListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetNetworkWardensList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_GetNetworkWardensList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetNetworkWardensList(ctx, req.(*AdminServiceGetNetworkWardensListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetNetworkWardenByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminServiceGetNetworkWardenByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetNetworkWardenByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_GetNetworkWardenByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetNetworkWardenByID(ctx, req.(*AdminServiceGetNetworkWardenByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "networkwarden.v1.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginAdmin",
			Handler:    _AdminService_LoginAdmin_Handler,
		},
		{
			MethodName: "RefreshAdminToken",
			Handler:    _AdminService_RefreshAdminToken_Handler,
		},
		{
			MethodName: "LogoutAdmin",
			Handler:    _AdminService_LogoutAdmin_Handler,
		},
		{
			MethodName: "ChangeAdminPassword",
			Handler:    _AdminService_ChangeAdminPassword_Handler,
		},
		{
			MethodName: "GetPersonalDataNodesList",
			Handler:    _AdminService_GetPersonalDataNodesList_Handler,
		},
		{
			MethodName: "GetPersonalDataNodeByID",
			Handler:    _AdminService_GetPersonalDataNodeByID_Handler,
		},
		{
			MethodName: "SetPersonalDataNodeStatus",
			Handler:    _AdminService_SetPersonalDataNodeStatus_Handler,
		},
		{
			MethodName: "GetNetworkNodesList",
			Handler:    _AdminService_GetNetworkNodesList_Handler,
		},
		{
			MethodName: "GetNetworkNodeByID",
			Handler:    _AdminService_GetNetworkNodeByID_Handler,
		},
		{
			MethodName: "SetNetworkNodeStatus",
			Handler:    _AdminService_SetNetworkNodeStatus_Handler,
		},
		{
			MethodName: "GetNetworkWardensList",
			Handler:    _AdminService_GetNetworkWardensList_Handler,
		},
		{
			MethodName: "GetNetworkWardenByID",
			Handler:    _AdminService_GetNetworkWardenByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "networkwarden/v1/admin_service.proto",
}
