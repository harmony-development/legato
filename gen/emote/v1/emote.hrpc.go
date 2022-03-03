package emotev1
import (
	"context"
	goserver "github.com/harmony-development/hrpc/pkg/go-server"
)
type EmoteService[T context.Context] interface {
	CreateEmotePack(T, *CreateEmotePackRequest) (*CreateEmotePackResponse, error)
	GetEmotePacks(T, *GetEmotePacksRequest) (*GetEmotePacksResponse, error)
	GetEmotePackEmotes(T, *GetEmotePackEmotesRequest) (*GetEmotePackEmotesResponse, error)
	AddEmoteToPack(T, *AddEmoteToPackRequest) (*AddEmoteToPackResponse, error)
	DeleteEmotePack(T, *DeleteEmotePackRequest) (*DeleteEmotePackResponse, error)
	DeleteEmoteFromPack(T, *DeleteEmoteFromPackRequest) (*DeleteEmoteFromPackResponse, error)
	DequipEmotePack(T, *DequipEmotePackRequest) (*DequipEmotePackResponse, error)
	EquipEmotePack(T, *EquipEmotePackRequest) (*EquipEmotePackResponse, error)
}
type EmoteServiceHandler[T context.Context] struct {
	Impl EmoteService[T]
}
func (h *EmoteServiceHandler[T]) Routes() map[string]goserver.UnaryMethodData[T] {
	return map[string]goserver.UnaryMethodData[T]{
		"protocol.emote.v1.EmoteService/CreateEmotePack": {
		MethodData: goserver.MethodData{Input: &CreateEmotePackRequest{}, Output: &CreateEmotePackResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.CreateEmotePack(c, msg.(*CreateEmotePackRequest))
					},
				},
				
		"protocol.emote.v1.EmoteService/GetEmotePacks": {
		MethodData: goserver.MethodData{Input: &GetEmotePacksRequest{}, Output: &GetEmotePacksResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GetEmotePacks(c, msg.(*GetEmotePacksRequest))
					},
				},
				
		"protocol.emote.v1.EmoteService/GetEmotePackEmotes": {
		MethodData: goserver.MethodData{Input: &GetEmotePackEmotesRequest{}, Output: &GetEmotePackEmotesResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GetEmotePackEmotes(c, msg.(*GetEmotePackEmotesRequest))
					},
				},
				
		"protocol.emote.v1.EmoteService/AddEmoteToPack": {
		MethodData: goserver.MethodData{Input: &AddEmoteToPackRequest{}, Output: &AddEmoteToPackResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.AddEmoteToPack(c, msg.(*AddEmoteToPackRequest))
					},
				},
				
		"protocol.emote.v1.EmoteService/DeleteEmotePack": {
		MethodData: goserver.MethodData{Input: &DeleteEmotePackRequest{}, Output: &DeleteEmotePackResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.DeleteEmotePack(c, msg.(*DeleteEmotePackRequest))
					},
				},
				
		"protocol.emote.v1.EmoteService/DeleteEmoteFromPack": {
		MethodData: goserver.MethodData{Input: &DeleteEmoteFromPackRequest{}, Output: &DeleteEmoteFromPackResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.DeleteEmoteFromPack(c, msg.(*DeleteEmoteFromPackRequest))
					},
				},
				
		"protocol.emote.v1.EmoteService/DequipEmotePack": {
		MethodData: goserver.MethodData{Input: &DequipEmotePackRequest{}, Output: &DequipEmotePackResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.DequipEmotePack(c, msg.(*DequipEmotePackRequest))
					},
				},
				
		"protocol.emote.v1.EmoteService/EquipEmotePack": {
		MethodData: goserver.MethodData{Input: &EquipEmotePackRequest{}, Output: &EquipEmotePackResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.EquipEmotePack(c, msg.(*EquipEmotePackRequest))
					},
				},
				
	}
}
func (h *EmoteServiceHandler[T]) StreamRoutes() map[string]goserver.StreamMethodData[T] {
	return map[string]goserver.StreamMethodData[T]{
	}
	}
