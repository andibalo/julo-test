package repositories

import (
	"gorm.io/gorm"
	"julo-test/internal/model"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {

	return &TransactionRepository{
		db: db,
	}
}

func (p *TransactionRepository) SaveTransaction(transaction *model.Transaction) error {

	err := p.db.Create(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (p *TransactionRepository) GetTransactionByRefID(refID string) (*model.Transaction, error) {

	var transaction *model.Transaction

	err := p.db.Where("reference_id = ?", refID).First(&transaction).Error

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (p *TransactionRepository) GetAllTransactions() (*[]model.Transaction, error) {

	var chats *[]model.Transaction

	err := p.db.Find(&chats).Error

	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (p *TransactionRepository) GetAllTransactionsByUserID(userID string) (*[]model.Transaction, error) {

	var chats *[]model.Transaction

	err := p.db.Where("user_id = ?", userID).Find(&chats).Error

	if err != nil {
		return nil, err
	}

	return chats, nil
}
