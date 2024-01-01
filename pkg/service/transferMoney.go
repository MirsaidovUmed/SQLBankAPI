package service

import (
	"fmt"
)

func (s *Service) TransferMoney() {
	var senderName, recipientName string
	var amount float64

	fmt.Println("Введите имя отправителя")

	fmt.Scan(&senderName)

	sender, err := s.Repository.GetAccount(senderName)

	if err != nil {
		fmt.Println("Отсуствует счет отправителя")
		return
	}

	fmt.Println("Введите имя получателя")
	fmt.Scan(&recipientName)

	recipient, err := s.Repository.GetAccount(recipientName)

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

	//change account balance
	s.Repository.ChangeAccountsBalance(sender, sender.Balance-amount)

	// change account balance
	s.Repository.ChangeAccountsBalance(recipient, recipient.Balance+(amount-comission))

	// top up profit account
	s.Repository.TopUpProfitAccount(comission)

	// addTransfer
	s.Repository.AddTransfer(sender, recipient, amount)

	fmt.Println("Успешно переведено")
}
