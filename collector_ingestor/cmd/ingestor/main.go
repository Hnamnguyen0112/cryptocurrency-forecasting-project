package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/config"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/database"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/entities"
	kafkaPkg "github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/kafka"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector_ingestor/pkg/response"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
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
				switch *e.TopicPartition.Topic {
				case "binance_ticker":
					value := response.BinanceTicker{}
					err := scr.Deser.DeserializeInto(*e.TopicPartition.Topic, e.Value, &value)
					if err != nil {
						log.Printf("Error deserializing binance ticker: %v\n", err)
					}
				case "binance_candlestick":
					value := response.BinanceCandlestick{}
					err := scr.Deser.DeserializeInto(*e.TopicPartition.Topic, e.Value, &value)
					if err != nil {
						log.Printf("Error deserializing binance candlestick: %v\n", err)
					}
				case "coinbase_ticker":
					value := response.CoinbaseTicker{}
					err := scr.Deser.DeserializeInto(*e.TopicPartition.Topic, e.Value, &value)
					if err != nil {
						log.Printf("Error deserializing coinbase ticker: %v\n", err)
					}
				case "coinbase_candles":
					value := response.CoinbaseCandles{}
					err := scr.Deser.DeserializeInto(*e.TopicPartition.Topic, e.Value, &value)
					if err != nil {
						log.Printf("Error deserializing coinbase candles: %v\n", err)
					}
				default:
				}
			case kafka.Error:
				log.Printf("Error: %v\n", e)
			default:
				log.Printf("Ignored: %v\n", e)
			}
		}
	}

	kafkaConsumer.Close()
}
