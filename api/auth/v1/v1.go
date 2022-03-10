package authv1impl

import (
	"fmt"

	"github.com/harmony-development/legato/api"
	"github.com/harmony-development/legato/api/auth"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db"
	"github.com/harmony-development/legato/gen"
	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	"github.com/matthewhartstonge/argon2"
	"github.com/samber/lo"
	"github.com/thanhpk/randstr"
)

type AuthV1 struct {
	db           *db.DB
	cfg          *config.Config
	formHandlers map[AuthStepType]formHandler
	hasher       argon2.Config
}

func New(db *db.DB, cfg *config.Config) *AuthV1 {
	inst := &AuthV1{
		db:     db,
		cfg:    cfg,
		hasher: argon2.DefaultConfig(),
	}

	inst.formHandlers = map[AuthStepType]formHandler{
		LoginStep:    inst.handleLogin,
		RegisterStep: inst.handleRegister,
	}

	return inst
}

func (v1 *AuthV1) generateID() string {
	return randstr.Hex(v1.cfg.Auth.TokenLength)
}

func (v1 *AuthV1) generateToken() string {
	return randstr.Hex(32)
}

func (v1 *AuthV1) Federate(c *api.LegatoContext, _ *authv1.FederateRequest) (*authv1.FederateResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *AuthV1) LoginFederated(_ *api.LegatoContext, _ *authv1.LoginFederatedRequest) (*authv1.LoginFederatedResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *AuthV1) Key(_ *api.LegatoContext, _ *authv1.KeyRequest) (*authv1.KeyResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (v1 *AuthV1) BeginAuth(c *api.LegatoContext, _ *authv1.BeginAuthRequest) (*authv1.BeginAuthResponse, error) {
	authId := v1.generateID()

	if err := v1.db.SetAuthStep(c, authId, int(InitialStep)); err != nil {
		return nil, fmt.Errorf("failed to set auth step: %w", err)
	}

	return &authv1.BeginAuthResponse{
		AuthId: authId,
	}, nil
}

func (v1 *AuthV1) NextStep(c *api.LegatoContext, req *authv1.NextStepRequest) (*authv1.NextStepResponse, error) {
	step, err := v1.db.GetAuthStep(c, req.AuthId)
	if err != nil {
		return nil, err
	}
	rawStep, ok := steps[AuthStepType(step)]
	if !ok {
		return nil, fmt.Errorf("step not found in map %w", api.ErrorInternalServerError)
	}
	if req.GetStep() == nil {
		return &authv1.NextStepResponse{
			Step: rawStep.ToProtoV1(),
		}, nil
	}
	switch currentStep := rawStep.(type) {
	case auth.ChoiceStep:
		choice := req.GetChoice()
		if choice == nil {
			return nil, api.ErrorStepMismatch
		}
		if !lo.Contains(currentStep.Choices, choice.Choice) {
			return nil, api.ErrorInvalidChoice
		}
		nextStep, err := fromString(choice.Choice)
		if err != nil {
			return nil, fmt.Errorf("choice %s not available: %w", choice.Choice, api.ErrorInternalServerError)
		}
		if err := v1.db.SetAuthStep(c, req.AuthId, int(nextStep)); err != nil {
			return nil, api.ErrorInternalServerError
		} else {
			return &authv1.NextStepResponse{
				Step: steps[nextStep].ToProtoV1(),
			}, nil
		}
	case auth.FormStep:
		f := req.GetForm()
		if f == nil {
			return nil, api.ErrorStepMismatch
		}
		handler, ok := v1.formHandlers[AuthStepType(step)]
		if !ok {
			return nil, fmt.Errorf("form handler not found for step %w", api.ErrorInternalServerError)
		}
		return handler(c, req.AuthId, f.Fields)
	default:
		return nil, api.ErrorInternalServerError
	}
}

func (v1 *AuthV1) StepBack(c *api.LegatoContext, req *authv1.StepBackRequest) (*authv1.StepBackResponse, error) {
	step, err := v1.db.GetAuthStep(c, req.AuthId)
	if err != nil {
		return nil, err
	}
	rawStep, ok := steps[AuthStepType(step)]
	if !ok {
		return nil, fmt.Errorf("step not found in map %w", api.ErrorInternalServerError)
	}
	if !rawStep.StepInfo().CanGoBack {
		return nil, api.ErrorNoStepBack
	}

	prevStep, err := fromString(rawStep.StepInfo().PreviousStep)
	if err != nil {
		return nil, fmt.Errorf("step %s not available: %w", rawStep.StepInfo().PreviousStep, api.ErrorInternalServerError)
	}

	if err := v1.db.SetAuthStep(c, req.AuthId, int(prevStep)); err != nil {
		return nil, fmt.Errorf("failed to go back from %s to %s: %w", rawStep.StepInfo().ID, rawStep.StepInfo().PreviousStep, api.ErrorInternalServerError)
	}

	return &authv1.StepBackResponse{
		Step: steps[prevStep].ToProtoV1(),
	}, nil
}

func (v1 *AuthV1) CheckLoggedIn(_ *api.LegatoContext, _ *authv1.CheckLoggedInRequest) (*authv1.CheckLoggedInResponse, error) {
	return nil, nil
}

func (v1 *AuthV1) StreamSteps(ctx *api.LegatoContext, stream chan *authv1.StreamStepsRequest, res chan *authv1.StreamStepsResponse) error {
	req := <-stream
	step := v1.db.StreamUserSteps(ctx, req.AuthId)
	for s := range step {
		switch msg := s.GetMessage().(type) {
		case *gen.AuthMessage_StepId:
			res <- &authv1.StreamStepsResponse{
				Step: steps[AuthStepType(msg.StepId)].ToProtoV1(),
			}
		case *gen.AuthMessage_Session_:
			res <- &authv1.StreamStepsResponse{
				Step: &authv1.AuthStep{
					CanGoBack:   false,
					FallbackUrl: "",
					Step: &authv1.AuthStep_Session{
						Session: &authv1.Session{
							SessionToken: msg.Session.SessionId,
							UserId:       msg.Session.UserId,
						},
					},
				},
			}
		}
	}
	return nil
}
