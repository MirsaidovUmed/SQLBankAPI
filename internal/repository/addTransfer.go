package repository

import (
	"context"
	"sqlBankCLI/internal/models"
)

func (repo *Repository) AddTransfer(sender *models.Account, recipient *models.Account, amount float64) (err error) {
	_, err = repo.Conn.Exec(context.Background(), `
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
