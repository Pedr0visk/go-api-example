package repository

import (
	"analytics/internal/domain"
	"context"
)

type PageRepository interface {
	Create(ctx context.Context, id, url string) (domain.Page, error)
	Find(ctx context.Context, id string) (domain.Page, error)
}
