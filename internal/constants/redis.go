package constants

const (
	UserWalletBalanceRedisKeyPrefix string = "usr-wlt-bln-"
)

func GetUserWalletBalanceRedisKey(custID string) string {
	return UserWalletBalanceRedisKeyPrefix + custID
}
