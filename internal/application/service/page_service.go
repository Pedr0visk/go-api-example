package service

import (
	"analytics/internal/application/repository"
	"analytics/internal/domain"
	"context"
)

type PageService struct {
	repo      repository.PageRepository
	msgBroker repository.PageMessageBrokerRepository
}

type RetrievePageMetadataOutput struct {
	PageID   string          `json:"page_id"`
	Metadata domain.Metadata `json:"metadata"`
}

func (svc *PageService) RetrievePageMetadata(ctx context.Context, url string) (domain.Page, error) {
	// defer
	var err error
	var page domain.Page

	page = svc.repo.GetByUrl(ctx, url)
	if page.ID == "" {
		err = svc.repo.Create(ctx, &domain.Page{Url: page.Url})
		if err != nil {
			return page, err
		}
	}

	return page, nil
}

func NewPageService(repo repository.PageRepository, msgBroker repository.TraceMessageBrokerRepository) *PageService {
	return &PageService{
		repo:      repo,
		msgBroker: msgBroker,
	}
}
