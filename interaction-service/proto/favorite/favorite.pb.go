// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: interaction-service/proto/favorite.proto

package favorite

import (
	"context"
	pb "demotest/douyin-api/proto/video"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"reflect"
	"sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FavoriteActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`             // 用户id
	VideoId    int64 `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`          // 视频id
	ActionType int32 `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"` // 1-点赞，2-取消点赞
}

func (x *FavoriteActionRequest) Reset() {
	*x = FavoriteActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interaction_service_proto_favorite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteActionRequest) ProtoMessage() {}

func (x *FavoriteActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interaction_service_proto_favorite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteActionRequest.ProtoReflect.Descriptor instead.
func (*FavoriteActionRequest) Descriptor() ([]byte, []int) {
	return file_interaction_service_proto_favorite_proto_rawDescGZIP(), []int{0}
}

func (x *FavoriteActionRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FavoriteActionRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *FavoriteActionRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type FavoriteActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，200-成功，其他值-失败
	StatusMsg  *string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
}

func (x *FavoriteActionResponse) Reset() {
	*x = FavoriteActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interaction_service_proto_favorite_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteActionResponse) ProtoMessage() {}

func (x *FavoriteActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interaction_service_proto_favorite_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteActionResponse.ProtoReflect.Descriptor instead.
func (*FavoriteActionResponse) Descriptor() ([]byte, []int) {
	return file_interaction_service_proto_favorite_proto_rawDescGZIP(), []int{1}
}

func (x *FavoriteActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FavoriteActionResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

type IsFavoriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`    // 用户id
	VideoId int64 `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"` // 视频id
}

func (x *IsFavoriteRequest) Reset() {
	*x = IsFavoriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interaction_service_proto_favorite_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsFavoriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsFavoriteRequest) ProtoMessage() {}

func (x *IsFavoriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interaction_service_proto_favorite_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsFavoriteRequest.ProtoReflect.Descriptor instead.
func (*IsFavoriteRequest) Descriptor() ([]byte, []int) {
	return file_interaction_service_proto_favorite_proto_rawDescGZIP(), []int{2}
}

func (x *IsFavoriteRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *IsFavoriteRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type IsFavoriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，200-成功，其他值-失败
	IsFavorite bool  `protobuf:"varint,2,opt,name=isFavorite,proto3" json:"isFavorite,omitempty"`                   // 是否喜欢结果
}

func (x *IsFavoriteResponse) Reset() {
	*x = IsFavoriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interaction_service_proto_favorite_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsFavoriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsFavoriteResponse) ProtoMessage() {}

func (x *IsFavoriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interaction_service_proto_favorite_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsFavoriteResponse.ProtoReflect.Descriptor instead.
func (*IsFavoriteResponse) Descriptor() ([]byte, []int) {
	return file_interaction_service_proto_favorite_proto_rawDescGZIP(), []int{3}
}

func (x *IsFavoriteResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *IsFavoriteResponse) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

type FavoriteListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *FavoriteListRequest) Reset() {
	*x = FavoriteListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interaction_service_proto_favorite_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteListRequest) ProtoMessage() {}

func (x *FavoriteListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interaction_service_proto_favorite_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteListRequest.ProtoReflect.Descriptor instead.
func (*FavoriteListRequest) Descriptor() ([]byte, []int) {
	return file_interaction_service_proto_favorite_proto_rawDescGZIP(), []int{4}
}

func (x *FavoriteListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FavoriteListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type FavoriteListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoList  []*pb.Video `protobuf:"bytes,1,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
	StatusCode int32       `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  *string     `protobuf:"bytes,3,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"`
}

func (x *FavoriteListResponse) Reset() {
	*x = FavoriteListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interaction_service_proto_favorite_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteListResponse) ProtoMessage() {}

func (x *FavoriteListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interaction_service_proto_favorite_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteListResponse.ProtoReflect.Descriptor instead.
func (*FavoriteListResponse) Descriptor() ([]byte, []int) {
	return file_interaction_service_proto_favorite_proto_rawDescGZIP(), []int{5}
}

func (x *FavoriteListResponse) GetVideoList() []*pb.Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

func (x *FavoriteListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FavoriteListResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

var File_interaction_service_proto_favorite_proto protoreflect.FileDescriptor

var file_interaction_service_proto_favorite_proto_rawDesc = []byte{
	0x0a, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x64, 0x6f, 0x75, 0x79,
	0x69, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e,
	0x0a, 0x17, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x6e,
	0x0a, 0x18, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x22, 0x49,
	0x0a, 0x13, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x14, 0x69, 0x73, 0x5f,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x22, 0x46, 0x0a, 0x15, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x96, 0x01, 0x0a, 0x16, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x76, 0x70, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x22, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73,
	0x67, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x6d, 0x73, 0x67, 0x32, 0xda, 0x01, 0x0a, 0x0f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x66, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3b, 0x0a, 0x0a, 0x49, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14,
	0x2e, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x0c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x26, 0x5a, 0x24, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interaction_service_proto_favorite_proto_rawDescOnce sync.Once
	file_interaction_service_proto_favorite_proto_rawDescData = file_interaction_service_proto_favorite_proto_rawDesc
)

func file_interaction_service_proto_favorite_proto_rawDescGZIP() []byte {
	file_interaction_service_proto_favorite_proto_rawDescOnce.Do(func() {
		file_interaction_service_proto_favorite_proto_rawDescData = protoimpl.X.CompressGZIP(file_interaction_service_proto_favorite_proto_rawDescData)
	})
	return file_interaction_service_proto_favorite_proto_rawDescData
}

var file_interaction_service_proto_favorite_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_interaction_service_proto_favorite_proto_goTypes = []interface{}{
	(*FavoriteActionRequest)(nil),  // 0: favorite_action_request
	(*FavoriteActionResponse)(nil), // 1: favorite_action_response
	(*IsFavoriteRequest)(nil),      // 2: is_favorite_request
	(*IsFavoriteResponse)(nil),     // 3: is_favorite_response
	(*FavoriteListRequest)(nil),    // 4: favorite_list_request
	(*FavoriteListResponse)(nil),   // 5: favorite_list_response
	(*pb.Video)(nil),               // 6: vp.Video
}
var file_interaction_service_proto_favorite_proto_depIdxs = []int32{
	6, // 0: favorite_list_response.video_list:type_name -> vp.Video
	0, // 1: FavoriteService.FavoriteAction:input_type -> favorite_action_request
	2, // 2: FavoriteService.IsFavorite:input_type -> is_favorite_request
	4, // 3: FavoriteService.FavoriteList:input_type -> favorite_list_request
	1, // 4: FavoriteService.FavoriteAction:output_type -> favorite_action_response
	3, // 5: FavoriteService.IsFavorite:output_type -> is_favorite_response
	5, // 6: FavoriteService.FavoriteList:output_type -> favorite_list_response
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_interaction_service_proto_favorite_proto_init() }
func file_interaction_service_proto_favorite_proto_init() {
	if File_interaction_service_proto_favorite_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interaction_service_proto_favorite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteActionRequest); i {
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
		file_interaction_service_proto_favorite_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteActionResponse); i {
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
		file_interaction_service_proto_favorite_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsFavoriteRequest); i {
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
		file_interaction_service_proto_favorite_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsFavoriteResponse); i {
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
		file_interaction_service_proto_favorite_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteListRequest); i {
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
		file_interaction_service_proto_favorite_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteListResponse); i {
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
	file_interaction_service_proto_favorite_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_interaction_service_proto_favorite_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_interaction_service_proto_favorite_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interaction_service_proto_favorite_proto_goTypes,
		DependencyIndexes: file_interaction_service_proto_favorite_proto_depIdxs,
		MessageInfos:      file_interaction_service_proto_favorite_proto_msgTypes,
	}.Build()
	File_interaction_service_proto_favorite_proto = out.File
	file_interaction_service_proto_favorite_proto_rawDesc = nil
	file_interaction_service_proto_favorite_proto_goTypes = nil
	file_interaction_service_proto_favorite_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FavoriteServiceClient is the client API for FavoriteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FavoriteServiceClient interface {
	FavoriteAction(ctx context.Context, in *FavoriteActionRequest, opts ...grpc.CallOption) (*FavoriteActionResponse, error)
	IsFavorite(ctx context.Context, in *IsFavoriteRequest, opts ...grpc.CallOption) (*IsFavoriteResponse, error)
	FavoriteList(ctx context.Context, in *FavoriteListRequest, opts ...grpc.CallOption) (*FavoriteListResponse, error)
}

type favoriteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFavoriteServiceClient(cc grpc.ClientConnInterface) FavoriteServiceClient {
	return &favoriteServiceClient{cc}
}

func (c *favoriteServiceClient) FavoriteAction(ctx context.Context, in *FavoriteActionRequest, opts ...grpc.CallOption) (*FavoriteActionResponse, error) {
	out := new(FavoriteActionResponse)
	err := c.cc.Invoke(ctx, "/FavoriteService/FavoriteAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteServiceClient) IsFavorite(ctx context.Context, in *IsFavoriteRequest, opts ...grpc.CallOption) (*IsFavoriteResponse, error) {
	out := new(IsFavoriteResponse)
	err := c.cc.Invoke(ctx, "/FavoriteService/IsFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteServiceClient) FavoriteList(ctx context.Context, in *FavoriteListRequest, opts ...grpc.CallOption) (*FavoriteListResponse, error) {
	out := new(FavoriteListResponse)
	err := c.cc.Invoke(ctx, "/FavoriteService/FavoriteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavoriteServiceServer is the server API for FavoriteService service.
type FavoriteServiceServer interface {
	FavoriteAction(context.Context, *FavoriteActionRequest) (*FavoriteActionResponse, error)
	IsFavorite(context.Context, *IsFavoriteRequest) (*IsFavoriteResponse, error)
	FavoriteList(context.Context, *FavoriteListRequest) (*FavoriteListResponse, error)
}

// UnimplementedFavoriteServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFavoriteServiceServer struct {
}

func (*UnimplementedFavoriteServiceServer) FavoriteAction(context.Context, *FavoriteActionRequest) (*FavoriteActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteAction not implemented")
}
func (*UnimplementedFavoriteServiceServer) IsFavorite(context.Context, *IsFavoriteRequest) (*IsFavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFavorite not implemented")
}
func (*UnimplementedFavoriteServiceServer) FavoriteList(context.Context, *FavoriteListRequest) (*FavoriteListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteList not implemented")
}

func RegisterFavoriteServiceServer(s *grpc.Server, srv FavoriteServiceServer) {
	s.RegisterService(&_FavoriteService_serviceDesc, srv)
}

func _FavoriteService_FavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServiceServer).FavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FavoriteService/FavoriteAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServiceServer).FavoriteAction(ctx, req.(*FavoriteActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteService_IsFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServiceServer).IsFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FavoriteService/IsFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServiceServer).IsFavorite(ctx, req.(*IsFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteService_FavoriteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServiceServer).FavoriteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FavoriteService/FavoriteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServiceServer).FavoriteList(ctx, req.(*FavoriteListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FavoriteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FavoriteService",
	HandlerType: (*FavoriteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FavoriteAction",
			Handler:    _FavoriteService_FavoriteAction_Handler,
		},
		{
			MethodName: "IsFavorite",
			Handler:    _FavoriteService_IsFavorite_Handler,
		},
		{
			MethodName: "FavoriteList",
			Handler:    _FavoriteService_FavoriteList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interaction-service/proto/favorite.proto",
}
