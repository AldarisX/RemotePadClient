// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: proto/cmd.proto

package common

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

type CmdType int32

const (
	CmdType_T_Hello    CmdType = 0
	CmdType_T_Ping     CmdType = 1
	CmdType_T_Pad_Type CmdType = 2
	CmdType_T_Pad_Data CmdType = 3
)

// Enum value maps for CmdType.
var (
	CmdType_name = map[int32]string{
		0: "T_Hello",
		1: "T_Ping",
		2: "T_Pad_Type",
		3: "T_Pad_Data",
	}
	CmdType_value = map[string]int32{
		"T_Hello":    0,
		"T_Ping":     1,
		"T_Pad_Type": 2,
		"T_Pad_Data": 3,
	}
)

func (x CmdType) Enum() *CmdType {
	p := new(CmdType)
	*p = x
	return p
}

func (x CmdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CmdType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_cmd_proto_enumTypes[0].Descriptor()
}

func (CmdType) Type() protoreflect.EnumType {
	return &file_proto_cmd_proto_enumTypes[0]
}

func (x CmdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CmdType.Descriptor instead.
func (CmdType) EnumDescriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{0}
}

type MsgType int32

const (
	MsgType_Server MsgType = 0
	MsgType_Driver MsgType = 1
	MsgType_Pad    MsgType = 2
)

// Enum value maps for MsgType.
var (
	MsgType_name = map[int32]string{
		0: "Server",
		1: "Driver",
		2: "Pad",
	}
	MsgType_value = map[string]int32{
		"Server": 0,
		"Driver": 1,
		"Pad":    2,
	}
)

func (x MsgType) Enum() *MsgType {
	p := new(MsgType)
	*p = x
	return p
}

func (x MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_cmd_proto_enumTypes[1].Descriptor()
}

func (MsgType) Type() protoreflect.EnumType {
	return &file_proto_cmd_proto_enumTypes[1]
}

func (x MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgType.Descriptor instead.
func (MsgType) EnumDescriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{1}
}

type PadType int32

const (
	PadType_DS4     PadType = 0
	PadType_Xbox360 PadType = 1
)

// Enum value maps for PadType.
var (
	PadType_name = map[int32]string{
		0: "DS4",
		1: "Xbox360",
	}
	PadType_value = map[string]int32{
		"DS4":     0,
		"Xbox360": 1,
	}
)

func (x PadType) Enum() *PadType {
	p := new(PadType)
	*p = x
	return p
}

func (x PadType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PadType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_cmd_proto_enumTypes[2].Descriptor()
}

func (PadType) Type() protoreflect.EnumType {
	return &file_proto_cmd_proto_enumTypes[2]
}

func (x PadType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PadType.Descriptor instead.
func (PadType) EnumDescriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{2}
}

type PadBtn int32

const (
	PadBtn_A           PadBtn = 0
	PadBtn_B           PadBtn = 1
	PadBtn_Y           PadBtn = 2
	PadBtn_X           PadBtn = 3
	PadBtn_LB          PadBtn = 4
	PadBtn_RB          PadBtn = 5
	PadBtn_L3          PadBtn = 6
	PadBtn_R3          PadBtn = 7
	PadBtn_Start       PadBtn = 8
	PadBtn_Select      PadBtn = 9
	PadBtn_Xbox        PadBtn = 10
	PadBtn_L2          PadBtn = 11
	PadBtn_R2          PadBtn = 12
	PadBtn_LAX         PadBtn = 13
	PadBtn_LAY         PadBtn = 14
	PadBtn_RAX         PadBtn = 15
	PadBtn_RAY         PadBtn = 16
	PadBtn_DUP         PadBtn = 17
	PadBtn_DDOWN       PadBtn = 18
	PadBtn_DLEFT       PadBtn = 19
	PadBtn_DRIGHT      PadBtn = 20
	PadBtn_DS4TouchPad PadBtn = 21
)

// Enum value maps for PadBtn.
var (
	PadBtn_name = map[int32]string{
		0:  "A",
		1:  "B",
		2:  "Y",
		3:  "X",
		4:  "LB",
		5:  "RB",
		6:  "L3",
		7:  "R3",
		8:  "Start",
		9:  "Select",
		10: "Xbox",
		11: "L2",
		12: "R2",
		13: "LAX",
		14: "LAY",
		15: "RAX",
		16: "RAY",
		17: "DUP",
		18: "DDOWN",
		19: "DLEFT",
		20: "DRIGHT",
		21: "DS4TouchPad",
	}
	PadBtn_value = map[string]int32{
		"A":           0,
		"B":           1,
		"Y":           2,
		"X":           3,
		"LB":          4,
		"RB":          5,
		"L3":          6,
		"R3":          7,
		"Start":       8,
		"Select":      9,
		"Xbox":        10,
		"L2":          11,
		"R2":          12,
		"LAX":         13,
		"LAY":         14,
		"RAX":         15,
		"RAY":         16,
		"DUP":         17,
		"DDOWN":       18,
		"DLEFT":       19,
		"DRIGHT":      20,
		"DS4TouchPad": 21,
	}
)

func (x PadBtn) Enum() *PadBtn {
	p := new(PadBtn)
	*p = x
	return p
}

func (x PadBtn) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PadBtn) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_cmd_proto_enumTypes[3].Descriptor()
}

func (PadBtn) Type() protoreflect.EnumType {
	return &file_proto_cmd_proto_enumTypes[3]
}

func (x PadBtn) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PadBtn.Descriptor instead.
func (PadBtn) EnumDescriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{3}
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd     CmdType     `protobuf:"varint,1,opt,name=cmd,proto3,enum=common.CmdType" json:"cmd,omitempty"`
	Id      string      `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	MsgType MsgType     `protobuf:"varint,3,opt,name=msgType,proto3,enum=common.MsgType" json:"msgType,omitempty"`
	To      string      `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Hello   *Hello      `protobuf:"bytes,5,opt,name=hello,proto3" json:"hello,omitempty"`
	Ping    *Ping       `protobuf:"bytes,6,opt,name=ping,proto3" json:"ping,omitempty"`
	PadType PadType     `protobuf:"varint,7,opt,name=padType,proto3,enum=common.PadType" json:"padType,omitempty"`
	PadData *PadBtnData `protobuf:"bytes,8,opt,name=padData,proto3" json:"padData,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cmd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cmd_proto_msgTypes[0]
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
	return file_proto_cmd_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetCmd() CmdType {
	if x != nil {
		return x.Cmd
	}
	return CmdType_T_Hello
}

func (x *Data) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Data) GetMsgType() MsgType {
	if x != nil {
		return x.MsgType
	}
	return MsgType_Server
}

func (x *Data) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Data) GetHello() *Hello {
	if x != nil {
		return x.Hello
	}
	return nil
}

func (x *Data) GetPing() *Ping {
	if x != nil {
		return x.Ping
	}
	return nil
}

func (x *Data) GetPadType() PadType {
	if x != nil {
		return x.PadType
	}
	return PadType_DS4
}

func (x *Data) GetPadData() *PadBtnData {
	if x != nil {
		return x.PadData
	}
	return nil
}

type Hello struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group     string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Msg       string `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
	ServerMsg string `protobuf:"bytes,5,opt,name=serverMsg,proto3" json:"serverMsg,omitempty"`
	Order     bool   `protobuf:"varint,6,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *Hello) Reset() {
	*x = Hello{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cmd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hello) ProtoMessage() {}

func (x *Hello) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cmd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hello.ProtoReflect.Descriptor instead.
func (*Hello) Descriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{1}
}

func (x *Hello) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Hello) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Hello) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Hello) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Hello) GetServerMsg() string {
	if x != nil {
		return x.ServerMsg
	}
	return ""
}

func (x *Hello) GetOrder() bool {
	if x != nil {
		return x.Order
	}
	return false
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time int64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cmd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cmd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{2}
}

func (x *Ping) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type PadBtnData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BtnType PadBtn `protobuf:"varint,1,opt,name=btnType,proto3,enum=common.PadBtn" json:"btnType,omitempty"`
	BtnVal  int32  `protobuf:"varint,2,opt,name=btnVal,proto3" json:"btnVal,omitempty"`
}

func (x *PadBtnData) Reset() {
	*x = PadBtnData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cmd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PadBtnData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PadBtnData) ProtoMessage() {}

func (x *PadBtnData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cmd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PadBtnData.ProtoReflect.Descriptor instead.
func (*PadBtnData) Descriptor() ([]byte, []int) {
	return file_proto_cmd_proto_rawDescGZIP(), []int{3}
}

func (x *PadBtnData) GetBtnType() PadBtn {
	if x != nil {
		return x.BtnType
	}
	return PadBtn_A
}

func (x *PadBtnData) GetBtnVal() int32 {
	if x != nil {
		return x.BtnVal
	}
	return 0
}

var File_proto_cmd_proto protoreflect.FileDescriptor

var file_proto_cmd_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x94, 0x02, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x23, 0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x05,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x20, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x61, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x70, 0x61, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x70, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x64,
	0x42, 0x74, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x70, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x87, 0x01, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x1a, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x4e, 0x0a, 0x0a, 0x50, 0x61, 0x64, 0x42, 0x74, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x07, 0x62, 0x74, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x64, 0x42, 0x74, 0x6e, 0x52, 0x07, 0x62, 0x74, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x74, 0x6e, 0x56, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x62, 0x74, 0x6e, 0x56, 0x61, 0x6c, 0x2a, 0x42, 0x0a, 0x07, 0x43, 0x6d, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x5f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x54, 0x5f, 0x50, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x5f,
	0x50, 0x61, 0x64, 0x5f, 0x54, 0x79, 0x70, 0x65, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x5f,
	0x50, 0x61, 0x64, 0x5f, 0x44, 0x61, 0x74, 0x61, 0x10, 0x03, 0x2a, 0x2a, 0x0a, 0x07, 0x4d, 0x73,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x10, 0x01, 0x12, 0x07, 0x0a,
	0x03, 0x50, 0x61, 0x64, 0x10, 0x02, 0x2a, 0x1f, 0x0a, 0x07, 0x50, 0x61, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x53, 0x34, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x58, 0x62,
	0x6f, 0x78, 0x33, 0x36, 0x30, 0x10, 0x01, 0x2a, 0xd5, 0x01, 0x0a, 0x06, 0x50, 0x61, 0x64, 0x42,
	0x74, 0x6e, 0x12, 0x05, 0x0a, 0x01, 0x41, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x42, 0x10, 0x01,
	0x12, 0x05, 0x0a, 0x01, 0x59, 0x10, 0x02, 0x12, 0x05, 0x0a, 0x01, 0x58, 0x10, 0x03, 0x12, 0x06,
	0x0a, 0x02, 0x4c, 0x42, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x52, 0x42, 0x10, 0x05, 0x12, 0x06,
	0x0a, 0x02, 0x4c, 0x33, 0x10, 0x06, 0x12, 0x06, 0x0a, 0x02, 0x52, 0x33, 0x10, 0x07, 0x12, 0x09,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x10, 0x09, 0x12, 0x08, 0x0a, 0x04, 0x58, 0x62, 0x6f, 0x78, 0x10, 0x0a, 0x12,
	0x06, 0x0a, 0x02, 0x4c, 0x32, 0x10, 0x0b, 0x12, 0x06, 0x0a, 0x02, 0x52, 0x32, 0x10, 0x0c, 0x12,
	0x07, 0x0a, 0x03, 0x4c, 0x41, 0x58, 0x10, 0x0d, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x41, 0x59, 0x10,
	0x0e, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x41, 0x58, 0x10, 0x0f, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x41,
	0x59, 0x10, 0x10, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x55, 0x50, 0x10, 0x11, 0x12, 0x09, 0x0a, 0x05,
	0x44, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x12, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x4c, 0x45, 0x46, 0x54,
	0x10, 0x13, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x14, 0x12, 0x0f,
	0x0a, 0x0b, 0x44, 0x53, 0x34, 0x54, 0x6f, 0x75, 0x63, 0x68, 0x50, 0x61, 0x64, 0x10, 0x15, 0x42,
	0x2f, 0x0a, 0x12, 0x63, 0x6e, 0x2e, 0x6d, 0x69, 0x73, 0x61, 0x6b, 0x61, 0x6e, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0xaa, 0x02,
	0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x64, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cmd_proto_rawDescOnce sync.Once
	file_proto_cmd_proto_rawDescData = file_proto_cmd_proto_rawDesc
)

func file_proto_cmd_proto_rawDescGZIP() []byte {
	file_proto_cmd_proto_rawDescOnce.Do(func() {
		file_proto_cmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cmd_proto_rawDescData)
	})
	return file_proto_cmd_proto_rawDescData
}

var file_proto_cmd_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_proto_cmd_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_cmd_proto_goTypes = []interface{}{
	(CmdType)(0),       // 0: common.CmdType
	(MsgType)(0),       // 1: common.MsgType
	(PadType)(0),       // 2: common.PadType
	(PadBtn)(0),        // 3: common.PadBtn
	(*Data)(nil),       // 4: common.Data
	(*Hello)(nil),      // 5: common.Hello
	(*Ping)(nil),       // 6: common.Ping
	(*PadBtnData)(nil), // 7: common.PadBtnData
}
var file_proto_cmd_proto_depIdxs = []int32{
	0, // 0: common.Data.cmd:type_name -> common.CmdType
	1, // 1: common.Data.msgType:type_name -> common.MsgType
	5, // 2: common.Data.hello:type_name -> common.Hello
	6, // 3: common.Data.ping:type_name -> common.Ping
	2, // 4: common.Data.padType:type_name -> common.PadType
	7, // 5: common.Data.padData:type_name -> common.PadBtnData
	3, // 6: common.PadBtnData.btnType:type_name -> common.PadBtn
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_cmd_proto_init() }
func file_proto_cmd_proto_init() {
	if File_proto_cmd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cmd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_cmd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hello); i {
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
		file_proto_cmd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_proto_cmd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PadBtnData); i {
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
			RawDescriptor: file_proto_cmd_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_cmd_proto_goTypes,
		DependencyIndexes: file_proto_cmd_proto_depIdxs,
		EnumInfos:         file_proto_cmd_proto_enumTypes,
		MessageInfos:      file_proto_cmd_proto_msgTypes,
	}.Build()
	File_proto_cmd_proto = out.File
	file_proto_cmd_proto_rawDesc = nil
	file_proto_cmd_proto_goTypes = nil
	file_proto_cmd_proto_depIdxs = nil
}
