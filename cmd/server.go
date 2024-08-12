package cmd

import (
	"context"
	"embed"
	"io/fs"
	"os"
)

// PageData is the data structure for the HTML template
type PageData struct {
}

func Server(ctx context.Context, args []string, stdout, stderr *os.File, staticAssets embed.FS) error {

	staticAssets, err := getStaticAssets(staticAssets, "web/static")
	if err != nil {
		return err
	}

	return nil
}

// getStaticAssets returns the static assets from the embed.FS
func getStaticAssets(f embed.FS, filePath string) (fs.FS, error) {
	return fs.Sub(f, filePath)
}
