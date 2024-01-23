package transport

import (
	"fmt"
	"sqlBankCLI/internal/models"
)

func (t *Transport) WithdrawClientAccount() {
	var account models.Account
	var amount float64

	fmt.Println("Здравствуйте")

	fmt.Println("Введите номер телефона")

	fmt.Scan(&account.PhoneNumber)

	fmt.Println("Введите сумму которую хотите снять")

	fmt.Scan(&amount)

	err := t.Svc.WithdrawClientAccount(account, amount)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Деньги усешно сняты!")
}
