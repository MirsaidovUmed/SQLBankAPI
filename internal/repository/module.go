package repository

import (
	"sqlBankCLI/internal/models"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Conn *pgx.Conn
}

type RepositoryInterface interface {
	CreateBankAccount(account models.Account) (err error)
	CreateProfitAccount(fullName string, phoneNumber string, address string) (err error)
	GetAccountByName(name string) (account models.Account, err error)
	GetAccountByPhoneNumber(phoneNumber string) (account models.Account, err error)
	ChangeAccountBalance(account models.Account) (err error)
	GetPercent() float64
	AddTransfer(sender *models.Account, recipient *models.Account, amount float64) (err error)
}

func NewRepository(conn *pgx.Conn) RepositoryInterface {
	return &Repository{
		Conn: conn,
	}
}
