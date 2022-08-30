// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.13.0
// source: proto/backupAgent/esbak/es_history.proto

package esbak

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

type GetEsHistoryListInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info     string `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
	PageNo   int64  `protobuf:"varint,2,opt,name=PageNo,proto3" json:"PageNo,omitempty"`
	PageSize int64  `protobuf:"varint,3,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	Sort     string `protobuf:"bytes,4,opt,name=Sort,proto3" json:"Sort,omitempty"`
}

func (x *GetEsHistoryListInput) Reset() {
	*x = GetEsHistoryListInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_backupAgent_esbak_es_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEsHistoryListInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEsHistoryListInput) ProtoMessage() {}

func (x *GetEsHistoryListInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backupAgent_esbak_es_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEsHistoryListInput.ProtoReflect.Descriptor instead.
func (*GetEsHistoryListInput) Descriptor() ([]byte, []int) {
	return file_proto_backupAgent_esbak_es_history_proto_rawDescGZIP(), []int{0}
}

func (x *GetEsHistoryListInput) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *GetEsHistoryListInput) GetPageNo() int64 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *GetEsHistoryListInput) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetEsHistoryListInput) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

type ESHistoryIDInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ESHistoryIDInput) Reset() {
	*x = ESHistoryIDInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_backupAgent_esbak_es_history_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ESHistoryIDInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ESHistoryIDInput) ProtoMessage() {}

func (x *ESHistoryIDInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backupAgent_esbak_es_history_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ESHistoryIDInput.ProtoReflect.Descriptor instead.
func (*ESHistoryIDInput) Descriptor() ([]byte, []int) {
	return file_proto_backupAgent_esbak_es_history_proto_rawDescGZIP(), []int{1}
}

func (x *ESHistoryIDInput) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type ESHistoryOneMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	OK      bool   `protobuf:"varint,2,opt,name=OK,proto3" json:"OK,omitempty"`
}

func (x *ESHistoryOneMessage) Reset() {
	*x = ESHistoryOneMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_backupAgent_esbak_es_history_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ESHistoryOneMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ESHistoryOneMessage) ProtoMessage() {}

func (x *ESHistoryOneMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backupAgent_esbak_es_history_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ESHistoryOneMessage.ProtoReflect.Descriptor instead.
func (*ESHistoryOneMessage) Descriptor() ([]byte, []int) {
	return file_proto_backupAgent_esbak_es_history_proto_rawDescGZIP(), []int{2}
}

func (x *ESHistoryOneMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ESHistoryOneMessage) GetOK() bool {
	if x != nil {
		return x.OK
	}
	return false
}

type ESHistoryListOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total                int64                   `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	EsHistoryListOutItem []*ESHistoryListOutItem `protobuf:"bytes,2,rep,name=esHistoryListOutItem,proto3" json:"esHistoryListOutItem,omitempty"`
}

func (x *ESHistoryListOutput) Reset() {
	*x = ESHistoryListOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_backupAgent_esbak_es_history_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ESHistoryListOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ESHistoryListOutput) ProtoMessage() {}

func (x *ESHistoryListOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backupAgent_esbak_es_history_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ESHistoryListOutput.ProtoReflect.Descriptor instead.
func (*ESHistoryListOutput) Descriptor() ([]byte, []int) {
	return file_proto_backupAgent_esbak_es_history_proto_rawDescGZIP(), []int{3}
}

func (x *ESHistoryListOutput) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ESHistoryListOutput) GetEsHistoryListOutItem() []*ESHistoryListOutItem {
	if x != nil {
		return x.EsHistoryListOutItem
	}
	return nil
}

type ESHistoryListOutItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID               int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TaskID           int64  `protobuf:"varint,9,opt,name=TaskID,proto3" json:"TaskID,omitempty"`
	UUID             string `protobuf:"bytes,10,opt,name=UUID,proto3" json:"UUID,omitempty"`
	DurationInMillis int64  `protobuf:"varint,11,opt,name=DurationInMillis,proto3" json:"DurationInMillis,omitempty"`
	Snapshot         string `protobuf:"bytes,2,opt,name=Snapshot,proto3" json:"Snapshot,omitempty"`
	Repository       string `protobuf:"bytes,3,opt,name=Repository,proto3" json:"Repository,omitempty"`
	Indices          string `protobuf:"bytes,4,opt,name=Indices,proto3" json:"Indices,omitempty"`
	State            string `protobuf:"bytes,5,opt,name=State,proto3" json:"State,omitempty"`
	StartTime        string `protobuf:"bytes,6,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime          string `protobuf:"bytes,7,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	Message          string `protobuf:"bytes,8,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *ESHistoryListOutItem) Reset() {
	*x = ESHistoryListOutItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_backupAgent_esbak_es_history_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ESHistoryListOutItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ESHistoryListOutItem) ProtoMessage() {}

func (x *ESHistoryListOutItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backupAgent_esbak_es_history_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ESHistoryListOutItem.ProtoReflect.Descriptor instead.
func (*ESHistoryListOutItem) Descriptor() ([]byte, []int) {
	return file_proto_backupAgent_esbak_es_history_proto_rawDescGZIP(), []int{4}
}

func (x *ESHistoryListOutItem) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ESHistoryListOutItem) GetTaskID() int64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

func (x *ESHistoryListOutItem) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *ESHistoryListOutItem) GetDurationInMillis() int64 {
	if x != nil {
		return x.DurationInMillis
	}
	return 0
}

func (x *ESHistoryListOutItem) GetSnapshot() string {
	if x != nil {
		return x.Snapshot
	}
	return ""
}

func (x *ESHistoryListOutItem) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *ESHistoryListOutItem) GetIndices() string {
	if x != nil {
		return x.Indices
	}
	return ""
}

func (x *ESHistoryListOutItem) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *ESHistoryListOutItem) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *ESHistoryListOutItem) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *ESHistoryListOutItem) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_backupAgent_esbak_es_history_proto protoreflect.FileDescriptor

var file_proto_backupAgent_esbak_es_history_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x2f, 0x65, 0x73, 0x62, 0x61, 0x6b, 0x2f, 0x65, 0x73, 0x5f, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x73, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x45,
	0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6f, 0x72,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x22, 0x22, 0x0a,
	0x10, 0x45, 0x53, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x44, 0x22, 0x3f, 0x0a, 0x13, 0x45, 0x53, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4f, 0x6e,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x4b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02,
	0x4f, 0x4b, 0x22, 0x93, 0x01, 0x0a, 0x13, 0x45, 0x53, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x66, 0x0a, 0x14, 0x65, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x75, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x53,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x14, 0x65, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x75, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x22, 0xbc, 0x02, 0x0a, 0x14, 0x45, 0x53, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x2a, 0x0a,
	0x10, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x4d, 0x69, 0x6c, 0x6c, 0x69,
	0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x88, 0x02, 0x0a, 0x10, 0x45, 0x73, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7c, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x45, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x33, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x45, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x53, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x76, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x53, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2e, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x53, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x31, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x53, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4f, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x45, 0x73, 0x42, 0x61, 0x6b, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_backupAgent_esbak_es_history_proto_rawDescOnce sync.Once
	file_proto_backupAgent_esbak_es_history_proto_rawDescData = file_proto_backupAgent_esbak_es_history_proto_rawDesc
)

func file_proto_backupAgent_esbak_es_history_proto_rawDescGZIP() []byte {
	file_proto_backupAgent_esbak_es_history_proto_rawDescOnce.Do(func() {
		file_proto_backupAgent_esbak_es_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_backupAgent_esbak_es_history_proto_rawDescData)
	})
	return file_proto_backupAgent_esbak_es_history_proto_rawDescData
}

var file_proto_backupAgent_esbak_es_history_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_backupAgent_esbak_es_history_proto_goTypes = []interface{}{
	(*GetEsHistoryListInput)(nil), // 0: go.micro.service.backupAgent.GetEsHistoryListInput
	(*ESHistoryIDInput)(nil),      // 1: go.micro.service.backupAgent.ESHistoryIDInput
	(*ESHistoryOneMessage)(nil),   // 2: go.micro.service.backupAgent.ESHistoryOneMessage
	(*ESHistoryListOutput)(nil),   // 3: go.micro.service.backupAgent.ESHistoryListOutput
	(*ESHistoryListOutItem)(nil),  // 4: go.micro.service.backupAgent.ESHistoryListOutItem
}
var file_proto_backupAgent_esbak_es_history_proto_depIdxs = []int32{
	4, // 0: go.micro.service.backupAgent.ESHistoryListOutput.esHistoryListOutItem:type_name -> go.micro.service.backupAgent.ESHistoryListOutItem
	0, // 1: go.micro.service.backupAgent.EsHistoryService.GetEsHistoryList:input_type -> go.micro.service.backupAgent.GetEsHistoryListInput
	1, // 2: go.micro.service.backupAgent.EsHistoryService.DeleteESHistory:input_type -> go.micro.service.backupAgent.ESHistoryIDInput
	3, // 3: go.micro.service.backupAgent.EsHistoryService.GetEsHistoryList:output_type -> go.micro.service.backupAgent.ESHistoryListOutput
	2, // 4: go.micro.service.backupAgent.EsHistoryService.DeleteESHistory:output_type -> go.micro.service.backupAgent.ESHistoryOneMessage
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_backupAgent_esbak_es_history_proto_init() }
func file_proto_backupAgent_esbak_es_history_proto_init() {
	if File_proto_backupAgent_esbak_es_history_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_backupAgent_esbak_es_history_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEsHistoryListInput); i {
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
		file_proto_backupAgent_esbak_es_history_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ESHistoryIDInput); i {
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
		file_proto_backupAgent_esbak_es_history_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ESHistoryOneMessage); i {
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
		file_proto_backupAgent_esbak_es_history_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ESHistoryListOutput); i {
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
		file_proto_backupAgent_esbak_es_history_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ESHistoryListOutItem); i {
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
			RawDescriptor: file_proto_backupAgent_esbak_es_history_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_backupAgent_esbak_es_history_proto_goTypes,
		DependencyIndexes: file_proto_backupAgent_esbak_es_history_proto_depIdxs,
		MessageInfos:      file_proto_backupAgent_esbak_es_history_proto_msgTypes,
	}.Build()
	File_proto_backupAgent_esbak_es_history_proto = out.File
	file_proto_backupAgent_esbak_es_history_proto_rawDesc = nil
	file_proto_backupAgent_esbak_es_history_proto_goTypes = nil
	file_proto_backupAgent_esbak_es_history_proto_depIdxs = nil
}
