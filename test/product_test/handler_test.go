package product_test

import (
	"boilerplate-echogo-dida/internal/product"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockHandler struct {
	getProductsFunc    func() ([]product.Products, error)
	getProductByIdFunc func(id int) ([]product.Products, error)
	createProductFunc  func(product.Products) error
	updateProductFunc  func(product.Products) error
	deleteProductFunc  func(id int) error
}

// GetProducts implements product handler.
func (m *mockHandler) GetProducts() ([]product.Products, error) {
	return m.getProductsFunc()
}

// GetProductById implements product handler.
func (m *mockHandler) GetProductById(id int) ([]product.Products, error) {
	return m.getProductByIdFunc(id)
}

// CreateProduct implements product handler.
func (m *mockHandler) CreateProduct(product product.Products) error {
	return m.createProductFunc(product)
}

// UpdateProduct implements product handler.
func (m *mockHandler) UpdateProduct(product product.Products) error {
	return m.updateProductFunc(product)
}

// DeleteProduct implements product handler.
func (m *mockHandler) DeleteProduct(id int) error {
	return m.deleteProductFunc(id)
}

func setupContext(method, path, body string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func TestGetProductsHandler(t *testing.T) {
	c, rec := setupContext(http.MethodGet, "/api/products", "")

	mock := &mockHandler{
		getProductsFunc: func() ([]product.Products, error) {
			return []product.Products{}, nil
		},
	}
	productHandler := product.NewProductHandler(mock)
	err := productHandler.GetProducts(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Get products success")
}

func TestGetProductByIdHandler(t *testing.T) {
	c, rec := setupContext(http.MethodGet, "/api/products/1", "")
	c.SetPath("/api/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mock := &mockHandler{
		getProductByIdFunc: func(id int) ([]product.Products, error) {
			return []product.Products{}, nil
		},
	}
	productHandler := product.NewProductHandler(mock)
	err := productHandler.GetProductById(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Get product by id success")
}

func TestCreateProductHandler(t *testing.T) {
	c, rec := setupContext(http.MethodPost, "/api/products", `{"name":"dida","price":123, "stock":123}`)

	mock := &mockHandler{
		createProductFunc: func(product product.Products) error {
			return nil
		},
	}

	productHandler := product.NewProductHandler(mock)
	err := productHandler.CreateProduct(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Create product success")
}

func TestUpdateProductHandler(t *testing.T) {
	c, rec := setupContext(http.MethodPut, "/api/products", `{"id":1,"name":"dida","price":123, "stock":123}`)

	mock := &mockHandler{
		updateProductFunc: func(product product.Products) error {
			return nil
		},
	}

	productHandler := product.NewProductHandler(mock)
	err := productHandler.UpdateProduct(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Update product success")
}

func TestDeleteProductHandler(t *testing.T) {
	c, rec := setupContext(http.MethodDelete, "/api/products/1", "")
	c.SetPath("/api/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mock := &mockHandler{
		deleteProductFunc: func(id int) error {
			return nil
		},
	}

	productHandler := product.NewProductHandler(mock)
	err := productHandler.DeleteProduct(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Delete product success")
}
