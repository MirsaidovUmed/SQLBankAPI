package pkg

import (
	"bankCLI/pkg/models"
	"fmt"
)

func TransferMoney() {
	var sender, recipient string
	var amount float64

	fmt.Println("Введите имя отправителя")

	fmt.Scan(&sender)

	var Sender *models.Account
	var has bool
	// проверка на наличие отправителя
	for _, val := range Accounts {
		if sender == val.Name {
			has = true
			Sender = val
		}
	}

	if !has {
		fmt.Println("Отсуствует счет отправителя")
		return
	}

	fmt.Println("Введите имя получателя")
	fmt.Scan(&recipient)

	has = false
	var Recipient *models.Account
	// проверка на наличие клиента
	for _, val := range Accounts {
		if recipient == val.Name {
			has = true
			Recipient = val
		}
	}

	if !has {
		fmt.Println("Отсуствует счет получателя")
		return
	}

	fmt.Println("Введите сумму перевода")
	fmt.Scan(&amount)

	if Sender.Balance < amount {
		fmt.Println("Недостаточно средств")
		return
	}
	var comission float64 = amount / 100 * Percent

	for _, val := range Accounts {
		if sender == val.Name {
			val.Balance = Sender.Balance - amount
		}
	}

	for _, val := range Accounts {
		if recipient == val.Name {
			val.Balance = Recipient.Balance + amount - comission
		}
	}
	has = false
	var balanceProfit float64
	for _, val := range Accounts {
		if "profit" == val.Name {
			balanceProfit = val.Balance
			has = true
		}
	}

	if !has {
		Accounts = append(Accounts, &models.Account{
			Name:        "profit",
			PhoneNumber: "554",
			Balance:     0,
		})
	}

	for _, val := range Accounts {
		if "profit" == val.Name {
			val.Balance = balanceProfit + comission
		}
	}

	Transfers = append(Transfers, models.Transfer{
		Sender:    Sender,
		Recipient: Recipient,
		Amount:    amount,
	})

	fmt.Println("Успешно переведено")
}
