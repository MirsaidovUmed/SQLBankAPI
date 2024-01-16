package transport

import (
	"fmt"
	"sqlBankCLI/pkg/models"
)

func (t *Transport) CreateBankAccount() {
	var account models.Account

	fmt.Println("Здравствуйте")

	fmt.Println("Введите ваше ФИО:")

	fmt.Scan(&account.FullName)

	fmt.Println("Введите ваш номер телефона:")

	fmt.Scan(&account.PhoneNumber)

	fmt.Println("Введите ваш адрес:")

	fmt.Scan(&account.Address)
	err := t.Svc.CreateBankAccount(account)

	if err != nil {
		fmt.Println(err)
		return
	}

}
