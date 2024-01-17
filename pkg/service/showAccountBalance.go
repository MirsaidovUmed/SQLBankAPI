package service

import (
	"sqlBankCLI/pkg/errors"
	"sqlBankCLI/pkg/models"
)

func (s *Service) ShowAccountBalance(account models.Account) (name string, balance float64, err error) {
	if len(account.PhoneNumber) != 9 {
		return "", 0, errors.ErrIncorrectPhoneNumber
	}
	account, err = s.Repo.GetAccountByPhoneNumber(account.PhoneNumber)
	if err != nil {
		return name, 0, err
	}
	return account.FullName, account.Balance, nil
}
