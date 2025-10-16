package main

import (
	"boilerplate-echogo-dida/configs"
	"boilerplate-echogo-dida/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Error().
			Str("function", "main starting app").
			Err(err).
			Msg("Gagal memuat file .env")
	}
	configs.Connect()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(configs.CORSConfig))
	e.Use(middleware.RateLimiterWithConfig(configs.RateLimiterConfig))
	// Routing
	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":6000"))
}
