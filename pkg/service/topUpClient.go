package service

import (
	"fmt"
)

func (s *Service) TopUpClientsAccount() {
	var name string
	var amount float64

	fmt.Println("Введите имя клиента")

	fmt.Scan(&name)

	//Get Account
	account, err := s.Repository.GetAccount(name)

	if err != nil {
		fmt.Println("Ошибка!, данного пользователя нет в нашей бд")
		return
	}

	fmt.Println("Введите сумму которую хотите пополнить")
	fmt.Scan(&amount)

	s.Repository.ChangeAccountsBalance(account, account.Balance+amount)

}
