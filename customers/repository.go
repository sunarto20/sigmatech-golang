package customers

import "gorm.io/gorm"

type Repository interface {
	Find() ([]Customer, error)
	FindById(ID int) (Customer, error)
	Create(customer Customer) (Customer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Find() ([]Customer, error) {
	var customers []Customer

	err := r.db.Find(&customers).Error

	return customers, err
}

func (r *repository) FindById(id int) (Customer, error) {
	var customer Customer
	err := r.db.First(&customer, "id = ?", id).Error
	return customer, err
}

func (r *repository) Create(customer Customer) (Customer, error) {
	err := r.db.Create(&customer).Error

	return customer, err
}
