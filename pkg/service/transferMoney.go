package service

func (s *Service) TransferMoney(senderPhone, recipientPhone string, amount float64) error {
	sender, err := s.Repo.GetAccountByPhoneNumber(senderPhone)
	if err != nil {
		return err
	}

	recipient, err := s.Repo.GetAccountByPhoneNumber(recipientPhone)
	if err != nil {
		return err
	}

	if sender.Balance < amount {
		return err
	}

	comission := amount / 100 * s.Repo.GetPercent()
	sender.Balance -= amount + comission

	err = s.Repo.ChangeAccountBalance(sender)
	if err != nil {
		return err
	}

	recipient.Balance += amount - comission
	err = s.Repo.ChangeAccountBalance(recipient)
	if err != nil {
		return err
	}

	profitAccount, err := s.Repo.GetAccountByName("profit")
	if err != nil {
		err = s.Repo.CreateProfitAccount("profit", "544", "address")
		if err != nil {
			return err
		}
	}

	profitAccount.Balance += comission
	err = s.Repo.ChangeAccountBalance(profitAccount)
	if err != nil {
		return err
	}

	s.Repo.AddTransfer(&sender, &recipient, amount)

	return nil
}
