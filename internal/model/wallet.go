package model

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"julo-test/internal/dto"
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

func (u *Wallet) ToDTO() *dto.Wallet {

	return &dto.Wallet{
		ID:         u.ID,
		OwnedBy:    u.OwnedBy,
		Status:     u.Status,
		EnabledAt:  u.EnabledAt.Time,
		Balance:    u.Balance,
		DisabledAt: u.DisabledAt.Time,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}
