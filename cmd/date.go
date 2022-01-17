package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "Print the current date.",
	Run: func(cmd *cobra.Command, args []string) {
		format, _ := cmd.Flags().GetString("format")
		fmt.Println(time.Now().Format(format))
	},
}

func init() {
	rootCmd.AddCommand(dateCmd)

	dateCmd.Flags().StringP("format", "f", "02 Jan 06", "specify a custom date format")
}
