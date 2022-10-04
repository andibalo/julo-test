package service

import (
	"go.uber.org/zap"
)

//go:generate mockery --name=CakeService --case underscore
type WalletService interface {
}

type TransactionService interface {
}

type Config interface {
	Logger() *zap.Logger
}
