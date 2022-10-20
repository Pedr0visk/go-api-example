package dto

type Track struct {
	UserWalletAddress string `json:"user_wallet_address"` // user wallet address
	Payload           string `json:"message"`             // we can send any data here in json format
	Date              int64  `json:"date"`                // date when it occurred in milliseconds
	Source            string `json:"source"`
}
