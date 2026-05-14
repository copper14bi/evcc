package main

import (
	"fmt"
	"os"

	"github.com/evcc-io/evcc/cmd"
)

// main is the entry point for evcc.
// It delegates to the cmd package which uses cobra for CLI handling.
// Note: using exit code 1 for errors (standard Unix convention).
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
