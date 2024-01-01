package repository

import "bankCLI/pkg/models"

func (repo *Repository) AddTransfer(sender, recipient *models.Account, amount float64) {
	repo.Database.Transfers = append(repo.Database.Transfers, &models.Transfer{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	})
}
