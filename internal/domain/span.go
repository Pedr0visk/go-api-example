package domain

type Span struct {
	ID                string `json:"id"`
	SessionID         string `json:"session_id"`
	UserWalletAddress string `json:"user_wallet_address"` // 0xb794f5ea0ba39494ce839613fffba74279579268
	Payload           string `json:"payload"`             // we can send any data here in json format
	Date              int64  `json:"date"`                // date when it occurred in milliseconds
	PageUrl           string `json:"page_url"`            // website's hostname+path
}

// /^0x[a-fA-F0-9]{40}$/g wallet address regex
