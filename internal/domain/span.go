package domain

type Url struct {
	Hostname string `json:"hostname"`
	Pathname string `json:"pathname"`
	Search   string `json:"search"`
}

type Span struct {
	ID        string `json:"id"`
	SessionID string `json:"session_id"`
	PageID    string `json:"page_id"`
	Date      int64  `json:"date"`
	Url       string `json:"url"`
	UserAgent string `json:"user_agent"`
}

// /^0x[a-fA-F0-9]{40}$/g wallet address regex
