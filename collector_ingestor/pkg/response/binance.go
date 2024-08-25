package response

type BinanceCommon struct {
	EventType string `json:"e"`
	EventTime int64  `json:"E"`
	Symbol    string `json:"s"`
}

type BinanceTicker struct {
	BinanceCommon
	PriceChange                 string `json:"p"`
	PriceChangePercent          string `json:"P"`
	WeightedAveragePrice        string `json:"w"`
	LastPrice                   string `json:"c"`
	LastQuantity                string `json:"Q"`
	OpenPrice                   string `json:"o"`
	HighPrice                   string `json:"h"`
	LowPrice                    string `json:"l"`
	TotalTradedBaseAssetVolume  string `json:"v"`
	TotalTradedQuoteAssetVolume string `json:"q"`
	StatisticsOpenTime          int64  `json:"O"`
	StatisticsCloseTime         int64  `json:"C"`
	FirstTradeId                int64  `json:"F"`
	LastTradeId                 int64  `json:"L"`
	TotalNumberOfTrades         int64  `json:"n"`
}

type Kline struct {
	KlineStartTime           int64  `json:"t"`
	KlineCloseTime           int64  `json:"T"`
	Symbol                   string `json:"s"`
	Interval                 string `json:"i"`
	FirstTradeId             int64  `json:"f"`
	LastTradeId              int64  `json:"L"`
	OpenPrice                string `json:"o"`
	ClosePrice               string `json:"c"`
	HighPrice                string `json:"h"`
	LowPrice                 string `json:"l"`
	BaseAssetVolume          string `json:"v"`
	NumberOfTrades           int64  `json:"n"`
	IsClosed                 bool   `json:"x"`
	QuoteAssetVolume         string `json:"q"`
	TakerBuyBaseAssetVolume  string `json:"V"`
	TakerBuyQuoteAssetVolume string `json:"Q"`
}

type BinanceCandlestick struct {
	BinanceCommon
	Kline Kline `json:"k"`
}
