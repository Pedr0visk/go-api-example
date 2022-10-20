package repository

import (
	"context"
	"hive-data-collector/internal/application/dto"
	"hive-data-collector/internal/domain"
)

type TrackRepository interface {
	Create(ctx context.Context, track dto.Track) (domain.Track, error)
}
