package service

import (
	"analytics/internal/domain"
	"context"
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

	if err := s.msgBroker.Created(ctx, domain.Span{
		ID:        "1",
		SessionID: params.SessionID,
		PageID:    params.PageID,
		Date:      int(params.Date),
		Url: domain.Url{
			Hostname: params.Hostname,
			Pathname: params.Pathname,
			Search:   params.Search,
		},
	}); err != nil {
		return err
	}

	return nil
}
