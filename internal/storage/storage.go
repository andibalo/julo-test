package storage

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
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
	txnRepository    *repositories.TransactionRepository
	walletRepository *repositories.WalletRepository
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
		databaseConfig := cfg.PgConfig()

		db, err := gorm.Open(postgres.Open(databaseConfig.GetDBConnectionString()), &gorm.Config{})
		if err != nil {
			cfg.Logger().Fatal("Could not connect to database: %s", zap.Error(err))
		}

		cfg.Logger().Info("Successfully Connected to Database")

		instance = db
	})

	return instance
}

type Storage interface {
}
