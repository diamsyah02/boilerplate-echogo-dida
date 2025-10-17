package product

import (
	"boilerplate-echogo-dida/configs"
	"errors"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductsRepository() ([]Products, error)
	GetProductByIdRepository(id int) (Products, bool, error)
	CreateProductRepository(product Products) error
	UpdateProductRepository(product Products) error
	DeleteProductRepository(id int) error
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (s *productRepository) GetProductsRepository() ([]Products, error) {
	var products []Products
	if err := configs.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productRepository) GetProductByIdRepository(id int) (Products, bool, error) {
	var product Products
	if err := configs.DB.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Products{}, false, nil
		}
		return Products{}, false, err
	}
	return product, true, nil
}

func (s *productRepository) CreateProductRepository(product Products) error {
	if err := configs.DB.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (s *productRepository) UpdateProductRepository(product Products) error {
	if err := configs.DB.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

func (s *productRepository) DeleteProductRepository(id int) error {
	if err := configs.DB.Delete(&Products{}, id).Error; err != nil {
		return err
	}
	return nil
}
