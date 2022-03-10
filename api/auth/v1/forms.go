package authv1impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/harmony-development/legato/api"
	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/matthewhartstonge/argon2"
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

	id, err := v1.db.CreateUser(ctx, email, username, passwordHash)
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

	return v1.genToken(ctx, sessionID, id)
}

func (v1 *AuthV1) handleLogin(ctx context.Context, sessionID string, fields []*authv1.NextStepRequest_FormFields) (*authv1.NextStepResponse, error) {
	if len(fields) < 2 {
		return nil, api.ErrorInvalidForm
	}
	email := fields[0].GetString_()
	password := fields[1].GetBytes()
	if email == "" || password == nil {
		return nil, api.ErrorMissingForm
	}

	user, err := v1.db.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, api.ErrorInvalidAuth
		}
		return nil, err
	}

	if ok, err := argon2.VerifyEncoded(password, user.PasswordHash); !ok {
		return nil, api.ErrorInvalidAuth
	} else if err != nil {
		return nil, fmt.Errorf("error while verifying password: %w", err)
	}

	return v1.genToken(ctx, sessionID, user.ID)
}

func (v1 *AuthV1) genToken(ctx context.Context, sessionID string, userID uint64) (*authv1.NextStepResponse, error) {
	token := v1.generateToken()

	if err := v1.db.SaveSession(ctx, sessionID, token, userID); err != nil {
		return nil, err
	}

	return &authv1.NextStepResponse{
		Step: &authv1.AuthStep{
			Step: &authv1.AuthStep_Session{
				Session: &authv1.Session{
					UserId:       userID,
					SessionToken: token,
				},
			},
		},
	}, nil
}
