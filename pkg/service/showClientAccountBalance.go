package service

import "fmt"

func (s *Service) ShowClientsAccount() {
	var name string

	fmt.Println("Введите имя клиента")

	fmt.Scan(&name)

	account, err := s.Repository.GetAccount(name)

	if err != nil {
		fmt.Println("Ошибка!, данного пользователя нет в нашей бд")
		return
	}

	fmt.Printf("Баланс счета %v является %v \n", name, account.Balance)
}
