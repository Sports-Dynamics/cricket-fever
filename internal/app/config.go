package app

import (
	"errors"
	"os"
)

type Config struct {
	Port          int
	ValidAPIToken string
}

func NewConfig() (Config, error) {
	port := 10000
	token := os.Getenv("SOV_API_TOKEN")
	if len(token) == 0 {
		return Config{}, errors.New("missing environmental variable sovApiToken")
	}
	return Config{Port: port, ValidAPIToken: token}, nil
}
