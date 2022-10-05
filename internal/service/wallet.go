package service

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"julo-test/internal/constants"
	"julo-test/internal/model"
	"julo-test/internal/request"
	"julo-test/internal/response"
	"julo-test/internal/storage"
	"julo-test/internal/util"
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

func (s *walletService) CreateWallet(initWalletReq *request.InitWalletRequest) (response.Code, string, error) {
	s.config.Logger().Info("CreateWallet: creating wallet")

	wallet := &model.Wallet{
		OwnedBy: initWalletReq.CustomerXID,
		Status:  constants.WalletDisabled,
		Balance: 0,
	}

	existingWallet, err := s.store.FetchWalletByCustID(initWalletReq.CustomerXID)

	if existingWallet != nil {

		jwt, err := util.GenerateToken(initWalletReq.CustomerXID)

		if err != nil {
			s.config.Logger().Error("CreateWallet: error generating jwt", zap.Error(err))
			return response.ServerError, "", err
		}

		return response.Success, jwt, nil
	}

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			s.config.Logger().Error("CreateWallet: error fetching wallet by cust id", zap.Error(err))
			return response.ServerError, "", err
		}
	}

	err = s.store.CreateWallet(wallet)

	if err != nil {

		s.config.Logger().Error("CreateWallet: error creating wallet", zap.Error(err))
		return response.ServerError, "", err

	}

	jwt, err := util.GenerateToken(initWalletReq.CustomerXID)

	if err != nil {
		s.config.Logger().Error("CreateWallet: error generating jwt", zap.Error(err))
		return response.ServerError, "", err
	}

	return response.Success, jwt, nil
}
