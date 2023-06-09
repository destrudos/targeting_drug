// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// source: determined/device/v1/device.proto

package devicev1

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

// The type of the Device.
type Type int32

const (
	// An unspecified device type.
	Type_TYPE_UNSPECIFIED Type = 0
	// A CPU device.
	Type_TYPE_CPU Type = 1
	// CUDA device.
	Type_TYPE_CUDA Type = 2
	// ROCM.
	Type_TYPE_ROCM Type = 3
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_CPU",
		2: "TYPE_CUDA",
		3: "TYPE_ROCM",
	}
	Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_CPU":         1,
		"TYPE_CUDA":        2,
		"TYPE_ROCM":        3,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_determined_device_v1_device_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_determined_device_v1_device_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_determined_device_v1_device_proto_rawDescGZIP(), []int{0}
}

// Device represents a single computational device on an agent.
type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The index of the device.
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The brand name of the device.
	Brand string `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	// The unique UUID of the device.
	Uuid string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// The type of the Device.
	Type Type `protobuf:"varint,4,opt,name=type,proto3,enum=determined.device.v1.Type" json:"type,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_device_v1_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_determined_device_v1_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_determined_device_v1_device_proto_rawDescGZIP(), []int{0}
}

func (x *Device) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Device) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Device) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Device) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_TYPE_UNSPECIFIED
}

var File_determined_device_v1_device_proto protoreflect.FileDescriptor

var file_determined_device_v1_device_proto_rawDesc = []byte{
	0x0a, 0x21, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x72, 0x0a, 0x06, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x2e, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x64, 0x65,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x48, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x50, 0x55, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x43, 0x55, 0x44, 0x41, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x52, 0x4f, 0x43, 0x4d, 0x10, 0x03, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64,
	0x2d, 0x61, 0x69, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_determined_device_v1_device_proto_rawDescOnce sync.Once
	file_determined_device_v1_device_proto_rawDescData = file_determined_device_v1_device_proto_rawDesc
)

func file_determined_device_v1_device_proto_rawDescGZIP() []byte {
	file_determined_device_v1_device_proto_rawDescOnce.Do(func() {
		file_determined_device_v1_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_determined_device_v1_device_proto_rawDescData)
	})
	return file_determined_device_v1_device_proto_rawDescData
}

var file_determined_device_v1_device_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_determined_device_v1_device_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_determined_device_v1_device_proto_goTypes = []interface{}{
	(Type)(0),      // 0: determined.device.v1.Type
	(*Device)(nil), // 1: determined.device.v1.Device
}
var file_determined_device_v1_device_proto_depIdxs = []int32{
	0, // 0: determined.device.v1.Device.type:type_name -> determined.device.v1.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_determined_device_v1_device_proto_init() }
func file_determined_device_v1_device_proto_init() {
	if File_determined_device_v1_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_determined_device_v1_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
			RawDescriptor: file_determined_device_v1_device_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_determined_device_v1_device_proto_goTypes,
		DependencyIndexes: file_determined_device_v1_device_proto_depIdxs,
		EnumInfos:         file_determined_device_v1_device_proto_enumTypes,
		MessageInfos:      file_determined_device_v1_device_proto_msgTypes,
	}.Build()
	File_determined_device_v1_device_proto = out.File
	file_determined_device_v1_device_proto_rawDesc = nil
	file_determined_device_v1_device_proto_goTypes = nil
	file_determined_device_v1_device_proto_depIdxs = nil
}
