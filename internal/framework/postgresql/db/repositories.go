package db

import (
	"context"
	"hive-data-collector/internal/application/dto"
	"hive-data-collector/internal/domain"
)

type TraceRepository struct{}

func (r *TraceRepository) Create(ctx context.Context, trace dto.Trace) (domain.Trace, error) {
	// db := GetDB()
	return domain.Trace{
		UserWalletAddress: "0xABC123",
		Payload:           "payload",
		Date:              123712631283,
		Source:            "test-url",
	}, nil
}

func NewTraceRepository() *TraceRepository {
	return &TraceRepository{}
}
