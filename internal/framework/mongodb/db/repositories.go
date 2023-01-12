package db

import (
	"analytics/internal/domain"
	"context"
	"fmt"
)

type PageRepository struct{}

var collectionName = "pages"

func (repo *PageRepository) Create(ctx context.Context, page *domain.Page) error {
	db := GetDB()

	doc := Page{
		Url: page.Url,
	}

	coll := db.Collection(collectionName)
	result, err := coll.InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	page.ID = fmt.Sprintf("%v", result.InsertedID)

	return nil
}

func (repo *PageRepository) GetByUrl(ctx context.Context, url string) domain.Page {
	return domain.Page{}
}

func NewPageRepository() *PageRepository {
	return &PageRepository{}
}
