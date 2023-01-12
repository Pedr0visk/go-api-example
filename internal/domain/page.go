package domain

type Page struct {
	ID        string     `json:"id"`
	Url       string     `json:"url"`
	Metadata  *Metadata  `json:"metadata"`
	Publisher *Publisher `json:"publisher_id"`
}

type Metadata struct {
	Taxonomies []Taxonomy `json:"taxonomies"`
}

type Report struct {
	ID        string `json:"id"`
	Year      int    `json:"year"`       // 2022,2023...
	Month     int    `json:"month"`      // 1,2,3,4...
	Hour      int    `json:"hour"`       // 1,2,3...23,0
	PageViews int    `json:"page_views"` // 0
}
