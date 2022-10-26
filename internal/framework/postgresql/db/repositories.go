package db

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"hive-data-collector/internal/domain"
)

type TraceRepository struct{}

func (r *TraceRepository) Create(ctx context.Context, trace domain.Trace) (domain.Trace, error) {
	db := GetDB()

	obj := Trace{
		ID:                uuid.NewV4(),
		PublisherUrl:      trace.PublisherUrl,
		Date:              trace.Date,
		UserWalletAddress: trace.UserWalletAddress,
		Payload:           trace.Payload,
	}

	if err := db.Create(&obj).Error; err != nil {
		return domain.Trace{}, domain.WrapErrorf(err, domain.ErrorCodeUnknown, "insert trace")
	}

	trace.ID = obj.ID.String()

	return trace, nil
}

func NewTraceRepository() *TraceRepository {
	return &TraceRepository{}
}
