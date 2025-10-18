package product_test

import (
	"boilerplate-echogo-dida/modules/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockService struct {
	getProductsFunc    func() ([]product.Products, error)
	getProductByIdFunc func(id int) (product.Products, bool, error)
	createProductFunc  func(product product.Products) error
	updateProductFunc  func(product product.Products) error
	deleteProductFunc  func(id int) error
}

// GetProducts implements product service.
func (m *mockService) GetProducts() ([]product.Products, error) {
	return m.getProductsFunc()
}

// GetProductById implements product service.
func (m *mockService) GetProductById(id int) (product.Products, bool, error) {
	return m.getProductByIdFunc(id)
}

// CreateProduct implements product service.
func (m *mockService) CreateProduct(product product.Products) error {
	return m.createProductFunc(product)
}

// UpdateProduct implements product service.
func (m *mockService) UpdateProduct(product product.Products) error {
	return m.updateProductFunc(product)
}

// DeleteProduct implements product service.
func (m *mockService) DeleteProduct(id int) error {
	return m.deleteProductFunc(id)
}

func TestGetProductsService(t *testing.T) {
	mockRepo := &mockService{}
	service := product.NewProductService(mockRepo)

	mockRepo.getProductsFunc = func() ([]product.Products, error) {
		return []product.Products{}, nil
	}

	products, err := service.GetProducts()
	assert.NoError(t, err)
	assert.Equal(t, []product.Products{}, products)
}

func TestGetProductByIdService(t *testing.T) {
	mockRepo := &mockService{}
	service := product.NewProductService(mockRepo)

	mockRepo.getProductByIdFunc = func(id int) (product.Products, bool, error) {
		return product.Products{}, true, nil
	}

	product, err := service.GetProductById(1)
	assert.NoError(t, err)
	assert.Equal(t, product, product)
}

func TestCreateProductService(t *testing.T) {
	mockRepo := &mockService{}
	service := product.NewProductService(mockRepo)

	mockRepo.createProductFunc = func(product product.Products) error {
		return nil
	}

	err := service.CreateProduct(product.Products{})
	assert.NoError(t, err)
}

func TestUpdateProductService(t *testing.T) {
	mockRepo := &mockService{}
	service := product.NewProductService(mockRepo)

	mockRepo.updateProductFunc = func(product product.Products) error {
		return nil
	}

	err := service.UpdateProduct(product.Products{})
	assert.NoError(t, err)
}

func TestDeleteProductService(t *testing.T) {
	mockRepo := &mockService{}
	service := product.NewProductService(mockRepo)

	mockRepo.deleteProductFunc = func(id int) error {
		return nil
	}

	err := service.DeleteProduct(1)
	assert.NoError(t, err)
}
