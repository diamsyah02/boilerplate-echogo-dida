package v1

import (
	"boilerplate-echogo-dida/internal/user"
	"boilerplate-echogo-dida/pkg/configs"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Group) {
	userRepo := user.NewUserRepository()
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	u := e.Group("/users")
	u.Use(configs.EchoJWTConfig)
	u.GET("", userHandler.GetUsers)
	u.PUT("", userHandler.UpdateUser)
}
