package product_test

import (
	"boilerplate-echogo-dida/internal/product"
	"strings"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockProductRoute struct {
	getProductsFunc    func() ([]product.Products, error)
	getProductByIdFunc func(id int) ([]product.Products, error)
	createProductFunc  func(product product.Products) error
	updateProductFunc  func(product product.Products) error
	deleteProductFunc  func(id int) error
}

// GetProducts implements product handler.
func (m *mockProductRoute) GetProducts() ([]product.Products, error) {
	return m.getProductsFunc()
}

// GetProductById implements product handler.
func (m *mockProductRoute) GetProductById(id int) ([]product.Products, error) {
	return m.getProductByIdFunc(id)
}

// CreateProduct implements product handler.
func (m *mockProductRoute) CreateProduct(product product.Products) error {
	return m.createProductFunc(product)
}

// UpdateProduct implements product handler.
func (m *mockProductRoute) UpdateProduct(product product.Products) error {
	return m.updateProductFunc(product)
}

// DeleteProduct implements product handler.
func (m *mockProductRoute) DeleteProduct(id int) error {
	return m.deleteProductFunc(id)
}

func TestGetProductsRoute(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mock := &mockProductRoute{
		getProductsFunc: func() ([]product.Products, error) {
			return []product.Products{}, nil
		},
	}
	productRoute := product.NewProductHandler(mock)
	err := productRoute.GetProducts(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Get products success")
}

func TestGetProductByIdRoute(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/products/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mock := &mockProductRoute{
		getProductByIdFunc: func(id int) ([]product.Products, error) {
			return []product.Products{}, nil
		},
	}
	productRoute := product.NewProductHandler(mock)
	err := productRoute.GetProductById(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Get product by id success")
}

func TestCreateProductRoute(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/products", strings.NewReader(`{"name":"dida","price":123, "stock":123}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mock := &mockProductRoute{
		createProductFunc: func(product product.Products) error {
			return nil
		},
	}
	productRoute := product.NewProductHandler(mock)
	err := productRoute.CreateProduct(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Create product success")
}

func TestUpdateProductRoute(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/api/products", strings.NewReader(`{"name":"dida","price":123, "stock":123}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mock := &mockProductRoute{
		updateProductFunc: func(product product.Products) error {
			return nil
		},
	}
	productRoute := product.NewProductHandler(mock)
	err := productRoute.UpdateProduct(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Update product success")
}

func TestDeleteProductRoute(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/api/products/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mock := &mockProductRoute{
		deleteProductFunc: func(id int) error {
			return nil
		},
	}
	productRoute := product.NewProductHandler(mock)
	err := productRoute.DeleteProduct(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Delete product success")
}
