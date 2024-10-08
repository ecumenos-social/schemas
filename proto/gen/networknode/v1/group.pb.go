// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: networknode/v1/group.proto

package v1

import (
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

type Group_Privacy int32

const (
	Group_PRIVACY_UNKNOWN_UNSPECIFIED Group_Privacy = 0
	Group_PRIVACY_PUBLIC              Group_Privacy = 1
	Group_PRIVACY_PRIVATE             Group_Privacy = 2
)

// Enum value maps for Group_Privacy.
var (
	Group_Privacy_name = map[int32]string{
		0: "PRIVACY_UNKNOWN_UNSPECIFIED",
		1: "PRIVACY_PUBLIC",
		2: "PRIVACY_PRIVATE",
	}
	Group_Privacy_value = map[string]int32{
		"PRIVACY_UNKNOWN_UNSPECIFIED": 0,
		"PRIVACY_PUBLIC":              1,
		"PRIVACY_PRIVATE":             2,
	}
)

func (x Group_Privacy) Enum() *Group_Privacy {
	p := new(Group_Privacy)
	*p = x
	return p
}

func (x Group_Privacy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Group_Privacy) Descriptor() protoreflect.EnumDescriptor {
	return file_networknode_v1_group_proto_enumTypes[0].Descriptor()
}

func (Group_Privacy) Type() protoreflect.EnumType {
	return &file_networknode_v1_group_proto_enumTypes[0]
}

func (x Group_Privacy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Group_Privacy.Descriptor instead.
func (Group_Privacy) EnumDescriptor() ([]byte, []int) {
	return file_networknode_v1_group_proto_rawDescGZIP(), []int{0, 0}
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt           string        `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastModifiedAt      *string       `protobuf:"bytes,3,opt,name=last_modified_at,json=lastModifiedAt,proto3,oneof" json:"last_modified_at,omitempty"`
	Name                string        `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	NameLanguage        *string       `protobuf:"bytes,5,opt,name=name_language,json=nameLanguage,proto3,oneof" json:"name_language,omitempty"`
	Description         string        `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	DescriptionLanguage *string       `protobuf:"bytes,7,opt,name=description_language,json=descriptionLanguage,proto3,oneof" json:"description_language,omitempty"`
	ProfileImageUrl     string        `protobuf:"bytes,8,opt,name=profile_image_url,json=profileImageUrl,proto3" json:"profile_image_url,omitempty"`
	HeaderImageUrl      string        `protobuf:"bytes,9,opt,name=header_image_url,json=headerImageUrl,proto3" json:"header_image_url,omitempty"`
	IsAnonymous         bool          `protobuf:"varint,10,opt,name=is_anonymous,json=isAnonymous,proto3" json:"is_anonymous,omitempty"`
	Privacy             Group_Privacy `protobuf:"varint,11,opt,name=privacy,proto3,enum=networknode.v1.Group_Privacy" json:"privacy,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networknode_v1_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_networknode_v1_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_networknode_v1_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Group) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Group) GetLastModifiedAt() string {
	if x != nil && x.LastModifiedAt != nil {
		return *x.LastModifiedAt
	}
	return ""
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetNameLanguage() string {
	if x != nil && x.NameLanguage != nil {
		return *x.NameLanguage
	}
	return ""
}

func (x *Group) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Group) GetDescriptionLanguage() string {
	if x != nil && x.DescriptionLanguage != nil {
		return *x.DescriptionLanguage
	}
	return ""
}

func (x *Group) GetProfileImageUrl() string {
	if x != nil {
		return x.ProfileImageUrl
	}
	return ""
}

func (x *Group) GetHeaderImageUrl() string {
	if x != nil {
		return x.HeaderImageUrl
	}
	return ""
}

func (x *Group) GetIsAnonymous() bool {
	if x != nil {
		return x.IsAnonymous
	}
	return false
}

func (x *Group) GetPrivacy() Group_Privacy {
	if x != nil {
		return x.Privacy
	}
	return Group_PRIVACY_UNKNOWN_UNSPECIFIED
}

var File_networknode_v1_group_proto protoreflect.FileDescriptor

var file_networknode_v1_group_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xc4, 0x04, 0x0a,
	0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2d, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0d, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x14, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x02, 0x52, 0x13, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x11,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f,
	0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x6e, 0x6f, 0x6e,
	0x79, 0x6d, 0x6f, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x22, 0x53,
	0x0a, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x52, 0x49,
	0x56, 0x41, 0x43, 0x59, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x52,
	0x49, 0x56, 0x41, 0x43, 0x59, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x01, 0x12, 0x13,
	0x0a, 0x0f, 0x50, 0x52, 0x49, 0x56, 0x41, 0x43, 0x59, 0x5f, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54,
	0x45, 0x10, 0x02, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x42, 0xb6, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x6f, 0x73, 0x2d, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x6c, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x6e, 0x6f,
	0x64, 0x65, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4e, 0x58, 0x58, 0xaa, 0x02, 0x0e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x6e, 0x6f, 0x64, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x6e, 0x6f, 0x64, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x6e, 0x6f, 0x64, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_networknode_v1_group_proto_rawDescOnce sync.Once
	file_networknode_v1_group_proto_rawDescData = file_networknode_v1_group_proto_rawDesc
)

func file_networknode_v1_group_proto_rawDescGZIP() []byte {
	file_networknode_v1_group_proto_rawDescOnce.Do(func() {
		file_networknode_v1_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_networknode_v1_group_proto_rawDescData)
	})
	return file_networknode_v1_group_proto_rawDescData
}

var file_networknode_v1_group_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_networknode_v1_group_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_networknode_v1_group_proto_goTypes = []any{
	(Group_Privacy)(0), // 0: networknode.v1.Group.Privacy
	(*Group)(nil),      // 1: networknode.v1.Group
}
var file_networknode_v1_group_proto_depIdxs = []int32{
	0, // 0: networknode.v1.Group.privacy:type_name -> networknode.v1.Group.Privacy
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_networknode_v1_group_proto_init() }
func file_networknode_v1_group_proto_init() {
	if File_networknode_v1_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_networknode_v1_group_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Group); i {
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
	file_networknode_v1_group_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_networknode_v1_group_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_networknode_v1_group_proto_goTypes,
		DependencyIndexes: file_networknode_v1_group_proto_depIdxs,
		EnumInfos:         file_networknode_v1_group_proto_enumTypes,
		MessageInfos:      file_networknode_v1_group_proto_msgTypes,
	}.Build()
	File_networknode_v1_group_proto = out.File
	file_networknode_v1_group_proto_rawDesc = nil
	file_networknode_v1_group_proto_goTypes = nil
	file_networknode_v1_group_proto_depIdxs = nil
}
