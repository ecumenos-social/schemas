// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: personaldatanode/v1/blockchain_block.proto

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

type BlockchainBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *BlockchainBlock_Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   *BlockchainTransaction  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BlockchainBlock) Reset() {
	*x = BlockchainBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personaldatanode_v1_blockchain_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainBlock) ProtoMessage() {}

func (x *BlockchainBlock) ProtoReflect() protoreflect.Message {
	mi := &file_personaldatanode_v1_blockchain_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainBlock.ProtoReflect.Descriptor instead.
func (*BlockchainBlock) Descriptor() ([]byte, []int) {
	return file_personaldatanode_v1_blockchain_block_proto_rawDescGZIP(), []int{0}
}

func (x *BlockchainBlock) GetHeader() *BlockchainBlock_Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BlockchainBlock) GetData() *BlockchainTransaction {
	if x != nil {
		return x.Data
	}
	return nil
}

type BlockchainBlock_Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash            string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ParentBlockHash string `protobuf:"bytes,2,opt,name=parent_block_hash,json=parentBlockHash,proto3" json:"parent_block_hash,omitempty"`
	Timestamp       string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *BlockchainBlock_Header) Reset() {
	*x = BlockchainBlock_Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personaldatanode_v1_blockchain_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainBlock_Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainBlock_Header) ProtoMessage() {}

func (x *BlockchainBlock_Header) ProtoReflect() protoreflect.Message {
	mi := &file_personaldatanode_v1_blockchain_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainBlock_Header.ProtoReflect.Descriptor instead.
func (*BlockchainBlock_Header) Descriptor() ([]byte, []int) {
	return file_personaldatanode_v1_blockchain_block_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BlockchainBlock_Header) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *BlockchainBlock_Header) GetParentBlockHash() string {
	if x != nil {
		return x.ParentBlockHash
	}
	return ""
}

func (x *BlockchainBlock_Header) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_personaldatanode_v1_blockchain_block_proto protoreflect.FileDescriptor

var file_personaldatanode_v1_blockchain_block_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f,
	0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x30, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x6e,
	0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x01, 0x0a, 0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x43, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x66, 0x0a, 0x06,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0xde, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x14, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x6f, 0x73, 0x2d, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x6c, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x64,
	0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58,
	0xaa, 0x02, 0x13, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x14, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_personaldatanode_v1_blockchain_block_proto_rawDescOnce sync.Once
	file_personaldatanode_v1_blockchain_block_proto_rawDescData = file_personaldatanode_v1_blockchain_block_proto_rawDesc
)

func file_personaldatanode_v1_blockchain_block_proto_rawDescGZIP() []byte {
	file_personaldatanode_v1_blockchain_block_proto_rawDescOnce.Do(func() {
		file_personaldatanode_v1_blockchain_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_personaldatanode_v1_blockchain_block_proto_rawDescData)
	})
	return file_personaldatanode_v1_blockchain_block_proto_rawDescData
}

var file_personaldatanode_v1_blockchain_block_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_personaldatanode_v1_blockchain_block_proto_goTypes = []any{
	(*BlockchainBlock)(nil),        // 0: personaldatanode.v1.BlockchainBlock
	(*BlockchainBlock_Header)(nil), // 1: personaldatanode.v1.BlockchainBlock.Header
	(*BlockchainTransaction)(nil),  // 2: personaldatanode.v1.BlockchainTransaction
}
var file_personaldatanode_v1_blockchain_block_proto_depIdxs = []int32{
	1, // 0: personaldatanode.v1.BlockchainBlock.header:type_name -> personaldatanode.v1.BlockchainBlock.Header
	2, // 1: personaldatanode.v1.BlockchainBlock.data:type_name -> personaldatanode.v1.BlockchainTransaction
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_personaldatanode_v1_blockchain_block_proto_init() }
func file_personaldatanode_v1_blockchain_block_proto_init() {
	if File_personaldatanode_v1_blockchain_block_proto != nil {
		return
	}
	file_personaldatanode_v1_blockchain_transaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_personaldatanode_v1_blockchain_block_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BlockchainBlock); i {
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
		file_personaldatanode_v1_blockchain_block_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BlockchainBlock_Header); i {
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
			RawDescriptor: file_personaldatanode_v1_blockchain_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_personaldatanode_v1_blockchain_block_proto_goTypes,
		DependencyIndexes: file_personaldatanode_v1_blockchain_block_proto_depIdxs,
		MessageInfos:      file_personaldatanode_v1_blockchain_block_proto_msgTypes,
	}.Build()
	File_personaldatanode_v1_blockchain_block_proto = out.File
	file_personaldatanode_v1_blockchain_block_proto_rawDesc = nil
	file_personaldatanode_v1_blockchain_block_proto_goTypes = nil
	file_personaldatanode_v1_blockchain_block_proto_depIdxs = nil
}