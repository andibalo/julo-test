package v1

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"gorm.io/gorm"
	voerrors "julo-test/internal/apperrors"
	"julo-test/internal/config"
	"julo-test/internal/constants"
	"julo-test/internal/model"
	"julo-test/internal/response"
	"julo-test/internal/service"
	"julo-test/internal/storage"
	"julo-test/internal/util"
	"net/http"
)

type WalletController struct {
	cfg           config.Config
	walletService service.WalletService
	validator     util.Validator
	store         storage.Storage
}

func NewWalletController(cfg config.Config, walletService service.WalletService, store storage.Storage) *WalletController {

	validator := util.GetNewValidator()

	return &WalletController{
		cfg:           cfg,
		walletService: walletService,
		validator:     validator,
		store:         store,
	}
}

func (h *WalletController) AddRoutes(e *echo.Echo) {
	r := e.Group(constants.WalletBasePath, middleware.JWTWithConfig(h.cfg.JWTConfig()))

	r.POST("", h.EnableWallet)
	r.GET("", h.GetWalletBalance)
}

func (h *WalletController) GetWalletBalance(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.Claims)

	cxid := claims.CustomerXID

	code, wallet, err := h.walletService.FetchWalletBalance(cxid)
	if err != nil {
		h.cfg.Logger().Error("GetWalletBalance: error fetching wallet balance", zap.Error(err))

		errorMsg := "error fetching wallet balance"

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

	enableWalletResp := response.FetchWalletBalanceResponse{Wallet: walletInfo}

	resp := response.NewResponse(code, enableWalletResp)

	resp.SetResponseMessage("Successfully fetched wallet")

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
