package database

import "bankCLI/pkg/models"

type Database struct {
	Accounts  []*models.Account
	Cities    []*models.City
	Transfers []*models.Transfer

	Percent float64
}

func NewDatabase(percent float64) *Database {
	return &Database{
		Accounts:  make([]*models.Account, 0),
		Cities:    make([]*models.City, 0),
		Transfers: make([]*models.Transfer, 0),
		Percent:   percent,
	}
}
