package service

import (
	"julo-test/internal/storage"
)

type walletService struct {
	config Config
	store  storage.Storage
}

func NewWalletService(config Config, store storage.Storage) *walletService {

	return &walletService{
		config: config,
		store:  store,
	}
}
