// Package config handle user credentials and config fetching/saving
package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

var service string = "gira"

// Config represents the application configuration
type Config struct {
	Debug       bool              `toml:"-"` // Populated from General.Debug for backward compatibility
	General     GeneralConfig     `toml:"general"`
	Credentials CredentialsConfig `toml:"credentials"`
}

// GeneralConfig holds general application settings
type GeneralConfig struct {
	Debug bool `toml:"debug"`
}

// CredentialsConfig holds user credentials
type CredentialsConfig struct {
	Email  string `toml:"email"`
	Secret string `toml:"secret"`
	Domain string `toml:"domain"`
}

// Credentials is an alias for backward compatibility
type Credentials = CredentialsConfig

// Load reads and parses the configuration file from ~/.config/gira/credentials.toml
func Load() (Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return Config{}, fmt.Errorf("failed to get config path: %w", err)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return Config{}, fmt.Errorf("config file not found at %s. Please create it from credentials.toml.example", configPath)
		}
		return Config{}, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := toml.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Validate required fields
	if err := validateConfig(&config); err != nil {
		return Config{}, err
	}

	// Populate Debug field for backward compatibility
	config.Debug = config.General.Debug

	return config, nil
}

// getConfigPath returns the path to the credentials file
func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	return filepath.Join(homeDir, ".config", "gira", "credentials.toml"), nil
}

// validateConfig checks that all required fields are set
func validateConfig(config *Config) error {
	if len(strings.TrimSpace(config.Credentials.Email)) == 0 {
		return errors.New("email is not set or empty in credentials.toml")
	}

	if len(strings.TrimSpace(config.Credentials.Secret)) == 0 {
		return errors.New("secret is not set or empty in credentials.toml")
	}

	if len(strings.TrimSpace(config.Credentials.Domain)) == 0 {
		return errors.New("domain is not set or empty in credentials.toml")
	}

	return nil
}
