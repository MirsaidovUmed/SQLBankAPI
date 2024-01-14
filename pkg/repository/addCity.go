package repository

import (
	"context"
	"errors"
)

func (repo *Repository) AddCity(name, region string) (err error) {
	_, err = repo.GetCity(name)

	if err != nil && err.Error() != "not found" {
		return errors.New("already exists")
	}

	_, err = repo.Database.Exec(context.Background(), `
		INSERT INTO city(
			name, region
		) VALUES (
			$1, $2
		)
	`, name, region)

	return err
}
