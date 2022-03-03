package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/server"
)

type StartCmd struct{}

func (cmd *StartCmd) Name() string { return "start" }

func (cmd *StartCmd) Synopsis() string { return "starts the server." }

func (cmd *StartCmd) Usage() string {
	return `start`
}

func (cmd *StartCmd) SetFlags(f *flag.FlagSet) {}

// Start starts the server.
func (cmd *StartCmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := cmd.start(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return subcommands.ExitFailure
	} else {
		return subcommands.ExitSuccess
	}
}

func (cmd *StartCmd) start() error {
	cfg, err := config.ReadConfig()
	if err != nil {
		return err
	}

	s := server.New(cfg)

	return s.Start()
}
