package repository

func (repo *Repository) TopUpProfitAccount(amount float64) {
	account, err := repo.GetAccount("profit", "0")

	if err != nil {
		city, err := repo.GetCity("CityName") // Замените "CityName" на имя города
		if err != nil {
			// Обработка ошибки
			return
		}

		err = repo.AddAccount("profit", 0, "544", "0", city)
		if err != nil {
			// Обработка ошибки
			return
		}

		account, err = repo.GetAccount("profit", "0")
		if err != nil {
			// Обработка ошибки
			return
		}
	}

	account.Balance += amount
	err = repo.ChangeAccountsBalance(account)
	if err != nil {
		// Обработка ошибки
	}
}
