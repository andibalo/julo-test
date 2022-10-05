package config

import (
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	logger2 "julo-test/internal/logger"
	"julo-test/internal/util"
	"os"
)

const envVarEnvironment = "ENV"

func InitConfig() *AppConfig {
	logger := logger2.NewMainLoggerSingleton()

	return &AppConfig{
		logger:      logger,
		environment: os.Getenv(envVarEnvironment),
	}
}

type AppConfig struct {
	logger      *zap.Logger
	environment string
}

type Config interface {
	Logger() *zap.Logger
	ServerAddress() string
	StorageAddress() string
	JWTSecret() string
	JWTConfig() middleware.JWTConfig
}

func (a *AppConfig) Logger() *zap.Logger {
	return a.logger
}

func (a *AppConfig) ServerAddress() string {
	return viper.GetString("SERVER_PORT")
}

func (a *AppConfig) JWTSecret() string {

	return viper.GetString("JWT_SECRET")
}

func (a *AppConfig) StorageAddress() string {

	return fmt.Sprintf("%s?parseTime=true", viper.GetString("STORAGE_DSN"))
}

func (a *AppConfig) JWTConfig() middleware.JWTConfig {

	return util.DefaultJWTConfig()
}
