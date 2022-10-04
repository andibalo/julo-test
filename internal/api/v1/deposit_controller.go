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

type DepositController struct {
	cfg        config.Config
	txnService service.TransactionService
	validator  util.Validator
}

func NewDepositController(cfg config.Config, txnService service.TransactionService) *DepositController {

	validator := util.GetNewValidator()

	return &DepositController{
		cfg:        cfg,
		txnService: txnService,
		validator:  validator,
	}
}

func (h *DepositController) AddRoutes(e *echo.Echo) {
	_ = e.Group(constants.WalletBasePath)

}

func (h *DepositController) failedDepositResponse(c echo.Context, code response.Code, err error, errorMsg string) error {
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
