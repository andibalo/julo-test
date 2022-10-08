package model

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID          string       `json:"id" gorm:"primaryKey"`
	TxnType     string       `json:"txn_type" gorm:"not null;type:varchar(20)"`
	DepositedBy string       `json:"deposited_by"`
	DepositedAt sql.NullTime `json:"deposited_at"`
	WithdrawnBy string       `json:"withdrawn_by"`
	WithdrawnAt sql.NullTime `json:"withdrawn_at"`
	Status      string       `json:"status"`
	Amount      int          `json:"amount"`
	ReferenceId string       `json:"reference_id" gorm:"uniqueIndex;type:varchar(255)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u *Transaction) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	return nil
}
