package product

type ProductService interface {
	GetProductsService() ([]Products, error)
	GetProductByIdService(id int) ([]Products, error)
	CreateProductService(product Products) error
	UpdateProductService(product Products) error
	DeleteProductService(id int) error
}

type productService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) GetProductsService() ([]Products, error) {
	result, err := s.repo.GetProductsRepository()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *productService) GetProductByIdService(id int) ([]Products, error) {
	result, found, err := s.repo.GetProductByIdRepository(id)
	if err != nil {
		return []Products{}, err
	}
	if !found {
		return []Products{}, nil
	}
	return []Products{result}, nil
}

func (s *productService) CreateProductService(product Products) error {
	err := s.repo.CreateProductRepository(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *productService) UpdateProductService(product Products) error {
	err := s.repo.UpdateProductRepository(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *productService) DeleteProductService(id int) error {
	err := s.repo.DeleteProductRepository(id)
	if err != nil {
		return err
	}
	return nil
}
