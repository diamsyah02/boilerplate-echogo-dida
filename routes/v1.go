package routes

import (
	"net/http"

	"boilerplate-echogo-dida/configs"
	"boilerplate-echogo-dida/modules/auth"
	"boilerplate-echogo-dida/modules/product"
	"boilerplate-echogo-dida/modules/user"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	api := e.Group("/api")
	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Boilerplate Rest Api Echo Golang")
	})
	// auth
	api.POST("/login", auth.LoginHandler)
	api.POST("/register", auth.RegisterHandler)

	// user
	routeUser := api.Group("/user")
	routeUser.Use(configs.EchoJWTConfig)
	routeUser.GET("/", user.GetUsersHandler)
	routeUser.PUT("/", user.UpdateUserHandler)

	// product
	routeProduct := api.Group("/product")
	routeProduct.Use(configs.EchoJWTConfig)
	routeProduct.GET("/", product.GetProductsHandler)
	routeProduct.GET("/:id", product.GetProductByIdHandler)
	routeProduct.POST("/", product.CreateProductHandler)
	routeProduct.PUT("/", product.UpdateProductHandler)
	routeProduct.DELETE("/:id", product.DeleteProductHandler)
}
