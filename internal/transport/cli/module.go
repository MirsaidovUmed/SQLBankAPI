package transport

import "sqlBankCLI/internal/service"

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
