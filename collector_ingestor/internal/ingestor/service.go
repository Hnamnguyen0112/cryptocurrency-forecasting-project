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

		binanceTickerEntity := entities.BinanceTicker{}

		transformBinanceTicker(value, &binanceTickerEntity)

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

		binanceCandlestickEntity := entities.BinanceCandlestick{}

		transformBinanceCandlestick(value, &binanceCandlestickEntity)

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

		parseTime, err := time.Parse(time.RFC3339Nano, value.Timestamp)
		if err != nil {
			log.Printf("Error parsing time: %v\n", err)
		}

		for _, event := range value.Events {
			for _, ticker := range event.Tickers {
				coinbaseTickerEntity := entities.CoinbaseTicker{
					Timestamp: parseTime,
				}

				transformCoinbaseTicker(ticker, &coinbaseTickerEntity)

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

		parseTime, err := time.Parse(time.RFC3339Nano, value.Timestamp)
		if err != nil {
			log.Printf("Error parsing time: %v\n", err)
		}

		for _, event := range value.Events {
			for _, candle := range event.Candles {
				coinbaseCandleEntity := entities.CoinbaseCandle{
					Timestamp: parseTime,
				}

				transformCoinbaseCandle(candle, &coinbaseCandleEntity)

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
