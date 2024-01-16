package repository

func (repo *Repository) TopUpProfitAccount(amount float64) error {
	account, err := repo.GetAccountByName("profit")
	if err != nil {
		return err
	}

	account.Balance += amount
	err = repo.ChangeAccountBalance(account)
	if err != nil {
		return err
	}
	return nil
}
