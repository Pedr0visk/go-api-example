package repository

import (
	"analytics/internal/domain"
	"context"
)

type PageReportSearchRepository interface {
	Search(ctx context.Context, pageID string, year, month uint) (domain.PageReport, error)
}
