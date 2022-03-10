package db

import (
	"context"
	"strconv"
	"strings"

	goserver "github.com/harmony-development/hrpc/pkg/go-server"
)

const (
	sessionsPrefix = "sessions"

	authStepPrefix = "auth_state"
	chatPrefix     = "chat"
	profilePrefix  = "profile"
	subPrefix      = "subs"
)

func subkey(key string, subkeys ...string) string {
	return key + ":" + strings.Join(subkeys, ":")
}

func idsubkey(key string, id uint64) string {
	return subkey(key, strconv.FormatUint(id, 10))
}

func subscribeKey(userID uint64) string {
	return idsubkey(subPrefix, userID)
}

func guildKey(guildID uint64) string {
	return idsubkey(chatPrefix, guildID)
}

func profileKey(userID uint64) string {
	return idsubkey(profilePrefix, userID)
}

func (db *DB) publishStreamEvent(ctx context.Context, topic string, event goserver.VTProtoMessage) error {
	serialized, err := event.MarshalVT()
	if err != nil {
		return Wrap(err, "failed to serialize stream event")
	}

	return TryWrap(db.Rdb.Publish(ctx, topic, serialized).Err(), "failed to publish stream event")
}
