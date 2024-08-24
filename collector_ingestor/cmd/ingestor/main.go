package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/config"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/database"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/entities"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/kafka"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	dbUser := config.Config("DB_USER")
	dbPassword := config.Config("DB_PASSWORD")
	dbHost := config.Config("DB_HOST")
	dbPort := config.Config("DB_PORT")
	dbName := config.Config("DB_NAME")

	connectParams := database.ConnectParams{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		Name:     dbName,
	}

	database.Connect(connectParams)

	database.DB.AutoMigrate(
		&entities.BinanceTicker{},
		&entities.BinanceCandlestick{},
	)

	kafkaConsumer := kafka.NewKafkaConsumer(
		config.Config("KAFKA_BOOTSTRAP_SERVERS"),
		"ingestor_group",
		[]string{"binance_ticker", "binance_candlestick", "coinbase_ticker", "coinbase_candles"},
		interrupt,
	)
	defer kafkaConsumer.Close()

	go kafkaConsumer.Consume()

	kafkaConsumer.HandleInterrupt()
}
