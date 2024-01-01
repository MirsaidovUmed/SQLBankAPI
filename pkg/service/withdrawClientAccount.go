package service

import "fmt"

func (s *Service) WithdrawClientsAccount() {
	var name string
	var amount float64

	fmt.Println("Введите имя клиента")

	fmt.Scan(&name)

	account, err := s.Repository.GetAccount(name)

	if err != nil {
		fmt.Println("Ошибка! данного пользователя нет в нашей бд")
		return
	}

	fmt.Println("Введите сумму которую хотите снять")
	fmt.Scan(&amount)

	if account.Balance < amount {
		fmt.Println("Ошибка! недостаточно средств на балансе")
		return
	}

	s.Repository.ChangeAccountsBalance(account, account.Balance-amount)
}
