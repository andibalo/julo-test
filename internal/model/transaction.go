package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	TxnType     string    `json:"txn_type" gorm:"primaryKey"`
	DepositedBy string    `json:"deposited_by"`
	DepositedAt time.Time `json:"deposited_at"`
	OwnedBy     string    `json:"owned_by" gorm:"uniqueIndex;not null;type:varchar(64)"`
	Status      string    `json:"status"`
	EnabledAt   time.Time `json:"enabled_at"`
	WithdrawnBy string    `json:"withdrawn_by"`
	WithdrawnAt time.Time `json:"withdrawn_at"`
	Amount      int       `json:"amount"`
	ReferenceId string    `json:"reference_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u *Transaction) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	return nil
}
