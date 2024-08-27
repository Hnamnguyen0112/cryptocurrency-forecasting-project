package ingestor

import (
	"time"

	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/entities"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/response"
)

func transformBinanceTicker(message response.BinanceTicker, entity *entities.BinanceTicker) {
	entity.EventType = message.EventType
	entity.EventTime = time.Unix(0, message.EventTime*int64(time.Millisecond))
	entity.Symbol = message.Symbol
	entity.PriceChange = message.PriceChange
	entity.PriceChangePercent = message.PriceChangePercent
	entity.WeightedAveragePrice = message.WeightedAveragePrice
	entity.LastPrice = message.LastPrice
	entity.LastQuantity = message.LastQuantity
	entity.OpenPrice = message.OpenPrice
	entity.HighPrice = message.HighPrice
	entity.LowPrice = message.LowPrice
	entity.TotalTradedBaseAssetVolume = message.TotalTradedBaseAssetVolume
	entity.TotalTradedQuoteAssetVolume = message.TotalTradedQuoteAssetVolume
	entity.StatisticsOpenTime = time.Unix(0, message.StatisticsOpenTime*int64(time.Millisecond))
	entity.StatisticsCloseTime = time.Unix(0, message.StatisticsCloseTime*int64(time.Millisecond))
	entity.FirstTradeId = message.FirstTradeId
	entity.LastTradeId = message.LastTradeId
	entity.TotalNumberOfTrades = message.TotalNumberOfTrades
}

func transformBinanceCandlestick(
	message response.BinanceCandlestick,
	entity *entities.BinanceCandlestick,
) {
	entity.EventType = message.EventType
	entity.EventTime = time.Unix(0, message.EventTime*int64(time.Millisecond))
	entity.Symbol = message.Symbol
	entity.KlineStartTime = time.Unix(0, message.Kline.KlineStartTime*int64(time.Millisecond))
	entity.KlineCloseTime = time.Unix(0, message.Kline.KlineCloseTime*int64(time.Millisecond))
	entity.Interval = message.Kline.Interval
	entity.FirstTradeId = message.Kline.FirstTradeId
	entity.LastTradeId = message.Kline.LastTradeId
	entity.OpenPrice = message.Kline.OpenPrice
	entity.ClosePrice = message.Kline.ClosePrice
	entity.HighPrice = message.Kline.HighPrice
	entity.LowPrice = message.Kline.LowPrice
	entity.BaseAssetVolume = message.Kline.BaseAssetVolume
	entity.NumberOfTrades = message.Kline.NumberOfTrades
	entity.IsClosed = message.Kline.IsClosed
	entity.QuoteAssetVolume = message.Kline.QuoteAssetVolume
	entity.TakerBuyBaseAssetVolume = message.Kline.TakerBuyBaseAssetVolume
	entity.TakerBuyQuoteAssetVolume = message.Kline.TakerBuyQuoteAssetVolume
}

func transformCoinbaseTicker(message response.CoinbaseTicker, entity *entities.CoinbaseTicker) {
	entity.ProductId = message.ProductId
	entity.Price = message.Price
	entity.Volume24h = message.Volume24H
	entity.Low24h = message.Low24H
	entity.High24h = message.High24H
	entity.Low52w = message.Low52W
	entity.High52w = message.High52W
	entity.PricePercentChg24h = message.PricePercentChange24H
	entity.BestBid = message.BestBid
	entity.BestBidQuantity = message.BestBidQuantity
	entity.BestAsk = message.BestAsk
	entity.BestAskQuantity = message.BestAskQuantity
}

func transformCoinbaseCandle(message response.CoinbaseCandle, entity *entities.CoinbaseCandle) {
	entity.ProductId = message.ProductId
	entity.Start = message.Start
	entity.High = message.High
	entity.Low = message.Low
	entity.Open = message.Open
	entity.Close = message.Close
	entity.Volume = message.Volume
}
