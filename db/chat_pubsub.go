package db

import (
	"context"
	"fmt"
	"strconv"

	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	"github.com/samber/lo"
)

func (db *DB) PublishChatEvent(ctx context.Context, guildID uint64, event *chatv1.StreamEvent) error {
	return db.publishStreamEvent(ctx, subkey(chatPrefix, strconv.FormatUint(guildID, 10)), event)
}

func (db *DB) StreamChatEvents(ctx context.Context, userID uint64, guilds []uint64) chan *chatv1.StreamEvent {
	ch := make(chan *chatv1.StreamEvent)

	go func() {
		sub := db.Rdb.Subscribe(
			ctx,
			append(
				lo.Map(guilds, func(guildID uint64, _ int) string {
					return subkey(chatPrefix, strconv.FormatUint(guildID, 10))
				}),
				subkey(profilePrefix, strconv.FormatUint(userID, 10)),
			)...,
		)

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
