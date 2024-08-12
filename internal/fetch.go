package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// CoinPrice is a struct that represents the price of a coin. It contains the name of the coin and the price in U.S. dollars.
type CoinPrice struct {
	// Name is the name of the coin
	Name string `json:"-"`
	// USD is the price of the coin in U.S. dollars
	USD float64 `json:"USD"`
}

type NewsFeed struct {
	Type      int           `json:"Type"`
	Message   string        `json:"Message"`
	Promoted  []interface{} `json:"Promoted"`
	Data      []News        `json:"Data"`
	RateLimit struct {
	} `json:"RateLimit"`
	HasWarning bool `json:"HasWarning"`
}

type News struct {
	ID          string `json:"id"`
	GUID        string `json:"guid"`
	PublishedOn int    `json:"published_on"`
	Imageurl    string `json:"imageurl"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	Body        string `json:"body"`
	Tags        string `json:"tags"`
	Lang        string `json:"lang"`
	Upvotes     string `json:"upvotes"`
	Downvotes   string `json:"downvotes"`
	Categories  string `json:"categories"`
	SourceInfo  struct {
		Name string `json:"name"`
		Img  string `json:"img"`
		Lang string `json:"lang"`
	} `json:"source_info"`
}

// FetchCoinPrice fetches the current price of a coin. The return value is a CoinPrice struct.
// The symbol parameter is the coin symbol (e.g. BTC, ETH, USDC).
/*
  Example:
  btc, err := internal.FetchCoinPrice("BTC")
*/
func FetchCoinPrice(symbol string) (CoinPrice, error) {

	var (
		output CoinPrice
	)
	client := http.DefaultClient

	url := fmt.Sprintf("https://min-api.cryptocompare.com/data/price?fsym=%s&tsyms=USD", symbol)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return output, err
	}
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return output, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return output, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return output, err
	}
	err = json.Unmarshal(body, &output)
	if err != nil {
		return output, fmt.Errorf("unable to unmarshal coin data %w", err)
	}

	output.Name = getCoinName(symbol)

	return output, nil
}

func getCoinName(symbol string) string {
	switch strings.ToUpper(symbol) {
	case "BTC":
		return "Bitcoin"
	case "ETH":
		return "Ethereum"
	case "USDC":
		return "USDC"
	default:
		return "Unknown"
	}
}

// FetchNews fetches the latest news articles related to cryptocurrency. The return value is a slice of News structs.
// The function returns the top 5 news articles.
func FetchNews() ([]News, error) {
	var (
		output []News
	)
	client := http.DefaultClient

	url := "https://min-api.cryptocompare.com/data/v2/news/?lang=EN"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return output, err
	}
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return output, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return output, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return output, err
	}

	var newsList NewsFeed
	err = json.Unmarshal(body, &newsList)
	if err != nil {
		return output, fmt.Errorf("unable to unmarshal news data %w", err)
	}

	if len(newsList.Data) < 5 {
		return newsList.Data, nil
	}

	output = newsList.Data[:5]

	return output, nil
}
