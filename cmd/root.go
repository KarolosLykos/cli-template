package cmd

import (
	"os"
	"os/signal"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// dateCmd command to print the date.
var rootCmd = &cobra.Command{
	Use:     "cli-template",
	Short:   "This cli template shows nothing",
	Long:    `This is a template CLI application, which can be used as a boilerplate for awesome CLI tools written in Go.`,
	Example: `cli-template`,
	Version: "v0.0.1", // <--VERSION-->
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		pterm.Print("exiting....")
		os.Exit(0)
	}()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
