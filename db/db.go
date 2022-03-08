// Package db contains the database implementation for legato.
package db

import (
	"context"
	"embed"

	"github.com/go-redis/redis/v8"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

//go:embed sqlc/migrations/*.sql
var Migrations embed.FS

// DB is the database structure.
type DB struct {
	Rdb      *redis.Client
	Postgres *pgxpool.Pool
	models   *models.Queries
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
		Rdb:      rdb,
		Postgres: conn,
		models:   models.New(conn),
	}, nil
}
