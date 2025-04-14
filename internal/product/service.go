package product

type Service interface {
	CreateProduct(req CreateProductRequest) (Product, error)
	GetAllProducts() ([]Product, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) CreateProduct(req CreateProductRequest) (Product, error) {
	if errs := req.Validate(); errs != nil {
		return Product{}, &ValidationError{Errors: errs}
	}

	product := Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}

	return s.repo.Create(product)
}

func (s *service) GetAllProducts() ([]Product, error) {
	// add validation
	return s.repo.FindAll()
}
