package fetcherservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)
func GetMarketDataGetMarketData(symbol string) (MarketData, error) {
	url := fmt.Sprintf("%s?symbol=%s", "https://api.binance.com/api/v3/ticker/bookTicker", symbol)

	// Make the HTTP request to Binance API
	resp, err := http.Get(url)
	if err != nil {
		return MarketData{}, fmt.Errorf("failed to fetch market data: %v", err)
	}
	defer resp.Body.Close()

	var data struct {
		Symbol   string  `json:"symbol"`
		BidPrice string  `json:"bidPrice"`
		AskPrice string  `json:"askPrice"`
		LastPrice string `json:"lastPrice"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return MarketData{}, fmt.Errorf("failed to decode market data response: %v", err)
	}

	// Convert string prices to float64
	bidPrice, err := strconv.ParseFloat(data.BidPrice, 64)
	if err != nil {
		return MarketData{}, fmt.Errorf("invalid bid price: %v", err)
	}

	askPrice, err := strconv.ParseFloat(data.AskPrice, 64)
	if err != nil {
		return MarketData{}, fmt.Errorf("invalid ask price: %v", err)
	}

	lastPrice, err := strconv.ParseFloat(data.LastPrice, 64)
	if err != nil {
		return MarketData{}, fmt.Errorf("invalid last price: %v", err)
	}

	// Return the parsed market data
	return MarketData{
		Symbol:   data.Symbol,
		BidPrice: bidPrice,
		AskPrice: askPrice,
		LastPrice: lastPrice,
	}, nil
}