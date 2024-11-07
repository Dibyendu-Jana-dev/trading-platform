package dto

type MarketDataResponse struct {
    Symbol    string  `json:"symbol"`
    Price     string  `json:"price"`
    HighPrice string  `json:"highPrice"`
    LowPrice  string  `json:"lowPrice"`
    Volume    string  `json:"volume"`
    Timestamp int64   `json:"timestamp"`
}