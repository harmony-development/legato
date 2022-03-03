package mediaproxyv1
import (
	"context"
	goserver "github.com/harmony-development/hrpc/pkg/go-server"
)
type MediaProxyService[T context.Context] interface {
	FetchLinkMetadata(T, *FetchLinkMetadataRequest) (*FetchLinkMetadataResponse, error)
	InstantView(T, *InstantViewRequest) (*InstantViewResponse, error)
	CanInstantView(T, *CanInstantViewRequest) (*CanInstantViewResponse, error)
}
type MediaProxyServiceHandler[T context.Context] struct {
	Impl MediaProxyService[T]
}
func (h *MediaProxyServiceHandler[T]) Routes() map[string]goserver.UnaryMethodData[T] {
	return map[string]goserver.UnaryMethodData[T]{
		"protocol.mediaproxy.v1.MediaProxyService/FetchLinkMetadata": {
		MethodData: goserver.MethodData{Input: &FetchLinkMetadataRequest{}, Output: &FetchLinkMetadataResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.FetchLinkMetadata(c, msg.(*FetchLinkMetadataRequest))
					},
				},
				
		"protocol.mediaproxy.v1.MediaProxyService/InstantView": {
		MethodData: goserver.MethodData{Input: &InstantViewRequest{}, Output: &InstantViewResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.InstantView(c, msg.(*InstantViewRequest))
					},
				},
				
		"protocol.mediaproxy.v1.MediaProxyService/CanInstantView": {
		MethodData: goserver.MethodData{Input: &CanInstantViewRequest{}, Output: &CanInstantViewResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.CanInstantView(c, msg.(*CanInstantViewRequest))
					},
				},
				
	}
}
func (h *MediaProxyServiceHandler[T]) StreamRoutes() map[string]goserver.StreamMethodData[T] {
	return map[string]goserver.StreamMethodData[T]{
	}
	}
