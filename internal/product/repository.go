package product

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Product, error)
	Create(product Product) (Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	return products, err
}
