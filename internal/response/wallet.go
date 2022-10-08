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

type DisableWalletDetail struct {
	ID         string    `json:"id"`
	OwnedBy    string    `json:"owned_by"`
	Status     string    `json:"status"`
	DisabledAt time.Time `json:"disabled_at"`
	Balance    int       `json:"balance"`
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
