// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.13.0
// source: proto/backupAgent/bakhistory/bakhistory.proto

package bakhistory

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

type HistoryListInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"info" form:"info" uri:"info"
	Info string `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
	// @inject_tag: json:"page_no" form:"page_no" uri:"page_no"
	PageNo int64 `protobuf:"varint,2,opt,name=PageNo,proto3" json:"PageNo,omitempty"`
	// @inject_tag: json:"page_size" form:"page_size" uri:"page_size"
	PageSize int64  `protobuf:"varint,3,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	Sort     string `protobuf:"bytes,4,opt,name=Sort,proto3" json:"Sort,omitempty"`
}

func (x *HistoryListInput) Reset() {
	*x = HistoryListInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_backupAgent_bakhistory_bakhistory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryListInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryListInput) ProtoMessage() {}

func (x *HistoryListInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backupAgent_bakhistory_bakhistory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryListInput.ProtoReflect.Descriptor instead.
func (*HistoryListInput) Descriptor() ([]byte, []int) {
	return file_proto_backupAgent_bakhistory_bakhistory_proto_rawDescGZIP(), []int{0}
}

func (x *HistoryListInput) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *HistoryListInput) GetPageNo() int64 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *HistoryListInput) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *HistoryListInput) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

type HistoryIDInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *HistoryIDInput) Reset() {
	*x = HistoryIDInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_backupAgent_bakhistory_bakhistory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryIDInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryIDInput) ProtoMessage() {}

func (x *HistoryIDInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backupAgent_bakhistory_bakhistory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryIDInput.ProtoReflect.Descriptor instead.
func (*HistoryIDInput) Descriptor() ([]byte, []int) {
	return file_proto_backupAgent_bakhistory_bakhistory_proto_rawDescGZIP(), []int{1}
}

func (x *HistoryIDInput) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type HistoryOneMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"message" form:"message" uri:"message"
	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	// @inject_tag: json:"content" form:"content" uri:"content"
	OK bool `protobuf:"varint,2,opt,name=OK,proto3" json:"OK,omitempty"`
}

func (x *HistoryOneMessage) Reset() {
	*x = HistoryOneMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_backupAgent_bakhistory_bakhistory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryOneMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryOneMessage) ProtoMessage() {}

func (x *HistoryOneMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backupAgent_bakhistory_bakhistory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryOneMessage.ProtoReflect.Descriptor instead.
func (*HistoryOneMessage) Descriptor() ([]byte, []int) {
	return file_proto_backupAgent_bakhistory_bakhistory_proto_rawDescGZIP(), []int{2}
}

func (x *HistoryOneMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HistoryOneMessage) GetOK() bool {
	if x != nil {
		return x.OK
	}
	return false
}

type HistoryListOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total              int64                 `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	HistoryListOutItem []*HistoryListOutItem `protobuf:"bytes,2,rep,name=historyListOutItem,proto3" json:"historyListOutItem,omitempty"`
}

func (x *HistoryListOutput) Reset() {
	*x = HistoryListOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_backupAgent_bakhistory_bakhistory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryListOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryListOutput) ProtoMessage() {}

func (x *HistoryListOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backupAgent_bakhistory_bakhistory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryListOutput.ProtoReflect.Descriptor instead.
func (*HistoryListOutput) Descriptor() ([]byte, []int) {
	return file_proto_backupAgent_bakhistory_bakhistory_proto_rawDescGZIP(), []int{3}
}

func (x *HistoryListOutput) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *HistoryListOutput) GetHistoryListOutItem() []*HistoryListOutItem {
	if x != nil {
		return x.HistoryListOutItem
	}
	return nil
}

type HistoryListOutItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Host       string `protobuf:"bytes,2,opt,name=Host,proto3" json:"Host,omitempty"`
	DBName     string `protobuf:"bytes,3,opt,name=DBName,proto3" json:"DBName,omitempty"`
	DingStatus int64  `protobuf:"varint,4,opt,name=DingStatus,proto3" json:"DingStatus,omitempty"`
	OSSStatus  int64  `protobuf:"varint,5,opt,name=OSSStatus,proto3" json:"OSSStatus,omitempty"`
	Message    string `protobuf:"bytes,6,opt,name=Message,proto3" json:"Message,omitempty"`
	FileSize   string `protobuf:"bytes,7,opt,name=FileSize,proto3" json:"FileSize,omitempty"`
	FileName   string `protobuf:"bytes,8,opt,name=FileName,proto3" json:"FileName,omitempty"`
	BakTime    string `protobuf:"bytes,9,opt,name=BakTime,proto3" json:"BakTime,omitempty"`
}

func (x *HistoryListOutItem) Reset() {
	*x = HistoryListOutItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_backupAgent_bakhistory_bakhistory_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryListOutItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryListOutItem) ProtoMessage() {}

func (x *HistoryListOutItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backupAgent_bakhistory_bakhistory_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryListOutItem.ProtoReflect.Descriptor instead.
func (*HistoryListOutItem) Descriptor() ([]byte, []int) {
	return file_proto_backupAgent_bakhistory_bakhistory_proto_rawDescGZIP(), []int{4}
}

func (x *HistoryListOutItem) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *HistoryListOutItem) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *HistoryListOutItem) GetDBName() string {
	if x != nil {
		return x.DBName
	}
	return ""
}

func (x *HistoryListOutItem) GetDingStatus() int64 {
	if x != nil {
		return x.DingStatus
	}
	return 0
}

func (x *HistoryListOutItem) GetOSSStatus() int64 {
	if x != nil {
		return x.OSSStatus
	}
	return 0
}

func (x *HistoryListOutItem) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HistoryListOutItem) GetFileSize() string {
	if x != nil {
		return x.FileSize
	}
	return ""
}

func (x *HistoryListOutItem) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *HistoryListOutItem) GetBakTime() string {
	if x != nil {
		return x.BakTime
	}
	return ""
}

var File_proto_backupAgent_bakhistory_bakhistory_proto protoreflect.FileDescriptor

var file_proto_backupAgent_bakhistory_bakhistory_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x2f, 0x62, 0x61, 0x6b, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x62,
	0x61, 0x6b, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1c, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x6e, 0x0a,
	0x10, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6f, 0x72,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x22, 0x20, 0x0a,
	0x0e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22,
	0x3d, 0x0a, 0x11, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4f, 0x6e, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x4f, 0x4b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x4f, 0x4b, 0x22, 0x8b,
	0x01, 0x0a, 0x11, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x60, 0x0a, 0x12, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x75, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x12, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x22, 0xfa, 0x01, 0x0a,
	0x12, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x75, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x42, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x42, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x44, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x44, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x4f, 0x53, 0x53, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x4f, 0x53, 0x53, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x42, 0x61, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x42, 0x61, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xf0, 0x01, 0x0a, 0x07, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x73, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2c, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x49, 0x44, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x2f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x4f, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x2f, 0x3b, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_backupAgent_bakhistory_bakhistory_proto_rawDescOnce sync.Once
	file_proto_backupAgent_bakhistory_bakhistory_proto_rawDescData = file_proto_backupAgent_bakhistory_bakhistory_proto_rawDesc
)

func file_proto_backupAgent_bakhistory_bakhistory_proto_rawDescGZIP() []byte {
	file_proto_backupAgent_bakhistory_bakhistory_proto_rawDescOnce.Do(func() {
		file_proto_backupAgent_bakhistory_bakhistory_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_backupAgent_bakhistory_bakhistory_proto_rawDescData)
	})
	return file_proto_backupAgent_bakhistory_bakhistory_proto_rawDescData
}

var file_proto_backupAgent_bakhistory_bakhistory_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_backupAgent_bakhistory_bakhistory_proto_goTypes = []interface{}{
	(*HistoryListInput)(nil),   // 0: go.micro.service.backupAgent.HistoryListInput
	(*HistoryIDInput)(nil),     // 1: go.micro.service.backupAgent.HistoryIDInput
	(*HistoryOneMessage)(nil),  // 2: go.micro.service.backupAgent.HistoryOneMessage
	(*HistoryListOutput)(nil),  // 3: go.micro.service.backupAgent.HistoryListOutput
	(*HistoryListOutItem)(nil), // 4: go.micro.service.backupAgent.HistoryListOutItem
}
var file_proto_backupAgent_bakhistory_bakhistory_proto_depIdxs = []int32{
	4, // 0: go.micro.service.backupAgent.HistoryListOutput.historyListOutItem:type_name -> go.micro.service.backupAgent.HistoryListOutItem
	0, // 1: go.micro.service.backupAgent.History.GetHistoryList:input_type -> go.micro.service.backupAgent.HistoryListInput
	1, // 2: go.micro.service.backupAgent.History.DeleteHistory:input_type -> go.micro.service.backupAgent.HistoryIDInput
	3, // 3: go.micro.service.backupAgent.History.GetHistoryList:output_type -> go.micro.service.backupAgent.HistoryListOutput
	2, // 4: go.micro.service.backupAgent.History.DeleteHistory:output_type -> go.micro.service.backupAgent.HistoryOneMessage
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_backupAgent_bakhistory_bakhistory_proto_init() }
func file_proto_backupAgent_bakhistory_bakhistory_proto_init() {
	if File_proto_backupAgent_bakhistory_bakhistory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_backupAgent_bakhistory_bakhistory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryListInput); i {
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
		file_proto_backupAgent_bakhistory_bakhistory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryIDInput); i {
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
		file_proto_backupAgent_bakhistory_bakhistory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryOneMessage); i {
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
		file_proto_backupAgent_bakhistory_bakhistory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryListOutput); i {
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
		file_proto_backupAgent_bakhistory_bakhistory_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryListOutItem); i {
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
			RawDescriptor: file_proto_backupAgent_bakhistory_bakhistory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_backupAgent_bakhistory_bakhistory_proto_goTypes,
		DependencyIndexes: file_proto_backupAgent_bakhistory_bakhistory_proto_depIdxs,
		MessageInfos:      file_proto_backupAgent_bakhistory_bakhistory_proto_msgTypes,
	}.Build()
	File_proto_backupAgent_bakhistory_bakhistory_proto = out.File
	file_proto_backupAgent_bakhistory_bakhistory_proto_rawDesc = nil
	file_proto_backupAgent_bakhistory_bakhistory_proto_goTypes = nil
	file_proto_backupAgent_bakhistory_bakhistory_proto_depIdxs = nil
}
