package product

import (
	"net/http"
	"strconv"

	"boilerplate-echogo-dida/utils"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type ProductHandler struct {
	service ProductService
}

func NewProductHandler(s ProductService) ProductHandler {
	return ProductHandler{s}
}

func (h *ProductHandler) GetProductsHandler(c echo.Context) error {
	logger := log.With().Str("function", "GetProductsHandler").Logger()

	products, err := h.service.GetProductsService()
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to get products")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Get products failed", nil))
	}
	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Get products success", products))
}

func (h *ProductHandler) GetProductByIdHandler(c echo.Context) error {
	logger := log.With().Str("function", "GetProductByIdHandler").Logger()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		logger.Error().
			Str("param", c.Param("id")).
			Err(err).
			Msg("Invalid product ID param")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Invalid product ID", nil))
	}
	result, err := h.service.GetProductByIdService(id)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to get product by id")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Get product by id failed", nil))
	}
	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Get product by id success", result))
}

func (h *ProductHandler) CreateProductHandler(c echo.Context) error {
	logger := log.With().Str("function", "CreateProductHandler").Logger()

	product := Products{}
	if err := c.Bind(&product); err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to bind product")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Create product failed", nil))
	}
	if err := h.service.CreateProductService(product); err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to create product")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Create product failed", nil))
	}
	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Create product success", nil))
}

func (h *ProductHandler) UpdateProductHandler(c echo.Context) error {
	logger := log.With().Str("function", "UpdateProductHandler").Logger()

	product := Products{}
	if err := c.Bind(&product); err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to bind product")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Update product failed", nil))
	}
	if err := h.service.UpdateProductService(product); err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to update product")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Update product failed", nil))
	}
	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Update product success", nil))
}

func (h *ProductHandler) DeleteProductHandler(c echo.Context) error {
	logger := log.With().Str("function", "DeleteProductHandler").Logger()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error().
			Str("param", c.Param("id")).
			Err(err).
			Msg("Invalid product ID param")
		return c.JSON(http.StatusBadRequest, utils.WebResponse(http.StatusBadRequest, "Invalid product ID", nil))
	}

	if err := h.service.DeleteProductService(id); err != nil {
		logger.Error().
			Int("product_id", id).
			Err(err).
			Msg("Failed to delete product")
		return c.JSON(http.StatusInternalServerError, utils.WebResponse(http.StatusInternalServerError, "Delete product failed", nil))
	}

	return c.JSON(http.StatusOK, utils.WebResponse(http.StatusOK, "Delete product success", nil))
}
