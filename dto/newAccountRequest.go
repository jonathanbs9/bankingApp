package dto

import "github.com/jonathanbs9/bankingApp/errs"

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("Para abrir una cuenta se necesita al menos $5000 ")
	}
	if r.AccountType != "saving" && r.AccountType != "checking"{
		return errs.NewValidationError("Tipo de cuenta debe ser 'saving' o 'checking'")
	}
	return nil
}
