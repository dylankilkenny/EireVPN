package db

import (
	"fmt"

	"github.com/caarlos0/env"
)

type DbConfig struct {
	User     string `env:"PG_USER"`
	Password string `env:"PG_PASSWORD"`
	Database string `env:"PG_DB"`
	Host     string `env:"PG_HOST"`
	Port     string `env:"PG_PORT"`
}

func (c *DbConfig) Load() {
	err := env.Parse(c)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
