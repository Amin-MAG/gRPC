// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.5
// source: apis/majority.proto

package apis

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

// The request message containing the user's name.
type ListMajorityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Numbers []int32 `protobuf:"varint,1,rep,packed,name=numbers,proto3" json:"numbers,omitempty"`
}

func (x *ListMajorityRequest) Reset() {
	*x = ListMajorityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_majority_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMajorityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMajorityRequest) ProtoMessage() {}

func (x *ListMajorityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_majority_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMajorityRequest.ProtoReflect.Descriptor instead.
func (*ListMajorityRequest) Descriptor() ([]byte, []int) {
	return file_apis_majority_proto_rawDescGZIP(), []int{0}
}

func (x *ListMajorityRequest) GetNumbers() []int32 {
	if x != nil {
		return x.Numbers
	}
	return nil
}

// The response message containing the greetings
type ListMajorityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message        string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	MajorityNumber int32  `protobuf:"varint,2,opt,name=majorityNumber,proto3" json:"majorityNumber,omitempty"`
}

func (x *ListMajorityReply) Reset() {
	*x = ListMajorityReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_majority_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMajorityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMajorityReply) ProtoMessage() {}

func (x *ListMajorityReply) ProtoReflect() protoreflect.Message {
	mi := &file_apis_majority_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMajorityReply.ProtoReflect.Descriptor instead.
func (*ListMajorityReply) Descriptor() ([]byte, []int) {
	return file_apis_majority_proto_rawDescGZIP(), []int{1}
}

func (x *ListMajorityReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListMajorityReply) GetMajorityNumber() int32 {
	if x != nil {
		return x.MajorityNumber
	}
	return 0
}

var File_apis_majority_proto protoreflect.FileDescriptor

var file_apis_majority_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x61, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x69, 0x65, 0x77, 0x22, 0x2f, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6a,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x55, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61,
	0x6a, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6d,
	0x61, 0x6a, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x6d, 0x0a,
	0x08, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x61, 0x0a, 0x15, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x23, 0x2e, 0x62, 0x61, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x69, 0x65, 0x77, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x61, 0x6c, 0x65, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6a,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13,
	0x62, 0x61, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apis_majority_proto_rawDescOnce sync.Once
	file_apis_majority_proto_rawDescData = file_apis_majority_proto_rawDesc
)

func file_apis_majority_proto_rawDescGZIP() []byte {
	file_apis_majority_proto_rawDescOnce.Do(func() {
		file_apis_majority_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_majority_proto_rawDescData)
	})
	return file_apis_majority_proto_rawDescData
}

var file_apis_majority_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apis_majority_proto_goTypes = []interface{}{
	(*ListMajorityRequest)(nil), // 0: bale_interview.ListMajorityRequest
	(*ListMajorityReply)(nil),   // 1: bale_interview.ListMajorityReply
}
var file_apis_majority_proto_depIdxs = []int32{
	0, // 0: bale_interview.Majority.CalculateListMajority:input_type -> bale_interview.ListMajorityRequest
	1, // 1: bale_interview.Majority.CalculateListMajority:output_type -> bale_interview.ListMajorityReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apis_majority_proto_init() }
func file_apis_majority_proto_init() {
	if File_apis_majority_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apis_majority_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMajorityRequest); i {
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
		file_apis_majority_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMajorityReply); i {
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
			RawDescriptor: file_apis_majority_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apis_majority_proto_goTypes,
		DependencyIndexes: file_apis_majority_proto_depIdxs,
		MessageInfos:      file_apis_majority_proto_msgTypes,
	}.Build()
	File_apis_majority_proto = out.File
	file_apis_majority_proto_rawDesc = nil
	file_apis_majority_proto_goTypes = nil
	file_apis_majority_proto_depIdxs = nil
}
