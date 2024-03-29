// Code generated by sqlc. DO NOT EDIT.
// source: guilds.sql

package models

import (
	"context"
	"time"
)

const addGuildMember = `-- name: AddGuildMember :exec
insert into guild_members (user_id, guild_id, owns_guild) values ($1, $2, $3)
`

type AddGuildMemberParams struct {
	UserID    uint64
	GuildID   uint64
	OwnsGuild bool
}

func (q *Queries) AddGuildMember(ctx context.Context, arg AddGuildMemberParams) error {
	_, err := q.db.Exec(ctx, addGuildMember, arg.UserID, arg.GuildID, arg.OwnsGuild)
	return err
}

const addGuildToList = `-- name: AddGuildToList :exec
insert into guild_list (user_id, guild_id, host, position) values ($1, $2, $3, $4)
`

type AddGuildToListParams struct {
	UserID   uint64
	GuildID  uint64
	Host     string
	Position string
}

func (q *Queries) AddGuildToList(ctx context.Context, arg AddGuildToListParams) error {
	_, err := q.db.Exec(ctx, addGuildToList,
		arg.UserID,
		arg.GuildID,
		arg.Host,
		arg.Position,
	)
	return err
}

const createChannel = `-- name: CreateChannel :one
insert into channels (id, guild_id, name, kind, position) values (generate_id(), $1, $2, $3, $4) returning id, created_at, guild_id, name, kind, position
`

type CreateChannelParams struct {
	GuildID  uint64
	Name     string
	Kind     int16
	Position string
}

func (q *Queries) CreateChannel(ctx context.Context, arg CreateChannelParams) (Channel, error) {
	row := q.db.QueryRow(ctx, createChannel,
		arg.GuildID,
		arg.Name,
		arg.Kind,
		arg.Position,
	)
	var i Channel
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.GuildID,
		&i.Name,
		&i.Kind,
		&i.Position,
	)
	return i, err
}

const createGuild = `-- name: CreateGuild :one
insert into guilds (id, name, picture, type) values (generate_id(), $1, $2, $3) returning id, name, picture, type, banned_users, created
`

type CreateGuildParams struct {
	Name    string
	Picture *string
	Type    int16
}

func (q *Queries) CreateGuild(ctx context.Context, arg CreateGuildParams) (Guild, error) {
	row := q.db.QueryRow(ctx, createGuild, arg.Name, arg.Picture, arg.Type)
	var i Guild
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Picture,
		&i.Type,
		&i.BannedUsers,
		&i.Created,
	)
	return i, err
}

const getChannel = `-- name: GetChannel :one
select id, created_at, guild_id, name, kind, position from channels where id = $1
`

func (q *Queries) GetChannel(ctx context.Context, id uint64) (Channel, error) {
	row := q.db.QueryRow(ctx, getChannel, id)
	var i Channel
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.GuildID,
		&i.Name,
		&i.Kind,
		&i.Position,
	)
	return i, err
}

const getChannelList = `-- name: GetChannelList :many
select id, created_at, guild_id, name, kind, position from channels where guild_id = $1 order by position desc
`

func (q *Queries) GetChannelList(ctx context.Context, guildID uint64) ([]Channel, error) {
	rows, err := q.db.Query(ctx, getChannelList, guildID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Channel
	for rows.Next() {
		var i Channel
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.GuildID,
			&i.Name,
			&i.Kind,
			&i.Position,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGuildList = `-- name: GetGuildList :many
select guild_id, host, position from guild_list where user_id = $1 order by position asc
`

type GetGuildListRow struct {
	GuildID  uint64
	Host     string
	Position string
}

func (q *Queries) GetGuildList(ctx context.Context, userID uint64) ([]GetGuildListRow, error) {
	rows, err := q.db.Query(ctx, getGuildList, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetGuildListRow
	for rows.Next() {
		var i GetGuildListRow
		if err := rows.Scan(&i.GuildID, &i.Host, &i.Position); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGuildMember = `-- name: GetGuildMember :one
select user_id, guild_id, owns_guild, joined from guild_members where user_id = $1 and guild_id = $2
`

type GetGuildMemberParams struct {
	UserID  uint64
	GuildID uint64
}

func (q *Queries) GetGuildMember(ctx context.Context, arg GetGuildMemberParams) (GuildMember, error) {
	row := q.db.QueryRow(ctx, getGuildMember, arg.UserID, arg.GuildID)
	var i GuildMember
	err := row.Scan(
		&i.UserID,
		&i.GuildID,
		&i.OwnsGuild,
		&i.Joined,
	)
	return i, err
}

const getGuildMembers = `-- name: GetGuildMembers :many
select user_id, owns_guild from guild_members where guild_id = $1
`

type GetGuildMembersRow struct {
	UserID    uint64
	OwnsGuild bool
}

func (q *Queries) GetGuildMembers(ctx context.Context, guildID uint64) ([]GetGuildMembersRow, error) {
	rows, err := q.db.Query(ctx, getGuildMembers, guildID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetGuildMembersRow
	for rows.Next() {
		var i GetGuildMembersRow
		if err := rows.Scan(&i.UserID, &i.OwnsGuild); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGuildsById = `-- name: GetGuildsById :many
select guilds.id, guilds.name, guilds.picture, guilds.type, guilds.banned_users, guilds.created, array_agg(guild_members.user_id)::bigint[] as owner_ids
	from guilds left join guild_members on guild_members.guild_id = guilds.id
	where guilds.id = any($1::bigint[]) and guild_members.owns_guild group by guilds.id
`

type GetGuildsByIdRow struct {
	ID          uint64
	Name        string
	Picture     *string
	Type        int16
	BannedUsers []uint64
	Created     time.Time
	OwnerIds    []uint64
}

func (q *Queries) GetGuildsById(ctx context.Context, dollar_1 []uint64) ([]GetGuildsByIdRow, error) {
	rows, err := q.db.Query(ctx, getGuildsById, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetGuildsByIdRow
	for rows.Next() {
		var i GetGuildsByIdRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Picture,
			&i.Type,
			&i.BannedUsers,
			&i.Created,
			&i.OwnerIds,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTopChannel = `-- name: GetTopChannel :one
select id, position from channels where guild_id = $1 order by position asc limit 1
`

type GetTopChannelRow struct {
	ID       uint64
	Position string
}

func (q *Queries) GetTopChannel(ctx context.Context, guildID uint64) (GetTopChannelRow, error) {
	row := q.db.QueryRow(ctx, getTopChannel, guildID)
	var i GetTopChannelRow
	err := row.Scan(&i.ID, &i.Position)
	return i, err
}

const getTopGuild = `-- name: GetTopGuild :one
select position from guild_list where user_id = $1 order by position asc limit 1
`

func (q *Queries) GetTopGuild(ctx context.Context, userID uint64) (string, error) {
	row := q.db.QueryRow(ctx, getTopGuild, userID)
	var position string
	err := row.Scan(&position)
	return position, err
}

const hasSharedGuilds = `-- name: HasSharedGuilds :one
select 
	(select array_agg(guild_id) from guild_members where guild_members.user_id = $1) 
	&& (select array_agg(guild_id) from guild_members where guild_members.user_id = $2) 
as exists
`

type HasSharedGuildsParams struct {
	UserID   uint64
	UserID_2 uint64
}

func (q *Queries) HasSharedGuilds(ctx context.Context, arg HasSharedGuildsParams) (interface{}, error) {
	row := q.db.QueryRow(ctx, hasSharedGuilds, arg.UserID, arg.UserID_2)
	var exists interface{}
	err := row.Scan(&exists)
	return exists, err
}

const isGuildMember = `-- name: IsGuildMember :one
select true as exists from guild_members where user_id = $1 and guild_id = $2
`

type IsGuildMemberParams struct {
	UserID  uint64
	GuildID uint64
}

func (q *Queries) IsGuildMember(ctx context.Context, arg IsGuildMemberParams) (bool, error) {
	row := q.db.QueryRow(ctx, isGuildMember, arg.UserID, arg.GuildID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const removeGuildFromList = `-- name: RemoveGuildFromList :exec
delete from guild_list where user_id = $1 and guild_id = $2
`

type RemoveGuildFromListParams struct {
	UserID  uint64
	GuildID uint64
}

func (q *Queries) RemoveGuildFromList(ctx context.Context, arg RemoveGuildFromListParams) error {
	_, err := q.db.Exec(ctx, removeGuildFromList, arg.UserID, arg.GuildID)
	return err
}

const removeGuildMember = `-- name: RemoveGuildMember :exec
delete from guild_members where user_id = $1 and guild_id = $2
`

type RemoveGuildMemberParams struct {
	UserID  uint64
	GuildID uint64
}

func (q *Queries) RemoveGuildMember(ctx context.Context, arg RemoveGuildMemberParams) error {
	_, err := q.db.Exec(ctx, removeGuildMember, arg.UserID, arg.GuildID)
	return err
}
