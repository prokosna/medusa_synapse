package infra

import (
	"fmt"

	"encoding/json"

	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/prokosna/medusa_synapse/domain"
	"github.com/prokosna/medusa_synapse/exception"
)

type PublisherKafka struct {
	config   domain.Config
	producer *kafka.Producer
}

func NewPublisherKafka(config domain.Config) *PublisherKafka {
	// See. https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": strings.Join(config.Brokers, ","),
	})
	if err != nil {
		panic(err.Error())
	}
	return &PublisherKafka{
		config:   config,
		producer: p,
	}
}

func (p *PublisherKafka) Publish(key string, img domain.Image) error {
	message, err := json.Marshal(img)
	if err != nil {
		return exception.NewBadRequestError(err.Error())
	}
	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)
	err = p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &key, Partition: kafka.PartitionAny},
		Value:          message,
	}, deliveryChan)
	if err != nil {
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return fmt.Errorf("delivery message failed: %+v\n", m.TopicPartition.Error)
	}
	return nil
}
