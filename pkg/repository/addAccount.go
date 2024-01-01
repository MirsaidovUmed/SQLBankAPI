package repository

import "bankCLI/pkg/models"

func (repo *Repository) AddAccount(name string, balance float64, phone_number string, city *models.City) {
	repo.Database.Accounts = append(repo.Database.Accounts, &models.Account{
		Name:        name,
		Balance:     balance,
		PhoneNumber: phone_number,
		City:        city,
	})
}
