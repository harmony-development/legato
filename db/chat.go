package db

import (
	"context"
	"fmt"
	"strconv"

	"github.com/harmony-development/legato/db/models"
	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	"github.com/jackc/pgx/v4"
	"github.com/samber/lo"
	"github.com/xissy/lexorank"
)

type GuildKind int16

const (
	GuildKindGuild = iota
	GuildKindRoom
	GuildKindDM
)

func (db *DB) CreateGuild(ctx context.Context, name string, picture *string, ownerID uint64, type_ GuildKind) (models.Guild, error) {
	var guild models.Guild
	err := db.Postgres.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		m := db.models.WithTx(tx)

		pos, err := m.GetTopGuild(ctx, ownerID)
		if err != nil {
			pos = ""
		}

		g, err := m.CreateGuild(ctx, models.CreateGuildParams{
			Name:    name,
			Picture: picture,
			Type:    int16(type_),
		})
		if err != nil {
			return Wrap(err, "failed to create guild")
		}

		guild = g

		if err := m.AddGuildMember(ctx, models.AddGuildMemberParams{
			GuildID:   guild.ID,
			UserID:    ownerID,
			OwnsGuild: true,
		}); err != nil {
			return Wrap(err, "failed to add guild owner")
		}

		newPos, _ := lexorank.Rank("", pos)

		return TryWrap(m.AddGuildToList(ctx, models.AddGuildToListParams{
			UserID:   ownerID,
			GuildID:  guild.ID,
			Host:     "",
			Position: newPos,
		}), "failed to add guild to list")
	})

	return guild, err
}

func (db *DB) GetGuildList(ctx context.Context, userID uint64) ([]models.GetGuildListRow, error) {
	guilds, err := db.models.GetGuildList(ctx, userID)
	return guilds, TryWrap(err, "failed to get guild list for %d", userID)
}

func (db *DB) GetGuildsById(ctx context.Context, guildIds []uint64) ([]models.Guild, error) {
	guilds, err := db.models.GetGuildsById(ctx, guildIds)
	if err != nil {
		return nil, Wrap(err, "failed to get guilds by ids %v", guildIds)
	}

	return lo.Map(guilds, func(row models.GetGuildsByIdRow, _ int) models.Guild {
		return models.Guild{
			ID:          row.ID,
			Name:        row.Name,
			Picture:     row.Picture,
			Type:        row.Type,
			BannedUsers: row.BannedUsers,
			Created:     row.Created,
		}
	}), nil
}

func (db *DB) CreateChannel(ctx context.Context) {
}

func (db *DB) GetChannel(ctx context.Context) {
}

func (db *DB) GetChannelList(ctx context.Context, guildID uint64) ([]models.Channel, error) {
	channels, err := db.models.GetChannelList(ctx, guildID)
	return channels, TryWrap(err, "failed to get channel list for %d", guildID)
}

func (db *DB) GetGuildMembers(ctx context.Context) {
}

func (db *DB) IsGuildMember(ctx context.Context) {
}

func (db *DB) GetGuildMember(ctx context.Context) {
}

func (db *DB) HasSharedGuilds(ctx context.Context) {
}

func (db *DB) RemoveGuildMember(ctx context.Context) {
}

func (db *DB) RemoveGuildFromList(ctx context.Context) {
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
				subkey(chatPrefix, strconv.FormatUint(userID, 10)),
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
