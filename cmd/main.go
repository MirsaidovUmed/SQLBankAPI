package main

import (
	"fmt"
	"sqlBankCLI/pkg/config"
	"sqlBankCLI/pkg/database"
	"sqlBankCLI/pkg/repository"
	"sqlBankCLI/pkg/service"
	"sqlBankCLI/pkg/transport"
)

func main() {
	conf := config.NewConfig()
	db := database.NewDatabase(conf)
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	transp := transport.NewTransport(svc)

	for {
		var choice int

		fmt.Println("<<BANK CLI WITH SQL>>")
		fmt.Println("1. Создать счет в банке")
		fmt.Println("2. Пополнить счёт клиента")
		fmt.Println("3. Снять деньги со счёт клиента")
		fmt.Println("4. Показать счёт клиента")
		fmt.Println("5. Перевод денег")
		fmt.Println("6. Выйти")

		fmt.Scan(&choice)

		if choice == 1 {
			transp.CreateBankAccount()
		} else if choice == 2 {
			transp.TopUpClientsAccount()
		} else if choice == 3 {
			transp.WithdrawClientAccount()
		} else if choice == 4 {
			transp.ShowAccountBalance()
		} else if choice == 5 {
			transp.TransferMoney()
		} else if choice == 6 {
			return
		}
	}
}
