package authv1impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/harmony-development/legato/api"
	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	"github.com/jackc/pgconn"
)

type formHandler func(ctx context.Context, sessionID string, fields []*authv1.NextStepRequest_FormFields) (*authv1.NextStepResponse, error)

func (v1 *AuthV1) handleRegister(ctx context.Context, sessionID string, fields []*authv1.NextStepRequest_FormFields) (*authv1.NextStepResponse, error) {
	if len(fields) < 3 {
		return nil, api.ErrorInvalidForm
	}
	email := fields[0].GetString_()
	username := fields[1].GetString_()
	password := fields[2].GetBytes()
	if email == "" || username == "" || password == nil {
		return nil, api.ErrorMissingForm
	}

	rawHash, err := v1.hasher.Hash(password, nil)
	if err != nil {
		return nil, err
	}

	passwordHash := rawHash.Encode()

	id, err := v1.db.SaveUser(ctx, email, username, passwordHash)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Message {
			case "duplicate key value violates unique constraint \"users_email_key\"":
				return nil, api.ErrorEmailAlreadyUsed
			case "duplicate key value violates unique constraint \"users_username_key\"":
				return nil, api.ErrorUsernameAlreadyUsed
			}
		}
		return nil, fmt.Errorf("failed to save user on register: %w", err)
	}

	token := v1.generateToken()

	v1.db.SaveSession(ctx, sessionID, token, id)

	return &authv1.NextStepResponse{
		Step: &authv1.AuthStep{
			Step: &authv1.AuthStep_Session{
				Session: &authv1.Session{
					UserId:       id,
					SessionToken: token,
				},
			},
		},
	}, nil
}

// func (v1 *AuthV1) handleLogin(ctx context.Context, fields []*authv1.NextStepRequest_FormFields) (*authv1.NextStepResponse, error) {
// }
