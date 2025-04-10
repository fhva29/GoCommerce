package product

type Service interface {
	GetAllProducts() ([]Product, error)
}

type service struct {
	repo Repository
}

func (s *service) GetAllProducts() ([]Product, error) {
	// add validation
	return s.repo.FindAll()
}

func NewService(r Repository) Service {
	return &service{r}
}
