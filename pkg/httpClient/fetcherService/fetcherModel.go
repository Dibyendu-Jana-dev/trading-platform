// MarketData represents the structure of the market data
package fetcherservice

type MarketData struct {
	Symbol   string  `json:"symbol"`
	BidPrice float64  `json:"bidPrice"`
	AskPrice float64  `json:"askPrice"`
	LastPrice float64 `json:"lastPrice"`
}