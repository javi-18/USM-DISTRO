// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.0
// 	protoc        v3.12.4
// source: servidoreshextech.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// Representa una entrada en el reloj vectorial
type VectorClockEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`               // Identificador del nodo
	Values []int32 `protobuf:"varint,2,rep,packed,name=values,proto3" json:"values,omitempty"` // Valores del reloj vectorial
}

func (x *VectorClockEntry) Reset() {
	*x = VectorClockEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servidoreshextech_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VectorClockEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VectorClockEntry) ProtoMessage() {}

func (x *VectorClockEntry) ProtoReflect() protoreflect.Message {
	mi := &file_servidoreshextech_proto_msgTypes[0]
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
	return file_servidoreshextech_proto_rawDescGZIP(), []int{0}
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

// Solicitud para agregar o actualizar productos
type ProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region   string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`      // Región del producto
	Product  string `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`    // Nombre del producto
	Quantity int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"` // Cantidad (o valor asociado)
}

func (x *ProductRequest) Reset() {
	*x = ProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servidoreshextech_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRequest) ProtoMessage() {}

func (x *ProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_servidoreshextech_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRequest.ProtoReflect.Descriptor instead.
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return file_servidoreshextech_proto_rawDescGZIP(), []int{1}
}

func (x *ProductRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *ProductRequest) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *ProductRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

// Solicitud para renombrar productos
type RenameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region  string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`                  // Región del producto
	OldName string `protobuf:"bytes,2,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"` // Nombre actual del producto
	NewName string `protobuf:"bytes,3,opt,name=new_name,json=newName,proto3" json:"new_name,omitempty"` // Nuevo nombre para el producto
}

func (x *RenameRequest) Reset() {
	*x = RenameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servidoreshextech_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameRequest) ProtoMessage() {}

func (x *RenameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_servidoreshextech_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameRequest.ProtoReflect.Descriptor instead.
func (*RenameRequest) Descriptor() ([]byte, []int) {
	return file_servidoreshextech_proto_rawDescGZIP(), []int{2}
}

func (x *RenameRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *RenameRequest) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *RenameRequest) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

// Solicitud para propagar cambios entre nodos
type PropagateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VectorClock []*VectorClockEntry `protobuf:"bytes,1,rep,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"` // Estado del reloj vectorial
	Logs        []string            `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty"`                                  // Registros de cambios realizados
}

func (x *PropagateRequest) Reset() {
	*x = PropagateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servidoreshextech_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropagateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropagateRequest) ProtoMessage() {}

func (x *PropagateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_servidoreshextech_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropagateRequest.ProtoReflect.Descriptor instead.
func (*PropagateRequest) Descriptor() ([]byte, []int) {
	return file_servidoreshextech_proto_rawDescGZIP(), []int{3}
}

func (x *PropagateRequest) GetVectorClock() []*VectorClockEntry {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

func (x *PropagateRequest) GetLogs() []string {
	if x != nil {
		return x.Logs
	}
	return nil
}

// Respuesta genérica
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`  // Mensaje de respuesta
	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"` // Indica si la operación fue exitosa
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servidoreshextech_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_servidoreshextech_proto_msgTypes[4]
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
	return file_servidoreshextech_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// Solicitud para sincronización entre nodos
type SyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VectorClock []*VectorClockEntry `protobuf:"bytes,1,rep,name=vector_clock,json=vectorClock,proto3" json:"vector_clock,omitempty"` // Estado del reloj vectorial para sincronizar
	SourceNode  string              `protobuf:"bytes,2,opt,name=source_node,json=sourceNode,proto3" json:"source_node,omitempty"`    // Nodo origen de la sincronización
}

func (x *SyncRequest) Reset() {
	*x = SyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servidoreshextech_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRequest) ProtoMessage() {}

func (x *SyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_servidoreshextech_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRequest.ProtoReflect.Descriptor instead.
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return file_servidoreshextech_proto_rawDescGZIP(), []int{5}
}

func (x *SyncRequest) GetVectorClock() []*VectorClockEntry {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

func (x *SyncRequest) GetSourceNode() string {
	if x != nil {
		return x.SourceNode
	}
	return ""
}

// Respuesta a una solicitud de sincronización
type SyncResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdatedClock []*VectorClockEntry `protobuf:"bytes,1,rep,name=updated_clock,json=updatedClock,proto3" json:"updated_clock,omitempty"` // Reloj vectorial actualizado
	UpdatedLogs  []string            `protobuf:"bytes,2,rep,name=updated_logs,json=updatedLogs,proto3" json:"updated_logs,omitempty"`    // Cambios sincronizados
}

func (x *SyncResponse) Reset() {
	*x = SyncResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servidoreshextech_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncResponse) ProtoMessage() {}

func (x *SyncResponse) ProtoReflect() protoreflect.Message {
	mi := &file_servidoreshextech_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncResponse.ProtoReflect.Descriptor instead.
func (*SyncResponse) Descriptor() ([]byte, []int) {
	return file_servidoreshextech_proto_rawDescGZIP(), []int{6}
}

func (x *SyncResponse) GetUpdatedClock() []*VectorClockEntry {
	if x != nil {
		return x.UpdatedClock
	}
	return nil
}

func (x *SyncResponse) GetUpdatedLogs() []string {
	if x != nil {
		return x.UpdatedLogs
	}
	return nil
}

var File_servidoreshextech_proto protoreflect.FileDescriptor

var file_servidoreshextech_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x68, 0x65, 0x78, 0x74,
	0x65, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x64, 0x6f, 0x72, 0x65, 0x73, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x10, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x5e, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x5d, 0x0a, 0x0d, 0x52, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e,
	0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e,
	0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x6e, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0c, 0x76, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x68, 0x65, 0x78,
	0x74, 0x65, 0x63, 0x68, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x3e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x76, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f,
	0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x2e,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0b, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x7b,
	0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48,
	0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72,
	0x65, 0x73, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x73, 0x32, 0x9e, 0x03, 0x0a, 0x0e,
	0x48, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c,
	0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x21, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x68, 0x65, 0x78, 0x74,
	0x65, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d,
	0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x20, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63,
	0x68, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x68, 0x65, 0x78, 0x74,
	0x65, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x21, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63,
	0x68, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x68, 0x65, 0x78,
	0x74, 0x65, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a,
	0x10, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x12, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x68, 0x65,
	0x78, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4c,
	0x0a, 0x09, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x2e,
	0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x68, 0x65, 0x78, 0x74, 0x65, 0x63, 0x68, 0x2e,
	0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_servidoreshextech_proto_rawDescOnce sync.Once
	file_servidoreshextech_proto_rawDescData = file_servidoreshextech_proto_rawDesc
)

func file_servidoreshextech_proto_rawDescGZIP() []byte {
	file_servidoreshextech_proto_rawDescOnce.Do(func() {
		file_servidoreshextech_proto_rawDescData = protoimpl.X.CompressGZIP(file_servidoreshextech_proto_rawDescData)
	})
	return file_servidoreshextech_proto_rawDescData
}

var file_servidoreshextech_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_servidoreshextech_proto_goTypes = []interface{}{
	(*VectorClockEntry)(nil), // 0: servidoreshextech.VectorClockEntry
	(*ProductRequest)(nil),   // 1: servidoreshextech.ProductRequest
	(*RenameRequest)(nil),    // 2: servidoreshextech.RenameRequest
	(*PropagateRequest)(nil), // 3: servidoreshextech.PropagateRequest
	(*Response)(nil),         // 4: servidoreshextech.Response
	(*SyncRequest)(nil),      // 5: servidoreshextech.SyncRequest
	(*SyncResponse)(nil),     // 6: servidoreshextech.SyncResponse
	(*empty.Empty)(nil),      // 7: google.protobuf.Empty
}
var file_servidoreshextech_proto_depIdxs = []int32{
	0, // 0: servidoreshextech.PropagateRequest.vector_clock:type_name -> servidoreshextech.VectorClockEntry
	0, // 1: servidoreshextech.SyncRequest.vector_clock:type_name -> servidoreshextech.VectorClockEntry
	0, // 2: servidoreshextech.SyncResponse.updated_clock:type_name -> servidoreshextech.VectorClockEntry
	1, // 3: servidoreshextech.HextechService.AddProduct:input_type -> servidoreshextech.ProductRequest
	2, // 4: servidoreshextech.HextechService.RenameProduct:input_type -> servidoreshextech.RenameRequest
	1, // 5: servidoreshextech.HextechService.DeleteProduct:input_type -> servidoreshextech.ProductRequest
	3, // 6: servidoreshextech.HextechService.PropagateChanges:input_type -> servidoreshextech.PropagateRequest
	5, // 7: servidoreshextech.HextechService.SyncState:input_type -> servidoreshextech.SyncRequest
	4, // 8: servidoreshextech.HextechService.AddProduct:output_type -> servidoreshextech.Response
	4, // 9: servidoreshextech.HextechService.RenameProduct:output_type -> servidoreshextech.Response
	4, // 10: servidoreshextech.HextechService.DeleteProduct:output_type -> servidoreshextech.Response
	7, // 11: servidoreshextech.HextechService.PropagateChanges:output_type -> google.protobuf.Empty
	6, // 12: servidoreshextech.HextechService.SyncState:output_type -> servidoreshextech.SyncResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_servidoreshextech_proto_init() }
func file_servidoreshextech_proto_init() {
	if File_servidoreshextech_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_servidoreshextech_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_servidoreshextech_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductRequest); i {
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
		file_servidoreshextech_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameRequest); i {
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
		file_servidoreshextech_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropagateRequest); i {
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
		file_servidoreshextech_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_servidoreshextech_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRequest); i {
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
		file_servidoreshextech_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncResponse); i {
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
			RawDescriptor: file_servidoreshextech_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_servidoreshextech_proto_goTypes,
		DependencyIndexes: file_servidoreshextech_proto_depIdxs,
		MessageInfos:      file_servidoreshextech_proto_msgTypes,
	}.Build()
	File_servidoreshextech_proto = out.File
	file_servidoreshextech_proto_rawDesc = nil
	file_servidoreshextech_proto_goTypes = nil
	file_servidoreshextech_proto_depIdxs = nil
}
