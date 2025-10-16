package configs

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

var RateLimiterConfig = middleware.RateLimiterConfig{
	Skipper: middleware.DefaultSkipper,
	Store: middleware.NewRateLimiterMemoryStoreWithConfig(
		middleware.RateLimiterMemoryStoreConfig{Rate: rate.Limit(10), Burst: 30, ExpiresIn: 3 * time.Minute},
	),
	IdentifierExtractor: func(ctx echo.Context) (string, error) {
		id := ctx.RealIP()
		return id, nil
	},
	ErrorHandler: func(c echo.Context, err error) error {
		return c.JSON(http.StatusForbidden, nil)
	},
	DenyHandler: func(c echo.Context, identifier string, err error) error {
		return c.JSON(http.StatusTooManyRequests, map[string]interface{}{
			"message": "Too many requests",
		})
	},
}
