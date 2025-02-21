// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.2
// source: protocol/protobuf/messages.proto

package im

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

// *
// 数据类型
type PackType int32

const (
	// *
	// 登陆
	PackType_LOGIN PackType = 0
	// *
	// 退出
	PackType_LOGOUT PackType = 1
	// *
	// 响应
	PackType_RESPONSE PackType = 2
	// *
	// 消息
	PackType_MESSAGE PackType = 3
)

// Enum value maps for PackType.
var (
	PackType_name = map[int32]string{
		0: "LOGIN",
		1: "LOGOUT",
		2: "RESPONSE",
		3: "MESSAGE",
	}
	PackType_value = map[string]int32{
		"LOGIN":    0,
		"LOGOUT":   1,
		"RESPONSE": 2,
		"MESSAGE":  3,
	}
)

func (x PackType) Enum() *PackType {
	p := new(PackType)
	*p = x
	return p
}

func (x PackType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PackType) Descriptor() protoreflect.EnumDescriptor {
	return file_protocol_protobuf_messages_proto_enumTypes[0].Descriptor()
}

func (PackType) Type() protoreflect.EnumType {
	return &file_protocol_protobuf_messages_proto_enumTypes[0]
}

func (x PackType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PackType.Descriptor instead.
func (PackType) EnumDescriptor() ([]byte, []int) {
	return file_protocol_protobuf_messages_proto_rawDescGZIP(), []int{0}
}

// *
// 聊天类型
type ChatType int32

const (
	// *
	// 单聊消息
	ChatType_CHAT ChatType = 0
	// *
	// 群聊消息
	ChatType_GROUP_CHAT ChatType = 1
	// *
	// 系统消息
	ChatType_SYSTEM ChatType = 2
)

// Enum value maps for ChatType.
var (
	ChatType_name = map[int32]string{
		0: "CHAT",
		1: "GROUP_CHAT",
		2: "SYSTEM",
	}
	ChatType_value = map[string]int32{
		"CHAT":       0,
		"GROUP_CHAT": 1,
		"SYSTEM":     2,
	}
)

func (x ChatType) Enum() *ChatType {
	p := new(ChatType)
	*p = x
	return p
}

func (x ChatType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatType) Descriptor() protoreflect.EnumDescriptor {
	return file_protocol_protobuf_messages_proto_enumTypes[1].Descriptor()
}

func (ChatType) Type() protoreflect.EnumType {
	return &file_protocol_protobuf_messages_proto_enumTypes[1]
}

func (x ChatType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatType.Descriptor instead.
func (ChatType) EnumDescriptor() ([]byte, []int) {
	return file_protocol_protobuf_messages_proto_rawDescGZIP(), []int{1}
}

// *
// 消息类型
type MessageType int32

const (
	// *
	// 文本
	MessageType_TEXT MessageType = 0
	// *
	// 图片
	MessageType_IMAGE MessageType = 1
	// *
	// 语音
	MessageType_AUDIO MessageType = 2
	// *
	// 视频
	MessageType_VEDIO MessageType = 3
	// *
	// 自定义消息
	MessageType_CUSTOM MessageType = 4
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "TEXT",
		1: "IMAGE",
		2: "AUDIO",
		3: "VEDIO",
		4: "CUSTOM",
	}
	MessageType_value = map[string]int32{
		"TEXT":   0,
		"IMAGE":  1,
		"AUDIO":  2,
		"VEDIO":  3,
		"CUSTOM": 4,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_protocol_protobuf_messages_proto_enumTypes[2].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_protocol_protobuf_messages_proto_enumTypes[2]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_protocol_protobuf_messages_proto_rawDescGZIP(), []int{2}
}

// *
// 一包数据
type PackData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                       // 消息id
	From      uint32   `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`                  // 发送者
	To        uint32   `protobuf:"varint,3,opt,name=to,proto3" json:"to,omitempty"`                      // 接收者
	Type      PackType `protobuf:"varint,4,opt,name=type,proto3,enum=im.PackType" json:"type,omitempty"` // 消息类型
	Payload   []byte   `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`             // 消息内容
	Timestamp uint64   `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`        // 时间戳
}

func (x *PackData) Reset() {
	*x = PackData{}
	mi := &file_protocol_protobuf_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PackData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackData) ProtoMessage() {}

func (x *PackData) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_protobuf_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackData.ProtoReflect.Descriptor instead.
func (*PackData) Descriptor() ([]byte, []int) {
	return file_protocol_protobuf_messages_proto_rawDescGZIP(), []int{0}
}

func (x *PackData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PackData) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *PackData) GetTo() uint32 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *PackData) GetType() PackType {
	if x != nil {
		return x.Type
	}
	return PackType_LOGIN
}

func (x *PackData) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *PackData) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// *
// 登陆请求体
type LoginPack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid        uint32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ClientType uint32 `protobuf:"varint,2,opt,name=clientType,proto3" json:"clientType,omitempty"`
	Password   string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginPack) Reset() {
	*x = LoginPack{}
	mi := &file_protocol_protobuf_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginPack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginPack) ProtoMessage() {}

func (x *LoginPack) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_protobuf_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginPack.ProtoReflect.Descriptor instead.
func (*LoginPack) Descriptor() ([]byte, []int) {
	return file_protocol_protobuf_messages_proto_rawDescGZIP(), []int{1}
}

func (x *LoginPack) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *LoginPack) GetClientType() uint32 {
	if x != nil {
		return x.ClientType
	}
	return 0
}

func (x *LoginPack) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// *
// 退出登录请求体
type LogoutPack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *LogoutPack) Reset() {
	*x = LogoutPack{}
	mi := &file_protocol_protobuf_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutPack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutPack) ProtoMessage() {}

func (x *LogoutPack) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_protobuf_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutPack.ProtoReflect.Descriptor instead.
func (*LogoutPack) Descriptor() ([]byte, []int) {
	return file_protocol_protobuf_messages_proto_rawDescGZIP(), []int{2}
}

func (x *LogoutPack) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LogoutPack) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

// *
// 通用相应体
type ResponsePack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ResponsePack) Reset() {
	*x = ResponsePack{}
	mi := &file_protocol_protobuf_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponsePack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsePack) ProtoMessage() {}

func (x *ResponsePack) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_protobuf_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsePack.ProtoReflect.Descriptor instead.
func (*ResponsePack) Descriptor() ([]byte, []int) {
	return file_protocol_protobuf_messages_proto_rawDescGZIP(), []int{3}
}

func (x *ResponsePack) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResponsePack) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

var File_protocol_protobuf_messages_proto protoreflect.FileDescriptor

var file_protocol_protobuf_messages_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x69, 0x6d, 0x22, 0x98, 0x01, 0x0a, 0x08, 0x50, 0x61, 0x63, 0x6b, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x69, 0x6d, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x59, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x63, 0x6b, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x38, 0x0a, 0x0a,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x3c, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x2a, 0x3c, 0x0a, 0x08, 0x50, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4c,
	0x4f, 0x47, 0x4f, 0x55, 0x54, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53, 0x50, 0x4f,
	0x4e, 0x53, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45,
	0x10, 0x03, 0x2a, 0x30, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x43, 0x48, 0x41, 0x54, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x52, 0x4f, 0x55,
	0x50, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x59, 0x53, 0x54,
	0x45, 0x4d, 0x10, 0x02, 0x2a, 0x44, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x55, 0x44, 0x49,
	0x4f, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x45, 0x44, 0x49, 0x4f, 0x10, 0x03, 0x12, 0x0a,
	0x0a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x04, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x69, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_protobuf_messages_proto_rawDescOnce sync.Once
	file_protocol_protobuf_messages_proto_rawDescData = file_protocol_protobuf_messages_proto_rawDesc
)

func file_protocol_protobuf_messages_proto_rawDescGZIP() []byte {
	file_protocol_protobuf_messages_proto_rawDescOnce.Do(func() {
		file_protocol_protobuf_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_protobuf_messages_proto_rawDescData)
	})
	return file_protocol_protobuf_messages_proto_rawDescData
}

var file_protocol_protobuf_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_protocol_protobuf_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protocol_protobuf_messages_proto_goTypes = []any{
	(PackType)(0),        // 0: im.PackType
	(ChatType)(0),        // 1: im.ChatType
	(MessageType)(0),     // 2: im.MessageType
	(*PackData)(nil),     // 3: im.PackData
	(*LoginPack)(nil),    // 4: im.LoginPack
	(*LogoutPack)(nil),   // 5: im.LogoutPack
	(*ResponsePack)(nil), // 6: im.ResponsePack
}
var file_protocol_protobuf_messages_proto_depIdxs = []int32{
	0, // 0: im.PackData.type:type_name -> im.PackType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protocol_protobuf_messages_proto_init() }
func file_protocol_protobuf_messages_proto_init() {
	if File_protocol_protobuf_messages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protocol_protobuf_messages_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_protobuf_messages_proto_goTypes,
		DependencyIndexes: file_protocol_protobuf_messages_proto_depIdxs,
		EnumInfos:         file_protocol_protobuf_messages_proto_enumTypes,
		MessageInfos:      file_protocol_protobuf_messages_proto_msgTypes,
	}.Build()
	File_protocol_protobuf_messages_proto = out.File
	file_protocol_protobuf_messages_proto_rawDesc = nil
	file_protocol_protobuf_messages_proto_goTypes = nil
	file_protocol_protobuf_messages_proto_depIdxs = nil
}
