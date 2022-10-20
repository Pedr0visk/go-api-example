package service

import (
	"context"
	"hive-data-collector/internal/application/dto"
	"hive-data-collector/internal/application/repository"
	"log"
)

type TrackService struct {
	repo      repository.TrackRepository
	msgBroker repository.TrackMessageBrokerRepository
}

func (s *TrackService) Create(ctx context.Context, userWalletAddress, source, payload string, date int64) error {
	// defer
	track, err := s.repo.Create(ctx, dto.Track{
		UserWalletAddress: userWalletAddress,
		Payload:           payload,
		Source:            source,
		Date:              date,
	})

	if err != nil {
		log.Fatalf("could not create track %v", err)
	}

	err = s.msgBroker.Created(ctx, track)
	if err != nil {
		// log error could not send kafka
		return nil
	}

	return nil
}

func NewTrackService(repo repository.TrackRepository, msgBroker repository.TrackMessageBrokerRepository) *TrackService {
	return &TrackService{
		repo:      repo,
		msgBroker: msgBroker,
	}
}
