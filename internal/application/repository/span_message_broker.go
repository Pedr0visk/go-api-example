package repository

import (
	"analytics/internal/domain"
	"context"
)

type SpanMessageBrokerRepository interface {
	Created(ctx context.Context, span domain.Span) error
}
