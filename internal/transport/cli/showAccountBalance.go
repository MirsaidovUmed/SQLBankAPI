package transport

import (
	"fmt"
	"sqlBankCLI/internal/models"
)

func (t *Transport) ShowAccountBalance() {
	var account models.Account

	fmt.Println("Здравствуйте!")

	fmt.Println("Введите номер телефона")

	fmt.Scan(&account.PhoneNumber)

	name, balance, err := t.Svc.ShowAccountBalance(account)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Баланс счета %v является %v \n", name, balance)
}
