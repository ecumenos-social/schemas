// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: networkwarden/v1/network_node.proto

package v1

import (
	v1 "github.com/ecumenos-social/schemas/proto/gen/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NetworkNode_Status int32

const (
	NetworkNode_STATUS_UNKNOWN_UNSPECIFIED NetworkNode_Status = 0
	NetworkNode_STATUS_APPROVED            NetworkNode_Status = 1
	NetworkNode_STATUS_PENDING             NetworkNode_Status = 2
	NetworkNode_STATUS_REJECTED            NetworkNode_Status = 3
)

// Enum value maps for NetworkNode_Status.
var (
	NetworkNode_Status_name = map[int32]string{
		0: "STATUS_UNKNOWN_UNSPECIFIED",
		1: "STATUS_APPROVED",
		2: "STATUS_PENDING",
		3: "STATUS_REJECTED",
	}
	NetworkNode_Status_value = map[string]int32{
		"STATUS_UNKNOWN_UNSPECIFIED": 0,
		"STATUS_APPROVED":            1,
		"STATUS_PENDING":             2,
		"STATUS_REJECTED":            3,
	}
)

func (x NetworkNode_Status) Enum() *NetworkNode_Status {
	p := new(NetworkNode_Status)
	*p = x
	return p
}

func (x NetworkNode_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkNode_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_networkwarden_v1_network_node_proto_enumTypes[0].Descriptor()
}

func (NetworkNode_Status) Type() protoreflect.EnumType {
	return &file_networkwarden_v1_network_node_proto_enumTypes[0]
}

func (x NetworkNode_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkNode_Status.Descriptor instead.
func (NetworkNode_Status) EnumDescriptor() ([]byte, []int) {
	return file_networkwarden_v1_network_node_proto_rawDescGZIP(), []int{0, 0}
}

type NetworkNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt        string             `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastModifiedAt   string             `protobuf:"bytes,3,opt,name=last_modified_at,json=lastModifiedAt,proto3" json:"last_modified_at,omitempty"`
	NwId             string             `protobuf:"bytes,4,opt,name=nw_id,json=nwId,proto3" json:"nw_id,omitempty"`
	Name             string             `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	DomainName       string             `protobuf:"bytes,6,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty"`
	Location         *v1.Geolocation    `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
	AccountsCapacity int64              `protobuf:"varint,8,opt,name=accounts_capacity,json=accountsCapacity,proto3" json:"accounts_capacity,omitempty"`
	Alive            bool               `protobuf:"varint,9,opt,name=alive,proto3" json:"alive,omitempty"`
	LastPingedAt     string             `protobuf:"bytes,10,opt,name=last_pinged_at,json=lastPingedAt,proto3" json:"last_pinged_at,omitempty"`
	IsOpen           bool               `protobuf:"varint,11,opt,name=is_open,json=isOpen,proto3" json:"is_open,omitempty"`
	OwnerHolderId    string             `protobuf:"bytes,12,opt,name=owner_holder_id,json=ownerHolderId,proto3" json:"owner_holder_id,omitempty"`
	Url              string             `protobuf:"bytes,13,opt,name=url,proto3" json:"url,omitempty"`
	Version          string             `protobuf:"bytes,14,opt,name=version,proto3" json:"version,omitempty"`
	RateLimit        float64            `protobuf:"fixed64,15,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	CrawlRateLimit   float64            `protobuf:"fixed64,16,opt,name=crawl_rate_limit,json=crawlRateLimit,proto3" json:"crawl_rate_limit,omitempty"`
	Status           NetworkNode_Status `protobuf:"varint,17,opt,name=status,proto3,enum=networkwarden.v1.NetworkNode_Status" json:"status,omitempty"`
}

func (x *NetworkNode) Reset() {
	*x = NetworkNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networkwarden_v1_network_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkNode) ProtoMessage() {}

func (x *NetworkNode) ProtoReflect() protoreflect.Message {
	mi := &file_networkwarden_v1_network_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkNode.ProtoReflect.Descriptor instead.
func (*NetworkNode) Descriptor() ([]byte, []int) {
	return file_networkwarden_v1_network_node_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkNode) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NetworkNode) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *NetworkNode) GetLastModifiedAt() string {
	if x != nil {
		return x.LastModifiedAt
	}
	return ""
}

func (x *NetworkNode) GetNwId() string {
	if x != nil {
		return x.NwId
	}
	return ""
}

func (x *NetworkNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NetworkNode) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

func (x *NetworkNode) GetLocation() *v1.Geolocation {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *NetworkNode) GetAccountsCapacity() int64 {
	if x != nil {
		return x.AccountsCapacity
	}
	return 0
}

func (x *NetworkNode) GetAlive() bool {
	if x != nil {
		return x.Alive
	}
	return false
}

func (x *NetworkNode) GetLastPingedAt() string {
	if x != nil {
		return x.LastPingedAt
	}
	return ""
}

func (x *NetworkNode) GetIsOpen() bool {
	if x != nil {
		return x.IsOpen
	}
	return false
}

func (x *NetworkNode) GetOwnerHolderId() string {
	if x != nil {
		return x.OwnerHolderId
	}
	return ""
}

func (x *NetworkNode) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *NetworkNode) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *NetworkNode) GetRateLimit() float64 {
	if x != nil {
		return x.RateLimit
	}
	return 0
}

func (x *NetworkNode) GetCrawlRateLimit() float64 {
	if x != nil {
		return x.CrawlRateLimit
	}
	return 0
}

func (x *NetworkNode) GetStatus() NetworkNode_Status {
	if x != nil {
		return x.Status
	}
	return NetworkNode_STATUS_UNKNOWN_UNSPECIFIED
}

var File_networkwarden_v1_network_node_proto protoreflect.FileDescriptor

var file_networkwarden_v1_network_node_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x77, 0x61,
	0x72, 0x64, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x65, 0x6f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x05, 0x0a, 0x0b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c,
	0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x12, 0x13, 0x0a,
	0x05, 0x6e, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x77,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x76,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x24,
	0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x70, 0x69, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x69, 0x6e, 0x67,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x26, 0x0a,
	0x0f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x48, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x28, 0x0a, 0x10, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x63, 0x72, 0x61, 0x77,
	0x6c, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x66, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x50, 0x50,
	0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x03,
	0x42, 0xc8, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x6f, 0x73, 0x2d, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4e,
	0x58, 0x58, 0xaa, 0x02, 0x10, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x77, 0x61, 0x72, 0x64,
	0x65, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x77,
	0x61, 0x72, 0x64, 0x65, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_networkwarden_v1_network_node_proto_rawDescOnce sync.Once
	file_networkwarden_v1_network_node_proto_rawDescData = file_networkwarden_v1_network_node_proto_rawDesc
)

func file_networkwarden_v1_network_node_proto_rawDescGZIP() []byte {
	file_networkwarden_v1_network_node_proto_rawDescOnce.Do(func() {
		file_networkwarden_v1_network_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_networkwarden_v1_network_node_proto_rawDescData)
	})
	return file_networkwarden_v1_network_node_proto_rawDescData
}

var file_networkwarden_v1_network_node_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_networkwarden_v1_network_node_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_networkwarden_v1_network_node_proto_goTypes = []interface{}{
	(NetworkNode_Status)(0), // 0: networkwarden.v1.NetworkNode.Status
	(*NetworkNode)(nil),     // 1: networkwarden.v1.NetworkNode
	(*v1.Geolocation)(nil),  // 2: common.v1.Geolocation
}
var file_networkwarden_v1_network_node_proto_depIdxs = []int32{
	2, // 0: networkwarden.v1.NetworkNode.location:type_name -> common.v1.Geolocation
	0, // 1: networkwarden.v1.NetworkNode.status:type_name -> networkwarden.v1.NetworkNode.Status
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_networkwarden_v1_network_node_proto_init() }
func file_networkwarden_v1_network_node_proto_init() {
	if File_networkwarden_v1_network_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_networkwarden_v1_network_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkNode); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_networkwarden_v1_network_node_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_networkwarden_v1_network_node_proto_goTypes,
		DependencyIndexes: file_networkwarden_v1_network_node_proto_depIdxs,
		EnumInfos:         file_networkwarden_v1_network_node_proto_enumTypes,
		MessageInfos:      file_networkwarden_v1_network_node_proto_msgTypes,
	}.Build()
	File_networkwarden_v1_network_node_proto = out.File
	file_networkwarden_v1_network_node_proto_rawDesc = nil
	file_networkwarden_v1_network_node_proto_goTypes = nil
	file_networkwarden_v1_network_node_proto_depIdxs = nil
}