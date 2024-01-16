package repository

import (
	"context"
)

func (repo *Repository) CreateProfitAccount(fullName string, phoneNumber string, address string) (err error) {

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
	`, fullName, phoneNumber, address)

	return
}
