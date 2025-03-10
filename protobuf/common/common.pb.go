//
// Protobuf file containing empty message definition

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: common/common.proto

package common

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

type NodeVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major      uint32  `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor      uint32  `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch      uint32  `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	Prerelease *string `protobuf:"bytes,4,opt,name=prerelease,proto3,oneof" json:"prerelease,omitempty"`
}

func (x *NodeVersion) Reset() {
	*x = NodeVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeVersion) ProtoMessage() {}

func (x *NodeVersion) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeVersion.ProtoReflect.Descriptor instead.
func (*NodeVersion) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *NodeVersion) GetMajor() uint32 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *NodeVersion) GetMinor() uint32 {
	if x != nil {
		return x.Minor
	}
	return 0
}

func (x *NodeVersion) GetPatch() uint32 {
	if x != nil {
		return x.Patch
	}
	return 0
}

func (x *NodeVersion) GetPrerelease() string {
	if x != nil && x.Prerelease != nil {
		return *x.Prerelease
	}
	return ""
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeVersion *NodeVersion `protobuf:"bytes,1,opt,name=node_version,json=nodeVersion,proto3" json:"node_version,omitempty"`
	BeaconID    string       `protobuf:"bytes,2,opt,name=beaconID,proto3" json:"beaconID,omitempty"`
	ChainHash   []byte       `protobuf:"bytes,3,opt,name=chain_hash,json=chainHash,proto3" json:"chain_hash,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *Metadata) GetNodeVersion() *NodeVersion {
	if x != nil {
		return x.NodeVersion
	}
	return nil
}

func (x *Metadata) GetBeaconID() string {
	if x != nil {
		return x.BeaconID
	}
	return ""
}

func (x *Metadata) GetChainHash() []byte {
	if x != nil {
		return x.ChainHash
	}
	return nil
}

var File_common_common_proto protoreflect.FileDescriptor

var file_common_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x83, 0x01,
	0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d, 0x61,
	0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x23, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x72, 0x65, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x22, 0x7d, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x36, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x6e, 0x6f, 0x64, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x65, 0x61, 0x63, 0x6f,
	0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x65, 0x61, 0x63, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x48, 0x61,
	0x73, 0x68, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2f, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_common_proto_rawDescOnce sync.Once
	file_common_common_proto_rawDescData = file_common_common_proto_rawDesc
)

func file_common_common_proto_rawDescGZIP() []byte {
	file_common_common_proto_rawDescOnce.Do(func() {
		file_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_common_proto_rawDescData)
	})
	return file_common_common_proto_rawDescData
}

var file_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_common_proto_goTypes = []interface{}{
	(*NodeVersion)(nil), // 0: common.NodeVersion
	(*Metadata)(nil),    // 1: common.Metadata
}
var file_common_common_proto_depIdxs = []int32{
	0, // 0: common.Metadata.node_version:type_name -> common.NodeVersion
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_common_proto_init() }
func file_common_common_proto_init() {
	if File_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeVersion); i {
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
		file_common_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
	file_common_common_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_common_proto_goTypes,
		DependencyIndexes: file_common_common_proto_depIdxs,
		MessageInfos:      file_common_common_proto_msgTypes,
	}.Build()
	File_common_common_proto = out.File
	file_common_common_proto_rawDesc = nil
	file_common_common_proto_goTypes = nil
	file_common_common_proto_depIdxs = nil
}
