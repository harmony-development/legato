package profilev1
import (
	"context"
	goserver "github.com/harmony-development/hrpc/pkg/go-server"
)
type ProfileService[T context.Context] interface {
	GetProfile(T, *GetProfileRequest) (*GetProfileResponse, error)
	UpdateProfile(T, *UpdateProfileRequest) (*UpdateProfileResponse, error)
	UpdateStatus(T, *UpdateStatusRequest) (*UpdateStatusResponse, error)
	GetAppData(T, *GetAppDataRequest) (*GetAppDataResponse, error)
	SetAppData(T, *SetAppDataRequest) (*SetAppDataResponse, error)
}
type ProfileServiceHandler[T context.Context] struct {
	Impl ProfileService[T]
}
func (h *ProfileServiceHandler[T]) Routes() map[string]goserver.UnaryMethodData[T] {
	return map[string]goserver.UnaryMethodData[T]{
		"protocol.profile.v1.ProfileService/GetProfile": {
		MethodData: goserver.MethodData{Input: &GetProfileRequest{}, Output: &GetProfileResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GetProfile(c, msg.(*GetProfileRequest))
					},
				},
				
		"protocol.profile.v1.ProfileService/UpdateProfile": {
		MethodData: goserver.MethodData{Input: &UpdateProfileRequest{}, Output: &UpdateProfileResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.UpdateProfile(c, msg.(*UpdateProfileRequest))
					},
				},
				
		"protocol.profile.v1.ProfileService/UpdateStatus": {
		MethodData: goserver.MethodData{Input: &UpdateStatusRequest{}, Output: &UpdateStatusResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.UpdateStatus(c, msg.(*UpdateStatusRequest))
					},
				},
				
		"protocol.profile.v1.ProfileService/GetAppData": {
		MethodData: goserver.MethodData{Input: &GetAppDataRequest{}, Output: &GetAppDataResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.GetAppData(c, msg.(*GetAppDataRequest))
					},
				},
				
		"protocol.profile.v1.ProfileService/SetAppData": {
		MethodData: goserver.MethodData{Input: &SetAppDataRequest{}, Output: &SetAppDataResponse{}},
		Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.Impl.SetAppData(c, msg.(*SetAppDataRequest))
					},
				},
				
	}
}
func (h *ProfileServiceHandler[T]) StreamRoutes() map[string]goserver.StreamMethodData[T] {
	return map[string]goserver.StreamMethodData[T]{
	}
	}
