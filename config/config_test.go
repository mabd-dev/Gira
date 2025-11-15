package config

import (
	"os"
	"path/filepath"
	"testing"
)

// setupTestConfig creates a temporary config file for testing
func setupTestConfig(t *testing.T, content string) (string, func()) {
	t.Helper()

	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "gira-config-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	// Create config directory structure
	configDir := filepath.Join(tmpDir, ".config", "gira")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		t.Fatalf("Failed to create config dir: %v", err)
	}

	// Write config file
	configPath := filepath.Join(configDir, "credentials.toml")
	if err := os.WriteFile(configPath, []byte(content), 0600); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Set HOME to temp directory
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)

	// Return cleanup function
	cleanup := func() {
		os.Setenv("HOME", originalHome)
		os.RemoveAll(tmpDir)
	}

	return configPath, cleanup
}

func TestLoad_ValidConfig(t *testing.T) {
	content := `[general]
debug = true

[credentials]
email = "test@example.com"
secret = "test-secret"
domain = "test.atlassian.net"
`
	_, cleanup := setupTestConfig(t, content)
	defer cleanup()

	config, err := Load()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Check general settings
	if !config.Debug {
		t.Error("Expected debug to be true")
	}
	if !config.General.Debug {
		t.Error("Expected General.Debug to be true")
	}

	// Check credentials
	if config.Credentials.Email != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got '%s'", config.Credentials.Email)
	}
	if config.Credentials.Secret != "test-secret" {
		t.Errorf("Expected secret 'test-secret', got '%s'", config.Credentials.Secret)
	}
	if config.Credentials.Domain != "test.atlassian.net" {
		t.Errorf("Expected domain 'test.atlassian.net', got '%s'", config.Credentials.Domain)
	}
}

func TestLoad_DebugFalse(t *testing.T) {
	content := `[general]
debug = false

[credentials]
email = "test@example.com"
secret = "test-secret"
domain = "test.atlassian.net"
`
	_, cleanup := setupTestConfig(t, content)
	defer cleanup()

	config, err := Load()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if config.Debug {
		t.Error("Expected debug to be false")
	}
	if config.General.Debug {
		t.Error("Expected General.Debug to be false")
	}
}

func TestLoad_MissingConfigFile(t *testing.T) {
	// Create temp dir but no config file
	tmpDir, err := os.MkdirTemp("", "gira-config-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	_, err = Load()
	if err == nil {
		t.Error("Expected error for missing config file, got nil")
	}
	if err != nil && !os.IsNotExist(err) {
		// Check if error message mentions config not found
		if err.Error() == "" || len(err.Error()) == 0 {
			t.Error("Expected meaningful error message")
		}
	}
}

func TestLoad_InvalidTOML(t *testing.T) {
	content := `[general]
debug = "this should be a boolean not a string
[credentials
email = "test@example.com"
`
	_, cleanup := setupTestConfig(t, content)
	defer cleanup()

	_, err := Load()
	if err == nil {
		t.Error("Expected error for invalid TOML, got nil")
	}
}

func TestLoad_MissingEmail(t *testing.T) {
	content := `[general]
debug = true

[credentials]
secret = "test-secret"
domain = "test.atlassian.net"
`
	_, cleanup := setupTestConfig(t, content)
	defer cleanup()

	_, err := Load()
	if err == nil {
		t.Error("Expected error for missing email, got nil")
	}
	if err != nil && err.Error() != "email is not set or empty in credentials.toml" {
		t.Errorf("Expected email validation error, got: %v", err)
	}
}

func TestLoad_EmptyEmail(t *testing.T) {
	content := `[general]
debug = true

[credentials]
email = "   "
secret = "test-secret"
domain = "test.atlassian.net"
`
	_, cleanup := setupTestConfig(t, content)
	defer cleanup()

	_, err := Load()
	if err == nil {
		t.Error("Expected error for empty email, got nil")
	}
	if err != nil && err.Error() != "email is not set or empty in credentials.toml" {
		t.Errorf("Expected email validation error, got: %v", err)
	}
}

func TestLoad_MissingSecret(t *testing.T) {
	content := `[general]
debug = true

[credentials]
email = "test@example.com"
domain = "test.atlassian.net"
`
	_, cleanup := setupTestConfig(t, content)
	defer cleanup()

	_, err := Load()
	if err == nil {
		t.Error("Expected error for missing secret, got nil")
	}
	if err != nil && err.Error() != "secret is not set or empty in credentials.toml" {
		t.Errorf("Expected secret validation error, got: %v", err)
	}
}

func TestLoad_EmptySecret(t *testing.T) {
	content := `[general]
debug = false

[credentials]
email = "test@example.com"
secret = ""
domain = "test.atlassian.net"
`
	_, cleanup := setupTestConfig(t, content)
	defer cleanup()

	_, err := Load()
	if err == nil {
		t.Error("Expected error for empty secret, got nil")
	}
	if err != nil && err.Error() != "secret is not set or empty in credentials.toml" {
		t.Errorf("Expected secret validation error, got: %v", err)
	}
}

func TestLoad_MissingDomain(t *testing.T) {
	content := `[general]
debug = true

[credentials]
email = "test@example.com"
secret = "test-secret"
`
	_, cleanup := setupTestConfig(t, content)
	defer cleanup()

	_, err := Load()
	if err == nil {
		t.Error("Expected error for missing domain, got nil")
	}
	if err != nil && err.Error() != "domain is not set or empty in credentials.toml" {
		t.Errorf("Expected domain validation error, got: %v", err)
	}
}

func TestLoad_EmptyDomain(t *testing.T) {
	content := `[general]
debug = true

[credentials]
email = "test@example.com"
secret = "test-secret"
domain = "   "
`
	_, cleanup := setupTestConfig(t, content)
	defer cleanup()

	_, err := Load()
	if err == nil {
		t.Error("Expected error for empty domain, got nil")
	}
	if err != nil && err.Error() != "domain is not set or empty in credentials.toml" {
		t.Errorf("Expected domain validation error, got: %v", err)
	}
}

func TestGetConfigPath(t *testing.T) {
	// Set a temporary HOME
	tmpDir, err := os.MkdirTemp("", "gira-config-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", originalHome)

	path, err := getConfigPath()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	expectedPath := filepath.Join(tmpDir, ".config", "gira", "credentials.toml")
	if path != expectedPath {
		t.Errorf("Expected path '%s', got '%s'", expectedPath, path)
	}
}

func TestValidateConfig_AllValid(t *testing.T) {
	config := &Config{
		General: GeneralConfig{Debug: true},
		Credentials: CredentialsConfig{
			Email:  "test@example.com",
			Secret: "test-secret",
			Domain: "test.atlassian.net",
		},
	}

	err := validateConfig(config)
	if err != nil {
		t.Errorf("Expected no error for valid config, got: %v", err)
	}
}

func TestValidateConfig_AllEmpty(t *testing.T) {
	config := &Config{
		General:     GeneralConfig{Debug: false},
		Credentials: CredentialsConfig{},
	}

	err := validateConfig(config)
	if err == nil {
		t.Error("Expected error for empty credentials, got nil")
	}
}

func TestCredentialsAlias(t *testing.T) {
	// Test that Credentials type alias works
	var cred Credentials
	cred.Email = "test@example.com"
	cred.Secret = "test-secret"
	cred.Domain = "test.atlassian.net"

	if cred.Email != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got '%s'", cred.Email)
	}
}
