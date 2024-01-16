package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"sqlBankCLI/pkg/errors"
	"sqlBankCLI/pkg/models"
)

func (repo *Repository) GetAccountByPhoneNumber(phoneNumber string) (account models.Account, err error) {
	row := repo.Conn.QueryRow(context.Background(), "SELECT id, full_name, phone_number, balance from account where phone_number = $1", phoneNumber)

	err = row.Scan(&account.Id, &account.FullName, &account.PhoneNumber, &account.Balance)

	if err != nil {
		if err == pgx.ErrNoRows {
			return account, errors.ErrDataNotFound
		}
	}

	return
}
