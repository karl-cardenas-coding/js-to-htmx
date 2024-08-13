package main

import (
	"context"
	"embed"
	"log/slog"
	"os"
	"time"

	"github.com/karl-cardenas-coding/js-to-htmx/cmd"
)

//go:embed web
var web embed.FS

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

	// Set the default logger to text format. Default level is info and time format is changed to "2006/01/02 15:04:05" using local time
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:       slog.LevelInfo,
		ReplaceAttr: changeTimeFormat,
	})))

	return cmd.Server(ctx, args, stdout, stderr, web)
}

// changeTimeFormat is a custom attribute replacer that changes the time format to "2006/01/02 15:04:05"
func changeTimeFormat(groups []string, a slog.Attr) slog.Attr {

	if a.Key == slog.TimeKey {
		a.Value = slog.StringValue(time.Now().Local().Format("2006/01/02 15:04:05"))
	}
	return a

}
