package main

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/config"
	binancecollector "github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/internal/binance_collector"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/kafka"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	u := url.URL{
		Scheme: "wss",
		Host:   "fstream.binance.com",
		Path:   "/ws",
	}
	log.Printf("Connecting to %s", u.String())

	ws := websocket.Connect(u.String(), interrupt)
	defer ws.Close()

	kafkaProducer := kafka.NewKafkaProducer(config.Config("KAFKA_BOOTSTRAP_SERVERS"), interrupt)
	defer kafkaProducer.Close()

	request := map[string]interface{}{
		"method": "SUBSCRIBE",
		"params": []string{"btcusdt@ticker", "btcusdt@kline_1m"},
		"id":     1,
	}

	ws.Subscribe(request)

	go func() {
		defer close(ws.Done)
		for {
			_, message, err := ws.Conn.ReadMessage()
			if err != nil {
				return
			}

			var binanceType binancecollector.BinanceCommon
			err = json.Unmarshal(message, &binanceType)
			if err != nil {
				return
			}

			switch binanceType.EventType {
			case "24hrTicker":
				kafkaProducer.Produce("binance_ticker", string(message))
			case "kline":
				kafkaProducer.Produce("binance_candlestick", string(message))
			default:
			}
		}
	}()

	ws.HandleInterrupt()
	kafkaProducer.HandleInterrupt()
}
