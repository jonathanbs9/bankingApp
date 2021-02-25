package dto

type CustomerResponse struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	City        string `json:"city"`
	ZipCode     string `json:"zip_code"`
	Status      string `json:"status"`
}
