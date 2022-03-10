package db

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	"github.com/samber/lo"
)

// SubscribeStream publishes a new command to subscribe a user to new topics.
func (db *DB) SubscribeStream(ctx context.Context, userID uint64, topics ...string) error {
	return TryWrap(db.Rdb.Publish(ctx, subscribeKey(userID), strings.Join(topics, ",")).Err(), "failed to publish subscribe command")
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

		defer func() {
			if err := sub.Close(); err != nil {
				fmt.Println("failed to close auth subscription", err)
			}
			close(ch)
		}()

		l := sync.RWMutex{}

		go func() {
			handler := db.Rdb.Subscribe(ctx, subscribeKey(userID))
			defer func() {
				if err := handler.Close(); err != nil {
					fmt.Println("failed to close handler: %w", err)
				}
			}()

			for msg := range handler.Channel() {
				l.Lock()
				sub.Subscribe(ctx, strings.Split(msg.Payload, ",")...)
				l.Unlock()
			}
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
