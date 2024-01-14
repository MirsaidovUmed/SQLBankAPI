package service

import "fmt"

func (s *Service) ShowClientsAccount() {
	var name, password string

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

	fmt.Printf("Баланс счета %v является %v \n", name, account.Balance)
}
