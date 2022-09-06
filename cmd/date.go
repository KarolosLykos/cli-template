package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var availableFormats = []string{
	"02 01 06",
	"02 01 06 15:04:05",
	"02 Jan 06",
	"02 Jan 06 15:04:05",
	"Mon Jan 06",
	"Mon Jan 06 15:04:05",
	"02 01 2006",
	"02 01 2006 15:04:05 MST",
}

var ErrBadFormat = errors.New("wrong date format")

func NewDateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "date",
		Short: "Print the current date.",
		RunE:  run,
	}

	cmd.Flags().StringP("format", "f", "02 Jan 06", "specify a custom date format")

	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	format, err := cmd.Flags().GetString("format")
	if err != nil {
		return err
	}

	if !checkFormat(format) {
		return ErrBadFormat
	}

	fmt.Fprintf(cmd.OutOrStdout(), "%s\n", time.Now().Format(format))

	return nil
}

func checkFormat(format string) bool {
	found := false

	for _, f := range availableFormats {
		if f == format {
			found = true

			break
		}
	}

	return found
}
