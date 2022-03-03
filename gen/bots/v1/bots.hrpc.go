package botsv1
import (
	"context"
	goserver "github.com/harmony-development/hrpc/pkg/go-server"
)
type BotsService[T context.Context] interface {
	MyBots(T, *MyBotsRequest) (*MyBotsResponse, error)
	GetBot(T, *GetBotRequest) (*GetBotResponse, error)
	CreateBot(T, *CreateBotRequest) (*CreateBotResponse, error)
	EditBot(T, *EditBotRequest) (*EditBotResponse, error)
	DeleteBot(T, *DeleteBotRequest) (*DeleteBotResponse, error)
	Policies(T, *PoliciesRequest) (*PoliciesResponse, error)
	AddBot(T, *AddBotRequest) (*AddBotResponse, error)
}
type BotsServiceHandler[T context.Context] struct {
	Impl BotsService[T]
}
func (h *BotsServiceHandler[T]) Routes() map[string]goserver.UnaryMethodData[T] {
	return map[string]goserver.UnaryMethodData[T]{
		"protocol.bots.v1.BotsService/MyBots": {
		MethodData: goserver.MethodData{Input: &MyBotsRequest{}, Output: &MyBotsResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.MyBots(c, msg.(*MyBotsRequest))
					},
				},
				
		"protocol.bots.v1.BotsService/GetBot": {
		MethodData: goserver.MethodData{Input: &GetBotRequest{}, Output: &GetBotResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GetBot(c, msg.(*GetBotRequest))
					},
				},
				
		"protocol.bots.v1.BotsService/CreateBot": {
		MethodData: goserver.MethodData{Input: &CreateBotRequest{}, Output: &CreateBotResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.CreateBot(c, msg.(*CreateBotRequest))
					},
				},
				
		"protocol.bots.v1.BotsService/EditBot": {
		MethodData: goserver.MethodData{Input: &EditBotRequest{}, Output: &EditBotResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.EditBot(c, msg.(*EditBotRequest))
					},
				},
				
		"protocol.bots.v1.BotsService/DeleteBot": {
		MethodData: goserver.MethodData{Input: &DeleteBotRequest{}, Output: &DeleteBotResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.DeleteBot(c, msg.(*DeleteBotRequest))
					},
				},
				
		"protocol.bots.v1.BotsService/Policies": {
		MethodData: goserver.MethodData{Input: &PoliciesRequest{}, Output: &PoliciesResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.Policies(c, msg.(*PoliciesRequest))
					},
				},
				
		"protocol.bots.v1.BotsService/AddBot": {
		MethodData: goserver.MethodData{Input: &AddBotRequest{}, Output: &AddBotResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.AddBot(c, msg.(*AddBotRequest))
					},
				},
				
	}
}
func (h *BotsServiceHandler[T]) StreamRoutes() map[string]goserver.StreamMethodData[T] {
	return map[string]goserver.StreamMethodData[T]{
	}
	}
