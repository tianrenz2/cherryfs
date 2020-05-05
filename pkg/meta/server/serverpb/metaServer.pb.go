// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        (unknown)
// source: metaServer.proto

package serverpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AskPutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Size       *int64  `protobuf:"varint,2,req,name=size" json:"size,omitempty"`
	Object     *string `protobuf:"bytes,3,req,name=object" json:"object,omitempty"`
	ObjectHash *string `protobuf:"bytes,4,req,name=objectHash" json:"objectHash,omitempty"`
}

func (x *AskPutRequest) Reset() {
	*x = AskPutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metaServer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskPutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskPutRequest) ProtoMessage() {}

func (x *AskPutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metaServer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskPutRequest.ProtoReflect.Descriptor instead.
func (*AskPutRequest) Descriptor() ([]byte, []int) {
	return file_metaServer_proto_rawDescGZIP(), []int{0}
}

func (x *AskPutRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *AskPutRequest) GetSize() int64 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

func (x *AskPutRequest) GetObject() string {
	if x != nil && x.Object != nil {
		return *x.Object
	}
	return ""
}

func (x *AskPutRequest) GetObjectHash() string {
	if x != nil && x.ObjectHash != nil {
		return *x.ObjectHash
	}
	return ""
}

type AskPutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   *int32  `protobuf:"varint,1,req,name=status" json:"status,omitempty"`
	DestAddr *string `protobuf:"bytes,2,req,name=destAddr" json:"destAddr,omitempty"`
	DestDir  *string `protobuf:"bytes,3,req,name=destDir" json:"destDir,omitempty"`
	Msg      *string `protobuf:"bytes,4,opt,name=msg" json:"msg,omitempty"`
}

func (x *AskPutResponse) Reset() {
	*x = AskPutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metaServer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskPutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskPutResponse) ProtoMessage() {}

func (x *AskPutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metaServer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskPutResponse.ProtoReflect.Descriptor instead.
func (*AskPutResponse) Descriptor() ([]byte, []int) {
	return file_metaServer_proto_rawDescGZIP(), []int{1}
}

func (x *AskPutResponse) GetStatus() int32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *AskPutResponse) GetDestAddr() string {
	if x != nil && x.DestAddr != nil {
		return *x.DestAddr
	}
	return ""
}

func (x *AskPutResponse) GetDestDir() string {
	if x != nil && x.DestDir != nil {
		return *x.DestDir
	}
	return ""
}

func (x *AskPutResponse) GetMsg() string {
	if x != nil && x.Msg != nil {
		return *x.Msg
	}
	return ""
}

var File_metaServer_proto protoreflect.FileDescriptor

var file_metaServer_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x65, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x62,
	0x22, 0x4f, 0x0a, 0x0d, 0x41, 0x73, 0x6b, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x12,
	0x0c, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x03, 0x12, 0x0e, 0x0a,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x12, 0x12, 0x0a,
	0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x02, 0x28,
	0x09, 0x22, 0x50, 0x0a, 0x0e, 0x41, 0x73, 0x6b, 0x50, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x05, 0x12, 0x10, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x09, 0x12, 0x0f, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x44, 0x69, 0x72,
	0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x12, 0x0b, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x32, 0x4f, 0x0a, 0x06, 0x41, 0x73, 0x6b, 0x50, 0x75, 0x74, 0x12, 0x45, 0x0a,
	0x06, 0x41, 0x73, 0x6b, 0x50, 0x75, 0x74, 0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x41, 0x73, 0x6b, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x41, 0x73, 0x6b, 0x50, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x70, 0x62,
}

var (
	file_metaServer_proto_rawDescOnce sync.Once
	file_metaServer_proto_rawDescData = file_metaServer_proto_rawDesc
)

func file_metaServer_proto_rawDescGZIP() []byte {
	file_metaServer_proto_rawDescOnce.Do(func() {
		file_metaServer_proto_rawDescData = protoimpl.X.CompressGZIP(file_metaServer_proto_rawDescData)
	})
	return file_metaServer_proto_rawDescData
}

var file_metaServer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_metaServer_proto_goTypes = []interface{}{
	(*AskPutRequest)(nil),  // 0: metaserverpb.AskPutRequest
	(*AskPutResponse)(nil), // 1: metaserverpb.AskPutResponse
}
var file_metaServer_proto_depIdxs = []int32{
	0, // 0: metaserverpb.AskPut.AskPut:input_type -> metaserverpb.AskPutRequest
	1, // 1: metaserverpb.AskPut.AskPut:output_type -> metaserverpb.AskPutResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_metaServer_proto_init() }
func file_metaServer_proto_init() {
	if File_metaServer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metaServer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskPutRequest); i {
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
		file_metaServer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskPutResponse); i {
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
			RawDescriptor: file_metaServer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metaServer_proto_goTypes,
		DependencyIndexes: file_metaServer_proto_depIdxs,
		MessageInfos:      file_metaServer_proto_msgTypes,
	}.Build()
	File_metaServer_proto = out.File
	file_metaServer_proto_rawDesc = nil
	file_metaServer_proto_goTypes = nil
	file_metaServer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AskPutClient is the client API for AskPut service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AskPutClient interface {
	AskPut(ctx context.Context, in *AskPutRequest, opts ...grpc.CallOption) (*AskPutResponse, error)
}

type askPutClient struct {
	cc grpc.ClientConnInterface
}

func NewAskPutClient(cc grpc.ClientConnInterface) AskPutClient {
	return &askPutClient{cc}
}

func (c *askPutClient) AskPut(ctx context.Context, in *AskPutRequest, opts ...grpc.CallOption) (*AskPutResponse, error) {
	out := new(AskPutResponse)
	err := c.cc.Invoke(ctx, "/metaserverpb.AskPut/AskPut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AskPutServer is the server API for AskPut service.
type AskPutServer interface {
	AskPut(context.Context, *AskPutRequest) (*AskPutResponse, error)
}

// UnimplementedAskPutServer can be embedded to have forward compatible implementations.
type UnimplementedAskPutServer struct {
}

func (*UnimplementedAskPutServer) AskPut(context.Context, *AskPutRequest) (*AskPutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskPut not implemented")
}

func RegisterAskPutServer(s *grpc.Server, srv AskPutServer) {
	s.RegisterService(&_AskPut_serviceDesc, srv)
}

func _AskPut_AskPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AskPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AskPutServer).AskPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metaserverpb.AskPut/AskPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AskPutServer).AskPut(ctx, req.(*AskPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AskPut_serviceDesc = grpc.ServiceDesc{
	ServiceName: "metaserverpb.AskPut",
	HandlerType: (*AskPutServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AskPut",
			Handler:    _AskPut_AskPut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metaServer.proto",
}