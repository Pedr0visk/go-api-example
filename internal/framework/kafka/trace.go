package kafka

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"hive-data-collector/internal/domain"
)

type TraceMessageBroker struct {
	producer  *kafka.Producer
	topicName string
}

type event struct {
	Type  string
	Value domain.Trace
}

func (r *TraceMessageBroker) Created(ctx context.Context, trace domain.Trace) error {
	return nil
}

func NewTraceMessageBroker() *TraceMessageBroker {
	return &TraceMessageBroker{}
}

func (t *TraceMessageBroker) publish(ctx context.Context, spanName, msgType string, trace domain.Trace) error {
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
		Value: trace,
	}

	if err := json.NewEncoder(&b).Encode(evt); err != nil {
		return domain.WrapErrorf(err, domain.ErrorCodeUnknown, "json.Encode")
	}

	if err := t.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &t.topicName,
			Partition: kafka.PartitionAny,
		},
		Value: b.Bytes(),
	}, nil); err != nil {
		return domain.WrapErrorf(err, domain.ErrorCodeUnknown, "product.Producer")
	}

	return nil
}
