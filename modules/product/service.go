package product

func GetProductsService() ([]Products, error) {
	result, err := GetProductsRepository()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetProductByIdService(id int) ([]Products, error) {
	result, found, err := GetProductByIdRepository(id)
	if err != nil {
		return []Products{}, err
	}
	if !found {
		return []Products{}, nil
	}
	return []Products{result}, nil
}

func CreateProductService(product Products) error {
	err := CreateProductRepository(product)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProductService(product Products) error {
	err := UpdateProductRepository(product)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProductService(id int) error {
	err := DeleteProductRepository(id)
	if err != nil {
		return err
	}
	return nil
}
