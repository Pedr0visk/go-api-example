package kafka

import (
	"analytics/internal/domain"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Page struct {
	producer  *kafka.Producer
	topicName string
}

type PageEvent struct {
	Type  string
	Value domain.Page
}
