-- name: CreateGuild :one
insert into guilds (id, name, picture, type) values (generate_id(), $1, $2, $3) returning *;

-- name: AddGuildMember :exec
insert into guild_members (user_id, guild_id, owns_guild) values ($1, $2, $3);

-- name: AddGuildToList :exec
insert into guild_list (user_id, guild_id, host, position) values ($1, $2, $3, $4);

-- name: GetGuildList :many
select guild_id, host, position from guild_list where user_id = $1 order by position asc;

-- name: GetGuildsById :many
select guilds.*, array_agg(guild_members.user_id)::bigint[] as owner_ids
	from guilds left join guild_members on guild_members.guild_id = guilds.id
	where guilds.id = any($1::bigint[]) and guild_members.owns_guild group by guilds.id;

-- name: GetTopGuild :one
select position from guild_list where user_id = $1 order by position asc limit 1;

-- name: GetTopChannel :one
select id, position from channels where guild_id = $1 order by position asc limit 1;

-- name: CreateChannel :one
insert into channels (id, guild_id, name, kind, position) values (generate_id(), $1, $2, $3, $4) returning *;

-- name: GetChannel :one
select * from channels where id = $1;

-- name: GetChannelList :many
select * from channels where guild_id = $1 order by position desc;

-- name: GetGuildMembers :many
select user_id, owns_guild from guild_members where guild_id = $1;

-- name: IsGuildMember :one
select true as exists from guild_members where user_id = $1 and guild_id = $2;

-- name: GetGuildMember :one
select * from guild_members where user_id = $1 and guild_id = $2;

-- name: HasSharedGuilds :one
select 
	(select array_agg(guild_id) from guild_members where guild_members.user_id = $1) 
	&& (select array_agg(guild_id) from guild_members where guild_members.user_id = $2) 
as exists;

-- name: RemoveGuildMember :exec
delete from guild_members where user_id = $1 and guild_id = $2;

-- name: RemoveGuildFromList :exec
delete from guild_list where user_id = $1 and guild_id = $2;
