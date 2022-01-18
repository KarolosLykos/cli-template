package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var DateCmd = &cobra.Command{
	Use:   "date",
	Short: "Print the current date.",
	Run: func(cmd *cobra.Command, args []string) {
		format, _ := cmd.Flags().GetString("format")
		fmt.Fprintf(cmd.OutOrStdout(), time.Now().Format(format))
	},
}

func init() {
	RootCmd.AddCommand(DateCmd)

	DateCmd.Flags().StringP("format", "f", "02 Jan 06", "specify a custom date format")
}
