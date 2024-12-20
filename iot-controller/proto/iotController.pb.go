// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.29.2
// source: iotController.proto

package iot_controller

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetPackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId        int32                  `protobuf:"varint,1,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	Timestamp       *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SomeUsefulField int32                  `protobuf:"varint,3,opt,name=someUsefulField,proto3" json:"someUsefulField,omitempty"`
}

func (x *GetPackageRequest) Reset() {
	*x = GetPackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iotController_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackageRequest) ProtoMessage() {}

func (x *GetPackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iotController_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackageRequest.ProtoReflect.Descriptor instead.
func (*GetPackageRequest) Descriptor() ([]byte, []int) {
	return file_iotController_proto_rawDescGZIP(), []int{0}
}

func (x *GetPackageRequest) GetDeviceId() int32 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *GetPackageRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *GetPackageRequest) GetSomeUsefulField() int32 {
	if x != nil {
		return x.SomeUsefulField
	}
	return 0
}

type PackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PackageResponse) Reset() {
	*x = PackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iotController_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageResponse) ProtoMessage() {}

func (x *PackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iotController_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageResponse.ProtoReflect.Descriptor instead.
func (*PackageResponse) Descriptor() ([]byte, []int) {
	return file_iotController_proto_rawDescGZIP(), []int{1}
}

func (x *PackageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_iotController_proto protoreflect.FileDescriptor

var file_iotController_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x6f, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x69, 0x6f, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x6f, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x66, 0x75, 0x6c, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x6f, 0x6d, 0x65,
	0x55, 0x73, 0x65, 0x66, 0x75, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x63, 0x0a, 0x0a, 0x49, 0x6f, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x57, 0x69, 0x74, 0x68, 0x41, 0x63, 0x6b, 0x12, 0x20, 0x2e, 0x69, 0x6f, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6f, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x12, 0x5a,
	0x10, 0x2e, 0x2f, 0x69, 0x6f, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iotController_proto_rawDescOnce sync.Once
	file_iotController_proto_rawDescData = file_iotController_proto_rawDesc
)

func file_iotController_proto_rawDescGZIP() []byte {
	file_iotController_proto_rawDescOnce.Do(func() {
		file_iotController_proto_rawDescData = protoimpl.X.CompressGZIP(file_iotController_proto_rawDescData)
	})
	return file_iotController_proto_rawDescData
}

var file_iotController_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_iotController_proto_goTypes = []any{
	(*GetPackageRequest)(nil),     // 0: iotController.GetPackageRequest
	(*PackageResponse)(nil),       // 1: iotController.PackageResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_iotController_proto_depIdxs = []int32{
	2, // 0: iotController.GetPackageRequest.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: iotController.IotService.StreamWithAck:input_type -> iotController.GetPackageRequest
	1, // 2: iotController.IotService.StreamWithAck:output_type -> iotController.PackageResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_iotController_proto_init() }
func file_iotController_proto_init() {
	if File_iotController_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iotController_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetPackageRequest); i {
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
		file_iotController_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PackageResponse); i {
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
			RawDescriptor: file_iotController_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iotController_proto_goTypes,
		DependencyIndexes: file_iotController_proto_depIdxs,
		MessageInfos:      file_iotController_proto_msgTypes,
	}.Build()
	File_iotController_proto = out.File
	file_iotController_proto_rawDesc = nil
	file_iotController_proto_goTypes = nil
	file_iotController_proto_depIdxs = nil
}