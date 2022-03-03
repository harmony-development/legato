package authv1impl

import (
	"github.com/harmony-development/legato/api"
	authv1 "github.com/harmony-development/legato/gen/auth/v1"
)

type AuthV1 struct{}

func (v1 *AuthV1) Federate(c api.LegatoContext, _ *authv1.FederateRequest) (*authv1.FederateResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *AuthV1) LoginFederated(_ api.LegatoContext, _ *authv1.LoginFederatedRequest) (*authv1.LoginFederatedResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *AuthV1) Key(_ api.LegatoContext, _ *authv1.KeyRequest) (*authv1.KeyResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *AuthV1) BeginAuth(_ api.LegatoContext, _ *authv1.BeginAuthRequest) (*authv1.BeginAuthResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *AuthV1) NextStep(_ api.LegatoContext, _ *authv1.NextStepRequest) (*authv1.NextStepResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *AuthV1) StepBack(_ api.LegatoContext, _ *authv1.StepBackRequest) (*authv1.StepBackResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *AuthV1) CheckLoggedIn(_ api.LegatoContext, _ *authv1.CheckLoggedInRequest) (*authv1.CheckLoggedInResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *AuthV1) StreamSteps(_ api.LegatoContext, _ chan *authv1.StreamStepsRequest) (chan *authv1.StreamStepsResponse, error) {
	panic("not implemented") // TODO: Implement
}
