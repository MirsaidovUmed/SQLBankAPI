package pkg

import "fmt"

func WithdrawClientsAccount() {
	var name string
	var amount float64

	fmt.Println("Введите имя клиента")

	fmt.Scan(&name)

	var balance float64
	var has bool
	// проверка на наличие клиента
	for _, val := range Accounts {
		if name == val.Name {
			balance = val.Balance
			has = true
		}
	}

	if !has {
		fmt.Println("Ошибка! данного пользователя нет в нашей бд")
		return
	}

	fmt.Println("Введите сумму которую хотите снять")
	fmt.Scan(&amount)

	if balance < amount {
		fmt.Println("Ошибка! недостаточно средств на балансе")
		return
	}

	for _, val := range Accounts {
		if name == val.Name {
			val.Balance = balance - amount
		}
	}
}
