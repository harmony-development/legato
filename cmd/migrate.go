// Package cmd contains commands for legato.
package cmd

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/google/subcommands"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type MigrateCmd struct{}

func (cmd *MigrateCmd) Name() string { return "migrate" }

func (cmd *MigrateCmd) Synopsis() string { return "runs a migration on the database." }

func (cmd *MigrateCmd) Usage() string {
	return `migrate`
}

func (cmd *MigrateCmd) SetFlags(f *flag.FlagSet) {}

func (cmd *MigrateCmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := cmd.RunMigrations(context.TODO()); err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}
	fmt.Println("migrations ran successfully")
	return subcommands.ExitSuccess
}

func (cmd *MigrateCmd) RunMigrations(ctx context.Context) error {
	m, err := newMigrate()
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

func newMigrate() (*migrate.Migrate, error) {
	cfg, err := config.ReadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	conn, err := sql.Open("pgx", cfg.Postgres.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres connection: %w", err)
	}

	driver, err := pgx.WithInstance(conn, &pgx.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx driver: %w", err)
	}

	migrations, err := iofs.New(db.Migrations, "sqlc/migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to create migrations source: %w", err)
	}

	m, err := migrate.NewWithInstance("iofs", migrations, "pgx", driver)
	if err != nil {
		return nil, fmt.Errorf("failed to load migrations: %w", err)
	}
	return m, err
}
