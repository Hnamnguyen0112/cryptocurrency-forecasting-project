package main

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/config"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/kafka"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/response"
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

	scr := kafka.NewSchemaRegistry(config.Config("SCHEMA_REGISTRY_URL"))

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
				log.Printf("Error in reading message: %v", err)
				return
			}

			var binanceType response.BinanceCommon
			err = json.Unmarshal(message, &binanceType)
			if err != nil {
				log.Printf("Error in unmarshalling message: %v", err)
				return
			}

			switch binanceType.EventType {
			case "24hrTicker":
				var binanceTicker response.BinanceTicker
				err = json.Unmarshal(message, &binanceTicker)
				if err != nil {
					log.Printf("Error in unmarshalling binance ticker: %v", err)
				}

				payload, err := scr.Ser.Serialize("binance_ticker", &binanceTicker)
				if err != nil {
					log.Printf("Error in serializing binance ticker: %v", err)
				}

				kafkaProducer.Produce("binance_ticker", payload)
			case "kline":
				var binanceKline response.BinanceCandlestick
				err = json.Unmarshal(message, &binanceKline)
				if err != nil {
					log.Printf("Error in unmarshalling binance candlestick: %v", err)
				}

				payload, err := scr.Ser.Serialize("binance_candlestick", &binanceKline)
				if err != nil {
					log.Printf("Error in serializing binance candlestick: %v", err)
				}

				kafkaProducer.Produce("binance_candlestick", payload)
			default:
			}
		}
	}()

	ws.HandleInterrupt()
	kafkaProducer.HandleInterrupt()
}
