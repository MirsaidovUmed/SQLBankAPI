package pkg

import "fmt"

func ShowClientsAccount() {
	var name string

	fmt.Println("Введите имя клиента")

	fmt.Scan(&name)

	// проверка на наличие клиента
	balance, ok := Database[name]

	if !ok {
		fmt.Println("Ошибка!, данного пользователя нет в нашей бд")
		return
	}

	fmt.Printf("Баланс счета %v является %v \n", name, balance)
}
