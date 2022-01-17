package cmd

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "cli-template",
	Short:   "This cli template shows nothing",
	Long:    `This is a template CLI application, which can be used as a boilerplate for awesome CLI tools written in Go.`,
	Example: `cli-template`,
	Version: "v0.0.1", // <--VERSION-->
}

func Execute() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		fmt.Println("exiting....")
		os.Exit(0)
	}()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
