package db

import (
	"context"
	"strconv"
	"strings"

	goserver "github.com/harmony-development/hrpc/pkg/go-server"
)

const (
	authStepPrefix = "auth_state"
	sessionsPrefix = "sessions"
	chatPrefix     = "chat"
	profilePrefix  = "profile"
)

func subkey(key string, subkeys ...string) string {
	return key + ":" + strings.Join(subkeys, ":")
}

func guildKey(guildID uint64) string {
	return subkey(chatPrefix, strconv.FormatUint(guildID, 10))
}

func profileKey(userID uint64) string {
	return subkey(chatPrefix, strconv.FormatUint(userID, 10))
}

func (db *DB) publishStreamEvent(ctx context.Context, topic string, event goserver.VTProtoMessage) error {
	serialized, err := event.MarshalVT()
	if err != nil {
		return Wrap(err, "failed to serialize stream event")
	}

	return TryWrap(db.Rdb.Publish(ctx, topic, serialized).Err(), "failed to publish stream event")
}
