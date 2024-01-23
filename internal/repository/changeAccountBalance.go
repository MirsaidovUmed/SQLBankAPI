package repository

import (
	"context"
	"sqlBankCLI/internal/models"
)

func (repo *Repository) ChangeAccountBalance(account models.Account) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
	UPDATE account
	SET balance = $1
	WHERE id = $2`, account.Balance, account.Id)
	if err != nil {
		return err
	}
	return
}
