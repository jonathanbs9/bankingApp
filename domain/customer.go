package domain

import "github.com/jonathanbs9/bankingApp/errs"

type Customer struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	City        string `json:"city"`
	ZipCode     string `json:"zip_code"`
	Status      bool  `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	GetCustomerById(string) (*Customer, *errs.AppError)
}
