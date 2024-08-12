package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CoinPrice struct {
	USD float64 `json:"USD"`
}

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

	return output, nil
}
