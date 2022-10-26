package repository

import (
	"context"
	"hive-data-collector/internal/domain"
)

type TraceRepository interface {
	Create(ctx context.Context, trace domain.Trace) (domain.Trace, error)
}
