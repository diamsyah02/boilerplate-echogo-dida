package main

import (
	"boilerplate-echogo-dida/configs"
	"boilerplate-echogo-dida/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Gagal memuat file .env: %v", err)
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
