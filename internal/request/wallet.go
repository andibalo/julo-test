package request

type InitWalletRequest struct {
	CustomerXID string `json:"customer_xid" validate:"required"`
}

type DepositWalletRequest struct {
	Amount int    `json:"amount" validate:"required"`
	RefID  string `json:"reference_id" validate:"required"`
}
