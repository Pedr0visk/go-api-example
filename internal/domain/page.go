package domain

type Page struct {
	ID          string
	PublisherID string
	Url         string
	Metadata    *Metadata
	Publisher   *Publisher
}

type Metadata struct {
	Taxonomies []Taxonomy
}

type Report struct {
	ID        string
	Year      int
	Month     int
	Hour      int
	PageViews int
}
