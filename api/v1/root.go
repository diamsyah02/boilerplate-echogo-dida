package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	api := e.Group("/api")
	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Boilerplate Rest Api Echo Golang")
	})

	RouteAuth(api)    // auth
	RouteUser(api)    // user
	RouteProduct(api) // product
}
