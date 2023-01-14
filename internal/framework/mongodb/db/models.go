package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Page struct {
	ID          primitive.ObjectID `bson:"_id"`
	Url         string             `bson:"url"`
	PublisherID string             `bson:"publisher_id"`
}

type PageMetrics struct {
	PageViews int
}
