package repository

import (
	"context"
	"hive-data-collector/internal/domain"
)

type TrackMessageBrokerRepository interface {
	Created(ctx context.Context, track domain.Track) error
}
