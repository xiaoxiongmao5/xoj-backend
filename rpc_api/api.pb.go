// protoc --go_out=. --go-triple_out=. ./api.proto
// EDIT IT, change to your package, service and message

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: api.proto

package rpc_api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QuestionGetByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionId int64 `protobuf:"varint,1,opt,name=questionId,proto3" json:"questionId,omitempty"`
}

func (x *QuestionGetByIdReq) Reset() {
	*x = QuestionGetByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionGetByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionGetByIdReq) ProtoMessage() {}

func (x *QuestionGetByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionGetByIdReq.ProtoReflect.Descriptor instead.
func (*QuestionGetByIdReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *QuestionGetByIdReq) GetQuestionId() int64 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

type RpcQuestionObj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content     string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Tags        string                 `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	Answer      string                 `protobuf:"bytes,5,opt,name=answer,proto3" json:"answer,omitempty"`
	SubmitNum   int32                  `protobuf:"varint,6,opt,name=submitNum,proto3" json:"submitNum,omitempty"`
	AcceptedNum int32                  `protobuf:"varint,7,opt,name=acceptedNum,proto3" json:"acceptedNum,omitempty"`
	JudgeCase   string                 `protobuf:"bytes,8,opt,name=judgeCase,proto3" json:"judgeCase,omitempty"`
	JudgeConfig string                 `protobuf:"bytes,9,opt,name=judgeConfig,proto3" json:"judgeConfig,omitempty"`
	ThumbNum    int32                  `protobuf:"varint,10,opt,name=thumbNum,proto3" json:"thumbNum,omitempty"`
	FavourNum   int32                  `protobuf:"varint,11,opt,name=favourNum,proto3" json:"favourNum,omitempty"`
	UserId      int64                  `protobuf:"varint,12,opt,name=userId,proto3" json:"userId,omitempty"`
	CreateTime  *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime  *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	IsDelete    int32                  `protobuf:"varint,15,opt,name=isDelete,proto3" json:"isDelete,omitempty"`
}

func (x *RpcQuestionObj) Reset() {
	*x = RpcQuestionObj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcQuestionObj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcQuestionObj) ProtoMessage() {}

func (x *RpcQuestionObj) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcQuestionObj.ProtoReflect.Descriptor instead.
func (*RpcQuestionObj) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *RpcQuestionObj) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RpcQuestionObj) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RpcQuestionObj) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *RpcQuestionObj) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *RpcQuestionObj) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *RpcQuestionObj) GetSubmitNum() int32 {
	if x != nil {
		return x.SubmitNum
	}
	return 0
}

func (x *RpcQuestionObj) GetAcceptedNum() int32 {
	if x != nil {
		return x.AcceptedNum
	}
	return 0
}

func (x *RpcQuestionObj) GetJudgeCase() string {
	if x != nil {
		return x.JudgeCase
	}
	return ""
}

func (x *RpcQuestionObj) GetJudgeConfig() string {
	if x != nil {
		return x.JudgeConfig
	}
	return ""
}

func (x *RpcQuestionObj) GetThumbNum() int32 {
	if x != nil {
		return x.ThumbNum
	}
	return 0
}

func (x *RpcQuestionObj) GetFavourNum() int32 {
	if x != nil {
		return x.FavourNum
	}
	return 0
}

func (x *RpcQuestionObj) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RpcQuestionObj) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *RpcQuestionObj) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *RpcQuestionObj) GetIsDelete() int32 {
	if x != nil {
		return x.IsDelete
	}
	return 0
}

type QuestionSubmitGetByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionSubmitId int64 `protobuf:"varint,1,opt,name=questionSubmitId,proto3" json:"questionSubmitId,omitempty"`
}

func (x *QuestionSubmitGetByIdReq) Reset() {
	*x = QuestionSubmitGetByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionSubmitGetByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionSubmitGetByIdReq) ProtoMessage() {}

func (x *QuestionSubmitGetByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionSubmitGetByIdReq.ProtoReflect.Descriptor instead.
func (*QuestionSubmitGetByIdReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *QuestionSubmitGetByIdReq) GetQuestionSubmitId() int64 {
	if x != nil {
		return x.QuestionSubmitId
	}
	return 0
}

type RpcQuestionSubmitObj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Language   string                 `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	Code       string                 `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	JudgeInfo  string                 `protobuf:"bytes,4,opt,name=judgeInfo,proto3" json:"judgeInfo,omitempty"`
	Status     int32                  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	QuestionId int64                  `protobuf:"varint,6,opt,name=questionId,proto3" json:"questionId,omitempty"`
	UserId     int64                  `protobuf:"varint,7,opt,name=userId,proto3" json:"userId,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	IsDelete   int32                  `protobuf:"varint,10,opt,name=isDelete,proto3" json:"isDelete,omitempty"`
}

func (x *RpcQuestionSubmitObj) Reset() {
	*x = RpcQuestionSubmitObj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcQuestionSubmitObj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcQuestionSubmitObj) ProtoMessage() {}

func (x *RpcQuestionSubmitObj) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcQuestionSubmitObj.ProtoReflect.Descriptor instead.
func (*RpcQuestionSubmitObj) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *RpcQuestionSubmitObj) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RpcQuestionSubmitObj) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *RpcQuestionSubmitObj) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *RpcQuestionSubmitObj) GetJudgeInfo() string {
	if x != nil {
		return x.JudgeInfo
	}
	return ""
}

func (x *RpcQuestionSubmitObj) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RpcQuestionSubmitObj) GetQuestionId() int64 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

func (x *RpcQuestionSubmitObj) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RpcQuestionSubmitObj) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *RpcQuestionSubmitObj) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *RpcQuestionSubmitObj) GetIsDelete() int32 {
	if x != nil {
		return x.IsDelete
	}
	return 0
}

type CommonUpdateByIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CommonUpdateByIdResp) Reset() {
	*x = CommonUpdateByIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonUpdateByIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonUpdateByIdResp) ProtoMessage() {}

func (x *CommonUpdateByIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonUpdateByIdResp.ProtoReflect.Descriptor instead.
func (*CommonUpdateByIdResp) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *CommonUpdateByIdResp) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x70, 0x63,
	0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xe2, 0x03, 0x0a, 0x0e,
	0x52, 0x70, 0x63, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x43, 0x61, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a,
	0x75, 0x64, 0x67, 0x65, 0x43, 0x61, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6a,
	0x75, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x4e, 0x75, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x61, 0x76, 0x6f, 0x75, 0x72,
	0x4e, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66, 0x61, 0x76, 0x6f, 0x75,
	0x72, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x22, 0x46, 0x0a, 0x18, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a, 0x10,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x22, 0xd8, 0x02, 0x0a, 0x14, 0x52, 0x70, 0x63,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4f, 0x62,
	0x6a, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x22, 0x2e, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x32, 0x9a, 0x01, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x41, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x72, 0x70,
	0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x70, 0x63, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62,
	0x6a, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x31, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x17, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x70, 0x63, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x1a,
	0x1d, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x32, 0xad, 0x01, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x12, 0x4d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x21,
	0x2e, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x1a, 0x1d, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x70, 0x63, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4f, 0x62, 0x6a,
	0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x1d, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x70, 0x63, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4f, 0x62, 0x6a, 0x1a,
	0x1d, 0x2e, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_goTypes = []interface{}{
	(*QuestionGetByIdReq)(nil),       // 0: rpc_api.QuestionGetByIdReq
	(*RpcQuestionObj)(nil),           // 1: rpc_api.RpcQuestionObj
	(*QuestionSubmitGetByIdReq)(nil), // 2: rpc_api.QuestionSubmitGetByIdReq
	(*RpcQuestionSubmitObj)(nil),     // 3: rpc_api.RpcQuestionSubmitObj
	(*CommonUpdateByIdResp)(nil),     // 4: rpc_api.CommonUpdateByIdResp
	(*timestamppb.Timestamp)(nil),    // 5: google.protobuf.Timestamp
}
var file_api_proto_depIdxs = []int32{
	5, // 0: rpc_api.RpcQuestionObj.createTime:type_name -> google.protobuf.Timestamp
	5, // 1: rpc_api.RpcQuestionObj.updateTime:type_name -> google.protobuf.Timestamp
	5, // 2: rpc_api.RpcQuestionSubmitObj.createTime:type_name -> google.protobuf.Timestamp
	5, // 3: rpc_api.RpcQuestionSubmitObj.updateTime:type_name -> google.protobuf.Timestamp
	0, // 4: rpc_api.Question.GetById:input_type -> rpc_api.QuestionGetByIdReq
	1, // 5: rpc_api.Question.Add1AcceptedNum:input_type -> rpc_api.RpcQuestionObj
	2, // 6: rpc_api.QuestionSubmit.GetById:input_type -> rpc_api.QuestionSubmitGetByIdReq
	3, // 7: rpc_api.QuestionSubmit.UpdateById:input_type -> rpc_api.RpcQuestionSubmitObj
	1, // 8: rpc_api.Question.GetById:output_type -> rpc_api.RpcQuestionObj
	4, // 9: rpc_api.Question.Add1AcceptedNum:output_type -> rpc_api.CommonUpdateByIdResp
	3, // 10: rpc_api.QuestionSubmit.GetById:output_type -> rpc_api.RpcQuestionSubmitObj
	4, // 11: rpc_api.QuestionSubmit.UpdateById:output_type -> rpc_api.CommonUpdateByIdResp
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionGetByIdReq); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcQuestionObj); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionSubmitGetByIdReq); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcQuestionSubmitObj); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonUpdateByIdResp); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
