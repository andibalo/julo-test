package config

import (
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	logger2 "julo-test/internal/logger"
	"julo-test/internal/util"
	"os"
	"time"
)

const envVarEnvironment = "ENV"

func InitConfig() *AppConfig {
	logger := logger2.NewMainLoggerSingleton()

	redisClient := InitRedisClient(viper.GetString("REDIS_HOST"), "")

	return &AppConfig{
		logger:      logger,
		environment: os.Getenv(envVarEnvironment),
		redisClient: redisClient,
	}
}

type AppConfig struct {
	logger      *zap.Logger
	environment string
	redisClient *redis.Client
}

type Config interface {
	Logger() *zap.Logger
	ServerAddress() string
	StorageAddress() string
	RedisClient() *redis.Client
	RedisGetUserWalletBalanceTTL() time.Duration
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

func (a *AppConfig) RedisClient() *redis.Client {

	return a.redisClient
}

func (a *AppConfig) RedisGetUserWalletBalanceTTL() time.Duration {
	return time.Duration(5) * time.Second
}

func (a *AppConfig) JWTConfig() middleware.JWTConfig {

	return util.DefaultJWTConfig()
}
