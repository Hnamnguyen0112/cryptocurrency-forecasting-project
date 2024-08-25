package kafka

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaConsumer struct {
	Consumer  *kafka.Consumer
	Interrupt chan os.Signal
}

func NewKafkaConsumer(
	address string,
	groupID string,
	topics []string,
	interrupt chan os.Signal,
) *KafkaConsumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": address,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Printf("Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	c.SubscribeTopics(topics, nil)

	return &KafkaConsumer{
		Consumer:  c,
		Interrupt: interrupt,
	}
}

func (kc *KafkaConsumer) Consume() {
	for {
		ev := kc.Consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			log.Printf("Message on %s: %s\n", e.TopicPartition, string(e.Value))
		case kafka.Error:
			log.Printf("Error: %v\n", e)
		}
	}
}

func (kc *KafkaConsumer) HandleInterrupt() {
	for {
		select {
		case <-kc.Interrupt:
			log.Println("Interrupt received, kafka consumer closing connection ...")
			kc.Consumer.Close()
			return
		}
	}
}

func (kc *KafkaConsumer) Close() {
	log.Println("Closing kafka consumer connection ...")
	kc.Consumer.Close()
	log.Println("Kafka consumer connection closed")
}
