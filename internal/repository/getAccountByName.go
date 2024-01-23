package repository

import (
	"context"
	"sqlBankCLI/internal/models"
	"sqlBankCLI/pkg/errors"

	"github.com/jackc/pgx/v5"
)

func (repo *Repository) GetAccountByName(name string) (account models.Account, err error) {
	row := repo.Conn.QueryRow(context.Background(), `
		SELECT id, full_name, phone_number, address, balance FROM account WHERE full_name = $1
	`, name)

	err = row.Scan(
		&account.Id,
		&account.FullName,
		&account.PhoneNumber,
		&account.Address,
		&account.Balance,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return account, errors.ErrDataNotFound
		}
	}

	return
}
