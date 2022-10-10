// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: pb/cs_logic.proto

package server

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

type RegisterResult int32

const (
	RegisterResult_REG_SUCCEEDED RegisterResult = 0
	RegisterResult_REG_DUPLICATE RegisterResult = 1
)

// Enum value maps for RegisterResult.
var (
	RegisterResult_name = map[int32]string{
		0: "REG_SUCCEEDED",
		1: "REG_DUPLICATE",
	}
	RegisterResult_value = map[string]int32{
		"REG_SUCCEEDED": 0,
		"REG_DUPLICATE": 1,
	}
)

func (x RegisterResult) Enum() *RegisterResult {
	p := new(RegisterResult)
	*p = x
	return p
}

func (x RegisterResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegisterResult) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_cs_logic_proto_enumTypes[0].Descriptor()
}

func (RegisterResult) Type() protoreflect.EnumType {
	return &file_pb_cs_logic_proto_enumTypes[0]
}

func (x RegisterResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegisterResult.Descriptor instead.
func (RegisterResult) EnumDescriptor() ([]byte, []int) {
	return file_pb_cs_logic_proto_rawDescGZIP(), []int{0}
}

type LockResult int32

const (
	LockResult_LOCK_SUCCEEDED LockResult = 0
	LockResult_LOCK_FAILED    LockResult = 1
)

// Enum value maps for LockResult.
var (
	LockResult_name = map[int32]string{
		0: "LOCK_SUCCEEDED",
		1: "LOCK_FAILED",
	}
	LockResult_value = map[string]int32{
		"LOCK_SUCCEEDED": 0,
		"LOCK_FAILED":    1,
	}
)

func (x LockResult) Enum() *LockResult {
	p := new(LockResult)
	*p = x
	return p
}

func (x LockResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LockResult) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_cs_logic_proto_enumTypes[1].Descriptor()
}

func (LockResult) Type() protoreflect.EnumType {
	return &file_pb_cs_logic_proto_enumTypes[1]
}

func (x LockResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LockResult.Descriptor instead.
func (LockResult) EnumDescriptor() ([]byte, []int) {
	return file_pb_cs_logic_proto_rawDescGZIP(), []int{1}
}

// 2
type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pass string `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cs_logic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cs_logic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_pb_cs_logic_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterReq) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

// 2
type RegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result       RegisterResult `protobuf:"varint,1,opt,name=result,proto3,enum=server.RegisterResult" json:"result,omitempty"`
	ErrorMessage string         `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cs_logic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cs_logic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_pb_cs_logic_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterReply) GetResult() RegisterResult {
	if x != nil {
		return x.Result
	}
	return RegisterResult_REG_SUCCEEDED
}

func (x *RegisterReply) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

// 3
type LockNodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LockNode *Node `protobuf:"bytes,1,opt,name=lockNode,proto3" json:"lockNode,omitempty"`
}

func (x *LockNodeReq) Reset() {
	*x = LockNodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cs_logic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockNodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockNodeReq) ProtoMessage() {}

func (x *LockNodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cs_logic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockNodeReq.ProtoReflect.Descriptor instead.
func (*LockNodeReq) Descriptor() ([]byte, []int) {
	return file_pb_cs_logic_proto_rawDescGZIP(), []int{2}
}

func (x *LockNodeReq) GetLockNode() *Node {
	if x != nil {
		return x.LockNode
	}
	return nil
}

// 3
type LockNodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result       LockResult `protobuf:"varint,1,opt,name=result,proto3,enum=server.LockResult" json:"result,omitempty"`
	ErrorMessage string     `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *LockNodeReply) Reset() {
	*x = LockNodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cs_logic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockNodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockNodeReply) ProtoMessage() {}

func (x *LockNodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cs_logic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockNodeReply.ProtoReflect.Descriptor instead.
func (*LockNodeReply) Descriptor() ([]byte, []int) {
	return file_pb_cs_logic_proto_rawDescGZIP(), []int{3}
}

func (x *LockNodeReply) GetResult() LockResult {
	if x != nil {
		return x.Result
	}
	return LockResult_LOCK_SUCCEEDED
}

func (x *LockNodeReply) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

// 4
type InputTextReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputText *InputText `protobuf:"bytes,1,opt,name=inputText,proto3" json:"inputText,omitempty"`
}

func (x *InputTextReq) Reset() {
	*x = InputTextReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cs_logic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputTextReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputTextReq) ProtoMessage() {}

func (x *InputTextReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cs_logic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputTextReq.ProtoReflect.Descriptor instead.
func (*InputTextReq) Descriptor() ([]byte, []int) {
	return file_pb_cs_logic_proto_rawDescGZIP(), []int{4}
}

func (x *InputTextReq) GetInputText() *InputText {
	if x != nil {
		return x.InputText
	}
	return nil
}

// 4
type InputTextReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result       LockResult `protobuf:"varint,1,opt,name=result,proto3,enum=server.LockResult" json:"result,omitempty"`
	ErrorMessage string     `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *InputTextReply) Reset() {
	*x = InputTextReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cs_logic_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputTextReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputTextReply) ProtoMessage() {}

func (x *InputTextReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cs_logic_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputTextReply.ProtoReflect.Descriptor instead.
func (*InputTextReply) Descriptor() ([]byte, []int) {
	return file_pb_cs_logic_proto_rawDescGZIP(), []int{5}
}

func (x *InputTextReply) GetResult() LockResult {
	if x != nil {
		return x.Result
	}
	return LockResult_LOCK_SUCCEEDED
}

func (x *InputTextReply) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_pb_cs_logic_proto protoreflect.FileDescriptor

var file_pb_cs_logic_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x62, 0x2f, 0x63, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x16, 0x70, 0x62, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x73, 0x73, 0x22, 0x63, 0x0a, 0x0d, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x37, 0x0a, 0x0b, 0x4c, 0x6f, 0x63, 0x6b, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x28,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x5f, 0x0a, 0x0d, 0x4c, 0x6f, 0x63, 0x6b,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3f, 0x0a, 0x0c, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2f, 0x0a, 0x09, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52,
	0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x65, 0x78, 0x74, 0x22, 0x60, 0x0a, 0x0e, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x36, 0x0a, 0x0e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x11,
	0x0a, 0x0d, 0x52, 0x45, 0x47, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x47, 0x5f, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41,
	0x54, 0x45, 0x10, 0x01, 0x2a, 0x31, 0x0a, 0x0a, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45,
	0x45, 0x44, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_cs_logic_proto_rawDescOnce sync.Once
	file_pb_cs_logic_proto_rawDescData = file_pb_cs_logic_proto_rawDesc
)

func file_pb_cs_logic_proto_rawDescGZIP() []byte {
	file_pb_cs_logic_proto_rawDescOnce.Do(func() {
		file_pb_cs_logic_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_cs_logic_proto_rawDescData)
	})
	return file_pb_cs_logic_proto_rawDescData
}

var file_pb_cs_logic_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pb_cs_logic_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pb_cs_logic_proto_goTypes = []interface{}{
	(RegisterResult)(0),    // 0: server.RegisterResult
	(LockResult)(0),        // 1: server.LockResult
	(*RegisterReq)(nil),    // 2: server.RegisterReq
	(*RegisterReply)(nil),  // 3: server.RegisterReply
	(*LockNodeReq)(nil),    // 4: server.LockNodeReq
	(*LockNodeReply)(nil),  // 5: server.LockNodeReply
	(*InputTextReq)(nil),   // 6: server.InputTextReq
	(*InputTextReply)(nil), // 7: server.InputTextReply
	(*Node)(nil),           // 8: server.Node
	(*InputText)(nil),      // 9: server.InputText
}
var file_pb_cs_logic_proto_depIdxs = []int32{
	0, // 0: server.RegisterReply.result:type_name -> server.RegisterResult
	8, // 1: server.LockNodeReq.lockNode:type_name -> server.Node
	1, // 2: server.LockNodeReply.result:type_name -> server.LockResult
	9, // 3: server.InputTextReq.inputText:type_name -> server.InputText
	1, // 4: server.InputTextReply.result:type_name -> server.LockResult
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pb_cs_logic_proto_init() }
func file_pb_cs_logic_proto_init() {
	if File_pb_cs_logic_proto != nil {
		return
	}
	file_pb_global_define_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_cs_logic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_pb_cs_logic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReply); i {
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
		file_pb_cs_logic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockNodeReq); i {
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
		file_pb_cs_logic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockNodeReply); i {
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
		file_pb_cs_logic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputTextReq); i {
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
		file_pb_cs_logic_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputTextReply); i {
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
			RawDescriptor: file_pb_cs_logic_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_cs_logic_proto_goTypes,
		DependencyIndexes: file_pb_cs_logic_proto_depIdxs,
		EnumInfos:         file_pb_cs_logic_proto_enumTypes,
		MessageInfos:      file_pb_cs_logic_proto_msgTypes,
	}.Build()
	File_pb_cs_logic_proto = out.File
	file_pb_cs_logic_proto_rawDesc = nil
	file_pb_cs_logic_proto_goTypes = nil
	file_pb_cs_logic_proto_depIdxs = nil
}
