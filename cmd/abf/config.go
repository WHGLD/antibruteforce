package main

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	GRPC struct {
		Host string `env:"GRPC_HOST" envDefault:"localhost"`
		Port int    `env:"GRPC_PORT" envDefault:"50051"`
	}
	DB struct {
		Username  string `env:"DB_USERNAME" envDefault:"admin"`
		Password  string `env:"DB_PASSWORD" envDefault:"admin"`
		Name      string `env:"DB_NAME" envDefault:"abf"`
		Host      string `env:"DB_HOST" envDefault:"localhost"`
		Port      int    `env:"DB_PORT" envDefault:"5432"`
		SSLEnable bool   `env:"DB_SSL_ENABLE" envDefault:"false"`
	}
	Redis struct {
		Password string `env:"REDIS_PASSWORD" envDefault:"redis_pw"`
		Host     string `env:"REDIS_HOST" envDefault:"localhost"`
		Port     int    `env:"REDIS_PORT" envDefault:"6379"`
	}
	Limits struct {
		Login    int `env:"LIMIT_LOGIN" envDefault:"10"`
		Password int `env:"LIMIT_PASSWORD" envDefault:"100"`
		IP       int `env:"LIMIT_IP" envDefault:"1000"`
	}
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
