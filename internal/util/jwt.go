package util

import (
	"errors"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"log"

	"github.com/golang-jwt/jwt"
)

func NewJWTConfig() middleware.JWTConfig {
	config := middleware.JWTConfig{
		AuthScheme: "Token",
	}

	return config
}

func GenerateToken(cxid string) (tokenString string, err error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cxid": cxid,
	})

	tokenString, err = token.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to sign JWT token")
		return "", err
	}

	return tokenString, nil
}
