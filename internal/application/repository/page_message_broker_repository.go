package repository

import (
	"analytics/internal/domain"
	"context"
)

type PageMessageBrokerRepository interface {
	Created(ctx context.Context, page domain.Page) error
}
