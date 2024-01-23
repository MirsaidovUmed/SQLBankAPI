package config

type Config struct {
	ServerAddress       string
	ServerPort          int
	PostgresURL         string
	ComissionOfTransfer float64
}

func NewConfig() *Config {
	return &Config{
		ServerAddress:       "127.0.0.1",
		ServerPort:          8000,
		PostgresURL:         "postgres://postgres:postgres@localhost:5432/bank_cli_sql",
		ComissionOfTransfer: 5.0,
	}
}
