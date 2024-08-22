package coinbaseworker

type CoinbaseCommon struct {
	Channel     string          `json:"channel"`
	ClientId    string          `json:"client_id"`
	Timestamp   string          `json:"timestamp"`
	SequenceNum int64           `json:"sequence_num"`
	Events      []CoinbaseEvent `json:"events"`
}

type CoinbaseEvent struct {
	Type    string            `json:"type"`
	Candles []CoinbaseCandles `json:"candles"`
	Ticker  []CoinbaseTicker  `json:"ticker"`
}

type CoinbaseTicker struct {
	Type                  string `json:"type"`
	ProductId             string `json:"product_id"`
	Price                 string `json:"price"`
	Volume24H             string `json:"volume_24_h"`
	Low24H                string `json:"low_24_h"`
	High24H               string `json:"high_24_h"`
	Low24W                string `json:"low_24_w"`
	High24W               string `json:"high_24_w"`
	PricePercentChange24H string `json:"price_percent_chg_24_h"`
	BestBid               string `json:"best_bid"`
	BestBidQuantity       string `json:"best_bid_quantity"`
	BestAsk               string `json:"best_ask"`
	BestAskQuantity       string `json:"best_ask_quantity"`
}

type CoinbaseCandles struct {
	Start     string `json:"start"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Open      string `json:"open"`
	Close     string `json:"close"`
	Volume    string `json:"volume"`
	ProductId string `json:"product_id"`
}
