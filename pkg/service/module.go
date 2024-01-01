package service

import "bankCLI/pkg/repository"

type Service struct {
	Repository repository.RepositoryInterface
}

type ServiceInterface interface {
	FillCities() (err error)
	RegisterClient()
	ShowClientsAccount()
	TopUpClientsAccount()
	TransferMoney()
	WithdrawClientsAccount()
}

func NewService(repo repository.RepositoryInterface) ServiceInterface {
	return &Service{
		Repository: repo,
	}
}
