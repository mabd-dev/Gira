package config

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var service string = "gira"

type Credentials struct {
	Email  string
	Secret string
	Domain string
}

func Load() (Credentials, error) {
	err := godotenv.Load()
	if err != nil {
		return Credentials{}, err
	}
	email := os.Getenv("email")
	if len(strings.TrimSpace(email)) == 0 {
		return Credentials{}, errors.New("email is not set. Please set it in .env")
	}

	secret := os.Getenv("secret")
	if len(strings.TrimSpace(secret)) == 0 {
		return Credentials{}, errors.New("secret is not set. Please set it in .env")
	}

	domain := os.Getenv("domain")
	if len(strings.TrimSpace(secret)) == 0 {
		return Credentials{}, errors.New("domain is not set. Please set it in .env")
	}

	return Credentials{
		Email:  email,
		Secret: secret,
		Domain: domain,
	}, nil
}
