package domain

import "github.com/jonathanbs9/bankingApp/errs"

type Customer struct {
	Id          string `db:"customer_id",json:"id"`
	FirstName   string `db:"first_name",json:"first_name"`
	LastName    string `db:"last_name",json:"last_name"`
	DateOfBirth string `db:"date_birth",json:"date_of_birth"`
	City        string `db:"city",json:"city"`
	ZipCode     string `db:"zipcode",json:"zip_code"`
	Status      bool   `db:"status",json:"status"`
}

type CustomerRepository interface {
	// Status  1 = active | 0 = inactive | "" (empty)
	FindAll(status string) ([]Customer, *errs.AppError)
	GetCustomerById(string) (*Customer, *errs.AppError)
}
