package main

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/config"
	coinbasecollector "github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/internal/coinbase_collector"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/kafka"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	u := url.URL{
		Scheme: "wss",
		Host:   "advanced-trade-ws.coinbase.com",
		Path:   "/",
	}
	log.Printf("Connecting to %s", u.String())

	ws := websocket.Connect(u.String(), interrupt)
	defer ws.Close()

	kafkaProducer := kafka.NewKafkaProducer(config.Config("KAFKA_BOOTSTRAP_SERVERS"), interrupt)
	defer kafkaProducer.Close()

	scr := kafka.NewSchemaRegistry(config.Config("SCHEMA_REGISTRY_URL"))

	tickerRequest := map[string]interface{}{
		"type":        "subscribe",
		"product_ids": []string{"BTC-USDT"},
		"channel":     "ticker",
	}
	ws.Subscribe(tickerRequest)

	candlesRequest := map[string]interface{}{
		"type":        "subscribe",
		"product_ids": []string{"BTC-USDT"},
		"channel":     "candles",
	}
	ws.Subscribe(candlesRequest)

	go func() {
		defer close(ws.Done)
		for {
			_, message, err := ws.Conn.ReadMessage()
			if err != nil {
				return
			}

			var coinbaseType coinbasecollector.CoinbaseCommon
			err = json.Unmarshal(message, &coinbaseType)
			if err != nil {
				return
			}

			switch coinbaseType.Channel {
			case "ticker":
				var coinbaseTicker coinbasecollector.CoinbaseTicker
				err = json.Unmarshal(message, &coinbaseTicker)
				if err != nil {
					return
				}

				payload, err := scr.Serde.Serialize("coinbase_ticker", &coinbaseTicker)
				if err != nil {
					log.Printf("Error in serializing coinbase ticker: %v", err)
					return
				}

				kafkaProducer.Produce("coinbase_ticker", payload)
			case "candles":
				var coinbaseCandles coinbasecollector.CoinbaseCandles
				err = json.Unmarshal(message, &coinbaseCandles)
				if err != nil {
					return
				}

				payload, err := scr.Serde.Serialize("coinbase_candles", &coinbaseCandles)
				if err != nil {
					log.Printf("Error in serializing coinbase candles: %v", err)
					return
				}

				kafkaProducer.Produce("coinbase_candles", payload)
			default:
			}
		}
	}()

	ws.HandleInterrupt()
	kafkaProducer.HandleInterrupt()
}
