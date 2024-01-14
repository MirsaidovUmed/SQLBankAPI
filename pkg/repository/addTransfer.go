package repository

import (
	"bankCLI/pkg/models"
	"context"
)

func (repo *Repository) AddTransfer(sender *models.Account, recipient *models.Account, amount float64) (err error) {
	_, err = repo.Database.Exec(context.Background(), `
	INSERT INTO transfer(
		sender_id,
		recipient_id,
		amount
	)VALUES(
		$1,
		$2,
		$3);
`, sender.Id, recipient.Id, amount)
	return
}
