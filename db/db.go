// Package db contains the database implementation for legato.
package db

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/harmony-development/legato/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

// DB is the database structure.
type DB struct {
	rdb  *redis.Client
	conn *pgxpool.Pool
}

// New creates a new DB instance.
func New(ctx context.Context, cfg *config.Config) (*DB, error) {
	redisOptions, err := redis.ParseURL(cfg.Redis.URL)
	if err != nil {
		return nil, Wrap(err, "failed to parse redis url")
	}

	rdb := redis.NewClient(redisOptions)

	conn, err := pgxpool.Connect(ctx, cfg.Postgres.URL)
	if err != nil {
		return nil, Wrap(err, "failed to connect to postgres")
	}

	return &DB{
		rdb:  rdb,
		conn: conn,
	}, nil
}

func (db *DB) GetCurrentMigration(ctx context.Context) (int, error) {
	var migration int
	err := db.conn.QueryRow(ctx, "SELECT current_migration FROM meta").Scan(&migration)

	return migration, TryWrap(err, "failed to get current migration")
}

func (db *DB) RawExec(c context.Context, query string, args ...interface{}) error {
	_, err := db.conn.Exec(c, query, args...)
	if err != nil {
		return fmt.Errorf("failed to run raw query: %w", err)
	}

	return nil
}
