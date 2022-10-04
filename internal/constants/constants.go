package constants

const (
	HeaderRequestID string = "X-Request-ID"
)

const (
	V1BasePath         = "/v1"
	WalletBasePath     = V1BasePath + "/wallet"
	DepositBasePath    = WalletBasePath + "/deposits"
	WithdrawalBasePath = WalletBasePath + "/withdrawals"
)
