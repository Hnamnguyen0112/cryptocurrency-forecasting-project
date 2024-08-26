package entities

import "time"

type CoinbaseTicker struct {
	ProductId          string    `json:"product_id"`
	Price              string    `json:"price"`
	Volume24h          string    `json:"volume_24_h"`
	Low24h             string    `json:"low_24_h"`
	High24h            string    `json:"high_24_h"`
	Low52w             string    `json:"low_52_w"`
	High52w            string    `json:"high_52_w"`
	PricePercentChg24h string    `json:"price_percent_chg_24_h"`
	BestBid            string    `json:"best_bid"`
	BestBidQuantity    string    `json:"best_bid_quantity"`
	BestAsk            string    `json:"best_ask"`
	BestAskQuantity    string    `json:"best_ask_quantity"`
	Timestamp          time.Time `json:"timestamp"`
}
