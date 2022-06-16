package service

import (
	"UdemyREST/domain"
	"UdemyREST/dto"
	"UdemyREST/errs"
	"time"
)

type AccountService interface {
	//define the request information
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	//define secondary port
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	//pass dto to repository
	//validate the incoming request 1
	err := req.Validate()
	// test for an error 2
	if err != nil {
		return nil, err
	}
	//on success create a domain account object 3
	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	//pass to repo and save to create a new acc 4
	newAccount, err := s.repo.Save(a)
	//check for err on this submission 5
	if err != nil {
		return nil, err
	}
	//if not err then we will see the new account dto 5
	response := newAccount.ToNewAccountResponseDto()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
