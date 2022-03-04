package chatv1
import (
	"context"
	goserver "github.com/harmony-development/hrpc/pkg/go-server"
)
type ChatService[T context.Context] interface {
	CreateGuild(T, *CreateGuildRequest) (*CreateGuildResponse, error)
	CreateRoom(T, *CreateRoomRequest) (*CreateRoomResponse, error)
	CreateDirectMessage(T, *CreateDirectMessageRequest) (*CreateDirectMessageResponse, error)
	UpgradeRoomToGuild(T, *UpgradeRoomToGuildRequest) (*UpgradeRoomToGuildResponse, error)
	CreateInvite(T, *CreateInviteRequest) (*CreateInviteResponse, error)
	CreateChannel(T, *CreateChannelRequest) (*CreateChannelResponse, error)
	GetGuildList(T, *GetGuildListRequest) (*GetGuildListResponse, error)
	InviteUserToGuild(T, *InviteUserToGuildRequest) (*InviteUserToGuildResponse, error)
	GetPendingInvites(T, *GetPendingInvitesRequest) (*GetPendingInvitesResponse, error)
	RejectPendingInvite(T, *RejectPendingInviteRequest) (*RejectPendingInviteResponse, error)
	IgnorePendingInvite(T, *IgnorePendingInviteRequest) (*IgnorePendingInviteResponse, error)
	GetGuild(T, *GetGuildRequest) (*GetGuildResponse, error)
	GetGuildInvites(T, *GetGuildInvitesRequest) (*GetGuildInvitesResponse, error)
	GetGuildMembers(T, *GetGuildMembersRequest) (*GetGuildMembersResponse, error)
	GetGuildChannels(T, *GetGuildChannelsRequest) (*GetGuildChannelsResponse, error)
	GetChannelMessages(T, *GetChannelMessagesRequest) (*GetChannelMessagesResponse, error)
	GetMessage(T, *GetMessageRequest) (*GetMessageResponse, error)
	UpdateGuildInformation(T, *UpdateGuildInformationRequest) (*UpdateGuildInformationResponse, error)
	UpdateChannelInformation(T, *UpdateChannelInformationRequest) (*UpdateChannelInformationResponse, error)
	UpdateChannelOrder(T, *UpdateChannelOrderRequest) (*UpdateChannelOrderResponse, error)
	UpdateAllChannelOrder(T, *UpdateAllChannelOrderRequest) (*UpdateAllChannelOrderResponse, error)
	UpdateMessageContent(T, *UpdateMessageContentRequest) (*UpdateMessageContentResponse, error)
	DeleteGuild(T, *DeleteGuildRequest) (*DeleteGuildResponse, error)
	DeleteInvite(T, *DeleteInviteRequest) (*DeleteInviteResponse, error)
	DeleteChannel(T, *DeleteChannelRequest) (*DeleteChannelResponse, error)
	DeleteMessage(T, *DeleteMessageRequest) (*DeleteMessageResponse, error)
	JoinGuild(T, *JoinGuildRequest) (*JoinGuildResponse, error)
	LeaveGuild(T, *LeaveGuildRequest) (*LeaveGuildResponse, error)
	TriggerAction(T, *TriggerActionRequest) (*TriggerActionResponse, error)
	SendMessage(T, *SendMessageRequest) (*SendMessageResponse, error)
	HasPermission(T, *HasPermissionRequest) (*HasPermissionResponse, error)
	SetPermissions(T, *SetPermissionsRequest) (*SetPermissionsResponse, error)
	GetPermissions(T, *GetPermissionsRequest) (*GetPermissionsResponse, error)
	MoveRole(T, *MoveRoleRequest) (*MoveRoleResponse, error)
	GetGuildRoles(T, *GetGuildRolesRequest) (*GetGuildRolesResponse, error)
	AddGuildRole(T, *AddGuildRoleRequest) (*AddGuildRoleResponse, error)
	ModifyGuildRole(T, *ModifyGuildRoleRequest) (*ModifyGuildRoleResponse, error)
	DeleteGuildRole(T, *DeleteGuildRoleRequest) (*DeleteGuildRoleResponse, error)
	ManageUserRoles(T, *ManageUserRolesRequest) (*ManageUserRolesResponse, error)
	GetUserRoles(T, *GetUserRolesRequest) (*GetUserRolesResponse, error)
	Typing(T, *TypingRequest) (*TypingResponse, error)
	PreviewGuild(T, *PreviewGuildRequest) (*PreviewGuildResponse, error)
	GetBannedUsers(T, *GetBannedUsersRequest) (*GetBannedUsersResponse, error)
	BanUser(T, *BanUserRequest) (*BanUserResponse, error)
	KickUser(T, *KickUserRequest) (*KickUserResponse, error)
	UnbanUser(T, *UnbanUserRequest) (*UnbanUserResponse, error)
	GetPinnedMessages(T, *GetPinnedMessagesRequest) (*GetPinnedMessagesResponse, error)
	PinMessage(T, *PinMessageRequest) (*PinMessageResponse, error)
	UnpinMessage(T, *UnpinMessageRequest) (*UnpinMessageResponse, error)
	AddReaction(T, *AddReactionRequest) (*AddReactionResponse, error)
	RemoveReaction(T, *RemoveReactionRequest) (*RemoveReactionResponse, error)
	GrantOwnership(T, *GrantOwnershipRequest) (*GrantOwnershipResponse, error)
	GiveUpOwnership(T, *GiveUpOwnershipRequest) (*GiveUpOwnershipResponse, error)
	StreamEvents(T, chan *StreamEventsRequest, chan *StreamEventsResponse) error
}
type ChatServiceHandler[T context.Context] struct {
	Impl ChatService[T]
}
func (h *ChatServiceHandler[T]) Routes() map[string]goserver.UnaryMethodData[T] {
	return map[string]goserver.UnaryMethodData[T]{
		"protocol.chat.v1.ChatService/CreateGuild": {
		MethodData: goserver.MethodData{Input: &CreateGuildRequest{}, Output: &CreateGuildResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.CreateGuild(c, msg.(*CreateGuildRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/CreateRoom": {
		MethodData: goserver.MethodData{Input: &CreateRoomRequest{}, Output: &CreateRoomResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.CreateRoom(c, msg.(*CreateRoomRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/CreateDirectMessage": {
		MethodData: goserver.MethodData{Input: &CreateDirectMessageRequest{}, Output: &CreateDirectMessageResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.CreateDirectMessage(c, msg.(*CreateDirectMessageRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/UpgradeRoomToGuild": {
		MethodData: goserver.MethodData{Input: &UpgradeRoomToGuildRequest{}, Output: &UpgradeRoomToGuildResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.UpgradeRoomToGuild(c, msg.(*UpgradeRoomToGuildRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/CreateInvite": {
		MethodData: goserver.MethodData{Input: &CreateInviteRequest{}, Output: &CreateInviteResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.CreateInvite(c, msg.(*CreateInviteRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/CreateChannel": {
		MethodData: goserver.MethodData{Input: &CreateChannelRequest{}, Output: &CreateChannelResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.CreateChannel(c, msg.(*CreateChannelRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/GetGuildList": {
		MethodData: goserver.MethodData{Input: &GetGuildListRequest{}, Output: &GetGuildListResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GetGuildList(c, msg.(*GetGuildListRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/InviteUserToGuild": {
		MethodData: goserver.MethodData{Input: &InviteUserToGuildRequest{}, Output: &InviteUserToGuildResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.InviteUserToGuild(c, msg.(*InviteUserToGuildRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/GetPendingInvites": {
		MethodData: goserver.MethodData{Input: &GetPendingInvitesRequest{}, Output: &GetPendingInvitesResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GetPendingInvites(c, msg.(*GetPendingInvitesRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/RejectPendingInvite": {
		MethodData: goserver.MethodData{Input: &RejectPendingInviteRequest{}, Output: &RejectPendingInviteResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.RejectPendingInvite(c, msg.(*RejectPendingInviteRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/IgnorePendingInvite": {
		MethodData: goserver.MethodData{Input: &IgnorePendingInviteRequest{}, Output: &IgnorePendingInviteResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.IgnorePendingInvite(c, msg.(*IgnorePendingInviteRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/GetGuild": {
		MethodData: goserver.MethodData{Input: &GetGuildRequest{}, Output: &GetGuildResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GetGuild(c, msg.(*GetGuildRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/GetGuildInvites": {
		MethodData: goserver.MethodData{Input: &GetGuildInvitesRequest{}, Output: &GetGuildInvitesResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GetGuildInvites(c, msg.(*GetGuildInvitesRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/GetGuildMembers": {
		MethodData: goserver.MethodData{Input: &GetGuildMembersRequest{}, Output: &GetGuildMembersResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GetGuildMembers(c, msg.(*GetGuildMembersRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/GetGuildChannels": {
		MethodData: goserver.MethodData{Input: &GetGuildChannelsRequest{}, Output: &GetGuildChannelsResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GetGuildChannels(c, msg.(*GetGuildChannelsRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/GetChannelMessages": {
		MethodData: goserver.MethodData{Input: &GetChannelMessagesRequest{}, Output: &GetChannelMessagesResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GetChannelMessages(c, msg.(*GetChannelMessagesRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/GetMessage": {
		MethodData: goserver.MethodData{Input: &GetMessageRequest{}, Output: &GetMessageResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GetMessage(c, msg.(*GetMessageRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/UpdateGuildInformation": {
		MethodData: goserver.MethodData{Input: &UpdateGuildInformationRequest{}, Output: &UpdateGuildInformationResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.UpdateGuildInformation(c, msg.(*UpdateGuildInformationRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/UpdateChannelInformation": {
		MethodData: goserver.MethodData{Input: &UpdateChannelInformationRequest{}, Output: &UpdateChannelInformationResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.UpdateChannelInformation(c, msg.(*UpdateChannelInformationRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/UpdateChannelOrder": {
		MethodData: goserver.MethodData{Input: &UpdateChannelOrderRequest{}, Output: &UpdateChannelOrderResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.UpdateChannelOrder(c, msg.(*UpdateChannelOrderRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/UpdateAllChannelOrder": {
		MethodData: goserver.MethodData{Input: &UpdateAllChannelOrderRequest{}, Output: &UpdateAllChannelOrderResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.UpdateAllChannelOrder(c, msg.(*UpdateAllChannelOrderRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/UpdateMessageContent": {
		MethodData: goserver.MethodData{Input: &UpdateMessageContentRequest{}, Output: &UpdateMessageContentResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.UpdateMessageContent(c, msg.(*UpdateMessageContentRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/DeleteGuild": {
		MethodData: goserver.MethodData{Input: &DeleteGuildRequest{}, Output: &DeleteGuildResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.DeleteGuild(c, msg.(*DeleteGuildRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/DeleteInvite": {
		MethodData: goserver.MethodData{Input: &DeleteInviteRequest{}, Output: &DeleteInviteResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.DeleteInvite(c, msg.(*DeleteInviteRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/DeleteChannel": {
		MethodData: goserver.MethodData{Input: &DeleteChannelRequest{}, Output: &DeleteChannelResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.DeleteChannel(c, msg.(*DeleteChannelRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/DeleteMessage": {
		MethodData: goserver.MethodData{Input: &DeleteMessageRequest{}, Output: &DeleteMessageResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.DeleteMessage(c, msg.(*DeleteMessageRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/JoinGuild": {
		MethodData: goserver.MethodData{Input: &JoinGuildRequest{}, Output: &JoinGuildResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.JoinGuild(c, msg.(*JoinGuildRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/LeaveGuild": {
		MethodData: goserver.MethodData{Input: &LeaveGuildRequest{}, Output: &LeaveGuildResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.LeaveGuild(c, msg.(*LeaveGuildRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/TriggerAction": {
		MethodData: goserver.MethodData{Input: &TriggerActionRequest{}, Output: &TriggerActionResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.TriggerAction(c, msg.(*TriggerActionRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/SendMessage": {
		MethodData: goserver.MethodData{Input: &SendMessageRequest{}, Output: &SendMessageResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.SendMessage(c, msg.(*SendMessageRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/HasPermission": {
		MethodData: goserver.MethodData{Input: &HasPermissionRequest{}, Output: &HasPermissionResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.HasPermission(c, msg.(*HasPermissionRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/SetPermissions": {
		MethodData: goserver.MethodData{Input: &SetPermissionsRequest{}, Output: &SetPermissionsResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.SetPermissions(c, msg.(*SetPermissionsRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/GetPermissions": {
		MethodData: goserver.MethodData{Input: &GetPermissionsRequest{}, Output: &GetPermissionsResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GetPermissions(c, msg.(*GetPermissionsRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/MoveRole": {
		MethodData: goserver.MethodData{Input: &MoveRoleRequest{}, Output: &MoveRoleResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.MoveRole(c, msg.(*MoveRoleRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/GetGuildRoles": {
		MethodData: goserver.MethodData{Input: &GetGuildRolesRequest{}, Output: &GetGuildRolesResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GetGuildRoles(c, msg.(*GetGuildRolesRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/AddGuildRole": {
		MethodData: goserver.MethodData{Input: &AddGuildRoleRequest{}, Output: &AddGuildRoleResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.AddGuildRole(c, msg.(*AddGuildRoleRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/ModifyGuildRole": {
		MethodData: goserver.MethodData{Input: &ModifyGuildRoleRequest{}, Output: &ModifyGuildRoleResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.ModifyGuildRole(c, msg.(*ModifyGuildRoleRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/DeleteGuildRole": {
		MethodData: goserver.MethodData{Input: &DeleteGuildRoleRequest{}, Output: &DeleteGuildRoleResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.DeleteGuildRole(c, msg.(*DeleteGuildRoleRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/ManageUserRoles": {
		MethodData: goserver.MethodData{Input: &ManageUserRolesRequest{}, Output: &ManageUserRolesResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.ManageUserRoles(c, msg.(*ManageUserRolesRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/GetUserRoles": {
		MethodData: goserver.MethodData{Input: &GetUserRolesRequest{}, Output: &GetUserRolesResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GetUserRoles(c, msg.(*GetUserRolesRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/Typing": {
		MethodData: goserver.MethodData{Input: &TypingRequest{}, Output: &TypingResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.Typing(c, msg.(*TypingRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/PreviewGuild": {
		MethodData: goserver.MethodData{Input: &PreviewGuildRequest{}, Output: &PreviewGuildResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.PreviewGuild(c, msg.(*PreviewGuildRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/GetBannedUsers": {
		MethodData: goserver.MethodData{Input: &GetBannedUsersRequest{}, Output: &GetBannedUsersResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GetBannedUsers(c, msg.(*GetBannedUsersRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/BanUser": {
		MethodData: goserver.MethodData{Input: &BanUserRequest{}, Output: &BanUserResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.BanUser(c, msg.(*BanUserRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/KickUser": {
		MethodData: goserver.MethodData{Input: &KickUserRequest{}, Output: &KickUserResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.KickUser(c, msg.(*KickUserRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/UnbanUser": {
		MethodData: goserver.MethodData{Input: &UnbanUserRequest{}, Output: &UnbanUserResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.UnbanUser(c, msg.(*UnbanUserRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/GetPinnedMessages": {
		MethodData: goserver.MethodData{Input: &GetPinnedMessagesRequest{}, Output: &GetPinnedMessagesResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GetPinnedMessages(c, msg.(*GetPinnedMessagesRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/PinMessage": {
		MethodData: goserver.MethodData{Input: &PinMessageRequest{}, Output: &PinMessageResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.PinMessage(c, msg.(*PinMessageRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/UnpinMessage": {
		MethodData: goserver.MethodData{Input: &UnpinMessageRequest{}, Output: &UnpinMessageResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.UnpinMessage(c, msg.(*UnpinMessageRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/AddReaction": {
		MethodData: goserver.MethodData{Input: &AddReactionRequest{}, Output: &AddReactionResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.AddReaction(c, msg.(*AddReactionRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/RemoveReaction": {
		MethodData: goserver.MethodData{Input: &RemoveReactionRequest{}, Output: &RemoveReactionResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.RemoveReaction(c, msg.(*RemoveReactionRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/GrantOwnership": {
		MethodData: goserver.MethodData{Input: &GrantOwnershipRequest{}, Output: &GrantOwnershipResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GrantOwnership(c, msg.(*GrantOwnershipRequest))
					},
				},
				
		"protocol.chat.v1.ChatService/GiveUpOwnership": {
		MethodData: goserver.MethodData{Input: &GiveUpOwnershipRequest{}, Output: &GiveUpOwnershipResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GiveUpOwnership(c, msg.(*GiveUpOwnershipRequest))
					},
				},
				
	}
}
func (h *ChatServiceHandler[T]) StreamRoutes() map[string]goserver.StreamMethodData[T] {
	return map[string]goserver.StreamMethodData[T]{
		"protocol.chat.v1.ChatService/StreamEvents": {
			MethodData: goserver.MethodData{Input: &StreamEventsRequest{}, Output: &StreamEventsResponse{}},
			Handler: func(c T, msg chan goserver.VTProtoMessage, res chan goserver.VTProtoMessage) (error) {
					ch := make(chan *StreamEventsResponse)
					goserver.ToProtoChannel(ch, res)
					err := h.Impl.StreamEvents(c, goserver.FromProtoChannel[*StreamEventsRequest](msg), ch)
					return err
				},
		},
	}
	}
