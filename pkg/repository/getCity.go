package repository

import (
	"bankCLI/pkg/models"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetCity(name string) (city *models.City, err error) {
	city = &models.City{}
	row := r.Database.QueryRow(context.Background(), `
		SELECT 
			id, name, region
		FROM city
		WHERE name = $1
	`, name)

	err = row.Scan(&city.Id, &city.Name, &city.Region)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("not found")
		}
	}

	return city, err
}
