// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: sync/v1/sync.proto

package syncv1

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

// Authentication data that will be sent in a `harmonytypes.v1.Token`.
type AuthData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The server ID of the server initiating the transaction. For Pull,
	// this tells the server being connected to which homeservers' events it should send.
	// For Push, this tells the server being connected to which homeservers' events it is
	// receiving.
	ServerId string `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// The UTC UNIX time in seconds of when the request is started. Servers should reject
	// tokens with a time too far from the current time, at their discretion. A recommended
	// variance is 1 minute.
	Time uint64 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *AuthData) Reset() {
	*x = AuthData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthData) ProtoMessage() {}

func (x *AuthData) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthData.ProtoReflect.Descriptor instead.
func (*AuthData) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_proto_rawDescGZIP(), []int{0}
}

func (x *AuthData) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *AuthData) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

// Object representing a postbox event.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The kind and data of this event.
	//
	// Types that are assignable to Kind:
	//	*Event_UserRemovedFromGuild_
	//	*Event_UserAddedToGuild_
	//	*Event_UserInvited_
	//	*Event_UserRejectedInvite_
	Kind isEvent_Kind `protobuf_oneof:"kind"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_proto_rawDescGZIP(), []int{1}
}

func (m *Event) GetKind() isEvent_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Event) GetUserRemovedFromGuild() *Event_UserRemovedFromGuild {
	if x, ok := x.GetKind().(*Event_UserRemovedFromGuild_); ok {
		return x.UserRemovedFromGuild
	}
	return nil
}

func (x *Event) GetUserAddedToGuild() *Event_UserAddedToGuild {
	if x, ok := x.GetKind().(*Event_UserAddedToGuild_); ok {
		return x.UserAddedToGuild
	}
	return nil
}

func (x *Event) GetUserInvited() *Event_UserInvited {
	if x, ok := x.GetKind().(*Event_UserInvited_); ok {
		return x.UserInvited
	}
	return nil
}

func (x *Event) GetUserRejectedInvite() *Event_UserRejectedInvite {
	if x, ok := x.GetKind().(*Event_UserRejectedInvite_); ok {
		return x.UserRejectedInvite
	}
	return nil
}

type isEvent_Kind interface {
	isEvent_Kind()
}

type Event_UserRemovedFromGuild_ struct {
	// User removed from a guild.
	UserRemovedFromGuild *Event_UserRemovedFromGuild `protobuf:"bytes,1,opt,name=user_removed_from_guild,json=userRemovedFromGuild,proto3,oneof"`
}

type Event_UserAddedToGuild_ struct {
	// User added to a guild.
	UserAddedToGuild *Event_UserAddedToGuild `protobuf:"bytes,2,opt,name=user_added_to_guild,json=userAddedToGuild,proto3,oneof"`
}

type Event_UserInvited_ struct {
	// User invited to a guild.
	UserInvited *Event_UserInvited `protobuf:"bytes,3,opt,name=user_invited,json=userInvited,proto3,oneof"`
}

type Event_UserRejectedInvite_ struct {
	// User rejected a guild invitation.
	UserRejectedInvite *Event_UserRejectedInvite `protobuf:"bytes,4,opt,name=user_rejected_invite,json=userRejectedInvite,proto3,oneof"`
}

func (*Event_UserRemovedFromGuild_) isEvent_Kind() {}

func (*Event_UserAddedToGuild_) isEvent_Kind() {}

func (*Event_UserInvited_) isEvent_Kind() {}

func (*Event_UserRejectedInvite_) isEvent_Kind() {}

// Used in `Pull` endpoint.
type PullRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PullRequest) Reset() {
	*x = PullRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullRequest) ProtoMessage() {}

func (x *PullRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullRequest.ProtoReflect.Descriptor instead.
func (*PullRequest) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_proto_rawDescGZIP(), []int{2}
}

// Used in `Pull` endpoint.
type PullResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The events that were not processed yet.
	EventQueue []*Event `protobuf:"bytes,1,rep,name=event_queue,json=eventQueue,proto3" json:"event_queue,omitempty"`
}

func (x *PullResponse) Reset() {
	*x = PullResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullResponse) ProtoMessage() {}

func (x *PullResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullResponse.ProtoReflect.Descriptor instead.
func (*PullResponse) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_proto_rawDescGZIP(), []int{3}
}

func (x *PullResponse) GetEventQueue() []*Event {
	if x != nil {
		return x.EventQueue
	}
	return nil
}

// Used in `Push` endpoint.
type PushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The event to push to the server.
	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_proto_rawDescGZIP(), []int{4}
}

func (x *PushRequest) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

// Used in `Push` endpoint.
type PushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushResponse) Reset() {
	*x = PushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushResponse) ProtoMessage() {}

func (x *PushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushResponse.ProtoReflect.Descriptor instead.
func (*PushResponse) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_proto_rawDescGZIP(), []int{5}
}

// Used in `NotifyNewId` endpoint.
type NotifyNewIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The new server ID of the server.
	NewServerId string `protobuf:"bytes,1,opt,name=new_server_id,json=newServerId,proto3" json:"new_server_id,omitempty"`
}

func (x *NotifyNewIdRequest) Reset() {
	*x = NotifyNewIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyNewIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyNewIdRequest) ProtoMessage() {}

func (x *NotifyNewIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyNewIdRequest.ProtoReflect.Descriptor instead.
func (*NotifyNewIdRequest) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_proto_rawDescGZIP(), []int{6}
}

func (x *NotifyNewIdRequest) GetNewServerId() string {
	if x != nil {
		return x.NewServerId
	}
	return ""
}

// Used in `NotifyNewId` endpoint.
type NotifyNewIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyNewIdResponse) Reset() {
	*x = NotifyNewIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyNewIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyNewIdResponse) ProtoMessage() {}

func (x *NotifyNewIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyNewIdResponse.ProtoReflect.Descriptor instead.
func (*NotifyNewIdResponse) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_proto_rawDescGZIP(), []int{7}
}

// Event sent when a user is removed from a guild.
type Event_UserRemovedFromGuild struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User ID of the user that was removed.
	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Guild ID of the guild where the user was.
	GuildId uint64 `protobuf:"varint,2,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
}

func (x *Event_UserRemovedFromGuild) Reset() {
	*x = Event_UserRemovedFromGuild{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event_UserRemovedFromGuild) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event_UserRemovedFromGuild) ProtoMessage() {}

func (x *Event_UserRemovedFromGuild) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event_UserRemovedFromGuild.ProtoReflect.Descriptor instead.
func (*Event_UserRemovedFromGuild) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Event_UserRemovedFromGuild) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Event_UserRemovedFromGuild) GetGuildId() uint64 {
	if x != nil {
		return x.GuildId
	}
	return 0
}

// Event sent when a user is added to a guild.
type Event_UserAddedToGuild struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User ID of the user that was added.
	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Guild ID of the guild where the user will be.
	GuildId uint64 `protobuf:"varint,2,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
}

func (x *Event_UserAddedToGuild) Reset() {
	*x = Event_UserAddedToGuild{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event_UserAddedToGuild) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event_UserAddedToGuild) ProtoMessage() {}

func (x *Event_UserAddedToGuild) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event_UserAddedToGuild.ProtoReflect.Descriptor instead.
func (*Event_UserAddedToGuild) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Event_UserAddedToGuild) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Event_UserAddedToGuild) GetGuildId() uint64 {
	if x != nil {
		return x.GuildId
	}
	return 0
}

// Event sent when a user is invited to a guild.
type Event_UserInvited struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User ID of the invitee.
	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// User ID of the user that invited.
	InviterId uint64 `protobuf:"varint,2,opt,name=inviter_id,json=inviterId,proto3" json:"inviter_id,omitempty"`
	// The unique identifier of a user's invite to another
	// user to join a given guild.
	InviteId string `protobuf:"bytes,3,opt,name=invite_id,json=inviteId,proto3" json:"invite_id,omitempty"`
}

func (x *Event_UserInvited) Reset() {
	*x = Event_UserInvited{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event_UserInvited) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event_UserInvited) ProtoMessage() {}

func (x *Event_UserInvited) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event_UserInvited.ProtoReflect.Descriptor instead.
func (*Event_UserInvited) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_proto_rawDescGZIP(), []int{1, 2}
}

func (x *Event_UserInvited) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Event_UserInvited) GetInviterId() uint64 {
	if x != nil {
		return x.InviterId
	}
	return 0
}

func (x *Event_UserInvited) GetInviteId() string {
	if x != nil {
		return x.InviteId
	}
	return ""
}

// Event sent when a user rejects a guild invitation.
type Event_UserRejectedInvite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Guild ID of the guild the invitee rejected an invite for.
	GuildId uint64 `protobuf:"varint,1,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
	// User ID of the invitee that rejected the invitation.
	UserId uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Invite ID of the invite that was rejected.
	InviteId string `protobuf:"bytes,3,opt,name=invite_id,json=inviteId,proto3" json:"invite_id,omitempty"`
}

func (x *Event_UserRejectedInvite) Reset() {
	*x = Event_UserRejectedInvite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_v1_sync_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event_UserRejectedInvite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event_UserRejectedInvite) ProtoMessage() {}

func (x *Event_UserRejectedInvite) ProtoReflect() protoreflect.Message {
	mi := &file_sync_v1_sync_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event_UserRejectedInvite.ProtoReflect.Descriptor instead.
func (*Event_UserRejectedInvite) Descriptor() ([]byte, []int) {
	return file_sync_v1_sync_proto_rawDescGZIP(), []int{1, 3}
}

func (x *Event_UserRejectedInvite) GetGuildId() uint64 {
	if x != nil {
		return x.GuildId
	}
	return 0
}

func (x *Event_UserRejectedInvite) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Event_UserRejectedInvite) GetInviteId() string {
	if x != nil {
		return x.InviteId
	}
	return ""
}

var File_sync_v1_sync_proto protoreflect.FileDescriptor

var file_sync_v1_sync_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x22, 0x3b, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x22, 0xda, 0x05, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x65, 0x0a,
	0x17, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x48, 0x00, 0x52, 0x14,
	0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x47,
	0x75, 0x69, 0x6c, 0x64, 0x12, 0x59, 0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64,
	0x65, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x79, 0x6e,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x65, 0x64, 0x54, 0x6f, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x48, 0x00, 0x52, 0x10, 0x75,
	0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x65, 0x64, 0x54, 0x6f, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x12,
	0x48, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0b, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64, 0x12, 0x5e, 0x0a, 0x14, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x48, 0x00, 0x52, 0x12, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6a, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x1a, 0x4a, 0x0a, 0x14, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x75, 0x69, 0x6c,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x75,
	0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x75,
	0x69, 0x6c, 0x64, 0x49, 0x64, 0x1a, 0x46, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x65, 0x64, 0x54, 0x6f, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x1a, 0x62, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x49,
	0x64, 0x1a, 0x65, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x75, 0x69, 0x6c, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x49, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x48, 0x0a, 0x0c, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x38, 0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x22, 0x3c, 0x0a, 0x0b, 0x50, 0x75, 0x73,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x4e, 0x65, 0x77, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0d, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x15, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4e, 0x65, 0x77, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x80, 0x02, 0x0a, 0x0e, 0x50, 0x6f, 0x73,
	0x74, 0x62, 0x6f, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x04, 0x50,
	0x75, 0x6c, 0x6c, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x79,
	0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x1d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a,
	0x0b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4e, 0x65, 0x77, 0x49, 0x64, 0x12, 0x24, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4e, 0x65, 0x77, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x79,
	0x6e, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4e, 0x65, 0x77, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xbd, 0x01, 0x0a, 0x14,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x79, 0x6e,
	0x63, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61,
	0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2d, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x79, 0x6e,
	0x63, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x79, 0x6e, 0x63, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x53,
	0x58, 0xaa, 0x02, 0x10, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x79, 0x6e,
	0x63, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5c,
	0x53, 0x79, 0x6e, 0x63, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x5c, 0x53, 0x79, 0x6e, 0x63, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x3a, 0x3a, 0x53, 0x79, 0x6e, 0x63, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_sync_v1_sync_proto_rawDescOnce sync.Once
	file_sync_v1_sync_proto_rawDescData = file_sync_v1_sync_proto_rawDesc
)

func file_sync_v1_sync_proto_rawDescGZIP() []byte {
	file_sync_v1_sync_proto_rawDescOnce.Do(func() {
		file_sync_v1_sync_proto_rawDescData = protoimpl.X.CompressGZIP(file_sync_v1_sync_proto_rawDescData)
	})
	return file_sync_v1_sync_proto_rawDescData
}

var file_sync_v1_sync_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_sync_v1_sync_proto_goTypes = []interface{}{
	(*AuthData)(nil),                   // 0: protocol.sync.v1.AuthData
	(*Event)(nil),                      // 1: protocol.sync.v1.Event
	(*PullRequest)(nil),                // 2: protocol.sync.v1.PullRequest
	(*PullResponse)(nil),               // 3: protocol.sync.v1.PullResponse
	(*PushRequest)(nil),                // 4: protocol.sync.v1.PushRequest
	(*PushResponse)(nil),               // 5: protocol.sync.v1.PushResponse
	(*NotifyNewIdRequest)(nil),         // 6: protocol.sync.v1.NotifyNewIdRequest
	(*NotifyNewIdResponse)(nil),        // 7: protocol.sync.v1.NotifyNewIdResponse
	(*Event_UserRemovedFromGuild)(nil), // 8: protocol.sync.v1.Event.UserRemovedFromGuild
	(*Event_UserAddedToGuild)(nil),     // 9: protocol.sync.v1.Event.UserAddedToGuild
	(*Event_UserInvited)(nil),          // 10: protocol.sync.v1.Event.UserInvited
	(*Event_UserRejectedInvite)(nil),   // 11: protocol.sync.v1.Event.UserRejectedInvite
}
var file_sync_v1_sync_proto_depIdxs = []int32{
	8,  // 0: protocol.sync.v1.Event.user_removed_from_guild:type_name -> protocol.sync.v1.Event.UserRemovedFromGuild
	9,  // 1: protocol.sync.v1.Event.user_added_to_guild:type_name -> protocol.sync.v1.Event.UserAddedToGuild
	10, // 2: protocol.sync.v1.Event.user_invited:type_name -> protocol.sync.v1.Event.UserInvited
	11, // 3: protocol.sync.v1.Event.user_rejected_invite:type_name -> protocol.sync.v1.Event.UserRejectedInvite
	1,  // 4: protocol.sync.v1.PullResponse.event_queue:type_name -> protocol.sync.v1.Event
	1,  // 5: protocol.sync.v1.PushRequest.event:type_name -> protocol.sync.v1.Event
	2,  // 6: protocol.sync.v1.PostboxService.Pull:input_type -> protocol.sync.v1.PullRequest
	4,  // 7: protocol.sync.v1.PostboxService.Push:input_type -> protocol.sync.v1.PushRequest
	6,  // 8: protocol.sync.v1.PostboxService.NotifyNewId:input_type -> protocol.sync.v1.NotifyNewIdRequest
	3,  // 9: protocol.sync.v1.PostboxService.Pull:output_type -> protocol.sync.v1.PullResponse
	5,  // 10: protocol.sync.v1.PostboxService.Push:output_type -> protocol.sync.v1.PushResponse
	7,  // 11: protocol.sync.v1.PostboxService.NotifyNewId:output_type -> protocol.sync.v1.NotifyNewIdResponse
	9,  // [9:12] is the sub-list for method output_type
	6,  // [6:9] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_sync_v1_sync_proto_init() }
func file_sync_v1_sync_proto_init() {
	if File_sync_v1_sync_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sync_v1_sync_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthData); i {
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
		file_sync_v1_sync_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_sync_v1_sync_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullRequest); i {
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
		file_sync_v1_sync_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullResponse); i {
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
		file_sync_v1_sync_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRequest); i {
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
		file_sync_v1_sync_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushResponse); i {
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
		file_sync_v1_sync_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyNewIdRequest); i {
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
		file_sync_v1_sync_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyNewIdResponse); i {
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
		file_sync_v1_sync_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event_UserRemovedFromGuild); i {
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
		file_sync_v1_sync_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event_UserAddedToGuild); i {
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
		file_sync_v1_sync_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event_UserInvited); i {
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
		file_sync_v1_sync_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event_UserRejectedInvite); i {
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
	file_sync_v1_sync_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Event_UserRemovedFromGuild_)(nil),
		(*Event_UserAddedToGuild_)(nil),
		(*Event_UserInvited_)(nil),
		(*Event_UserRejectedInvite_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sync_v1_sync_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sync_v1_sync_proto_goTypes,
		DependencyIndexes: file_sync_v1_sync_proto_depIdxs,
		MessageInfos:      file_sync_v1_sync_proto_msgTypes,
	}.Build()
	File_sync_v1_sync_proto = out.File
	file_sync_v1_sync_proto_rawDesc = nil
	file_sync_v1_sync_proto_goTypes = nil
	file_sync_v1_sync_proto_depIdxs = nil
}
