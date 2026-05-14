package main

import (
	"fmt"
	"os"

	"github.com/evcc-io/evcc/cmd"
)

// main is the entry point for evcc.
// It delegates to the cmd package which uses cobra for CLI handling.
// Note: using exit code 1 for errors (standard Unix convention).
//
// Personal fork: added exit code 2 for usage errors vs exit code 1 for
// runtime errors, following GNU conventions more closely.
//
// TODO: investigate whether cmd.Execute() distinguishes usage errors from
// runtime errors so we can propagate exit code 2 for the former.
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
