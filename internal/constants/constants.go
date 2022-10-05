package constants

const (
	HeaderRequestID string = "X-Request-ID"
)

const (
	V1BasePath         = "/api/v1"
	InitBasePath       = V1BasePath + "/init"
	WalletBasePath     = V1BasePath + "/wallet"
	DepositBasePath    = WalletBasePath + "/deposits"
	WithdrawalBasePath = WalletBasePath + "/withdrawals"
)
