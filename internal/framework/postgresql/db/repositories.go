package db

import "hive-data-collector/internal/application/dto"

type TrackRepository struct{}

func (r *TrackRepository) Create(track dto.Track) error {
	// db := GetDB()
	return nil
}

func NewTrackRepository() *TrackRepository {
	return &TrackRepository{}
}
