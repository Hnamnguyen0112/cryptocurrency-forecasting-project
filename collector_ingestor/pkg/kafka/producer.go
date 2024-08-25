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

func NewKafkaProducer(address string, interrupt chan os.Signal) *KafkaProducer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": address})
	if err != nil {
		log.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	return &KafkaProducer{
		Producer:  p,
		Interrupt: interrupt,
	}
}

// value is serialized to avro format
func (kp *KafkaProducer) Produce(topic string, value []byte) {
	err := kp.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)

	if err != nil {
		log.Printf("Failed to produce message: %v", err)
	}
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
