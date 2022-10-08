package response

import (
	"database/sql"
)

type InitWalletResponse struct {
	Token string `json:"token"`
}

type EnableWalletResponse struct {
	Wallet DefaultWalletInfo `json:"token"`
}

type DefaultWalletInfo struct {
	ID        string       `json:"id"`
	OwnedBy   string       `json:"owned_by"`
	Status    string       `json:"status"`
	EnabledAt sql.NullTime `json:"enabled_at"`
	Balance   int          `json:"balance"`
}
