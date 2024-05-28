// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: networkwarden/v1/network_warden_service.proto

package v1

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
	NetworkWardenService_CheckEmails_FullMethodName                              = "/networkwarden.v1.NetworkWardenService/CheckEmails"
	NetworkWardenService_RegisterHolder_FullMethodName                           = "/networkwarden.v1.NetworkWardenService/RegisterHolder"
	NetworkWardenService_ConfirmHolderRegistration_FullMethodName                = "/networkwarden.v1.NetworkWardenService/ConfirmHolderRegistration"
	NetworkWardenService_ResendConfirmationCode_FullMethodName                   = "/networkwarden.v1.NetworkWardenService/ResendConfirmationCode"
	NetworkWardenService_LoginHolder_FullMethodName                              = "/networkwarden.v1.NetworkWardenService/LoginHolder"
	NetworkWardenService_LogoutHolder_FullMethodName                             = "/networkwarden.v1.NetworkWardenService/LogoutHolder"
	NetworkWardenService_RefreshHolderToken_FullMethodName                       = "/networkwarden.v1.NetworkWardenService/RefreshHolderToken"
	NetworkWardenService_ChangeHolderPassword_FullMethodName                     = "/networkwarden.v1.NetworkWardenService/ChangeHolderPassword"
	NetworkWardenService_ModifyHolder_FullMethodName                             = "/networkwarden.v1.NetworkWardenService/ModifyHolder"
	NetworkWardenService_GetHolder_FullMethodName                                = "/networkwarden.v1.NetworkWardenService/GetHolder"
	NetworkWardenService_DeleteHolder_FullMethodName                             = "/networkwarden.v1.NetworkWardenService/DeleteHolder"
	NetworkWardenService_GetPersonalDataNodesList_FullMethodName                 = "/networkwarden.v1.NetworkWardenService/GetPersonalDataNodesList"
	NetworkWardenService_JoinPersonalDataNodeRegistrationWaitlist_FullMethodName = "/networkwarden.v1.NetworkWardenService/JoinPersonalDataNodeRegistrationWaitlist"
	NetworkWardenService_RegisterPersonalDataNode_FullMethodName                 = "/networkwarden.v1.NetworkWardenService/RegisterPersonalDataNode"
	NetworkWardenService_GetNetworkNodesList_FullMethodName                      = "/networkwarden.v1.NetworkWardenService/GetNetworkNodesList"
	NetworkWardenService_JoinNetworkNodeRegistrationWaitlist_FullMethodName      = "/networkwarden.v1.NetworkWardenService/JoinNetworkNodeRegistrationWaitlist"
	NetworkWardenService_RegisterNetworkNode_FullMethodName                      = "/networkwarden.v1.NetworkWardenService/RegisterNetworkNode"
	NetworkWardenService_GetNetworkWardensList_FullMethodName                    = "/networkwarden.v1.NetworkWardenService/GetNetworkWardensList"
	NetworkWardenService_RegisterNetworkWarden_FullMethodName                    = "/networkwarden.v1.NetworkWardenService/RegisterNetworkWarden"
)

// NetworkWardenServiceClient is the client API for NetworkWardenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkWardenServiceClient interface {
	CheckEmails(ctx context.Context, in *CheckEmailsRequest, opts ...grpc.CallOption) (*CheckEmailsResponse, error)
	RegisterHolder(ctx context.Context, in *RegisterHolderRequest, opts ...grpc.CallOption) (*RegisterHolderResponse, error)
	ConfirmHolderRegistration(ctx context.Context, in *ConfirmHolderRegistrationRequest, opts ...grpc.CallOption) (*ConfirmHolderRegistrationResponse, error)
	ResendConfirmationCode(ctx context.Context, in *ResendConfirmationCodeRequest, opts ...grpc.CallOption) (*ResendConfirmationCodeResponse, error)
	LoginHolder(ctx context.Context, in *LoginHolderRequest, opts ...grpc.CallOption) (*LoginHolderResponse, error)
	LogoutHolder(ctx context.Context, in *LogoutHolderRequest, opts ...grpc.CallOption) (*LogoutHolderResponse, error)
	RefreshHolderToken(ctx context.Context, in *RefreshHolderTokenRequest, opts ...grpc.CallOption) (*RefreshHolderTokenResponse, error)
	ChangeHolderPassword(ctx context.Context, in *ChangeHolderPasswordRequest, opts ...grpc.CallOption) (*ChangeHolderPasswordResponse, error)
	ModifyHolder(ctx context.Context, in *ModifyHolderRequest, opts ...grpc.CallOption) (*ModifyHolderResponse, error)
	GetHolder(ctx context.Context, in *GetHolderRequest, opts ...grpc.CallOption) (*GetHolderResponse, error)
	DeleteHolder(ctx context.Context, in *DeleteHolderRequest, opts ...grpc.CallOption) (*DeleteHolderResponse, error)
	GetPersonalDataNodesList(ctx context.Context, in *GetPersonalDataNodesListRequest, opts ...grpc.CallOption) (*GetPersonalDataNodesListResponse, error)
	JoinPersonalDataNodeRegistrationWaitlist(ctx context.Context, in *JoinPersonalDataNodeRegistrationWaitlistRequest, opts ...grpc.CallOption) (*JoinPersonalDataNodeRegistrationWaitlistResponse, error)
	RegisterPersonalDataNode(ctx context.Context, in *RegisterPersonalDataNodeRequest, opts ...grpc.CallOption) (*RegisterPersonalDataNodeResponse, error)
	GetNetworkNodesList(ctx context.Context, in *GetNetworkNodesListRequest, opts ...grpc.CallOption) (*GetNetworkNodesListResponse, error)
	JoinNetworkNodeRegistrationWaitlist(ctx context.Context, in *JoinNetworkNodeRegistrationWaitlistRequest, opts ...grpc.CallOption) (*JoinNetworkNodeRegistrationWaitlistResponse, error)
	RegisterNetworkNode(ctx context.Context, in *RegisterNetworkNodeRequest, opts ...grpc.CallOption) (*RegisterNetworkNodeResponse, error)
	GetNetworkWardensList(ctx context.Context, in *GetNetworkWardensListRequest, opts ...grpc.CallOption) (*GetNetworkWardensListResponse, error)
	RegisterNetworkWarden(ctx context.Context, in *RegisterNetworkWardenRequest, opts ...grpc.CallOption) (*RegisterNetworkWardenResponse, error)
}

type networkWardenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkWardenServiceClient(cc grpc.ClientConnInterface) NetworkWardenServiceClient {
	return &networkWardenServiceClient{cc}
}

func (c *networkWardenServiceClient) CheckEmails(ctx context.Context, in *CheckEmailsRequest, opts ...grpc.CallOption) (*CheckEmailsResponse, error) {
	out := new(CheckEmailsResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_CheckEmails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkWardenServiceClient) RegisterHolder(ctx context.Context, in *RegisterHolderRequest, opts ...grpc.CallOption) (*RegisterHolderResponse, error) {
	out := new(RegisterHolderResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_RegisterHolder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkWardenServiceClient) ConfirmHolderRegistration(ctx context.Context, in *ConfirmHolderRegistrationRequest, opts ...grpc.CallOption) (*ConfirmHolderRegistrationResponse, error) {
	out := new(ConfirmHolderRegistrationResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_ConfirmHolderRegistration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkWardenServiceClient) ResendConfirmationCode(ctx context.Context, in *ResendConfirmationCodeRequest, opts ...grpc.CallOption) (*ResendConfirmationCodeResponse, error) {
	out := new(ResendConfirmationCodeResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_ResendConfirmationCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkWardenServiceClient) LoginHolder(ctx context.Context, in *LoginHolderRequest, opts ...grpc.CallOption) (*LoginHolderResponse, error) {
	out := new(LoginHolderResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_LoginHolder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkWardenServiceClient) LogoutHolder(ctx context.Context, in *LogoutHolderRequest, opts ...grpc.CallOption) (*LogoutHolderResponse, error) {
	out := new(LogoutHolderResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_LogoutHolder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkWardenServiceClient) RefreshHolderToken(ctx context.Context, in *RefreshHolderTokenRequest, opts ...grpc.CallOption) (*RefreshHolderTokenResponse, error) {
	out := new(RefreshHolderTokenResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_RefreshHolderToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkWardenServiceClient) ChangeHolderPassword(ctx context.Context, in *ChangeHolderPasswordRequest, opts ...grpc.CallOption) (*ChangeHolderPasswordResponse, error) {
	out := new(ChangeHolderPasswordResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_ChangeHolderPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkWardenServiceClient) ModifyHolder(ctx context.Context, in *ModifyHolderRequest, opts ...grpc.CallOption) (*ModifyHolderResponse, error) {
	out := new(ModifyHolderResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_ModifyHolder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkWardenServiceClient) GetHolder(ctx context.Context, in *GetHolderRequest, opts ...grpc.CallOption) (*GetHolderResponse, error) {
	out := new(GetHolderResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_GetHolder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkWardenServiceClient) DeleteHolder(ctx context.Context, in *DeleteHolderRequest, opts ...grpc.CallOption) (*DeleteHolderResponse, error) {
	out := new(DeleteHolderResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_DeleteHolder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkWardenServiceClient) GetPersonalDataNodesList(ctx context.Context, in *GetPersonalDataNodesListRequest, opts ...grpc.CallOption) (*GetPersonalDataNodesListResponse, error) {
	out := new(GetPersonalDataNodesListResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_GetPersonalDataNodesList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkWardenServiceClient) JoinPersonalDataNodeRegistrationWaitlist(ctx context.Context, in *JoinPersonalDataNodeRegistrationWaitlistRequest, opts ...grpc.CallOption) (*JoinPersonalDataNodeRegistrationWaitlistResponse, error) {
	out := new(JoinPersonalDataNodeRegistrationWaitlistResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_JoinPersonalDataNodeRegistrationWaitlist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkWardenServiceClient) RegisterPersonalDataNode(ctx context.Context, in *RegisterPersonalDataNodeRequest, opts ...grpc.CallOption) (*RegisterPersonalDataNodeResponse, error) {
	out := new(RegisterPersonalDataNodeResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_RegisterPersonalDataNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkWardenServiceClient) GetNetworkNodesList(ctx context.Context, in *GetNetworkNodesListRequest, opts ...grpc.CallOption) (*GetNetworkNodesListResponse, error) {
	out := new(GetNetworkNodesListResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_GetNetworkNodesList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkWardenServiceClient) JoinNetworkNodeRegistrationWaitlist(ctx context.Context, in *JoinNetworkNodeRegistrationWaitlistRequest, opts ...grpc.CallOption) (*JoinNetworkNodeRegistrationWaitlistResponse, error) {
	out := new(JoinNetworkNodeRegistrationWaitlistResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_JoinNetworkNodeRegistrationWaitlist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkWardenServiceClient) RegisterNetworkNode(ctx context.Context, in *RegisterNetworkNodeRequest, opts ...grpc.CallOption) (*RegisterNetworkNodeResponse, error) {
	out := new(RegisterNetworkNodeResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_RegisterNetworkNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkWardenServiceClient) GetNetworkWardensList(ctx context.Context, in *GetNetworkWardensListRequest, opts ...grpc.CallOption) (*GetNetworkWardensListResponse, error) {
	out := new(GetNetworkWardensListResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_GetNetworkWardensList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkWardenServiceClient) RegisterNetworkWarden(ctx context.Context, in *RegisterNetworkWardenRequest, opts ...grpc.CallOption) (*RegisterNetworkWardenResponse, error) {
	out := new(RegisterNetworkWardenResponse)
	err := c.cc.Invoke(ctx, NetworkWardenService_RegisterNetworkWarden_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkWardenServiceServer is the server API for NetworkWardenService service.
// All implementations must embed UnimplementedNetworkWardenServiceServer
// for forward compatibility
type NetworkWardenServiceServer interface {
	CheckEmails(context.Context, *CheckEmailsRequest) (*CheckEmailsResponse, error)
	RegisterHolder(context.Context, *RegisterHolderRequest) (*RegisterHolderResponse, error)
	ConfirmHolderRegistration(context.Context, *ConfirmHolderRegistrationRequest) (*ConfirmHolderRegistrationResponse, error)
	ResendConfirmationCode(context.Context, *ResendConfirmationCodeRequest) (*ResendConfirmationCodeResponse, error)
	LoginHolder(context.Context, *LoginHolderRequest) (*LoginHolderResponse, error)
	LogoutHolder(context.Context, *LogoutHolderRequest) (*LogoutHolderResponse, error)
	RefreshHolderToken(context.Context, *RefreshHolderTokenRequest) (*RefreshHolderTokenResponse, error)
	ChangeHolderPassword(context.Context, *ChangeHolderPasswordRequest) (*ChangeHolderPasswordResponse, error)
	ModifyHolder(context.Context, *ModifyHolderRequest) (*ModifyHolderResponse, error)
	GetHolder(context.Context, *GetHolderRequest) (*GetHolderResponse, error)
	DeleteHolder(context.Context, *DeleteHolderRequest) (*DeleteHolderResponse, error)
	GetPersonalDataNodesList(context.Context, *GetPersonalDataNodesListRequest) (*GetPersonalDataNodesListResponse, error)
	JoinPersonalDataNodeRegistrationWaitlist(context.Context, *JoinPersonalDataNodeRegistrationWaitlistRequest) (*JoinPersonalDataNodeRegistrationWaitlistResponse, error)
	RegisterPersonalDataNode(context.Context, *RegisterPersonalDataNodeRequest) (*RegisterPersonalDataNodeResponse, error)
	GetNetworkNodesList(context.Context, *GetNetworkNodesListRequest) (*GetNetworkNodesListResponse, error)
	JoinNetworkNodeRegistrationWaitlist(context.Context, *JoinNetworkNodeRegistrationWaitlistRequest) (*JoinNetworkNodeRegistrationWaitlistResponse, error)
	RegisterNetworkNode(context.Context, *RegisterNetworkNodeRequest) (*RegisterNetworkNodeResponse, error)
	GetNetworkWardensList(context.Context, *GetNetworkWardensListRequest) (*GetNetworkWardensListResponse, error)
	RegisterNetworkWarden(context.Context, *RegisterNetworkWardenRequest) (*RegisterNetworkWardenResponse, error)
	mustEmbedUnimplementedNetworkWardenServiceServer()
}

// UnimplementedNetworkWardenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkWardenServiceServer struct {
}

func (UnimplementedNetworkWardenServiceServer) CheckEmails(context.Context, *CheckEmailsRequest) (*CheckEmailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckEmails not implemented")
}
func (UnimplementedNetworkWardenServiceServer) RegisterHolder(context.Context, *RegisterHolderRequest) (*RegisterHolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterHolder not implemented")
}
func (UnimplementedNetworkWardenServiceServer) ConfirmHolderRegistration(context.Context, *ConfirmHolderRegistrationRequest) (*ConfirmHolderRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmHolderRegistration not implemented")
}
func (UnimplementedNetworkWardenServiceServer) ResendConfirmationCode(context.Context, *ResendConfirmationCodeRequest) (*ResendConfirmationCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResendConfirmationCode not implemented")
}
func (UnimplementedNetworkWardenServiceServer) LoginHolder(context.Context, *LoginHolderRequest) (*LoginHolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginHolder not implemented")
}
func (UnimplementedNetworkWardenServiceServer) LogoutHolder(context.Context, *LogoutHolderRequest) (*LogoutHolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogoutHolder not implemented")
}
func (UnimplementedNetworkWardenServiceServer) RefreshHolderToken(context.Context, *RefreshHolderTokenRequest) (*RefreshHolderTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshHolderToken not implemented")
}
func (UnimplementedNetworkWardenServiceServer) ChangeHolderPassword(context.Context, *ChangeHolderPasswordRequest) (*ChangeHolderPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeHolderPassword not implemented")
}
func (UnimplementedNetworkWardenServiceServer) ModifyHolder(context.Context, *ModifyHolderRequest) (*ModifyHolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyHolder not implemented")
}
func (UnimplementedNetworkWardenServiceServer) GetHolder(context.Context, *GetHolderRequest) (*GetHolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHolder not implemented")
}
func (UnimplementedNetworkWardenServiceServer) DeleteHolder(context.Context, *DeleteHolderRequest) (*DeleteHolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHolder not implemented")
}
func (UnimplementedNetworkWardenServiceServer) GetPersonalDataNodesList(context.Context, *GetPersonalDataNodesListRequest) (*GetPersonalDataNodesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonalDataNodesList not implemented")
}
func (UnimplementedNetworkWardenServiceServer) JoinPersonalDataNodeRegistrationWaitlist(context.Context, *JoinPersonalDataNodeRegistrationWaitlistRequest) (*JoinPersonalDataNodeRegistrationWaitlistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinPersonalDataNodeRegistrationWaitlist not implemented")
}
func (UnimplementedNetworkWardenServiceServer) RegisterPersonalDataNode(context.Context, *RegisterPersonalDataNodeRequest) (*RegisterPersonalDataNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPersonalDataNode not implemented")
}
func (UnimplementedNetworkWardenServiceServer) GetNetworkNodesList(context.Context, *GetNetworkNodesListRequest) (*GetNetworkNodesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkNodesList not implemented")
}
func (UnimplementedNetworkWardenServiceServer) JoinNetworkNodeRegistrationWaitlist(context.Context, *JoinNetworkNodeRegistrationWaitlistRequest) (*JoinNetworkNodeRegistrationWaitlistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinNetworkNodeRegistrationWaitlist not implemented")
}
func (UnimplementedNetworkWardenServiceServer) RegisterNetworkNode(context.Context, *RegisterNetworkNodeRequest) (*RegisterNetworkNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNetworkNode not implemented")
}
func (UnimplementedNetworkWardenServiceServer) GetNetworkWardensList(context.Context, *GetNetworkWardensListRequest) (*GetNetworkWardensListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkWardensList not implemented")
}
func (UnimplementedNetworkWardenServiceServer) RegisterNetworkWarden(context.Context, *RegisterNetworkWardenRequest) (*RegisterNetworkWardenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNetworkWarden not implemented")
}
func (UnimplementedNetworkWardenServiceServer) mustEmbedUnimplementedNetworkWardenServiceServer() {}

// UnsafeNetworkWardenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkWardenServiceServer will
// result in compilation errors.
type UnsafeNetworkWardenServiceServer interface {
	mustEmbedUnimplementedNetworkWardenServiceServer()
}

func RegisterNetworkWardenServiceServer(s grpc.ServiceRegistrar, srv NetworkWardenServiceServer) {
	s.RegisterService(&NetworkWardenService_ServiceDesc, srv)
}

func _NetworkWardenService_CheckEmails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckEmailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).CheckEmails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_CheckEmails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).CheckEmails(ctx, req.(*CheckEmailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkWardenService_RegisterHolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterHolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).RegisterHolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_RegisterHolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).RegisterHolder(ctx, req.(*RegisterHolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkWardenService_ConfirmHolderRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmHolderRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).ConfirmHolderRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_ConfirmHolderRegistration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).ConfirmHolderRegistration(ctx, req.(*ConfirmHolderRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkWardenService_ResendConfirmationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResendConfirmationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).ResendConfirmationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_ResendConfirmationCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).ResendConfirmationCode(ctx, req.(*ResendConfirmationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkWardenService_LoginHolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginHolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).LoginHolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_LoginHolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).LoginHolder(ctx, req.(*LoginHolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkWardenService_LogoutHolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutHolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).LogoutHolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_LogoutHolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).LogoutHolder(ctx, req.(*LogoutHolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkWardenService_RefreshHolderToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshHolderTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).RefreshHolderToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_RefreshHolderToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).RefreshHolderToken(ctx, req.(*RefreshHolderTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkWardenService_ChangeHolderPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeHolderPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).ChangeHolderPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_ChangeHolderPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).ChangeHolderPassword(ctx, req.(*ChangeHolderPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkWardenService_ModifyHolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyHolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).ModifyHolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_ModifyHolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).ModifyHolder(ctx, req.(*ModifyHolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkWardenService_GetHolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).GetHolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_GetHolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).GetHolder(ctx, req.(*GetHolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkWardenService_DeleteHolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).DeleteHolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_DeleteHolder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).DeleteHolder(ctx, req.(*DeleteHolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkWardenService_GetPersonalDataNodesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonalDataNodesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).GetPersonalDataNodesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_GetPersonalDataNodesList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).GetPersonalDataNodesList(ctx, req.(*GetPersonalDataNodesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkWardenService_JoinPersonalDataNodeRegistrationWaitlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinPersonalDataNodeRegistrationWaitlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).JoinPersonalDataNodeRegistrationWaitlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_JoinPersonalDataNodeRegistrationWaitlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).JoinPersonalDataNodeRegistrationWaitlist(ctx, req.(*JoinPersonalDataNodeRegistrationWaitlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkWardenService_RegisterPersonalDataNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterPersonalDataNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).RegisterPersonalDataNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_RegisterPersonalDataNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).RegisterPersonalDataNode(ctx, req.(*RegisterPersonalDataNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkWardenService_GetNetworkNodesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkNodesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).GetNetworkNodesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_GetNetworkNodesList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).GetNetworkNodesList(ctx, req.(*GetNetworkNodesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkWardenService_JoinNetworkNodeRegistrationWaitlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinNetworkNodeRegistrationWaitlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).JoinNetworkNodeRegistrationWaitlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_JoinNetworkNodeRegistrationWaitlist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).JoinNetworkNodeRegistrationWaitlist(ctx, req.(*JoinNetworkNodeRegistrationWaitlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkWardenService_RegisterNetworkNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNetworkNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).RegisterNetworkNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_RegisterNetworkNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).RegisterNetworkNode(ctx, req.(*RegisterNetworkNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkWardenService_GetNetworkWardensList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkWardensListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).GetNetworkWardensList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_GetNetworkWardensList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).GetNetworkWardensList(ctx, req.(*GetNetworkWardensListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkWardenService_RegisterNetworkWarden_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNetworkWardenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkWardenServiceServer).RegisterNetworkWarden(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkWardenService_RegisterNetworkWarden_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkWardenServiceServer).RegisterNetworkWarden(ctx, req.(*RegisterNetworkWardenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkWardenService_ServiceDesc is the grpc.ServiceDesc for NetworkWardenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkWardenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "networkwarden.v1.NetworkWardenService",
	HandlerType: (*NetworkWardenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckEmails",
			Handler:    _NetworkWardenService_CheckEmails_Handler,
		},
		{
			MethodName: "RegisterHolder",
			Handler:    _NetworkWardenService_RegisterHolder_Handler,
		},
		{
			MethodName: "ConfirmHolderRegistration",
			Handler:    _NetworkWardenService_ConfirmHolderRegistration_Handler,
		},
		{
			MethodName: "ResendConfirmationCode",
			Handler:    _NetworkWardenService_ResendConfirmationCode_Handler,
		},
		{
			MethodName: "LoginHolder",
			Handler:    _NetworkWardenService_LoginHolder_Handler,
		},
		{
			MethodName: "LogoutHolder",
			Handler:    _NetworkWardenService_LogoutHolder_Handler,
		},
		{
			MethodName: "RefreshHolderToken",
			Handler:    _NetworkWardenService_RefreshHolderToken_Handler,
		},
		{
			MethodName: "ChangeHolderPassword",
			Handler:    _NetworkWardenService_ChangeHolderPassword_Handler,
		},
		{
			MethodName: "ModifyHolder",
			Handler:    _NetworkWardenService_ModifyHolder_Handler,
		},
		{
			MethodName: "GetHolder",
			Handler:    _NetworkWardenService_GetHolder_Handler,
		},
		{
			MethodName: "DeleteHolder",
			Handler:    _NetworkWardenService_DeleteHolder_Handler,
		},
		{
			MethodName: "GetPersonalDataNodesList",
			Handler:    _NetworkWardenService_GetPersonalDataNodesList_Handler,
		},
		{
			MethodName: "JoinPersonalDataNodeRegistrationWaitlist",
			Handler:    _NetworkWardenService_JoinPersonalDataNodeRegistrationWaitlist_Handler,
		},
		{
			MethodName: "RegisterPersonalDataNode",
			Handler:    _NetworkWardenService_RegisterPersonalDataNode_Handler,
		},
		{
			MethodName: "GetNetworkNodesList",
			Handler:    _NetworkWardenService_GetNetworkNodesList_Handler,
		},
		{
			MethodName: "JoinNetworkNodeRegistrationWaitlist",
			Handler:    _NetworkWardenService_JoinNetworkNodeRegistrationWaitlist_Handler,
		},
		{
			MethodName: "RegisterNetworkNode",
			Handler:    _NetworkWardenService_RegisterNetworkNode_Handler,
		},
		{
			MethodName: "GetNetworkWardensList",
			Handler:    _NetworkWardenService_GetNetworkWardensList_Handler,
		},
		{
			MethodName: "RegisterNetworkWarden",
			Handler:    _NetworkWardenService_RegisterNetworkWarden_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "networkwarden/v1/network_warden_service.proto",
}
