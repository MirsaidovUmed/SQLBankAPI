package repository

import (
	"bankCLI/pkg/database"
	"bankCLI/pkg/models"
)

type Repository struct {
	Database *database.Database
}

type RepositoryInterface interface {
	AddAccount(name string, balance float64, phone_number string, city *models.City)
	AddCity(name, region string) (err error)
	AddTransfer(sender, recipient *models.Account, amount float64)
	ChangeAccountsBalance(account *models.Account, newBalance float64)
	GetAccount(name string) (account *models.Account, err error)
	GetCity(name string) (city *models.City, err error)
	GetPercent() float64
	TopUpProfitAccount(amount float64)
}

func NewRepository(db *database.Database) RepositoryInterface {
	return &Repository{
		Database: db,
	}
}
