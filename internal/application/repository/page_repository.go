package repository

import (
	"analytics/internal/domain"
	"context"
)

type PageRepository interface {
	GetByUrl(ctx context.Context, url string) domain.Page
	Create(ctx context.Context, page *domain.Page) error
}
