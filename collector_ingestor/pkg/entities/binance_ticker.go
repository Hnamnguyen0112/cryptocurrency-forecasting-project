package entities

import "time"

type BinanceTicker struct {
	EventType                   string    `json:"eventType"`
	EventTime                   time.Time `json:"eventTime"`
	Symbol                      string    `json:"symbol"`
	PriceChange                 string    `json:"priceChange"`
	PriceChangePercent          string    `json:"priceChangePercent"`
	WeightedAveragePrice        string    `json:"weightedAveragePrice"`
	LastPrice                   string    `json:"lastPrice"`
	LastQuantity                string    `json:"lastQuantity"`
	OpenPrice                   string    `json:"openPrice"`
	HighPrice                   string    `json:"highPrice"`
	LowPrice                    string    `json:"lowPrice"`
	TotalTradedBaseAssetVolume  string    `json:"totalTradedBaseAssetVolume"`
	TotalTradedQuoteAssetVolume string    `json:"totalTradedQuoteAssetVolume"`
	StatisticsOpenTime          time.Time `json:"statisticsOpenTime"`
	StatisticsCloseTime         time.Time `json:"statisticsCloseTime"`
	FirstTradeId                int64     `json:"firstTradeId"`
	LastTradeId                 int64     `json:"lastTradeId"`
	TotalNumberOfTrades         int64     `json:"totalNumberOfTrades"`
}
