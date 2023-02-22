package infrastructure

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/config"
)

func NewKafka(config config.Config) (*kafka.Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.Kafka.Address,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return producer, nil
}
