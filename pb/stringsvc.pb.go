// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: pb/stringsvc.proto

package pb

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

// The uppercase request contains a parameter.
type UppercaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S string `protobuf:"bytes,1,opt,name=S,proto3" json:"S,omitempty"`
}

func (x *UppercaseRequest) Reset() {
	*x = UppercaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_stringsvc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UppercaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UppercaseRequest) ProtoMessage() {}

func (x *UppercaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_stringsvc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UppercaseRequest.ProtoReflect.Descriptor instead.
func (*UppercaseRequest) Descriptor() ([]byte, []int) {
	return file_pb_stringsvc_proto_rawDescGZIP(), []int{0}
}

func (x *UppercaseRequest) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

// The uppercase response contains two parameters.
type UppercaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V   string `protobuf:"bytes,1,opt,name=V,proto3" json:"V,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=Err,proto3" json:"Err,omitempty"`
}

func (x *UppercaseResponse) Reset() {
	*x = UppercaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_stringsvc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UppercaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UppercaseResponse) ProtoMessage() {}

func (x *UppercaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_stringsvc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UppercaseResponse.ProtoReflect.Descriptor instead.
func (*UppercaseResponse) Descriptor() ([]byte, []int) {
	return file_pb_stringsvc_proto_rawDescGZIP(), []int{1}
}

func (x *UppercaseResponse) GetV() string {
	if x != nil {
		return x.V
	}
	return ""
}

func (x *UppercaseResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

// The count request contains a parameter.
type CountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S string `protobuf:"bytes,1,opt,name=S,proto3" json:"S,omitempty"`
}

func (x *CountRequest) Reset() {
	*x = CountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_stringsvc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountRequest) ProtoMessage() {}

func (x *CountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_stringsvc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountRequest.ProtoReflect.Descriptor instead.
func (*CountRequest) Descriptor() ([]byte, []int) {
	return file_pb_stringsvc_proto_rawDescGZIP(), []int{2}
}

func (x *CountRequest) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

// The count response contains a parameters.
type CountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V int32 `protobuf:"varint,1,opt,name=V,proto3" json:"V,omitempty"`
}

func (x *CountResponse) Reset() {
	*x = CountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_stringsvc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountResponse) ProtoMessage() {}

func (x *CountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_stringsvc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountResponse.ProtoReflect.Descriptor instead.
func (*CountResponse) Descriptor() ([]byte, []int) {
	return file_pb_stringsvc_proto_rawDescGZIP(), []int{3}
}

func (x *CountResponse) GetV() int32 {
	if x != nil {
		return x.V
	}
	return 0
}

var File_pb_stringsvc_proto protoreflect.FileDescriptor

var file_pb_stringsvc_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x76, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x20, 0x0a, 0x10, 0x55, 0x70, 0x70, 0x65,
	0x72, 0x63, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01,
	0x53, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x53, 0x22, 0x33, 0x0a, 0x11, 0x55, 0x70,
	0x70, 0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0c, 0x0a, 0x01, 0x56, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x56, 0x12, 0x10, 0x0a,
	0x03, 0x45, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x72, 0x72, 0x22,
	0x1c, 0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0c, 0x0a, 0x01, 0x53, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x53, 0x22, 0x1d, 0x0a,
	0x0d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c,
	0x0a, 0x01, 0x56, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x56, 0x32, 0x77, 0x0a, 0x09,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x76, 0x63, 0x12, 0x3a, 0x0a, 0x09, 0x55, 0x70, 0x70,
	0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x70, 0x65,
	0x72, 0x63, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x70, 0x70, 0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_stringsvc_proto_rawDescOnce sync.Once
	file_pb_stringsvc_proto_rawDescData = file_pb_stringsvc_proto_rawDesc
)

func file_pb_stringsvc_proto_rawDescGZIP() []byte {
	file_pb_stringsvc_proto_rawDescOnce.Do(func() {
		file_pb_stringsvc_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_stringsvc_proto_rawDescData)
	})
	return file_pb_stringsvc_proto_rawDescData
}

var file_pb_stringsvc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_stringsvc_proto_goTypes = []interface{}{
	(*UppercaseRequest)(nil),  // 0: pb.UppercaseRequest
	(*UppercaseResponse)(nil), // 1: pb.UppercaseResponse
	(*CountRequest)(nil),      // 2: pb.CountRequest
	(*CountResponse)(nil),     // 3: pb.CountResponse
}
var file_pb_stringsvc_proto_depIdxs = []int32{
	0, // 0: pb.StringSvc.Uppercase:input_type -> pb.UppercaseRequest
	2, // 1: pb.StringSvc.Count:input_type -> pb.CountRequest
	1, // 2: pb.StringSvc.Uppercase:output_type -> pb.UppercaseResponse
	3, // 3: pb.StringSvc.Count:output_type -> pb.CountResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_stringsvc_proto_init() }
func file_pb_stringsvc_proto_init() {
	if File_pb_stringsvc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_stringsvc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UppercaseRequest); i {
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
		file_pb_stringsvc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UppercaseResponse); i {
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
		file_pb_stringsvc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountRequest); i {
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
		file_pb_stringsvc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountResponse); i {
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
			RawDescriptor: file_pb_stringsvc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_stringsvc_proto_goTypes,
		DependencyIndexes: file_pb_stringsvc_proto_depIdxs,
		MessageInfos:      file_pb_stringsvc_proto_msgTypes,
	}.Build()
	File_pb_stringsvc_proto = out.File
	file_pb_stringsvc_proto_rawDesc = nil
	file_pb_stringsvc_proto_goTypes = nil
	file_pb_stringsvc_proto_depIdxs = nil
}
