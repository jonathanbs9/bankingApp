package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", FirstName: "Jonathan", LastName: "Brull S.", DateOfBirth: "1986-06-23", City: "Mar del Plata",
			ZipCode: "7600", Status: true},
		{Id: "2", FirstName: "Ruben", LastName: "De Ronde", DateOfBirth: "1980-02-20", City: "Amsterdam",
			ZipCode: "7600", Status: true},
	}
	return CustomerRepositoryStub{customers: customers}
}
