package service

import (
	"sqlBankCLI/internal/models"
	"sqlBankCLI/pkg/errors"
)

func (s *Service) CreateBankAccount(account models.Account) error {
	if len(account.PhoneNumber) != 9 {
		return errors.ErrIncorrectPhoneNumber
	}

	account2, err := s.Repo.GetAccountByName(account.FullName)

	if err != errors.ErrDataNotFound {
		return err
	}

	if account2.PhoneNumber == account.PhoneNumber {
		return errors.ErrAlreadyHasAccount
	}

	err = s.Repo.CreateBankAccount(account)

	if err != nil {
		return err
	}

	return nil
}
