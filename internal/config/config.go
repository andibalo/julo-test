package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	logger2 "julo-test/internal/logger"
	"os"
)

const envVarEnvironment = "ENV"

func InitConfig() *AppConfig {
	logger := logger2.NewMainLoggerSingleton()

	pgConf := LoadDBConfig()

	return &AppConfig{
		logger:      logger,
		environment: os.Getenv(envVarEnvironment),
		pgConf:      pgConf,
	}
}

type AppConfig struct {
	logger      *zap.Logger
	environment string
	pgConf      *DBConfig
}

type Config interface {
	Logger() *zap.Logger
	ServerAddress() string
	PgConfig() *DBConfig
	JWTSecret() string
}

func (a *AppConfig) Logger() *zap.Logger {
	return a.logger
}

func (a *AppConfig) ServerAddress() string {
	return viper.GetString("SERVER_PORT")
}

func (a *AppConfig) UserDataFilePath() string {
	return viper.GetString("FILE_PATH")
}

func (a *AppConfig) PgConfig() *DBConfig {

	return a.pgConf
}

func (a *AppConfig) JWTSecret() string {

	return viper.GetString("JWT_SECRET")
}
