// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: apps/video/pb/video.proto

package video

import (
	user "github.com/Go-To-Byte/DouSheng/user_center/apps/user"
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

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 视频唯一标识
	// @gotags: json:"id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 视频作者
	// @gotags: json:"author"
	Author *user.User `protobuf:"bytes,2,opt,name=author,proto3" json:"author"`
	// 视频播放地址
	// @gotags: json:"play_url"
	PlayUrl string `protobuf:"bytes,3,opt,name=play_url,json=playUrl,proto3" json:"play_url"`
	// 视频封面地址
	// @gotags: json:"cover_url"
	CoverUrl string `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url"`
	// 视频的点赞总数
	// @gotags: json:"favorite_count"
	FavoriteCount int64 `protobuf:"varint,5,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count"`
	// 视频的评论总数
	// @gotags: json:"comment_count"
	CommentCount int64 `protobuf:"varint,6,opt,name=comment_count,json=commentCount,proto3" json:"comment_count"`
	// true-已点赞,false-未点赞
	// @gotags: json:"is_favorite"
	IsFavorite bool `protobuf:"varint,7,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite"`
	// 视频标题
	// @gotags: json:"title"
	Title string `protobuf:"bytes,8,opt,name=title,proto3" json:"title"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_video_pb_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_apps_video_pb_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_apps_video_pb_video_proto_rawDescGZIP(), []int{0}
}

func (x *Video) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetAuthor() *user.User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Video) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Video) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type VideoPo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 视频唯一标识
	// @gotags: json:"id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 视频作者ID TODO：若从视频列表中获取数据过多，这里可以直接用User[二次遍历]
	// 视频播放地址
	// @gotags: json:"play_url"
	PlayUrl string `protobuf:"bytes,2,opt,name=play_url,json=playUrl,proto3" json:"play_url"`
	// 视频封面地址
	// @gotags: json:"cover_url"
	CoverUrl string `protobuf:"bytes,3,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url"`
	// 视频标题
	// @gotags: json:"title"
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title"`
	// @gotags: json:"author_id"
	AuthorId int64 `protobuf:"varint,5,opt,name=author_id,json=authorId,proto3" json:"author_id"`
	// @gotags: json:"created_at"
	CreatedAt int64 `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
}

func (x *VideoPo) Reset() {
	*x = VideoPo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_video_pb_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoPo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoPo) ProtoMessage() {}

func (x *VideoPo) ProtoReflect() protoreflect.Message {
	mi := &file_apps_video_pb_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoPo.ProtoReflect.Descriptor instead.
func (*VideoPo) Descriptor() ([]byte, []int) {
	return file_apps_video_pb_video_proto_rawDescGZIP(), []int{1}
}

func (x *VideoPo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VideoPo) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *VideoPo) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *VideoPo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *VideoPo) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *VideoPo) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

// 分页请求参数 TODO：可以放在公共模块
type PageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 每页数据数量
	// @gotags: json:"page_size" form:"page_size"
	PageSize uint64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size" form:"page_size"`
	// 第几页
	// @gotags: json:"page_number" form:"page_number"
	PageNumber uint64 `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number" form:"page_number"`
	// 偏移量
	// @gotags: json:"offset" form:"offset"
	Offset int64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset" form:"offset"`
}

func (x *PageRequest) Reset() {
	*x = PageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_video_pb_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRequest) ProtoMessage() {}

func (x *PageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_video_pb_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRequest.ProtoReflect.Descriptor instead.
func (*PageRequest) Descriptor() ([]byte, []int) {
	return file_apps_video_pb_video_proto_rawDescGZIP(), []int{2}
}

func (x *PageRequest) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageRequest) GetPageNumber() uint64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *PageRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type FeedVideosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	// @gotags: json:"latest_time" form:"latest_time"
	LatestTime *int64 `protobuf:"varint,2,opt,name=latest_time,json=latestTime,proto3,oneof" json:"latest_time" form:"latest_time"`
	// 可选参数，登录用户设置
	// @gotags: json:"token" form:"token"
	Token *string `protobuf:"bytes,3,opt,name=token,proto3,oneof" json:"token" form:"token"`
}

func (x *FeedVideosRequest) Reset() {
	*x = FeedVideosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_video_pb_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedVideosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedVideosRequest) ProtoMessage() {}

func (x *FeedVideosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_video_pb_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedVideosRequest.ProtoReflect.Descriptor instead.
func (*FeedVideosRequest) Descriptor() ([]byte, []int) {
	return file_apps_video_pb_video_proto_rawDescGZIP(), []int{3}
}

func (x *FeedVideosRequest) GetPage() *PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *FeedVideosRequest) GetLatestTime() int64 {
	if x != nil && x.LatestTime != nil {
		return *x.LatestTime
	}
	return 0
}

func (x *FeedVideosRequest) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

type PublishVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户鉴权Token
	// @gotags: json:"token" form:"token" binding:"required"
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token" form:"token" binding:"required"`
	// 视频标题
	// @gotags: json:"title" form:"title" binding:"required" validate:"required"
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title" form:"title" binding:"required" validate:"required"`
	// 视频播放地址
	// @gotags: json:"play_url"
	PlayUrl string `protobuf:"bytes,3,opt,name=play_url,json=playUrl,proto3" json:"play_url"`
	// 视频封面地址
	// @gotags: json:"cover_url"
	CoverUrl string `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url"`
}

func (x *PublishVideoRequest) Reset() {
	*x = PublishVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_video_pb_video_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishVideoRequest) ProtoMessage() {}

func (x *PublishVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_video_pb_video_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishVideoRequest.ProtoReflect.Descriptor instead.
func (*PublishVideoRequest) Descriptor() ([]byte, []int) {
	return file_apps_video_pb_video_proto_rawDescGZIP(), []int{4}
}

func (x *PublishVideoRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *PublishVideoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PublishVideoRequest) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *PublishVideoRequest) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

// 发布列表的 请求 model
type PublishListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户ID
	// @gotags: json:"user_id" form:"token" validate:"required" binding:"required"
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id" form:"token" validate:"required" binding:"required"`
	// 用户鉴权Token
	// @gotags: json:"token" form:"token" binding:"required"
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token" form:"token" binding:"required"`
}

func (x *PublishListRequest) Reset() {
	*x = PublishListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_video_pb_video_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishListRequest) ProtoMessage() {}

func (x *PublishListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_video_pb_video_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishListRequest.ProtoReflect.Descriptor instead.
func (*PublishListRequest) Descriptor() ([]byte, []int) {
	return file_apps_video_pb_video_proto_rawDescGZIP(), []int{5}
}

func (x *PublishListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PublishListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type FeedSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 本次返回的视频中，发布最早的时间，作为下次请求的latest_time
	// @gotags: json:"next_time"
	NextTime *int64 `protobuf:"varint,2,opt,name=next_time,json=nextTime,proto3,oneof" json:"next_time"`
	// 视频列表
	// @gotags: json:"video_list"
	VideoList []*Video `protobuf:"bytes,1,rep,name=video_list,json=videoList,proto3" json:"video_list"`
}

func (x *FeedSetResponse) Reset() {
	*x = FeedSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_video_pb_video_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedSetResponse) ProtoMessage() {}

func (x *FeedSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_video_pb_video_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedSetResponse.ProtoReflect.Descriptor instead.
func (*FeedSetResponse) Descriptor() ([]byte, []int) {
	return file_apps_video_pb_video_proto_rawDescGZIP(), []int{6}
}

func (x *FeedSetResponse) GetNextTime() int64 {
	if x != nil && x.NextTime != nil {
		return *x.NextTime
	}
	return 0
}

func (x *FeedSetResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type PublishVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 可以携带一些额外属性
	// @gotags: json:"mate"
	Mate map[string]string `protobuf:"bytes,1,rep,name=mate,proto3" json:"mate" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PublishVideoResponse) Reset() {
	*x = PublishVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_video_pb_video_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishVideoResponse) ProtoMessage() {}

func (x *PublishVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_video_pb_video_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishVideoResponse.ProtoReflect.Descriptor instead.
func (*PublishVideoResponse) Descriptor() ([]byte, []int) {
	return file_apps_video_pb_video_proto_rawDescGZIP(), []int{7}
}

func (x *PublishVideoResponse) GetMate() map[string]string {
	if x != nil {
		return x.Mate
	}
	return nil
}

type PublishListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户发布的视频列表
	// @gotags: json:"video_list"
	VideoList []*Video `protobuf:"bytes,1,rep,name=video_list,json=videoList,proto3" json:"video_list"`
}

func (x *PublishListResponse) Reset() {
	*x = PublishListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_video_pb_video_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishListResponse) ProtoMessage() {}

func (x *PublishListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_video_pb_video_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishListResponse.ProtoReflect.Descriptor instead.
func (*PublishListResponse) Descriptor() ([]byte, []int) {
	return file_apps_video_pb_video_proto_rawDescGZIP(), []int{8}
}

func (x *PublishListResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

var File_apps_video_pb_video_proto protoreflect.FileDescriptor

var file_apps_video_pb_video_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x70, 0x62, 0x2f,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x6f, 0x75,
	0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x1a, 0x42, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x2d, 0x54, 0x6f, 0x2d, 0x42, 0x79,
	0x74, 0x65, 0x2f, 0x44, 0x6f, 0x75, 0x53, 0x68, 0x65, 0x6e, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xff, 0x01, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x6f, 0x75, 0x73,
	0x68, 0x65, 0x6e, 0x67, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72,
	0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x25,
	0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73,
	0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x22, 0xa3, 0x01, 0x0a, 0x07, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x63, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x9f, 0x01, 0x0a,
	0x11, 0x46, 0x65, 0x65, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0a, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x79,
	0x0a, 0x13, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x22, 0x43, 0x0a, 0x12, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x77,
	0x0a, 0x0f, 0x46, 0x65, 0x65, 0x64, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65,
	0x6e, 0x67, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x14, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x42, 0x0a, 0x04, 0x6d, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x6d, 0x61, 0x74, 0x65, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4b, 0x0a,
	0x13, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68,
	0x65, 0x6e, 0x67, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52,
	0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x8e, 0x02, 0x0a, 0x07, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0a, 0x46, 0x65, 0x65, 0x64, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x73, 0x12, 0x21, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65,
	0x6e, 0x67, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x23, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68,
	0x65, 0x6e, 0x67, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x22, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e,
	0x67, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x2d, 0x54, 0x6f, 0x2d,
	0x42, 0x79, 0x74, 0x65, 0x2f, 0x44, 0x6f, 0x75, 0x53, 0x68, 0x65, 0x6e, 0x67, 0x2f, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73,
	0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_video_pb_video_proto_rawDescOnce sync.Once
	file_apps_video_pb_video_proto_rawDescData = file_apps_video_pb_video_proto_rawDesc
)

func file_apps_video_pb_video_proto_rawDescGZIP() []byte {
	file_apps_video_pb_video_proto_rawDescOnce.Do(func() {
		file_apps_video_pb_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_video_pb_video_proto_rawDescData)
	})
	return file_apps_video_pb_video_proto_rawDescData
}

var file_apps_video_pb_video_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_apps_video_pb_video_proto_goTypes = []interface{}{
	(*Video)(nil),                // 0: dousheng.video.Video
	(*VideoPo)(nil),              // 1: dousheng.video.VideoPo
	(*PageRequest)(nil),          // 2: dousheng.video.PageRequest
	(*FeedVideosRequest)(nil),    // 3: dousheng.video.FeedVideosRequest
	(*PublishVideoRequest)(nil),  // 4: dousheng.video.PublishVideoRequest
	(*PublishListRequest)(nil),   // 5: dousheng.video.PublishListRequest
	(*FeedSetResponse)(nil),      // 6: dousheng.video.FeedSetResponse
	(*PublishVideoResponse)(nil), // 7: dousheng.video.PublishVideoResponse
	(*PublishListResponse)(nil),  // 8: dousheng.video.PublishListResponse
	nil,                          // 9: dousheng.video.PublishVideoResponse.MateEntry
	(*user.User)(nil),            // 10: dousheng.user.User
}
var file_apps_video_pb_video_proto_depIdxs = []int32{
	10, // 0: dousheng.video.Video.author:type_name -> dousheng.user.User
	2,  // 1: dousheng.video.FeedVideosRequest.page:type_name -> dousheng.video.PageRequest
	0,  // 2: dousheng.video.FeedSetResponse.video_list:type_name -> dousheng.video.Video
	9,  // 3: dousheng.video.PublishVideoResponse.mate:type_name -> dousheng.video.PublishVideoResponse.MateEntry
	0,  // 4: dousheng.video.PublishListResponse.video_list:type_name -> dousheng.video.Video
	3,  // 5: dousheng.video.Service.FeedVideos:input_type -> dousheng.video.FeedVideosRequest
	4,  // 6: dousheng.video.Service.PublishVideo:input_type -> dousheng.video.PublishVideoRequest
	5,  // 7: dousheng.video.Service.PublishList:input_type -> dousheng.video.PublishListRequest
	6,  // 8: dousheng.video.Service.FeedVideos:output_type -> dousheng.video.FeedSetResponse
	7,  // 9: dousheng.video.Service.PublishVideo:output_type -> dousheng.video.PublishVideoResponse
	8,  // 10: dousheng.video.Service.PublishList:output_type -> dousheng.video.PublishListResponse
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_apps_video_pb_video_proto_init() }
func file_apps_video_pb_video_proto_init() {
	if File_apps_video_pb_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_video_pb_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_apps_video_pb_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoPo); i {
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
		file_apps_video_pb_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageRequest); i {
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
		file_apps_video_pb_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedVideosRequest); i {
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
		file_apps_video_pb_video_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishVideoRequest); i {
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
		file_apps_video_pb_video_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishListRequest); i {
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
		file_apps_video_pb_video_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedSetResponse); i {
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
		file_apps_video_pb_video_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishVideoResponse); i {
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
		file_apps_video_pb_video_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishListResponse); i {
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
	file_apps_video_pb_video_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_apps_video_pb_video_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_video_pb_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_video_pb_video_proto_goTypes,
		DependencyIndexes: file_apps_video_pb_video_proto_depIdxs,
		MessageInfos:      file_apps_video_pb_video_proto_msgTypes,
	}.Build()
	File_apps_video_pb_video_proto = out.File
	file_apps_video_pb_video_proto_rawDesc = nil
	file_apps_video_pb_video_proto_goTypes = nil
	file_apps_video_pb_video_proto_depIdxs = nil
}
