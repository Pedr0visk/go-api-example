package db

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Trace struct {
	ID                uuid.UUID      `gorm:"column:id"`
	UserWalletAddress string         `gorm:"column:user_wallet_address"`
	Payload           string         `gorm:"column:payload"`
	Date              int64          `gorm:"column:date"`
	PublisherUrl      string         `gorm:"column:publisher_url"`
	CreatedAt         time.Time      `gorm:"column:created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at"`
}
