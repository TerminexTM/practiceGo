package service

import (
	"UdemyREST/domain"
	"UdemyREST/dto"
	"UdemyREST/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

//business logic here to have dependencie of the repository
type DefaultCustomerService struct {
	repo domain.CustomerRepository //grab that interface!
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	//declaring a slice of my dto
	var myDto = []dto.CustomerResponse{}
	//simple business logic that mutates the url string when it is Received.

	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	c, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	//break apart my domain object and map to dto
	for _, item := range c {
		convert := item.ToDto()
		myDto = append(myDto, convert)
	}
	return myDto, nil
}

//connects primary port with secondary port pairing repo to serice layer
func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	//map the domain object to our dto and return it -- responsilibity of making a dto is now on the domain
	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
