package database

import (
	"context"
	"sqlBankCLI/pkg/config"

	"github.com/jackc/pgx/v5"
)

func NewDatabase(config *config.Config) *pgx.Conn {
	con, err := pgx.Connect(context.Background(), config.PostgresURL)

	if err != nil {
		panic(err)
	}

	return con
}
