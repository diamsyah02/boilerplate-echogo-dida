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

	p := e.Group("/products")
	p.Use(configs.EchoJWTConfig)
	p.GET("", productHandler.GetProducts)
	p.GET("/:id", productHandler.GetProductById)
	p.POST("", productHandler.CreateProduct)
	p.PUT("", productHandler.UpdateProduct)
	p.DELETE("/:id", productHandler.DeleteProduct)
}
