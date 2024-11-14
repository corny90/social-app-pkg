// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: grpc-proto/behaviour.proto

package grpc_proto

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

// PING
type PingApiBehaviourRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingApiBehaviourRequest) Reset() {
	*x = PingApiBehaviourRequest{}
	mi := &file_grpc_proto_behaviour_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingApiBehaviourRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingApiBehaviourRequest) ProtoMessage() {}

func (x *PingApiBehaviourRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_behaviour_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingApiBehaviourRequest.ProtoReflect.Descriptor instead.
func (*PingApiBehaviourRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_behaviour_proto_rawDescGZIP(), []int{0}
}

type PingApiBehaviourResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PingApiBehaviourResponse) Reset() {
	*x = PingApiBehaviourResponse{}
	mi := &file_grpc_proto_behaviour_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingApiBehaviourResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingApiBehaviourResponse) ProtoMessage() {}

func (x *PingApiBehaviourResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_behaviour_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingApiBehaviourResponse.ProtoReflect.Descriptor instead.
func (*PingApiBehaviourResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_behaviour_proto_rawDescGZIP(), []int{1}
}

func (x *PingApiBehaviourResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// POST CREATE
type BehaviourFetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BehaviourFetchRequest) Reset() {
	*x = BehaviourFetchRequest{}
	mi := &file_grpc_proto_behaviour_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BehaviourFetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BehaviourFetchRequest) ProtoMessage() {}

func (x *BehaviourFetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_behaviour_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BehaviourFetchRequest.ProtoReflect.Descriptor instead.
func (*BehaviourFetchRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_behaviour_proto_rawDescGZIP(), []int{2}
}

func (x *BehaviourFetchRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type BehaviourFetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string         `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	StatisticData *StatisticData `protobuf:"bytes,2,opt,name=statistic_data,json=statisticData,proto3" json:"statistic_data,omitempty"`
}

func (x *BehaviourFetchResponse) Reset() {
	*x = BehaviourFetchResponse{}
	mi := &file_grpc_proto_behaviour_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BehaviourFetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BehaviourFetchResponse) ProtoMessage() {}

func (x *BehaviourFetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_behaviour_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BehaviourFetchResponse.ProtoReflect.Descriptor instead.
func (*BehaviourFetchResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_behaviour_proto_rawDescGZIP(), []int{3}
}

func (x *BehaviourFetchResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BehaviourFetchResponse) GetStatisticData() *StatisticData {
	if x != nil {
		return x.StatisticData
	}
	return nil
}

type TrackRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PostsId       string   `protobuf:"bytes,2,opt,name=posts_id,json=postsId,proto3" json:"posts_id,omitempty"`
	ActionView    bool     `protobuf:"varint,3,opt,name=action_view,json=actionView,proto3" json:"action_view,omitempty"`
	ActionLike    bool     `protobuf:"varint,4,opt,name=action_like,json=actionLike,proto3" json:"action_like,omitempty"`
	ActionComment bool     `protobuf:"varint,5,opt,name=action_comment,json=actionComment,proto3" json:"action_comment,omitempty"`
	ViewDuration  int32    `protobuf:"varint,6,opt,name=view_duration,json=viewDuration,proto3" json:"view_duration,omitempty"`
	Clicks        *Clicks  `protobuf:"bytes,7,opt,name=clicks,proto3" json:"clicks,omitempty"`
	CreatedAt     string   `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Keywords      []string `protobuf:"bytes,9,rep,name=keywords,proto3" json:"keywords,omitempty"`
	Score         int32    `protobuf:"varint,10,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *TrackRecord) Reset() {
	*x = TrackRecord{}
	mi := &file_grpc_proto_behaviour_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrackRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackRecord) ProtoMessage() {}

func (x *TrackRecord) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_behaviour_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackRecord.ProtoReflect.Descriptor instead.
func (*TrackRecord) Descriptor() ([]byte, []int) {
	return file_grpc_proto_behaviour_proto_rawDescGZIP(), []int{4}
}

func (x *TrackRecord) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TrackRecord) GetPostsId() string {
	if x != nil {
		return x.PostsId
	}
	return ""
}

func (x *TrackRecord) GetActionView() bool {
	if x != nil {
		return x.ActionView
	}
	return false
}

func (x *TrackRecord) GetActionLike() bool {
	if x != nil {
		return x.ActionLike
	}
	return false
}

func (x *TrackRecord) GetActionComment() bool {
	if x != nil {
		return x.ActionComment
	}
	return false
}

func (x *TrackRecord) GetViewDuration() int32 {
	if x != nil {
		return x.ViewDuration
	}
	return 0
}

func (x *TrackRecord) GetClicks() *Clicks {
	if x != nil {
		return x.Clicks
	}
	return nil
}

func (x *TrackRecord) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *TrackRecord) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *TrackRecord) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type Clicks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostUser bool `protobuf:"varint,1,opt,name=post_user,json=postUser,proto3" json:"post_user,omitempty"`
	Views    bool `protobuf:"varint,2,opt,name=views,proto3" json:"views,omitempty"`
	Likes    bool `protobuf:"varint,3,opt,name=likes,proto3" json:"likes,omitempty"`
	Comments bool `protobuf:"varint,4,opt,name=comments,proto3" json:"comments,omitempty"`
}

func (x *Clicks) Reset() {
	*x = Clicks{}
	mi := &file_grpc_proto_behaviour_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Clicks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Clicks) ProtoMessage() {}

func (x *Clicks) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_behaviour_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Clicks.ProtoReflect.Descriptor instead.
func (*Clicks) Descriptor() ([]byte, []int) {
	return file_grpc_proto_behaviour_proto_rawDescGZIP(), []int{5}
}

func (x *Clicks) GetPostUser() bool {
	if x != nil {
		return x.PostUser
	}
	return false
}

func (x *Clicks) GetViews() bool {
	if x != nil {
		return x.Views
	}
	return false
}

func (x *Clicks) GetLikes() bool {
	if x != nil {
		return x.Likes
	}
	return false
}

func (x *Clicks) GetComments() bool {
	if x != nil {
		return x.Comments
	}
	return false
}

type KeywordStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword    string `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	TotalViews int32  `protobuf:"varint,2,opt,name=total_views,json=totalViews,proto3" json:"total_views,omitempty"`
	TotalTime  int32  `protobuf:"varint,3,opt,name=total_time,json=totalTime,proto3" json:"total_time,omitempty"`
}

func (x *KeywordStats) Reset() {
	*x = KeywordStats{}
	mi := &file_grpc_proto_behaviour_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeywordStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeywordStats) ProtoMessage() {}

func (x *KeywordStats) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_behaviour_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeywordStats.ProtoReflect.Descriptor instead.
func (*KeywordStats) Descriptor() ([]byte, []int) {
	return file_grpc_proto_behaviour_proto_rawDescGZIP(), []int{6}
}

func (x *KeywordStats) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *KeywordStats) GetTotalViews() int32 {
	if x != nil {
		return x.TotalViews
	}
	return 0
}

func (x *KeywordStats) GetTotalTime() int32 {
	if x != nil {
		return x.TotalTime
	}
	return 0
}

type StatisticData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopRecords  []*TrackRecord  `protobuf:"bytes,1,rep,name=topRecords,proto3" json:"topRecords,omitempty"`
	TopPosts    []*TrackRecord  `protobuf:"bytes,2,rep,name=topPosts,proto3" json:"topPosts,omitempty"`
	TopKeywords []*KeywordStats `protobuf:"bytes,3,rep,name=topKeywords,proto3" json:"topKeywords,omitempty"`
}

func (x *StatisticData) Reset() {
	*x = StatisticData{}
	mi := &file_grpc_proto_behaviour_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatisticData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticData) ProtoMessage() {}

func (x *StatisticData) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_behaviour_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticData.ProtoReflect.Descriptor instead.
func (*StatisticData) Descriptor() ([]byte, []int) {
	return file_grpc_proto_behaviour_proto_rawDescGZIP(), []int{7}
}

func (x *StatisticData) GetTopRecords() []*TrackRecord {
	if x != nil {
		return x.TopRecords
	}
	return nil
}

func (x *StatisticData) GetTopPosts() []*TrackRecord {
	if x != nil {
		return x.TopPosts
	}
	return nil
}

func (x *StatisticData) GetTopKeywords() []*KeywordStats {
	if x != nil {
		return x.TopKeywords
	}
	return nil
}

var File_grpc_proto_behaviour_proto protoreflect.FileDescriptor

var file_grpc_proto_behaviour_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x22, 0x19, 0x0a, 0x17, 0x50, 0x69, 0x6e, 0x67, 0x41,
	0x70, 0x69, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x32, 0x0a, 0x18, 0x50, 0x69, 0x6e, 0x67, 0x41, 0x70, 0x69, 0x42, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x30, 0x0a, 0x15, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x75, 0x72, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x72, 0x0a, 0x16, 0x42, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x75, 0x72, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0e, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0d, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x22, 0xcb, 0x02, 0x0a,
	0x0b, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x69, 0x65,
	0x77, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x6b, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x6b, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x76, 0x69, 0x65, 0x77, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29,
	0x0a, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b,
	0x73, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x6d, 0x0a, 0x06, 0x43, 0x6c,
	0x69, 0x63, 0x6b, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x68, 0x0a, 0x0c, 0x4b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56,
	0x69, 0x65, 0x77, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x44, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x0a, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x75, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x0a, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x32, 0x0a,
	0x08, 0x74, 0x6f, 0x70, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x08, 0x74, 0x6f, 0x70, 0x50, 0x6f, 0x73, 0x74,
	0x73, 0x12, 0x39, 0x0a, 0x0b, 0x74, 0x6f, 0x70, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x75, 0x72, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x0b, 0x74, 0x6f, 0x70, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x32, 0xc6, 0x01, 0x0a,
	0x10, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5b, 0x0a, 0x10, 0x50, 0x69, 0x6e, 0x67, 0x41, 0x70, 0x69, 0x42, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x75, 0x72, 0x12, 0x22, 0x2e, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75,
	0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x41, 0x70, 0x69, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x75, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x41, 0x70, 0x69, 0x42, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55,
	0x0a, 0x0e, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x12, 0x20, 0x2e, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x2e, 0x42, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x2e, 0x42,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x18, 0x5a, 0x16, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_behaviour_proto_rawDescOnce sync.Once
	file_grpc_proto_behaviour_proto_rawDescData = file_grpc_proto_behaviour_proto_rawDesc
)

func file_grpc_proto_behaviour_proto_rawDescGZIP() []byte {
	file_grpc_proto_behaviour_proto_rawDescOnce.Do(func() {
		file_grpc_proto_behaviour_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_behaviour_proto_rawDescData)
	})
	return file_grpc_proto_behaviour_proto_rawDescData
}

var file_grpc_proto_behaviour_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_grpc_proto_behaviour_proto_goTypes = []any{
	(*PingApiBehaviourRequest)(nil),  // 0: behaviour.PingApiBehaviourRequest
	(*PingApiBehaviourResponse)(nil), // 1: behaviour.PingApiBehaviourResponse
	(*BehaviourFetchRequest)(nil),    // 2: behaviour.BehaviourFetchRequest
	(*BehaviourFetchResponse)(nil),   // 3: behaviour.BehaviourFetchResponse
	(*TrackRecord)(nil),              // 4: behaviour.TrackRecord
	(*Clicks)(nil),                   // 5: behaviour.Clicks
	(*KeywordStats)(nil),             // 6: behaviour.KeywordStats
	(*StatisticData)(nil),            // 7: behaviour.StatisticData
}
var file_grpc_proto_behaviour_proto_depIdxs = []int32{
	7, // 0: behaviour.BehaviourFetchResponse.statistic_data:type_name -> behaviour.StatisticData
	5, // 1: behaviour.TrackRecord.clicks:type_name -> behaviour.Clicks
	4, // 2: behaviour.StatisticData.topRecords:type_name -> behaviour.TrackRecord
	4, // 3: behaviour.StatisticData.topPosts:type_name -> behaviour.TrackRecord
	6, // 4: behaviour.StatisticData.topKeywords:type_name -> behaviour.KeywordStats
	0, // 5: behaviour.BehaviourService.PingApiBehaviour:input_type -> behaviour.PingApiBehaviourRequest
	2, // 6: behaviour.BehaviourService.BehaviourFetch:input_type -> behaviour.BehaviourFetchRequest
	1, // 7: behaviour.BehaviourService.PingApiBehaviour:output_type -> behaviour.PingApiBehaviourResponse
	3, // 8: behaviour.BehaviourService.BehaviourFetch:output_type -> behaviour.BehaviourFetchResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_grpc_proto_behaviour_proto_init() }
func file_grpc_proto_behaviour_proto_init() {
	if File_grpc_proto_behaviour_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_proto_behaviour_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_behaviour_proto_goTypes,
		DependencyIndexes: file_grpc_proto_behaviour_proto_depIdxs,
		MessageInfos:      file_grpc_proto_behaviour_proto_msgTypes,
	}.Build()
	File_grpc_proto_behaviour_proto = out.File
	file_grpc_proto_behaviour_proto_rawDesc = nil
	file_grpc_proto_behaviour_proto_goTypes = nil
	file_grpc_proto_behaviour_proto_depIdxs = nil
}
