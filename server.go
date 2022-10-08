package julo

import (
	"github.com/labstack/echo/v4"
	"julo-test/internal/api"
	v1 "julo-test/internal/api/v1"
	"julo-test/internal/config"
	"julo-test/internal/service"
	"julo-test/internal/storage"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
}

func NewServer(cfg *config.AppConfig) *Server {

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.New().String()
		},
	}))

	store := storage.New(cfg)

	walletService := service.NewWalletService(cfg, store)
	txnService := service.NewTransactionService(cfg, store)

	walletHandler := v1.NewWalletController(cfg, walletService, txnService, store)
	initHandler := v1.NewInitController(cfg, walletService)

	registerHandlers(e, &api.HealthCheck{}, walletHandler, initHandler)

	return &Server{
		echo: e,
	}
}

func (s *Server) Start(addr string) error {
	return s.echo.Start(":" + addr)
}

func registerHandlers(e *echo.Echo, handlers ...api.Handler) {
	for _, handler := range handlers {
		handler.AddRoutes(e)
	}
}
