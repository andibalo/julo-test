package repositories

import (
	"gorm.io/gorm"
	"julo-test/internal/model"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {

	return &WalletRepository{
		db: db,
	}
}

func (p *WalletRepository) SaveWallet(chat *model.Wallet) error {

	err := p.db.Create(chat).Error

	if err != nil {
		return err
	}

	return nil
}

func (p *WalletRepository) GetAllWallets() (*[]model.Wallet, error) {

	var chats *[]model.Wallet

	err := p.db.Find(&chats).Error

	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (p *WalletRepository) GetAllWalletsByUserID(userID string) (*[]model.Wallet, error) {

	var chats *[]model.Wallet

	err := p.db.Where("user_id = ?", userID).Find(&chats).Error

	if err != nil {
		return nil, err
	}

	return chats, nil
}
