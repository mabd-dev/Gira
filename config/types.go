package config

// Config represents the application configuration
type Config struct {
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
	return e.Message
}
