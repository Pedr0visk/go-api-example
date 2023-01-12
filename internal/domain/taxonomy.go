package domain

type Category struct {
	ID         int    `json:"id"`
	TaxonomyID int    `json:"taxonomy_id"`
	Label      string `json:"label"`
}

type Taxonomy struct {
	ID         int        `json:"id"`
	Label      string     `json:"label"`
	Categories []Category `json:"categories"`
}
