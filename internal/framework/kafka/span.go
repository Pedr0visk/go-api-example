package kafka

import (
	"analytics/internal"
	"analytics/internal/domain"
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type SpanMessageBroker struct {
	producer  *kafka.Producer
	topicName string
}

type event struct {
	Type  string
	Value domain.Span
}

func NewSpanMessageBroker(producer *kafka.Producer, topicName string) *SpanMessageBroker {
	return &SpanMessageBroker{
		producer:  producer,
		topicName: topicName,
	}
}

func (s *SpanMessageBroker) Created(ctx context.Context, span domain.Span) error {
	return s.publish(ctx, "Span.Created", "spans.event.created", span)
}

func (s *SpanMessageBroker) publish(ctx context.Context, spanName, msgType string, span domain.Span) error {
	fmt.Println(spanName)
	// monitoring
	//_, span := otel.Tracer(otelName).Start(ctx, spanName)
	//defer span.End()

	//span.SetAttributes(
	//	attribute.KeyValue{
	//		Key:   semconv.MessagingSystemKey,
	//		Value: attribute.StringValue("kafka"),
	//	},
	//	attribute.KeyValue{
	//		Key:   semconv.MessagingDestinationKey,
	//		Value: attribute.StringValue(t.topicName),
	//	},
	//)

	//-

	var b bytes.Buffer

	evt := event{
		Type:  msgType,
		Value: span,
	}

	if err := json.NewEncoder(&b).Encode(evt); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "json.Encode")
	}

	if err := s.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &s.topicName,
			Partition: kafka.PartitionAny,
		},
		Value: b.Bytes(),
		Key:   []byte(evt.Value.PageID),
	}, nil); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "span.Producer")
	}

	return nil
}
