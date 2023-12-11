package pkg

import "fmt"

func TopUpClientsAccount() {
	var name string
	var amount float64

	fmt.Println("Введите имя клиента")

	fmt.Scan(&name)

	// проверка на наличие клиента
	balance, ok := Database[name]

	if !ok {
		fmt.Println("Ошибка!, данного пользователя нет в нашей бд")
		return
	}

	fmt.Println("Введите сумму которую хотите пополнить")
	fmt.Scan(&amount)

	Database[name] = balance + amount

}
