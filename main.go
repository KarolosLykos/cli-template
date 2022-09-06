package main

import (
	"fmt"
	"os"

	"github.com/KarolosLykos/cli-template/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
