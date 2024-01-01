package main

import (
	"bankCLI/pkg/database"
	"bankCLI/pkg/repository"
	"bankCLI/pkg/service"
	"fmt"
)

func main() {
	db := database.NewDatabase(10.0)

	repo := repository.NewRepository(db)

	svc := service.NewService(repo)

	svc.FillCities()

	for {
		var choice int

		fmt.Println("1. Создание клиента, и его счет")
		fmt.Println("2. Пополнить счет клиента")
		fmt.Println("3. Посмотреть баланс клиента")
		fmt.Println("4. Снять деньги с баланса")
		fmt.Println("5. Перевод денег")
		fmt.Println("6. Выйти")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			svc.RegisterClient()
		case 2:
			svc.TopUpClientsAccount()
		case 3:
			svc.ShowClientsAccount()
		case 4:
			svc.WithdrawClientsAccount()
		case 5:
			svc.TransferMoney()
		case 6:
			return
		}
	}
}
