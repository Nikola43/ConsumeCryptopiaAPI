package main
type MarketTrade struct {
	Success bool        `json:"Success"`
	Message interface{} `json:"Message"`
	Data    struct {
		TradePairID    int     `json:"TradePairId"`
		Label          string  `json:"Label"`
		AskPrice       float64 `json:"AskPrice"`
		BidPrice       float64 `json:"BidPrice"`
		Low            float64 `json:"Low"`
		High           float64 `json:"High"`
		Volume         float64 `json:"Volume"`
		LastPrice      float64 `json:"LastPrice"`
		BuyVolume      float64 `json:"BuyVolume"`
		SellVolume     float64 `json:"SellVolume"`
		Change         float64 `json:"Change"`
		Open           float64 `json:"Open"`
		Close          float64 `json:"Close"`
		BaseVolume     float64 `json:"BaseVolume"`
		BuyBaseVolume  float64 `json:"BuyBaseVolume"`
		SellBaseVolume float64 `json:"SellBaseVolume"`
	} `json:"Data"`
	Error interface{} `json:"Error"`
}
