package config

type Config struct {
	PostgresURL         string
	ComissionOfTransfer float64
}

func NewConfig() *Config {
	return &Config{
		PostgresURL:         "postgres://postgres:postgres@localhost:5432/bank_cli_sql",
		ComissionOfTransfer: 5.0,
	}
}
