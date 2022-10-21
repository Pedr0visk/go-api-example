package db

import "time"

type Trace struct {
	ID                string    `gorm:"column:id"`
	UserWalletAddress string    `gorm:"column:user_wallet_address"`
	Payload           string    `gorm:"column:payload"`
	Date              int64     `gorm:"column:date"`
	Source            string    `gorm:"column:source"`
	CreatedAt         time.Time `gorm:"column:created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at"`
	DeletedAt         time.Time `gorm:"column:deleted_at"`
}
