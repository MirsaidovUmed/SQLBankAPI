package repository

import "bankCLI/pkg/models"

func (repo *Repository) ChangeAccountsBalance(account *models.Account, newBalance float64) {
	account.Balance = newBalance
}
