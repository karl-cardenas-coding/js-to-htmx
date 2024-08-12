package main

import (
	"context"
	"os"

	"github.com/karl-cardenas-coding/js-to-htmx/cmd"
)

func main() {

	ctx := context.Background()
	err := run(ctx, os.Args, os.Stdin, os.Stderr)
	if err != nil {
		os.Exit(1)
	}

}

func run(
	ctx context.Context,
	args []string,
	stdout,
	stderr *os.File,
) error {

	return cmd.Server(ctx, args, stdout, stderr)
}
