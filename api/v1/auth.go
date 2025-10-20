package v1

import (
	"boilerplate-echogo-dida/internal/auth"
	"boilerplate-echogo-dida/pkg/tasks"
	"os"

	"github.com/labstack/echo/v4"
)

func RouteAuth(e *echo.Group) {
	redisAddr := os.Getenv("REDIS_ADDRESS")
	distributor := tasks.NewTaskDistributor(redisAddr)
	authRepo := auth.NewAuthRepository()
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService, distributor)

	e.POST("/login", authHandler.Login)
	e.POST("/register", authHandler.Register)
}
