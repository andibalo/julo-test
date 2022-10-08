package v1

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"gorm.io/gorm"
	voerrors "julo-test/internal/apperrors"
	"julo-test/internal/config"
	"julo-test/internal/constants"
	customMiddleware "julo-test/internal/middlewares"
	"julo-test/internal/model"
	"julo-test/internal/request"
	"julo-test/internal/response"
	"julo-test/internal/service"
	"julo-test/internal/storage"
	"julo-test/internal/util"
	"net/http"
)

type WalletController struct {
	cfg           config.Config
	walletService service.WalletService
	txnService    service.TransactionService
	validator     util.Validator
	store         storage.Storage
}

func NewWalletController(cfg config.Config, walletService service.WalletService, txnService service.TransactionService,
	store storage.Storage) *WalletController {

	validator := util.GetNewValidator()

	return &WalletController{
		cfg:           cfg,
		walletService: walletService,
		txnService:    txnService,
		validator:     validator,
		store:         store,
	}
}

func (h *WalletController) AddRoutes(e *echo.Echo) {
	r := e.Group(constants.WalletBasePath, middleware.JWTWithConfig(h.cfg.JWTConfig()))

	r.POST("", h.EnableWallet, customMiddleware.ValidateWalletIsDisabled(h.store))
	r.GET("", h.GetWalletBalance, customMiddleware.ValidateWalletIsEnabled(h.store))
	r.PATCH("", h.DisableWallet, customMiddleware.ValidateWalletIsEnabled(h.store))
	r.POST(constants.DepositBasePath, h.DepositWallet, customMiddleware.ValidateWalletIsEnabled(h.store))
	r.POST(constants.WithdrawalBasePath, h.WithdrawFromWallet, customMiddleware.ValidateWalletIsEnabled(h.store))
}

func (h *WalletController) WithdrawFromWallet(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.Claims)

	cxid := claims.CustomerXID

	withdrawWalletReq := &request.WithdrawWalletRequest{}

	if err := c.Bind(withdrawWalletReq); err != nil {
		h.cfg.Logger().Error("WithdrawFromWallet: error binding withdraw from wallet request", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, nil)
	}

	err := h.validator.Validate(withdrawWalletReq)

	if err != nil {
		validationErrorMessage := err.Error()
		return h.failedWalletResponse(c, response.BadRequest, err, validationErrorMessage)
	}

	code, wallet, txn, err := h.walletService.WithdrawFromWallet(withdrawWalletReq.Amount, cxid, withdrawWalletReq.RefID)
	if err != nil {
		h.cfg.Logger().Error("WithdrawFromWallet: error depositing to wallet", zap.Error(err))

		errorMsg := "error withdrawing from wallet"

		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorMsg = "wallet not found"
		}

		if errors.Is(err, voerrors.ErrInsufficientBalance) {
			errorMsg = "not enough balance"
		}

		return h.failedWalletResponse(c, code, err, errorMsg)
	}

	withdrawInfo := response.WithdrawWalletDetail{
		ID:          txn.ID,
		WithdrawnBy: wallet.OwnedBy,
		Status:      txn.Status,
		WithdrawnAt: txn.WithdrawnAt.Time,
		Amount:      txn.Amount,
		RefID:       txn.ReferenceId,
	}

	withdrawWalletResp := response.WithdrawWalletResponse{Withdrawal: withdrawInfo}

	resp := response.NewResponse(code, withdrawWalletResp)

	resp.SetResponseMessage("Successfully withdrawn from wallet")

	return c.JSON(http.StatusOK, resp)
}

func (h *WalletController) DepositWallet(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.Claims)

	cxid := claims.CustomerXID

	depositWalletReq := &request.DepositWalletRequest{}

	if err := c.Bind(depositWalletReq); err != nil {
		h.cfg.Logger().Error("DepositWallet: error binding dpeeosit wallet request", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, nil)
	}

	err := h.validator.Validate(depositWalletReq)

	if err != nil {
		validationErrorMessage := err.Error()
		return h.failedWalletResponse(c, response.BadRequest, err, validationErrorMessage)
	}

	code, wallet, txn, err := h.walletService.DepositWallet(depositWalletReq.Amount, cxid, depositWalletReq.RefID)
	if err != nil {
		h.cfg.Logger().Error("DepositWallet: error depositing to wallet", zap.Error(err))

		errorMsg := "error depositing to wallet"

		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorMsg = "wallet not found"
		}

		return h.failedWalletResponse(c, code, err, errorMsg)
	}

	depositInfo := response.DepositWalletDetail{
		ID:          txn.ID,
		DepositedBy: wallet.OwnedBy,
		Status:      txn.Status,
		DepositedAt: txn.DepositedAt.Time,
		Amount:      txn.Amount,
		RefID:       txn.ReferenceId,
	}

	depositWalletResp := response.DepositWalletResponse{Deposit: depositInfo}

	resp := response.NewResponse(code, depositWalletResp)

	resp.SetResponseMessage("Successfully deposited to wallet")

	return c.JSON(http.StatusOK, resp)
}

func (h *WalletController) GetWalletBalance(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.Claims)

	cxid := claims.CustomerXID

	wallet := &model.Wallet{}

	cachedWallet, err := h.cfg.RedisClient().Get(context.Background(), constants.GetUserWalletBalanceRedisKey(cxid)).Result()
	if err != nil {
		h.cfg.Logger().Error("GetWalletBalance: error fetching wallet balance from cache", zap.Error(err))
	}

	if cachedWallet != "" {
		h.cfg.Logger().Info("GetWalletBalance: fetched wallet from cache")

		if err = util.Deserialize(cachedWallet, wallet); err != nil {
			h.cfg.Logger().Error("GetWalletBalance: failed deserializing user wallet cache", zap.Error(err))

			return h.failedWalletResponse(c, response.ServerError, err, "failed deserializing user wallet cache")
		}

		resp := h.buildGetWalletBalanceResp(wallet, response.Success, "Successfully fetched wallet")

		return c.JSON(http.StatusOK, resp)
	}

	code, wallet, err := h.walletService.FetchWalletBalance(cxid)

	if err != nil {
		h.cfg.Logger().Error("GetWalletBalance: error fetching wallet balance", zap.Error(err))

		errorMsg := "error fetching wallet balance"

		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorMsg = "wallet not found"
		}

		return h.failedWalletResponse(c, code, err, errorMsg)
	}

	serializedWallet, err := util.Serialize(wallet)
	if err != nil {
		h.cfg.Logger().Error("GetWalletBalance: failed serializing user wallet", zap.Error(err))
	}

	if serializedWallet != "" {
		cacheKey := constants.GetUserWalletBalanceRedisKey(cxid)
		if err := h.cfg.RedisClient().Set(context.Background(), cacheKey,
			serializedWallet, h.cfg.RedisGetUserWalletBalanceTTL()).Err(); err != nil {

			h.cfg.Logger().Error("GetWalletBalance: failed caching user wallet", zap.Error(err))

			return err
		}
	}

	resp := h.buildGetWalletBalanceResp(wallet, code, "Successfully fetched wallet")

	return c.JSON(http.StatusOK, resp)
}

func (h *WalletController) buildGetWalletBalanceResp(wallet *model.Wallet, code response.Code, msg string) *response.Wrapper {

	walletInfo := response.DefaultWalletInfo{
		ID:        wallet.ID,
		OwnedBy:   wallet.OwnedBy,
		Status:    wallet.Status,
		EnabledAt: wallet.EnabledAt.Time,
		Balance:   wallet.Balance,
	}

	enableWalletResp := response.FetchWalletBalanceResponse{Wallet: walletInfo}

	resp := response.NewResponse(code, enableWalletResp)

	resp.SetResponseMessage(msg)

	return resp
}

func (h *WalletController) DisableWallet(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.Claims)

	cxid := claims.CustomerXID

	code, wallet, err := h.walletService.DisableWallet(cxid)
	if err != nil {
		h.cfg.Logger().Error("DisableWallet: error disabling wallet", zap.Error(err))

		errorMsg := "error disabling wallet"

		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorMsg = "wallet not found"
		}

		return h.failedWalletResponse(c, code, err, errorMsg)
	}

	walletInfo := response.DisableWalletDetail{
		ID:         wallet.ID,
		OwnedBy:    wallet.OwnedBy,
		Status:     wallet.Status,
		DisabledAt: wallet.DisabledAt.Time,
		Balance:    wallet.Balance,
	}

	enableWalletResp := response.DisableWalletResponse{Wallet: walletInfo}

	resp := response.NewResponse(code, enableWalletResp)

	resp.SetResponseMessage("Successfully disabled wallet")

	return c.JSON(http.StatusOK, resp)
}

func (h *WalletController) EnableWallet(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.Claims)

	cxid := claims.CustomerXID

	code, wallet, err := h.walletService.EnableWallet(cxid)
	if err != nil {
		h.cfg.Logger().Error("EnableWallet: error enabling wallet", zap.Error(err))

		errorMsg := "error enabling wallet"

		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorMsg = "wallet not found"
		}

		return h.failedWalletResponse(c, code, err, errorMsg)
	}

	walletInfo := response.DefaultWalletInfo{
		ID:        wallet.ID,
		OwnedBy:   wallet.OwnedBy,
		Status:    wallet.Status,
		EnabledAt: wallet.EnabledAt.Time,
		Balance:   wallet.Balance,
	}

	enableWalletResp := response.EnableWalletResponse{Wallet: walletInfo}

	resp := response.NewResponse(code, enableWalletResp)

	resp.SetResponseMessage("Successfully enabled wallet")

	return c.JSON(http.StatusOK, resp)
}

func (h *WalletController) failedWalletResponse(c echo.Context, code response.Code, err error, errorMsg string) error {
	if code == "" {
		code = voerrors.MapErrorsToCode(err)
	}

	resp := response.Wrapper{
		ResponseCode: code,
		Status:       code.GetStatus(),
		Message:      code.GetMessage(),
	}

	if errorMsg != "" {
		resp.SetResponseMessage(errorMsg)
	}

	return c.JSON(voerrors.MapErrorsToStatusCode(err), resp)
}
