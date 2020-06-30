// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: chunkServer.proto

package pb

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

type ObjectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Targets []*Target `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"`
	Name    string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Hash    string    `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *ObjectInfo) Reset() {
	*x = ObjectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chunkServer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectInfo) ProtoMessage() {}

func (x *ObjectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_chunkServer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectInfo.ProtoReflect.Descriptor instead.
func (*ObjectInfo) Descriptor() ([]byte, []int) {
	return file_chunkServer_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectInfo) GetTargets() []*Target {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *ObjectInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ObjectInfo) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type PutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*PutRequest_Info
	//	*PutRequest_Content
	Data isPutRequest_Data `protobuf_oneof:"data"`
}

func (x *PutRequest) Reset() {
	*x = PutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chunkServer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRequest) ProtoMessage() {}

func (x *PutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chunkServer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRequest.ProtoReflect.Descriptor instead.
func (*PutRequest) Descriptor() ([]byte, []int) {
	return file_chunkServer_proto_rawDescGZIP(), []int{1}
}

func (m *PutRequest) GetData() isPutRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *PutRequest) GetInfo() *ObjectInfo {
	if x, ok := x.GetData().(*PutRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (x *PutRequest) GetContent() []byte {
	if x, ok := x.GetData().(*PutRequest_Content); ok {
		return x.Content
	}
	return nil
}

type isPutRequest_Data interface {
	isPutRequest_Data()
}

type PutRequest_Info struct {
	Info *ObjectInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type PutRequest_Content struct {
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3,oneof"`
}

func (*PutRequest_Info) isPutRequest_Data() {}

func (*PutRequest_Content) isPutRequest_Data() {}

type CopyObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*CopyObject_Info
	//	*CopyObject_Content
	Data isCopyObject_Data `protobuf_oneof:"data"`
}

func (x *CopyObject) Reset() {
	*x = CopyObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chunkServer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyObject) ProtoMessage() {}

func (x *CopyObject) ProtoReflect() protoreflect.Message {
	mi := &file_chunkServer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyObject.ProtoReflect.Descriptor instead.
func (*CopyObject) Descriptor() ([]byte, []int) {
	return file_chunkServer_proto_rawDescGZIP(), []int{2}
}

func (m *CopyObject) GetData() isCopyObject_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *CopyObject) GetInfo() *ObjectInfo {
	if x, ok := x.GetData().(*CopyObject_Info); ok {
		return x.Info
	}
	return nil
}

func (x *CopyObject) GetContent() []byte {
	if x, ok := x.GetData().(*CopyObject_Content); ok {
		return x.Content
	}
	return nil
}

type isCopyObject_Data interface {
	isCopyObject_Data()
}

type CopyObject_Info struct {
	Info *ObjectInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type CopyObject_Content struct {
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3,oneof"`
}

func (*CopyObject_Info) isCopyObject_Data() {}

func (*CopyObject_Content) isCopyObject_Data() {}

type PutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *PutResponse) Reset() {
	*x = PutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chunkServer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutResponse) ProtoMessage() {}

func (x *PutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chunkServer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutResponse.ProtoReflect.Descriptor instead.
func (*PutResponse) Descriptor() ([]byte, []int) {
	return file_chunkServer_proto_rawDescGZIP(), []int{3}
}

func (x *PutResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PutResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Dir  string `protobuf:"bytes,2,opt,name=dir,proto3" json:"dir,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chunkServer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chunkServer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_chunkServer_proto_rawDescGZIP(), []int{4}
}

func (x *GetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetRequest) GetDir() string {
	if x != nil {
		return x.Dir
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chunkServer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chunkServer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_chunkServer_proto_rawDescGZIP(), []int{5}
}

func (x *GetResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type TaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskType int32  `protobuf:"varint,1,opt,name=taskType,proto3" json:"taskType,omitempty"`
	Value    []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TaskRequest) Reset() {
	*x = TaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chunkServer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRequest) ProtoMessage() {}

func (x *TaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chunkServer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRequest.ProtoReflect.Descriptor instead.
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return file_chunkServer_proto_rawDescGZIP(), []int{6}
}

func (x *TaskRequest) GetTaskType() int32 {
	if x != nil {
		return x.TaskType
	}
	return 0
}

func (x *TaskRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type TaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chunkServer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chunkServer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_chunkServer_proto_rawDescGZIP(), []int{7}
}

func (x *TaskResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_chunkServer_proto protoreflect.FileDescriptor

var file_chunkServer_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x10, 0x6d, 0x65, 0x74, 0x61, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x0a, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x56, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x56, 0x0a,
	0x0a, 0x43, 0x6f, 0x70, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0x1a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x06, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3b, 0x0a, 0x0b, 0x50, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x32, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x64, 0x69, 0x72, 0x22, 0x27, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x3f, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x26, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xd9, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x09, 0x50, 0x75, 0x74, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x31, 0x0a, 0x0a, 0x43, 0x6f,
	0x70, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x30, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x33, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12,
	0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chunkServer_proto_rawDescOnce sync.Once
	file_chunkServer_proto_rawDescData = file_chunkServer_proto_rawDesc
)

func file_chunkServer_proto_rawDescGZIP() []byte {
	file_chunkServer_proto_rawDescOnce.Do(func() {
		file_chunkServer_proto_rawDescData = protoimpl.X.CompressGZIP(file_chunkServer_proto_rawDescData)
	})
	return file_chunkServer_proto_rawDescData
}

var file_chunkServer_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_chunkServer_proto_goTypes = []interface{}{
	(*ObjectInfo)(nil),   // 0: pb.ObjectInfo
	(*PutRequest)(nil),   // 1: pb.PutRequest
	(*CopyObject)(nil),   // 2: pb.CopyObject
	(*PutResponse)(nil),  // 3: pb.PutResponse
	(*GetRequest)(nil),   // 4: pb.GetRequest
	(*GetResponse)(nil),  // 5: pb.GetResponse
	(*TaskRequest)(nil),  // 6: pb.TaskRequest
	(*TaskResponse)(nil), // 7: pb.TaskResponse
	(*Target)(nil),       // 8: pb.Target
}
var file_chunkServer_proto_depIdxs = []int32{
	8, // 0: pb.ObjectInfo.targets:type_name -> pb.Target
	0, // 1: pb.PutRequest.info:type_name -> pb.ObjectInfo
	0, // 2: pb.CopyObject.info:type_name -> pb.ObjectInfo
	1, // 3: pb.ChunkServer.PutObject:input_type -> pb.PutRequest
	1, // 4: pb.ChunkServer.CopyObject:input_type -> pb.PutRequest
	4, // 5: pb.ChunkServer.GetObject:input_type -> pb.GetRequest
	6, // 6: pb.ChunkServer.TaskReceiver:input_type -> pb.TaskRequest
	3, // 7: pb.ChunkServer.PutObject:output_type -> pb.PutResponse
	3, // 8: pb.ChunkServer.CopyObject:output_type -> pb.PutResponse
	5, // 9: pb.ChunkServer.GetObject:output_type -> pb.GetResponse
	7, // 10: pb.ChunkServer.TaskReceiver:output_type -> pb.TaskResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_chunkServer_proto_init() }
func file_chunkServer_proto_init() {
	if File_chunkServer_proto != nil {
		return
	}
	file_metaServer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chunkServer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectInfo); i {
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
		file_chunkServer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutRequest); i {
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
		file_chunkServer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyObject); i {
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
		file_chunkServer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutResponse); i {
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
		file_chunkServer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_chunkServer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_chunkServer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRequest); i {
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
		file_chunkServer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResponse); i {
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
	file_chunkServer_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PutRequest_Info)(nil),
		(*PutRequest_Content)(nil),
	}
	file_chunkServer_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CopyObject_Info)(nil),
		(*CopyObject_Content)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chunkServer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chunkServer_proto_goTypes,
		DependencyIndexes: file_chunkServer_proto_depIdxs,
		MessageInfos:      file_chunkServer_proto_msgTypes,
	}.Build()
	File_chunkServer_proto = out.File
	file_chunkServer_proto_rawDesc = nil
	file_chunkServer_proto_goTypes = nil
	file_chunkServer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChunkServerClient is the client API for ChunkServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChunkServerClient interface {
	PutObject(ctx context.Context, opts ...grpc.CallOption) (ChunkServer_PutObjectClient, error)
	CopyObject(ctx context.Context, opts ...grpc.CallOption) (ChunkServer_CopyObjectClient, error)
	GetObject(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (ChunkServer_GetObjectClient, error)
	TaskReceiver(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
}

type chunkServerClient struct {
	cc grpc.ClientConnInterface
}

func NewChunkServerClient(cc grpc.ClientConnInterface) ChunkServerClient {
	return &chunkServerClient{cc}
}

func (c *chunkServerClient) PutObject(ctx context.Context, opts ...grpc.CallOption) (ChunkServer_PutObjectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChunkServer_serviceDesc.Streams[0], "/pb.ChunkServer/PutObject", opts...)
	if err != nil {
		return nil, err
	}
	x := &chunkServerPutObjectClient{stream}
	return x, nil
}

type ChunkServer_PutObjectClient interface {
	Send(*PutRequest) error
	CloseAndRecv() (*PutResponse, error)
	grpc.ClientStream
}

type chunkServerPutObjectClient struct {
	grpc.ClientStream
}

func (x *chunkServerPutObjectClient) Send(m *PutRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chunkServerPutObjectClient) CloseAndRecv() (*PutResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PutResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chunkServerClient) CopyObject(ctx context.Context, opts ...grpc.CallOption) (ChunkServer_CopyObjectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChunkServer_serviceDesc.Streams[1], "/pb.ChunkServer/CopyObject", opts...)
	if err != nil {
		return nil, err
	}
	x := &chunkServerCopyObjectClient{stream}
	return x, nil
}

type ChunkServer_CopyObjectClient interface {
	Send(*PutRequest) error
	CloseAndRecv() (*PutResponse, error)
	grpc.ClientStream
}

type chunkServerCopyObjectClient struct {
	grpc.ClientStream
}

func (x *chunkServerCopyObjectClient) Send(m *PutRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chunkServerCopyObjectClient) CloseAndRecv() (*PutResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PutResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chunkServerClient) GetObject(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (ChunkServer_GetObjectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChunkServer_serviceDesc.Streams[2], "/pb.ChunkServer/GetObject", opts...)
	if err != nil {
		return nil, err
	}
	x := &chunkServerGetObjectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChunkServer_GetObjectClient interface {
	Recv() (*GetResponse, error)
	grpc.ClientStream
}

type chunkServerGetObjectClient struct {
	grpc.ClientStream
}

func (x *chunkServerGetObjectClient) Recv() (*GetResponse, error) {
	m := new(GetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chunkServerClient) TaskReceiver(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/pb.ChunkServer/TaskReceiver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChunkServerServer is the server API for ChunkServer service.
type ChunkServerServer interface {
	PutObject(ChunkServer_PutObjectServer) error
	CopyObject(ChunkServer_CopyObjectServer) error
	GetObject(*GetRequest, ChunkServer_GetObjectServer) error
	TaskReceiver(context.Context, *TaskRequest) (*TaskResponse, error)
}

// UnimplementedChunkServerServer can be embedded to have forward compatible implementations.
type UnimplementedChunkServerServer struct {
}

func (*UnimplementedChunkServerServer) PutObject(ChunkServer_PutObjectServer) error {
	return status.Errorf(codes.Unimplemented, "method PutObject not implemented")
}
func (*UnimplementedChunkServerServer) CopyObject(ChunkServer_CopyObjectServer) error {
	return status.Errorf(codes.Unimplemented, "method CopyObject not implemented")
}
func (*UnimplementedChunkServerServer) GetObject(*GetRequest, ChunkServer_GetObjectServer) error {
	return status.Errorf(codes.Unimplemented, "method GetObject not implemented")
}
func (*UnimplementedChunkServerServer) TaskReceiver(context.Context, *TaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskReceiver not implemented")
}

func RegisterChunkServerServer(s *grpc.Server, srv ChunkServerServer) {
	s.RegisterService(&_ChunkServer_serviceDesc, srv)
}

func _ChunkServer_PutObject_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChunkServerServer).PutObject(&chunkServerPutObjectServer{stream})
}

type ChunkServer_PutObjectServer interface {
	SendAndClose(*PutResponse) error
	Recv() (*PutRequest, error)
	grpc.ServerStream
}

type chunkServerPutObjectServer struct {
	grpc.ServerStream
}

func (x *chunkServerPutObjectServer) SendAndClose(m *PutResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chunkServerPutObjectServer) Recv() (*PutRequest, error) {
	m := new(PutRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChunkServer_CopyObject_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChunkServerServer).CopyObject(&chunkServerCopyObjectServer{stream})
}

type ChunkServer_CopyObjectServer interface {
	SendAndClose(*PutResponse) error
	Recv() (*PutRequest, error)
	grpc.ServerStream
}

type chunkServerCopyObjectServer struct {
	grpc.ServerStream
}

func (x *chunkServerCopyObjectServer) SendAndClose(m *PutResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chunkServerCopyObjectServer) Recv() (*PutRequest, error) {
	m := new(PutRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChunkServer_GetObject_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChunkServerServer).GetObject(m, &chunkServerGetObjectServer{stream})
}

type ChunkServer_GetObjectServer interface {
	Send(*GetResponse) error
	grpc.ServerStream
}

type chunkServerGetObjectServer struct {
	grpc.ServerStream
}

func (x *chunkServerGetObjectServer) Send(m *GetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ChunkServer_TaskReceiver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServerServer).TaskReceiver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ChunkServer/TaskReceiver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServerServer).TaskReceiver(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChunkServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ChunkServer",
	HandlerType: (*ChunkServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TaskReceiver",
			Handler:    _ChunkServer_TaskReceiver_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PutObject",
			Handler:       _ChunkServer_PutObject_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CopyObject",
			Handler:       _ChunkServer_CopyObject_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetObject",
			Handler:       _ChunkServer_GetObject_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chunkServer.proto",
}
