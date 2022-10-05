package v1

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	voerrors "julo-test/internal/apperrors"
	"julo-test/internal/config"
	"julo-test/internal/constants"
	"julo-test/internal/request"
	"julo-test/internal/response"
	"julo-test/internal/service"
	"julo-test/internal/util"
	"net/http"
)

type InitController struct {
	cfg           config.Config
	walletService service.WalletService
	validator     util.Validator
}

func NewInitController(cfg config.Config, walletService service.WalletService) *InitController {

	validator := util.GetNewValidator()

	return &InitController{
		cfg:           cfg,
		walletService: walletService,
		validator:     validator,
	}
}

func (h *InitController) AddRoutes(e *echo.Echo) {
	r := e.Group(constants.InitBasePath)

	r.POST("", h.InitializeWallet)
}

func (h *InitController) InitializeWallet(c echo.Context) error {
	initWalletReq := &request.InitWalletRequest{}

	if err := c.Bind(initWalletReq); err != nil {
		h.cfg.Logger().Error("InitializeWallet: error binding init wallet request", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, nil)
	}

	err := h.validator.Validate(initWalletReq)

	if err != nil {
		validationErrorMessage := err.Error()
		return h.failedInitResponse(c, response.BadRequest, err, validationErrorMessage)
	}

	code, token, err := h.walletService.CreateWallet(initWalletReq)

	if err != nil {
		h.cfg.Logger().Error("InitializeWallet: error creating wallet", zap.Error(err))

		return h.failedInitResponse(c, code, err, "error creating wallet")
	}

	initWalletResp := response.InitWalletResponse{Token: token}

	resp := response.NewResponse(code, initWalletResp)

	resp.SetResponseMessage("Successfully initialized wallet")

	return c.JSON(http.StatusOK, resp)
}

func (h *InitController) failedInitResponse(c echo.Context, code response.Code, err error, errorMsg string) error {
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
