package customers

import "time"

type ServiceInterface interface {
	Find() ([]Customer, error)
	FindById(id int) (*Customer, error)
	Create(customerRequest CustomerRequest) (Customer, error)
}

type service struct {
	repository repository
}

func NewService(repository repository) *service {

	return &service{repository}
}

func (s *service) Find() ([]Customer, error) {
	return s.repository.Find()
}

func (s *service) FindById(id int) (*Customer, error) {

	customer, err := s.repository.FindById(id)

	return &customer, err
}
func (s *service) Create(customerRequest CustomerRequest) (Customer, error) {

	tanggalLahir, _ := time.Parse("2006-01-02", customerRequest.TanggalLahir)

	customer := Customer{
		Name:         customerRequest.Name,
		KTP:          customerRequest.KTP,
		FullName:     customerRequest.FullName,
		LegalName:    customerRequest.LegalName,
		TempatLahir:  customerRequest.TempatLahir,
		TanggalLahir: tanggalLahir,
		Gaji:         customerRequest.Gaji,
		FotoKTP:      customerRequest.FotoKTP,
		FotoSelfie:   customerRequest.FotoSelfie,
	}

	customer, err := s.repository.Create(customer)

	return customer, err
}
