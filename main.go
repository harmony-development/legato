package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	"github.com/harmony-development/legato/cmd"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&cmd.StartCmd{}, "")
	subcommands.Register(&cmd.MigrateCmd{}, "")

	flag.Parse()

	os.Exit(int(subcommands.Execute(context.Background())))
}
