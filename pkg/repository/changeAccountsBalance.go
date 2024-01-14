package repository

import (
	"bankCLI/pkg/models"
	"context"
)

func (repo *Repository) ChangeAccountsBalance(account *models.Account) (err error) {
	//account.Balance = newBalance
	_, err = repo.Database.Exec(context.Background(), `
	UPDATE account
	SET balance = $1
	WHERE id = $2`, account.Balance, account.Id)
	if err != nil {
		return err
	}
	return
}
