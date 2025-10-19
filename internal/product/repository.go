package product

import (
	"boilerplate-echogo-dida/pkg/configs"
	"errors"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProducts() ([]Products, error)
	GetProductById(id int) (Products, bool, error)
	CreateProduct(product Products) error
	UpdateProduct(product Products) error
	DeleteProduct(id int) error
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (s *productRepository) GetProducts() ([]Products, error) {
	var products []Products
	if err := configs.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productRepository) GetProductById(id int) (Products, bool, error) {
	var product Products
	if err := configs.DB.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Products{}, false, nil
		}
		return Products{}, false, err
	}
	return product, true, nil
}

func (s *productRepository) CreateProduct(product Products) error {
	if err := configs.DB.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (s *productRepository) UpdateProduct(product Products) error {
	if err := configs.DB.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

func (s *productRepository) DeleteProduct(id int) error {
	if err := configs.DB.Delete(&Products{}, id).Error; err != nil {
		return err
	}
	return nil
}
