package service

import (
	"fmt"
)

func (s *Service) TopUpClientsAccount() {
	var name, password string
	var amount float64

	fmt.Println("Введите имя клиента")

	fmt.Scan(&name)

	fmt.Println("Введите пароль")

	fmt.Scan(&password)
	account, err := s.Repository.GetAccount(name, password)

	if err != nil {
		fmt.Println("Ошибка!, данного пользователя нет в нашей бд")
		return
	}
	if password != account.Password {
		fmt.Println("Неправильный пароль")
		return
	}

	fmt.Println("Введите сумму которую хотите пополнить")
	fmt.Scan(&amount)

	// Update the account balance
	account.Balance += amount

	// Now update the account in the database
	err = s.Repository.ChangeAccountsBalance(account)
	if err != nil {
		fmt.Println("Ошибка при обновлении баланса аккаунта: ", err)
		return
	}
}
