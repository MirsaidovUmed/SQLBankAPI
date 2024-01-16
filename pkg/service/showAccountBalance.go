package service

import "sqlBankCLI/pkg/models"

func (s *Service) ShowAccountBalance(account models.Account) (name string, balance float64, err error) {
	account, err = s.Repo.GetAccountByPhoneNumber(account.PhoneNumber)
	if err != nil {
		return name, 0, err
	}
	return account.FullName, account.Balance, nil
}
