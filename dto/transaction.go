package dto

import "github.com/jonathanbs9/bankingApp/errs"

const (
	WITHDRAWAL = "withdrawal"
	DEPOSIT    = "deposit"
)

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerId      string  `json:"-"`
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

/*  Validate func => Condiciones para validar la transacion. En este caso:
Tipo de transaccion => 'deposit' y 'withdraw'.
Monto no puede ser menor o igual a cero
*/
func (r TransactionRequest) Validate() *errs.AppError {
	if r.TransactionType != WITHDRAWAL && r.TransactionType != DEPOSIT {
		return errs.NewValidationError("Tipo de transacción inválida...")
	}
	if r.Amount <= 0 {
		return errs.NewValidationError("Monto no puede ser menor o igual a cero... ")
	}
	return nil
}

func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	if r.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}
