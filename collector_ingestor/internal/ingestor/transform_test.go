package ingestor

import (
	"testing"
	"time"

	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/entities"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/response"
)

func TestTransformBinanceTicker(t *testing.T) {
	message := response.BinanceTicker{
		BinanceCommon: response.BinanceCommon{
			EventType: "eventType",
			EventTime: 1625152800000,
			Symbol:    "BTCUSDT",
		},
		PriceChange:                 "0.001",
		PriceChangePercent:          "0.5",
		WeightedAveragePrice:        "50000",
		LastPrice:                   "50000.5",
		LastQuantity:                "0.001",
		OpenPrice:                   "49900",
		HighPrice:                   "51000",
		LowPrice:                    "49500",
		TotalTradedBaseAssetVolume:  "1000",
		TotalTradedQuoteAssetVolume: "50000000",
		StatisticsOpenTime:          1625149200000,
		StatisticsCloseTime:         1625235600000,
		FirstTradeId:                1,
		LastTradeId:                 100,
		TotalNumberOfTrades:         1000,
	}

	expected := entities.BinanceTicker{
		EventType:                   "eventType",
		EventTime:                   time.Unix(0, 1625152800000*int64(time.Millisecond)),
		Symbol:                      "BTCUSDT",
		PriceChange:                 "0.001",
		PriceChangePercent:          "0.5",
		WeightedAveragePrice:        "50000",
		LastPrice:                   "50000.5",
		LastQuantity:                "0.001",
		OpenPrice:                   "49900",
		HighPrice:                   "51000",
		LowPrice:                    "49500",
		TotalTradedBaseAssetVolume:  "1000",
		TotalTradedQuoteAssetVolume: "50000000",
		StatisticsOpenTime:          time.Unix(0, 1625149200000*int64(time.Millisecond)),
		StatisticsCloseTime:         time.Unix(0, 1625235600000*int64(time.Millisecond)),
		FirstTradeId:                1,
		LastTradeId:                 100,
		TotalNumberOfTrades:         1000,
	}

	var entity entities.BinanceTicker

	transformBinanceTicker(message, &entity)

	if entity != expected {
		t.Errorf("Expected %+v, got %+v", expected, entity)
	}
}

func TestTransformBinanceCandlestick(t *testing.T) {
	message := response.BinanceCandlestick{
		BinanceCommon: response.BinanceCommon{
			EventType: "eventType",
			EventTime: 1625152800000,
			Symbol:    "BTCUSDT",
		},
		Kline: response.Kline{
			KlineStartTime:           1625149200000,
			KlineCloseTime:           1625156400000,
			Interval:                 "1m",
			FirstTradeId:             1,
			LastTradeId:              100,
			OpenPrice:                "50000",
			ClosePrice:               "50000.5",
			HighPrice:                "51000",
			LowPrice:                 "49500",
			BaseAssetVolume:          "1000",
			NumberOfTrades:           1000,
			IsClosed:                 true,
			QuoteAssetVolume:         "50000000",
			TakerBuyBaseAssetVolume:  "500",
			TakerBuyQuoteAssetVolume: "25000000",
		},
	}

	expected := entities.BinanceCandlestick{
		EventType:                "eventType",
		EventTime:                time.Unix(0, 1625152800000*int64(time.Millisecond)),
		Symbol:                   "BTCUSDT",
		KlineStartTime:           time.Unix(0, 1625149200000*int64(time.Millisecond)),
		KlineCloseTime:           time.Unix(0, 1625156400000*int64(time.Millisecond)),
		Interval:                 "1m",
		FirstTradeId:             1,
		LastTradeId:              100,
		OpenPrice:                "50000",
		ClosePrice:               "50000.5",
		HighPrice:                "51000",
		LowPrice:                 "49500",
		BaseAssetVolume:          "1000",
		NumberOfTrades:           1000,
		IsClosed:                 true,
		QuoteAssetVolume:         "50000000",
		TakerBuyBaseAssetVolume:  "500",
		TakerBuyQuoteAssetVolume: "25000000",
	}

	var entity entities.BinanceCandlestick

	transformBinanceCandlestick(message, &entity)

	if entity != expected {
		t.Errorf("Expected %+v, got %+v", expected, entity)
	}
}

func TestTransformCoinbaseTicker(t *testing.T) {
	message := response.CoinbaseTicker{
		ProductId:             "BTC-USD",
		Price:                 "50000",
		Volume24H:             "1000",
		Low24H:                "49500",
		High24H:               "51000",
		Low52W:                "30000",
		High52W:               "60000",
		PricePercentChange24H: "1.5",
		BestBid:               "49950",
		BestBidQuantity:       "0.5",
		BestAsk:               "50050",
		BestAskQuantity:       "0.7",
	}

	// Expected output entity
	expected := entities.CoinbaseTicker{
		ProductId:          "BTC-USD",
		Price:              "50000",
		Volume24h:          "1000",
		Low24h:             "49500",
		High24h:            "51000",
		Low52w:             "30000",
		High52w:            "60000",
		PricePercentChg24h: "1.5",
		BestBid:            "49950",
		BestBidQuantity:    "0.5",
		BestAsk:            "50050",
		BestAskQuantity:    "0.7",
	}

	var entity entities.CoinbaseTicker

	transformCoinbaseTicker(message, &entity)

	if entity != expected {
		t.Errorf("Expected %+v, got %+v", expected, entity)
	}
}

func TestTransformCoinbaseCandle(t *testing.T) {
	message := response.CoinbaseCandle{
		ProductId: "BTC-USD",
		Start:     "1625152800",
		High:      "51000",
		Low:       "49500",
		Open:      "50000",
		Close:     "50500",
		Volume:    "1000",
	}

	// Expected output entity
	expected := entities.CoinbaseCandle{
		ProductId: "BTC-USD",
		Start:     "1625152800",
		High:      "51000",
		Low:       "49500",
		Open:      "50000",
		Close:     "50500",
		Volume:    "1000",
	}

	var entity entities.CoinbaseCandle

	transformCoinbaseCandle(message, &entity)

	if entity != expected {
		t.Errorf("Expected %+v, got %+v", expected, entity)
	}
}
