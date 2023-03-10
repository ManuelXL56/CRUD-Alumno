// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: Config_Token.proto

package grpc_token

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

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Config_Token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_Config_Token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_Config_Token_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type DataTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey string `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	Data      string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DataTokenRequest) Reset() {
	*x = DataTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Config_Token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataTokenRequest) ProtoMessage() {}

func (x *DataTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Config_Token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataTokenRequest.ProtoReflect.Descriptor instead.
func (*DataTokenRequest) Descriptor() ([]byte, []int) {
	return file_Config_Token_proto_rawDescGZIP(), []int{1}
}

func (x *DataTokenRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *DataTokenRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type DataTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Matricule string `protobuf:"bytes,1,opt,name=matricule,proto3" json:"matricule,omitempty"`
	Result    string `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DataTokenReply) Reset() {
	*x = DataTokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Config_Token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataTokenReply) ProtoMessage() {}

func (x *DataTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_Config_Token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataTokenReply.ProtoReflect.Descriptor instead.
func (*DataTokenReply) Descriptor() ([]byte, []int) {
	return file_Config_Token_proto_rawDescGZIP(), []int{2}
}

func (x *DataTokenReply) GetMatricule() string {
	if x != nil {
		return x.Matricule
	}
	return ""
}

func (x *DataTokenReply) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_Config_Token_proto protoreflect.FileDescriptor

var file_Config_Token_proto_rawDesc = []byte{
	0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x1a, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x44, 0x0a, 0x10,
	0x44, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x46, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x63, 0x75, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x63, 0x75,
	0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x9a, 0x01, 0x0a, 0x0a, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0d, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x57, 0x54, 0x2d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_Config_Token_proto_rawDescOnce sync.Once
	file_Config_Token_proto_rawDescData = file_Config_Token_proto_rawDesc
)

func file_Config_Token_proto_rawDescGZIP() []byte {
	file_Config_Token_proto_rawDescOnce.Do(func() {
		file_Config_Token_proto_rawDescData = protoimpl.X.CompressGZIP(file_Config_Token_proto_rawDescData)
	})
	return file_Config_Token_proto_rawDescData
}

var file_Config_Token_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_Config_Token_proto_goTypes = []interface{}{
	(*Data)(nil),             // 0: grpc_token.Data
	(*DataTokenRequest)(nil), // 1: grpc_token.DataTokenRequest
	(*DataTokenReply)(nil),   // 2: grpc_token.DataTokenReply
}
var file_Config_Token_proto_depIdxs = []int32{
	1, // 0: grpc_token.RouteGuide.CreateToken:input_type -> grpc_token.DataTokenRequest
	1, // 1: grpc_token.RouteGuide.ValidateToken:input_type -> grpc_token.DataTokenRequest
	0, // 2: grpc_token.RouteGuide.CreateToken:output_type -> grpc_token.Data
	2, // 3: grpc_token.RouteGuide.ValidateToken:output_type -> grpc_token.DataTokenReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Config_Token_proto_init() }
func file_Config_Token_proto_init() {
	if File_Config_Token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Config_Token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_Config_Token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataTokenRequest); i {
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
		file_Config_Token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataTokenReply); i {
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
			RawDescriptor: file_Config_Token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Config_Token_proto_goTypes,
		DependencyIndexes: file_Config_Token_proto_depIdxs,
		MessageInfos:      file_Config_Token_proto_msgTypes,
	}.Build()
	File_Config_Token_proto = out.File
	file_Config_Token_proto_rawDesc = nil
	file_Config_Token_proto_goTypes = nil
	file_Config_Token_proto_depIdxs = nil
}
