package transport

import "sqlBankCLI/pkg/service"

type Transport struct {
	Svc service.ServiceInterface
}

type TransportInterface interface {
	CreateBankAccount()
	ShowAccountBalance()
	TopUpClientsAccount()
	WithdrawClientAccount()
	TransferMoney()
}

func NewTransport(svc service.ServiceInterface) TransportInterface {
	return &Transport{
		Svc: svc,
	}
}
