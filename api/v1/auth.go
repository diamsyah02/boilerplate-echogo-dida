package v1

import (
	"boilerplate-echogo-dida/internal/auth"

	"github.com/labstack/echo/v4"
)

func RouteAuth(e *echo.Group) {
	authRepo := auth.NewAuthRepository()
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService)

	e.POST("/login", authHandler.Login)
	e.POST("/register", authHandler.Register)
}
