package auth

import (
	"hafiztri123/hv1-job-tracker/internal/utils"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {
	auth := c.Get("Authorization")

	if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
		slog.Error("Authorization")
		return utils.NewResponse(
			c,
			utils.WithStatus(http.StatusUnauthorized),
			utils.WithMessage("Unauthorized"),
		)
	}

	tokenString := strings.TrimPrefix(auth, "Bearer ")

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret()), nil
	})

	if err != nil || !token.Valid {
		return utils.NewResponse(
			c,
			utils.WithMessage("Unauthorized"),
			utils.WithStatus(http.StatusUnauthorized),
			utils.WithError(err.Error()),
		)
	}

	c.Locals("userId", claims.UserId)
	c.Locals("email", claims.Email)

	return c.Next()

}
