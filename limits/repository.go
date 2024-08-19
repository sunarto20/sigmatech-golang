package limits

import "gorm.io/gorm"

type Repository interface {
	FindByCustomerId(CustomerId int) ([]Limit, error)
	Create(customer Limit) (Limit, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByCustomerId(customerId int) ([]Limit, error) {
	var limits []Limit
	err := r.db.Find(&limits, "customer_id = ?", customerId).Error
	return limits, err
}

func (r *repository) Create(limit Limit) (Limit, error) {
	err := r.db.Create(&limit).Error

	return limit, err
}
