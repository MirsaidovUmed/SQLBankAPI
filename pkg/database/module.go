package database

import (
	"bankCLI/pkg/config"
	"context"

	"github.com/jackc/pgx/v5"
)

func NewDatabase(config *config.Config) *pgx.Conn {
	db, err := pgx.Connect(context.Background(), config.PostgresURL)

	if err != nil {
		panic(err)
	}

	return db
}
