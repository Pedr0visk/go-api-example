package repository

import (
	"analytics/internal/domain"
	"context"
)

type SpanRepository interface {
	Create(ctx context.Context, span domain.Span) (domain.Span, error)
}
