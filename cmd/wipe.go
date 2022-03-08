// Package cmd contains commands for legato.
package cmd

import (
	"context"
	"flag"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/subcommands"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type WipeCMD struct {
	yes bool
}

func (cmd *WipeCMD) Name() string { return "wipe" }

func (cmd *WipeCMD) Synopsis() string { return "wipes the database." }

func (cmd *WipeCMD) Usage() string {
	return `wipe`
}

func (cmd *WipeCMD) SetFlags(f *flag.FlagSet) {
	f.Bool("yes", false, "wipe the database without asking")
}

func (cmd *WipeCMD) Execute(_ context.Context, flags *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if !cmd.yes {
		fmt.Println("are you sure?")
	}

	m, err := newMigrate()
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	if err := m.Down(); err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	fmt.Println("database wiped")

	return subcommands.ExitSuccess
}
