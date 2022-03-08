package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	"github.com/harmony-development/legato/cmd"
	"github.com/harmony-development/legato/util/panic_fmt"
	"github.com/joho/godotenv"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			panic_fmt.RecoverError(os.Getenv("RAW_LOG") != "true", r)
		}
	}()

	_ = godotenv.Load()

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&cmd.StartCmd{}, "")
	subcommands.Register(&cmd.MigrateCmd{}, "")
	subcommands.Register(&cmd.WipeCMD{}, "")

	flag.Parse()

	os.Exit(int(subcommands.Execute(context.Background())))
}
