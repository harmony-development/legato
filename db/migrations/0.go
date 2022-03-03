package migrations

import (
	"context"

	"github.com/harmony-development/legato/db"
)

func Migration0(db *db.DB) error {
	return db.RawExec(context.Background(), Migration0Query)
}

var Migration0Query = `
create table users (
	id bigint not null primary key,
	username text unique not null,
	avatar text,
	user_status smallint not null default 0,
	host text,
	remote_id bigint,
	email text unique,
	password_hash bytea,
	created timestamp not null default (current_timestamp at time zone 'utc'),
	check (
			(
					host is not null and remote_id is not null and 
					email is null and password_hash is null
			) or (
					host is null and remote_id is null and
					email is not null and password_hash is not null
			)
	)
);
create table guilds (
	id bigint not null primary key,
	name text not null,
	picture text,
	type smallint not null,
	banned_users bigint[] not null default '{}',
	created timestamp not null default (current_timestamp at time zone 'utc')
);
create table guild_members (
	user_id bigint not null references users(id) on delete cascade,
	guild_id bigint not null references guilds(id) on delete cascade,
	owns_guild boolean not null default false,
	joined timestamp not null default (current_timestamp at time zone 'utc'),
	primary key (user_id, guild_id)
);
create table guild_list (
	user_id bigint not null references users(id) on delete cascade,
	guild_id bigint not null,
	host text not null default '',
	position text not null,
	primary key (user_id, guild_id, host)
);
create table if not exists channels (
	id bigint primary key unique not null,
	created_at timestamp not null default (current_timestamp at time zone 'utc'),
	guild_id bigint references guilds (id) on delete cascade,
	name text not null,
	kind smallint not null,
	position text not null
);
create type message_override as (
	username text,
	avatar text,
	reason text
);
create table if not exists messages (
	id bigint primary key,
	channel_id bigint references channels (id) on delete cascade,
	created_at timestamp not null default (current_timestamp at time zone 'utc'),
	reply_to_id bigint,
	author_id bigint not null,
	edited_at timestamp,
	content text not null,
	-- todo: add formatting
	-- todo: add embeds
	-- todo: add actions
	override message_override,
	attachments text []
);
create table meta (
	current_migration int not null
);
insert into meta values (0);
CREATE SEQUENCE public.global_id_seq;
CREATE OR REPLACE FUNCTION public.generate_id()
	RETURNS bigint
	LANGUAGE 'plpgsql'
AS $BODY$
DECLARE
	our_epoch bigint := 1314220021721;
	seq_id bigint;
	now_millis bigint;
	-- the id of this DB shard, must be set for each
	-- schema shard you have - you could pass this as a parameter too
	shard_id int := 1;
	result bigint:= 0;
BEGIN
	SELECT nextval('public.global_id_seq') % 1024 INTO seq_id;
	SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp()) * 1000) INTO now_millis;
	result := (now_millis - our_epoch) << 23;
	result := result | (shard_id << 10);
	result := result | (seq_id);
return result;
END;
$BODY$;
`
