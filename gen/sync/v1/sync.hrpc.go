package syncv1
import (
	"context"
	goserver "github.com/harmony-development/hrpc/pkg/go-server"
)
type PostboxService[T context.Context] interface {
	Pull(T, *PullRequest) (*PullResponse, error)
	Push(T, *PushRequest) (*PushResponse, error)
	NotifyNewId(T, *NotifyNewIdRequest) (*NotifyNewIdResponse, error)
}
type PostboxServiceHandler[T context.Context] struct {
	Impl PostboxService[T]
}
func (h *PostboxServiceHandler[T]) Routes() map[string]goserver.UnaryMethodData[T] {
	return map[string]goserver.UnaryMethodData[T]{
		"protocol.sync.v1.PostboxService/Pull": {
		MethodData: goserver.MethodData{Input: &PullRequest{}, Output: &PullResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.Pull(c, msg.(*PullRequest))
					},
				},
				
		"protocol.sync.v1.PostboxService/Push": {
		MethodData: goserver.MethodData{Input: &PushRequest{}, Output: &PushResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.Push(c, msg.(*PushRequest))
					},
				},
				
		"protocol.sync.v1.PostboxService/NotifyNewId": {
		MethodData: goserver.MethodData{Input: &NotifyNewIdRequest{}, Output: &NotifyNewIdResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.NotifyNewId(c, msg.(*NotifyNewIdRequest))
					},
				},
				
	}
}
func (h *PostboxServiceHandler[T]) StreamRoutes() map[string]goserver.StreamMethodData[T] {
	return map[string]goserver.StreamMethodData[T]{
	}
	}
