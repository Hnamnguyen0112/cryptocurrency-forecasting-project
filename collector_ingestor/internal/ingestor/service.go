package ingestor

import (
	"log"
	"time"

	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/database"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/entities"
	kafkaPkg "github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/kafka"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/response"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func HandleMessage(e *kafka.Message, scr *kafkaPkg.SchemaRegistry) {
	DB := database.DB
	switch *e.TopicPartition.Topic {
	case "binance_ticker":
		value := response.BinanceTicker{}
		err := scr.Deser.DeserializeInto(*e.TopicPartition.Topic, e.Value, &value)
		if err != nil {
			log.Printf("Error deserializing binance ticker: %v\n", err)
		}

		binanceTickerEntity := entities.BinanceTicker{
			EventType:                   value.EventType,
			EventTime:                   time.Unix(0, value.EventTime*int64(time.Millisecond)),
			Symbol:                      value.Symbol,
			PriceChange:                 value.PriceChange,
			PriceChangePercent:          value.PriceChangePercent,
			WeightedAveragePrice:        value.WeightedAveragePrice,
			LastPrice:                   value.LastPrice,
			LastQuantity:                value.LastQuantity,
			OpenPrice:                   value.OpenPrice,
			HighPrice:                   value.HighPrice,
			LowPrice:                    value.LowPrice,
			TotalTradedBaseAssetVolume:  value.TotalTradedBaseAssetVolume,
			TotalTradedQuoteAssetVolume: value.TotalTradedQuoteAssetVolume,
			StatisticsOpenTime: time.Unix(
				0,
				value.StatisticsOpenTime*int64(time.Millisecond),
			),
			StatisticsCloseTime: time.Unix(
				0,
				value.StatisticsCloseTime*int64(time.Millisecond),
			),
			FirstTradeId:        value.FirstTradeId,
			LastTradeId:         value.LastTradeId,
			TotalNumberOfTrades: value.TotalNumberOfTrades,
		}

		err = DB.Create(&binanceTickerEntity).Error
		if err != nil {
			log.Printf("Error creating binance ticker entity: %v\n", err)
		}
	case "binance_candlestick":
		value := response.BinanceCandlestick{}
		err := scr.Deser.DeserializeInto(*e.TopicPartition.Topic, e.Value, &value)
		if err != nil {
			log.Printf("Error deserializing binance candlestick: %v\n", err)
		}

		binanceCandlestickEntity := entities.BinanceCandlestick{
			EventType: value.EventType,
			EventTime: time.Unix(0, value.EventTime*int64(time.Millisecond)),
			Symbol:    value.Symbol,
			KlineStartTime: time.Unix(
				0,
				value.Kline.KlineStartTime*int64(time.Millisecond),
			),
			KlineCloseTime: time.Unix(
				0,
				value.Kline.KlineCloseTime*int64(time.Millisecond),
			),
			Interval:                 value.Kline.Interval,
			FirstTradeId:             value.Kline.FirstTradeId,
			LastTradeId:              value.Kline.LastTradeId,
			OpenPrice:                value.Kline.OpenPrice,
			ClosePrice:               value.Kline.ClosePrice,
			HighPrice:                value.Kline.HighPrice,
			LowPrice:                 value.Kline.LowPrice,
			BaseAssetVolume:          value.Kline.BaseAssetVolume,
			NumberOfTrades:           value.Kline.NumberOfTrades,
			IsClosed:                 value.Kline.IsClosed,
			QuoteAssetVolume:         value.Kline.QuoteAssetVolume,
			TakerBuyBaseAssetVolume:  value.Kline.TakerBuyBaseAssetVolume,
			TakerBuyQuoteAssetVolume: value.Kline.TakerBuyQuoteAssetVolume,
		}

		err = DB.Create(&binanceCandlestickEntity).Error
		if err != nil {
			log.Printf("Error creating binance candlestick entity: %v\n", err)
		}
	case "coinbase_ticker":
		value := response.CoinbaseCommon{}
		err := scr.Deser.DeserializeInto(*e.TopicPartition.Topic, e.Value, &value)
		if err != nil {
			log.Printf("Error deserializing coinbase ticker: %v\n", err)
		}

		var coinbaseTickerEntities []entities.CoinbaseTicker

		for _, event := range value.Events {
			for _, ticker := range event.Tickers {
				parseTime, err := time.Parse(time.RFC3339Nano, value.Timestamp)
				if err != nil {
					log.Printf("Error parsing time: %v\n", err)
				}
				coinbaseTickerEntity := entities.CoinbaseTicker{
					ProductId:          ticker.ProductId,
					Price:              ticker.Price,
					Volume24h:          ticker.Volume24H,
					Low24h:             ticker.Low24H,
					High24h:            ticker.High24H,
					Low52w:             ticker.Low52W,
					High52w:            ticker.High52W,
					PricePercentChg24h: ticker.PricePercentChange24H,
					BestBid:            ticker.BestBid,
					BestBidQuantity:    ticker.BestBidQuantity,
					BestAsk:            ticker.BestAsk,
					BestAskQuantity:    ticker.BestAskQuantity,
					Timestamp:          parseTime,
				}

				coinbaseTickerEntities = append(coinbaseTickerEntities, coinbaseTickerEntity)
			}
		}

		err = DB.Create(&coinbaseTickerEntities).Error
		if err != nil {
			log.Printf("Error creating coinbase ticker entity: %v\n", err)
		}
	case "coinbase_candles":
		value := response.CoinbaseCommon{}
		err := scr.Deser.DeserializeInto(*e.TopicPartition.Topic, e.Value, &value)
		if err != nil {
			log.Printf("Error deserializing coinbase candles: %v\n", err)
		}

		var coinbaseCandleEntities []entities.CoinbaseCandle

		for _, event := range value.Events {
			for _, candle := range event.Candles {
				parseTime, err := time.Parse(time.RFC3339Nano, value.Timestamp)
				if err != nil {
					log.Printf("Error parsing time: %v\n", err)
				}

				coinbaseCandleEntity := entities.CoinbaseCandle{
					Start:     candle.Start,
					High:      candle.High,
					Low:       candle.Low,
					Open:      candle.Open,
					Close:     candle.Close,
					Volume:    candle.Volume,
					ProductId: candle.ProductId,
					Timestamp: parseTime,
				}

				coinbaseCandleEntities = append(coinbaseCandleEntities, coinbaseCandleEntity)
			}
		}

		err = DB.Create(&coinbaseCandleEntities).Error
		if err != nil {
			log.Printf("Error creating coinbase candle entity: %v\n", err)
		}
	default:
	}
}
