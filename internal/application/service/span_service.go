package service

import (
	"analytics/internal/domain"
	"context"
)

type SpanMessageBroker interface {
	Created(ctx context.Context, span domain.Span) error
}

type SpanService struct {
	msgBroker SpanMessageBroker
}

type SpanCreate struct {
	PageUrl   string `json:"page_url"`
	Agent     string `json:"agent"`
	Date      int    `json:"date"`
	SessionID string `json:"session_id"`
}

func (s *SpanService) Create(ctx context.Context, inputParams SpanCreate) error {

	if err := s.msgBroker.Created(ctx, domain.Span{}); err != nil {
		return err
	}

	return nil
}
