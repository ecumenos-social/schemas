// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: networkwarden/v1/personal_data_node.proto

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

type PersonalDataNode_Status int32

const (
	PersonalDataNode_STATUS_UNKNOWN_UNSPECIFIED PersonalDataNode_Status = 0
	PersonalDataNode_STATUS_APPROVED            PersonalDataNode_Status = 1
	PersonalDataNode_STATUS_PENDING             PersonalDataNode_Status = 2
	PersonalDataNode_STATUS_REJECTED            PersonalDataNode_Status = 3
)

// Enum value maps for PersonalDataNode_Status.
var (
	PersonalDataNode_Status_name = map[int32]string{
		0: "STATUS_UNKNOWN_UNSPECIFIED",
		1: "STATUS_APPROVED",
		2: "STATUS_PENDING",
		3: "STATUS_REJECTED",
	}
	PersonalDataNode_Status_value = map[string]int32{
		"STATUS_UNKNOWN_UNSPECIFIED": 0,
		"STATUS_APPROVED":            1,
		"STATUS_PENDING":             2,
		"STATUS_REJECTED":            3,
	}
)

func (x PersonalDataNode_Status) Enum() *PersonalDataNode_Status {
	p := new(PersonalDataNode_Status)
	*p = x
	return p
}

func (x PersonalDataNode_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PersonalDataNode_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_networkwarden_v1_personal_data_node_proto_enumTypes[0].Descriptor()
}

func (PersonalDataNode_Status) Type() protoreflect.EnumType {
	return &file_networkwarden_v1_personal_data_node_proto_enumTypes[0]
}

func (x PersonalDataNode_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PersonalDataNode_Status.Descriptor instead.
func (PersonalDataNode_Status) EnumDescriptor() ([]byte, []int) {
	return file_networkwarden_v1_personal_data_node_proto_rawDescGZIP(), []int{0, 0}
}

type PersonalDataNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            string                  `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastModifiedAt       string                  `protobuf:"bytes,3,opt,name=last_modified_at,json=lastModifiedAt,proto3" json:"last_modified_at,omitempty"`
	NwId                 string                  `protobuf:"bytes,4,opt,name=nw_id,json=nwId,proto3" json:"nw_id,omitempty"`
	Address              string                  `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Label                string                  `protobuf:"bytes,6,opt,name=label,proto3" json:"label,omitempty"`
	Name                 string                  `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Description          string                  `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	Location             *v1.Geolocation         `protobuf:"bytes,9,opt,name=location,proto3" json:"location,omitempty"`
	AccountsCapacity     int64                   `protobuf:"varint,10,opt,name=accounts_capacity,json=accountsCapacity,proto3" json:"accounts_capacity,omitempty"`
	Alive                bool                    `protobuf:"varint,11,opt,name=alive,proto3" json:"alive,omitempty"`
	LastPingedAt         string                  `protobuf:"bytes,12,opt,name=last_pinged_at,json=lastPingedAt,proto3" json:"last_pinged_at,omitempty"`
	IsOpen               bool                    `protobuf:"varint,13,opt,name=is_open,json=isOpen,proto3" json:"is_open,omitempty"`
	IsInviteCodeRequired bool                    `protobuf:"varint,14,opt,name=is_invite_code_required,json=isInviteCodeRequired,proto3" json:"is_invite_code_required,omitempty"`
	OwnerHolderId        string                  `protobuf:"bytes,15,opt,name=owner_holder_id,json=ownerHolderId,proto3" json:"owner_holder_id,omitempty"`
	Url                  string                  `protobuf:"bytes,16,opt,name=url,proto3" json:"url,omitempty"`
	Version              string                  `protobuf:"bytes,17,opt,name=version,proto3" json:"version,omitempty"`
	RateLimit            *v1.RateLimit           `protobuf:"bytes,18,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	CrawlRateLimit       *v1.RateLimit           `protobuf:"bytes,19,opt,name=crawl_rate_limit,json=crawlRateLimit,proto3" json:"crawl_rate_limit,omitempty"`
	IdGenNode            int64                   `protobuf:"varint,20,opt,name=id_gen_node,json=idGenNode,proto3" json:"id_gen_node,omitempty"`
	Status               PersonalDataNode_Status `protobuf:"varint,21,opt,name=status,proto3,enum=networkwarden.v1.PersonalDataNode_Status" json:"status,omitempty"`
}

func (x *PersonalDataNode) Reset() {
	*x = PersonalDataNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networkwarden_v1_personal_data_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonalDataNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonalDataNode) ProtoMessage() {}

func (x *PersonalDataNode) ProtoReflect() protoreflect.Message {
	mi := &file_networkwarden_v1_personal_data_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonalDataNode.ProtoReflect.Descriptor instead.
func (*PersonalDataNode) Descriptor() ([]byte, []int) {
	return file_networkwarden_v1_personal_data_node_proto_rawDescGZIP(), []int{0}
}

func (x *PersonalDataNode) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PersonalDataNode) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PersonalDataNode) GetLastModifiedAt() string {
	if x != nil {
		return x.LastModifiedAt
	}
	return ""
}

func (x *PersonalDataNode) GetNwId() string {
	if x != nil {
		return x.NwId
	}
	return ""
}

func (x *PersonalDataNode) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PersonalDataNode) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *PersonalDataNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PersonalDataNode) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PersonalDataNode) GetLocation() *v1.Geolocation {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *PersonalDataNode) GetAccountsCapacity() int64 {
	if x != nil {
		return x.AccountsCapacity
	}
	return 0
}

func (x *PersonalDataNode) GetAlive() bool {
	if x != nil {
		return x.Alive
	}
	return false
}

func (x *PersonalDataNode) GetLastPingedAt() string {
	if x != nil {
		return x.LastPingedAt
	}
	return ""
}

func (x *PersonalDataNode) GetIsOpen() bool {
	if x != nil {
		return x.IsOpen
	}
	return false
}

func (x *PersonalDataNode) GetIsInviteCodeRequired() bool {
	if x != nil {
		return x.IsInviteCodeRequired
	}
	return false
}

func (x *PersonalDataNode) GetOwnerHolderId() string {
	if x != nil {
		return x.OwnerHolderId
	}
	return ""
}

func (x *PersonalDataNode) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PersonalDataNode) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PersonalDataNode) GetRateLimit() *v1.RateLimit {
	if x != nil {
		return x.RateLimit
	}
	return nil
}

func (x *PersonalDataNode) GetCrawlRateLimit() *v1.RateLimit {
	if x != nil {
		return x.CrawlRateLimit
	}
	return nil
}

func (x *PersonalDataNode) GetIdGenNode() int64 {
	if x != nil {
		return x.IdGenNode
	}
	return 0
}

func (x *PersonalDataNode) GetStatus() PersonalDataNode_Status {
	if x != nil {
		return x.Status
	}
	return PersonalDataNode_STATUS_UNKNOWN_UNSPECIFIED
}

var File_networkwarden_v1_personal_data_node_proto protoreflect.FileDescriptor

var file_networkwarden_v1_personal_data_node_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x06, 0x0a, 0x10, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x6e, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x77, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x32, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x6f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x5f,
	0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x70,
	0x69, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x65, 0x64, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69,
	0x73, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x35, 0x0a, 0x17, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0f,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x48, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x33, 0x0a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x3e, 0x0a, 0x10, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x0e, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x52, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x5f,
	0x6e, 0x6f, 0x64, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x64, 0x47, 0x65,
	0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x77,
	0x61, 0x72, 0x64, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x66, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x50, 0x50,
	0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x03,
	0x42, 0xcd, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x6f, 0x73, 0x2d, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2f, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x4e, 0x58, 0x58, 0xaa, 0x02, 0x10, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x77, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_networkwarden_v1_personal_data_node_proto_rawDescOnce sync.Once
	file_networkwarden_v1_personal_data_node_proto_rawDescData = file_networkwarden_v1_personal_data_node_proto_rawDesc
)

func file_networkwarden_v1_personal_data_node_proto_rawDescGZIP() []byte {
	file_networkwarden_v1_personal_data_node_proto_rawDescOnce.Do(func() {
		file_networkwarden_v1_personal_data_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_networkwarden_v1_personal_data_node_proto_rawDescData)
	})
	return file_networkwarden_v1_personal_data_node_proto_rawDescData
}

var file_networkwarden_v1_personal_data_node_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_networkwarden_v1_personal_data_node_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_networkwarden_v1_personal_data_node_proto_goTypes = []any{
	(PersonalDataNode_Status)(0), // 0: networkwarden.v1.PersonalDataNode.Status
	(*PersonalDataNode)(nil),     // 1: networkwarden.v1.PersonalDataNode
	(*v1.Geolocation)(nil),       // 2: common.v1.Geolocation
	(*v1.RateLimit)(nil),         // 3: common.v1.RateLimit
}
var file_networkwarden_v1_personal_data_node_proto_depIdxs = []int32{
	2, // 0: networkwarden.v1.PersonalDataNode.location:type_name -> common.v1.Geolocation
	3, // 1: networkwarden.v1.PersonalDataNode.rate_limit:type_name -> common.v1.RateLimit
	3, // 2: networkwarden.v1.PersonalDataNode.crawl_rate_limit:type_name -> common.v1.RateLimit
	0, // 3: networkwarden.v1.PersonalDataNode.status:type_name -> networkwarden.v1.PersonalDataNode.Status
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_networkwarden_v1_personal_data_node_proto_init() }
func file_networkwarden_v1_personal_data_node_proto_init() {
	if File_networkwarden_v1_personal_data_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_networkwarden_v1_personal_data_node_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PersonalDataNode); i {
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
			RawDescriptor: file_networkwarden_v1_personal_data_node_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_networkwarden_v1_personal_data_node_proto_goTypes,
		DependencyIndexes: file_networkwarden_v1_personal_data_node_proto_depIdxs,
		EnumInfos:         file_networkwarden_v1_personal_data_node_proto_enumTypes,
		MessageInfos:      file_networkwarden_v1_personal_data_node_proto_msgTypes,
	}.Build()
	File_networkwarden_v1_personal_data_node_proto = out.File
	file_networkwarden_v1_personal_data_node_proto_rawDesc = nil
	file_networkwarden_v1_personal_data_node_proto_goTypes = nil
	file_networkwarden_v1_personal_data_node_proto_depIdxs = nil
}
