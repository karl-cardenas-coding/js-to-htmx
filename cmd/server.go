package cmd

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/karl-cardenas-coding/js-to-htmx/internal"
)

// PageData is the data structure for the HTML template
type PageData struct {
	CoinName    string
	Price       float64
	LastUpdated string
}

func Server(ctx context.Context, args []string, stdout, stderr *os.File) error {

	// Set the default logger to text format. Default level is info and time format is changed to "2006/01/02 15:04:05" using local time
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:       slog.LevelInfo,
		ReplaceAttr: changeTimeFormat,
	})))

	// Serve static files from the web/static directory at /static/
	fs := http.FileServer(http.Dir("web/static"))
	// Strip the /static/ prefix from the URL path so that the files are served from / instead of /static/
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", landingPageHandler("web/index.html", PageData{}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	serverPort := "http://localhost:" + port
	slog.Info("Server started", "URL", serverPort)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return err
	}
	return nil
}

// landingPageHandler handles the landing page and writes the authentication URL to the page
func landingPageHandler(indexFile string, data PageData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmp, err := template.ParseFiles(indexFile)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			slog.Error("unable to parse template", "error", err)
		}

		btc, err := internal.FetchCoinPrice("BTC")
		if err != nil {
			slog.Error("unable to fetch coin price", "error", err)
		}

		data.CoinName = "Bitcoin"
		data.Price = btc.USD
		data.LastUpdated = time.Now().Local().Format("15:04:05")

		tmpl := template.Must(tmp, err)
		err = tmpl.Execute(w, data)
		if err != nil {
			slog.Error("unable to execute template", "error", err)
		}

	}
}

func changeTimeFormat(groups []string, a slog.Attr) slog.Attr {

	if a.Key == slog.TimeKey {
		a.Value = slog.StringValue(time.Now().Local().Format("2006/01/02 15:04:05"))
	}
	return a

}
