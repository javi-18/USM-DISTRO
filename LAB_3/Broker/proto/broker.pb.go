// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.0
// 	protoc        v3.12.4
// source: broker.proto

package proto

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

type VectorClockEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values []int32 `protobuf:"varint,2,rep,packed,name=values,proto3" json:"values,omitempty"`
}

func (x *VectorClockEntry) Reset() {
	*x = VectorClockEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VectorClockEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VectorClockEntry) ProtoMessage() {}

func (x *VectorClockEntry) ProtoReflect() protoreflect.Message {
	mi := &file_broker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VectorClockEntry.ProtoReflect.Descriptor instead.
func (*VectorClockEntry) Descriptor() ([]byte, []int) {
	return file_broker_proto_rawDescGZIP(), []int{0}
}

func (x *VectorClockEntry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *VectorClockEntry) GetValues() []int32 {
	if x != nil {
		return x.Values
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandType string              `protobuf:"bytes,1,opt,name=command_type,json=commandType,proto3" json:"command_type,omitempty"`
	Region      string              `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	Product     string              `protobuf:"bytes,3,opt,name=product,proto3" json:"product,omitempty"`
	Value       int32               `protobuf:"varint,4,opt,name=value,proto3" json:"value,omitempty"`
	NewName     string              `protobuf:"bytes,5,opt,name=new_name,json=newName,proto3" json:"new_name,omitempty"`
	VectorClock []*VectorClockEntry `protobuf:"bytes,6,rep,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_broker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_broker_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetCommandType() string {
	if x != nil {
		return x.CommandType
	}
	return ""
}

func (x *Request) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Request) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *Request) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Request) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

func (x *Request) GetVectorClock() []*VectorClockEntry {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerAddress string              `protobuf:"bytes,1,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"`
	VectorClock   []*VectorClockEntry `protobuf:"bytes,2,rep,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
	Message       string              `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_broker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_broker_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetServerAddress() string {
	if x != nil {
		return x.ServerAddress
	}
	return ""
}

func (x *Response) GetVectorClock() []*VectorClockEntry {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type InconsistencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region      string              `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Product     string              `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	VectorClock []*VectorClockEntry `protobuf:"bytes,3,rep,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"`
}

func (x *InconsistencyRequest) Reset() {
	*x = InconsistencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InconsistencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InconsistencyRequest) ProtoMessage() {}

func (x *InconsistencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_broker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InconsistencyRequest.ProtoReflect.Descriptor instead.
func (*InconsistencyRequest) Descriptor() ([]byte, []int) {
	return file_broker_proto_rawDescGZIP(), []int{3}
}

func (x *InconsistencyRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *InconsistencyRequest) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *InconsistencyRequest) GetVectorClock() []*VectorClockEntry {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_broker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_broker_proto_rawDescGZIP(), []int{4}
}

var File_broker_proto protoreflect.FileDescriptor

var file_broker_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x22, 0x3c, 0x0a, 0x10, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x22, 0xcc, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e,
	0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e,
	0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63,
	0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c,
	0x6f, 0x63, 0x6b, 0x22, 0x88, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3b, 0x0a, 0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f,
	0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x85,
	0x01, 0x0a, 0x14, 0x49, 0x6e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x3b, 0x0a, 0x0c, 0x76, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43,
	0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32,
	0x8c, 0x01, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x34, 0x0a, 0x0f, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x13, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x49, 0x6e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1c,
	0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xbf,
	0x01, 0x0a, 0x0e, 0x48, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x32, 0x0a, 0x0d, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0f, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x49, 0x6e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1c, 0x2e,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x62, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x0d, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10,
	0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_broker_proto_rawDescOnce sync.Once
	file_broker_proto_rawDescData = file_broker_proto_rawDesc
)

func file_broker_proto_rawDescGZIP() []byte {
	file_broker_proto_rawDescOnce.Do(func() {
		file_broker_proto_rawDescData = protoimpl.X.CompressGZIP(file_broker_proto_rawDescData)
	})
	return file_broker_proto_rawDescData
}

var file_broker_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_broker_proto_goTypes = []interface{}{
	(*VectorClockEntry)(nil),     // 0: broker.VectorClockEntry
	(*Request)(nil),              // 1: broker.Request
	(*Response)(nil),             // 2: broker.Response
	(*InconsistencyRequest)(nil), // 3: broker.InconsistencyRequest
	(*Empty)(nil),                // 4: broker.Empty
}
var file_broker_proto_depIdxs = []int32{
	0, // 0: broker.Request.vector_clock:type_name -> broker.VectorClockEntry
	0, // 1: broker.Response.vector_clock:type_name -> broker.VectorClockEntry
	0, // 2: broker.InconsistencyRequest.vector_clock:type_name -> broker.VectorClockEntry
	1, // 3: broker.BrokerService.RedirectRequest:input_type -> broker.Request
	3, // 4: broker.BrokerService.HandleInconsistency:input_type -> broker.InconsistencyRequest
	1, // 5: broker.HextechService.HandleRequest:input_type -> broker.Request
	3, // 6: broker.HextechService.ResolveInconsistency:input_type -> broker.InconsistencyRequest
	4, // 7: broker.HextechService.GetVectorClock:input_type -> broker.Empty
	2, // 8: broker.BrokerService.RedirectRequest:output_type -> broker.Response
	2, // 9: broker.BrokerService.HandleInconsistency:output_type -> broker.Response
	2, // 10: broker.HextechService.HandleRequest:output_type -> broker.Response
	2, // 11: broker.HextechService.ResolveInconsistency:output_type -> broker.Response
	2, // 12: broker.HextechService.GetVectorClock:output_type -> broker.Response
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_broker_proto_init() }
func file_broker_proto_init() {
	if File_broker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_broker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VectorClockEntry); i {
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
		file_broker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_broker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_broker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InconsistencyRequest); i {
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
		file_broker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_broker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_broker_proto_goTypes,
		DependencyIndexes: file_broker_proto_depIdxs,
		MessageInfos:      file_broker_proto_msgTypes,
	}.Build()
	File_broker_proto = out.File
	file_broker_proto_rawDesc = nil
	file_broker_proto_goTypes = nil
	file_broker_proto_depIdxs = nil
}
