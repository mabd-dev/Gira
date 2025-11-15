package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad_ValidConfig(t *testing.T) {
	// Clean up environment variables
	defer os.Unsetenv("email")
	defer os.Unsetenv("secret")
	defer os.Unsetenv("domain")
	defer os.Unsetenv("debug")

	// Create a temporary directory for test
	tempDir := t.TempDir()
	envFile := filepath.Join(tempDir, ".env")

	// Write a valid .env file
	envContent := `email=test@example.com
secret=test-secret-token
domain=test-domain
debug=1`
	err := os.WriteFile(envFile, []byte(envContent), 0644)
	require.NoError(t, err)

	// Change to temp directory
	originalDir, _ := os.Getwd()
	err = os.Chdir(tempDir)
	require.NoError(t, err)
	defer os.Chdir(originalDir)

	// Test Load
	cfg, err := Load()
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", cfg.Email)
	assert.Equal(t, "test-secret-token", cfg.Secret)
	assert.Equal(t, "test-domain", cfg.Domain)
	assert.True(t, cfg.Debug)
}

func TestLoad_DebugFalse(t *testing.T) {
	defer os.Unsetenv("email")
	defer os.Unsetenv("secret")
	defer os.Unsetenv("domain")
	defer os.Unsetenv("debug")

	tempDir := t.TempDir()
	envFile := filepath.Join(tempDir, ".env")

	envContent := `email=test@example.com
secret=test-secret-token
domain=test-domain
debug=0`
	err := os.WriteFile(envFile, []byte(envContent), 0644)
	require.NoError(t, err)

	originalDir, _ := os.Getwd()
	err = os.Chdir(tempDir)
	require.NoError(t, err)
	defer os.Chdir(originalDir)

	cfg, err := Load()
	assert.NoError(t, err)
	assert.False(t, cfg.Debug)
}

func TestLoad_DebugNotSet(t *testing.T) {
	defer os.Unsetenv("email")
	defer os.Unsetenv("secret")
	defer os.Unsetenv("domain")
	defer os.Unsetenv("debug")

	tempDir := t.TempDir()
	envFile := filepath.Join(tempDir, ".env")

	envContent := `email=test@example.com
secret=test-secret-token
domain=test-domain`
	err := os.WriteFile(envFile, []byte(envContent), 0644)
	require.NoError(t, err)

	originalDir, _ := os.Getwd()
	err = os.Chdir(tempDir)
	require.NoError(t, err)
	defer os.Chdir(originalDir)

	cfg, err := Load()
	assert.NoError(t, err)
	assert.False(t, cfg.Debug)
}

func TestLoad_MissingEnvFile(t *testing.T) {
	defer os.Unsetenv("email")
	defer os.Unsetenv("secret")
	defer os.Unsetenv("domain")
	defer os.Unsetenv("debug")

	tempDir := t.TempDir()
	originalDir, _ := os.Getwd()
	err := os.Chdir(tempDir)
	require.NoError(t, err)
	defer os.Chdir(originalDir)

	_, err = Load()
	assert.Error(t, err)
}

func TestLoad_MissingEmail(t *testing.T) {
	defer os.Unsetenv("email")
	defer os.Unsetenv("secret")
	defer os.Unsetenv("domain")
	defer os.Unsetenv("debug")

	tempDir := t.TempDir()
	envFile := filepath.Join(tempDir, ".env")

	envContent := `secret=test-secret-token
domain=test-domain`
	err := os.WriteFile(envFile, []byte(envContent), 0644)
	require.NoError(t, err)

	originalDir, _ := os.Getwd()
	err = os.Chdir(tempDir)
	require.NoError(t, err)
	defer os.Chdir(originalDir)

	_, err = Load()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "email")
}

func TestLoad_EmptyEmail(t *testing.T) {
	defer os.Unsetenv("email")
	defer os.Unsetenv("secret")
	defer os.Unsetenv("domain")
	defer os.Unsetenv("debug")

	tempDir := t.TempDir()
	envFile := filepath.Join(tempDir, ".env")

	envContent := `email=
secret=test-secret-token
domain=test-domain`
	err := os.WriteFile(envFile, []byte(envContent), 0644)
	require.NoError(t, err)

	originalDir, _ := os.Getwd()
	err = os.Chdir(tempDir)
	require.NoError(t, err)
	defer os.Chdir(originalDir)

	_, err = Load()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "email")
}

func TestLoad_MissingSecret(t *testing.T) {
	defer os.Unsetenv("email")
	defer os.Unsetenv("secret")
	defer os.Unsetenv("domain")
	defer os.Unsetenv("debug")

	tempDir := t.TempDir()
	envFile := filepath.Join(tempDir, ".env")

	envContent := `email=test@example.com
domain=test-domain`
	err := os.WriteFile(envFile, []byte(envContent), 0644)
	require.NoError(t, err)

	originalDir, _ := os.Getwd()
	err = os.Chdir(tempDir)
	require.NoError(t, err)
	defer os.Chdir(originalDir)

	_, err = Load()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "secret")
}

func TestLoad_EmptySecret(t *testing.T) {
	defer os.Unsetenv("email")
	defer os.Unsetenv("secret")
	defer os.Unsetenv("domain")
	defer os.Unsetenv("debug")

	tempDir := t.TempDir()
	envFile := filepath.Join(tempDir, ".env")

	envContent := `email=test@example.com
secret=
domain=test-domain`
	err := os.WriteFile(envFile, []byte(envContent), 0644)
	require.NoError(t, err)

	originalDir, _ := os.Getwd()
	err = os.Chdir(tempDir)
	require.NoError(t, err)
	defer os.Chdir(originalDir)

	_, err = Load()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "secret")
}

func TestLoad_MissingDomain(t *testing.T) {
	defer os.Unsetenv("email")
	defer os.Unsetenv("secret")
	defer os.Unsetenv("domain")
	defer os.Unsetenv("debug")

	tempDir := t.TempDir()
	envFile := filepath.Join(tempDir, ".env")

	envContent := `email=test@example.com
secret=test-secret-token`
	err := os.WriteFile(envFile, []byte(envContent), 0644)
	require.NoError(t, err)

	originalDir, _ := os.Getwd()
	err = os.Chdir(tempDir)
	require.NoError(t, err)
	defer os.Chdir(originalDir)

	_, err = Load()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "domain")
}

func TestLoad_EmptyDomain(t *testing.T) {
	defer os.Unsetenv("email")
	defer os.Unsetenv("secret")
	defer os.Unsetenv("domain")
	defer os.Unsetenv("debug")

	tempDir := t.TempDir()
	envFile := filepath.Join(tempDir, ".env")

	envContent := `email=test@example.com
secret=test-secret-token
domain=   `
	err := os.WriteFile(envFile, []byte(envContent), 0644)
	require.NoError(t, err)

	originalDir, _ := os.Getwd()
	err = os.Chdir(tempDir)
	require.NoError(t, err)
	defer os.Chdir(originalDir)

	_, err = Load()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "domain")
}
