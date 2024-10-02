package models

// QuoteResponse represents the response from the API for fetching market quotes.
type QuoteResponse struct {
	Data struct {
		Symbol string  `json:"symbol"`
		LTP    float64 `json:"ltp"`
	} `json:"data"`
	Status string `json:"status"`
}

// MarketDepthResponse represents the response from the API for fetching market depth.
type MarketDepthResponse struct {
	Data struct {
		Bids []struct {
			Price float64 `json:"price"`
			Qty   int     `json:"qty"`
		} `json:"bids"`
		Asks []struct {
			Price float64 `json:"price"`
			Qty   int     `json:"qty"`
		} `json:"asks"`
	} `json:"data"`
	Status string `json:"status"`
}

// HistoricalDataResponse represents the response from the API for fetching historical data.
type HistoricalDataResponse struct {
	Data []struct {
		Time   int64   `json:"time"`
		Open   float64 `json:"open"`
		High   float64 `json:"high"`
		Low    float64 `json:"low"`
		Close  float64 `json:"close"`
		Volume int64   `json:"volume"`
	} `json:"data"`
	Status string `json:"status"`
}
