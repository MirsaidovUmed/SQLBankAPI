package transport

import "fmt"

func (t *Transport) TransferMoney() {
	var senderPhone, recipientPhone string
	var amount float64

	fmt.Println("Здравсвуйте")

	fmt.Println("Введите номер отправителя")

	fmt.Scan(&senderPhone)

	fmt.Println("Введите номер получателя")
	fmt.Scan(&recipientPhone)

	fmt.Println("Введите сумму перевода")

	fmt.Scan(&amount)

	err := t.Svc.TransferMoney(senderPhone, recipientPhone, amount)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Успешно переведено")

}
