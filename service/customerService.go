package service

import "github.com/jonathanbs9/bankingApp/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomerById(string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// Here we link service with repository
func (s DefaultCustomerService) GetCustomerById(id string) (*domain.Customer, error) {
	return s.repo.GetCustomerById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
