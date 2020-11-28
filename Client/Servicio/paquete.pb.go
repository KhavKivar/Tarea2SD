// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: paquete.proto

package Servicio

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

type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_Ok      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

// Enum value maps for UploadStatusCode.
var (
	UploadStatusCode_name = map[int32]string{
		0: "Unknown",
		1: "Ok",
		2: "Failed",
	}
	UploadStatusCode_value = map[string]int32{
		"Unknown": 0,
		"Ok":      1,
		"Failed":  2,
	}
)

func (x UploadStatusCode) Enum() *UploadStatusCode {
	p := new(UploadStatusCode)
	*p = x
	return p
}

func (x UploadStatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadStatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_paquete_proto_enumTypes[0].Descriptor()
}

func (UploadStatusCode) Type() protoreflect.EnumType {
	return &file_paquete_proto_enumTypes[0]
}

func (x UploadStatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadStatusCode.Descriptor instead.
func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return file_paquete_proto_rawDescGZIP(), []int{0}
}

type Books struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book []string `protobuf:"bytes,1,rep,name=Book,proto3" json:"Book,omitempty"`
}

func (x *Books) Reset() {
	*x = Books{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paquete_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Books) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Books) ProtoMessage() {}

func (x *Books) ProtoReflect() protoreflect.Message {
	mi := &file_paquete_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Books.ProtoReflect.Descriptor instead.
func (*Books) Descriptor() ([]byte, []int) {
	return file_paquete_proto_rawDescGZIP(), []int{0}
}

func (x *Books) GetBook() []string {
	if x != nil {
		return x.Book
	}
	return nil
}

type ChunkDes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book string `protobuf:"bytes,1,opt,name=Book,proto3" json:"Book,omitempty"`
	Part string `protobuf:"bytes,2,opt,name=Part,proto3" json:"Part,omitempty"`
}

func (x *ChunkDes) Reset() {
	*x = ChunkDes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paquete_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkDes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkDes) ProtoMessage() {}

func (x *ChunkDes) ProtoReflect() protoreflect.Message {
	mi := &file_paquete_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkDes.ProtoReflect.Descriptor instead.
func (*ChunkDes) Descriptor() ([]byte, []int) {
	return file_paquete_proto_rawDescGZIP(), []int{1}
}

func (x *ChunkDes) GetBook() string {
	if x != nil {
		return x.Book
	}
	return ""
}

func (x *ChunkDes) GetPart() string {
	if x != nil {
		return x.Part
	}
	return ""
}

type ChunkBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contenido []byte `protobuf:"bytes,1,opt,name=Contenido,proto3" json:"Contenido,omitempty"`
}

func (x *ChunkBook) Reset() {
	*x = ChunkBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paquete_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkBook) ProtoMessage() {}

func (x *ChunkBook) ProtoReflect() protoreflect.Message {
	mi := &file_paquete_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkBook.ProtoReflect.Descriptor instead.
func (*ChunkBook) Descriptor() ([]byte, []int) {
	return file_paquete_proto_rawDescGZIP(), []int{2}
}

func (x *ChunkBook) GetContenido() []byte {
	if x != nil {
		return x.Contenido
	}
	return nil
}

type ListChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkList []string `protobuf:"bytes,1,rep,name=ChunkList,proto3" json:"ChunkList,omitempty"`
	Ext       string   `protobuf:"bytes,2,opt,name=ext,proto3" json:"ext,omitempty"`
}

func (x *ListChunk) Reset() {
	*x = ListChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paquete_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChunk) ProtoMessage() {}

func (x *ListChunk) ProtoReflect() protoreflect.Message {
	mi := &file_paquete_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChunk.ProtoReflect.Descriptor instead.
func (*ListChunk) Descriptor() ([]byte, []int) {
	return file_paquete_proto_rawDescGZIP(), []int{3}
}

func (x *ListChunk) GetChunkList() []string {
	if x != nil {
		return x.ChunkList
	}
	return nil
}

func (x *ListChunk) GetExt() string {
	if x != nil {
		return x.Ext
	}
	return ""
}

type BookToDownload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book string `protobuf:"bytes,1,opt,name=Book,proto3" json:"Book,omitempty"`
}

func (x *BookToDownload) Reset() {
	*x = BookToDownload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paquete_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookToDownload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookToDownload) ProtoMessage() {}

func (x *BookToDownload) ProtoReflect() protoreflect.Message {
	mi := &file_paquete_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookToDownload.ProtoReflect.Descriptor instead.
func (*BookToDownload) Descriptor() ([]byte, []int) {
	return file_paquete_proto_rawDescGZIP(), []int{4}
}

func (x *BookToDownload) GetBook() string {
	if x != nil {
		return x.Book
	}
	return ""
}

type ChunkSendToServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contenido   []byte `protobuf:"bytes,1,opt,name=Contenido,proto3" json:"Contenido,omitempty"`
	TotalChunks int32  `protobuf:"varint,2,opt,name=totalChunks,proto3" json:"totalChunks,omitempty"`
	NumeroChunk int32  `protobuf:"varint,3,opt,name=numeroChunk,proto3" json:"numeroChunk,omitempty"`
	Nombre      string `protobuf:"bytes,4,opt,name=nombre,proto3" json:"nombre,omitempty"`
}

func (x *ChunkSendToServer) Reset() {
	*x = ChunkSendToServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paquete_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChunkSendToServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkSendToServer) ProtoMessage() {}

func (x *ChunkSendToServer) ProtoReflect() protoreflect.Message {
	mi := &file_paquete_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkSendToServer.ProtoReflect.Descriptor instead.
func (*ChunkSendToServer) Descriptor() ([]byte, []int) {
	return file_paquete_proto_rawDescGZIP(), []int{5}
}

func (x *ChunkSendToServer) GetContenido() []byte {
	if x != nil {
		return x.Contenido
	}
	return nil
}

func (x *ChunkSendToServer) GetTotalChunks() int32 {
	if x != nil {
		return x.TotalChunks
	}
	return 0
}

func (x *ChunkSendToServer) GetNumeroChunk() int32 {
	if x != nil {
		return x.NumeroChunk
	}
	return 0
}

func (x *ChunkSendToServer) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

type Mensaje struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *Mensaje) Reset() {
	*x = Mensaje{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paquete_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mensaje) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mensaje) ProtoMessage() {}

func (x *Mensaje) ProtoReflect() protoreflect.Message {
	mi := &file_paquete_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mensaje.ProtoReflect.Descriptor instead.
func (*Mensaje) Descriptor() ([]byte, []int) {
	return file_paquete_proto_rawDescGZIP(), []int{6}
}

func (x *Mensaje) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Respuesta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje            string `protobuf:"bytes,1,opt,name=Mensaje,proto3" json:"Mensaje,omitempty"`
	ChunkSendToServer1 int32  `protobuf:"varint,3,opt,name=ChunkSendToServer1,proto3" json:"ChunkSendToServer1,omitempty"`
	ChunkSendToServer2 int32  `protobuf:"varint,4,opt,name=ChunkSendToServer2,proto3" json:"ChunkSendToServer2,omitempty"`
	ChunkSendToServer3 int32  `protobuf:"varint,5,opt,name=ChunkSendToServer3,proto3" json:"ChunkSendToServer3,omitempty"`
}

func (x *Respuesta) Reset() {
	*x = Respuesta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paquete_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Respuesta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Respuesta) ProtoMessage() {}

func (x *Respuesta) ProtoReflect() protoreflect.Message {
	mi := &file_paquete_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Respuesta.ProtoReflect.Descriptor instead.
func (*Respuesta) Descriptor() ([]byte, []int) {
	return file_paquete_proto_rawDescGZIP(), []int{7}
}

func (x *Respuesta) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

func (x *Respuesta) GetChunkSendToServer1() int32 {
	if x != nil {
		return x.ChunkSendToServer1
	}
	return 0
}

func (x *Respuesta) GetChunkSendToServer2() int32 {
	if x != nil {
		return x.ChunkSendToServer2
	}
	return 0
}

func (x *Respuesta) GetChunkSendToServer3() int32 {
	if x != nil {
		return x.ChunkSendToServer3
	}
	return 0
}

type Propuesta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book               string `protobuf:"bytes,1,opt,name=Book,proto3" json:"Book,omitempty"`
	TotalChunks        int32  `protobuf:"varint,2,opt,name=totalChunks,proto3" json:"totalChunks,omitempty"`
	ChunkSendToServer1 int32  `protobuf:"varint,3,opt,name=ChunkSendToServer1,proto3" json:"ChunkSendToServer1,omitempty"`
	ChunkSendToServer2 int32  `protobuf:"varint,4,opt,name=ChunkSendToServer2,proto3" json:"ChunkSendToServer2,omitempty"`
	ChunkSendToServer3 int32  `protobuf:"varint,5,opt,name=ChunkSendToServer3,proto3" json:"ChunkSendToServer3,omitempty"`
	Ext                string `protobuf:"bytes,6,opt,name=ext,proto3" json:"ext,omitempty"`
}

func (x *Propuesta) Reset() {
	*x = Propuesta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paquete_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Propuesta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Propuesta) ProtoMessage() {}

func (x *Propuesta) ProtoReflect() protoreflect.Message {
	mi := &file_paquete_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Propuesta.ProtoReflect.Descriptor instead.
func (*Propuesta) Descriptor() ([]byte, []int) {
	return file_paquete_proto_rawDescGZIP(), []int{8}
}

func (x *Propuesta) GetBook() string {
	if x != nil {
		return x.Book
	}
	return ""
}

func (x *Propuesta) GetTotalChunks() int32 {
	if x != nil {
		return x.TotalChunks
	}
	return 0
}

func (x *Propuesta) GetChunkSendToServer1() int32 {
	if x != nil {
		return x.ChunkSendToServer1
	}
	return 0
}

func (x *Propuesta) GetChunkSendToServer2() int32 {
	if x != nil {
		return x.ChunkSendToServer2
	}
	return 0
}

func (x *Propuesta) GetChunkSendToServer3() int32 {
	if x != nil {
		return x.ChunkSendToServer3
	}
	return 0
}

func (x *Propuesta) GetExt() string {
	if x != nil {
		return x.Ext
	}
	return ""
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contenido   []byte `protobuf:"bytes,1,opt,name=Contenido,proto3" json:"Contenido,omitempty"`
	TotalChunks int32  `protobuf:"varint,2,opt,name=totalChunks,proto3" json:"totalChunks,omitempty"`
	NumeroChunk int32  `protobuf:"varint,3,opt,name=numeroChunk,proto3" json:"numeroChunk,omitempty"`
	Nombre      string `protobuf:"bytes,4,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Ext         string `protobuf:"bytes,5,opt,name=ext,proto3" json:"ext,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paquete_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_paquete_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_paquete_proto_rawDescGZIP(), []int{9}
}

func (x *Chunk) GetContenido() []byte {
	if x != nil {
		return x.Contenido
	}
	return nil
}

func (x *Chunk) GetTotalChunks() int32 {
	if x != nil {
		return x.TotalChunks
	}
	return 0
}

func (x *Chunk) GetNumeroChunk() int32 {
	if x != nil {
		return x.NumeroChunk
	}
	return 0
}

func (x *Chunk) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Chunk) GetExt() string {
	if x != nil {
		return x.Ext
	}
	return ""
}

type UploadStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mensaje string           `protobuf:"bytes,1,opt,name=Mensaje,proto3" json:"Mensaje,omitempty"`
	Code    UploadStatusCode `protobuf:"varint,2,opt,name=Code,proto3,enum=paquete.UploadStatusCode" json:"Code,omitempty"`
}

func (x *UploadStatus) Reset() {
	*x = UploadStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paquete_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadStatus) ProtoMessage() {}

func (x *UploadStatus) ProtoReflect() protoreflect.Message {
	mi := &file_paquete_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadStatus.ProtoReflect.Descriptor instead.
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return file_paquete_proto_rawDescGZIP(), []int{10}
}

func (x *UploadStatus) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

func (x *UploadStatus) GetCode() UploadStatusCode {
	if x != nil {
		return x.Code
	}
	return UploadStatusCode_Unknown
}

var File_paquete_proto protoreflect.FileDescriptor

var file_paquete_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x22, 0x1b, 0x0a, 0x05, 0x42, 0x6f, 0x6f, 0x6b,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x32, 0x0a, 0x08, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x65,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x72, 0x74, 0x22, 0x29, 0x0a, 0x09, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x69, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x69, 0x64, 0x6f, 0x22, 0x3b, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x78,
	0x74, 0x22, 0x24, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x54, 0x6f, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x8d, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x69, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x69, 0x64, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x22, 0x1b, 0x0a, 0x07, 0x4d, 0x65, 0x6e, 0x73, 0x61,
	0x6a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4d, 0x73, 0x67, 0x22, 0xb5, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73,
	0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x12, 0x2e, 0x0a, 0x12,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53,
	0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x31, 0x12, 0x2e, 0x0a, 0x12,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53,
	0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x32, 0x12, 0x2e, 0x0a, 0x12,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x33, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53,
	0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x33, 0x22, 0xe3, 0x01, 0x0a,
	0x09, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f,
	0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x20,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73,
	0x12, 0x2e, 0x0a, 0x12, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x31,
	0x12, 0x2e, 0x0a, 0x12, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x32,
	0x12, 0x2e, 0x0a, 0x12, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x33, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x33,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65,
	0x78, 0x74, 0x22, 0x93, 0x01, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x69, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x69, 0x64, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x6f, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x78, 0x74, 0x22, 0x57, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x6e, 0x73,
	0x61, 0x6a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x6e, 0x73, 0x61,
	0x6a, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x2a, 0x33, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x6b, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x32, 0xbe, 0x03, 0x0a, 0x17, 0x65, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x61, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x64, 0x61, 0x12, 0x30, 0x0a, 0x05, 0x53, 0x75, 0x62, 0x69, 0x72, 0x12, 0x0e, 0x2e, 0x70, 0x61,
	0x71, 0x75, 0x65, 0x74, 0x65, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x15, 0x2e, 0x70, 0x61,
	0x71, 0x75, 0x65, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0f, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x50, 0x72,
	0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12, 0x12, 0x2e, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74,
	0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x1a, 0x12, 0x2e, 0x70, 0x61,
	0x71, 0x75, 0x65, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x22,
	0x00, 0x12, 0x3f, 0x0a, 0x17, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x72, 0x45, 0x73,
	0x74, 0x61, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x12, 0x10, 0x2e, 0x70,
	0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x1a, 0x10,
	0x2e, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0b, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x1a, 0x2e, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x2e, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x10, 0x2e,
	0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x42, 0x61, 0x6a, 0x61, 0x72, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x6f, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x54, 0x6f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x12, 0x2e, 0x70, 0x61, 0x71,
	0x75, 0x65, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x0a, 0x42, 0x61, 0x6a, 0x61, 0x72, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x11,
	0x2e, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x65,
	0x73, 0x1a, 0x12, 0x2e, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x2e, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x18, 0x4f, 0x62, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x4c, 0x69, 0x62, 0x72, 0x6f, 0x73, 0x44, 0x69, 0x73, 0x70, 0x6f, 0x6e, 0x69, 0x62,
	0x6c, 0x65, 0x73, 0x12, 0x10, 0x2e, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x2e, 0x4d, 0x65,
	0x6e, 0x73, 0x61, 0x6a, 0x65, 0x1a, 0x0e, 0x2e, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x54, 0x61, 0x72,
	0x65, 0x61, 0x32, 0x53, 0x44, 0x2f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x69, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_paquete_proto_rawDescOnce sync.Once
	file_paquete_proto_rawDescData = file_paquete_proto_rawDesc
)

func file_paquete_proto_rawDescGZIP() []byte {
	file_paquete_proto_rawDescOnce.Do(func() {
		file_paquete_proto_rawDescData = protoimpl.X.CompressGZIP(file_paquete_proto_rawDescData)
	})
	return file_paquete_proto_rawDescData
}

var file_paquete_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_paquete_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_paquete_proto_goTypes = []interface{}{
	(UploadStatusCode)(0),     // 0: paquete.UploadStatusCode
	(*Books)(nil),             // 1: paquete.Books
	(*ChunkDes)(nil),          // 2: paquete.ChunkDes
	(*ChunkBook)(nil),         // 3: paquete.ChunkBook
	(*ListChunk)(nil),         // 4: paquete.ListChunk
	(*BookToDownload)(nil),    // 5: paquete.BookToDownload
	(*ChunkSendToServer)(nil), // 6: paquete.ChunkSendToServer
	(*Mensaje)(nil),           // 7: paquete.Mensaje
	(*Respuesta)(nil),         // 8: paquete.Respuesta
	(*Propuesta)(nil),         // 9: paquete.Propuesta
	(*Chunk)(nil),             // 10: paquete.Chunk
	(*UploadStatus)(nil),      // 11: paquete.UploadStatus
}
var file_paquete_proto_depIdxs = []int32{
	0,  // 0: paquete.UploadStatus.Code:type_name -> paquete.UploadStatusCode
	10, // 1: paquete.estructura_centralizada.Subir:input_type -> paquete.Chunk
	9,  // 2: paquete.estructura_centralizada.EnviarPropuesta:input_type -> paquete.Propuesta
	7,  // 3: paquete.estructura_centralizada.VerificarEstadoServidor:input_type -> paquete.Mensaje
	6,  // 4: paquete.estructura_centralizada.EnviarChunk:input_type -> paquete.ChunkSendToServer
	5,  // 5: paquete.estructura_centralizada.BajarArchivo:input_type -> paquete.BookToDownload
	2,  // 6: paquete.estructura_centralizada.BajarChunk:input_type -> paquete.ChunkDes
	7,  // 7: paquete.estructura_centralizada.ObtenerLibrosDisponibles:input_type -> paquete.Mensaje
	11, // 8: paquete.estructura_centralizada.Subir:output_type -> paquete.UploadStatus
	8,  // 9: paquete.estructura_centralizada.EnviarPropuesta:output_type -> paquete.Respuesta
	7,  // 10: paquete.estructura_centralizada.VerificarEstadoServidor:output_type -> paquete.Mensaje
	7,  // 11: paquete.estructura_centralizada.EnviarChunk:output_type -> paquete.Mensaje
	4,  // 12: paquete.estructura_centralizada.BajarArchivo:output_type -> paquete.ListChunk
	3,  // 13: paquete.estructura_centralizada.BajarChunk:output_type -> paquete.ChunkBook
	1,  // 14: paquete.estructura_centralizada.ObtenerLibrosDisponibles:output_type -> paquete.Books
	8,  // [8:15] is the sub-list for method output_type
	1,  // [1:8] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_paquete_proto_init() }
func file_paquete_proto_init() {
	if File_paquete_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_paquete_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Books); i {
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
		file_paquete_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkDes); i {
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
		file_paquete_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkBook); i {
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
		file_paquete_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChunk); i {
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
		file_paquete_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookToDownload); i {
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
		file_paquete_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChunkSendToServer); i {
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
		file_paquete_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mensaje); i {
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
		file_paquete_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Respuesta); i {
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
		file_paquete_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Propuesta); i {
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
		file_paquete_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_paquete_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadStatus); i {
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
			RawDescriptor: file_paquete_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_paquete_proto_goTypes,
		DependencyIndexes: file_paquete_proto_depIdxs,
		EnumInfos:         file_paquete_proto_enumTypes,
		MessageInfos:      file_paquete_proto_msgTypes,
	}.Build()
	File_paquete_proto = out.File
	file_paquete_proto_rawDesc = nil
	file_paquete_proto_goTypes = nil
	file_paquete_proto_depIdxs = nil
}
