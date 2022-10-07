package repository

import (
	"fmt"
	"testing"

	"github.com/GajanSoorian/parallax-invoicing/invoice-service/internal/config"
	"github.com/stretchr/testify/assert"
)

// test fixtures
func getDefaultEnv() *config.Config {
	return config.NewDefaultConfig()
}

// Checks if Default DB connection is successful. If this fails, check NewDefaultConfig which stores development
// DB setup information.
func TestDefaultDevDbSetup(t *testing.T) {
	db, err := SetupDbConnection(getDefaultEnv())
	assert.Nil(t, err)
	assert.Nil(t, checkDbConnection(db))
}

// Test behavior when DB driver is not supported.
func TestSetupDatabaseDriverNotSupported(t *testing.T) {
	wrongEnv := getDefaultEnv()
	wrongEnv.RepoDriver = "wrongDriver"
	_, err := SetupDbConnection(wrongEnv)
	assert.ErrorContains(t, err, "sql: unknown driver")
}

// Test behavior when DB connect failed unsupported Environment file format.
func TestSetupDatabaseFailedToConnect(t *testing.T) {
	_, err := SetupDbConnection(&config.Config{})
	assert.NotNil(t, err)
	fmt.Println(err)
}
