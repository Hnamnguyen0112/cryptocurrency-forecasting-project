package entities

import "time"

type CoinbaseCandle struct {
	ProductId string    `json:"product_id"`
	Start     string    `json:"start"`
	High      string    `json:"high"`
	Low       string    `json:"low"`
	Open      string    `json:"open"`
	Close     string    `json:"close"`
	Volume    string    `json:"volume"`
	Timestamp time.Time `json:"timestamp"`
}
