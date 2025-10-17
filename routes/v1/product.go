package v1

import (
	"boilerplate-echogo-dida/configs"
	"boilerplate-echogo-dida/modules/product"

	"github.com/labstack/echo/v4"
)

func RouteProduct(e *echo.Group) {
	productRepo := product.NewProductRepository()
	productService := product.NewProductService(productRepo)
	productHandler := product.NewProductHandler(productService)

	e.Use(configs.EchoJWTConfig)
	e.GET("", productHandler.GetProductsHandler)
	e.GET("/:id", productHandler.GetProductByIdHandler)
	e.POST("", productHandler.CreateProductHandler)
	e.PUT("", productHandler.UpdateProductHandler)
	e.DELETE("/:id", productHandler.DeleteProductHandler)
}
