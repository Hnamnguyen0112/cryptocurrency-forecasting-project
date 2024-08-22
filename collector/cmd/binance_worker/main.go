package main

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/config"
	binanceworker "github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/internal/binance_worker"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/pkg/database"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/pkg/entities"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/pkg/kafka"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/pkg/websocket"
)

const idleTimeout = 5 * time.Second

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	dbUser := config.Config("BINANCE_DB_USER")
	dbPassword := config.Config("BINANCE_DB_PASSWORD")
	dbHost := config.Config("BINANCE_DB_HOST")
	dbPort := config.Config("BINANCE_DB_PORT")
	dbName := config.Config("BINANCE_DB_NAME")

	connectParams := database.ConnectParams{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		Name:     dbName,
	}

	database.Connect(connectParams)

	database.DB.AutoMigrate(&entities.BinanceTicker{}, &entities.BinanceCandlestick{})

	u := url.URL{
		Scheme: "wss",
		Host:   "fstream.binance.com",
		Path:   "/ws",
	}
	log.Printf("Connecting to %s", u.String())

	ws := websocket.Connect(u.String(), interrupt)
	defer ws.Close()

	kafkaProducer := kafka.NewKafkaProducer(interrupt)
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

			var binanceType binanceworker.BinanceCommon
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
