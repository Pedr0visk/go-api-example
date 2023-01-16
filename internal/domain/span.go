package domain

type Url struct {
	Hostname string
	Pathname string
	Search   string
}

type Span struct {
	ID        string
	SessionID string
	PageID    string
	Date      int
	Url       Url
	UserAgent string
}

// /^0x[a-fA-F0-9]{40}$/g wallet address regex
