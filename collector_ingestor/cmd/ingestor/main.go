package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/config"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/internal/ingestor"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/database"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/entities"
	kafkaPkg "github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"gorm.io/driver/postgres"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	dbUser := config.Config("DB_USER")
	dbPassword := config.Config("DB_PASSWORD")
	dbHost := config.Config("DB_HOST")
	dbPort := config.Config("DB_PORT")
	dbName := config.Config("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)

	dialector := postgres.New(postgres.Config{
		DSN: dsn,
	})

	DB, err := database.Connect(dialector)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		os.Exit(1)
	}

	DB.AutoMigrate(
		&entities.BinanceTicker{},
		&entities.BinanceCandlestick{},
		&entities.CoinbaseTicker{},
		&entities.CoinbaseCandle{},
	)

	kafkaConsumer := kafkaPkg.NewKafkaConsumer(
		config.Config("KAFKA_BOOTSTRAP_SERVERS"),
		"ingestor_group",
		[]string{"binance_ticker", "binance_candlestick", "coinbase_ticker", "coinbase_candles"},
		interrupt,
	)

	scr := kafkaPkg.NewSchemaRegistry(config.Config("SCHEMA_REGISTRY_URL"))

	log.Println("Starting consuming messages from Kafka")

	run := true

	for run {
		select {
		case sig := <-interrupt:
			log.Printf("Received signal: %v\n", sig)
			run = false
		default:
			ev := kafkaConsumer.Consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				go func(e *kafka.Message) {
					ingestor.HandleMessage(e, scr)
				}(e)
			case kafka.Error:
				log.Printf("Error: %v\n", e)
			default:
				log.Printf("Ignored: %v\n", e)
			}
		}
	}

	kafkaConsumer.Close()
}
