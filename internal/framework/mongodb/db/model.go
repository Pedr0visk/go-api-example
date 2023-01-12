package db

type Page struct {
	Url string `bson:"url"`
}

type PageMetrics struct {
	PageViews int
}
