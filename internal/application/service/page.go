package service

import (
	"analytics/internal"
	"analytics/internal/application/repository"
	"context"
)

type PageService struct {
	pageRepo             repository.PageRepository
	pageReportRepo       repository.PageReportRepository
	pageReportSearchRepo repository.PageReportSearchRepository
}

func NewPageService(pageRepo repository.PageRepository, pageReportRepo repository.PageReportRepository) *PageService {
	return &PageService{
		pageRepo:       pageRepo,
		pageReportRepo: pageReportRepo,
	}
}

func (p *PageService) UpdatePageViews(ctx context.Context, id string, pageViews uint) error {
	page, err := p.pageRepo.Find(ctx, id)
	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "pageRepo.Find")
	}

	// get month and year
	pageReport, err := p.pageReportSearchRepo.Search(ctx, page.ID, 2023, 1)
	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "pageReportSearchRepo.Search")
	}

	err = p.pageReportRepo.Update(ctx, pageReport.ID, pageViews)
	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "pageReportRepo.Update")
	}

	return nil
}
