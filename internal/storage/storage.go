package storage

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"julo-test/internal/config"
	"julo-test/internal/model"
	"julo-test/internal/storage/repositories"
	"sync"
)

var onceDb sync.Once

var instance *gorm.DB

type Store struct {
	logger           *zap.Logger
	txnRepository    TransactionRepository
	walletRepository WalletRepository
}

func New(cfg *config.AppConfig) *Store {
	db := InitDB(cfg)

	migrateDB(db)

	txnRepo := repositories.NewTransactionRepository(db)
	walletRepo := repositories.NewWalletRepository(db)

	return &Store{
		logger:           cfg.Logger(),
		txnRepository:    txnRepo,
		walletRepository: walletRepo,
	}
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.Transaction{}, &model.Wallet{})
}

func InitDB(cfg *config.AppConfig) *gorm.DB {
	onceDb.Do(func() {

		db, err := gorm.Open(mysql.Open(cfg.StorageAddress()), &gorm.Config{})
		if err != nil {
			cfg.Logger().Fatal("Could not connect to database: %s", zap.Error(err))
		}

		cfg.Logger().Info("Successfully Connected to Database")

		instance = db
	})

	return instance
}

type Storage interface {
	CreateWallet(wallet *model.Wallet) error
	FetchWalletByCustID(custID string) (*model.Wallet, error)
	UpdateWalletStatusByCustID(custID, status string) error
	UpdateWalletBalanceByCustID(custID string, balance int, transaction *model.Transaction) error
	CreateTransaction(transaction *model.Transaction) error
	FetchTransactionByRefID(refID string) (*model.Transaction, error)
}

type WalletRepository interface {
	GetWalletByCustID(custID string) (*model.Wallet, error)
	SaveWallet(wallet *model.Wallet) error
	UpdateWalletStatusByCustID(custID, status string) error
	UpdateWalletBalanceByCustID(custID string, balance int, transaction *model.Transaction) error
}

type TransactionRepository interface {
	SaveTransaction(transaction *model.Transaction) error
	GetTransactionByRefID(refID string) (*model.Transaction, error)
}
