package dto

import (
	"time"
)

type Transaction struct {
	ID          string    `json:"id"`
	TxnType     string    `json:"txn_type"`
	DepositedBy string    `json:"deposited_by"`
	DepositedAt time.Time `json:"deposited_at"`
	WithdrawnBy string    `json:"withdrawn_by"`
	WithdrawnAt time.Time `json:"withdrawn_at"`
	Status      string    `json:"status"`
	Amount      int       `json:"amount"`
	ReferenceId string    `json:"reference_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
