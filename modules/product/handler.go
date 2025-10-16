package product

import (
	"net/http"
	"strconv"

	"boilerplate-echogo-dida/utils"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func GetProductsHandler(c echo.Context) error {
	products, err := GetProductsService()
	if err != nil {
		log.Error().
			Str("function", "GetProductsHandler").
			Err(err).
			Msg("Failed to get products")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Get products failed", nil))
	}
	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Get products success", products))
}

func GetProductByIdHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error().
			Str("function", "GetProductByIdHandler").
			Str("param", c.Param("id")).
			Err(err).
			Msg("Invalid product ID param")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Invalid product ID", nil))
	}
	product, err := GetProductByIdService(id)
	if err != nil {
		log.Error().
			Str("function", "GetProductByIdHandler").
			Err(err).
			Msg("Failed to get product by id")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Get product by id failed", nil))
	}
	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Get product by id success", product))
}

func CreateProductHandler(c echo.Context) error {
	product := Products{}
	if err := c.Bind(&product); err != nil {
		log.Error().
			Str("function", "CreateProductHandler").
			Err(err).
			Msg("Failed to bind product")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Create product failed", nil))
	}
	if err := CreateProductService(product); err != nil {
		log.Error().
			Str("function", "CreateProductHandler").
			Err(err).
			Msg("Failed to create product")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Create product failed", nil))
	}
	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Create product success", nil))
}

func UpdateProductHandler(c echo.Context) error {
	product := Products{}
	if err := c.Bind(&product); err != nil {
		log.Error().
			Str("function", "UpdateProductHandler").
			Err(err).
			Msg("Failed to bind product")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Update product failed", nil))
	}
	if err := UpdateProductService(product); err != nil {
		log.Error().
			Str("function", "UpdateProductHandler").
			Err(err).
			Msg("Failed to update product")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Update product failed", nil))
	}
	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Update product success", nil))
}

func DeleteProductHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error().
			Str("function", "DeleteProductHandler").
			Str("param", c.Param("id")).
			Err(err).
			Msg("Invalid product ID param")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Invalid product ID", nil))
	}

	if err := DeleteProductService(id); err != nil {
		log.Error().
			Str("function", "DeleteProductHandler").
			Int("product_id", id).
			Err(err).
			Msg("Failed to delete product")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Delete product failed", nil))
	}

	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Delete product success", nil))
}
