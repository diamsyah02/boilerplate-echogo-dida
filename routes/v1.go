package routes

import (
	"net/http"

	"boilerplate-echogo-dida/modules/auth"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	api := e.Group("/api")
	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Boilerplate Rest Api Echo Golang")
	})
	api.POST("/login", auth.LoginHandler)
	api.POST("/register", auth.RegisterHandler)
}
