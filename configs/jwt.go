package configs

import (
	"net/http"
	"os"

	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var EchoJWTConfig = echoJwt.WithConfig(echoJwt.Config{
	SigningKey: []byte(os.Getenv("JWT_SECRET")),
	ErrorHandler: func(c echo.Context, err error) error {
		if err != nil && err.Error() == "token is expired" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Token expired, silakan login ulang.",
			})
		}
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Token tidak valid atau hilang.",
		})
	},
})
