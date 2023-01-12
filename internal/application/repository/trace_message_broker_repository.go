package repository

import (
	"analytics/internal/domain"
	"context"
)

type TraceMessageBrokerRepository interface {
	Created(ctx context.Context, trace domain.Span) error
}
