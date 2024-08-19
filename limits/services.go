package limits

type Service interface {
	Create(customerId int, customerRequest LimitRequest) (Limit, error)
	FindByCustomerId(customerId int) ([]Limit, error)
}

type service struct {
	repository repository
}

func NewService(repository repository) *service {
	return &service{repository}
}

func (s *service) FindByCustomerId(id int) ([]Limit, error) {

	customer, err := s.repository.FindByCustomerId(id)

	return customer, err
}

func (s *service) Create(customerId int, customerRequest LimitRequest) (Limit, error) {

	customer := Limit{
		CustomerId: customerId,
		Tenor:      customerRequest.Tenor,
		Amount:     customerRequest.Amount,
	}

	customer, err := s.repository.Create(customer)

	return customer, err
}
