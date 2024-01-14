package repository

import (
	"bankCLI/pkg/config"
	"bankCLI/pkg/models"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Database *pgx.Conn
	Config   *config.Config
}

type RepositoryInterface interface {
	AddAccount(name string, balance float64, phone_number, password string, city *models.City) (err error)
	AddCity(name, region string) (err error)
	AddTransfer(sender *models.Account, recipient *models.Account, amount float64) (err error)
	ChangeAccountsBalance(account *models.Account) (err error)
	GetAccount(name, password string) (account *models.Account, err error)
	GetCity(name string) (city *models.City, err error)
	GetPercent() float64
	TopUpProfitAccount(amount float64)
}

func NewRepository(db *pgx.Conn, config *config.Config) RepositoryInterface {
	return &Repository{
		Database: db,
		Config:   config,
	}
}
