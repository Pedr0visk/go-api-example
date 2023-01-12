package domain

import "time"

type Session struct {
	ID         string    `json:"id"` // cookieID
	UserID     string    `json:"-"`
	LastAccess time.Time `json:"last_access"`
}
