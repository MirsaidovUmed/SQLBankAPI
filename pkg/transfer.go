package pkg

import "fmt"

// Проверка двух счетов
// Подсчет комиссии
// Перевод (
//	Снятие денег с 1ого счета,
//	Пополнение денег на 2ой счет,
//	Пополенение Прибыль в банк
//)

func TransferMoney() {
	var sender, recipient string
	var amount float64

	fmt.Println("Введите имя отправителя")

	fmt.Scan(&sender)

	balanceSender, ok := Database[sender]

	if !ok {
		fmt.Println("Отсуствует счет отправителя")
		return
	}

	fmt.Println("Введите имя получателя")
	fmt.Scan(&recipient)

	balanceRecipient, ok := Database[recipient]

	if !ok {
		fmt.Println("Отсуствует счет получателя")
		return
	}

	fmt.Println("Введите сумму перевода")
	fmt.Scan(&amount)

	if balanceSender < amount {
		fmt.Println("Недостаточно средств")
		return
	}
	var comission float64 = amount / 100 * Percent

	Database[sender] = balanceSender - amount
	Database[recipient] = (balanceRecipient + amount) - comission

	balanceProfit, ok := Database["Profit"]

	if !ok {
		Database["Profit"] = 0
	}

	Database["Profit"] = balanceProfit + comission

	fmt.Println("Успешно переведено")
}
