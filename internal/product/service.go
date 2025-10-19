package product

type ProductService interface {
	GetProducts() ([]Products, error)
	GetProductById(id int) ([]Products, error)
	CreateProduct(product Products) error
	UpdateProduct(product Products) error
	DeleteProduct(id int) error
}

type productService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) GetProducts() ([]Products, error) {
	result, err := s.repo.GetProducts()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *productService) GetProductById(id int) ([]Products, error) {
	result, found, err := s.repo.GetProductById(id)
	if err != nil {
		return []Products{}, err
	}
	if !found {
		return []Products{}, nil
	}
	return []Products{result}, nil
}

func (s *productService) CreateProduct(product Products) error {
	err := s.repo.CreateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *productService) UpdateProduct(product Products) error {
	err := s.repo.UpdateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *productService) DeleteProduct(id int) error {
	err := s.repo.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}
