package service

import (
	"context"
	"hive-data-collector/internal/application/dto"
	"hive-data-collector/internal/application/repository"
	"hive-data-collector/internal/domain"
	"log"
)

type TraceService struct {
	repo      repository.TraceRepository
	msgBroker repository.TraceMessageBrokerRepository
}

func (s *TraceService) Create(ctx context.Context, params dto.TraceCreateParams) (domain.Trace, error) {
	// defer

	trace, err := s.repo.Create(ctx, dto.Trace{
		UserWalletAddress: params.UserWalletAddress,
		Payload:           params.Payload,
		PublisherUrl:      params.PublisherUrl,
		Date:              params.Date,
	})

	if err != nil {
		log.Fatalf("could not create trace %v", err)
	}

	err = s.msgBroker.Created(ctx, trace)
	if err != nil {
		// log error could not send kafka
		return domain.Trace{}, nil
	}

	return trace, nil
}

func NewTraceService(repo repository.TraceRepository, msgBroker repository.TraceMessageBrokerRepository) *TraceService {
	return &TraceService{
		repo:      repo,
		msgBroker: msgBroker,
	}
}
