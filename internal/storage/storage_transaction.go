package storage

import "julo-test/internal/model"

func (s *Store) CreateTransaction(transaction *model.Transaction) error {
	return s.txnRepository.SaveTransaction(transaction)
}

func (s *Store) FetchTransactionByRefID(refID string) (*model.Transaction, error) {
	return s.txnRepository.GetTransactionByRefID(refID)
}
