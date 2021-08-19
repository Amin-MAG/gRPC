// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.5
// source: api/echo.proto

package api

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

type CounterParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CounterParam) Reset() {
	*x = CounterParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CounterParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterParam) ProtoMessage() {}

func (x *CounterParam) ProtoReflect() protoreflect.Message {
	mi := &file_api_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterParam.ProtoReflect.Descriptor instead.
func (*CounterParam) Descriptor() ([]byte, []int) {
	return file_api_echo_proto_rawDescGZIP(), []int{0}
}

type RadioParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RadioParams) Reset() {
	*x = RadioParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadioParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadioParams) ProtoMessage() {}

func (x *RadioParams) ProtoReflect() protoreflect.Message {
	mi := &file_api_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadioParams.ProtoReflect.Descriptor instead.
func (*RadioParams) Descriptor() ([]byte, []int) {
	return file_api_echo_proto_rawDescGZIP(), []int{1}
}

type GlobalCounterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Counter int64 `protobuf:"varint,1,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (x *GlobalCounterResponse) Reset() {
	*x = GlobalCounterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_echo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalCounterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalCounterResponse) ProtoMessage() {}

func (x *GlobalCounterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_echo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalCounterResponse.ProtoReflect.Descriptor instead.
func (*GlobalCounterResponse) Descriptor() ([]byte, []int) {
	return file_api_echo_proto_rawDescGZIP(), []int{2}
}

func (x *GlobalCounterResponse) GetCounter() int64 {
	if x != nil {
		return x.Counter
	}
	return 0
}

type GlobalRadioResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkIndex int64  `protobuf:"varint,1,opt,name=chunkIndex,proto3" json:"chunkIndex,omitempty"`
	MusicBytes []byte `protobuf:"bytes,2,opt,name=musicBytes,proto3" json:"musicBytes,omitempty"`
}

func (x *GlobalRadioResponse) Reset() {
	*x = GlobalRadioResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_echo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalRadioResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalRadioResponse) ProtoMessage() {}

func (x *GlobalRadioResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_echo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalRadioResponse.ProtoReflect.Descriptor instead.
func (*GlobalRadioResponse) Descriptor() ([]byte, []int) {
	return file_api_echo_proto_rawDescGZIP(), []int{3}
}

func (x *GlobalRadioResponse) GetChunkIndex() int64 {
	if x != nil {
		return x.ChunkIndex
	}
	return 0
}

func (x *GlobalRadioResponse) GetMusicBytes() []byte {
	if x != nil {
		return x.MusicBytes
	}
	return nil
}

var File_api_echo_proto protoreflect.FileDescriptor

var file_api_echo_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x22, 0x0e, 0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x0d, 0x0a, 0x0b, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x31, 0x0a, 0x15, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x55, 0x0a, 0x13, 0x47, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1e, 0x0a, 0x0a, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x42, 0x79, 0x74, 0x65, 0x73, 0x32,
	0x8d, 0x01, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x44, 0x0a, 0x0d, 0x47, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x65, 0x63, 0x68, 0x6f,
	0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x1b, 0x2e,
	0x65, 0x63, 0x68, 0x6f, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3f,
	0x0a, 0x0b, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x12, 0x11, 0x2e,
	0x65, 0x63, 0x68, 0x6f, 0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x1a, 0x19, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x61,
	0x64, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x2f, 0x0a, 0x10, 0x69, 0x72, 0x2e, 0x6d, 0x61, 0x67, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65,
	0x63, 0x68, 0x6f, 0x42, 0x09, 0x45, 0x63, 0x68, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x08, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0xa2, 0x02, 0x03, 0x45, 0x43, 0x48,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_echo_proto_rawDescOnce sync.Once
	file_api_echo_proto_rawDescData = file_api_echo_proto_rawDesc
)

func file_api_echo_proto_rawDescGZIP() []byte {
	file_api_echo_proto_rawDescOnce.Do(func() {
		file_api_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_echo_proto_rawDescData)
	})
	return file_api_echo_proto_rawDescData
}

var file_api_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_echo_proto_goTypes = []interface{}{
	(*CounterParam)(nil),          // 0: echo.CounterParam
	(*RadioParams)(nil),           // 1: echo.RadioParams
	(*GlobalCounterResponse)(nil), // 2: echo.GlobalCounterResponse
	(*GlobalRadioResponse)(nil),   // 3: echo.GlobalRadioResponse
}
var file_api_echo_proto_depIdxs = []int32{
	0, // 0: echo.Echo.GlobalCounter:input_type -> echo.CounterParam
	1, // 1: echo.Echo.GlobalRadio:input_type -> echo.RadioParams
	2, // 2: echo.Echo.GlobalCounter:output_type -> echo.GlobalCounterResponse
	3, // 3: echo.Echo.GlobalRadio:output_type -> echo.GlobalRadioResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_echo_proto_init() }
func file_api_echo_proto_init() {
	if File_api_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CounterParam); i {
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
		file_api_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadioParams); i {
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
		file_api_echo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalCounterResponse); i {
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
		file_api_echo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalRadioResponse); i {
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
			RawDescriptor: file_api_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_echo_proto_goTypes,
		DependencyIndexes: file_api_echo_proto_depIdxs,
		MessageInfos:      file_api_echo_proto_msgTypes,
	}.Build()
	File_api_echo_proto = out.File
	file_api_echo_proto_rawDesc = nil
	file_api_echo_proto_goTypes = nil
	file_api_echo_proto_depIdxs = nil
}
