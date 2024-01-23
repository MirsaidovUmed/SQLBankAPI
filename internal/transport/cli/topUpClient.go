package transport

import (
	"fmt"
	"sqlBankCLI/internal/models"
)

func (t *Transport) TopUpClientsAccount() {
	var account models.Account
	var amount float64

	fmt.Println("Здравствуйте")

	fmt.Println("Введите номер телефона")

	fmt.Scan(&account.PhoneNumber)

	fmt.Println("Введите сумму которую хотите пополнить")

	fmt.Scan(&amount)

	err := t.Svc.TopUpClientsAccount(account, amount)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Счет успешно пополнен!")
}
