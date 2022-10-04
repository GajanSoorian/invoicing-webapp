package repository

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/GajanSoorian/parallax-invoicing/invoice-service/internal/config"
	"github.com/stretchr/testify/assert"
)

// test fixtures
func getDefaultEnv() *config.Config {
	return config.NewDefaultConfig()
}

// Test behavior when DB driver is not supported.
func TestSetupDatabaseDriverNotSupported(t *testing.T) {
	_, err := SetupDatabase(getDefaultEnv())
	assert.NotNil(t, err)
	fmt.Println(err)
}

// Test behavior when DB connect failed due to failed Ping test
func TestSetupDatabaseFailedToConnect(t *testing.T) {
	_, err := SetupDatabase(getDefaultEnv())
	assert.NotNil(t, err)
	fmt.Println(err)
}

// Test building DB connection parameters based on environment variables.
func TestBuildDbConnectionParam(t *testing.T) {

	DbConnectionString := BuildDbConnectionParam(nil)
	// If Environment object if nil, returned connection string should be empty
	assert.Equal(t, "", DbConnectionString)

	DbConnectionString = BuildDbConnectionParam(getDefaultEnv())
	// Check if returned string has connect_timeout set to the default value
	assert.Contains(t, DbConnectionString, "connect_timeout="+strconv.Itoa(config.DEFAULT_TIMEOUT))
}
