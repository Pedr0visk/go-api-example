package service

type SpanCreateParams struct {
	PageID    string
	SessionID string
	Date      int
	UserAgent string
	Hostname  string
	Pathname  string
	Search    string
}
