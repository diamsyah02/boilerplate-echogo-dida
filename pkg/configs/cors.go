package configs

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var CORSConfig = middleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
}
