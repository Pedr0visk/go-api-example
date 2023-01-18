package domain

type Session struct {
	ID       string `json:"id"` // cookieID
	Metadata string
}
