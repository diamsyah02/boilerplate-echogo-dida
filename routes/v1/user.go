package v1

import (
	"boilerplate-echogo-dida/configs"
	"boilerplate-echogo-dida/modules/user"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Group) {
	userRepo := user.NewUserRepository()
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	e.Use(configs.EchoJWTConfig)
	e.GET("", userHandler.GetUsersHandler)
	e.PUT("", userHandler.UpdateUserHandler)
}
