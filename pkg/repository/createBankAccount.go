package repository

import (
	"context"
	"sqlBankCLI/pkg/models"
)

func (repo *Repository) CreateBankAccount(account models.Account) (err error) {

	_, err = repo.Conn.Exec(context.Background(), `
		INSERT INTO account(
			full_name,
			phone_number,
			address
		) VALUES (
			$1,
			$2,
			$3
		)
	`, account.FullName, account.PhoneNumber, account.Address)

	return
}
