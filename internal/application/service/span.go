package service

type SpanCreateParams struct {
	PageID    string
	SessionID string
	Date      int64
	UserAgent string
	Hostname  string
	Pathname  string
	Search    string
}
