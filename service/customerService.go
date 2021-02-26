package service

import (
	"github.com/jonathanbs9/bankingApp/domain"
	"github.com/jonathanbs9/bankingApp/dto"
	"github.com/jonathanbs9/bankingApp/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomerById(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// New Customer Service
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

// Get All Customers
func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active"{
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	customers, err := s.repo.FindAll(status)
	if err != nil{
		return nil, err
	}

	response := make([]dto.CustomerResponse,0)
	for _ , c := range customers {
		response = append(response, c.ToDto())
	}

	return response, err
}

// Get Customer by ID
// Here we link service with repository
func (s DefaultCustomerService) GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.GetCustomerById(id)
	if err != nil {
		return nil, err
	}
	// convierto lo que me traigo de la BD a un DTO con la func ToDto()
	response := c.ToDto()
	return &response, nil
}


