package domain

type Customer struct {
	Id          string
	FirstName   string
	LastName    string
	DateOfBirth string
	City        string
	ZipCode     string
	Status      bool
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
