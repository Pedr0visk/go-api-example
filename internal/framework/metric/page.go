package metric

import (
	"analytics/internal"
	"analytics/internal/application/service"
	"context"
)

type PageMetric struct {
	svc service.PageService
}

func NewPageMetric(svc service.PageService) *PageMetric {
	return &PageMetric{svc: svc}
}

func (p *PageMetric) Index(ctx context.Context, pageID string, pageViews uint) error {
	err := p.svc.UpdatePageViews(ctx, pageID, pageViews)
	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "PageMetric.Index")
	}

	return nil
}
