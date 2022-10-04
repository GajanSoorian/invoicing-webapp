package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Todo: have a test env file locally.

// Test the error handling for wrong file name or path.
func TestWrongFileNameOrPath(t *testing.T) {
	config, err := GetEnvVariables("WrongFile.env")

	assert.Nil(t, config)
	assert.NotNil(t, err)
}

// Test validate the invoice-service.env file which stores the environment variables of invoice-service.
// Test case should pass if the environment variable file is valid.
func TestGetEnvVariables(t *testing.T) {
	currentPath, _ := os.Getwd()
	config, err := GetEnvVariables(currentPath + "/../../invoice-service.env")
	assert.NotNil(t, config)
	assert.Nil(t, err)
}
