package repository

import (
	"bankCLI/pkg/models"
	"errors"
)

func (r *Repository) GetCity(name string) (city *models.City, err error) {
	for _, v := range r.Database.Cities {
		if v.Name == name {
			city = v
			return
		}
	}

	return nil, errors.New("not found")
}
