package service

import (
	"sqlBankCLI/internal/models"
	"sqlBankCLI/pkg/errors"
)

func (s *Service) WithdrawClientAccount(account models.Account, amount float64) (err error) {
	if len(account.PhoneNumber) != 9 {
		return errors.ErrIncorrectPhoneNumber
	}

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
