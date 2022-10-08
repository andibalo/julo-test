package util

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"julo-test/internal/model"
	"log"
)

func DefaultJWTConfig() middleware.JWTConfig {
	config := middleware.JWTConfig{
		Claims:     &model.Claims{},
		SigningKey: []byte(viper.GetString("JWT_SECRET")),
	}

	return config
}

func GenerateToken(cxid string) (tokenString string, err error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.Claims{
		CustomerXID: cxid,
	})

	tokenString, err = token.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to sign JWT token")
		return "", err
	}

	return tokenString, nil
}
