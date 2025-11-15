// Package config handle user credentials and config fetching/saving
package config

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var service string = "gira"

type Config struct {
	Debug bool
	Credentials
}

type Credentials struct {
	Email  string
	Secret string
	Domain string
}

func Load() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	debug := false
	if value, err := strconv.Atoi(os.Getenv("debug")); err == nil {
		debug = value != 0
	}

	email := os.Getenv("email")
	if len(strings.TrimSpace(email)) == 0 {
		return Config{}, errors.New("email is not set or empty. Set it in .env")
	}

	secret := os.Getenv("secret")
	if len(strings.TrimSpace(secret)) == 0 {
		return Config{}, errors.New("secret is not set or empty. Set it in .env")
	}

	domain := os.Getenv("domain")
	if len(strings.TrimSpace(domain)) == 0 {
		return Config{}, errors.New("domain is not set or empty. Set it in .env")
	}

	return Config{
		Debug: debug,
		Credentials: Credentials{
			Email:  email,
			Secret: secret,
			Domain: domain,
		},
	}, nil
}
