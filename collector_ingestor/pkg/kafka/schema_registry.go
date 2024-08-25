package kafka

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/jsonschema"
)

type SchemaRegistry struct {
	Client schemaregistry.Client
	Serde  *jsonschema.Serializer
}

func NewSchemaRegistry(url string) *SchemaRegistry {
	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(url))
	if err != nil {
		log.Printf("Error creating schema registry client: %v\n", err)
		os.Exit(1)
	}

	ser, err := jsonschema.NewSerializer(client, serde.ValueSerde, jsonschema.NewSerializerConfig())

	if err != nil {
		log.Printf("Error creating schema registry serializer: %v\n", err)
		os.Exit(1)
	}

	return &SchemaRegistry{
		Client: client,
		Serde:  ser,
	}
}
