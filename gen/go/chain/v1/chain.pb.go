// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: chain/v1/chain.proto

package chainv1

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

type StorageType int32

const (
	StorageType_STORAGE_TYPE_UNSPECIFIED            StorageType = 0
	StorageType_STORAGE_TYPE_STANDARD_STORE         StorageType = 1
	StorageType_STORAGE_TYPE_METADATA_ENCODED_STORE StorageType = 2
)

// Enum value maps for StorageType.
var (
	StorageType_name = map[int32]string{
		0: "STORAGE_TYPE_UNSPECIFIED",
		1: "STORAGE_TYPE_STANDARD_STORE",
		2: "STORAGE_TYPE_METADATA_ENCODED_STORE",
	}
	StorageType_value = map[string]int32{
		"STORAGE_TYPE_UNSPECIFIED":            0,
		"STORAGE_TYPE_STANDARD_STORE":         1,
		"STORAGE_TYPE_METADATA_ENCODED_STORE": 2,
	}
)

func (x StorageType) Enum() *StorageType {
	p := new(StorageType)
	*p = x
	return p
}

func (x StorageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorageType) Descriptor() protoreflect.EnumDescriptor {
	return file_chain_v1_chain_proto_enumTypes[0].Descriptor()
}

func (StorageType) Type() protoreflect.EnumType {
	return &file_chain_v1_chain_proto_enumTypes[0]
}

func (x StorageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorageType.Descriptor instead.
func (StorageType) EnumDescriptor() ([]byte, []int) {
	return file_chain_v1_chain_proto_rawDescGZIP(), []int{0}
}

type IndexEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IndexEntry) Reset() {
	*x = IndexEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_v1_chain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexEntry) ProtoMessage() {}

func (x *IndexEntry) ProtoReflect() protoreflect.Message {
	mi := &file_chain_v1_chain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexEntry.ProtoReflect.Descriptor instead.
func (*IndexEntry) Descriptor() ([]byte, []int) {
	return file_chain_v1_chain_proto_rawDescGZIP(), []int{0}
}

func (x *IndexEntry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *IndexEntry) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Storage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         StorageType            `protobuf:"varint,1,opt,name=type,proto3,enum=chain.v1.StorageType" json:"type,omitempty"`
	ReverseIndex map[string]*IndexEntry `protobuf:"bytes,2,rep,name=reverse_index,json=reverseIndex,proto3" json:"reverse_index,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Storage) Reset() {
	*x = Storage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_v1_chain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Storage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Storage) ProtoMessage() {}

func (x *Storage) ProtoReflect() protoreflect.Message {
	mi := &file_chain_v1_chain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Storage.ProtoReflect.Descriptor instead.
func (*Storage) Descriptor() ([]byte, []int) {
	return file_chain_v1_chain_proto_rawDescGZIP(), []int{1}
}

func (x *Storage) GetType() StorageType {
	if x != nil {
		return x.Type
	}
	return StorageType_STORAGE_TYPE_UNSPECIFIED
}

func (x *Storage) GetReverseIndex() map[string]*IndexEntry {
	if x != nil {
		return x.ReverseIndex
	}
	return nil
}

var File_chain_v1_chain_proto protoreflect.FileDescriptor

var file_chain_v1_chain_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x22, 0x34, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xd5, 0x01, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x15, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x48, 0x0a,
	0x0d, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x72, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x55, 0x0a, 0x11, 0x52, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x75,
	0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a,
	0x18, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x53,
	0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x4e,
	0x44, 0x41, 0x52, 0x44, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x10, 0x01, 0x12, 0x27, 0x0a, 0x23,
	0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x54,
	0x41, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x45, 0x4e, 0x43, 0x4f, 0x44, 0x45, 0x44, 0x5f, 0x53, 0x54,
	0x4f, 0x52, 0x45, 0x10, 0x02, 0x32, 0x10, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x70, 0x68, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chain_v1_chain_proto_rawDescOnce sync.Once
	file_chain_v1_chain_proto_rawDescData = file_chain_v1_chain_proto_rawDesc
)

func file_chain_v1_chain_proto_rawDescGZIP() []byte {
	file_chain_v1_chain_proto_rawDescOnce.Do(func() {
		file_chain_v1_chain_proto_rawDescData = protoimpl.X.CompressGZIP(file_chain_v1_chain_proto_rawDescData)
	})
	return file_chain_v1_chain_proto_rawDescData
}

var file_chain_v1_chain_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chain_v1_chain_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chain_v1_chain_proto_goTypes = []interface{}{
	(StorageType)(0),   // 0: chain.v1.StorageType
	(*IndexEntry)(nil), // 1: chain.v1.IndexEntry
	(*Storage)(nil),    // 2: chain.v1.Storage
	nil,                // 3: chain.v1.Storage.ReverseIndexEntry
}
var file_chain_v1_chain_proto_depIdxs = []int32{
	0, // 0: chain.v1.Storage.type:type_name -> chain.v1.StorageType
	3, // 1: chain.v1.Storage.reverse_index:type_name -> chain.v1.Storage.ReverseIndexEntry
	1, // 2: chain.v1.Storage.ReverseIndexEntry.value:type_name -> chain.v1.IndexEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_chain_v1_chain_proto_init() }
func file_chain_v1_chain_proto_init() {
	if File_chain_v1_chain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chain_v1_chain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexEntry); i {
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
		file_chain_v1_chain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Storage); i {
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
			RawDescriptor: file_chain_v1_chain_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chain_v1_chain_proto_goTypes,
		DependencyIndexes: file_chain_v1_chain_proto_depIdxs,
		EnumInfos:         file_chain_v1_chain_proto_enumTypes,
		MessageInfos:      file_chain_v1_chain_proto_msgTypes,
	}.Build()
	File_chain_v1_chain_proto = out.File
	file_chain_v1_chain_proto_rawDesc = nil
	file_chain_v1_chain_proto_goTypes = nil
	file_chain_v1_chain_proto_depIdxs = nil
}