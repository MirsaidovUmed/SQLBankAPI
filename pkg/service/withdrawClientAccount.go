package service

import "sqlBankCLI/pkg/models"

func (s *Service) WithdrawClientAccount(account models.Account, amount float64) (err error) {
	account, err = s.Repo.GetAccountByPhoneNumber(account.PhoneNumber)

	if err != nil {
		return
	}

	if account.Balance < amount {
		return err
	}

	account.Balance -= amount

	err = s.Repo.ChangeAccountBalance(account)
	if err != nil {
		return
	}

	return
}
