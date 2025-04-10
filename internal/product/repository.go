package product

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Product, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	return products, err
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}
