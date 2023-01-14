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

func (svc *PageService) RetrievePageMetadata(ctx context.Context, publisherID, url string) (domain.Page, error) {
	// defer
	var err error
	var page domain.Page

	page, err = svc.repo.FindByUrl(ctx, url)
	if err != nil {
		return page, err
	}

	if page.ID == "" {
		page, err = svc.repo.Create(ctx, publisherID, url)
		if err != nil {
			return page, err
		}
	}

	return page, nil
}

func NewPageService(repo repository.PageRepository, msgBroker repository.PageMessageBrokerRepository) *PageService {
	return &PageService{
		repo:      repo,
		msgBroker: msgBroker,
	}
}
