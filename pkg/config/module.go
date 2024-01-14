package config

type Config struct {
	PostgresURL         string  `json:"postgres_url"`
	ComissionOfTransfer float64 `json:"comission_of_transfer"`
}

func NewConfig() *Config {
	return &Config{
		PostgresURL:         "postgres://postgres:postgres@localhost:5432/bankcli",
		ComissionOfTransfer: 10.0,
	}
}
