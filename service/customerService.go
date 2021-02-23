package service

import (
	"github.com/jonathanbs9/bankingApp/domain"
	"github.com/jonathanbs9/bankingApp/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomerById(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// Here we link service with repository
func (s DefaultCustomerService) GetCustomerById(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.GetCustomerById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
