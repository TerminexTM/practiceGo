package service

import "UdemyREST/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

//business logic here to have dependencie of the repository
type DefaultCustomerService struct {
	repo domain.CustomerRepository //grab that interface!
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
