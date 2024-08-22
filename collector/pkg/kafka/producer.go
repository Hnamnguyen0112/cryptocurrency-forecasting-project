package kafka

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaProducer struct {
	Producer  *kafka.Producer
	Interrupt chan os.Signal
}

func NewKafkaProducer(interrupt chan os.Signal) *KafkaProducer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:29092"})
	if err != nil {
		panic(err)
	}

	return &KafkaProducer{
		Producer:  p,
		Interrupt: interrupt,
	}
}

func (kp *KafkaProducer) Produce(topic string, message string) error {
	err := kp.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	if err != nil {
		return err
	}

	return nil
}

func (kp *KafkaProducer) HandleInterrupt() {
	for {
		select {
		case <-kp.Interrupt:
			log.Println("Interrupt received, kafka producer closing connection ...")
			kp.Producer.Flush(15 * 1000)
			return
		}
	}
}

func (kp *KafkaProducer) Close() {
	log.Println("Closing kafka producer connection ...")
	kp.Producer.Close()
	log.Println("Kafka producer connection closed")
}
