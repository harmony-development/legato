package db

import (
	"context"
	"fmt"
	"strconv"

	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	"github.com/samber/lo"
)

func (db *DB) SubscribeStream(ctx context.Context, userID uint64, topics ...string) error {
	sub, ok := db.streams.Load(userID)
	if ok {
		return sub.Subscribe(ctx, topics...)
	}
	return nil // TODO: return error if not found...?
}

func (db *DB) PublishChatEvent(ctx context.Context, guildID uint64, event *chatv1.StreamEvent) error {
	return db.publishStreamEvent(ctx, subkey(chatPrefix, strconv.FormatUint(guildID, 10)), event)
}

func (db *DB) PublishUserEvent(ctx context.Context, userID uint64, event *chatv1.StreamEvent) error {
	return db.publishStreamEvent(ctx, subkey(profilePrefix, strconv.FormatUint(userID, 10)), event)
}

func (db *DB) StreamChatEvents(ctx context.Context, userID uint64, guilds []uint64) chan *chatv1.StreamEvent {
	ch := make(chan *chatv1.StreamEvent)

	go func() {
		topics := append(
			lo.Map(guilds, func(guildID uint64, _ int) string {
				return guildKey(guildID)
			}),
			profileKey(userID),
		)

		sub := db.Rdb.Subscribe(
			ctx,
			topics...,
		)

		db.streams.Store(userID, sub)

		defer func() {
			if err := sub.Close(); err != nil {
				fmt.Println("failed to close auth subscription", err)
			}
			close(ch)
		}()

		for msg := range sub.Channel() {
			res := &chatv1.StreamEvent{}

			if err := res.UnmarshalVT([]byte(msg.Payload)); err != nil {
				continue
			}

			ch <- res
		}
	}()

	return ch
}
