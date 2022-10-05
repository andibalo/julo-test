package model

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Wallet struct {
	ID         string       `json:"id" gorm:"primaryKey"`
	OwnedBy    string       `json:"owned_by" gorm:"uniqueIndex;not null;type:varchar(64)"`
	Status     string       `json:"status" gorm:"not null;type:varchar(64)"`
	EnabledAt  sql.NullTime `json:"enabled_at"`
	DisabledAt sql.NullTime `json:"disabled_at"`
	Balance    int          `json:"balance"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (u *Wallet) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	return nil
}