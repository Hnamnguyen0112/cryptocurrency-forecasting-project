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
				log.Printf("Error in reading message: %v", err)
				return
			}

			var coinbaseType response.CoinbaseCommon
			err = json.Unmarshal(message, &coinbaseType)
			if err != nil {
				log.Printf("Error in unmarshalling message: %v", err)
				return
			}

			switch coinbaseType.Channel {
			case "ticker":
				var coinbaseTicker response.CoinbaseCommon
				err = json.Unmarshal(message, &coinbaseTicker)
				if err != nil {
					log.Printf("Error in unmarshalling coinbase ticker: %v", err)
				}

				payload, err := scr.Ser.Serialize("coinbase_ticker", &coinbaseTicker)
				if err != nil {
					log.Printf("Error in serializing coinbase ticker: %v", err)
				}

				kafkaProducer.Produce("coinbase_ticker", payload)
			case "candles":
				var coinbaseCandles response.CoinbaseCommon
				err = json.Unmarshal(message, &coinbaseCandles)
				if err != nil {
					log.Printf("Error in unmarshalling coinbase candles: %v", err)
				}

				payload, err := scr.Ser.Serialize("coinbase_candles", &coinbaseCandles)
				if err != nil {
					log.Printf("Error in serializing coinbase candles: %v", err)
				}

				kafkaProducer.Produce("coinbase_candles", payload)
			default:
			}
		}
	}()

	ws.HandleInterrupt()
	kafkaProducer.HandleInterrupt()
}
