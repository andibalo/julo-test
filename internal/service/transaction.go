package service

import (
	"julo-test/internal/storage"
)

type transactionService struct {
	config Config
	store  storage.Storage
}

func NewTransactionService(config Config, store storage.Storage) *transactionService {

	return &transactionService{
		config: config,
		store:  store,
	}
}
