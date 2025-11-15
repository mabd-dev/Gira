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

// Load reads and parses the configuration file from ~/.config/gira/credentials.toml
// If the file doesn't exist, it creates it with example values
func Load(basePath string) (Config, error) {
	configPath, err := getConfigPath(basePath)
	if err != nil {
		return Config{}, fmt.Errorf("Failed to get config path: %w", err)
	}

	// Check if config file exists, if not create it with example values
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := initializeConfig(configPath); err != nil {
			return Config{}, fmt.Errorf("failed to initialize config: %w", err)
		}
		firstTimeError := FirstTimeError{
			Message: fmt.Sprintf("config file created at %s. Please edit it with your Jira credentials", configPath),
		}
		return Config{}, &firstTimeError
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
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

// initializeConfig creates the config directory and file with example values
func initializeConfig(configPath string) error {
	// Create directory with secure permissions (0700 - only owner can access)
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Example configuration content
	exampleConfig := `# Gira Configuration File
# This file was automatically created. Please update it with your Jira credentials.
#
# Security: This file contains sensitive credentials
#   - Keep this file private (permissions: 600)
#   - Never commit this file to version control

[general]
# Enable debug mode (uses mock API client instead of real Jira API)
debug = true

[credentials]
# Your Jira account email
email = "your.email@example.com"

# Your Jira API token
# Get it from: https://id.atlassian.com/manage-profile/security/api-tokens
secret = "your_jira_api_token"

# Your Jira domain (e.g., "yourcompany" if your URL is yourcompany.atlassian.net)
domain = "your-domain"
`

	// Write file with secure permissions (0600 - only owner can read/write)
	if err := os.WriteFile(configPath, []byte(exampleConfig), 0600); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// GetConfigPath returns the path to the credentials file
func getConfigPath(basePath string) (string, error) {
	return filepath.Join(basePath, ".config", "gira", "credentials.toml"), nil
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

	domain := strings.TrimSpace(config.Credentials.Domain)
	if strings.HasPrefix(domain, "http://") || strings.HasPrefix(domain, "https://") {
		return errors.New("domain should not include http:// or https:// prefix")
	}

	return nil
}
