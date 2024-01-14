package service

import (
	"fmt"
)

func (s *Service) TransferMoney() {
	var senderName, recipientName, senderPassword, recipientPassword string
	var amount float64

	fmt.Println("Введите имя отправителя")

	fmt.Scan(&senderName)

	fmt.Println("Введите пароль")

	fmt.Scan(&senderPassword)

	sender, err := s.Repository.GetAccount(senderName, senderPassword)

	if err != nil {
		fmt.Println("Отсуствует счет отправителя")
		return
	}

	fmt.Println("Введите имя получателя")
	fmt.Scan(&recipientName)

	fmt.Println("Введите пароль")

	fmt.Scan(&recipientPassword)

	recipient, err := s.Repository.GetAccount(recipientName, recipientPassword)

	if err != nil {
		fmt.Println("Отсуствует счет получателя")
		return
	}

	fmt.Println("Введите сумму перевода")

	fmt.Scan(&amount)

	if sender.Balance < amount {
		fmt.Println("Недостаточно средств")
		return
	}

	var comission float64 = amount / 100 * s.Repository.GetPercent()

	sender.Balance -= amount

	//change account balance
	s.Repository.ChangeAccountsBalance(sender)

	recipient.Balance += amount - comission

	// change account balance
	s.Repository.ChangeAccountsBalance(recipient)

	// top up profit account
	s.Repository.TopUpProfitAccount(comission)

	// addTransfer
	s.Repository.AddTransfer(sender, recipient, amount)

	fmt.Println("Успешно переведено")
}
