package repository

func (repo *Repository) TopUpProfitAccount(amount float64) {
	account, err := repo.GetAccount("profit")

	if err != nil {
		repo.AddAccount("profit", 0, "544", repo.Database.Cities[0])

		account, _ = repo.GetAccount("profit")
	}

	account.Balance += amount
}
