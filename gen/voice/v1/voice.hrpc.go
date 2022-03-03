package voicev1
import (
	"context"
	goserver "github.com/harmony-development/hrpc/pkg/go-server"
)
type VoiceService[T context.Context] interface {
	StreamMessage(T, chan *StreamMessageRequest) (chan *StreamMessageResponse, error)
}
type VoiceServiceHandler[T context.Context] struct {
	Impl VoiceService[T]
}
func (h *VoiceServiceHandler[T]) Routes() map[string]goserver.UnaryMethodData[T] {
	return map[string]goserver.UnaryMethodData[T]{
	}
}
func (h *VoiceServiceHandler[T]) StreamRoutes() map[string]goserver.StreamMethodData[T] {
	return map[string]goserver.StreamMethodData[T]{
		"protocol.voice.v1.VoiceService/StreamMessage": {
			MethodData: goserver.MethodData{Input: &StreamMessageRequest{}, Output: &StreamMessageResponse{}},
			Handler: func(c T, msg chan goserver.VTProtoMessage) (chan goserver.VTProtoMessage, error) {
					res, err := h.Impl.StreamMessage(c, goserver.FromProtoChannel[*StreamMessageRequest](msg))
					return goserver.ToProtoChannel(res), err
				},
		},
	}
	}
