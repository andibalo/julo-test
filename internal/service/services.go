package service

import (
	"go.uber.org/zap"
	"julo-test/internal/model"
	"julo-test/internal/request"
	"julo-test/internal/response"
)

//go:generate mockery --name=CakeService --case underscore
type WalletService interface {
	CreateWallet(initWalletReq *request.InitWalletRequest) (response.Code, string, error)
	EnableWallet(custID string) (response.Code, *model.Wallet, error)
	FetchWalletBalance(custID string) (response.Code, *model.Wallet, error)
	DisableWallet(custID string) (response.Code, *model.Wallet, error)
	DepositWallet(amount int, custID, refID string) (response.Code, *model.Wallet, *model.Transaction, error)
	WithdrawFromWallet(amount int, custID, refID string) (response.Code, *model.Wallet, *model.Transaction, error)
}

type TransactionService interface {
}

type Config interface {
	Logger() *zap.Logger
}
