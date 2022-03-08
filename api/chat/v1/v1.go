// Package chatv1impl contains the implementation of the chat/v1 API.
package chatv1impl

import (
	"github.com/harmony-development/legato/api"
	"github.com/harmony-development/legato/db"
	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
)

type ChatV1 struct {
	db *db.DB
}

func New(db *db.DB) *ChatV1 {
	return &ChatV1{db: db}
}

func (v1 *ChatV1) CreateGuild(c api.LegatoContext, req *chatv1.CreateGuildRequest) (*chatv1.CreateGuildResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) CreateRoom(_ api.LegatoContext, _ *chatv1.CreateRoomRequest) (*chatv1.CreateRoomResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) CreateDirectMessage(_ api.LegatoContext, _ *chatv1.CreateDirectMessageRequest) (*chatv1.CreateDirectMessageResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) UpgradeRoomToGuild(_ api.LegatoContext, _ *chatv1.UpgradeRoomToGuildRequest) (*chatv1.UpgradeRoomToGuildResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) CreateInvite(_ api.LegatoContext, _ *chatv1.CreateInviteRequest) (*chatv1.CreateInviteResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) CreateChannel(_ api.LegatoContext, _ *chatv1.CreateChannelRequest) (*chatv1.CreateChannelResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) GetGuildList(_ api.LegatoContext, _ *chatv1.GetGuildListRequest) (*chatv1.GetGuildListResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) InviteUserToGuild(_ api.LegatoContext, _ *chatv1.InviteUserToGuildRequest) (*chatv1.InviteUserToGuildResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) GetPendingInvites(_ api.LegatoContext, _ *chatv1.GetPendingInvitesRequest) (*chatv1.GetPendingInvitesResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) RejectPendingInvite(_ api.LegatoContext, _ *chatv1.RejectPendingInviteRequest) (*chatv1.RejectPendingInviteResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) IgnorePendingInvite(_ api.LegatoContext, _ *chatv1.IgnorePendingInviteRequest) (*chatv1.IgnorePendingInviteResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) GetGuild(_ api.LegatoContext, _ *chatv1.GetGuildRequest) (*chatv1.GetGuildResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) GetGuildInvites(_ api.LegatoContext, _ *chatv1.GetGuildInvitesRequest) (*chatv1.GetGuildInvitesResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) GetGuildMembers(_ api.LegatoContext, _ *chatv1.GetGuildMembersRequest) (*chatv1.GetGuildMembersResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) GetGuildChannels(_ api.LegatoContext, _ *chatv1.GetGuildChannelsRequest) (*chatv1.GetGuildChannelsResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) GetChannelMessages(_ api.LegatoContext, _ *chatv1.GetChannelMessagesRequest) (*chatv1.GetChannelMessagesResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) GetMessage(_ api.LegatoContext, _ *chatv1.GetMessageRequest) (*chatv1.GetMessageResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) UpdateGuildInformation(_ api.LegatoContext, _ *chatv1.UpdateGuildInformationRequest) (*chatv1.UpdateGuildInformationResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) UpdateChannelInformation(_ api.LegatoContext, _ *chatv1.UpdateChannelInformationRequest) (*chatv1.UpdateChannelInformationResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) UpdateChannelOrder(_ api.LegatoContext, _ *chatv1.UpdateChannelOrderRequest) (*chatv1.UpdateChannelOrderResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) UpdateAllChannelOrder(_ api.LegatoContext, _ *chatv1.UpdateAllChannelOrderRequest) (*chatv1.UpdateAllChannelOrderResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) UpdateMessageContent(_ api.LegatoContext, _ *chatv1.UpdateMessageContentRequest) (*chatv1.UpdateMessageContentResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) DeleteGuild(_ api.LegatoContext, _ *chatv1.DeleteGuildRequest) (*chatv1.DeleteGuildResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) DeleteInvite(_ api.LegatoContext, _ *chatv1.DeleteInviteRequest) (*chatv1.DeleteInviteResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) DeleteChannel(_ api.LegatoContext, _ *chatv1.DeleteChannelRequest) (*chatv1.DeleteChannelResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) DeleteMessage(_ api.LegatoContext, _ *chatv1.DeleteMessageRequest) (*chatv1.DeleteMessageResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) JoinGuild(_ api.LegatoContext, _ *chatv1.JoinGuildRequest) (*chatv1.JoinGuildResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) LeaveGuild(_ api.LegatoContext, _ *chatv1.LeaveGuildRequest) (*chatv1.LeaveGuildResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) TriggerAction(_ api.LegatoContext, _ *api.LegatoContext) (*api.LegatoContext, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) SendMessage(_ api.LegatoContext, _ *chatv1.SendMessageRequest) (*chatv1.SendMessageResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) HasPermission(_ api.LegatoContext, _ *chatv1.HasPermissionRequest) (*chatv1.HasPermissionResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) SetPermissions(_ api.LegatoContext, _ *chatv1.SetPermissionsRequest) (*chatv1.SetPermissionsResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) GetPermissions(_ api.LegatoContext, _ *chatv1.GetPermissionsRequest) (*chatv1.GetPermissionsResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) MoveRole(_ api.LegatoContext, _ *chatv1.MoveRoleRequest) (*chatv1.MoveRoleResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) GetGuildRoles(_ api.LegatoContext, _ *chatv1.GetGuildRolesRequest) (*chatv1.GetGuildRolesResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) AddGuildRole(_ api.LegatoContext, _ *chatv1.AddGuildRoleRequest) (*chatv1.AddGuildRoleResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) ModifyGuildRole(_ api.LegatoContext, _ *chatv1.ModifyGuildRoleRequest) (*chatv1.ModifyGuildRoleResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) DeleteGuildRole(_ api.LegatoContext, _ *chatv1.DeleteGuildRoleRequest) (*chatv1.DeleteGuildRoleResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) ManageUserRoles(_ api.LegatoContext, _ *chatv1.ManageUserRolesRequest) (*chatv1.ManageUserRolesResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) GetUserRoles(_ api.LegatoContext, _ *chatv1.GetUserRolesRequest) (*chatv1.GetUserRolesResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) Typing(_ api.LegatoContext, _ *api.LegatoContext) (*api.LegatoContext, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) PreviewGuild(_ api.LegatoContext, _ *chatv1.PreviewGuildRequest) (*chatv1.PreviewGuildResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) GetBannedUsers(_ api.LegatoContext, _ *chatv1.GetBannedUsersRequest) (*chatv1.GetBannedUsersResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) BanUser(_ api.LegatoContext, _ *chatv1.BanUserRequest) (*chatv1.BanUserResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) KickUser(_ api.LegatoContext, _ *chatv1.KickUserRequest) (*chatv1.KickUserResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) UnbanUser(_ api.LegatoContext, _ *chatv1.UnbanUserRequest) (*chatv1.UnbanUserResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) GetPinnedMessages(_ api.LegatoContext, _ *chatv1.GetPinnedMessagesRequest) (*chatv1.GetPinnedMessagesResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) PinMessage(_ api.LegatoContext, _ *chatv1.PinMessageRequest) (*chatv1.PinMessageResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) UnpinMessage(_ api.LegatoContext, _ *chatv1.UnpinMessageRequest) (*chatv1.UnpinMessageResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) AddReaction(_ api.LegatoContext, _ *chatv1.AddReactionRequest) (*chatv1.AddReactionResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) RemoveReaction(_ api.LegatoContext, _ *chatv1.RemoveReactionRequest) (*chatv1.RemoveReactionResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) GrantOwnership(_ api.LegatoContext, _ *chatv1.GrantOwnershipRequest) (*chatv1.GrantOwnershipResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) GiveUpOwnership(_ api.LegatoContext, _ *chatv1.GiveUpOwnershipRequest) (*chatv1.GiveUpOwnershipResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *ChatV1) StreamEvents(_ api.LegatoContext, _ chan *chatv1.StreamEventsRequest, _ chan *chatv1.StreamEventsResponse) error {
	panic("not implemented") // TODO: Implement
}
