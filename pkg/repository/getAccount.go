package repository

import (
	"bankCLI/pkg/models"
	"errors"
)

func (r *Repository) GetAccount(name string) (account *models.Account, err error) {
	for _, v := range r.Database.Accounts {
		if v.Name == name {
			account = v
			return
		}
	}

	return nil, errors.New("not found")
}
