package auth

import (
	"hafiztri123/hv1-job-tracker/internal/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id, email string) (string, error) {

	claims := &Claims{
		UserId: id,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   id,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func secret() string {
	return utils.GetEnv("JWT_SECRET", "secret")
}
