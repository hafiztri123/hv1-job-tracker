package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
