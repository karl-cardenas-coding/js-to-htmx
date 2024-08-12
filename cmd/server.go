package cmd

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/karl-cardenas-coding/js-to-htmx/internal"
)

// PageData is the data structure for the HTML template
type PageData struct {
	CoinName     string
	Price        float64
	LastUpdated  string
	CoinLogoPath string
	News         []internal.News
}

// Server starts the server and serves the web pages
func Server(ctx context.Context, args []string, stdout, stderr *os.File) error {
	// Serve static files from the web/static directory at /static/
	fs := http.FileServer(http.Dir("web/static"))
	// Strip the /static/ prefix from the URL path so that the files are served from / instead of /static/
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", landingPageHandler("web/index.html", PageData{}))
	http.HandleFunc("/coin", coinPriceHandler("web/coin.html"))
	http.HandleFunc("/news", newsHandler("web/news.html"))

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
		w.Header().Set("Content-Type", "text/html")

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
		data.LastUpdated = time.Now().Local().Format("15:04:05 PM")
		data.CoinLogoPath = "/static/images/bitcoin.png"

		tmpl := template.Must(tmp, err)
		err = tmpl.Execute(w, data)
		if err != nil {
			slog.Error("unable to execute template", "error", err)
		}

	}
}

// coinPriceHandler handles the coin price page and writes the coin price to the page.
func coinPriceHandler(templateFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html")

		tmp, err := template.ParseFiles(templateFile)
		if err != nil {
			slog.Error("unable to parse template", "error", err)
		}

		coinSymbol := r.URL.Query().Get("symbol")
		if coinSymbol == "" {
			http.Error(w, "missing coin symbol", http.StatusBadRequest)
		}

		coin, err := internal.FetchCoinPrice(coinSymbol)
		if err != nil {
			http.Error(w, "unable to fetch coin price", http.StatusInternalServerError)
			return
		}

		data := PageData{
			CoinName:     coin.Name,
			Price:        coin.USD,
			LastUpdated:  time.Now().Local().Format("15:04:05 PM"),
			CoinLogoPath: "/static/images/" + strings.ToLower(coinSymbol) + ".png",
		}

		tmpl := template.Must(tmp, err)
		err = tmpl.Execute(w, data)
		if err != nil {
			slog.Error("unable to execute template", "error", err)
		}

	}

}

func newsHandler(templateFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html")

		tmp, err := template.ParseFiles(templateFile)
		if err != nil {
			slog.Error("unable to parse template", "error", err)
		}

		news, err := internal.FetchNews()
		if err != nil {
			http.Error(w, "unable to fetch news", http.StatusInternalServerError)
			return
		}

		data := PageData{
			News: news,
		}

		tmpl := template.Must(tmp, err)
		err = tmpl.Execute(w, data)
		if err != nil {
			slog.Error("unable to execute template", "error", err)
		}

	}

}
