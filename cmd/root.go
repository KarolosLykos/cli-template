package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd command to print the date.
var RootCmd = &cobra.Command{
	Use:     "cli-template",
	Short:   "This cli template shows nothing",
	Long:    `This is a template CLI application, which can be used as a boilerplate for awesome CLI tools written in Go.`,
	Example: `cli-template`,
	Version: "v0.0.1", // <--VERSION-->
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return RootCmd.Execute()
}
