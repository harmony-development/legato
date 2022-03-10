package db

import (
	"context"
	"fmt"
	"time"

	"github.com/harmony-development/legato/db/models"
	"github.com/harmony-development/legato/gen"
)

func (db *DB) SaveSession(ctx context.Context, sessionID string, token string, userID uint64) error {
	if err := db.PublishAuth(ctx, sessionID, &gen.AuthMessage{
		Message: &gen.AuthMessage_Session_{
			Session: &gen.AuthMessage_Session{
				SessionId: token,
				UserId:    userID,
			},
		},
	}); err != nil {
		return Wrap(err, "failed to publish session")
	}

	return TryWrap(db.Rdb.Set(ctx, subkey(sessionsPrefix, token), userID, 0).Err(), "failed to save session")
}

func (db *DB) GetUserForSession(ctx context.Context, sessionID string) (uint64, error) {
	res, err := db.Rdb.Get(ctx, subkey(sessionsPrefix, sessionID)).Uint64()

	return res, TryWrap(err, "failed to get user for session")
}

func (db *DB) PublishAuth(ctx context.Context, session string, msg *gen.AuthMessage) error {
	msgBytes, err := msg.MarshalVT()
	if err != nil {
		return Wrap(err, "failed to marshal auth publish message")
	}

	return TryWrap(
		db.Rdb.Publish(ctx, subkey(authStepPrefix, session), string(msgBytes)).Err(),
		"failed to publish session",
	)
}

// SetAuthStep saves to redis the auth session and step the user is on.
func (db *DB) SetAuthStep(ctx context.Context, sessionID string, step int) error {
	if err := db.PublishAuth(ctx, sessionID, &gen.AuthMessage{
		Message: &gen.AuthMessage_StepId{
			StepId: int32(step),
		},
	}); err != nil {
		return err
	}

	return TryWrap(
		db.Rdb.Set(
			ctx,
			subkey(authStepPrefix, sessionID),
			step,
			6*time.Hour,
		).Err(),
		"failed to save auth session",
	)
}

func (db *DB) GetAuthStep(ctx context.Context, sessionID string) (int, error) {
	res, err := db.Rdb.Get(ctx, subkey(authStepPrefix, sessionID)).Int()
	if err != nil {
		return 0, fmt.Errorf("failed to get auth step for %s: %w", sessionID, err)
	}

	return res, nil
}

func (db *DB) StreamUserSteps(ctx context.Context, sessionID string) chan *gen.AuthMessage {
	ch := make(chan *gen.AuthMessage)

	go func() {
		sub := db.Rdb.Subscribe(ctx, subkey(authStepPrefix, sessionID)).Channel()
		for msg := range sub {
			res := &gen.AuthMessage{}

			if err := res.UnmarshalVT([]byte(msg.Payload)); err != nil {
				continue
			}

			ch <- res
		}

		close(ch)
	}()

	return ch
}

// CreateUser saves a user to the database.
func (db *DB) CreateUser(ctx context.Context, email string, username string, passwordHash []byte) (uint64, error) {
	userID, err := db.models.CreateUser(ctx, models.CreateUserParams{
		Email:        &email,
		Username:     username,
		PasswordHash: passwordHash,
	})

	return userID, TryWrap(err, "failed to create user")
}

func (db *DB) GetUserByID(ctx context.Context, id uint64) (models.User, error) {
	user, err := db.models.GetUserById(ctx, id)
	return user, TryWrap(err, "failed to get user by id")
}

func (db *DB) GetUserByEmail(ctx context.Context, email string) (models.GetUserByEmailRow, error) {
	user, err := db.models.GetUserByEmail(ctx, &email)
	return user, TryWrap(err, "failed to get user by email")
}
