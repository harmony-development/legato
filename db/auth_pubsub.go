package db

import (
	"context"

	"github.com/harmony-development/legato/gen"
)

func (db *DB) PublishAuth(ctx context.Context, sessionID string, event *gen.AuthMessage) error {
	return db.publishStreamEvent(ctx, subkey(authStepPrefix, sessionID), event)
}
