package product

import (
	"boilerplate-echogo-dida/configs"
	"errors"

	"gorm.io/gorm"
)

func GetProductsRepository() ([]Products, error) {
	var products []Products
	if err := configs.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductByIdRepository(id int) (Products, bool, error) {
	var product Products
	if err := configs.DB.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Products{}, false, nil
		}
		return Products{}, false, err
	}
	return product, true, nil
}

func CreateProductRepository(product Products) error {
	if err := configs.DB.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProductRepository(product Products) error {
	if err := configs.DB.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProductRepository(id int) error {
	if err := configs.DB.Delete(&Products{}, id).Error; err != nil {
		return err
	}
	return nil
}
