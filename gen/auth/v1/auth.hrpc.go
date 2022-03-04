package authv1

import (
	"context"

	goserver "github.com/harmony-development/hrpc/pkg/go-server"
)

type AuthService[T context.Context] interface {
	Federate(T, *FederateRequest) (*FederateResponse, error)
	LoginFederated(T, *LoginFederatedRequest) (*LoginFederatedResponse, error)
	Key(T, *KeyRequest) (*KeyResponse, error)
	BeginAuth(T, *BeginAuthRequest) (*BeginAuthResponse, error)
	NextStep(T, *NextStepRequest) (*NextStepResponse, error)
	StepBack(T, *StepBackRequest) (*StepBackResponse, error)
	CheckLoggedIn(T, *CheckLoggedInRequest) (*CheckLoggedInResponse, error)
	StreamSteps(T, chan *StreamStepsRequest, chan *StreamStepsResponse) error
}

type AuthServiceHandler[T context.Context] struct {
	Impl AuthService[T]
}

func (h *AuthServiceHandler[T]) Routes() map[string]goserver.UnaryMethodData[T] {
	return map[string]goserver.UnaryMethodData[T]{
		"protocol.auth.v1.AuthService/Federate": {
			MethodData: goserver.MethodData{Input: &FederateRequest{}, Output: &FederateResponse{}},
			Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
				return h.Impl.Federate(c, msg.(*FederateRequest))
			},
		},

		"protocol.auth.v1.AuthService/LoginFederated": {
			MethodData: goserver.MethodData{Input: &LoginFederatedRequest{}, Output: &LoginFederatedResponse{}},
			Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
				return h.Impl.LoginFederated(c, msg.(*LoginFederatedRequest))
			},
		},

		"protocol.auth.v1.AuthService/Key": {
			MethodData: goserver.MethodData{Input: &KeyRequest{}, Output: &KeyResponse{}},
			Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
				return h.Impl.Key(c, msg.(*KeyRequest))
			},
		},

		"protocol.auth.v1.AuthService/BeginAuth": {
			MethodData: goserver.MethodData{Input: &BeginAuthRequest{}, Output: &BeginAuthResponse{}},
			Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
				return h.Impl.BeginAuth(c, msg.(*BeginAuthRequest))
			},
		},

		"protocol.auth.v1.AuthService/NextStep": {
			MethodData: goserver.MethodData{Input: &NextStepRequest{}, Output: &NextStepResponse{}},
			Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
				return h.Impl.NextStep(c, msg.(*NextStepRequest))
			},
		},

		"protocol.auth.v1.AuthService/StepBack": {
			MethodData: goserver.MethodData{Input: &StepBackRequest{}, Output: &StepBackResponse{}},
			Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
				return h.Impl.StepBack(c, msg.(*StepBackRequest))
			},
		},

		"protocol.auth.v1.AuthService/CheckLoggedIn": {
			MethodData: goserver.MethodData{Input: &CheckLoggedInRequest{}, Output: &CheckLoggedInResponse{}},
			Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
				return h.Impl.CheckLoggedIn(c, msg.(*CheckLoggedInRequest))
			},
		},
	}
}

func (h *AuthServiceHandler[T]) StreamRoutes() map[string]goserver.StreamMethodData[T] {
	return map[string]goserver.StreamMethodData[T]{
		"protocol.auth.v1.AuthService/StreamSteps": {
			MethodData: goserver.MethodData{Input: &StreamStepsRequest{}, Output: &StreamStepsResponse{}},
			Handler: func(c T, msg chan goserver.VTProtoMessage, res chan goserver.VTProtoMessage) error {
				ch := make(chan *StreamStepsResponse)
				goserver.ToProtoChannel(ch, res)
				err := h.Impl.StreamSteps(c, goserver.FromProtoChannel[*StreamStepsRequest](msg), ch)
				return err
			},
		},
	}
}
