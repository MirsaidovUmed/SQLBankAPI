package repository

import (
	"bankCLI/pkg/models"
	"context"
)

func (repo *Repository) AddAccount(name string, balance float64, phone_number, password string, city *models.City) (err error) {
	_, err = repo.Database.Exec(context.Background(), `
		INSERT INTO account(
			name,
			balance,
			phone_number,
			city_id,
			password
		)
		VALUES(
			$1,
			$2,
			$3,
			$4,
			$5
		);
	`, name, balance, phone_number, city.Id, password)

	return
}
