package response

import (
	"time"
)

type InitWalletResponse struct {
	Token string `json:"token"`
}

type EnableWalletResponse struct {
	Wallet DefaultWalletInfo `json:"wallet"`
}

type FetchWalletBalanceResponse struct {
	Wallet DefaultWalletInfo `json:"wallet"`
}

type DisableWalletResponse struct {
	Wallet DisableWalletDetail `json:"wallet"`
}

type DepositWalletResponse struct {
	Deposit DepositWalletDetail `json:"deposit"`
}

type WithdrawWalletResponse struct {
	Withdrawal WithdrawWalletDetail `json:"withdrawal"`
}

type DisableWalletDetail struct {
	ID         string    `json:"id"`
	OwnedBy    string    `json:"owned_by"`
	Status     string    `json:"status"`
	DisabledAt time.Time `json:"disabled_at"`
	Balance    int       `json:"balance"`
}

type WithdrawWalletDetail struct {
	ID          string    `json:"id"`
	WithdrawnBy string    `json:"withdrawn_by"`
	Status      string    `json:"status"`
	WithdrawnAt time.Time `json:"withdrawn_at"`
	Amount      int       `json:"amount"`
	RefID       string    `json:"reference_id"`
}

type DepositWalletDetail struct {
	ID          string    `json:"id"`
	DepositedBy string    `json:"deposited_by"`
	Status      string    `json:"status"`
	DepositedAt time.Time `json:"deposited_at"`
	Amount      int       `json:"amount"`
	RefID       string    `json:"reference_id"`
}

type DefaultWalletInfo struct {
	ID        string    `json:"id"`
	OwnedBy   string    `json:"owned_by"`
	Status    string    `json:"status"`
	EnabledAt time.Time `json:"enabled_at"`
	Balance   int       `json:"balance"`
}
