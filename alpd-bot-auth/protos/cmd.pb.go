// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: cmd.proto

package protos

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CmdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Context *RequestContext `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	Name    string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Args    []string        `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *CmdRequest) Reset() {
	*x = CmdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdRequest) ProtoMessage() {}

func (x *CmdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdRequest.ProtoReflect.Descriptor instead.
func (*CmdRequest) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{0}
}

func (x *CmdRequest) GetContext() *RequestContext {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *CmdRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CmdRequest) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

type RequestContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User       string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	RemoteAddr string `protobuf:"bytes,2,opt,name=remote_addr,json=remoteAddr,proto3" json:"remote_addr,omitempty"`
}

func (x *RequestContext) Reset() {
	*x = RequestContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestContext) ProtoMessage() {}

func (x *RequestContext) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestContext.ProtoReflect.Descriptor instead.
func (*RequestContext) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{1}
}

func (x *RequestContext) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *RequestContext) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

type CmdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CmdReply) Reset() {
	*x = CmdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdReply) ProtoMessage() {}

func (x *CmdReply) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdReply.ProtoReflect.Descriptor instead.
func (*CmdReply) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{2}
}

func (x *CmdReply) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_cmd_proto protoreflect.FileDescriptor

var file_cmd_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x6c, 0x70,
	0x64, 0x22, 0x64, 0x0a, 0x0a, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x61, 0x6c, 0x70, 0x64, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x45, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x22, 0x1e,
	0x0a, 0x08, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x48,
	0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x36, 0x0a, 0x10, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x56, 0x31, 0x12, 0x10, 0x2e, 0x61, 0x6c, 0x70, 0x64, 0x2e, 0x43, 0x6d, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x6c, 0x70, 0x64, 0x2e, 0x43, 0x6d,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cmd_proto_rawDescOnce sync.Once
	file_cmd_proto_rawDescData = file_cmd_proto_rawDesc
)

func file_cmd_proto_rawDescGZIP() []byte {
	file_cmd_proto_rawDescOnce.Do(func() {
		file_cmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_proto_rawDescData)
	})
	return file_cmd_proto_rawDescData
}

var file_cmd_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cmd_proto_goTypes = []interface{}{
	(*CmdRequest)(nil),     // 0: alpd.CmdRequest
	(*RequestContext)(nil), // 1: alpd.RequestContext
	(*CmdReply)(nil),       // 2: alpd.CmdReply
}
var file_cmd_proto_depIdxs = []int32{
	1, // 0: alpd.CmdRequest.context:type_name -> alpd.RequestContext
	0, // 1: alpd.CommandService.ExecuteCommandV1:input_type -> alpd.CmdRequest
	2, // 2: alpd.CommandService.ExecuteCommandV1:output_type -> alpd.CmdReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cmd_proto_init() }
func file_cmd_proto_init() {
	if File_cmd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdRequest); i {
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
		file_cmd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestContext); i {
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
		file_cmd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdReply); i {
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
			RawDescriptor: file_cmd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmd_proto_goTypes,
		DependencyIndexes: file_cmd_proto_depIdxs,
		MessageInfos:      file_cmd_proto_msgTypes,
	}.Build()
	File_cmd_proto = out.File
	file_cmd_proto_rawDesc = nil
	file_cmd_proto_goTypes = nil
	file_cmd_proto_depIdxs = nil
}