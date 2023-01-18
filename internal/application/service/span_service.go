package service

import (
	"analytics/internal/domain"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type SpanMessageBrokerRepository interface {
	Created(ctx context.Context, span domain.Span) error
}

type SpanService struct {
	msgBroker SpanMessageBrokerRepository
}

func NewSpanService(msgBroker SpanMessageBrokerRepository) *SpanService {
	return &SpanService{
		msgBroker: msgBroker,
	}
}

func (s *SpanService) Create(ctx context.Context, params SpanCreateParams) error {
	ID, _ := uuid.NewUUID()

	if err := s.msgBroker.Created(ctx, domain.Span{
		ID:        ID.String(),
		SessionID: params.SessionID,
		PageID:    params.PageID,
		Date:      params.Date,
		UserAgent: params.UserAgent,
		Url:       fmt.Sprintf("%s%s%s", params.Hostname, params.Pathname, params.Search),
	}); err != nil {
		return err
	}

	return nil
}
