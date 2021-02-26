package service

import (
	"github.com/jonathanbs9/bankingApp/domain"
	"github.com/jonathanbs9/bankingApp/dto"
	"github.com/jonathanbs9/bankingApp/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

// New Account Service
func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}

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
