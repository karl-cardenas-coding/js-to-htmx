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
