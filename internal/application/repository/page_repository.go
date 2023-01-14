package repository

import (
	"analytics/internal/domain"
	"context"
)

type PageRepository interface {
	FindByUrl(ctx context.Context, url string) (domain.Page, error)
	Find(ctx context.Context, id string) (domain.Page, error)
	Create(ctx context.Context, url, publisherID string) (domain.Page, error)
	Update(ctx context.Context, id, url string) error
}
