package repositories

import (
	"gorm.io/gorm"
	"julo-test/internal/constants"
	"julo-test/internal/model"
	"time"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {

	return &WalletRepository{
		db: db,
	}
}

func (p *WalletRepository) SaveWallet(wallet *model.Wallet) error {

	err := p.db.Create(wallet).Error

	if err != nil {
		return err
	}

	return nil
}

func (p *WalletRepository) UpdateWalletStatusByCustID(custID, status string) error {

	updatesMap := map[string]interface{}{"status": status}

	if status == constants.WalletEnabled {
		updatesMap["enabled_at"] = time.Now()
	} else {
		updatesMap["disabled_at"] = time.Now()
	}

	err := p.db.Model(&model.Wallet{}).Where("owned_by = ?", custID).Updates(updatesMap).Error

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

func (p *WalletRepository) GetWalletByCustID(custID string) (*model.Wallet, error) {

	var wallet *model.Wallet

	err := p.db.Where("owned_by = ?", custID).First(&wallet).Error

	if err != nil {
		return nil, err
	}

	return wallet, nil
}
