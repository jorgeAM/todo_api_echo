package middlewares

import (
	"os"

	"github.com/jorgeAM/echo-api/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var key = os.Getenv("JWT_KEY")

// AuthMiddleware verify jwt in authorization header
func AuthMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &models.Claim{},
		SigningKey: []byte(key),
	}

	return middleware.JWTWithConfig(config)
}
