package main

import (
	"github.com/caarlos0/env/v6"
)

// Config confuguration struct.
type Config struct {
	DB struct {
		Username  string `env:"DB_USERNAME" envDefault:"admin"`
		Password  string `env:"DB_PASSWORD" envDefault:"admin"`
		Name      string `env:"DB_NAME" envDefault:"abf"`
		Host      string `env:"DB_HOST" envDefault:"localhost"`
		Port      int    `env:"DB_PORT" envDefault:"5432"`
		SSLEnable bool   `env:"DB_SSL_ENABLE" envDefault:"false"`
	}
}

// NewConfig returns configuration.
func NewConfig() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
