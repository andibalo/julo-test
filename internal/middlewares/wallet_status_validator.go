package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"julo-test/internal/constants"
	"julo-test/internal/model"
	"julo-test/internal/storage"
	"net/http"
)

func ValidateWalletIsDisabled(store storage.Storage) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*model.Claims)

			cxid := claims.CustomerXID

			wallet, err := store.FetchWalletByCustID(cxid)

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return echo.NewHTTPError(http.StatusNotFound, "Wallet not found")
				}

				return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching wallet")
			}

			if wallet.Status == constants.WalletEnabled {
				return echo.NewHTTPError(http.StatusForbidden, "wallet needs to be disabled to use this resource")
			}

			return next(c)

		}
	}
}

func ValidateWalletIsEnabled(store storage.Storage) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*model.Claims)

			cxid := claims.CustomerXID

			wallet, err := store.FetchWalletByCustID(cxid)

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return echo.NewHTTPError(http.StatusNotFound, "Wallet not found")
				}

				return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching wallet")
			}

			if wallet.Status == constants.WalletDisabled {
				return echo.NewHTTPError(http.StatusForbidden, "wallet needs to be enabled to use this resource")
			}

			return next(c)

		}
	}
}
