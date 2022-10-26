package domain

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

type Trace struct {
	ID                string    `json:"id"`
	UserWalletAddress string    `json:"user_wallet_address"` // 0xb794f5ea0ba39494ce839613fffba74279579268
	Payload           string    `json:"payload"`             // we can send any data here in json format
	Date              int64     `json:"date"`                // date when it occurred in milliseconds
	PublisherUrl      string    `json:"publisher_url"`       // website's hostname+path
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         time.Time `json:"deleted_at"`
}

// Validate ...
func (t Trace) Validate() error {
	if err := validation.ValidateStruct(&t,
		validation.Field(&t.UserWalletAddress, validation.Required),
		validation.Field(&t.Payload, validation.Required),
		validation.Field(&t.Date, validation.Required),
		validation.Field(&t.PublisherUrl),
	); err != nil {
		return WrapErrorf(err, ErrorCodeInvalidArgument, "invalid values")
	}

	return nil
}

// /^0x[a-fA-F0-9]{40}$/g wallet address regex
