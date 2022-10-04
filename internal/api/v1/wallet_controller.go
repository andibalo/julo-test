package v1

import (
	"github.com/labstack/echo/v4"
	voerrors "julo-test/internal/apperrors"
	"julo-test/internal/config"
	"julo-test/internal/constants"
	"julo-test/internal/response"
	"julo-test/internal/service"
	"julo-test/internal/util"
)

type WalletController struct {
	cfg           config.Config
	walletService service.WalletService
	validator     util.Validator
}

func NewWalletController(cfg config.Config, walletService service.WalletService) *WalletController {

	validator := util.GetNewValidator()

	return &WalletController{
		cfg:           cfg,
		walletService: walletService,
		validator:     validator,
	}
}

func (h *WalletController) AddRoutes(e *echo.Echo) {
	_ = e.Group(constants.WalletBasePath)

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
