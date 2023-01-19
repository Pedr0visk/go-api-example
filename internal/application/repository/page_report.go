package repository

import (
	"analytics/internal/domain"
	"context"
)

type PageReportRepository interface {
	Create(ctx context.Context, pageID string) (domain.PageReport, error)
	Update(ctx context.Context, id string, pageViews uint) error
}
