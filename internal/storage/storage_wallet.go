package storage

import "julo-test/internal/model"

func (s *Store) CreateWallet(wallet *model.Wallet) error {
	return s.walletRepository.SaveWallet(wallet)
}

func (s *Store) FetchWalletByCustID(custID string) (*model.Wallet, error) {
	return s.walletRepository.GetWalletByCustID(custID)
}
