package service

import (
	"sqlBankCLI/pkg/models"
	"sqlBankCLI/pkg/repository"
)

type Service struct {
	Repo repository.RepositoryInterface
}

type ServiceInterface interface {
	CreateBankAccount(account models.Account) error
	ShowAccountBalance(account models.Account) (name string, balance float64, err error)
	TopUpClientsAccount(account models.Account, amount float64) (err error)
	WithdrawClientAccount(account models.Account, amount float64) (err error)
	TransferMoney(senderPhone, recipientPhone string, amount float64) error
}

func NewService(repo repository.RepositoryInterface) ServiceInterface {
	return &Service{
		Repo: repo,
	}
}
