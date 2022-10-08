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

type DefaultWalletInfo struct {
	ID        string    `json:"id"`
	OwnedBy   string    `json:"owned_by"`
	Status    string    `json:"status"`
	EnabledAt time.Time `json:"enabled_at"`
	Balance   int       `json:"balance"`
}
