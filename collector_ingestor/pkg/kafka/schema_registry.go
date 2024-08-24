package kafka

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/avro"
)

type SchemaRegistry struct {
	client schemaregistry.Client
	serde  *avro.SpecificSerializer
}

func NewSchemaRegistry(url string) *SchemaRegistry {
	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(url))
	if err != nil {
		fmt.Printf("Failed to create client: %s\n", err)
		os.Exit(1)
	}

	ser, err := avro.NewSpecificSerializer(client, serde.ValueSerde, avro.NewSerializerConfig())

	if err != nil {
		fmt.Printf("Failed to create serializer: %s\n", err)
		os.Exit(1)
	}

	return &SchemaRegistry{
		client: client,
		serde:  ser,
	}
}
