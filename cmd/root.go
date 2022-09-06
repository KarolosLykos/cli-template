package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cli-template",
		Short:   "This cli template shows nothing",
		Long:    `This is a template CLI application, which can be used as a boilerplate for awesome CLI tools written in Go.`,
		Example: `cli-template`,
		Version: "v0.0.1", // <--VERSION-->
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(cmd.UsageString())

			return nil
		},
	}

	cmd.AddCommand(NewDateCmd())    // version subcommand
	cmd.AddCommand(NewVersionCmd()) // version subcommand

	return cmd
}

// Execute invokes the command.
func Execute() error {
	if err := NewRootCmd().Execute(); err != nil {
		return fmt.Errorf("error executing root command: %w", err)
	}

	return nil
}
