package service

import (
	"go.uber.org/zap"
	"julo-test/internal/request"
	"julo-test/internal/response"
)

//go:generate mockery --name=CakeService --case underscore
type WalletService interface {
	CreateWallet(initWalletReq *request.InitWalletRequest) (response.Code, string, error)
}

type TransactionService interface {
}

type Config interface {
	Logger() *zap.Logger
}
