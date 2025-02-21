package commands

import (
	"github.com/marcusfreisleben/gitlab-tart-executor/internal/commands/cleanup"
	"github.com/marcusfreisleben/gitlab-tart-executor/internal/commands/config"
	"github.com/marcusfreisleben/gitlab-tart-executor/internal/commands/prepare"
	"github.com/marcusfreisleben/gitlab-tart-executor/internal/commands/run"
	"github.com/marcusfreisleben/gitlab-tart-executor/internal/version"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	command := &cobra.Command{
		Use:           "executor",
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       version.FullVersion,
	}

	command.AddCommand(
		config.NewCommand(),
		prepare.NewCommand(),
		run.NewCommand(),
		cleanup.NewCommand(),
	)

	return command
}
