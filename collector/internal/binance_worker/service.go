package binanceworker

import (
	"encoding/json"
	"time"

	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/pkg/database"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/pkg/entities"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/pkg/websocket"
)

type Service struct {
  WS *websocket.Websocket
  TotalSubscribed int
}

func NewService(ws *websocket.Websocket) *Service {
  return &Service{
    WS: ws,
    TotalSubscribed: 0,
  }
}

func (s *Service) Subcribe(symbol string, channel string) error {
  request := Request{
    Method: "SUBSCRIBE",
    Params: []string{symbol + "@" + channel},
    ID: s.TotalSubscribed + 1,
  }

  if err := s.WS.Conn.WriteJSON(request); err != nil {
    return err
  }

  s.TotalSubscribed++

  return nil
}

func (s *Service) Listen() {
  for {
    _, message, err := s.WS.Conn.ReadMessage()
    if err != nil {
      return
    }

    var binanceType BinanceCommon
    err = json.Unmarshal(message, &binanceType)
    if err != nil {
      return
    }

    switch binanceType.EventType {
    case "24hrTicker":
      ticker := BinanceTicker{}
      if err := json.Unmarshal(message, &ticker); err != nil {
        return
      }

      if err := SaveTicker(ticker); err != nil {
        return
      }
    case "kline":
      candlestick := BinanceCandlestick{}
      if err := json.Unmarshal(message, &candlestick); err != nil { 
        return
      }
      
      if err := SaveCandlestick(candlestick); err != nil {
        return
      }
    default:
    }
  }
}

func SaveTicker(ticker BinanceTicker) error {
  ticketEntity := entities.BinanceTicker{
    EventType: ticker.EventType,
    EventTime: time.Unix(0, ticker.EventTime*int64(time.Millisecond)),
    Symbol: ticker.Symbol,
    PriceChange: ticker.PriceChange,
    PriceChangePercent: ticker.PriceChangePercent,
    WeightedAveragePrice: ticker.WeightedAveragePrice,
    LastPrice: ticker.LastPrice,
    LastQuantity: ticker.LastQuantity,
    OpenPrice: ticker.OpenPrice,
    HighPrice: ticker.HighPrice,
    LowPrice: ticker.LowPrice,
    TotalTradedBaseAssetVolume: ticker.TotalTradedBaseAssetVolume,
    TotalTradedQuoteAssetVolume: ticker.TotalTradedQuoteAssetVolume,
    StatisticsOpenTime: time.Unix(0, ticker.StatisticsOpenTime*int64(time.Millisecond)),
    StatisticsCloseTime: time.Unix(0, ticker.StatisticsCloseTime*int64(time.Millisecond)),
    FirstTradeId: ticker.FirstTradeId,
    LastTradeId: ticker.LastTradeId,
    TotalNumberOfTrades: ticker.TotalNumberOfTrades,
  }

  db := database.DB

  if err := db.Create(&ticketEntity).Error; err != nil {
    return err
  }

  return nil
}

func SaveCandlestick(candlestick BinanceCandlestick) error {
  candlestickEntity := entities.BinanceCandlestick{
    EventType: candlestick.EventType,
    EventTime: time.Unix(0, candlestick.EventTime*int64(time.Millisecond)),
    Symbol: candlestick.Symbol,
    KlineStartTime: time.Unix(0, candlestick.Kline.KlineStartTime*int64(time.Millisecond)),
    KlineCloseTime: time.Unix(0, candlestick.Kline.KlineCloseTime*int64(time.Millisecond)),
    Interval: candlestick.Kline.Interval,
    FirstTradeId: candlestick.Kline.FirstTradeId,
    LastTradeId: candlestick.Kline.LastTradeId,
    OpenPrice: candlestick.Kline.OpenPrice,
    ClosePrice: candlestick.Kline.ClosePrice,
    HighPrice: candlestick.Kline.HighPrice,
    LowPrice: candlestick.Kline.LowPrice,
    BaseAssetVolume: candlestick.Kline.BaseAssetVolume,
    NumberOfTrades: candlestick.Kline.NumberOfTrades,
    IsClosed: candlestick.Kline.IsClosed,
    QuoteAssetVolume: candlestick.Kline.QuoteAssetVolume,
    TakerBuyBaseAssetVolume: candlestick.Kline.TakerBuyBaseAssetVolume,
    TakerBuyQuoteAssetVolume: candlestick.Kline.TakerBuyQuoteAssetVolume,
  }

  db := database.DB

  if err := db.Create(&candlestickEntity).Error; err != nil {
    return err
  }

  return nil
}
