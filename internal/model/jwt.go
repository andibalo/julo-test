package model

import "github.com/golang-jwt/jwt"

type Claims struct {
	CustomerXID string `json:"cxid"`
	jwt.StandardClaims
}
