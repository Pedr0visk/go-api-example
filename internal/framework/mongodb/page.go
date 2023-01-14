package mongodb

import (
	"analytics/internal/domain"
	"analytics/internal/framework/mongodb/db"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PageRepository struct {
	Db *mongo.Database
}

func NewPageRepository(database *mongo.Database) *PageRepository {
	return &PageRepository{
		Db: database,
	}
}

var collectionName = "pages"

func (p *PageRepository) Create(ctx context.Context, url, publisherID string) (domain.Page, error) {
	var page domain.Page

	coll := p.Db.Collection(collectionName)

	doc := db.Page{
		Url:         url,
		PublisherID: publisherID,
	}

	result, err := coll.InsertOne(ctx, doc)
	if err != nil {
		return domain.Page{}, err
	}

	page.ID = fmt.Sprintf("%v", result.InsertedID)
	page.PublisherID = publisherID
	page.Url = url

	return page, nil
}

func (p *PageRepository) Find(ctx context.Context, id string) (domain.Page, error) {
	var page domain.Page

	coll := p.Db.Collection(collectionName)

	err := coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&page)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Page{}, nil
		}
		return domain.Page{}, err
	}

	return page, nil
}
