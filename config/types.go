package config

import "fmt"

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

type FirstTimeError struct {
	Message string
}

func (e *FirstTimeError) Error() string {
	return fmt.Sprintln(e.Message)
}
