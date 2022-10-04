package util

import (
	"errors"
	"fmt"
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

func ParseToken(tokenString string) (claim model.Claims, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("JWT_SECRET")), nil
	})
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to parse JWT Token")
		return claim, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claim.Name = fmt.Sprintf("%v", claims["cxid"])
	} else {
		log.Println(err)
		err = errors.New("Failed to parse private claims")
		return claim, err
	}

	return claim, nil

}
