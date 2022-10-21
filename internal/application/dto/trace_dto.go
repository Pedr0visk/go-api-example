package dto

type Trace struct {
	UserWalletAddress string `json:"user_wallet_address"` // user wallet address
	Payload           string `json:"message"`             // we can send any data here in json format
	Date              int64  `json:"date"`                // date when it occurred in milliseconds
	PublisherUrl      string `json:"publisher_url"`
}

type TraceCreateParams struct {
	UserWalletAddress string `json:"user_wallet_address"` // user wallet address
	Payload           string `json:"message"`             // we can send any data here in json format
	Date              int64  `json:"date"`                // date when it occurred in milliseconds
	PublisherUrl      string `json:"publisher_url"`
}
