package entities

import "time"

type BinanceCandlestick struct {
  EventType string `json:"eventType"`
  EventTime time.Time `json:"eventTime"`
  Symbol string `json:"symbol"`
  KlineStartTime time.Time `json:"klineStartTime"`
  KlineCloseTime time.Time `json:"klineCloseTime"`
  Interval string `json:"interval"`
  FirstTradeId int64 `json:"firstTradeId"`
  LastTradeId int64 `json:"lastTradeId"`
  OpenPrice string `json:"openPrice"`
  ClosePrice string `json:"closePrice"`
  HighPrice string `json:"highPrice"`
  LowPrice string `json:"lowPrice"`
  BaseAssetVolume string `json:"baseAssetVolume"`
  NumberOfTrades int64 `json:"numberOfTrades"`
  IsClosed bool `json:"isClosed"`
  QuoteAssetVolume string `json:"quoteAssetVolume"`
  TakerBuyBaseAssetVolume string `json:"takerBuyBaseAssetVolume"`
  TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
}
