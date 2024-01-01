package repository

import (
	"bankCLI/pkg/models"
	"errors"
)

func (repo *Repository) AddCity(name, region string) (err error) {
	_, err = repo.GetCity(name)

	if err.Error() != "not found" {
		return errors.New("already exists")
	}

	repo.Database.Cities = append(repo.Database.Cities, &models.City{
		Name:   name,
		Region: region,
	})

	return

}
