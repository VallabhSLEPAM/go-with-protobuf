// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: my-protobuf/proto/car/car.proto

package car

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId          string `protobuf:"bytes,1,opt,name=car_id,proto3" json:"car_id,omitempty"`
	Brand          string `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Model          string `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	Price          uint32 `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	ManufatureYear uint32 `protobuf:"varint,5,opt,name=manufature_year,json=manufatureYear,proto3" json:"manufature_year,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	mi := &file_my_protobuf_proto_car_car_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_my_protobuf_proto_car_car_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_my_protobuf_proto_car_car_proto_rawDescGZIP(), []int{0}
}

func (x *Car) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

func (x *Car) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Car) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Car) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Car) GetManufatureYear() uint32 {
	if x != nil {
		return x.ManufatureYear
	}
	return 0
}

var File_my_protobuf_proto_car_car_proto protoreflect.FileDescriptor

var file_my_protobuf_proto_car_car_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x2f, 0x63, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x29, 0x6d, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a,
	0x03, 0x43, 0x61, 0x72, 0x12, 0x20, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x06,
	0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xfa, 0x42, 0x1f, 0x72, 0x1d, 0x32, 0x1b, 0x28, 0x3f,
	0x69, 0x29, 0x5e, 0x54, 0x6f, 0x79, 0x6f, 0x74, 0x6f, 0x7c, 0x48, 0x6f, 0x6e, 0x64, 0x61, 0x7c,
	0x46, 0x6f, 0x72, 0x64, 0x7c, 0x42, 0x4d, 0x57, 0x24, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x12, 0x1d, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x1e, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x1d, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x34,
	0x0a, 0x0f, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x79, 0x65, 0x61,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x2a, 0x06, 0x18, 0xe8,
	0x0f, 0x28, 0xe4, 0x0f, 0x52, 0x0e, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x59, 0x65, 0x61, 0x72, 0x42, 0x1a, 0x5a, 0x18, 0x6d, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x61, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_my_protobuf_proto_car_car_proto_rawDescOnce sync.Once
	file_my_protobuf_proto_car_car_proto_rawDescData = file_my_protobuf_proto_car_car_proto_rawDesc
)

func file_my_protobuf_proto_car_car_proto_rawDescGZIP() []byte {
	file_my_protobuf_proto_car_car_proto_rawDescOnce.Do(func() {
		file_my_protobuf_proto_car_car_proto_rawDescData = protoimpl.X.CompressGZIP(file_my_protobuf_proto_car_car_proto_rawDescData)
	})
	return file_my_protobuf_proto_car_car_proto_rawDescData
}

var file_my_protobuf_proto_car_car_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_my_protobuf_proto_car_car_proto_goTypes = []any{
	(*Car)(nil), // 0: Car
}
var file_my_protobuf_proto_car_car_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_my_protobuf_proto_car_car_proto_init() }
func file_my_protobuf_proto_car_car_proto_init() {
	if File_my_protobuf_proto_car_car_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_my_protobuf_proto_car_car_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_my_protobuf_proto_car_car_proto_goTypes,
		DependencyIndexes: file_my_protobuf_proto_car_car_proto_depIdxs,
		MessageInfos:      file_my_protobuf_proto_car_car_proto_msgTypes,
	}.Build()
	File_my_protobuf_proto_car_car_proto = out.File
	file_my_protobuf_proto_car_car_proto_rawDesc = nil
	file_my_protobuf_proto_car_car_proto_goTypes = nil
	file_my_protobuf_proto_car_car_proto_depIdxs = nil
}
