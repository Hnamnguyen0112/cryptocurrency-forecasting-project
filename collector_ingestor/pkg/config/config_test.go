package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	envDir := "collector_ingestor"
	err := os.MkdirAll(envDir, 0755)
	assert.NoError(t, err, "Error creating directory")

	envFilePath := filepath.Join(envDir, ".env")
	envContent := "TEST_KEY=TEST_VALUE\n"
	err = os.WriteFile(envFilePath, []byte(envContent), 0644)
	assert.NoError(t, err, "Error creating temporary .env file")
	defer os.Remove(envFilePath)
	defer os.Remove(envDir)

	key := "TEST_KEY"
	expectedValue := "TEST_VALUE"
	actualValue := Config(key)
	assert.Equal(t, expectedValue, actualValue, "The value should match the expected value")

	nonExistentKey := "NON_EXISTENT_KEY"
	actualValue = Config(nonExistentKey)
	assert.Equal(t, "", actualValue, "The value should be an empty string for a non-existent key")

	os.Remove(envFilePath)
	os.Unsetenv("TEST_KEY")
	actualValue = Config(key)
	assert.Equal(
		t,
		"",
		actualValue,
		"The value should be an empty string when .env file is missing",
	)
}
