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

type WithdrawalController struct {
	cfg        config.Config
	txnService service.TransactionService
	validator  util.Validator
}

func NewWithdrawalController(cfg config.Config, txnService service.TransactionService) *WithdrawalController {

	validator := util.GetNewValidator()

	return &WithdrawalController{
		cfg:        cfg,
		txnService: txnService,
		validator:  validator,
	}
}

func (h *WithdrawalController) AddRoutes(e *echo.Echo) {
	_ = e.Group(constants.WalletBasePath)

}

func (h *WithdrawalController) failedWalletResponse(c echo.Context, code response.Code, err error, errorMsg string) error {
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
