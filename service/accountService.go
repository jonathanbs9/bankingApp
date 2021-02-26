package service

import (
	"github.com/jonathanbs9/bankingApp/domain"
	"github.com/jonathanbs9/bankingApp/dto"
	"github.com/jonathanbs9/bankingApp/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

// New Account Service
func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}

// New Account
func (s DefaultAccountService) NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.Account{
		AccountId:   "",
		CustomerId:  request.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	// Case  new account
	response := newAccount.ToNewAccountResponseDto()
	return &response, nil
}

// Make Transaction
func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	// Validate incoming request
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	// Checking the availability balance in the account
	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.Findby(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Saldo insuficiente en la cuenta")
		}
	}

	// If everything is ok => Build domain object & save transaction
	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05Z07:00"),
	}
	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, err
	}
	response := transaction.ToDto()
	return &response, nil
}
