// Package cmd contains commands for legato.
package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db"
	"github.com/harmony-development/legato/db/migrations"
)

type MigrateCmd struct{}

func (cmd *MigrateCmd) Name() string { return "migrate" }

func (cmd *MigrateCmd) Synopsis() string { return "runs a migration on the database." }

func (cmd *MigrateCmd) Usage() string {
	return `migrate`
}

func (cmd *MigrateCmd) SetFlags(f *flag.FlagSet) {}

func (cmd *MigrateCmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := RunMigrations(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return subcommands.ExitFailure
	} else {
		return subcommands.ExitSuccess
	}
}

func RunMigrations() error {
	cfg, err := config.ReadConfig()
	if err != nil {
		return err
	}

	conn, err := db.New(context.Background(), cfg)
	if err != nil {
		return err
	}

	current, err := conn.GetCurrentMigration(context.Background())
	if err != nil {
		current = -1
	}
	for i := current + 1; i <= len(migrations.Migrations)-1; i++ {
		fmt.Printf("Running migration (%d/%d)\n", i, len(migrations.Migrations))
		if err := migrations.Migrations[i](conn); err != nil {
			return fmt.Errorf("failed to run migration %d: %w", i, err)
		}
	}

	fmt.Println("Migrations complete")

	return nil
}
