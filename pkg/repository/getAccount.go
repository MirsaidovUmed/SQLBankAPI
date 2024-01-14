package repository

import (
	"bankCLI/pkg/models"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetAccount(name string) (account *models.Account, err error) {
	account = &models.Account{}
	row := r.Database.QueryRow(context.Background(), "SELECT id, balance from account where name = $1", name)

	row.Scan(&account.Id, &account.Balance)

	if err != nil {
		if err == pgx.ErrNoRows {
			return account, errors.New("ПОЛЬЗОВАТЕЛЬ НЕ НАЙДЕН!! ")
		}
	}

	return account, err
}
