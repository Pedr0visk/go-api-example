package repository

import (
	"context"
	"hive-data-collector/internal/domain"
)

type TraceMessageBrokerRepository interface {
	Created(ctx context.Context, trace domain.Trace) error
}
