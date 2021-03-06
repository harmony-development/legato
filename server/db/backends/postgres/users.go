package postgres

import (
	"database/sql"
	"errors"
	"math"
	"time"

	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/db/queries"
	"github.com/harmony-development/legato/server/db/types"
	"github.com/harmony-development/legato/server/db/utilities"
	"github.com/ztrue/tracerr"
)

// GetLocalGuilds gets the guilds a user is in
func (db *database) GetLocalGuilds(userID uint64) ([]uint64, error) {
	data, err := db.queries.GuildsForUser(ctx, userID)
	err = tracerr.Wrap(err)
	return data, err
}

// SessionToUserID gets the user ID from a session
func (db *database) SessionToUserID(session string) (uint64, error) {
	userID, exists := db.SessionCache.Get(session)
	s, ok := userID.(uint64)
	if !exists || !ok {
		userID, err := db.queries.SessionToUserID(ctx, session)
		if err != nil {
			err = tracerr.Wrap(err)
			db.Logger.CheckException(err)
		}
		return userID, err
	}
	return s, nil
}

// GetUser gets a user with their email
func (db *database) GetUserByEmail(email string) (queries.GetUserByEmailRow, error) {
	ret, err := db.queries.GetUserByEmail(ctx, email)
	err = tracerr.Wrap(err)
	return ret, err
}

// GetUserByID gets a user with their ID and their home server
func (db *database) GetUserByID(userID uint64) (queries.GetUserRow, error) {
	ret, err := db.queries.GetUser(ctx, userID)
	err = tracerr.Wrap(err)
	return ret, err
}

// AddSession persists a session into the DB
func (db *database) AddSession(userID uint64, session string) error {
	db.SessionCache.Add(session, userID)
	return tracerr.Wrap(db.queries.AddSession(ctx, queries.AddSessionParams{
		UserID:     userID,
		Session:    session,
		Expiration: time.Now().UTC().Add(db.Config.Server.Policies.Sessions.Duration).Unix(),
	}))
}

// GetLocalUserForForeignUser gets a local user from the foreign users database
func (db *database) GetLocalUserForForeignUser(userID uint64, homeserver string) (uint64, error) {
	ret, err := db.queries.GetLocalUserID(ctx, queries.GetLocalUserIDParams{
		UserID:     userID,
		HomeServer: homeserver,
	})
	err = tracerr.Wrap(err)
	return ret, err
}

// AddLocalUser adds a user to the DB that contains login information (not foreign)
func (db *database) AddLocalUser(userID uint64, email, username string, passwordHash []byte) error {
	tx, err := db.Begin()
	if err != nil {
		err = tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return err
	}
	tq := db.queries.WithTx(tx)
	if err := tq.AddUser(ctx, userID); err != nil {
		err = tracerr.Wrap(err)
		return err
	}
	if err := tq.AddLocalUser(ctx, queries.AddLocalUserParams{
		UserID:   userID,
		Email:    email,
		Password: passwordHash,
	}); err != nil {
		err = tracerr.Wrap(err)
		return err
	}
	if err := tq.AddProfile(ctx, queries.AddProfileParams{
		UserID:   userID,
		Username: username,
		Avatar:   sql.NullString{},
		Status:   int16(harmonytypesv1.UserStatus_USER_STATUS_OFFLINE),
		IsBot:    false,
	}); err != nil {
		err = tracerr.Wrap(err)
		return err
	}
	if err := tx.Commit(); err != nil {
		return tracerr.Wrap(tx.Rollback())
	}
	return nil
}

// AddForeignUser inserts
func (db *database) AddForeignUser(homeServer string, userID, localUserID uint64, username, avatar string) (uint64, error) {
	tx, err := db.Begin()
	if err != nil {
		err = tracerr.Wrap(err)
		db.Logger.CheckException(err)
		return 0, err
	}
	tq := db.queries.WithTx(tx)
	if err := tq.AddUser(ctx, localUserID); err != nil {
		err = tracerr.Wrap(err)
		return 0, err
	}
	if err := tq.AddProfile(ctx, queries.AddProfileParams{
		UserID:   localUserID,
		Username: username,
		Avatar:   utilities.ToSqlString(avatar),
		Status:   int16(harmonytypesv1.UserStatus_USER_STATUS_OFFLINE),
		IsBot:    false,
	}); err != nil {
		err = tracerr.Wrap(err)
		return 0, err
	}
	if userID, err = tq.AddForeignUser(ctx, queries.AddForeignUserParams{
		UserID:      userID,
		HomeServer:  homeServer,
		LocalUserID: localUserID,
	}); err != nil {
		err = tracerr.Wrap(err)
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			err = tracerr.Wrap(err)
			return 0, err
		}
		err = tracerr.Wrap(err)
		return 0, err
	}
	return userID, nil
}

func (db *database) EmailExists(email string) (bool, error) {
	count, err := db.queries.EmailExists(ctx, email)
	err = tracerr.Wrap(err)
	return count > 0, err
}

func (db *database) ExpireSessions() error {
	if err := db.queries.ExpireSessions(ctx, time.Now().UTC().Unix()); err != nil {
		err = tracerr.Wrap(err)
		return err
	}
	return nil
}

func (db *database) ExtendSession(session string) error {
	return tracerr.Wrap(db.queries.AddTimeToSession(ctx, session))
}

func (db *database) UpdateUsername(userID uint64, username string) error {
	return tracerr.Wrap(db.queries.UpdateUsername(ctx, queries.UpdateUsernameParams{
		Username: username,
		UserID:   userID,
	}))
}

func (db *database) GetAvatar(userID uint64) (sql.NullString, error) {
	ret, err := db.queries.GetAvatar(ctx, userID)
	err = tracerr.Wrap(err)
	return ret, err
}

func (db *database) UpdateAvatar(userID uint64, avatar string) error {
	return tracerr.Wrap(db.queries.UpdateAvatar(ctx, queries.UpdateAvatarParams{
		Avatar: utilities.ToSqlString(avatar),
		UserID: userID,
	}))
}

func (db *database) SetStatus(userID uint64, status harmonytypesv1.UserStatus) error {
	return tracerr.Wrap(db.queries.SetStatus(ctx, queries.SetStatusParams{
		Status: int16(status), // lol shut up it's an int16
		UserID: userID,
	}))
}

func (db *database) SetUsername(userID uint64, username string) error {
	return tracerr.Wrap(db.queries.UpdateUsername(ctx, queries.UpdateUsernameParams{
		UserID:   userID,
		Username: username,
	}))
}

func (db *database) SetAvatar(userID uint64, avatar string) error {
	return tracerr.Wrap(db.queries.UpdateAvatar(ctx, queries.UpdateAvatarParams{
		UserID: userID,
		Avatar: utilities.ToSqlString(avatar),
	}))
}

func (db *database) GetUserMetadata(userID uint64, appID string) (string, error) {
	metadata, err := db.queries.GetUserMetadata(ctx, queries.GetUserMetadataParams{
		UserID: userID,
		AppID:  appID,
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return metadata, err
}

func (db *database) GetNonceInfo(nonce string) (queries.GetNonceInfoRow, error) {
	info, err := db.queries.GetNonceInfo(ctx, nonce)
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return info, err
}

func (db *database) AddNonce(nonce string, userID uint64, homeServer string) error {
	err := db.queries.AddNonce(ctx, queries.AddNonceParams{
		Nonce:      nonce,
		UserID:     userID,
		HomeServer: homeServer,
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return err
}

func (db *database) GetGuildList(userID uint64) ([]queries.GetGuildListRow, error) {
	guilds, err := db.queries.GetGuildList(ctx, userID)
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return guilds, err
}

func (db *database) GetGuildListPosition(userID, guildID uint64, homeServer string) (string, error) {
	position, err := db.queries.GetGuildListPosition(ctx, queries.GetGuildListPositionParams{
		UserID:     userID,
		GuildID:    guildID,
		HomeServer: homeServer,
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return position, err
}

func (db *database) AddGuildToList(userID, guildID uint64, homeServer string) error {
	pos, err := db.queries.GetLastGuildPositionInList(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pos = ""
		} else {
			err = tracerr.Wrap(err)
			db.Logger.Exception(err)
			return err
		}
	}

	err = db.queries.AddToGuildList(ctx, queries.AddToGuildListParams{
		UserID:     userID,
		GuildID:    guildID,
		HomeServer: homeServer,
		Position:   Rank(pos, ""),
	})
	err = tracerr.Wrap(err)

	db.Logger.CheckException(err)
	return err
}

func (db *database) MoveGuild(userID, guildID uint64, homeServer string, nextGuildID, prevGuildID uint64, nextHomeServer, prevHomeServer string) error {
	nextPos, err := db.queries.GetGuildListPosition(ctx, queries.GetGuildListPositionParams{
		UserID:     userID,
		GuildID:    nextGuildID,
		HomeServer: nextHomeServer,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			nextPos = ""
		} else {
			err = tracerr.Wrap(err)
			db.Logger.Exception(err)
			return err
		}
	}

	prevPos, err := db.queries.GetGuildListPosition(ctx, queries.GetGuildListPositionParams{
		UserID:     userID,
		GuildID:    prevGuildID,
		HomeServer: prevHomeServer,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			nextPos = ""
		} else {
			err = tracerr.Wrap(err)
			db.Logger.Exception(err)
			return err
		}
	}

	err = db.queries.MoveGuild(ctx, queries.MoveGuildParams{
		Position:   Rank(prevPos, nextPos),
		GuildID:    guildID,
		HomeServer: homeServer,
	})
	err = tracerr.Wrap(err)

	db.Logger.CheckException(err)

	return err
}

func (db database) RemoveGuildFromList(userID, guildID uint64, homeServer string) error {
	err := db.queries.RemoveGuildFromList(ctx, queries.RemoveGuildFromListParams{
		UserID:     userID,
		GuildID:    guildID,
		HomeServer: homeServer,
	})
	err = tracerr.Wrap(err)
	db.Logger.CheckException(err)
	return err
}

func (db database) UserIsLocal(userID uint64) error {
	ok, err := db.queries.UserIsLocal(ctx, userID)
	if err == nil && !ok {
		err = types.ErrNotLocal
	}
	if err != types.ErrNotLocal {
		err = tracerr.Wrap(err)
	}
	return err
}

func (db database) GetAllMutuals(userID uint64) ([]uint64, error) {
	mutuals, err := db.queries.GetAllMutuals(ctx, userID)
	err = tracerr.Wrap(err)
	return mutuals, err
}

func (db database) SetIsBot(userID uint64, isBot bool) error {
	var newExpiration int64

	if err := db.queries.SetBot(ctx, queries.SetBotParams{
		IsBot:  isBot,
		UserID: userID,
	}); err != nil {
		return err
	}

	if isBot {
		newExpiration = math.MaxInt64
	} else {
		newExpiration = time.Now().UTC().Add(db.Config.Server.Policies.Sessions.Duration).Unix()
	}

	return db.queries.SetExpiration(ctx, queries.SetExpirationParams{
		UserID:     userID,
		Expiration: newExpiration,
	})
}

func (db database) IsBot(userID uint64) (bool, error) {
	return db.queries.IsBot(ctx, userID)
}
