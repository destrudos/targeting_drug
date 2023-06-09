// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// source: determined/api/v1/resourcepool.proto

package apiv1

import (
	resourcepoolv1 "github.com/determined-ai/determined/proto/pkg/resourcepoolv1"
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

// Get the list of resource pools from the cluster.
type GetResourcePoolsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Skip the number of resource pools before returning results. Negative values
	// denote number of resource pools to skip from the end before returning
	// results.
	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	// Limit the number of resource pools. A value of 0 denotes no limit.
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetResourcePoolsRequest) Reset() {
	*x = GetResourcePoolsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_resourcepool_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourcePoolsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcePoolsRequest) ProtoMessage() {}

func (x *GetResourcePoolsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_resourcepool_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcePoolsRequest.ProtoReflect.Descriptor instead.
func (*GetResourcePoolsRequest) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_resourcepool_proto_rawDescGZIP(), []int{0}
}

func (x *GetResourcePoolsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetResourcePoolsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

// Response to GetResourcePoolsRequest.
type GetResourcePoolsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of returned resource pools.
	ResourcePools []*resourcepoolv1.ResourcePool `protobuf:"bytes,1,rep,name=resource_pools,json=resourcePools,proto3" json:"resource_pools,omitempty"`
	// Pagination information of the full dataset.
	Pagination *Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetResourcePoolsResponse) Reset() {
	*x = GetResourcePoolsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_api_v1_resourcepool_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourcePoolsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcePoolsResponse) ProtoMessage() {}

func (x *GetResourcePoolsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_determined_api_v1_resourcepool_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcePoolsResponse.ProtoReflect.Descriptor instead.
func (*GetResourcePoolsResponse) Descriptor() ([]byte, []int) {
	return file_determined_api_v1_resourcepool_proto_rawDescGZIP(), []int{1}
}

func (x *GetResourcePoolsResponse) GetResourcePools() []*resourcepoolv1.ResourcePool {
	if x != nil {
		return x.ResourcePools
	}
	return nil
}

func (x *GetResourcePoolsResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_determined_api_v1_resourcepool_proto protoreflect.FileDescriptor

var file_determined_api_v1_resourcepool_proto_rawDesc = []byte{
	0x0a, 0x24, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x70, 0x6f, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x64, 0x65, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x64,
	0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xaa, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70,
	0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x65, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f,
	0x6f, 0x6c, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2d, 0x61, 0x69, 0x2f, 0x64,
	0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_determined_api_v1_resourcepool_proto_rawDescOnce sync.Once
	file_determined_api_v1_resourcepool_proto_rawDescData = file_determined_api_v1_resourcepool_proto_rawDesc
)

func file_determined_api_v1_resourcepool_proto_rawDescGZIP() []byte {
	file_determined_api_v1_resourcepool_proto_rawDescOnce.Do(func() {
		file_determined_api_v1_resourcepool_proto_rawDescData = protoimpl.X.CompressGZIP(file_determined_api_v1_resourcepool_proto_rawDescData)
	})
	return file_determined_api_v1_resourcepool_proto_rawDescData
}

var file_determined_api_v1_resourcepool_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_determined_api_v1_resourcepool_proto_goTypes = []interface{}{
	(*GetResourcePoolsRequest)(nil),     // 0: determined.api.v1.GetResourcePoolsRequest
	(*GetResourcePoolsResponse)(nil),    // 1: determined.api.v1.GetResourcePoolsResponse
	(*resourcepoolv1.ResourcePool)(nil), // 2: determined.resourcepool.v1.ResourcePool
	(*Pagination)(nil),                  // 3: determined.api.v1.Pagination
}
var file_determined_api_v1_resourcepool_proto_depIdxs = []int32{
	2, // 0: determined.api.v1.GetResourcePoolsResponse.resource_pools:type_name -> determined.resourcepool.v1.ResourcePool
	3, // 1: determined.api.v1.GetResourcePoolsResponse.pagination:type_name -> determined.api.v1.Pagination
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_determined_api_v1_resourcepool_proto_init() }
func file_determined_api_v1_resourcepool_proto_init() {
	if File_determined_api_v1_resourcepool_proto != nil {
		return
	}
	file_determined_api_v1_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_determined_api_v1_resourcepool_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourcePoolsRequest); i {
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
		file_determined_api_v1_resourcepool_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourcePoolsResponse); i {
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
			RawDescriptor: file_determined_api_v1_resourcepool_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_determined_api_v1_resourcepool_proto_goTypes,
		DependencyIndexes: file_determined_api_v1_resourcepool_proto_depIdxs,
		MessageInfos:      file_determined_api_v1_resourcepool_proto_msgTypes,
	}.Build()
	File_determined_api_v1_resourcepool_proto = out.File
	file_determined_api_v1_resourcepool_proto_rawDesc = nil
	file_determined_api_v1_resourcepool_proto_goTypes = nil
	file_determined_api_v1_resourcepool_proto_depIdxs = nil
}
